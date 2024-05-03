package cmd

import (
	"github.com/phuonganhniie/botbot-leetcode/config"
	"github.com/phuonganhniie/botbot-leetcode/internal/api"
	"github.com/phuonganhniie/botbot-leetcode/internal/format"
	"github.com/phuonganhniie/botbot-leetcode/internal/logger"
	"github.com/phuonganhniie/botbot-leetcode/internal/telegram"
)

func init() {
	logger.InitLogger()
	config.LoadConfig()
}

func Start() {
	challenge, err := api.FetchDailyChallenge(config.GetConfig().LeetCodeDailyURL)
	if err != nil {
		logger.Errorf("Fetch daily challenge error: %v", err)
		return
	}

	var formatter format.MessageFormatter
	formatter = &format.TelegramFormatter{}

	message := formatter.FormatMessage(&challenge)
	err = telegram.SendChallenge(config.GetConfig().TelegramBotToken, config.GetConfig().TelegramChatID, message)
	if err != nil {
		logger.Errorf("Failed to send Telegram message: %v", err)
		return
	}
	logger.Info("Message sent successfully to Telegram")
}
