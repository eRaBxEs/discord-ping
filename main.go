package main

import (
	"fmt"

	"github.com/erabxes/discord-ping/bot"
	"github.com/erabxes/discord-ping/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
