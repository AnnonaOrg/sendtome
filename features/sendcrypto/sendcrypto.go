package sendcrypto

import (
	"os"

	"github.com/dbidib/sendtome/features"
	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/sendcrypto", OnSendCrypto)
	features.RegisterFeature("/sendUSDT", OnSendCryptoUSDT)
}

// Command: /sendcrypto <PAYLOAD>
func OnSendCrypto(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	welcomeMsg := os.Getenv("SEND_CRYPTO_MSG")
	if len(welcomeMsg) == 0 {
		welcomeMsg = "Send crypto (发送加密货币)?" + "\n" +
			"请直接联系获取地址！！！"
	}
	if len(welcomeMsg) > 0 {
		return c.Reply(welcomeMsg, tele.ModeHTML)
	}
	return nil
}

// Command: /sendcrypto <PAYLOAD>
func OnSendCryptoUSDT(c tele.Context) error {
	if !c.Message().Private() {
		return nil
	}

	welcomeMsg := os.Getenv("SEND_CRYPTO_USDT_MSG")
	if len(welcomeMsg) == 0 {
		welcomeMsg = "Send crypto USDT(发送加密货币USDT)?" + "\n" +
			"请直接联系获取地址！！！"
	}
	if len(welcomeMsg) > 0 {
		return c.Reply(welcomeMsg, tele.ModeHTML)
	}
	return nil
}
