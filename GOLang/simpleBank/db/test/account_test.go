package db

import (
	"context"
	"testing"
	"time"

	util "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/util"
	"github.com/stretchr/testify/require"
)

// Function to test CreateAccount
func TestCreateAccount(t *testing.T) {
	account, err := util.CreateRandomAccount(testQueries)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account1, _ := util.CreateRandomAccount(testQueries)
	err := util.DeleteRandomAccount(testQueries, account1)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}

func TestGetAccount(t *testing.T) {
	account1, _ := util.CreateRandomAccount(testQueries)
	account2, err := util.GetRandomAccount(testQueries, account1)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// Function to test ListAccounts
func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		util.CreateRandomAccount(testQueries)
	}

	accounts, err := util.ListRandomAccounts(testQueries, 5, 0)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
