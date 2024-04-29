package telegram

import "main/clients/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
}

func New(Client *telegram.Client, storage)
