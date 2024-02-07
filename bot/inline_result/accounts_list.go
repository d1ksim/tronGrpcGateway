package inline_result

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/d1mpi/tronGrpcGateway/database"
	"github.com/d1mpi/tronGrpcGateway/database/models"
)

func AccountsList(b *gotgbot.Bot, ctx *ext.Context) error {
	var account models.Accounts
	database.DataBase.First(&account, ctx.ChosenInlineResult.ResultId)

	payment := "Оплаты отсутствуют"
	if account.LastPayment != nil {
		payment = account.LastPayment.String()
	}

	_, err := b.SendMessage(ctx.ChosenInlineResult.From.Id, fmt.Sprintf(
		"<b>Пользователь</b> @%s\n"+
			"<b>Telegram ID:</b> <code>%d</code>\n\n"+
			"<b>Последняя оплата:</b> %s\n"+
			"<b>Оплачивает в месяц:</b> $%d",
		account.Username, account.TelegramID, payment, account.SubscribePrice,
	),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{
						Text: "Заблокировать/разблокировать", CallbackData: fmt.Sprintf("ban_user:%s", ctx.ChosenInlineResult.ResultId),
					},
				}, {
					{
						Text: "Изменить абонку", CallbackData: fmt.Sprintf("change_price:%s", ctx.ChosenInlineResult.ResultId),
					},
				}},
			},
			ParseMode: "html",
		})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
