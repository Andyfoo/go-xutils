// Copyright 2019 Andyfoo
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package xpage

import (
	"fmt"
	"strings"
)

type PageJsonS struct {
	TotalNum int    `json:"totalNum"`
	PageSize int    `json:"pageSize"`
	CurPage  int    `json:"curPage"`
	PageNum  int    `json:"pageNum"`
	PageUrl  string `json:"pageUrl"`
}

func Page(curPage int, pageSize int, totalNum int, url string, parms ...interface{}) PageJsonS {
	pageJson := PageJsonS{}
	pageJson.CurPage = curPage
	pageJson.PageSize = pageSize
	pageJson.TotalNum = totalNum
	pageJson.PageNum = totalNum / pageSize
	if totalNum%pageSize > 0 {
		pageJson.PageNum++
	}
	pageJson.PageUrl = url
	if len(parms) > 0 {
		parmStr, ok := parms[0].(string)
		if strings.Contains(pageJson.PageUrl, "?") {
			pageJson.PageUrl += "&"
		} else if !strings.Contains(pageJson.PageUrl, "?") {
			pageJson.PageUrl += "?"
		}

		if ok {
			pageJson.PageUrl += parmStr
		}

		parmMap, ok := parms[0].(map[string]string)
		if ok {

			for k, v := range parmMap {
				pageJson.PageUrl = fmt.Sprintf("%s%s=%s&", pageJson.PageUrl, k, v)
			}
		}
		pos := len(pageJson.PageUrl) - 1
		if pageJson.PageUrl[pos:] == "&" {
			pageJson.PageUrl = pageJson.PageUrl[0:pos]
		}
	}

	return pageJson
}
