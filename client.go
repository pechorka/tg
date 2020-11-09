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
	token    string
	parseMod ParseMod
}

// Options is options for tg client
type Options struct {
	// ParseMod which parse mod to use. Optional
	ParseMod ParseMod
}

// NewClient create new tg client
func NewClient(token string, opts *Options) *Client {
	c := Client{
		token: token,
	}
	if opts != nil {
		c.parseMod = opts.ParseMod
	}
	return &c
}

// SendMsg send message in chat
func (c *Client) SendMsg(chatID int64, text string) error {
	req := map[string]interface{}{
		"chat_id":    chatID,
		"text":       text,
		"parse_mode": c.parseMod,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	res, err := http.Post("https://api.telegram.org/bot"+c.token+"/sendMessage", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
