package config

import (
	"os"

	"github.com/phuonganhniie/botbot-leetcode/internal/logger"
	"github.com/spf13/viper"
)

type Config struct {
	LeetCodeDailyURL string `mapstructure:"LEETCODE_DAILY_URL"`
	TelegramBotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TelegramChatID   string `mapstructure:"TELEGRAM_CHAT_ID"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if _, err := os.Stat("config/config.env"); os.IsNotExist(err) {
		logger.Info("No local .env file found. Relying on environment variables.")
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("config")

		if err := viper.ReadInConfig(); err != nil {
			logger.Errorf("Error reading config file: %v\n", err)
		} else {
			logger.Info("Config file loaded")
		}
	}

	viper.BindEnv("LEETCODE_DAILY_URL")
	viper.BindEnv("TELEGRAM_BOT_TOKEN")
	viper.BindEnv("TELEGRAM_CHAT_ID")

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
