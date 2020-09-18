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

// DescribeLiveStreamOnlineUserNum invokes the cdn.DescribeLiveStreamOnlineUserNum API synchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamonlineusernum.html
func (client *Client) DescribeLiveStreamOnlineUserNum(request *DescribeLiveStreamOnlineUserNumRequest) (response *DescribeLiveStreamOnlineUserNumResponse, err error) {
	response = CreateDescribeLiveStreamOnlineUserNumResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamOnlineUserNumWithChan invokes the cdn.DescribeLiveStreamOnlineUserNum API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamonlineusernum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamOnlineUserNumWithChan(request *DescribeLiveStreamOnlineUserNumRequest) (<-chan *DescribeLiveStreamOnlineUserNumResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamOnlineUserNumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamOnlineUserNum(request)
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

// DescribeLiveStreamOnlineUserNumWithCallback invokes the cdn.DescribeLiveStreamOnlineUserNum API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamonlineusernum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamOnlineUserNumWithCallback(request *DescribeLiveStreamOnlineUserNumRequest, callback func(response *DescribeLiveStreamOnlineUserNumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamOnlineUserNumResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamOnlineUserNum(request)
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

// DescribeLiveStreamOnlineUserNumRequest is the request struct for api DescribeLiveStreamOnlineUserNum
type DescribeLiveStreamOnlineUserNumRequest struct {
	*requests.RpcRequest
	StartTime     string           `position:"Query" name:"StartTime"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	HlsSwitch     string           `position:"Query" name:"HlsSwitch"`
	StreamName    string           `position:"Query" name:"StreamName"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamOnlineUserNumResponse is the response struct for api DescribeLiveStreamOnlineUserNum
type DescribeLiveStreamOnlineUserNumResponse struct {
	*responses.BaseResponse
	RequestId       string         `json:"RequestId" xml:"RequestId"`
	TotalUserNumber int64          `json:"TotalUserNumber" xml:"TotalUserNumber"`
	OnlineUserInfo  OnlineUserInfo `json:"OnlineUserInfo" xml:"OnlineUserInfo"`
}

// CreateDescribeLiveStreamOnlineUserNumRequest creates a request to invoke DescribeLiveStreamOnlineUserNum API
func CreateDescribeLiveStreamOnlineUserNumRequest() (request *DescribeLiveStreamOnlineUserNumRequest) {
	request = &DescribeLiveStreamOnlineUserNumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamOnlineUserNum", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveStreamOnlineUserNumResponse creates a response to parse from DescribeLiveStreamOnlineUserNum response
func CreateDescribeLiveStreamOnlineUserNumResponse() (response *DescribeLiveStreamOnlineUserNumResponse) {
	response = &DescribeLiveStreamOnlineUserNumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
