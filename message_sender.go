package telesend

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	sendMsgURL = "https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s"
)

// BotError defines struct if bot error sending message to telegram
type BotError struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

// Bot defines struct for telegram bot
type Bot struct {
	token  string
	chatID string
}

// NewBot will create new instance of bot in a certain telegram group or channel
func NewBot(newToken, newChatID string) *Bot {
	return &Bot{
		token:  newToken,
		chatID: newChatID,
	}
}

// SendMessage will send message to private telegram group or channel through bot
func (b *Bot) SendMessage(v interface{}) error {
	msg := toString(v)
	url := fmt.Sprintf(sendMsgURL, b.token, b.chatID, msg)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusBadRequest {
		bodyStr := string(body)
		errData := &BotError{}

		err = json.Unmarshal([]byte(bodyStr), &errData)
		if err != nil {
			return err
		}
		return errors.New(errData.Description)
	}
	return nil
}
