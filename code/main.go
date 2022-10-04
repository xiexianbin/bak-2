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

package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"

	_ "github.com/xiexianbin/fc-aliyun-cdn-404/code/models"
	_ "github.com/xiexianbin/fc-aliyun-cdn-404/code/routers"
)

func init() {
	orm.Debug = true
	_ = orm.RegisterDriver("sqlite3", orm.DRSqlite)
	_ = orm.RegisterDataBase("default", "sqlite3", "db.sqlite3")

	// sync db struct
	_ = orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
