# Auto-YunZhanYi-Go

> 浙江工商大学 每日自动云战役打卡

## 目录

* [配置微信推送](#配置微信推送)
* [配置config文件](#配置config文件)
* [使用帮助](#使用帮助)
* [Docker食用](#Docker食用)
* [联系](#联系)

## 

## 配置微信推送

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

## 配置config文件

①请参考config.json里的信息填写，详细帮助请使用如下命令查看

```bash
Auto-YunZhanYi-Go -h
```

②该程序支持多用户，在config.json文件中增加多条信息即可

## 使用帮助

①自行编译或直接下载release版本配合Linux定时任务使用

交叉编译到Linux：

```bash
make build
```

交叉编译到Arm Linux:

```bash
make build_arm
```

②将Auto-YunZhanYi-Go和config.json放入服务器

③授予Auto-YunZhanYi-Go 777权限  

```bash
chmod 777 Auto-Yun-ZhanYi-Go
```

注意：执行的时候会读取当前目录下的config.json,如果用定时任务执行Auto-YunZhanYi-Go请确保 ~/ 目录下存放了config.json

## Docker食用

```bash
docker run -d --network=bridge -v /home/config.json:/root/config.json --restart=always --name yunzhanyi-go myleo1/yunzhanyi-go
```

替换/home/config.json为config.json所在目录即可

注意：Docker版本打卡时间为每天早晨7点04分

## 联系

Created by [@myleo1](https://github.com/myleo1) - feel free to contact me!