package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTelegramHandleFromUserMS(userId string) (string, error) {
	var telegramHandle string

	// oops... its hardcoded, abit lazy
	userTeleEndpoint := fmt.Sprintf("http://user-ms:3005/api/users/telegram?userId=%v", userId)

	resp, err := http.Get(userTeleEndpoint)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var responseBody struct {
		TelegramHandle string `json:"telegram_handle"`
	}
	if err = json.Unmarshal(body, &responseBody); err != nil {
		return "", err
	}

	telegramHandle = responseBody.TelegramHandle

	return telegramHandle, nil
}
