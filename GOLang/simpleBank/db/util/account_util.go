package util

import (
	"context"

	sqlc "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/sqlc"
)

// Function to create a random account
func CreateRandomAccount(queries *sqlc.Queries) (sqlc.Account, error) {
	arg := sqlc.CreateAccountParams{
		Owner:    RandomOwner(),
		Balance:  RandomMoney(),
		Currency: RandomCurrency(),
	}
	account, err := queries.CreateAccount(context.Background(), arg)
	if err != nil {
		return sqlc.Account{}, err
	}

	return account, nil
}

// Function to delete a random account
func DeleteRandomAccount(queries *sqlc.Queries, account sqlc.Account) error {
	err := queries.DeleteAccount(context.Background(), account.ID)
	if err != nil {
		return err
	}

	return nil
}

// Function to get a random account
func GetRandomAccount(queries *sqlc.Queries, account1 sqlc.Account) (sqlc.Account, error) {
	account2, err := queries.GetAccount(context.Background(), account1.ID)
	if err != nil {
		return sqlc.Account{}, err
	}

	return account2, nil
}

// Function to list random accounts
func ListRandomAccounts(queries *sqlc.Queries, limit, offset int32) ([]sqlc.Account, error) {
	arg := sqlc.ListAccountsParams{
		Limit:  limit,
		Offset: offset,
	}

	accounts, err := queries.ListAccounts(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
