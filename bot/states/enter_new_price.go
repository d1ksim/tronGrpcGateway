package states

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/d1mpi/tronGrpcGateway/bot/callbacks"
	"github.com/d1mpi/tronGrpcGateway/bot/commands"
	"html"
)

const (
	PRICE = "price"
)

func RegisterEnterNewPriceState(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandler(handlers.NewConversation(
		[]ext.Handler{handlers.NewCallback(callbackquery.Prefix("change_price"), callbacks.ChangePriceCallback)},
		map[string][]ext.Handler{
			PRICE: {handlers.NewMessage(NoCommands, price)},
		},
		&handlers.ConversationOpts{
			Exits:        []ext.Handler{handlers.NewCommand("cancel", commands.CancelCommand)},
			StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
			AllowReEntry: true,
		},
	))
}

func price(b *gotgbot.Bot, ctx *ext.Context) error {
	inputName := ctx.EffectiveMessage.Text
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Nice to meet you, %s!\n\nAnd how old are you?", html.EscapeString(inputName)), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send name message: %w", err)
	}

	println(ctx.Data["x"])

	return handlers.EndConversation()
}
