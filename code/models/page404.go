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

package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Page404 struct {
	ID      int `orm:"column(id);pk;auto"`
	Ip      string
	Method  string
	Url     string
	Code1   string
	Code2   string
	Status  string
	LogTime *time.Time `orm:"type(datetime);null"`
	Created *time.Time `orm:"auto_now_add;type(datetime)"`
}

func NewPage404(ip, method, url, code1, code2, status string, logTime time.Time) (*Page404, error) {
	if url == "" {
		return nil, errors.New("url is empty")
	}

	p := &Page404{
		Ip:      ip,
		Method:  method,
		Url:     url,
		Code1:   code1,
		Code2:   code2,
		Status:  status,
		LogTime: &logTime,
	}
	o := orm.NewOrm()
	// o := orm.NewOrmUsingDB("default")
	_, _, err := o.ReadOrCreate(p, "Ip", "Url", "LogTime")
	if err != nil {
		return nil, err
	}
	return p, err
}

func ListPage404() ([]*Page404, error) {
	o := orm.NewOrm()
	var pages404s []*Page404
	queryset := o.QueryTable(&Page404{})
	_, err := queryset.All(&pages404s)
	if err != nil {
		return nil, err
	}

	return pages404s, nil
}

func ListPage404GroupBy() (interface{}, error) {
	o := orm.NewOrm()
	rawSeter := o.Raw("select url, count(1) as count from page404 group by url order by count desc")
	var result []orm.Params
	_, err := rawSeter.Values(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteAllPage404() error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM page404").Exec()

	return err
}

func init() {
	orm.RegisterModel(new(Page404))
}
