package tg

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const api = `https://api.telegram.org`

// Client for telegram bot
type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
	}
}

// SendMsg send message in chat
func (c *Client) SendMsg(chatID int64, text string) error {
	type req struct {
		ChatID int64  `json:"chat_id"`
		Text   string `json:"text"`
	}

	reqBody := req{
		ChatID: chatID,
		Text:   text,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	res, err := http.Post("https://api.telegram.org/bot"+c.token+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
