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

package controllers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/xiexianbin/fc-aliyun-cdn-404/code/jobs"
)

// TaskController operations for Task
type TaskController struct {
	BaseController
}

// Get ...
// @Title Get
// @Description get Task
// @Param	last		param 	int	true		"last log to get"
// @Success 200 {object} ResponseOK
// @Failure 403 :last is empty
// @router / [get]
func (c *TaskController) Get() {
	var err error
	offset, _ := c.GetInt("offset", 0)
	last, err := c.GetInt("last", 0)
	if err != nil {
		c.Data["json"] = ResponseError(err)
	} else if last == 0 {
		c.Data["json"] = ResponseError(errors.New("last must be greater than 0"))
	} else {
		now := time.Now()
		var days []string
		var errs []error
		for i := offset; i < last; i++ {
			duration, err := time.ParseDuration(fmt.Sprintf("-%dh", i*24))
			day := now.Add(duration).Format("2006-01-02")
			err = jobs.DoParseCDNLog(day)
			if err != nil {
				errs = append(errs, err)
			} else {
				days = append(days, day)
			}
		}

		resp := &Response{
			Status: "ok",
		}
		if len(errs) > 0 {
			resp.Data = errs
		}
		if len(days) > 0 {
			resp.Message = fmt.Sprintf("task [%s] is new created.", strings.Join(days, ", "))
		}

		c.Data["json"] = resp
	}

	_ = c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Task by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} ResponseOK
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TaskController) GetOne() {
	var err error
	day := c.Ctx.Input.Query(":day")
	err = jobs.DoParseCDNLog(day)
	if err != nil {
		c.Data["json"] = ResponseError(err)
	} else {
		c.Data["json"] = ResponseOK(fmt.Sprintf("task [%s] is new created.", day))
	}

	_ = c.ServeJSON()
}
