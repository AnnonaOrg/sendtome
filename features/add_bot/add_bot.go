package add_bot

import (
	"fmt"
	// "log"
	"os"
	"strings"

	"github.com/umfaka/sendtome/features"
	"github.com/umfaka/sendtome/internal/utils"
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
	senderID := fmt.Sprintf("%d", c.Message().Sender.ID)
	managerID := os.Getenv("BOT_MANAGER_ID")
	if isManger := strings.EqualFold(senderID, managerID); !isManger {
		// log.Printf("非法用户: %s, Admin: %s", senderID, managerID)
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

	webhookURL := fmt.Sprintf("%s/%s", os.Getenv("BOT_TELEGRAM_WEBHOOK_URL"), botToken)
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
