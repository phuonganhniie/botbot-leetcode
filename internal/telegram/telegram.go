package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

func GetAndStoreChatIds(token string, filePath string) (err error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", token)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.Errorf("Failed to create request for Telegram API: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Failed to get updates via Telegram API: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Telegram API returned non-OK status: %v", err)
	}

	body, _ := io.ReadAll(resp.Body)

	var teleResp model.TelegramResponse
	if err = json.Unmarshal(body, &teleResp); err != nil {
		return err
	}

	newChatIds := []int64{}
	for _, rs := range teleResp.Result {
		newChatIds = append(newChatIds, rs.Message.Chat.ID)
	}

	// Load existing chat IDs from the JSON file
	var existChatIds []int64
	if filePath != "" {
		file, err := os.Open(filePath)
		if os.IsNotExist(err) {
			os.Create(filePath)
			file.Chmod(0777)
		}
		if err == nil {
			defer file.Close()
			_ = json.NewDecoder(file).Decode(&existChatIds)
		}
		defer file.Close()
	}

	// If the file path is empty, use a default name
	if filePath == "" {
		logger.Infof("Warning: chatIDsFile is empty. Using default path: default_chat_ids.json")
		filePath = "default_chat_ids.json"
	}

	// Merge new chat IDs with existing ones, ensuring uniqueness
	allChatIds := append(existChatIds, newChatIds...)
	uniqueChatIds := uniqueChatIds(allChatIds)

	// Write unique chat IDs to the JSON file
	if len(newChatIds) > 0 {
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		file.Chmod(0777)
		defer file.Close()

		if err := json.NewEncoder(file).Encode(uniqueChatIds); err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}
	return nil
}

func LoadChatIds(filePath string) ([]int64, error) {
	if filePath == "" {
		logger.Infof("Warning: chatIDsFile is empty. Using default path: default_chat_ids.json")
		filePath = "default_chat_ids.json"
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	chatIds := []int64{}
	if err := json.NewDecoder(file).Decode(&chatIds); err != nil {
		return nil, fmt.Errorf("failed to decode chat IDs: %v", err)
	}

	return chatIds, nil
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
