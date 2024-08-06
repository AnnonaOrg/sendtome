package info

import (
	"fmt"
	"strings"

	"github.com/umfaka/sendtome/internal/features"

	tele "gopkg.in/telebot.v3"
)

func init() {
	features.RegisterFeature("/info", OnInfo)
}
func OnInfo(c tele.Context) error {
	payload := c.Message().Payload
	chatUsername := ""
	if strings.HasPrefix(payload, "@") {
		chatUsername = payload
	} else if strings.HasPrefix(payload, "https://t.me/") {
		if chatPath, isFound := strings.CutPrefix(payload, "https://t.me/"); isFound && len(chatPath) > 0 {
			if chatName, _, isFound := strings.Cut(chatPath, "/"); isFound && len(chatName) > 0 {
				chatUsername = "@" + chatName
			} else {
				chatUsername = "@" + chatPath
			}
		} else {
			return c.Reply("格式错误🙅:" + payload)
		}
	}
	if len(chatUsername) == 0 {
		return c.Reply("请输入username，\n例如 `/info @公开群组频道用户名`", tele.ModeMarkdownV2)
	}

	chat, err := c.Bot().ChatByUsername(chatUsername)
	if err != nil {
		return c.Reply("本次查询失败(请使用公开群组或频道链接查询): " + err.Error())
	}

	text := fmt.Sprintf("%s id: %d", chat.Type, chat.ID)
	if chat.LinkedChatID != 0 {
		text = text + "\n" + fmt.Sprintf("连接群id: %d", chat.LinkedChatID)
	}
	text = text + "\n" + chat.Title
	if len(chat.Bio) > 0 {
		text = text + "\n" + "Bio: " + chat.Bio
	}
	if len(chat.Description) > 0 {
		text = text + "\n" + "Description: " + chat.Description
	}

	if len(chat.Username) > 0 {
		text = text + "\n" + "@" + chat.Username
	}
	return c.Reply(text)
}
