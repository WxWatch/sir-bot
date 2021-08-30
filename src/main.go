package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"wxwatch.dev/bot/src/discord"
	"wxwatch.dev/bot/src/logx"
	"wxwatch.dev/bot/src/server"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	logger := logx.GetLogger()

	discordOptions := &discord.Options{
		Token: Token,
	}

	options := &server.Options{
		DiscordOptions: discordOptions,
	}

	server := server.NewServer(options)

	err := server.Start()
	if err != nil {
		logger.Infof("Error starting server: ", err)
		return
	}

	// Wait here until sigterm
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	server.Stop()
}
