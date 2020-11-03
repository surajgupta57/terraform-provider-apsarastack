package alidns

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

// AddrPool is a nested struct in alidns response
type AddrPool struct {
	AddrCount           int    `json:"AddrCount" xml:"AddrCount"`
	Name                string `json:"Name" xml:"Name"`
	MinAvailableAddrNum int    `json:"MinAvailableAddrNum" xml:"MinAvailableAddrNum"`
	UpdateTime          string `json:"UpdateTime" xml:"UpdateTime"`
	MonitorConfigId     string `json:"MonitorConfigId" xml:"MonitorConfigId"`
	CreateTimestamp     int64  `json:"CreateTimestamp" xml:"CreateTimestamp"`
	CreateTime          string `json:"CreateTime" xml:"CreateTime"`
	AddrPoolName        string `json:"AddrPoolName" xml:"AddrPoolName"`
	AddrPoolId          string `json:"AddrPoolId" xml:"AddrPoolId"`
	UpdateTimestamp     int64  `json:"UpdateTimestamp" xml:"UpdateTimestamp"`
	Status              string `json:"Status" xml:"Status"`
	MonitorStatus       string `json:"MonitorStatus" xml:"MonitorStatus"`
	Type                string `json:"Type" xml:"Type"`
}
