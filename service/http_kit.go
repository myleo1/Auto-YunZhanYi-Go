package service

import (
	"bytes"
	"crypto/tls"
	"github.com/mizuki1412/go-core-kit/library/jsonkit"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

var client *http.Client

func init() {
	//跳过校验证书
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar
}

// 填写FormData、JsonData时可缺省contentType
type Req struct {
	Method      string
	Url         string
	Header      map[string]string
	ContentType string
	FormData    map[string]string
	JsonData    interface{}
	BinaryData  []byte
	Timeout     int // seconds
}

const ContentTypeForm = "application/x-www-form-urlencoded"
const ContentTypeJSON = "application/json"

func Request(reqBean Req) (*http.Cookie, string) {
	if reqBean.Method == "" {
		reqBean.Method = http.MethodPost
	}
	var req *http.Request
	var err error
	if reqBean.BinaryData != nil {
		req, err = http.NewRequest(reqBean.Method, reqBean.Url, bytes.NewBuffer(reqBean.BinaryData))
	} else if reqBean.JsonData != nil {
		req, err = http.NewRequest(reqBean.Method, reqBean.Url, bytes.NewBuffer([]byte(jsonkit.ToString(reqBean.JsonData))))
	} else {
		data := make(url.Values)
		for key, val := range reqBean.FormData {
			data.Add(key, val)
		}
		req, err = http.NewRequest(reqBean.Method, reqBean.Url, strings.NewReader(data.Encode()))
	}
	if err != nil {
		panic("reqErr")
	}
	if reqBean.ContentType == "" {
		if reqBean.FormData != nil {
			req.Header.Set("Content-Type", ContentTypeForm)
		} else if reqBean.JsonData != nil {
			req.Header.Set("Content-Type", ContentTypeJSON)
		}
	}
	for key, val := range reqBean.Header {
		req.Header.Set(key, val)
	}
	if reqBean.Timeout > 0 {
		client.Timeout = time.Duration(reqBean.Timeout) * time.Second
	}
	resp, err := client.Do(req)
	if err != nil {
		panic("client.DoErr")
	}
	defer resp.Body.Close()
	var cookies *http.Cookie
	if len(resp.Cookies()) > 0 {
		cookies = resp.Cookies()[0]
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("readBodyErr")
	}
	return cookies, string(body)
}
