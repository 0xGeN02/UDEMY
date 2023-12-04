package db

import (
	"context"
	"testing"
	"time"

	lib "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/lib"
	require "github.com/stretchr/testify/require"
)

// Function to test CreateAccount
func TestCreateAccount(t *testing.T) {
	account, err := lib.CreateRandomAccount(testQueries)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account1, _ := lib.CreateRandomAccount(testQueries)
	err := lib.DeleteRandomAccount(testQueries, account1)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}

func TestGetAccount(t *testing.T) {
	account1, _ := lib.CreateRandomAccount(testQueries)
	account2, err := lib.GetRandomAccount(testQueries, account1)
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
		lib.CreateRandomAccount(testQueries)
	}

	accounts, err := lib.ListRandomAccounts(testQueries, 5, 0)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
