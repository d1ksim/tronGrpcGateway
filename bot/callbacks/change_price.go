package callbacks

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"strconv"
	"strings"
)

func (c *Client) ChangePriceCallback(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	accountId, _ := strconv.Atoi(strings.Split(cb.Data, ":")[1])

	_, err := ctx.EffectiveMessage.Reply(b, "<b>Пожалуйста, введите новую цену за подписку:</b>", &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})

	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	c.SetUserData(ctx, "accountId", accountId)
	return handlers.NextConversationState("price")
}
