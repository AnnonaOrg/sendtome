package start

import (
	"os"

	"github.com/umfaka/sendtome/core/features"

	tele "gopkg.in/telebot.v3"
)

// Command: /start <PAYLOAD>
func Onstart(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	welcomeMsg := os.Getenv("WELCOME_MSG")
	if len(welcomeMsg) == 0 {
		welcomeMsg = "What can this bot do?" + "\n" +
			"不是机器人，不是机器人，不是机器人！！！" + "\n" +
			"我就是本人，这里是专门为双向用户无法联系我而搭建的沟通渠道！" + "\n" +
			"双向用户直接在这里发消息就可以，我可以收到，并且回复你！"
	}
	if len(welcomeMsg) > 0 {
		// return c.Send(welcomeMsg)
		return c.Reply(welcomeMsg)
	}
	return nil
}

func init() {
	features.RegisterFeature("/start", Onstart)
}
