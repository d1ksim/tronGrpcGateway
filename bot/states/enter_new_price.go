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
	"github.com/d1mpi/tronGrpcGateway/database"
	"github.com/d1mpi/tronGrpcGateway/database/models"
	"strconv"
)

const (
	PRICE = "price"
)

type CallbackStorage struct {
	storage callbacks.Client
}

func RegisterEnterNewPriceState(dispatcher *ext.Dispatcher) {
	client := CallbackStorage{}

	dispatcher.AddHandler(handlers.NewConversation(
		[]ext.Handler{handlers.NewCallback(callbackquery.Prefix("change_price"), client.storage.ChangePriceCallback)},
		map[string][]ext.Handler{
			PRICE: {handlers.NewMessage(NoCommands, client.price)},
		},
		&handlers.ConversationOpts{
			Exits:        []ext.Handler{handlers.NewCommand("cancel", commands.CancelCommand)},
			StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
			AllowReEntry: true,
		},
	))
}

func (c *CallbackStorage) price(b *gotgbot.Bot, ctx *ext.Context) error {
	accountId, _ := c.storage.GetUserData(ctx, "accountId")
	subscribePrice, _ := strconv.Atoi(ctx.EffectiveMessage.Text)

	var account models.Accounts
	database.DataBase.First(&account, accountId)
	database.DataBase.Model(&account).Update("subscribe_price", subscribePrice)

	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Пользователю @%s установлена новая плата за подписку: $%d", account.Username, subscribePrice), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send name message: %w", err)
	}

	return handlers.EndConversation()
}
