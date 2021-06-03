package main

import (
	"commands"
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var DB string = "metadata.db"
var BotAPI string = os.Args[1]

func main() {

	if len(os.Args) != 2 {
		fmt.Println("You need a token.")
		os.Exit(0)
	}

	bot, err := tgbotapi.NewBotAPI(BotAPI)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {

			i, err := strconv.Atoi(update.Message.Command())
			if err != nil {
				fmt.Println("Not a book number")
			}

			if i > 0 && i < 500000 {
				msg.Text = commands.GetBookDescription(DB, update.Message.Command())

				imagePath := commands.GetImage(DB, update.Message.Command())
				photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, imagePath)
				bot.Send(photo)

				bookPath := commands.GetBook(DB, update.Message.Command())
				book := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, bookPath)
				bot.Send(book)

			} else {
				switch update.Message.Command() {
				case "author", "a":
					msg.Text = commands.SearchAuthor(DB, update.Message.CommandArguments())
				case "help", "start", "h":
					msg.Text = commands.Help()
				case "last", "l":
					msg.Text = commands.LastBook(DB)
				case "r":
					msg.Text = commands.RandBook(DB)
				case "search", "s":
					if len(update.Message.CommandArguments()) < 3 {
						msg.Text = "Enter more whan 3 symbols"
					} else {
						msg.Text = commands.SearchBook(DB, update.Message.CommandArguments())
					}
				case "stat":
					msg.Text = commands.Statistic(DB)
				default:
					msg.Text = "Commant not found, press /help"
				}
			}

			bot.Send(msg)

		}
	}
}
