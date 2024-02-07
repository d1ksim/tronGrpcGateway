package callbacks

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func AccountsCallback(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery

	switchInlineText := "accounts_list 1"

	_, err := ctx.EffectiveChat.SendMessage(b, "<b>Выберите действие</b>",
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{
						Text: "Список пользователей", SwitchInlineQueryCurrentChat: &switchInlineText,
					},
				}, {
					{
						Text: "Поиск по username", CallbackData: "search_by_username",
					},
				}},
			},
			ParseMode: "html",
		})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	_, err = cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{})
	if err != nil {
		return fmt.Errorf("failed to answer start callback query: %w", err)
	}
	return nil
}
