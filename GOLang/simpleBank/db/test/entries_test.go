package db

import (
	"testing"
	"time"

	lib "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/lib"
	require "github.com/stretchr/testify/require"
)

// Test CreateEntry
func TestCreateEntry(t *testing.T) {
	account, _ := lib.CreateRandomAccount(testQueries)
	entry, err := lib.CreateRandomEntry(testQueries, account)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

// Test GetEntry
func TestGetEntry(t *testing.T) {
	account, _ := lib.CreateRandomAccount(testQueries)
	entry1, _ := lib.CreateRandomEntry(testQueries, account)
	entry2, err := lib.GetRandomEntry(testQueries, entry1)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

// Test ListEntries
func TestListEntries(t *testing.T) {
	account, _ := lib.CreateRandomAccount(testQueries)
	for i := 0; i < 10; i++ {
		lib.CreateRandomEntry(testQueries, account)
	}

	entries, err := lib.ListRandomEntries(testQueries, account.ID, 5, 0)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
