/*
Copyright © 2022 xiexianbin
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aliyun

import (
	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
)

type CDNClient struct {
	client *cdn20180510.Client
}

func NewCDNClient(accessKeyId string, accessKeySecret string) (*CDNClient, error) {
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("cdn.aliyuncs.com")
	_result := &cdn20180510.Client{}
	_result, err := cdn20180510.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &CDNClient{
		client: _result,
	}, nil
}

func (c *CDNClient) LogsRequest(domain, startTime, endTime string) (*cdn20180510.DescribeCdnDomainLogsResponse, error) {

	describeCdnDomainLogsRequest := &cdn20180510.DescribeCdnDomainLogsRequest{
		DomainName: tea.String(domain),
		StartTime:  tea.String(startTime), // "2022-05-16T12:00:00Z"
		EndTime:    tea.String(endTime),   // "2022-05-16T13:00:00Z"
	}
	describeCdnDomainLogsResponse, err := c.client.DescribeCdnDomainLogs(describeCdnDomainLogsRequest)
	if err != nil {
		return nil, err
	}

	return describeCdnDomainLogsResponse, err
}
