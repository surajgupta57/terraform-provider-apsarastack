package cms

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

// DeleteContactGroup invokes the cms.DeleteContactGroup API synchronously
func (client *Client) DeleteContactGroup(request *DeleteContactGroupRequest) (response *DeleteContactGroupResponse, err error) {
	response = CreateDeleteContactGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteContactGroupWithChan invokes the cms.DeleteContactGroup API asynchronously
func (client *Client) DeleteContactGroupWithChan(request *DeleteContactGroupRequest) (<-chan *DeleteContactGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteContactGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteContactGroup(request)
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

// DeleteContactGroupWithCallback invokes the cms.DeleteContactGroup API asynchronously
func (client *Client) DeleteContactGroupWithCallback(request *DeleteContactGroupRequest, callback func(response *DeleteContactGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteContactGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteContactGroup(request)
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

// DeleteContactGroupRequest is the request struct for api DeleteContactGroup
type DeleteContactGroupRequest struct {
	*requests.RpcRequest
	ContactGroupName string `position:"Query" name:"ContactGroupName"`
}

// DeleteContactGroupResponse is the response struct for api DeleteContactGroup
type DeleteContactGroupResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteContactGroupRequest creates a request to invoke DeleteContactGroup API
func CreateDeleteContactGroupRequest() (request *DeleteContactGroupRequest) {
	request = &DeleteContactGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteContactGroup", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteContactGroupResponse creates a response to parse from DeleteContactGroup response
func CreateDeleteContactGroupResponse() (response *DeleteContactGroupResponse) {
	response = &DeleteContactGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
