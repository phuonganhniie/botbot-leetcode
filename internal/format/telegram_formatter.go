package format

import (
	"fmt"
	"time"

	"github.com/phuonganhniie/botbot-leetcode/model"
)

type TelegramFormatter struct{}

func (f *TelegramFormatter) FormatMessage(challenge *model.Challenge) string {
	currentDate := time.Now().Format("January 2, 2006")
	return fmt.Sprintf("Gút mỏning, hôm nay là %s, làm nhẹ bài LeetCode nhá. Si sợ không? Không sợ thì chiến!\n\nID: %s\nVấn đề: %s\nĐộ căng: %s\nZô: %s",
		currentDate, challenge.Id, challenge.Name, challenge.Difficulty, challenge.QuestionLink)
}
