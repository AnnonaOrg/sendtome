package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dbidib/sendtome/common"
	"github.com/dbidib/sendtome/features"
	_ "github.com/dbidib/sendtome/main/distro/all"
	"github.com/dbidib/sendtome/utils"
	tele "gopkg.in/telebot.v3"
)

var (
	bot *tele.Bot
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
	body, err := io.ReadAll(r.Body)
	common.Must(err)
	log.Println(string(body))

	var u tele.Update
	common.Must(json.Unmarshal(body, &u))

	bot.ProcessUpdate(u)
}

func init() {
	var err error
	botToken := os.Getenv("BOT_TELEGRAM_TOKEN")
	bot, err = tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: true,
	})
	common.Must(err)
	if strings.EqualFold(os.Getenv("ENABLE_SET_WEBHOOK"), "true") {
		// log.Println("ENABLE_SET_WEBHOOK: true")
		commands := []tele.Command{
			{
				Text:        "/start",
				Description: "Start",
			},
			{
				Text:        "/id",
				Description: "查询自己的用户id信息",
			},
			{
				Text:        "/info",
				Description: "查询公开群组频道信息",
			},
			{
				Text:        "/ping",
				Description: "Ping",
			},
		}

		if len(os.Getenv("SEND_CRYPTO_MSG")) > 0 {
			commands = append(commands, tele.Command{
				Text:        "/sendcrypto",
				Description: "Send crypto (发送加密货币)",
			})
		}

		bot.SetCommands(commands)

		webhookURL := os.Getenv("BOT_TELEGRAM_WEBHOOK_URL")
		if len(webhookURL) > 0 && strings.HasPrefix(webhookURL, "https") {
			utils.SetTelegramWebhook(botToken, webhookURL)
		}

	}

	features.Handle(bot)
}
