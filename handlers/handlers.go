package handlers

import (
	"context"
	"log"
	"os"

	maxbotapi "github.com/max-messenger/max-bot-api-client-go/v2"
	"github.com/max-messenger/max-bot-api-client-go/v2/model"
	"github.com/max-messenger/maxbot"
)

// Универсальная функция для загрузки и отправки любого локального изображения
func sendLocalImage(ctx context.Context, api *maxbot.Api, chatID int64, text, filePath, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	token, err := api.Client().Upload.Upload(ctx, model.UploadType(model.AttachImage), file, fileName, fileInfo.Size())
	if err != nil {
		return err
	}

	msg := maxbotapi.NewMessage().SetChat(chatID).SetText(text).AddAttachByToken(token, model.AttachImage)
	_, err = api.Client().Messages.Send(ctx, msg)
	return err
}

// OnTextHandler обрабатывает любые текстовые сообщения, отправленные пользователем боту.
// В зависимости от текста сообщения, бот отвечает определенной фразой или подтверждает получение текста.
func OnTextHandler(c maxbot.Context) error {
	text := c.Update().GetMessage().Body.Text
	log.Print(text)
	defaultAnswer := "Я бот ИПК, вот что я умею"

	kb := model.NewKeyboard()
	kb.AddRow().AddCallBack("Отправить документы", "/Docs")

	switch text {
	case "привет":
		return c.Send(SayHi())
	case "пока":
		return c.Send(SayBye())
	case "как дела?":
		return c.Send(SayHowAreYou())
	default:
		if err := c.Send(defaultAnswer, maxbot.WithKeyboard(kb)); err != nil {
			return err
		}
	}

	return nil
}

func SayHi() string {
	return "Ну привет челик..."
}

func SayBye() string {
	return "Пока челик!"
}

func SayHowAreYou() string {
	return "у меня всё хорошо, а у тебя"
}

func StartHandler(c maxbot.Context) error {
	return c.Send("Здравствуйте, чем могу помочь?")
}

// InfoHandler отображает информацию о боте при отправке команды /info.
func InfoHandler(c maxbot.Context) error {
	kb := model.NewKeyboard()
	kb.AddRow().AddLink("docs", "https://dev.max.ru/docs")
	err := c.Send("max мне в руки", maxbot.WithKeyboard(kb))
	if err != nil {
		return err
	}
	return nil
}

func SendDocs(c maxbot.Context) error {
	text := "Какой документ необходимо отправить?"

	kb := model.NewKeyboard()

	kb.AddRow().
		AddCallBack("Паспорт", "/Passport").
		AddCallBack("СНИЛС", "/SNILS")

	return c.Send(text, maxbot.WithKeyboard(kb))
}

func SendPassport(c maxbot.Context, api *maxbot.Api) error {
	chatID := c.Update().GetMessage().Recipient.ChatID
	return sendLocalImage(c.Context(), api, chatID, "Отправляйте паспорт", "images/passport.jpg", "passport.jpg")
}

func SendSNILS(c maxbot.Context, api *maxbot.Api) error {
	chatID := c.Update().GetMessage().Recipient.ChatID
	return sendLocalImage(c.Context(), api, chatID, "Отправляйте СНИЛС", "images/snils.jpg", "snils.jpg")
}
