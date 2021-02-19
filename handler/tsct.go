package api

import (
	"net/http"

	tsct "github.com/Aoang/TSCT"
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
	e := tsct.GetEnv()
	if e == nil {
		e = &tsct.Config{
			BotToken:   BotToken,
			TelegramID: TelegramID,
			QQSecret:   QQSecret,
		}
	}

	tsct.Load(e)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tsct.Handler(r)
}
