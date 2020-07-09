package core

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

type ChatWindow struct {
	messages []string
	currentMessage string
	sender MessageSender
}

func NewChatWindow(messages []string, sender MessageSender) ChatWindow {
	return ChatWindow{
		messages: messages,
		sender: sender,
		currentMessage: "",
	}
}

func (chat *ChatWindow) PrintMessage(message string) {
	chat.messages = append(chat.messages, message)
	chat.render()
}

func (chatWindow ChatWindow) render() {
	history := tui.NewVBox()

	for _, message := range chatWindow.messages {
		history.Append(tui.NewHBox(
			tui.NewLabel(message),
			tui.NewSpacer(),
		))
	}
	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)
	input.SetText(chatWindow.currentMessage)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(historyBox, input)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		chatWindow.sender.sendMessage(e.Text())
		input.SetText("")
	})

	input.OnChanged(func(e *tui.Entry) {
		chatWindow.currentMessage = e.Text()
	})

	ui, err := tui.New(chat)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
