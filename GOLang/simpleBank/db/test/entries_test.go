package db

import (
	"testing"
	"time"

	util "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/util"
	"github.com/stretchr/testify/require"
)

// Test CreateEntry
func TestCreateEntry(t *testing.T) {
	account, _ := util.CreateRandomAccount(testQueries)
	entry, err := util.CreateRandomEntry(testQueries, account)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

// Test GetEntry
func TestGetEntry(t *testing.T) {
	account, _ := util.CreateRandomAccount(testQueries)
	entry1, _ := util.CreateRandomEntry(testQueries, account)
	entry2, err := util.GetRandomEntry(testQueries, entry1)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

// Test ListEntries
func TestListEntries(t *testing.T) {
	account, _ := util.CreateRandomAccount(testQueries)
	for i := 0; i < 10; i++ {
		util.CreateRandomEntry(testQueries, account)
	}

	entries, err := util.ListRandomEntries(testQueries, account.ID, 5, 0)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
