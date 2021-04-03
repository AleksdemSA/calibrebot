package main

//импорт нужных библиотек
import (
	"commands"
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// название файла-базы данных Calibre и токен как аргумент в строке запуска
var DB string = "metadata.db"
var BotAPI string = os.Args[1]

// основная функция
func main() {

	//проверяем наличие аргумента и нет ли пробелов в нём
	if len(os.Args) != 2 {
		fmt.Println("В качестве аргумента нужно указать только токен.")
		os.Exit(0)
	}

	// запускаем бота
	bot, err := tgbotapi.NewBotAPI(BotAPI)
	if err != nil {
		log.Panic(err)
	}

	// дебажим в консоль вывод
	bot.Debug = true
	// сразу пишем в консоль на каком боте авторизованы
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

			//если задан номер книги, ищем её
			i, err := strconv.Atoi(update.Message.Command())
			if err != nil {
				fmt.Println("Not a book number")
			}

			if i > 0 && i < 500000 {
				msg.Text = commands.GetBookDescription(DB, update.Message.Command())

				// вытаскиваем фото
				imagePath := commands.GetImage(DB, update.Message.Command())
				photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, imagePath)
				bot.Send(photo)

				//вытаскиваем файл
				bookPath := commands.GetBook(DB, update.Message.Command())
				book := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, bookPath)
				bot.Send(book)

			} else {
				//иначе обрабатываем команды
				switch update.Message.Command() {
				// поиск по автору
				case "author", "a":
					msg.Text = commands.SearchAuthor(DB, update.Message.CommandArguments())
				// помощь и старт
				case "help", "start", "h":
					msg.Text = commands.Help()
				// выводим последние 20 книг
				case "last", "l":
					msg.Text = commands.LastBook(DB)
				// выводим случайную книгу
				case "r":
					msg.Text = commands.RandBook(DB)
				// поиск книги по базе
				case "search", "s":
					if len(update.Message.CommandArguments()) < 3 {
						msg.Text = "Поиск не может состоять меньше чем из 3 символов. Набери /search или /s и любое выражение."
					} else {
						msg.Text = commands.SearchBook(DB, update.Message.CommandArguments())
					}
				// выводим статистику по боту и книгам
				case "stat":
					msg.Text = commands.Statistic(DB)
				// и ответ по-умолчанию, если комманда некорректна
				default:
					msg.Text = "Команда неизвестна, набери или нажми /help"
				}
			}

			bot.Send(msg)

		}
	}
}
