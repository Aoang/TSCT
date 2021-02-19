package main

import (
	"net/http"

	tsct "github.com/Aoang/TSCT"
	"github.com/Aoang/TSCT/api"
)

func main() {
	tsct.Load(&tsct.Config{
		BotToken:   "",
		TelegramID: 0,
		QQSecret:   "",
	})

	http.HandleFunc("/handler", handler.Handler)
	_ = http.ListenAndServe("127.0.0.1:8443", nil)
}
