package util

import (
	"context"

	sqlc "github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/sqlc"
)

// Function to create a random entry
func CreateRandomEntry(queries *sqlc.Queries, account sqlc.Account) (sqlc.Entry, error) {
	arg := sqlc.CreateEntryParams{
		AccountID: account.ID,
		Amount:    RandomMoney(),
	}
	entry, err := queries.CreateEntry(context.Background(), arg)
	if err != nil {
		return sqlc.Entry{}, err
	}

	return entry, nil
}

// Function to get a random entry
func GetRandomEntry(queries *sqlc.Queries, entry1 sqlc.Entry) (sqlc.Entry, error) {
	entry2, err := queries.GetEntry(context.Background(), entry1.ID)
	if err != nil {
		return sqlc.Entry{}, err
	}

	return entry2, nil
}

// Function to list random entries
func ListRandomEntries(queries *sqlc.Queries, accountID int64, limit, offset int32) ([]sqlc.Entry, error) {
	arg := sqlc.ListEntriesParams{
		AccountID: accountID,
		Limit:     limit,
		Offset:    offset,
	}

	entries, err := queries.ListEntries(context.Background(), arg)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
