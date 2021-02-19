package handler

import (
	"net/http"

	t "github.com/Aoang/TSCT"
)

// -------------------------------- //
//  配置覆盖优先级
//    环境变量 > const 配置
const (
	BotToken   = ""
	TelegramID = 0
	QQSecret   = ""
)

// -------------------------------- //

func init() {
	e := t.GetEnv()
	if e == nil {
		e = &t.Config{
			BotToken:   BotToken,
			TelegramID: TelegramID,
			QQSecret:   QQSecret,
		}
	}

	t.Load(e)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	t.Handler(r)
}
