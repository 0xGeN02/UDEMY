package db

import (
	"testing"
	"time"

	util "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/util"
	"github.com/stretchr/testify/require"
)

// Test CreateTransfer
func TestCreateTransfer(t *testing.T) {
	account1, _ := util.CreateRandomAccount(testQueries)
	account2, _ := util.CreateRandomAccount(testQueries)
	transfer, err := util.CreateRandomTransfer(testQueries, account1, account2)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

// Test GetTransfer
func TestGetTransfer(t *testing.T) {
	account1, _ := util.CreateRandomAccount(testQueries)
	account2, _ := util.CreateRandomAccount(testQueries)
	transfer1, _ := util.CreateRandomTransfer(testQueries, account1, account2)
	transfer2, err := util.GetRandomTransfer(testQueries, transfer1)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

// Test ListTransfers
func TestListTransfers(t *testing.T) {
	account1, _ := util.CreateRandomAccount(testQueries)
	account2, _ := util.CreateRandomAccount(testQueries)
	for i := 0; i < 10; i++ {
		util.CreateRandomTransfer(testQueries, account1, account2)
	}

	transfers, err := util.ListRandomTransfers(testQueries, account1.ID, 5, 0)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
