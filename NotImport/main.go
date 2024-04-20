package main

import (
	"flag"
	"log"
	"main/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient = telegram.New(tgBotHost, mustToken())

	//fetcher = fetcher.New()

	//processor = processor.New()

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-bot-token",
		"",
		"Введите токен телеграм",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("токен не введёт или введён не верно")
	}

	return *token
}
