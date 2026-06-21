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

// До запуска основной функции подгружаем переменные окружения
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден, используются переменные окружения из системы")
	}
}

func main() {
	// Настраиваем HTTP клиент с таймаутом для запросов к API
	opts := []maxbot.Opt{
		maxbot.WithHTTPClient(&http.Client{Timeout: 25 * time.Second}),
	}

	// Получаем токен бота из переменных окружения
	token := os.Getenv("TOKEN")

	// Создаем экземпляр бота с указанным токеном и опциями
	bot, err := maxbot.NewApi(token, opts...)
	if err != nil {
		log.Fatal(err)
	}

	// Регистрируем обработчики для команд и текстовых сообщений

	bot.Handle("/start", handlers.StartHandler) // Отображение текста при старте бота
	bot.Handle("/info", handlers.InfoHandler)   // Отображаем информацию о боте по команде /info
	bot.HandleCallback("/Docs", handlers.SendDocs)
	// bot.HandleCallback("/Passport", handlers.SendPassport)
	bot.HandleCallback("/Passport", func(c maxbot.Context) error {
		// Внутри этой функции мы можем использовать переменную bot из main.go!
		return handlers.SendPassport(c, bot)
	})
	bot.HandleCallback("/SNILS", func(c maxbot.Context) error {
		return handlers.SendSNILS(c, bot)
	})
	bot.Handle(maxbot.OnText, handlers.OnTextHandler) // Обрабатываем любые текстовые сообщения с помощью функции OnTextHandler

	// Запускаем бота
	bot.Start()
}
