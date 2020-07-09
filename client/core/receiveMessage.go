package core

import (
	"github.com/marcusolsson/tui-go"
)

type MessageReceivedCallback struct {
	messageBox *tui.Box
}

func (callback MessageReceivedCallback) OnMessageReceived(message string) {
	callback.messageBox.Append(tui.NewHBox(
		tui.NewLabel(message),
		tui.NewSpacer(),
	))
}

