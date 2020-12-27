package lark

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (c *Client) BotHook(ctx context.Context, hookId string, text *TextMessage, post *PostMessage, card *CardMessage) (*BotResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", CustomBotHost, "bot/v2/hook", hookId)
	var body []byte
	var err error
	if text != nil {
		if body, err = json.Marshal(text); err != nil {
			return nil, err
		}
	} else if post != nil {
		if body, err = json.Marshal(post); err != nil {
			return nil, err
		}
	} else if card != nil {
		if body, err = json.Marshal(card); err != nil {
			return nil, err
		}
	}
	resp, err := c.Post(ctx, url, body, time.Duration(2)*time.Second)
	if err != nil {
		return nil, err
	}
	fmt.Printf("resp: %s", resp.Body())
	return nil, nil
}
