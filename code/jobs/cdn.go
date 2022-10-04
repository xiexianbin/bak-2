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

package jobs

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/xiexianbin/fc-aliyun-cdn-404/code/models"
	"github.com/xiexianbin/fc-aliyun-cdn-404/code/utils"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/xiexianbin/fc-aliyun-cdn-404/code/aliyun"
)

func DoParseCDNLog(day string) error {
	dayTime, err := time.Parse("2006-01-02", day)
	if err != nil {
		return errors.New(fmt.Sprintf("day [%s] layout not match `2006-01-02`", day))
	}

	lockKey := fmt.Sprintf("hasTask-%s", day)
	hasTask, _ := utils.BCache.Get(nil, lockKey)
	if hasTask == true {
		return errors.New(fmt.Sprintf("a task [%s] is already exist", day))
	} else {
		_ = utils.BCache.Put(nil, lockKey, true, 10*time.Minute)
		//go jobs.ParseCDNLog(dayTime)
		err := ParseCDNLog(dayTime)
		if err == nil {
			_ = utils.BCache.Delete(nil, lockKey)
		}
		return err
	}
}

func ParseCDNLog(dayTime time.Time) error {
	var err error
	aliyunCDNClient, err := aliyun.NewCDNClient(
		beego.AppConfig.DefaultString("aliyun::access_key_id", ""),
		beego.AppConfig.DefaultString("aliyun::access_key_secret", ""))
	if err != nil {
		logs.Info("create aliyun cdn client err:", err.Error())
	}

	domain := beego.AppConfig.DefaultString("aliyun::cdndomain", "www.xiexianbin.cn")
	startTime := dayTime.Format("2006-01-02T15:04:05Z")
	sub, _ := time.ParseDuration("23h59m59s")
	endTime := dayTime.Add(sub).Format("2006-01-02T15:04:05Z")
	logsRequest, err := aliyunCDNClient.LogsRequest(domain, startTime, endTime)
	if err != nil {
		logs.Debug("get aliyun CDN LogsRequest error %s", err.Error())
		return err
	}

	domainLogDetail := logsRequest.Body.DomainLogDetails.DomainLogDetail
	if len(domainLogDetail) == 0 {
		logs.Debug("aliyun CDN domainLogDetail is ", domainLogDetail)
		return nil
	}

	// create logs dir
	logDir := fmt.Sprintf("%s/%s", utils.CND_LOG_DIR, dayTime.Format("2006-01-02"))
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		logs.Warning("create dir %s error %s", logDir, err.Error())
		return err
	}
	// defer os.RemoveAll(logDir)

	for _, logInfoDetail := range domainLogDetail[0].LogInfos.LogInfoDetail {
		destFile := path.Join(logDir, *logInfoDetail.LogName)
		logPath := logInfoDetail.LogPath
		// download log
		if err := utils.DownloadFile(*logPath, destFile); err != nil {
			logs.Warning("download %s error %s", logPath, err.Error())
			continue
		}

		// read gzip log
		context, err := utils.ReadGzipFile(destFile)
		if err != nil {
			logs.Warning("read gzip file %s error %s", destFile, err.Error())
			continue
		}

		// parse log
		// [17/May/2022:00:11:28 +0800] 171.8.172.154 - 478 "-" "GET https://www.xiexianbin.cn/docker/2017-03-03-docker-container-ecosystem/" 302 404 7315 MISS "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Mobile Safari/537.36" "text/html; charset=utf-8"
		for _, line := range strings.Split(context, "\n") {
			if strings.Contains(line, "bingbot/2.0;") == true || strings.Contains(line, "DotBot/1.2;") == true ||
				strings.Contains(line, "Googlebot/2.1;") || strings.Contains(line, "Sogou web spider/4.0") || strings.Contains(line, "AhrefsBot/7.0;") ||
				strings.Contains(line, "UptimeRobot/2.0;") || strings.Contains(line, "Baiduspider") {
				continue
			}
			s := strings.Split(line, " ")
			if len(s) > 11 && s[8] == "302" && s[9] != "413" {
				ip := s[2]
				method := strings.Replace(s[6], "\"", "", -1)
				url := strings.Replace(s[7], "\"", "", -1)
				code1 := s[8]
				code2 := s[9]
				status := s[11]
				timeStr := strings.Replace(s[0], "[", "", -1)

				loc := time.FixedZone("UTC+8", +8*60*60)
				logTime, _ := time.ParseInLocation("02/Jan/2006:15:04:05", timeStr, loc)

				_, err = models.NewPage404(ip, method, url, code1, code2, status, logTime)
				if err != nil {
					logs.Warning("save now 404 line %s error %s", line, err.Error())
				}
			}
		}

		// 单个用户的调用频率为：100次/秒
		time.Sleep(600 * time.Microsecond)
	}

	return nil
}
