package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/phuonganhniie/botbot-leetcode/internal/logger"
	"github.com/phuonganhniie/botbot-leetcode/model"
)

func uniqueChatIds(chatIds []int64) []int64 {
	seen := make(map[int64]bool)
	uniqueIds := []int64{}

	for _, id := range chatIds {
		if _, found := seen[id]; !found {
			uniqueIds = append(uniqueIds, id)
			seen[id] = true
		}
	}
	return uniqueIds
}

func GetChatIds(token string) (chatIds []int64, err error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", token)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Errorf("Failed to create request for Telegram API: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Failed to send message via Telegram API: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Telegram API returned non-OK status: %v", err)
	}

	body, _ := io.ReadAll(resp.Body)

	var teleResp model.TelegramResponse
	if err = json.Unmarshal(body, &teleResp); err != nil {
		return nil, err
	}

	for _, rs := range teleResp.Result {
		chatIds = append(chatIds, rs.Message.Chat.ID)
	}

	uniqueChatIds := uniqueChatIds(chatIds)
	return uniqueChatIds, nil
}

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

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
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
		return fmt.Errorf("Telegram API returned non-OK status: %v", err)
	}

	return nil
}
