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

package xhttp

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"net/http/cookiejar"

	"github.com/Andyfoo/go-xutils/xlog"
)

const (
	UserAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	Accept         = "text/xml,application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5"
	AcceptEncoding = "gzip,deflate"
	AcceptLanguage = "zh-cn"
	AcceptCharset  = "gb2312,utf-8;q=0.7,*;q=0.7"
	Timeout        = 30

	ContentType_xml      = "application/xml"
	ContentType_json     = "application/json"
	ContentType_multipar = "multipart/form-data"
	ContentType_form     = "application/x-www-form-urlencoded"

	Method_GET    = "GET"
	Method_POST   = "POST"
	Method_PUT    = "PUT"
	Method_PATCH  = "PATCH"
	Method_DELETE = "DELETE"
)

type ReqParm struct {
	UserAgent string
	Referer   string
	MediaType string
	Charset   string
}
type HttpUtil struct {
	UserAgent      string
	Accept         string
	AcceptEncoding string
	AcceptLanguage string
	AcceptCharset  string
	Timeout        time.Duration
	Headers        map[string]string
	UseCookie      bool
	CookieJar      *cookiejar.Jar
}

func NewDefaultHttpUtil() HttpUtil {
	headers := make(map[string]string)
	return HttpUtil{
		UserAgent:      UserAgent,
		Accept:         Accept,
		AcceptEncoding: AcceptEncoding,
		AcceptLanguage: AcceptLanguage,
		AcceptCharset:  AcceptCharset,
		Timeout:        Timeout,
		Headers:        headers,
	}
}
func NewHttpUtil(UserAgent, Accept, AcceptEncoding, AcceptLanguage, AcceptCharset string, Timeout time.Duration) HttpUtil {
	headers := make(map[string]string)
	return HttpUtil{
		UserAgent:      UserAgent,
		Accept:         Accept,
		AcceptEncoding: AcceptEncoding,
		AcceptLanguage: AcceptLanguage,
		AcceptCharset:  AcceptCharset,
		Timeout:        Timeout,
		Headers:        headers,
	}
}

func (hUtil *HttpUtil) GetTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		Dial:                  hUtil.TimeoutDialer(hUtil.Timeout*time.Second, hUtil.Timeout*time.Second),
		TLSHandshakeTimeout:   hUtil.Timeout * time.Second,
		ResponseHeaderTimeout: hUtil.Timeout * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
func (hUtil *HttpUtil) TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func (hUtil *HttpUtil) addHeader(header *http.Header) {
	header.Set("User-Agent", hUtil.UserAgent)
	header.Set("Accept", hUtil.Accept)
	header.Set("Accept-Encoding", hUtil.AcceptEncoding)
	header.Set("Accept-Language", hUtil.AcceptLanguage)
	header.Set("Accept-Charset", hUtil.AcceptCharset)
	if hUtil.Headers != nil {
		for k, v := range hUtil.Headers {
			header.Set(k, v)
		}
	}
}

func (hUtil *HttpUtil) InitCookie() {
	if hUtil.CookieJar == nil {
		jar, err := cookiejar.New(nil)
		if err != nil {
			xlog.Errorf("cookiejar error: %v", err)
			return
		}
		hUtil.CookieJar = jar
	}
}
func (hUtil *HttpUtil) addCookie(client *http.Client) {
	if !hUtil.UseCookie {
		return
	}
	hUtil.InitCookie()
	if hUtil.CookieJar != nil {
		client.Jar = hUtil.CookieJar
	} else {
		xlog.Errorf("client.Jar is nil")
	}

}
func (hUtil *HttpUtil) Get(urlStr string, _reqParm ...ReqParm) string {
	xlog.Infof("get %s", urlStr)

	client := &http.Client{Transport: hUtil.GetTransport()}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		xlog.Errorf("http.NewRequest error: %v", err)
		return ""
	}
	hUtil.addHeader(&req.Header)
	hUtil.addCookie(client)

	var reqParm ReqParm
	if len(_reqParm) > 0 {
		reqParm = _reqParm[0]
	}
	if len(reqParm.UserAgent) > 0 {
		req.Header.Set("User-Agent", reqParm.UserAgent)
	}
	if len(reqParm.Referer) > 0 {
		req.Header.Set("Referer", reqParm.Referer)
	}

	//使用这个方式 需要自己复制 301、302跳转
	// tr := &http.Transport{}
	// response, err := tr.RoundTrip(req)
	// if err != nil {
	// 	xlog.Errorf("tr.RoundTrip: %v", err)
	// 	return ""
	// }
	// fmt.Println(response.Header)

	//fmt.Println(req.Header)
	resp, err := client.Do(req)
	if err != nil {
		xlog.Errorf("client.Do: %v", err)
		return ""
	}
	//fmt.Println(resp)
	//fmt.Println(resp.Header)
	//fmt.Println("Cookie", req.Header["Cookie"])

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		xlog.Errorf("http status error: %s,%d", resp.Status, resp.StatusCode)
		return ""
	}
	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			xlog.Errorf("http resp unzip is failed,err: %v", err)
			return ""
		}
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		xlog.Errorf("ioutil.ReadAll: %v", err)
		return ""
	}
	return string(data)
}
func (hUtil *HttpUtil) PostMap(urlStr string, postDataMap map[string]string, _reqParm ...ReqParm) string {
	var reqParm ReqParm
	if len(_reqParm) > 0 {
		reqParm = _reqParm[0]
	}
	postData := url.Values{}
	for k, v := range postDataMap {
		postData.Set(k, v)
	}
	return hUtil.Post(urlStr, postData, reqParm)
}
func (hUtil *HttpUtil) Post(urlStr string, postData url.Values, _reqParm ...ReqParm) string {
	xlog.Info(fmt.Sprintf("post %s", urlStr))
	client := &http.Client{Transport: hUtil.GetTransport()}
	bodyData := ioutil.NopCloser(strings.NewReader(postData.Encode()))
	req, err := http.NewRequest("POST", urlStr, bodyData)
	if err != nil {
		xlog.Errorf("http.NewRequest error: %v", err)
		return ""
	}

	hUtil.addHeader(&req.Header)
	hUtil.addCookie(client)

	var reqParm ReqParm
	if len(_reqParm) > 0 {
		reqParm = _reqParm[0]
	}

	if len(reqParm.MediaType) > 0 && len(reqParm.Charset) > 0 {
		req.Header.Set("Content-Type", fmt.Sprintf("%s;%s", reqParm.MediaType, reqParm.Charset))
	} else if len(reqParm.MediaType) > 0 {
		req.Header.Set("Content-Type", reqParm.MediaType)
	} else {
		req.Header.Set("Content-Type", ContentType_form)
	}

	hUtil.addHeader(&req.Header)
	if len(reqParm.UserAgent) > 0 {
		req.Header.Set("User-Agent", reqParm.UserAgent)
	}
	if len(reqParm.Referer) > 0 {
		req.Header.Set("Referer", reqParm.Referer)
	}

	resp, err := client.Do(req)
	if err != nil {
		xlog.Errorf("client.Do: %v", err)
		return ""
	}
	//fmt.Println(resp.Header)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		xlog.Errorf("http status error: %s,%d", resp.Status, resp.StatusCode)
		return ""
	}
	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			xlog.Errorf("http resp unzip is failed,err: %v", err)
			return ""
		}
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		xlog.Errorf("ioutil.ReadAll: %v", err)
		return ""
	}
	return string(data)
}

func (hUtil *HttpUtil) DownFile(urlStr string, _reqParm ...ReqParm) ([]byte, string) {
	xlog.Info(fmt.Sprintf("DownFile %s", urlStr))

	client := &http.Client{Transport: hUtil.GetTransport()}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		xlog.Errorf("http.NewRequest error: %v", err)
		return nil, ""
	}
	hUtil.addHeader(&req.Header)
	hUtil.addCookie(client)

	var reqParm ReqParm
	if len(_reqParm) > 0 {
		reqParm = _reqParm[0]
	}
	if len(reqParm.UserAgent) > 0 {
		req.Header.Set("User-Agent", reqParm.UserAgent)
	}
	if len(reqParm.Referer) > 0 {
		req.Header.Set("Referer", reqParm.Referer)
	}

	resp, err := client.Do(req)
	if err != nil {
		xlog.Errorf("client.Do: %v", err)
		return nil, ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		xlog.Errorf("http status error: %s,%d", resp.Status, resp.StatusCode)
		return nil, ""
	}
	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			xlog.Errorf("http resp unzip is failed,err: %v", err)
			return nil, ""
		}
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		xlog.Errorf("ioutil.ReadAll: %v", err)
		return nil, ""
	}
	filename := hUtil.getFileName(resp)
	return data, filename
}

func (hUtil *HttpUtil) getFileName(resp *http.Response) string {
	//fmt.Println(resp.Header)
	contentDisposition := resp.Header.Get("Content-Disposition")
	if len(contentDisposition) == 0 {
		return ""
	}
	re := regexp.MustCompile(`filename=["]*([^"]+)["]*`)
	matched := re.FindStringSubmatch(contentDisposition)
	if matched == nil || len(matched) == 0 || len(matched[0]) == 0 {
		//fmt.Println("######")
		return ""
	}

	//fmt.Println(matched)
	return matched[1]
}

func (hUtil *HttpUtil) GetBytes(urlStr string, reqParm ...ReqParm) []byte {
	xlog.Info(fmt.Sprintf("getBytes %s", urlStr))

	client := &http.Client{Transport: hUtil.GetTransport()}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		xlog.Errorf("http.NewRequest error: %v", err)
		return nil
	}
	hUtil.addHeader(&req.Header)
	hUtil.addCookie(client)

	if len(reqParm) > 0 && len(reqParm[0].UserAgent) > 0 {
		req.Header.Set("User-Agent", reqParm[0].UserAgent)
	}
	if len(reqParm) > 0 && len(reqParm[0].Referer) > 0 {
		req.Header.Set("Referer", reqParm[0].Referer)
	}

	resp, err := client.Do(req)
	if err != nil {
		xlog.Errorf("client.Do: %v", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		xlog.Errorf("http status error: %s,%d", resp.Status, resp.StatusCode)
		return nil
	}
	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			xlog.Errorf("http resp unzip is failed,err: %v", err)
			return nil
		}
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		xlog.Errorf("ioutil.ReadAll: %v", err)
		return nil
	}
	return data
}

func (hUtil *HttpUtil) Req(method string, urlStr string, postData []byte, _reqParm ...ReqParm) []byte {
	xlog.Info(fmt.Sprintf("Req %s, %s", method, urlStr))
	client := &http.Client{Transport: hUtil.GetTransport()}
	bodyData := ioutil.NopCloser(bytes.NewReader(postData))
	req, err := http.NewRequest(method, urlStr, bodyData)
	if err != nil {
		xlog.Errorf("http.NewRequest error: %v", err)
		return nil
	}
	var reqParm ReqParm
	if len(_reqParm) > 0 {
		reqParm = _reqParm[0]
	}

	if len(reqParm.MediaType) > 0 && len(reqParm.Charset) > 0 {
		req.Header.Set("Content-Type", fmt.Sprintf("%s;%s", reqParm.MediaType, reqParm.Charset))
	} else if len(reqParm.MediaType) > 0 {
		req.Header.Set("Content-Type", reqParm.MediaType)
	}

	hUtil.addHeader(&req.Header)
	hUtil.addCookie(client)

	if len(reqParm.UserAgent) > 0 {
		req.Header.Set("User-Agent", reqParm.UserAgent)
	}
	if len(reqParm.Referer) > 0 {
		req.Header.Set("Referer", reqParm.Referer)
	}

	resp, err := client.Do(req)
	if err != nil {
		xlog.Errorf("client.Do: %v", err)
		return nil
	}
	//fmt.Println(resp.Header)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		xlog.Errorf("http status error: %s,%d", resp.Status, resp.StatusCode)
		return nil
	}
	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			xlog.Errorf("http resp unzip is failed,err: %v", err)
			return nil
		}
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		xlog.Errorf("ioutil.ReadAll: %v", err)
		return nil
	}
	return data
}

func (hUtil *HttpUtil) Head(urlStr string, postData []byte, _reqParm ...ReqParm) http.Header {
	xlog.Info(fmt.Sprintf("Head %s", urlStr))
	client := &http.Client{Transport: hUtil.GetTransport()}
	bodyData := ioutil.NopCloser(bytes.NewReader(postData))
	req, err := http.NewRequest("HEAD", urlStr, bodyData)
	if err != nil {
		xlog.Errorf("http.NewRequest error: %v", err)
		return nil
	}
	var reqParm ReqParm
	if len(_reqParm) > 0 {
		reqParm = _reqParm[0]
	}

	if len(reqParm.MediaType) > 0 && len(reqParm.Charset) > 0 {
		req.Header.Set("Content-Type", fmt.Sprintf("%s;%s", reqParm.MediaType, reqParm.Charset))
	} else if len(reqParm.MediaType) > 0 {
		req.Header.Set("Content-Type", reqParm.MediaType)
	}

	hUtil.addHeader(&req.Header)
	hUtil.addCookie(client)

	if len(reqParm.UserAgent) > 0 {
		req.Header.Set("User-Agent", reqParm.UserAgent)
	}
	if len(reqParm.Referer) > 0 {
		req.Header.Set("Referer", reqParm.Referer)
	}

	resp, err := client.Do(req)
	if err != nil {
		xlog.Errorf("client.Do: %v", err)
		return nil
	}
	//fmt.Println(resp.Header)
	defer resp.Body.Close()

	return resp.Header
}
