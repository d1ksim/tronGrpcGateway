package inline_query

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"math/rand"
	"strconv"
)

func AccountsList(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{gotgbot.InlineQueryResultArticle{
		Id:      strconv.Itoa(rand.Int()),
		Title:   "Bot Library",
		Url:     "github.com/PaulSonOfLars/gotgbot",
		HideUrl: true,
		InputMessageContent: gotgbot.InputTextMessageContent{
			MessageText: "Bot library source code:\ngithub.com/PaulSonOfLars/gotgbot",
		},
		Description: "Link to the bot source code",
	}}, &gotgbot.AnswerInlineQueryOpts{
		IsPersonal: true,
	})
	if err != nil {
		return fmt.Errorf("failed to send source ILQ: %w", err)
	}
	return nil
}
