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
	bot.Handle(maxbot.OnText, handlers.OnTextHandler)                            // Обрабатываем любые текстовые сообщения с помощью функции OnTextHandler
	bot.HandleCallback("/MainInfo", handlers.OnTextHandler)                      // Обрабатываем нажатие на кнопку "Назад" возвращая к дефолтному ответу
	bot.HandleCallback("/AboutCollege", handlers.AboutCollege)                   // Обрабатываем нажатие на кнопку "Информация о колледже"
	bot.HandleCallback("/EducationalActivities", handlers.EducationalActivities) // Обрабатываем нажатие на кнопку "Места осуществления образовательной деятельности"

	// Запускаем бота
	bot.Start()
}
