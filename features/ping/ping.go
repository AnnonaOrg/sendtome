package ping

import (
	"fmt"

	"github.com/dbidib/sendtome/features"

	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/ping", OnPing)
}
func OnPing(c tele.Context) error {
	if !c.Message().Private() {
		return c.Reply("pong")
	}
	text := fmt.Sprintf("Pong! %s%s @%s(%d)",
		c.Message().Sender.FirstName, c.Message().Sender.LastName,
		c.Message().Sender.Username, c.Message().Sender.ID,
	)
	return c.Reply(text)
}
