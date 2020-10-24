# Auto-YunZhanYi-Go

浙江工商大学 每日自动云战役打卡

### 1、配置微信推送

「Server酱」，英文名「ServerChan」，是一款「程序员」和「服务器」之间的通信软件。

使用 Server酱 前提是已有了 GitHub 账号，登录获取到 key 值，并绑定微信即可。然后会把每日打卡的消息给你推送到微信中。

①打开 server 酱的官网[http://sc.ftqq.com/3.version]

②点击右上角的 `登入` 链接

![server-1](https://cdn.jsdelivr.net/gh/ruicky/ruicky.github.io/2020/06/05/jd-sign/server-1.jpg)

③会跳入 GitHub 授权页，在该页面填入你的 GitHub 账户即可

![server-2.jpg](https://cdn.jsdelivr.net/gh/ruicky/ruicky.github.io/2020/06/05/jd-sign/server-2.jpg)

④点击上方的 `微信推送` 链接， 然后点击页面中的 `开始绑定`

![server-3](https://cdn.jsdelivr.net/gh/ruicky/ruicky.github.io/2020/06/05/jd-sign/server-3.jpg)

⑤掏出手机，打开微信，扫描屏幕上的二维码，如果未关注，先关注，然后在绑定即可。

![server-4](https://cdn.jsdelivr.net/gh/ruicky/ruicky.github.io/2020/06/05/jd-sign/server-4.jpg)

⑥绑定后，点击上方的 `发送消息` 链接，就可以看到你自己的 key 值，保存下来，下面会用到。

![server-5](https://cdn.jsdelivr.net/gh/ruicky/ruicky.github.io/2020/06/05/jd-sign/server-5.jpg)

### 2、配置config文件

请参考config.json里的信息填写，该程序支持多用户，在配置json文件中增加多条信息即可。

注意：如果不懂userAgent是什么意思，保持默认即可

### 3、注意事项

1、使用go build 命令自行编译或直接下载release版本使用

2、将config.json和编译好后的可执行文件放在同一目录

3、请授予该目录777权限  chmod -R 777 /dir

4、配合Linux的定时任务使用