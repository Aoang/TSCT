package api

import (
	tsct "github.com/Aoang/TSCT"
	"net/http"
)

func init() {
	// 将配置文件写入代码来部署
	tsct.Load(&tsct.Config{
		BotToken:   "",
		TelegramID: 0,
		QQSecret:   "",
	})

	// 使用 Vercel 的环境变量来部署
	e := tsct.GetEnv()
	if e == nil {
		// 这好像是一个悲剧？
		return
	}
	tsct.Load(e)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tsct.Handler(w, r)
}
