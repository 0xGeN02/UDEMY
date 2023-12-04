package util

import (
	"context"

	sqlc "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/sqlc"
	util "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/util"
)

// Function to create a random transfer
func CreateRandomTransfer(queries *sqlc.Queries, fromAccount, toAccount sqlc.Account) (sqlc.Transfer, error) {
	arg := sqlc.CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := queries.CreateTransfer(context.Background(), arg)
	if err != nil {
		return sqlc.Transfer{}, err
	}

	return transfer, nil
}

// Function to get a random transfer
func GetRandomTransfer(queries *sqlc.Queries, transfer1 sqlc.Transfer) (sqlc.Transfer, error) {
	transfer2, err := queries.GetTransfer(context.Background(), transfer1.ID)
	if err != nil {
		return sqlc.Transfer{}, err
	}

	return transfer2, nil
}

// Function to list random transfers
func ListRandomTransfers(queries *sqlc.Queries, fromAccountID int64, limit, offset int32) ([]sqlc.Transfer, error) {
	arg := sqlc.ListTransfersParams{
		FromAccountID: fromAccountID,
		Limit:         limit,
		Offset:        offset,
	}

	transfers, err := queries.ListTransfers(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	return transfers, nil
}
