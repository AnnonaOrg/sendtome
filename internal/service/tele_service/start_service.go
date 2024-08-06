package tele_service

import (
	tele "gopkg.in/telebot.v3"
)

// Command: /start <PAYLOAD>
func Start(c tele.Context, note string) {
	if !c.Message().Private() {
		return
	}

	return
}
