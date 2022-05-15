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

package forms

type Page404 struct {
	Ip      string `json:"ip"`
	Method  string `json:"method"`
	Url     string `json:"url"`
	Code1   string `json:"code1"`
	Code2   string `json:"code2"`
	Status  string `json:"status"`
	LogTime string `json:"log_time"`
}
