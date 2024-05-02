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

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

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

func GetConfig() *Config {
	if config == nil {
		panic("load config first")
	}
	return config
}
