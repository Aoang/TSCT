package main

import (
	tsct "github.com/Aoang/TSCT"
	"github.com/Aoang/TSCT/api"
	"net/http"
)

func main() {
	tsct.Load(&tsct.Config{
		BotToken:   "",
		TelegramID: 0,
		QQSecret:   "",
	})

	http.HandleFunc("/api", api.Handler)
	_ = http.ListenAndServe("127.0.0.1:8443", nil)
}
