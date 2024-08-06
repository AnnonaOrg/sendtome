package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	_ "github.com/AnnonaOrg/sendtome/cmd/sendtome/distro/all"
	"github.com/AnnonaOrg/sendtome/common"
	"github.com/AnnonaOrg/sendtome/core/features"
	"github.com/AnnonaOrg/sendtome/core/utils"
	tele "gopkg.in/telebot.v3"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
	// log.Println("r.URL.Path", r.URL.Path)
	botToken := ""
	if tmpText, isFound := strings.CutPrefix(r.URL.Path, "/webhook/tele/"); isFound {
		if tmpText2, isFound := strings.CutSuffix(tmpText, "/"); isFound {
			tmpText = tmpText2
		}
		botToken = tmpText
	} else {
		return
	}

	body, err := io.ReadAll(r.Body)
	common.Must(err)
	var u tele.Update
	common.Must(json.Unmarshal(body, &u))

	bot, err := tele.NewBot(tele.Settings{
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
		webhookURL := fmt.Sprintf("%s/%s", os.Getenv("BOT_TELEGRAM_WEBHOOK_URL"), botToken)
		if len(webhookURL) > 0 && strings.HasPrefix(webhookURL, "https") {
			utils.SetTelegramWebhook(botToken, webhookURL)
		}
	}
	features.Handle(bot)
	bot.ProcessUpdate(u)
}
