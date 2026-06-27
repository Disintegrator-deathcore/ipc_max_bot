package handlers

import (
	"github.com/max-messenger/max-bot-api-client-go/v2/model"
	"github.com/max-messenger/maxbot"
)

func OnTextHandler(c maxbot.Context) error {
	defaultAnswer := "Я бот ИПК, вот что я умею"

	kb := model.NewKeyboard()
	kb.AddRow().
		AddCallBack("Информация о колледже", CmdAboutCollege).
		AddCallBack("Информация о поступлении", "/AboutAdmission")

	return c.Send(defaultAnswer, maxbot.WithKeyboard(kb))
}

func AboutCollege(c maxbot.Context) error {
	enterpriseCardLink := "https://ipcollege.ru/wp-content/uploads/2025/11/КАРТОЧКА-ПРЕДПРИЯТИЯ.pdf"

	kb := model.NewKeyboard()
	kb.AddRow().
		AddCallBack("Контактные данные", "/Contacts").
		AddCallBack("Место нахождения", "/Location").
		AddLink("Карточка предприятия", enterpriseCardLink)
	kb.AddRow().
		AddCallBack("Режим и график работы", "/WorkSchedule").
		AddCallBack("Места осуществления образовательной деятельности", CmdEducationalActivities).
		AddCallBack("Назад", CmdMainInfo)

	return c.Send(aboutCollegeText, maxbot.WithKeyboard(kb))
}

func EducationalActivities(c maxbot.Context) error {
	text := "Образование"

	kb := model.NewKeyboard()
	kb.AddRow().
		AddCallBack("Основная", CmdMainEducate).
		AddCallBack("Учебная практика", CmdEducationalPractice).
		AddCallBack("Производственная практика", CmdProductionPractice).
		AddCallBack("Назад", CmdAboutCollege)

	return c.Send(text, maxbot.WithKeyboard(kb))
}

func MainEducate(c maxbot.Context) error {
	kb := model.NewKeyboard()
	kb.AddRow().AddCallBack("Назад", CmdEducationalActivities)

	return c.Send(mainEducateText, maxbot.WithKeyboard(kb))
}

func EducationalPractice(c maxbot.Context) error {
	kb := model.NewKeyboard()
	kb.AddRow().AddCallBack("Назад", CmdEducationalActivities)

	return c.Send(educationalPracticeText, maxbot.WithKeyboard(kb))
}

func ProductionPractice(c maxbot.Context) error {
	kb := model.NewKeyboard()
	kb.AddRow().AddCallBack("Назад", CmdEducationalActivities)

	// Используем чистую константу из файла routes.go
	return c.Send(productionPracticeText, maxbot.WithKeyboard(kb))
}
