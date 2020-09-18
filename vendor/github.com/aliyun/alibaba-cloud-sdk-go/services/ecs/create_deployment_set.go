package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateDeploymentSet invokes the ecs.CreateDeploymentSet API synchronously
// api document: https://help.aliyun.com/api/ecs/createdeploymentset.html
func (client *Client) CreateDeploymentSet(request *CreateDeploymentSetRequest) (response *CreateDeploymentSetResponse, err error) {
	response = CreateCreateDeploymentSetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDeploymentSetWithChan invokes the ecs.CreateDeploymentSet API asynchronously
// api document: https://help.aliyun.com/api/ecs/createdeploymentset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDeploymentSetWithChan(request *CreateDeploymentSetRequest) (<-chan *CreateDeploymentSetResponse, <-chan error) {
	responseChan := make(chan *CreateDeploymentSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDeploymentSet(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateDeploymentSetWithCallback invokes the ecs.CreateDeploymentSet API asynchronously
// api document: https://help.aliyun.com/api/ecs/createdeploymentset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDeploymentSetWithCallback(request *CreateDeploymentSetRequest, callback func(response *CreateDeploymentSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDeploymentSetResponse
		var err error
		defer close(result)
		response, err = client.CreateDeploymentSet(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateDeploymentSetRequest is the request struct for api CreateDeploymentSet
type CreateDeploymentSetRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                      string           `position:"Query" name:"ClientToken"`
	Description                      string           `position:"Query" name:"Description"`
	GroupCount                       requests.Integer `position:"Query" name:"GroupCount"`
	ResourceOwnerAccount             string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                     string           `position:"Query" name:"OwnerAccount"`
	DeploymentSetName                string           `position:"Query" name:"DeploymentSetName"`
	OwnerId                          requests.Integer `position:"Query" name:"OwnerId"`
	OnUnableToRedeployFailedInstance string           `position:"Query" name:"OnUnableToRedeployFailedInstance"`
	Granularity                      string           `position:"Query" name:"Granularity"`
	Domain                           string           `position:"Query" name:"Domain"`
	Strategy                         string           `position:"Query" name:"Strategy"`
}

// CreateDeploymentSetResponse is the response struct for api CreateDeploymentSet
type CreateDeploymentSetResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	DeploymentSetId string `json:"DeploymentSetId" xml:"DeploymentSetId"`
}

// CreateCreateDeploymentSetRequest creates a request to invoke CreateDeploymentSet API
func CreateCreateDeploymentSetRequest() (request *CreateDeploymentSetRequest) {
	request = &CreateDeploymentSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateDeploymentSet", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDeploymentSetResponse creates a response to parse from CreateDeploymentSet response
func CreateCreateDeploymentSetResponse() (response *CreateDeploymentSetResponse) {
	response = &CreateDeploymentSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
