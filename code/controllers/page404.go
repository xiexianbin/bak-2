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
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/xiexianbin/fc-aliyun-cdn-404/code/forms"
	"github.com/xiexianbin/fc-aliyun-cdn-404/code/models"
	"github.com/xiexianbin/fc-aliyun-cdn-404/code/utils"
)

// Page404Controller operations for Page404
type Page404Controller struct {
	BaseController
}

// Post ...
// @Title Create
// @Description create Page404
// @Param	body		body 	forms.Page404	true		"body for Page404 content"
// @Success 201 {object} models.Page404
// @Failure 403 body is empty
// @router / [post]
func (c *Page404Controller) Post() {
	var form forms.Page404
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	if err != nil {
		c.Data["json"] = ResponseError(err)
	} else {
		page404, err := models.NewPage404(form.Ip, form.Method, form.Url, form.Code1, form.Code2, form.Status, time.Now())
		if err != nil {
			c.Data["json"] = ResponseError(err)
		} else {
			c.Data["json"] = ResponseOK(page404)
		}
	}
	_ = c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Page404
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Param	group	query	string	false	"Filter. e.g. url ..."
// @Success 200 {object} models.Page404
// @Failure 403
// @router / [get]
func (c *Page404Controller) GetAll() {
	cleandb, _ := c.GetBool("cleandb", false)
	group, _ := c.GetBool("group", false)
	if cleandb == true {
		c.Delete()
	} else if group == false {
		page404s, err := models.ListPage404()
		if err != nil {
			c.Data["json"] = ResponseError(err)
		} else {
			c.Data["json"] = ResponseOK(page404s)
		}
	} else {
		data, err := models.ListPage404GroupBy()
		if err != nil {
			c.Data["json"] = ResponseError(err)
		} else {
			c.Data["json"] = ResponseOK(data)
		}
	}

	_ = c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Abc
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Page404Controller) Delete() {
	// check auth
	authSecret := os.Getenv("SECRET")
	currentSecret := c.Ctx.Input.Header("secret")
	if authSecret != currentSecret || currentSecret != "" {
		c.Data["json"] = ResponseError(errors.New("secret not match"))
	} else {
		err := models.DeleteAllPage404()
		if err != nil {
			c.Data["json"] = ResponseError(err)
		} else {
			c.Data["json"] = ResponseOK(nil)
		}
	}

	err := os.RemoveAll(utils.CND_LOG_DIR)
	if err != nil {
		c.Data["json"] = ResponseError(err)
	}

	_ = c.ServeJSON()
}
