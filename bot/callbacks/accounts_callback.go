package callbacks

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func AccountsCallback(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery

	_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: "You pressed a button!",
	})
	if err != nil {
		return fmt.Errorf("failed to answer start callback query: %w", err)
	}

	_, _, err = cb.Message.EditText(b, "You edited the start message.", nil)
	if err != nil {
		return fmt.Errorf("failed to edit start message text: %w", err)
	}
	return nil
}
