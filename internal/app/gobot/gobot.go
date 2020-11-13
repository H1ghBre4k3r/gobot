package gobot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Gobot starts the bot.
func Gobot(Token string) {
	discord, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	discord.UpdateStatus(1, "Go-Kart")

	waitForProgramToClose()

	// Cleanly close down the Discord session.
	discord.Close()
}

// Handle messages.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == ",ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == ",pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == ",hi" {
		s.ChannelMessageSend(m.ChannelID, "Hey <@"+m.Author.ID+">!")
	}
}

// Wait here until CTRL-C or other term signal is received.
func waitForProgramToClose() {
	fmt.Println("Bot is now running. \nPress CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
