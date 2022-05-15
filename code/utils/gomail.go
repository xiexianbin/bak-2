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

package utils

import (
	"github.com/beego/beego/v2/server/web"
	"gopkg.in/gomail.v2"
)

func SendMail(subject, context string, to ...string) error {
	host := web.AppConfig.DefaultString("mail::host", "")
	port := web.AppConfig.DefaultInt("mail::port", 465)
	username := web.AppConfig.DefaultString("mail::from", "")
	password := web.AppConfig.DefaultString("mail::password", "")
	defaultTo := web.AppConfig.DefaultString("mail::to", "me@xiexianbin.cn")

	m := gomail.NewMessage()
	m.SetHeader("From", username)
	if len(to) > 0 {
		m.SetHeader("To", to...)
	} else {
		m.SetHeader("To", defaultTo)
	}

	m.SetHeader("Subject", subject)
	m.SetBody("text/html", context)

	d := gomail.NewDialer(host, port, username, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
