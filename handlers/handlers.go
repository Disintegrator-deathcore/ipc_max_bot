package handlers

import (
	"context"
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
	defaultAnswer := "Я бот ИПК, вот что я умею"

	kb := model.NewKeyboard()
	kb.AddRow().
		AddCallBack("Информация о колледже", "/AboutCollege").
		AddCallBack("Информация о поступлении", "/AboutAdmission")

	if err := c.Send(defaultAnswer, maxbot.WithKeyboard(kb)); err != nil {
		return err
	}

	return nil
}

// Функция обработки нажатия кнопки информации о колледже
func AboutCollege(c maxbot.Context) error {
	fullName := "\nПолное наименование ОО:\nБюджетное учреждение профессионального образования Ханты–Мансийского автономного округа–Югры «Игримский политехнический колледж»\n"
	abbriviatedName := "\nСокращенное наименование ОО:\nБУ «Игримский политехнический колледж»"
	text := "Основные сведения:\n"
	EnterpriseCardLink := "https://ipcollege.ru/wp-content/uploads/2025/11/КАРТОЧКА-ПРЕДПРИЯТИЯ.pdf"

	text += fullName + abbriviatedName // Соединяем весь текст о колледже в один

	kb := model.NewKeyboard()
	kb.AddRow().
		AddCallBack("Контактные данные", "/Contacts").
		AddCallBack("Место нахождения", "/Location").
		AddLink("Карточка предприятия", EnterpriseCardLink)
	kb.AddRow().
		AddCallBack("Режим и график работы", "/WorkSchedule").
		AddCallBack("Места осуществления образовательной деятельности", "/EducationalActivities")

	if err := c.Send(text, maxbot.WithKeyboard(kb)); err != nil {
		return err
	}

	return nil
}

func EducationalActivities(c maxbot.Context) error {
	text := "Образование"

	kb := model.NewKeyboard()

	kb.AddRow().
		AddCallBack("Основная", "/MainEducate").
		AddCallBack("Учебная практика", "/EducationalPractice").
		AddCallBack("Производственная практика", "/ProductionPractice")

	if err := c.Send(text, maxbot.WithKeyboard(kb)); err != nil {
		return err
	}

	return nil
}
