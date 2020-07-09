package core

import (
	"fmt"
	"log"
	"github.com/marcusolsson/tui-go"
)

type ChatWindow struct {
	sender MessageSender
}

func NewChatWindow(sender MessageSender) ChatWindow {
	return ChatWindow{
		sender: sender,
	}
}

func (chatWindow ChatWindow) Init(messageProcessor *MessageReceivedCallback) {
	history := tui.NewVBox()
	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)
	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	messageProcessor.messageBox = history

	inputBox := chatWindow.configureInput(messageProcessor)

	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	ui, err := tui.New(chat)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func (chatWindow ChatWindow) configureInput(callback *MessageReceivedCallback) *tui.Box {
	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	input.OnSubmit(func(e *tui.Entry) {
		chatWindow.sender.sendMessage(e.Text())
		callback.OnMessageReceived(fmt.Sprintf("%s", e.Text()))
		input.SetText("")
	})

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)
	return inputBox
}
