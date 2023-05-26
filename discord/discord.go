package discord

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zahraayafni/sample-discord-bot/calculator"
	"github.com/zahraayafni/sample-discord-bot/tinyurl"
)

func InitDiscord() (*discordgo.Session, error) {
	// initialize discord client
	discordToken := "Bot " + os.Getenv("DISCORD_TOKEN")
	sess, err := discordgo.New(discordToken)
	if err != nil {
		return nil, err
	}

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	return sess, nil
}

func RunDiscordHandler(sess *discordgo.Session) {
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		switch {
		case m.Content == "hello":
			s.ChannelMessageSend(m.ChannelID, "world!")
		case strings.HasPrefix(m.Content, "hitung") && strings.ContainsAny(m.Content, "+/-*"):
			input := strings.Replace(m.Content, "hitung ", "", -1)
			op, num1, num2, err := calculator.GenerateInput(input)
			if err != nil {
				log.Println("ERROR:", err)
				s.ChannelMessageSend(m.ChannelID, "try to input the correct number operation")
			} else {
				res := calculator.BasicCalculator(op, num1, num2)
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("result is %d", res))
			}
		case strings.HasPrefix(m.Content, "createShortLink"):
			commands := strings.Split(m.Content, " ")
			if len(commands) < 2 {
				s.ChannelMessageSend(m.ChannelID, "try to input the correct command")
			} else if len(commands[1]) > 0 {
				shortLink, err := tinyurl.CreateShortLink(commands[1])
				if err != nil {
					log.Println("ERROR:", err)
					s.ChannelMessageSend(m.ChannelID, "sorry, we fail to process your command. please try again!")
				} else {
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("here is your short link : %s", shortLink))
				}
			}
		}
	})
}
