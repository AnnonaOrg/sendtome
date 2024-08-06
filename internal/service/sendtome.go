package service

import (
	"os"

	"github.com/AnnonaOrg/osenv"
)

func GetSendToMeID(botID int64) string {
	if chatIDStr := osenv.GetBotReportChatID(); len(chatIDStr) > 0 {
		return chatIDStr
	} else if chatIDStr := os.Getenv("SENDTOME_ID"); len(chatIDStr) > 0 {
		return chatIDStr
	}
	return ""
}
