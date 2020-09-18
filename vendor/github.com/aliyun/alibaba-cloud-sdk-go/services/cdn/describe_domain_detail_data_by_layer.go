package cdn

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

// DescribeDomainDetailDataByLayer invokes the cdn.DescribeDomainDetailDataByLayer API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomaindetaildatabylayer.html
func (client *Client) DescribeDomainDetailDataByLayer(request *DescribeDomainDetailDataByLayerRequest) (response *DescribeDomainDetailDataByLayerResponse, err error) {
	response = CreateDescribeDomainDetailDataByLayerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainDetailDataByLayerWithChan invokes the cdn.DescribeDomainDetailDataByLayer API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomaindetaildatabylayer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainDetailDataByLayerWithChan(request *DescribeDomainDetailDataByLayerRequest) (<-chan *DescribeDomainDetailDataByLayerResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainDetailDataByLayerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainDetailDataByLayer(request)
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

// DescribeDomainDetailDataByLayerWithCallback invokes the cdn.DescribeDomainDetailDataByLayer API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomaindetaildatabylayer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainDetailDataByLayerWithCallback(request *DescribeDomainDetailDataByLayerRequest, callback func(response *DescribeDomainDetailDataByLayerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainDetailDataByLayerResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainDetailDataByLayer(request)
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

// DescribeDomainDetailDataByLayerRequest is the request struct for api DescribeDomainDetailDataByLayer
type DescribeDomainDetailDataByLayerRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	Layer          string           `position:"Query" name:"Layer"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Field          string           `position:"Query" name:"Field"`
}

// DescribeDomainDetailDataByLayerResponse is the response struct for api DescribeDomainDetailDataByLayer
type DescribeDomainDetailDataByLayerResponse struct {
	*responses.BaseResponse
	RequestId string                                `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeDomainDetailDataByLayer `json:"Data" xml:"Data"`
}

// CreateDescribeDomainDetailDataByLayerRequest creates a request to invoke DescribeDomainDetailDataByLayer API
func CreateDescribeDomainDetailDataByLayerRequest() (request *DescribeDomainDetailDataByLayerRequest) {
	request = &DescribeDomainDetailDataByLayerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainDetailDataByLayer", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeDomainDetailDataByLayerResponse creates a response to parse from DescribeDomainDetailDataByLayer response
func CreateDescribeDomainDetailDataByLayerResponse() (response *DescribeDomainDetailDataByLayerResponse) {
	response = &DescribeDomainDetailDataByLayerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
