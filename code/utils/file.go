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
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(url, destFile string) error {
	var err error

	if strings.HasSuffix(url, "http") == false {
		url = fmt.Sprintf("https://%s", url)
	}

	// download data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create new file
	f, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer f.Close()

	// write file
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
