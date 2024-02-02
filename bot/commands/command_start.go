package commands

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func StartCommand(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf(
		"Доброго времени суток @%s\n\n"+
			"<b>Этот бот является панелью управления Android Wallet приложением.\n"+
			"Пожалуйста, выберите действие.</b>", ctx.Message.From.Username),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{
						Text: "Аккаунты", CallbackData: "accounts",
					},
					{
						Text: "Статистика", CallbackData: "statistics",
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
