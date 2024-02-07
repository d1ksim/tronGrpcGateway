package callbacks

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func (c *Client) ChangePriceCallback(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello, I'm @%s.\nWhat is your name?.", b.User.Username), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})

	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}

	c.SetUserData(ctx, "test", "hello")

	return handlers.NextConversationState("price")
}
