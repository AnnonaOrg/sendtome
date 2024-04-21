package main

import (
	"os"
	"time"

	"github.com/dbidib/sendtome/common"
	"github.com/dbidib/sendtome/features"
	_ "github.com/dbidib/sendtome/internal/dotenv"
	_ "github.com/dbidib/sendtome/main/distro/all"

	tele "gopkg.in/telebot.v3"
)

func main() {
	b, err := tele.NewBot(tele.Settings{
		Token:  os.Getenv("BOT_TELEGRAM_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	common.Must(err)

	features.Handle(b)

	commands := []tele.Command{
		{
			Text:        "/id",
			Description: "Getid",
		},
		{
			Text:        "/ping",
			Description: "Ping",
		},
		{
			Text:        "/about",
			Description: "About",
		},
		{
			Text:        "/start",
			Description: "Start",
		},
	}
	b.SetCommands(commands)
	b.Start()
}
