package marketRateType

import "gopkg.in/telebot.v3"

type TelegramMesageBroker struct {
	To   telebot.Recipient
	What interface{}
	Opts []interface{}
}
