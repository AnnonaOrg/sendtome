package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func SetTelegramWebhook(botToken string, webhookURL string) (retText string, err error) {
	apiUrl := "https://api.telegram.org/bot" + botToken + "/setWebhook?url="
	apiUrl = apiUrl + url.QueryEscape(webhookURL)
	log.Println(apiUrl)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		retText = string(body)
		return
	} else {
		err = fmt.Errorf("response status code err:%d", resp.StatusCode)
	}
	return
}
