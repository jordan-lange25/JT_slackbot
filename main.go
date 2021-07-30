package main

import (
	"github.com/jordan-lange25/JT_slackbot/pkg/chunky"
)

func main() {
	token := ""

	chunky := chunky.Chunky{}

	chunky.InitalizeClient(token)

	chunky.PostMessage("#general", "Hello, wurld!")
}
