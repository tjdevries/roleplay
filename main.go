package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting Bot...")

	var args struct {
		ChannelID string `arg:"positional"`
	}
	arg.MustParse(&args)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed with error: %v", err)
	}

	message := "Hello I'm a Bot"

	if args.ChannelID == "" {
		panic("You must set the channelID variable above! Right-Click the Channel in Discord and hit 'Copy ID'")
	}

	if message == "" {
		panic("You must set the message variable above! What do you want your bot to say!")
	}

	// First Look to see if the environment variable is already set
	botToken := os.Getenv("DISCORD_BOT_TOKEN")

	// If the Discord Bot Token is not set, then try and load the .env file
	if botToken == "" {
		fmt.Println("Loading DISCORD_BOT_TOKEN from .env file")
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Could not find .env file with DISCORD_BOT_TOKEN Set. Error: %s", err.Error())
		}
		botToken = os.Getenv("DISCORD_BOT_TOKEN")
	}

	if botToken == "" {
		panic("You must set your DISCORD_BOT_TOKEN environment variable. See README.md for details")
	}

	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	dg.Identify.Intents = discordgo.IntentsAll

	// if err := dg.Open(); err != nil {
	// 	log.Fatalf("Bad connection: %v", err)
	// }
	defer dg.Close()

	_, err = dg.ChannelMessageSend(args.ChannelID, message)
	if err != nil {
		fmt.Printf("Error Sending Channel Message: %s", err.Error())
	}

	guildID := os.Getenv("GUILD_ID")
	userID := os.Getenv("USER_ID")
	roleID := "953754686345846854"

	// Hey, it me
	st, err := dg.GuildMember(guildID, userID)
	if err != nil {
		log.Fatalf("Bad Build: %v", err)
	}

	fmt.Println("Guild Member Roles:", st.Roles)
	fmt.Println(dg.GuildMemberRoleAdd(guildID, userID, roleID))
}
