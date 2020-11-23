package cs

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

// CancelWorkflow invokes the cs.CancelWorkflow API synchronously
func (client *Client) CancelWorkflow(request *CancelWorkflowRequest) (response *CancelWorkflowResponse, err error) {
	response = CreateCancelWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// CancelWorkflowWithChan invokes the cs.CancelWorkflow API asynchronously
func (client *Client) CancelWorkflowWithChan(request *CancelWorkflowRequest) (<-chan *CancelWorkflowResponse, <-chan error) {
	responseChan := make(chan *CancelWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelWorkflow(request)
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

// CancelWorkflowWithCallback invokes the cs.CancelWorkflow API asynchronously
func (client *Client) CancelWorkflowWithCallback(request *CancelWorkflowRequest, callback func(response *CancelWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelWorkflowResponse
		var err error
		defer close(result)
		response, err = client.CancelWorkflow(request)
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

// CancelWorkflowRequest is the request struct for api CancelWorkflow
type CancelWorkflowRequest struct {
	*requests.RoaRequest
	WorkflowName string `position:"Path" name:"workflowName"`
}

// CancelWorkflowResponse is the response struct for api CancelWorkflow
type CancelWorkflowResponse struct {
	*responses.BaseResponse
}

// CreateCancelWorkflowRequest creates a request to invoke CancelWorkflow API
func CreateCancelWorkflowRequest() (request *CancelWorkflowRequest) {
	request = &CancelWorkflowRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "CancelWorkflow", "/gs/workflow/[workflowName]", "", "")
	request.Method = requests.PUT
	return
}

// CreateCancelWorkflowResponse creates a response to parse from CancelWorkflow response
func CreateCancelWorkflowResponse() (response *CancelWorkflowResponse) {
	response = &CancelWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}