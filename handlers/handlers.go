package handlers

import (
	"fmt"

	"github.com/max-messenger/maxbot"
)

func OnTextHandler(c maxbot.Context) error {
	text := c.Update().GetMessage().Body.Text

	if text == "hi" {
		return c.Send(SayHi())
	}

	if err := c.Send(fmt.Sprintf("%s - принято", text)); err != nil {
		return err
	}
	return nil
}

func SayHi() string {
	return "Ну привет челик..."
}

func SayHiHandler(c maxbot.Context) error {
	return c.Send("Ну привет, челик...")
}
