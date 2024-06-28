package getid

import (
	"fmt"

	"github.com/dbidib/sendtome/features"

	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/id", OnGetID)
}
func OnGetID(c tele.Context) error {
	text := ""
	if c.Message().FromChannel() {
		text = fmt.Sprintf("%s\n%s(%d)", text,
			c.Message().Chat.Title, c.Message().Chat.ID,
		)
	} else {
		text = fmt.Sprintf("@%s(%d)",
			c.Message().Sender.Username, c.Message().Sender.ID,
		)
		if c.Message().FromGroup() {
			if len(c.Message().Chat.Username) > 0 {
				text = fmt.Sprintf("%s\n%s @%s(%d)", text,
					c.Message().Chat.Title, c.Message().Chat.Username, c.Message().Chat.ID,
				)
			} else {
				text = fmt.Sprintf("%s\n%s(%d)", text,
					c.Message().Chat.Title, c.Message().Chat.ID,
				)
			}
		}
	}

	return c.Reply(text)
}
