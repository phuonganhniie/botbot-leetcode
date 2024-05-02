package slack

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/phuonganhniie/botbot-leetcode/internal/logger"
	"github.com/phuonganhniie/botbot-leetcode/model"
)

func SendChallenge(challenge *model.Challenge, webhookUrl string) error {
	payload := map[string]string{
		"text": "Today's LeetCode Challenge: *" + challenge.Name + "* \nDifficulty: " + challenge.Difficulty + "\n" + challenge.QuestionLink,
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		logger.Errorf("Failed to marshal Slack payload: %v", err)
		return err
	}

	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		logger.Errorf("Failed to send Slack webhook: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("Slack webhook returned non-OK status: %v", err)
		return err
	}
	return nil
}
