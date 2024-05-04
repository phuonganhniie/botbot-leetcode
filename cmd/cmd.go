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

	challenge, err := api.FetchDailyChallenge(cfg.LeetCodeDailyURL)
	if err != nil {
		logger.Errorf("Fetch daily challenge error: %v", err)
		return
	}

	chatIDs, err := telegram.GetChatIds(cfg.TelegramBotToken)
	if err != nil {
		logger.Errorf("Failed to get Telegram list ChatIDs: %v", err)
		return
	}

	var formatter format.MessageFormatter
	formatter = &format.TelegramFormatter{}

	message := formatter.FormatMessage(&challenge)

	for _, chatId := range chatIDs {
		chatIdStr := strconv.Itoa(int(chatId))
		err = telegram.SendChallenge(cfg.TelegramBotToken, chatIdStr, message)
		if err != nil {
			logger.Errorf("Failed to send Telegram message: %v", err)
			return
		}
	}
	logger.Info("Message sent successfully to Telegram")
}
