package service

import (
	"strings"

	"github.com/AnnonaOrg/osenv"
	"github.com/umfaka/sendtome/core/log"
	"github.com/umfaka/sendtome/core/utils"
)

func SetBotFatherWebhook() {
	botToken := osenv.GetBotTelegramToken()
	webhookURL := osenv.GetBotTelegramWebhookURL() //os.Getenv("BOT_TELEGRAM_WEBHOOK_URL")
	if len(webhookURL) > 0 && strings.HasPrefix(webhookURL, "https") {
		if tmpText, err := utils.SetTelegramWebhook(botToken, webhookURL+"/"+botToken); err != nil {
			log.Errorf("SetTelegramWebhook(%s): %v", webhookURL, err)
		} else {
			log.Debugf("SetTelegramWebhook(%s): %s", webhookURL, tmpText)
		}
	} else {
		log.Debugf("SetBotFatherWebhook(%s,%s)", botToken, webhookURL)
	}
}
