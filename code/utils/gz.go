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
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadGzipFile(destFile string) (string, error) {
	f, err := os.Open(destFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		return "", err
	}
	defer gr.Close()

	bs, err := ioutil.ReadAll(gr)
	if err != nil {
		return "", err
	}

	// 以文本形式输出
	return fmt.Sprintf("%s", bs), nil
}
