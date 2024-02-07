package states

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

// NoCommands Create a matcher which only matches text which is not a command.
func NoCommands(msg *gotgbot.Message) bool {
	return message.Text(msg) && !message.Command(msg)
}
