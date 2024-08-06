package main

import (
	"time"

	"github.com/AnnonaOrg/osenv"
	_ "github.com/umfaka/sendtome/cmd/sendtome/distro/all"
	"github.com/umfaka/sendtome/common"
	_ "github.com/umfaka/sendtome/internal/dotenv"
	"github.com/umfaka/sendtome/internal/features"
	"github.com/umfaka/sendtome/internal/log"
	tele "gopkg.in/telebot.v3"
)

func mainBot() {
	botToken := osenv.GetBotTelegramToken()
	botAPIProxyURL := osenv.GetBotTelegramAPIProxyURL()
	log.Debugf("GetBotTelegramAPIProxyURL(): %s", botAPIProxyURL)
	bot, err := tele.NewBot(tele.Settings{
		URL:    botAPIProxyURL,
		Token:  botToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	common.Must(err)

	features.Handle(bot)

	commands := []tele.Command{
		{
			Text:        "/start",
			Description: "开始",
		},
		{
			Text:        "/id",
			Description: "获取ID",
		},
		{
			Text:        "/ping",
			Description: "Ping",
		},
		// {
		// 	Text:        "/about",
		// 	Description: "About",
		// },
		{
			Text:        "/version",
			Description: "查看版本",
		},
	}
	bot.SetCommands(commands)

	bot.Start()
}
