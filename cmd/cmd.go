package cmd

import (
	"github.com/phuonganhniie/botbot-leetcode/config"
	"github.com/phuonganhniie/botbot-leetcode/internal/api"
	"github.com/phuonganhniie/botbot-leetcode/internal/logger"
)

func init() {
	logger.InitLogger()
}

func Start() {
	_, err := config.LoadConfig("config")
	if err != nil {
		logger.Errorf("Load application config failed: %v", err)
		return
	}
	logger.Info("Load application config successfully")

	challenge, err := api.FetchDailyChallenge(config.GetConfig().LeetCodeDailyURL)
	if err != nil {
		logger.Errorf("Fetch daily challenge error: %v", err)
		return
	}
	logger.Infof("Today's LeetCode Challenge: %s * Difficulty: %s * Link: %s", challenge.Name, challenge.Difficulty, challenge.QuestionLink)
}
