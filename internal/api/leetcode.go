package api

import (
	"encoding/json"
	"net/http"

	"github.com/phuonganhniie/botbot-leetcode/model"
)

func FetchDailyChallenge(apiUrl string) (model.Challenge, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return model.Challenge{}, err
	}
	defer resp.Body.Close()

	var challenge model.Challenge
	err = json.NewDecoder(resp.Body).Decode(&challenge)
	if err != nil {
		return model.Challenge{}, err
	}

	return challenge, nil
}
