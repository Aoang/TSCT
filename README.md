# QQ 安全中心动态密钥 (Tencent QQ Security Center Token)

不得不说，有现成的技术方案不用，非要跑去魔改一个，真是国内众多大厂的特征。

整的没办法迁入 Bitwarden，只能玩骚操作来曲线救国。

提取密钥，放入 Vercel 的 Serverless 监听 Telegram。

这样使用的时候只需要给自己的机器人发条消息就好了。

### 前言

现有的 One Time Password 方案大多分为 HOTP 和 OTOP，前者基于事件，后者基于时间。

在两步验证的安全策略中，大多数都是采用基于时间的 OTOP 算法。

OTOP 算法一般来说都是通用的，因为都是遵循 [RFC6238](https://tools.ietf.org/html/rfc6238) 来实现的。

而 QQ 安全中心的算法是魔改 RFC6238 实现的，导致除开它自己的客户端，其他的 OTOP 客户端都没办法用。

### 提取 Secret
1. Android 手机
2. 系统已经 Root
3. 安装并登录 QQ 安全中心 (^6.9.10)
4. 关闭 QQ 安全中心，确保其不在后台
5. 提取文件
  1. 保存 `/data/data/com.tencent.token/shared_prefs/token_save_info.xml` 文件内 `.string` 的内容
  2. 复制 `/data/data/com.tencent.token/databases/mobiletoken.db` 至电脑
6. 解密 `mobiletoken.db`
7. 打开 `mobiletoken.db` 查看 `main.token_conf` 提取加密的 Secret 
8. 解密 Secret

解密的部分需要看 [HyperSine](https://github.com/HyperSine) 提供的 [教程工具](https://github.com/HyperSine/forensic-qqtoken#2-%E5%A6%82%E4%BD%95%E8%8E%B7%E5%8F%96secret) 。

### 安装准备
1. 使用 Github 账号 fork 本仓库
2. 使用 Vercel 账号导入 fork 过去的仓库
3. 使用 Telegram 账号获取自己账户的 ID
4. 申请一个 Telegram Bot 拿到 Token
5. 可选，给 Vercel 绑定一个域名

### 部署

给 Telegram Bot 设置 Webhook

可以直接修改 `cmd/webhook/main.go` 中的配置然后运行一次即可


##### 源码部署
！！！ 使用这种部署方式之前先将仓库转为私有库

更改 `api/tsct.go` 文件内的配置信息，例如
```go
const (
	BotToken   = "123456789:abcdefgh"
	TelegramID = 123
	QQSecret   = "987654321"
)
```


##### 环境变量部署

设置 Vercel 环境变量
```text
BOT_TOKEN=
WEBHOOK_URL=
QQ_SECRET=
TELEGRAM_ID=
```


### 使用方法

给 Telegram Bot 发送任意消息，会返回生成的时间及三个一次性密码。


2020/09/02 11:42:42 CST

~~582608~~  |  `463836`  |  `203893`


### 调用方式

```go
package main

import (
	tsct "github.com/Aoang/TSCT"
	"time"
	"fmt"
)

func main() {
	code := tsct.Token("asd",time.Now())
	fmt.Println(code)
}
```




### TODO
- 多 SECRET 支持






















