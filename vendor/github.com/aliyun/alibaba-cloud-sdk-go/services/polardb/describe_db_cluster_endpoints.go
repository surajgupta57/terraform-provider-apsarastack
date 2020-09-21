package polardb

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

// DescribeDBClusterEndpoints invokes the polardb.DescribeDBClusterEndpoints API synchronously
// api document: https://help.aliyun.com/api/polardb/describedbclusterendpoints.html
func (client *Client) DescribeDBClusterEndpoints(request *DescribeDBClusterEndpointsRequest) (response *DescribeDBClusterEndpointsResponse, err error) {
	response = CreateDescribeDBClusterEndpointsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterEndpointsWithChan invokes the polardb.DescribeDBClusterEndpoints API asynchronously
// api document: https://help.aliyun.com/api/polardb/describedbclusterendpoints.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBClusterEndpointsWithChan(request *DescribeDBClusterEndpointsRequest) (<-chan *DescribeDBClusterEndpointsResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterEndpointsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterEndpoints(request)
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

// DescribeDBClusterEndpointsWithCallback invokes the polardb.DescribeDBClusterEndpoints API asynchronously
// api document: https://help.aliyun.com/api/polardb/describedbclusterendpoints.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBClusterEndpointsWithCallback(request *DescribeDBClusterEndpointsRequest, callback func(response *DescribeDBClusterEndpointsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterEndpointsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterEndpoints(request)
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

// DescribeDBClusterEndpointsRequest is the request struct for api DescribeDBClusterEndpoints
type DescribeDBClusterEndpointsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBEndpointId         string           `position:"Query" name:"DBEndpointId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterEndpointsResponse is the response struct for api DescribeDBClusterEndpoints
type DescribeDBClusterEndpointsResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Items     []DBEndpoint `json:"Items" xml:"Items"`
}

// CreateDescribeDBClusterEndpointsRequest creates a request to invoke DescribeDBClusterEndpoints API
func CreateDescribeDBClusterEndpointsRequest() (request *DescribeDBClusterEndpointsRequest) {
	request = &DescribeDBClusterEndpointsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterEndpoints", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterEndpointsResponse creates a response to parse from DescribeDBClusterEndpoints response
func CreateDescribeDBClusterEndpointsResponse() (response *DescribeDBClusterEndpointsResponse) {
	response = &DescribeDBClusterEndpointsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}