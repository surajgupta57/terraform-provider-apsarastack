package apsarastack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/yundun_bastionhost"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	BATIONHOST_RELEASE_HANG_MINS  = 5
	BASTIONHOST_WAITING_FOR_START = 600
)

func resourceApsaraStackBastionhostInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceApsaraStackBastionhostInstanceCreate,
		Read:   resourceApsaraStackBastionhostInstanceRead,
		Update: resourceApsaraStackBastionhostInstanceUpdate,
		Delete: resourceApsaraStackBastionhostInstanceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 64),
			},
			"license_code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntInSlice([]int{1, 3, 6, 12, 24, 36}),
				Optional:     true,
				Default:      1,
				ForceNew:     true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_group_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": tagsSchema(),

			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceApsaraStackBastionhostInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	request := buildBastionhostCreateRequest(d, meta)
	raw, err := client.WithBssopenapiClient(func(bssopenapiClient *bssopenapi.Client) (interface{}, error) {
		return bssopenapiClient.CreateInstance(request)
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "apsarastack_yundun_bastionhost_instance", request.GetActionName(), ApsaraStackSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)

	response := raw.(*bssopenapi.CreateInstanceResponse)
	instanceId := response.Data.InstanceId
	if !response.Success {
		return WrapError(Error(response.Message))
	}
	d.SetId(instanceId)

	bastionhostService := bastionhostService{client}

	// check RAM policy
	if err := bastionhostService.ProcessRolePolicy(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	// wait for order complete
	stateConf := BuildStateConf([]string{}, []string{"PENDING"}, d.Timeout(schema.TimeoutCreate), 20*time.Second, bastionhostService.BastionhostInstanceRefreshFunc(d.Id(), []string{"UPGRADING", "UPGRADE_FAILED", "CREATE_FAILED"}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	rawSecurityGroupIds := d.Get("security_group_ids").([]interface{})
	securityGroupIds := make([]string, len(rawSecurityGroupIds))
	for index, rawSecurityGroupId := range rawSecurityGroupIds {
		securityGroupIds[index] = rawSecurityGroupId.(string)
	}
	// start instance
	if err := bastionhostService.StartBastionhostInstance(instanceId, d.Get("vswitch_id").(string), securityGroupIds); err != nil {
		return WrapError(err)
	}
	// wait for pending
	stateConf = BuildStateConf([]string{"PENDING", "CREATING"}, []string{"RUNNING"}, d.Timeout(schema.TimeoutCreate), BASTIONHOST_WAITING_FOR_START*time.Second, bastionhostService.BastionhostInstanceRefreshFunc(d.Id(), []string{"UPGRADING", "UPGRADE_FAILED", "CREATE_FAILED"}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return resourceApsaraStackBastionhostInstanceUpdate(d, meta)
}

func resourceApsaraStackBastionhostInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	BastionhostService := bastionhostService{client}
	instance, err := BastionhostService.DescribeBastionhostInstanceAttribute(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	d.Set("description", instance.Description)
	period, err := computePeriodByUnit(instance.StartTime/1000, instance.ExpireTime/1000, d.Get("period").(int), "Month")
	if err != nil {
		return WrapError(err)
	}
	d.Set("period", period)
	d.Set("license_code", instance.LicenseCode)
	d.Set("region_id", client.RegionId)
	d.Set("vswitch_id", instance.VswitchId)
	sgs := make([]string, 0, len(instance.ReferredSecurityGroups))
	for _, sg := range instance.ReferredSecurityGroups {
		sgs = append(sgs, sg)
	}
	d.Set("security_group_ids", sgs)

	tags, err := BastionhostService.DescribeTags(d.Id(), nil, TagResourceInstance)
	if err != nil {
		return WrapError(err)
	}
	d.Set("tags", BastionhostService.tagsToMap(tags))
	return nil
}

func resourceApsaraStackBastionhostInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	bastionhostService := bastionhostService{client}

	d.Partial(true)
	if d.HasChange("tags") {
		if err := bastionhostService.setInstanceTags(d, TagResourceInstance); err != nil {
			return WrapError(err)
		}
		d.SetPartial("tags")
	}

	if d.HasChange("description") {
		if err := bastionhostService.UpdateBastionhostInstanceDescription(d.Id(), d.Get("description").(string)); err != nil {
			return WrapError(err)
		}
		d.SetPartial("description")
	}

	if d.HasChange("resource_group_id") {
		if err := bastionhostService.UpdateResourceGroup(d.Id(), d.Get("resource_group_id").(string)); err != nil {
			return WrapError(err)
		}
		d.SetPartial("resource_group_id")
	}

	if d.IsNewResource() {
		d.Partial(false)
		return resourceApsaraStackBastionhostInstanceRead(d, meta)
	}

	if d.HasChange("license_code") {
		params := map[string]string{
			"LicenseCode": "license_code",
		}
		if err := bastionhostService.UpdateInstanceSpec(params, d, meta); err != nil {
			return WrapError(err)
		}
		stateConf := BuildStateConf([]string{"UPGRADING"}, []string{"PENDING", "RUNNING"}, d.Timeout(schema.TimeoutUpdate), 20*time.Second, bastionhostService.BastionhostInstanceRefreshFunc(d.Id(), []string{"CREATING", "UPGRADE_FAILED", "CREATE_FAILED"}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
		d.SetPartial("license_code")
	}

	if d.HasChange("security_group_ids") {
		securityGroupIds := d.Get("security_group_ids").([]interface{})
		sgs := make([]string, 0, len(securityGroupIds))
		for _, rawSecurityGroupId := range securityGroupIds {
			sgs = append(sgs, rawSecurityGroupId.(string))
		}
		if err := bastionhostService.UpdateBastionhostSecurityGroups(d.Id(), sgs); err != nil {
			return WrapError(err)
		}
		stateConf := BuildStateConf([]string{"UPGRADING"}, []string{"RUNNING"}, d.Timeout(schema.TimeoutUpdate), 20*time.Second, bastionhostService.BastionhostInstanceRefreshFunc(d.Id(), []string{"CREATING", "UPGRADE_FAILED", "CREATE_FAILED"}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
		d.SetPartial("security_group_ids")
	}

	d.Partial(false)
	// wait for order complete
	return resourceApsaraStackBastionhostInstanceRead(d, meta)
}

func resourceApsaraStackBastionhostInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	bastionhostService := bastionhostService{client}
	request := yundun_bastionhost.CreateRefundInstanceRequest()
	request.Headers = map[string]string{"RegionId": client.RegionId}
	request.QueryParams = map[string]string{"AccessKeySecret": client.SecretKey, "Product": "Yundun-bastionhost", "Department": client.Department, "ResourceGroup": client.ResourceGroup}
	request.InstanceId = d.Id()

	raw, err := bastionhostService.client.WithBastionhostClient(func(BastionhostClient *yundun_bastionhost.Client) (interface{}, error) {
		return BastionhostClient.RefundInstance(request)
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), ApsaraStackSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	// Wait for the release procedure of cloud resource dependencies. Instance can not be fetched through api as soon as release has
	// been invoked, however the resources have not been fully destroyed yet. Therefore, a certain amount time of waiting
	// is quite necessary (conservative estimation cloud be less then 3 minutes)
	time.Sleep(BATIONHOST_RELEASE_HANG_MINS * time.Minute)
	return WrapError(bastionhostService.WaitForYundunBastionhostInstance(d.Id(), Deleted, 0))
}

func buildBastionhostCreateRequest(d *schema.ResourceData, meta interface{}) *bssopenapi.CreateInstanceRequest {
	request := bssopenapi.CreateCreateInstanceRequest()
	request.ProductCode = "bastionhost"
	request.SubscriptionType = "Subscription"
	request.Period = requests.NewInteger(d.Get("period").(int))
	client := meta.(*connectivity.ApsaraStackClient)
	request.Headers = map[string]string{"RegionId": client.RegionId}
	request.QueryParams = map[string]string{"AccessKeySecret": client.SecretKey, "Product": "Yundun-bastionhost", "Department": client.Department, "ResourceGroup": client.ResourceGroup}

	request.Parameter = &[]bssopenapi.CreateInstanceParameter{
		// force to buy vpc version
		{
			Code:  "NetworkType",
			Value: "vpc",
		},
		{
			Code:  "LicenseCode",
			Value: d.Get("license_code").(string),
		},
		{
			Code:  "PlanCode",
			Value: "cloudbastion",
		},
		{
			Code:  "RegionId",
			Value: client.RegionId,
		},
	}
	return request
}
