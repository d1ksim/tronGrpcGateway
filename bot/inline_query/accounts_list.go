package inline_query

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/d1mpi/tronGrpcGateway/database"
	"github.com/d1mpi/tronGrpcGateway/database/models"
	"strconv"
)

func AccountsList(b *gotgbot.Bot, ctx *ext.Context) error {
	var usersResult []gotgbot.InlineQueryResult
	var accounts []models.Accounts

	database.DataBase.Find(&accounts)

	for _, v := range accounts {
		usersResult = append(usersResult, gotgbot.InlineQueryResultArticle{
			Id:    strconv.Itoa(int(v.ID)),
			Title: fmt.Sprintf("@%s (ID: %d)", v.Username, v.TelegramID),
			InputMessageContent: gotgbot.InputTextMessageContent{
				MessageText: "Получение данных пользователя...",
			},
		})
	}

	_, err := ctx.InlineQuery.Answer(b, usersResult, &gotgbot.AnswerInlineQueryOpts{
		IsPersonal: false,
		CacheTime:  15,
	})
	if err != nil {
		return fmt.Errorf("failed to send source ILQ: %w", err)
	}
	return nil
}
