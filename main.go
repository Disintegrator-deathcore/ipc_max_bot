package main

import (
	"log"
	"net/http"
	"os"
	"test_bot/handlers"
	"time"

	"github.com/joho/godotenv"
	"github.com/max-messenger/max-bot-api-client-go/v2/model"
	"github.com/max-messenger/maxbot"
)

// До запуска основной функции подгружаем переменные окружения
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}

func main() {
	opts := []maxbot.Opt{
		maxbot.WithHTTPClient(&http.Client{Timeout: 25 * time.Second}),
	}

	token := os.Getenv("TOKEN")

	bot, err := maxbot.NewApi(token, opts...)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/info", func(c maxbot.Context) error {
		kb := model.NewKeyboard()
		kb.AddRow().AddLink("docs", "https://dev.max.ru/docs")
		err = c.Send("max мне в руки", maxbot.WithKeyboard(kb))
		if err != nil {
			return err
		}

		return nil
	})

	bot.Handle(maxbot.OnText, handlers.OnTextHandler)

	bot.Start()
}
