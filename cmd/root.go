package cmd

import (
	"Auto-NCO-ZJGSU/service"
	"github.com/mizuki1412/go-core-kit/class/exception"
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"github.com/mizuki1412/go-core-kit/service/logkit"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	DefFlags(rootCmd)
}

var rootCmd = &cobra.Command{
	Use: "Auto-YunZhanYi-Go",
	Run: func(cmd *cobra.Command, args []string) {
		initkit.BindFlags(cmd)
		run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}

func DefFlags(cmd *cobra.Command) {
	cmd.Flags().String("trueName", "anonymous", "推送时显示的名字")
	cmd.Flags().String("name", "", "云战役登陆id(学号)")
	cmd.Flags().String("psswd", "", "云战役登陆密码(默认身份证后6位)")
	cmd.Flags().String("home", "浙江省杭州市江干区浙江工商大学金沙港", "打卡地址")
	cmd.Flags().String("userAgent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1", "UserAgent")
	cmd.Flags().String("wechatPushKey", "", "server酱提供的PushKey")
}

func run() {
	var users []interface{}
	users, err := cast.ToSliceE(configkit.Get("id", ""))
	if users == nil || err != nil {
		users = []interface{}{map[string]string{"trueName": configkit.GetStringD("trueName"), "name": configkit.GetStringD("name"), "psswd": configkit.GetStringD("psswd"), "home": configkit.GetStringD("home"), "userAgent": configkit.GetStringD("userAgent"), "wechatPushKey": configkit.GetStringD("wechatPushKey")}}
	}
	for _, v := range users {
		user, err := cast.ToStringMapStringE(v)
		if err != nil {
			continue
		}
		if user["name"] == "" || user["psswd"] == "" {
			continue
		}
		do(user)
		time.Sleep(1 * time.Second)
	}
}

//打卡
func do(user map[string]string) {
	//加recover避免某一用户打卡失败导致程序panic
	defer func() {
		if err := recover(); err != nil {
			var msg string
			if e, ok := err.(exception.Exception); ok {
				msg = e.Msg
				// 带代码位置信息
				logkit.Error(e.Error())
			} else {
				msg = cast.ToString(err)
				logkit.Error(msg)
			}
		}
	}()
	cookie := service.GetCookie(user["name"], user["psswd"], user["userAgent"], user["home"])
	result := service.PostInfo(cookie, user["userAgent"])
	if user["wechatPushKey"] != "" {
		service.Push2WeChat(user["wechatPushKey"], user["name"], user["trueName"], result)
	}
}
