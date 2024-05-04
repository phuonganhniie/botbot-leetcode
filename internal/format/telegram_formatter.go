package format

import (
	"fmt"
	"time"

	"github.com/phuonganhniie/botbot-leetcode/model"
)

type TelegramFormatter struct{}

func (f *TelegramFormatter) FormatMessage(challenge *model.Challenge) string {
	currentDate := time.Now().Format("January 2, 2006")
	return fmt.Sprintf("Gút mỏning, today is %s, and time to s(t)uck with this problem. But don't worry, you'll be fine, and you'll be learnt.\n\nID: %s\nProblem: %s\nDifficulty: %s\nLet's be in: %s",
		currentDate, challenge.Id, challenge.Name, challenge.Difficulty, challenge.QuestionLink)
}
