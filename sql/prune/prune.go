package prune

import (
	"context"
	"database/sql"
	"fmt"

	sfom_sql "github.com/sfomuseum/go-database/sql"
	"github.com/whosonfirst/go-whosonfirst-iterate/v3"
	"github.com/whosonfirst/go-whosonfirst-uri"
)

// PruneTables will remove all the records in 'to_prune'
func PruneTables(ctx context.Context, db *sql.DB, to_prune ...sfom_sql.Table) error {

	tx, err := db.Begin()

	if err != nil {
		return fmt.Errorf("Failed create transaction, because %w", err)
	}

	for _, t := range to_prune {

		sql := fmt.Sprintf("DELETE FROM %s", t.Name())
		stmt, err := tx.Prepare(sql)

		if err != nil {
			return fmt.Errorf("Failed to prepare statement (%s), because %w", sql, err)
		}

		_, err = stmt.Exec()

		if err != nil {
			return fmt.Errorf("Failed to execute statement (%s), because %w", sql, err)
		}
	}

	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("Failed to commit transaction, because %w", err)
	}

	return nil
}

// PruneTablesWithIterator will remove records emitted by an iterator (defined by 'iterator_uri' and 'iterator_source') from 'to_prune'.
func PruneTablesWithIterator(ctx context.Context, iterator_uri string, iterator_source string, db *sql.DB, to_prune ...sfom_sql.Table) error {

	iter, err := iterate.NewIterator(ctx, iterator_uri)

	if err != nil {
		return fmt.Errorf("Failed to create iterator, %v", err)
	}

	for rec, err := range iter.Iterate(ctx, iterator_source) {

		if err != nil {
			return fmt.Errorf("Failed to iterate URIs, %v", err)
		}

		defer rec.Body.Close()

		id, _, err := uri.ParseURI(rec.Path)

		if err != nil {
			return fmt.Errorf("Failed to parse %s, %w", rec.Path, err)
		}

		tx, err := db.Begin()

		if err != nil {
			return fmt.Errorf("Failed create transaction for pruning %d, because %w", id, err)
		}

		for _, t := range to_prune {

			sql := fmt.Sprintf("DELETE FROM %s WHERE id = ?", t.Name())
			stmt, err := tx.Prepare(sql)

			if err != nil {
				return fmt.Errorf("Failed to prepare statement (%s), because %w", sql, err)
			}

			_, err = stmt.Exec(id)

			if err != nil {
				return fmt.Errorf("Failed to execute statement (%s, %d), because %w", sql, id, err)
			}
		}

		err = tx.Commit()

		if err != nil {
			fmt.Errorf("Failed to commit transaction to pruning %d, because %w", id, err)
		}

		return nil
	}

	return nil
}
