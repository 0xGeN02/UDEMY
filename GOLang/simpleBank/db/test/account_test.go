package db

import (
	"context"
	"testing"
	"time"

	sqlc "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/sqlc"
	util "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/util"
	"github.com/stretchr/testify/require"
)

// Function to create a random account to use as test
func createRandomAccount(t *testing.T) sqlc.Account {
	arg := sqlc.CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// Function to test GetAccount
func GetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// Function to test CreateAccount
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// Function to test UpdateAccount
func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := sqlc.UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	acccount2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acccount2)

	require.Equal(t, account1.ID, acccount2.ID)
	require.Equal(t, account1.Owner, acccount2.Owner)
	require.Equal(t, arg.Balance, acccount2.Balance)
	require.Equal(t, account1.Currency, acccount2.Currency)
	require.WithinDuration(t, account1.CreatedAt, acccount2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := sqlc.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
