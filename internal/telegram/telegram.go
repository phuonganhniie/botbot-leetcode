package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/phuonganhniie/botbot-leetcode/internal/logger"
)

func SendChallenge(token string, chatID string, messageText string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload := map[string]interface{}{
		"chat_id": chatID,
		"text":    messageText,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		logger.Errorf("Failed to marshal message payload: %v", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		logger.Errorf("Failed to create request for Telegram API: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Failed to send message via Telegram API: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("Telegram API returned non-OK status: %v", err)
		return fmt.Errorf("Telegram API returned non-OK status: %s", resp.Status)
	}

	return nil
}
