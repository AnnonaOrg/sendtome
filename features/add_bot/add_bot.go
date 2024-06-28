package add_bot

import (
	"fmt"
	"os"
	"strings"

	"github.com/dbidib/sendtome/features"
	"github.com/dbidib/sendtome/utils"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/addbot", OnAddBot)
}

// Command: /start <PAYLOAD>
func OnAddBot(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}
	senderID := fmt.Sprintf("%s", c.Message().Sender.ID)
	if !strings.EqualFold(senderID, os.Getenv("SENDTOME_ID")) {
		return nil
	}
	botToken := c.Message().Payload
	botToken = strings.TrimSpace(botToken)

	bot, err := tele.NewBot(tele.Settings{
		Token:       botToken,
		Synchronous: false,
	})
	if err != nil {
		msgText := fmt.Sprintf("NewBot Err: %v", err)
		return c.Reply(msgText)
	}
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
		if _, err := utils.SetTelegramWebhook(botToken, webhookURL); err != nil {
			msgText := fmt.Sprintf("SetTelegramWebhook Err: %v", err)
			return c.Reply(msgText)
		} else {
			msgText := "SetTelegramWebhook Success"
			return c.Reply(msgText)
		}
	} else {
		return c.Reply("BOT_TELEGRAM_WEBHOOK_URL : " + webhookURL)
	}

	// return nil
}
