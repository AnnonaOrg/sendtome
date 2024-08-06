package utils

import (
	// "encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 自检openAPI服务是否正常运行
func CheckPingBaseURL(baseURL string) (retBool bool) {
	if !strings.HasPrefix(baseURL, "http") {
		return
	}
	apiURL := baseURL + "/ping"

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// var pingPong response.ExResponse
	// if err := json.Unmarshal(body, &pingPong); err != nil {
	// 	return
	// }
	// if pingPong.Code > 0 || pingPong.Message != "pong" {
	// 	return
	// }
	// if pingPong.Message == "pong" {
	// 	return true
	// }
	if strings.Contains(string(body), "pong") {
		return true
	} else {
		return false
	}
	return true
}
