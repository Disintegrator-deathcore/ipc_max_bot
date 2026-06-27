package main

import (
	"log"
	"net/http"
	"os"
	"test_bot/handlers"
	"time"

	"github.com/joho/godotenv"
	"github.com/max-messenger/maxbot"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден, используются переменные окружения из системы")
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

	// Выносим всю регистрацию в одну чистую функцию
	handlers.RegisterAll(bot)

	log.Println("Бот ИПК успешно запущен...")
	bot.Start()
}
