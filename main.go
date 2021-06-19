package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	markovapi "github.com/jvdsouza/marcord/MarkovAPI"
	discordshell "github.com/jvdsouza/marcord/discordShell"
)

func main() {
	var markovapi markovapi.MarkovAPI
	botToken := os.Getenv("BOT_TOKEN")
	discordshell.LaunchShell(botToken, markovapi)
}
