package service

import (
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/mizuki1412/go-core-kit/class/exception"
	"net/http"
	"strings"
)

var m = map[string]string{
	"uuid":           "",
	"locationInfo":   "浙江省杭州市",
	"currentResd":    "",
	"fromHbToZjDate": "",
	"fromHbToZj":     "C",
	"fromWtToHzDate": "",
	"fromWtToHz":     "B",
	"meetDate":       "",
	"meetCase":       "C",
	"travelDate":     "",
	"travelCase":     "D",
	"medObsvReason":  "",
	"medObsv":        "B",
	"belowCaseDesc":  "",
	"belowCase":      "D",
	"temperature":    "",
	"notApplyReason": "",
	"hzQRCode":       "A",
	"specialDesc":    "",
}

//获取cookie值,生成uuid并赋值,home赋值
func GetCookie(user, pwd, userAgent, home string) string {
	var cookie string
	cookies, _ := Request(Req{
		Url:    "https://nco.zjgsu.edu.cn/login",
		Method: http.MethodPost,
		Header: map[string]string{
			"User-Agent": userAgent,
		},
		FormData: map[string]string{
			"name":  user,
			"psswd": pwd,
		},
	})
	if cookies != nil {
		u, err := uuid.NewV4()
		if err != nil {
			panic(exception.New(err.Error()))
		}
		uid := u.String()
		m["uuid"] = uid
		m["currentResd"] = home
		cookie = cookies.Name + "=" + cookies.Value + ";" + " _ncov_uuid=" + uid + "; _ncov_username=" + user + "; _ncov_psswd=" + pwd
	}
	if cookie == "" {
		panic(exception.New("cookie null"))
	}
	return cookie
}

//post 报送表单json
func PostInfo(cookie, userAgent string) string {
	_, body := Request(Req{
		Url:    "https://nco.zjgsu.edu.cn/",
		Method: http.MethodPost,
		Header: map[string]string{
			"User-Agent": userAgent,
			"Cookie":     cookie,
		},
		JsonData: m,
	})
	if strings.Contains(body, "报送成功") {
		return "ok"
	}
	if strings.Contains(body, "当天已报送") {
		return "already"
	}
	return body
}

// 微信推送
func Push2WeChat(pushKey, id, name string, result string) {
	if result == "ok" {
		Request(Req{
			Url:    "https://sc.ftqq.com/" + pushKey + ".send",
			Method: http.MethodPost,
			FormData: map[string]string{
				"text": "打卡成功" + id + name,
			},
		})
		return
	}
	if result == "already" {
		Request(Req{
			Url:    "https://sc.ftqq.com/" + pushKey + ".send",
			Method: http.MethodPost,
			FormData: map[string]string{
				"text": "当天已经打卡" + id + name,
			},
		})
		return
	}
	Request(Req{
		Url:    "https://sc.ftqq.com/" + pushKey + ".send",
		Method: http.MethodPost,
		FormData: map[string]string{
			"text": "打卡失败" + id + name,
			"desp": result,
		},
	})
}
