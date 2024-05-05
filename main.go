package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/phuonganhniie/botbot-leetcode/cmd"
)

func main() {
	lambda.Start(cmd.Start)
}
