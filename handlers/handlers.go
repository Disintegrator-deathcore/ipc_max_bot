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
		AddCallBack("Информация о поступлении", CmdAboutAdmission)

	return c.Send(defaultAnswer, maxbot.WithKeyboard(kb))
}

func AboutCollege(c maxbot.Context) error {

	kb := model.NewKeyboard()
	kb.AddRow().
		AddCallBack("Контактные данные", CmdContacts).
		AddCallBack("Место нахождения", CmdLocation).
		AddLink("Карточка предприятия", enterpriseCardLink)
	kb.AddRow().
		AddCallBack("Режим и график работы", CmdWorkSchedule).
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

	return c.Send(productionPracticeText, maxbot.WithKeyboard(kb))
}

// Функции в разработке
func Contacts(c maxbot.Context) error {
	kb := model.NewKeyboard()
	kb.AddRow().AddCallBack("Назад", CmdAboutCollege)

	return c.Send("Функция в разработке", maxbot.WithKeyboard(kb))
}

func Location(c maxbot.Context) error {
	kb := model.NewKeyboard()
	kb.AddRow().AddCallBack("Назад", CmdAboutCollege)

	return c.Send("Функция в разработке", maxbot.WithKeyboard(kb))
}

func WorkSchedule(c maxbot.Context) error {
	kb := model.NewKeyboard()
	kb.AddRow().AddCallBack("Назад", CmdAboutCollege)

	return c.Send("Функция в разработке", maxbot.WithKeyboard(kb))
}

func AboutAdmission(c maxbot.Context) error {
	kb := model.NewKeyboard()
	kb.AddRow().AddCallBack("Назад", CmdMainInfo)

	return c.Send("Функция в разработке", maxbot.WithKeyboard(kb))
}
