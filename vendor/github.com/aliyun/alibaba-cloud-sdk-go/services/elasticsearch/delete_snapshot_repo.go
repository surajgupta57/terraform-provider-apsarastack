package elasticsearch

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

// DeleteSnapshotRepo invokes the elasticsearch.DeleteSnapshotRepo API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/deletesnapshotrepo.html
func (client *Client) DeleteSnapshotRepo(request *DeleteSnapshotRepoRequest) (response *DeleteSnapshotRepoResponse, err error) {
	response = CreateDeleteSnapshotRepoResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSnapshotRepoWithChan invokes the elasticsearch.DeleteSnapshotRepo API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/deletesnapshotrepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSnapshotRepoWithChan(request *DeleteSnapshotRepoRequest) (<-chan *DeleteSnapshotRepoResponse, <-chan error) {
	responseChan := make(chan *DeleteSnapshotRepoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSnapshotRepo(request)
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

// DeleteSnapshotRepoWithCallback invokes the elasticsearch.DeleteSnapshotRepo API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/deletesnapshotrepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSnapshotRepoWithCallback(request *DeleteSnapshotRepoRequest, callback func(response *DeleteSnapshotRepoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSnapshotRepoResponse
		var err error
		defer close(result)
		response, err = client.DeleteSnapshotRepo(request)
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

// DeleteSnapshotRepoRequest is the request struct for api DeleteSnapshotRepo
type DeleteSnapshotRepoRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
	RepoPath    string `position:"Query" name:"repoPath"`
}

// DeleteSnapshotRepoResponse is the response struct for api DeleteSnapshotRepo
type DeleteSnapshotRepoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateDeleteSnapshotRepoRequest creates a request to invoke DeleteSnapshotRepo API
func CreateDeleteSnapshotRepoRequest() (request *DeleteSnapshotRepoRequest) {
	request = &DeleteSnapshotRepoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DeleteSnapshotRepo", "/openapi/instances/[InstanceId]/snapshot-repos", "elasticsearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteSnapshotRepoResponse creates a response to parse from DeleteSnapshotRepo response
func CreateDeleteSnapshotRepoResponse() (response *DeleteSnapshotRepoResponse) {
	response = &DeleteSnapshotRepoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
