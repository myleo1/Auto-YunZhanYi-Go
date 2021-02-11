package main_test

import (
	"Auto-NCO-ZJGSU/service"
	"net/http"
	"testing"
)

func TestWechatPush(t *testing.T) {
	service.Request(service.Req{
		Url:    "http://111.0.80.195:40089/push",
		Method: http.MethodPost,
		Header: map[string]string{
			"token": "97eb40e3-1a41-4f47-a663-ae49613721ec",
		},
		FormData: map[string]string{
			"message": "wechat",
		},
	})
}
