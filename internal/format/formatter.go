package format

import "github.com/phuonganhniie/botbot-leetcode/model"

type MessageFormatter interface {
	FormatMessage(challenge *model.Challenge) string
}
