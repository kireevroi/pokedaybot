package main

import (
    "github.com/Syfaro/telegram-bot-api"
    "os"
		"github.com/joho/godotenv"
		"log"
)

func telegramBot() {
		err := godotenv.Load("token.env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
    if err != nil {
        panic(err)
    }
    //Устанавливаем время обновления
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    //Получаем обновления от бота
    updates, err := bot.GetUpdatesChan(u)
		if err == nil {
			log.Println("No error");
		}
    for update := range updates {
        if update.Message == nil {
            continue
        }
        //Проверяем что от пользователья пришло именно текстовое сообщение
        if update.Message.Text != "" {
					log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				}
    }
}

func main() {
    telegramBot()
}