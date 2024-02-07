package callbacks

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/d1mpi/tronGrpcGateway/database"
	"github.com/d1mpi/tronGrpcGateway/database/models"
	"strconv"
	"strings"
)

func BanUserCallback(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	accountId, _ := strconv.Atoi(strings.Split(cb.Data, ":")[1])

	var account models.Accounts
	database.DataBase.First(&account, accountId)
	database.DataBase.Model(&account).Update("active", !account.Active)

	activeMessage := "отключена"
	if !account.Active == false {
		activeMessage = "активирована"
	}

	_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: fmt.Sprintf("Запись успешно %s", activeMessage),
	})
	if err != nil {
		return fmt.Errorf("failed to answer start callback query: %w", err)
	}
	return nil
}
