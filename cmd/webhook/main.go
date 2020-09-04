package main

import (
	tsct "github.com/Aoang/TSCT"
)

func main() {
	tsct.Load(&tsct.Config{
		BotToken:   "",
		WebhookURL: "",
	})
}
