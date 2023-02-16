package databases

import (
	"context"
	"database/sql"
	"fmt"

	databases "github.com/mmatz101/go-odds/databases/sqlc"
)

type Store struct {
	*databases.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: databases.New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*databases.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := databases.New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", rberr, err)
		}
	}

	return tx.Commit()
}
