package main

import (
	tsct "github.com/Aoang/TSCT"
	"net/http"
)

func main() {
	tsct.Load(&tsct.Config{
		BotToken:   "",
		TelegramID: 0,
		QQSecret:   "",
	})

	http.HandleFunc("/api", tsct.Handler)
	_ = http.ListenAndServe("127.0.0.1:8443", nil)
}
