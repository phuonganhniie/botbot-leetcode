package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	LeetCodeDailyURL string `mapstructure:"LEETCODE_DAILY_URL"`
	TelegramBotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TelegramChatID   string `mapstructure:"TELEGRAM_CHAT_ID"`
}

var config *Config

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	viper.BindEnv("LEETCODE_DAILY_URL")
	viper.BindEnv("TELEGRAM_BOT_TOKEN")
	viper.BindEnv("TELEGRAM_CHAT_ID")

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
