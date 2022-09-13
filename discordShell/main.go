package discordshell

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type API interface {
	ProcessTextMessage(s string) string
}

func Launch(botToken string, api API) {
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal("Error creating new discord session: ", err)
	}

	discord.AddHandler(apiShell(api))

	err = discord.Open()
	if err != nil {
		log.Fatal("Error in opening discord session: ", err)
	}

	fmt.Println("Marcord is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func apiShell(api API) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		prefix := os.Getenv("BOT_PREFIX")
		message, _ := getMessage(s, m, prefix)

		processedMessage := api.ProcessTextMessage(message)
		if len(processedMessage) > 0 {
			sendMessage(s, m)(processedMessage)
		}
	}
}

func getMessage(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) (string, error) {
	if m.Author.ID == s.State.User.ID {
		return "", errors.New("Author is this bot")
	}

	if m.Content[0:len(prefix)] == prefix {
		return m.Content[len(prefix):len(m.Content)], nil
	}

	return "", nil
}

func sendMessage(s *discordgo.Session, m *discordgo.MessageCreate) func(m string) {
	return func(message string) {
		_, _ = s.ChannelMessageSend(m.ChannelID, message)
	}
}
