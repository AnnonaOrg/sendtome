package sendtome

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/dbidib/sendtome/features"
	tele "gopkg.in/telebot.v3"
)

var syncMap sync.Map

func init() {
	features.RegisterFeature(tele.OnText, OnPrivateSendToMe)
	features.RegisterFeature(tele.OnPhoto, OnPrivateSendToMeByPhoto)
	features.RegisterFeature(tele.OnAudio, OnPrivateSendToMeByAudio)
	features.RegisterFeature(tele.OnAnimation, OnPrivateSendToMeByAnimation)
	features.RegisterFeature(tele.OnDocument, OnPrivateSendToMeByDocument)
	features.RegisterFeature(tele.OnVideo, OnPrivateSendToMeByVideo)
	features.RegisterFeature(tele.OnVoice, OnPrivateSendToMeByVoice)
}

// OnText
func OnPrivateSendToMe(c tele.Context) error {
	if c.Message().FromChannel() {
		return nil
	}
	if !c.Message().FromGroup() && !c.Message().Private() && !c.Message().IsReply() {
		return nil
	}
	msgId := ""
	if len(c.Message().AlbumID) > 0 {
		msgId = fmt.Sprintf("%d_%s", c.Message().Chat.ID, c.Message().AlbumID)
	} else { // c.Message().ID > 0
		msgId = msgId + fmt.Sprintf("%d_%d", c.Message().Chat.ID, c.Message().ID)
	}
	if _, ok := syncMap.LoadOrStore(msgId, ""); ok {
		return nil
	}

	adminID := os.Getenv("SENDTOME_ID")
	if len(adminID) == 0 {
		return nil
	}
	senderID := fmt.Sprintf("%d", c.Message().Sender.ID)
	// 管理员回复信息
	if c.Message().IsReply() && strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到回复消息(err)：", c.Message())
		} else {
			fmt.Println("收到回复消息：", string(jsonText))
		}
		replyToText := c.Message().ReplyTo.Text
		if len(c.Message().ReplyTo.Caption) > 0 {
			replyToText = c.Message().ReplyTo.Caption
		}
		prefixLine, _, isFound := strings.Cut(replyToText+"\n", "\n")
		if !isFound {
			return c.Reply(
				fmt.Sprintf("回复消息格式异常(OnText)(prefixLine: %s): %+v", replyToText, c.Message().ReplyTo),
			)
		}
		_, sendToID, isFound := strings.Cut(prefixLine, "#id")
		if !isFound {
			return c.Reply("回复消息格式异常(OnText)(sendToID): " + fmt.Sprintf("%+v", c.Message().ReplyTo))
		}

		reciverId, err := strconv.ParseInt(sendToID, 10, 64)
		if err != nil {
			fmt.Print("回复消息格式异常(OnText): 待回复id %s\n%+v", sendToID, c.Message().ReplyTo)
		}
		reciver := &tele.User{
			ID: reciverId, //int64(reciverId),
		}

		if _, err := c.Bot().Send(reciver, c.Message().Text); err != nil {
			// return err
			return c.Reply("⚠️回复内容转投失败，请重试: " + err.Error())
		}
		// return nil
		return c.Reply("✅回复内容转投成功。")
	}

	// 收到私聊消息
	if c.Message().Private() && !strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到私聊消息(err)：", c.Message())
		} else {
			fmt.Println("收到私聊消息：", string(jsonText))
		}

		reciverId, err := strconv.ParseInt(adminID, 10, 64)
		if err != nil {
			fmt.Println("设置有误：环境变量(SENDTOME_ID)：", adminID)
		}
		reciver := &tele.User{
			ID: reciverId, //int64(reciverId),
		}

		newMsg := fmt.Sprintf("@%s #id%d\n%s",
			c.Message().Sender.Username,
			c.Message().Sender.ID,
			c.Message().Text)

		if _, err := c.Bot().Send(reciver, newMsg); err != nil {
			return err
			// return c.Reply("感谢！⚠️私聊内容转投失败。" + err.Error())
		}
		return nil
		// return c.Reply("感谢！✅私聊内容转投成功。")
	}
	what := fmt.Sprintf("@%s #id%d\n%s %s\nHi,Admin! 别逗了，宝！\n%s",
		c.Message().Sender.Username,
		c.Message().Sender.ID,
		c.Message().Sender.FirstName,
		c.Message().Sender.LastName,
		c.Message().Text)
	return c.Reply(what)
}

// OnPhoto
func OnPrivateSendToMeByPhoto(c tele.Context) error {
	if c.Message().FromChannel() || c.Message().FromGroup() {
		return nil
	}
	if !c.Message().Private() && !c.Message().IsReply() {
		return nil
	}
	msgId := ""
	if len(c.Message().AlbumID) > 0 {
		msgId = fmt.Sprintf("%d_%s", c.Message().Chat.ID, c.Message().AlbumID)
	} else { // c.Message().ID > 0
		msgId = msgId + fmt.Sprintf("%d_%d", c.Message().Chat.ID, c.Message().ID)
	}
	if _, ok := syncMap.LoadOrStore(msgId, ""); ok {
		return nil
	}

	adminID := os.Getenv("SENDTOME_ID")
	if len(adminID) == 0 {
		return nil
	}
	senderID := fmt.Sprintf("%d", c.Message().Sender.ID)
	// 管理员回复信息
	if c.Message().IsReply() && strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到回复消息(err)：", c.Message())
		} else {
			fmt.Println("收到回复消息：", string(jsonText))
		}
		replyToText := c.Message().ReplyTo.Text
		if len(c.Message().ReplyTo.Caption) > 0 {
			replyToText = c.Message().ReplyTo.Caption
		}
		prefixLine, _, isFound := strings.Cut(replyToText+"\n", "\n")
		if !isFound {
			return c.Reply(
				fmt.Sprintf("回复消息格式异常(OnPhoto)(prefixLine: %s): %+v", replyToText, c.Message().ReplyTo),
			)
		}
		_, sendToID, isFound := strings.Cut(prefixLine, "#id")
		if !isFound {
			return c.Reply("回复消息格式异常(OnPhoto)(sendToID): " + fmt.Sprintf("%+v", c.Message().ReplyTo))
		}

		reciverId, err := strconv.ParseInt(sendToID, 10, 64)
		if err != nil {
			fmt.Print("回复消息格式异常(OnPhoto): 待回复id %s\n%+v", sendToID, c.Message().ReplyTo)
		}
		reciver := &tele.User{
			ID: reciverId,
		}

		if c.Message().Photo != nil {
			newMsg := c.Message().Photo
			if len(c.Message().Caption) > 0 {
				newMsg.Caption = c.Message().Caption
			}
			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return c.Reply("⚠️回复内容转投失败，请重试。" + err.Error())
			}
			return c.Reply("✅回复内容(OnPhoto)转投成功")
		}
		return c.Reply("⚠️回复内容(OnPhoto)转投失败，请重试。" + fmt.Sprintf("获取图片信息失败: %+v", c.Message()))
	}

	// 收到私聊消息
	if c.Message().Private() && !strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到私聊消息(err)：", c.Message())
		} else {
			fmt.Println("收到私聊消息：", string(jsonText))
		}

		reciverId, err := strconv.ParseInt(adminID, 10, 64)
		if err != nil {
			fmt.Println("设置有误：环境变量(SENDTOME_ID)：", adminID)
		}
		reciver := &tele.User{
			ID: reciverId,
		}
		if c.Message().Photo != nil {
			newMsg := c.Message().Photo
			msgText := fmt.Sprintf("@%s #id%d\n%s",
				c.Message().Sender.Username,
				c.Message().Sender.ID,
				c.Message().Caption)
			newMsg.Caption = msgText

			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return err
				// return c.Reply("感谢！⚠️私聊内容转投失败。" + err.Error())
			}
			return nil
			// return c.Reply("感谢！✅私聊内容转投成功。")
		}
		return fmt.Errorf("获取图片信息失败")
	}

	what := fmt.Sprintf("@%s #id%d\n%s %s\nHi,Admin! 别逗了，宝！\n%s",
		c.Message().Sender.Username,
		c.Message().Sender.ID,
		c.Message().Sender.FirstName,
		c.Message().Sender.LastName,
		c.Message().Caption)
	return c.Reply(what)
}

// OnAudio
func OnPrivateSendToMeByAudio(c tele.Context) error {
	if c.Message().FromChannel() || c.Message().FromGroup() {
		return nil
	}
	if !c.Message().Private() && !c.Message().IsReply() {
		return nil
	}
	msgId := ""
	if len(c.Message().AlbumID) > 0 {
		msgId = fmt.Sprintf("%d_%s", c.Message().Chat.ID, c.Message().AlbumID)
	} else { // c.Message().ID > 0
		msgId = msgId + fmt.Sprintf("%d_%d", c.Message().Chat.ID, c.Message().ID)
	}
	if _, ok := syncMap.LoadOrStore(msgId, ""); ok {
		return nil
	}

	adminID := os.Getenv("SENDTOME_ID")
	if len(adminID) == 0 {
		return nil
	}
	senderID := fmt.Sprintf("%d", c.Message().Sender.ID)
	// 管理员回复信息
	if c.Message().IsReply() && strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到回复消息(err)：", c.Message())
		} else {
			fmt.Println("收到回复消息：", string(jsonText))
		}
		replyToText := c.Message().ReplyTo.Text
		if len(c.Message().ReplyTo.Caption) > 0 {
			replyToText = c.Message().ReplyTo.Caption
		}
		prefixLine, _, isFound := strings.Cut(replyToText+"\n", "\n")
		if !isFound {
			return c.Reply(
				fmt.Sprintf("回复消息格式异常(OnAudio)(prefixLine: %s): %+v", replyToText, c.Message().ReplyTo),
			)
		}
		_, sendToID, isFound := strings.Cut(prefixLine, "#id")
		if !isFound {
			return c.Reply("回复消息格式异常(OnAudio)(sendToID): " + fmt.Sprintf("%+v", c.Message().ReplyTo))
		}

		reciverId, err := strconv.ParseInt(sendToID, 10, 64)
		if err != nil {
			fmt.Print("回复消息格式异常(OnAudio): 待回复id %s\n%+v", sendToID, c.Message().ReplyTo)
		}
		reciver := &tele.User{
			ID: reciverId,
		}

		if c.Message().Audio != nil {
			newMsg := c.Message().Audio
			if len(c.Message().Caption) > 0 {
				newMsg.Caption = c.Message().Caption
			}
			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return c.Reply("⚠️回复内容(OnAudio)转投失败，请重试。" + err.Error())
			}
			return c.Reply("✅回复内容(OnAudio)转投成功。")
		}
		return c.Reply("⚠️回复内容(OnAudio)转投失败，请重试。" + fmt.Sprintf("获取图片信息失败: %+v", c.Message()))
	}

	// 收到私聊消息
	if c.Message().Private() && !strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到私聊消息(err)：", c.Message())
		} else {
			fmt.Println("收到私聊消息：", string(jsonText))
		}

		reciverId, err := strconv.ParseInt(adminID, 10, 64)
		if err != nil {
			fmt.Println("设置有误：环境变量(SENDTOME_ID)：", adminID)
		}
		reciver := &tele.User{
			ID: reciverId,
		}
		if c.Message().Audio != nil {
			newMsg := c.Message().Audio
			msgText := fmt.Sprintf("@%s #id%d\n%s",
				c.Message().Sender.Username,
				c.Message().Sender.ID,
				c.Message().Caption)
			newMsg.Caption = msgText

			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return err
				// return c.Reply("感谢！⚠️私聊内容转投失败。" + err.Error())
			}
			return nil
			// return c.Reply("感谢！✅私聊内容转投成功。")
		}
		return fmt.Errorf("获取图片信息失败")
	}

	what := fmt.Sprintf("@%s #id%d\n%s %s\nHi,Admin! 别逗了，宝！\n%s",
		c.Message().Sender.Username,
		c.Message().Sender.ID,
		c.Message().Sender.FirstName,
		c.Message().Sender.LastName,
		c.Message().Caption)
	return c.Reply(what)
}

// OnAnimation
func OnPrivateSendToMeByAnimation(c tele.Context) error {
	if c.Message().FromChannel() || c.Message().FromGroup() {
		return nil
	}
	if !c.Message().Private() && !c.Message().IsReply() {
		return nil
	}
	msgId := ""
	if len(c.Message().AlbumID) > 0 {
		msgId = fmt.Sprintf("%d_%s", c.Message().Chat.ID, c.Message().AlbumID)
	} else { // c.Message().ID > 0
		msgId = msgId + fmt.Sprintf("%d_%d", c.Message().Chat.ID, c.Message().ID)
	}
	if _, ok := syncMap.LoadOrStore(msgId, ""); ok {
		return nil
	}

	adminID := os.Getenv("SENDTOME_ID")
	if len(adminID) == 0 {
		return nil
	}
	senderID := fmt.Sprintf("%d", c.Message().Sender.ID)
	// 管理员回复信息
	if c.Message().IsReply() && strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到回复消息(err)：", c.Message())
		} else {
			fmt.Println("收到回复消息：", string(jsonText))
		}
		replyToText := c.Message().ReplyTo.Text
		if len(c.Message().ReplyTo.Caption) > 0 {
			replyToText = c.Message().ReplyTo.Caption
		}
		prefixLine, _, isFound := strings.Cut(replyToText+"\n", "\n")
		if !isFound {
			return c.Reply(
				fmt.Sprintf("回复消息格式异常(OnAnimation)(prefixLine: %s): %+v", replyToText, c.Message().ReplyTo),
			)
		}
		_, sendToID, isFound := strings.Cut(prefixLine, "#id")
		if !isFound {
			return c.Reply("回复消息格式异常(OnAnimation)(sendToID): " + fmt.Sprintf("%+v", c.Message().ReplyTo))
		}

		reciverId, err := strconv.ParseInt(sendToID, 10, 64)
		if err != nil {
			fmt.Print("回复消息格式异常(OnAnimation): 待回复id %s\n%+v", sendToID, c.Message().ReplyTo)
		}
		reciver := &tele.User{
			ID: reciverId,
		}

		if c.Message().Animation != nil {
			newMsg := c.Message().Animation
			if len(c.Message().Caption) > 0 {
				newMsg.Caption = c.Message().Caption
			}
			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return c.Reply("⚠️回复内容(OnAnimation)转投失败，请重试。" + err.Error())
			}
			return c.Reply("✅回复内容(OnAnimation)转投成功。")
		}
		return c.Reply("⚠️回复内容(OnAnimation)转投失败，请重试。" + fmt.Sprintf("获取图片信息失败: %+v", c.Message()))
	}

	// 收到私聊消息
	if c.Message().Private() && !strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到私聊消息(err)：", c.Message())
		} else {
			fmt.Println("收到私聊消息：", string(jsonText))
		}

		reciverId, err := strconv.ParseInt(adminID, 10, 64)
		if err != nil {
			fmt.Println("设置有误：环境变量(SENDTOME_ID)：", adminID)
		}
		reciver := &tele.User{
			ID: reciverId,
		}
		if c.Message().Animation != nil {
			newMsg := c.Message().Animation
			msgText := fmt.Sprintf("@%s #id%d\n%s",
				c.Message().Sender.Username,
				c.Message().Sender.ID,
				c.Message().Caption)
			newMsg.Caption = msgText

			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return err
				// return c.Reply("感谢！⚠️私聊内容转投失败。" + err.Error())
			}
			return nil
			// return c.Reply("感谢！✅私聊内容转投成功。")
		}
		return fmt.Errorf("获取图片信息失败")
	}

	what := fmt.Sprintf("@%s #id%d\n%s %s\nHi,Admin! 别逗了，宝！\n%s",
		c.Message().Sender.Username,
		c.Message().Sender.ID,
		c.Message().Sender.FirstName,
		c.Message().Sender.LastName,
		c.Message().Caption)
	return c.Reply(what)
}

// OnDocument
func OnPrivateSendToMeByDocument(c tele.Context) error {
	if c.Message().FromChannel() || c.Message().FromGroup() {
		return nil
	}
	if !c.Message().Private() && !c.Message().IsReply() {
		return nil
	}
	msgId := ""
	if len(c.Message().AlbumID) > 0 {
		msgId = fmt.Sprintf("%d_%s", c.Message().Chat.ID, c.Message().AlbumID)
	} else { // c.Message().ID > 0
		msgId = msgId + fmt.Sprintf("%d_%d", c.Message().Chat.ID, c.Message().ID)
	}
	if _, ok := syncMap.LoadOrStore(msgId, ""); ok {
		return nil
	}

	adminID := os.Getenv("SENDTOME_ID")
	if len(adminID) == 0 {
		return nil
	}
	senderID := fmt.Sprintf("%d", c.Message().Sender.ID)
	// 管理员回复信息
	if c.Message().IsReply() && strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到回复消息(err)：", c.Message())
		} else {
			fmt.Println("收到回复消息：", string(jsonText))
		}
		replyToText := c.Message().ReplyTo.Text
		if len(c.Message().ReplyTo.Caption) > 0 {
			replyToText = c.Message().ReplyTo.Caption
		}
		prefixLine, _, isFound := strings.Cut(replyToText+"\n", "\n")
		if !isFound {
			return c.Reply(
				fmt.Sprintf("回复消息格式异常(OnDocument)(prefixLine: %s): %+v", replyToText, c.Message().ReplyTo),
			)
		}
		_, sendToID, isFound := strings.Cut(prefixLine, "#id")
		if !isFound {
			return c.Reply("回复消息格式异常(OnDocument)(sendToID): " + fmt.Sprintf("%+v", c.Message().ReplyTo))
		}

		reciverId, err := strconv.ParseInt(sendToID, 10, 64)
		if err != nil {
			fmt.Print("回复消息格式异常(OnDocument): 待回复id %s\n%+v", sendToID, c.Message().ReplyTo)
		}
		reciver := &tele.User{
			ID: reciverId,
		}

		if c.Message().Document != nil {
			newMsg := c.Message().Document
			if len(c.Message().Caption) > 0 {
				newMsg.Caption = c.Message().Caption
			}
			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return c.Reply("⚠️回复内容(OnDocument)转投失败，请重试。" + err.Error())
			}
			return c.Reply("✅回复内容(OnDocument)转投成功。")
		}
		return c.Reply("⚠️回复内容转(OnDocument)投失败，请重试。" + fmt.Sprintf("获取图片信息失败: %+v", c.Message()))
	}

	// 收到私聊消息
	if c.Message().Private() && !strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到私聊消息(err)：", c.Message())
		} else {
			fmt.Println("收到私聊消息：", string(jsonText))
		}

		reciverId, err := strconv.ParseInt(adminID, 10, 64)
		if err != nil {
			fmt.Println("设置有误：环境变量(SENDTOME_ID)：", adminID)
		}
		reciver := &tele.User{
			ID: reciverId,
		}
		if c.Message().Document != nil {
			newMsg := c.Message().Document
			msgText := fmt.Sprintf("@%s #id%d\n%s",
				c.Message().Sender.Username,
				c.Message().Sender.ID,
				c.Message().Caption)
			newMsg.Caption = msgText

			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return err
				// return c.Reply("感谢！⚠️私聊内容转投失败。" + err.Error())
			}
			return nil
			// return c.Reply("感谢！✅私聊内容转投成功。")
		}
		return fmt.Errorf("获取图片信息失败")
	}

	what := fmt.Sprintf("@%s #id%d\n%s %s\nHi,Admin! 别逗了，宝！\n%s",
		c.Message().Sender.Username,
		c.Message().Sender.ID,
		c.Message().Sender.FirstName,
		c.Message().Sender.LastName,
		c.Message().Caption)
	return c.Reply(what)
}

// OnVideo
func OnPrivateSendToMeByVideo(c tele.Context) error {
	if c.Message().FromChannel() || c.Message().FromGroup() {
		return nil
	}
	if !c.Message().Private() && !c.Message().IsReply() {
		return nil
	}
	msgId := ""
	if len(c.Message().AlbumID) > 0 {
		msgId = fmt.Sprintf("%d_%s", c.Message().Chat.ID, c.Message().AlbumID)
	} else { // c.Message().ID > 0
		msgId = msgId + fmt.Sprintf("%d_%d", c.Message().Chat.ID, c.Message().ID)
	}
	if _, ok := syncMap.LoadOrStore(msgId, ""); ok {
		return nil
	}

	adminID := os.Getenv("SENDTOME_ID")
	if len(adminID) == 0 {
		return nil
	}
	senderID := fmt.Sprintf("%d", c.Message().Sender.ID)
	// 管理员回复信息
	if c.Message().IsReply() && strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到回复消息(err)：", c.Message())
		} else {
			fmt.Println("收到回复消息：", string(jsonText))
		}
		replyToText := c.Message().ReplyTo.Text
		if len(c.Message().ReplyTo.Caption) > 0 {
			replyToText = c.Message().ReplyTo.Caption
		}
		prefixLine, _, isFound := strings.Cut(replyToText+"\n", "\n")
		if !isFound {
			return c.Reply(
				fmt.Sprintf("回复消息格式异常(OnVideo)(prefixLine: %s): %+v", replyToText, c.Message().ReplyTo),
			)
		}
		_, sendToID, isFound := strings.Cut(prefixLine, "#id")
		if !isFound {
			return c.Reply("回复消息格式异常(OnVideo)(sendToID): " + fmt.Sprintf("%+v", c.Message().ReplyTo))
		}

		reciverId, err := strconv.ParseInt(sendToID, 10, 64)
		if err != nil {
			fmt.Print("回复消息格式异常(OnVideo): 待回复id %s\n%+v", sendToID, c.Message().ReplyTo)
		}
		reciver := &tele.User{
			ID: reciverId,
		}

		if c.Message().Video != nil {
			newMsg := c.Message().Video
			if len(c.Message().Caption) > 0 {
				newMsg.Caption = c.Message().Caption
			}
			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return c.Reply("⚠️回复内容(OnVideo)转投失败，请重试。" + err.Error())
			}
			return c.Reply(fmt.Sprintf("✅回复内容(OnVideo)转投成功,Message():%+v", c.Message()))
		}
		return c.Reply("⚠️回复内容(OnVideo)转投失败，请重试。" + fmt.Sprintf("获取图片信息失败: %+v", c.Message()))
	}

	// 收到私聊消息
	if c.Message().Private() && !strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到私聊消息(err)：", c.Message())
		} else {
			fmt.Println("收到私聊消息：", string(jsonText))
		}

		reciverId, err := strconv.ParseInt(adminID, 10, 64)
		if err != nil {
			fmt.Println("设置有误：环境变量(SENDTOME_ID)：", adminID)
		}
		reciver := &tele.User{
			ID: reciverId,
		}
		if c.Message().Video != nil {
			newMsg := c.Message().Video
			msgText := fmt.Sprintf("@%s #id%d\n%s",
				c.Message().Sender.Username,
				c.Message().Sender.ID,
				c.Message().Caption)
			newMsg.Caption = msgText

			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return err
				// return c.Reply("感谢！⚠️私聊内容转投失败。" + err.Error())
			}
			return nil
			// return c.Reply("感谢！✅私聊内容转投成功。")
		}
		return fmt.Errorf("获取图片信息失败")
	}

	what := fmt.Sprintf("@%s #id%d\n%s %s\nHi,Admin! 别逗了，宝！\n%s",
		c.Message().Sender.Username,
		c.Message().Sender.ID,
		c.Message().Sender.FirstName,
		c.Message().Sender.LastName,
		c.Message().Caption)
	return c.Reply(what)
}

// OnVoice
func OnPrivateSendToMeByVoice(c tele.Context) error {
	if c.Message().FromChannel() || c.Message().FromGroup() {
		return nil
	}
	if !c.Message().Private() && !c.Message().IsReply() {
		return nil
	}
	msgId := ""
	if len(c.Message().AlbumID) > 0 {
		msgId = fmt.Sprintf("%d_%s", c.Message().Chat.ID, c.Message().AlbumID)
	} else { // c.Message().ID > 0
		msgId = msgId + fmt.Sprintf("%d_%d", c.Message().Chat.ID, c.Message().ID)
	}
	if _, ok := syncMap.LoadOrStore(msgId, ""); ok {
		return nil
	}

	adminID := os.Getenv("SENDTOME_ID")
	if len(adminID) == 0 {
		return nil
	}
	senderID := fmt.Sprintf("%d", c.Message().Sender.ID)
	// 管理员回复信息
	if c.Message().IsReply() && strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到回复消息(err)：", c.Message())
		} else {
			fmt.Println("收到回复消息：", string(jsonText))
		}
		replyToText := c.Message().ReplyTo.Text
		if len(c.Message().ReplyTo.Caption) > 0 {
			replyToText = c.Message().ReplyTo.Caption
		}
		prefixLine, _, isFound := strings.Cut(replyToText+"\n", "\n")
		if !isFound {
			return c.Reply(
				fmt.Sprintf("回复消息格式异常(OnVoice)(prefixLine: %s): %+v", replyToText, c.Message().ReplyTo),
			)
		}
		_, sendToID, isFound := strings.Cut(prefixLine, "#id")
		if !isFound {
			return c.Reply("回复消息格式异常(OnVoice)(sendToID): " + fmt.Sprintf("%+v", c.Message().ReplyTo))
		}

		reciverId, err := strconv.ParseInt(sendToID, 10, 64)
		if err != nil {
			fmt.Print("回复消息格式异常(OnVoice): 待回复id %s\n%+v", sendToID, c.Message().ReplyTo)
		}
		reciver := &tele.User{
			ID: reciverId,
		}

		if c.Message().Voice != nil {
			newMsg := c.Message().Voice
			if len(c.Message().Caption) > 0 {
				newMsg.Caption = c.Message().Caption
			}
			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return c.Reply("⚠️回复内容(OnVoice)转投失败，请重试。" + err.Error())
			}
			return c.Reply("✅回复内容(OnVoice)转投成功。")
		}
		return c.Reply("⚠️回复内容(OnVoice)转投失败，请重试。" + fmt.Sprintf("获取图片信息失败: %+v", c.Message()))
	}

	// 收到私聊消息
	if c.Message().Private() && !strings.EqualFold(senderID, adminID) {
		if jsonText, err := json.Marshal(c.Message()); err != nil {
			fmt.Println("收到私聊消息(err)：", c.Message())
		} else {
			fmt.Println("收到私聊消息：", string(jsonText))
		}

		reciverId, err := strconv.ParseInt(adminID, 10, 64)
		if err != nil {
			fmt.Println("设置有误：环境变量(SENDTOME_ID)：", adminID)
		}
		reciver := &tele.User{
			ID: reciverId,
		}
		if c.Message().Voice != nil {
			newMsg := c.Message().Voice
			msgText := fmt.Sprintf("@%s #id%d\n%s",
				c.Message().Sender.Username,
				c.Message().Sender.ID,
				c.Message().Caption)
			newMsg.Caption = msgText

			if _, err := c.Bot().Send(reciver, newMsg); err != nil {
				return err
				// return c.Reply("感谢！⚠️私聊内容转投失败。" + err.Error())
			}
			return nil
			// return c.Reply("感谢！✅私聊内容转投成功。")
		}
		return fmt.Errorf("获取图片信息失败")
	}

	what := fmt.Sprintf("@%s #id%d\n%s %s\nHi,Admin! 别逗了，宝！\n%s",
		c.Message().Sender.Username,
		c.Message().Sender.ID,
		c.Message().Sender.FirstName,
		c.Message().Sender.LastName,
		c.Message().Caption)
	return c.Reply(what)
}
