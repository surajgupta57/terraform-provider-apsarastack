package apsarastack

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccApsaraStackYundunBastionhostInstanceDataSource_basic(t *testing.T) {
	rand := acctest.RandInt()
	resourceId := "data.apsarastack_yundun_bastionhost_instances.default"

	testAccConfig := dataSourceTestAccConfigFunc(resourceId, fmt.Sprintf("tf_testAcc%d", rand),
		dataSourceYundunBastionhostInstanceConfigDependency)

	idsConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"ids": []string{"${apsarastack_yundun_bastionhost_instance.default.id}"},
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"ids": []string{"${apsarastack_yundun_bastionhost_instance.default.id}-fake"},
		}),
	}

	nameRegexConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"description_regex": "${apsarastack_yundun_bastionhost_instance.default.description}",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"description_regex": "${apsarastack_yundun_bastionhost_instance.default.description}-fake",
		}),
	}

	tagsConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"ids": []string{"${apsarastack_yundun_bastionhost_instance.default.id}"},
			"tags": map[string]interface{}{
				"Created": "TF",
			},
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"ids": []string{"${apsarastack_yundun_bastionhost_instance.default.id}-fake"},
			"tags": map[string]interface{}{
				"Created": "TF-fake",
			},
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"description_regex": "${apsarastack_yundun_bastionhost_instance.default.description}",
			"ids":               []string{"${apsarastack_yundun_bastionhost_instance.default.id}"},
			"tags": map[string]interface{}{
				"For": "acceptance test",
			},
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"description_regex": "${apsarastack_yundun_bastionhost_instance.default.description}-fake",
			"ids":               []string{"${apsarastack_yundun_bastionhost_instance.default.id}-fake"},
			"tags": map[string]interface{}{
				"For": "acceptance test-fake",
			},
		}),
	}

	var existYundunBastionhostInstanceMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":                             "1",
			"descriptions.#":                    "1",
			"ids.0":                             CHECKSET,
			"descriptions.0":                    fmt.Sprintf("tf_testAcc%d", rand),
			"instances.#":                       "1",
			"instances.0.description":           fmt.Sprintf("tf_testAcc%d", rand),
			"instances.0.license_code":          "bhah_ent_50_asset",
			"instances.0.user_vswitch_id":       CHECKSET,
			"instances.0.public_network_access": "true",
			"instances.0.private_domain":        CHECKSET,
			"instances.0.instance_status":       CHECKSET,
			"instances.0.security_group_ids.#":  "1",
		}
	}
	var fakeYundunBastionhostInstanceMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":          "0",
			"descriptions.#": "0",
		}
	}
	var yundunBastionhostInstanceCheckInfo = dataSourceAttr{
		resourceId:   "data.apsarastack_yundun_bastionhost_instances.default",
		existMapFunc: existYundunBastionhostInstanceMapFunc,
		fakeMapFunc:  fakeYundunBastionhostInstanceMapFunc,
	}

	preCheck := func() {
		testAccPreCheckWithRegions(t, true, connectivity.YundunBastionhostSupportedRegions)
		testAccPreCheckWithAccountSiteType(t, DomesticSite)
	}

	yundunBastionhostInstanceCheckInfo.dataSourceTestCheckWithPreCheck(t, rand, preCheck, nameRegexConf, idsConf, tagsConf, allConf)

}

func dataSourceYundunBastionhostInstanceConfigDependency(description string) string {
	return fmt.Sprintf(
		`data "apsarastack_zones" "default" {
				  available_resource_creation = "VSwitch"
				}
				
				variable "name" {
				  default = "%s"
				}
				
				resource "apsarastack_vpc" "default" {
				  name       = "${var.name}"
				  cidr_block = "172.16.0.0/12"
				}
				
				resource "apsarastack_vswitch" "default" {
				  vpc_id            = "${apsarastack_vpc.default.id}"
				  cidr_block        = "172.16.0.0/21"
				  availability_zone = "${data.apsarastack_zones.default.zones.0.id}"
				  name              = "${var.name}"
				}
				
				resource "apsarastack_security_group" "default" {
				  name   = "${var.name}"
				  vpc_id = "${apsarastack_vpc.default.id}"
				}
				
				provider "apsarastack" {
				  endpoints {
					bssopenapi = "business.aliyuncs.com"
				  }
				}
				
				resource "apsarastack_yundun_bastionhost_instance" "default" {
				  description        = "${var.name}"
				  license_code       = "bhah_ent_50_asset"
				  period             = "1"
				  vswitch_id         = "${apsarastack_vswitch.default.id}"
				  security_group_ids = ["${apsarastack_security_group.default.id}"]
				  tags 				 = {
						Created = "TF"
						For 	= "acceptance test"
				  }
				}`, description)
}
