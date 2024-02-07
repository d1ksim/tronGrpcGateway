package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/choseninlineresult"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/inlinequery"
	cbs "github.com/d1mpi/tronGrpcGateway/bot/callbacks"
	cmds "github.com/d1mpi/tronGrpcGateway/bot/commands"
	"github.com/d1mpi/tronGrpcGateway/bot/inline_query"
	"github.com/d1mpi/tronGrpcGateway/bot/inline_result"
	"github.com/d1mpi/tronGrpcGateway/bot/states"
	"github.com/spf13/viper"
	"log"
	"time"
)

func InitTelegramBot() {
	bot, err := gotgbot.NewBot(viper.GetString("bot.token"), nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	// commands
	dispatcher.AddHandler(handlers.NewCommand("start", cmds.StartCommand))

	// callbacks
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("accounts"), cbs.AccountsCallback))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Prefix("ban_user"), cbs.BanUserCallback))

	// inline query
	dispatcher.AddHandler(handlers.NewInlineQuery(inlinequery.QueryPrefix("accounts_list"), inline_query.AccountsList))

	// chosen inline result
	dispatcher.AddHandler(handlers.NewChosenInlineResult(choseninlineresult.All, inline_result.AccountsList))

	// states
	states.RegisterEnterNewPriceState(dispatcher)

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("Bot %s has been started...\n", bot.User.Username)
}
