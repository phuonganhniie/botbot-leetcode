package cmd

import (
	"strconv"

	"github.com/phuonganhniie/botbot-leetcode/config"
	"github.com/phuonganhniie/botbot-leetcode/internal/api"
	"github.com/phuonganhniie/botbot-leetcode/internal/format"
	"github.com/phuonganhniie/botbot-leetcode/internal/logger"
	"github.com/phuonganhniie/botbot-leetcode/internal/telegram"
)

func init() {
	logger.InitLogger()
}

func Start() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Errorf("Failed to load configuration: %v", err)
		return
	}

	// Get the LeetCode daily challenge
	challenge, err := api.FetchDailyChallenge(cfg.LeetCodeDailyURL)
	if err != nil {
		logger.Errorf("Fetch daily challenge error: %v", err)
		return
	}

	// Retrieve and store chat IDs
	chatIDsFile := cfg.TelegramChatIDsFilePath
	err = telegram.GetAndStoreChatIds(cfg.TelegramBotToken, chatIDsFile)
	if err != nil {
		logger.Errorf("Failed to get Telegram list ChatIDs: %v", err)
		return
	}

	// Load chat IDs from file
	chatIds, err := telegram.LoadChatIds(chatIDsFile)
	if err != nil {
		logger.Errorf("Failed to get Telegram list ChatIDs: %v", err)
		return
	}

	// Send the message to all chat IDs
	var formatter format.MessageFormatter
	formatter = &format.TelegramFormatter{}

	message := formatter.FormatMessage(&challenge)

	for _, chatId := range chatIds {
		chatIdStr := strconv.Itoa(int(chatId))
		err = telegram.SendChallenge(cfg.TelegramBotToken, chatIdStr, message)
		if err != nil {
			logger.Errorf("Failed to send Telegram message: %v", err)
			return
		}
	}
	logger.Info("Message sent successfully to Telegram")
}
