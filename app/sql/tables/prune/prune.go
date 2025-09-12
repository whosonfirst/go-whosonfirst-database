package prune

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	_ "github.com/whosonfirst/go-whosonfirst-database/sql"
	
	database_sql "github.com/sfomuseum/go-database/sql"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-database/sql/prune"
	"github.com/whosonfirst/go-whosonfirst-database/sql/tables"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

// To do: Add RunWithOptions...

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	if spatial_tables {
		rtree = true
		geojson = true
		properties = true
		spr = true
	}

	if spelunker_tables {
		// rtree = true
		spr = true
		spelunker = true
		geojson = true
		concordances = true
		ancestors = true
		search = true
	}

	db, err := database_sql.OpenWithURI(ctx, database_uri)

	if err != nil {
		return err
	}

	defer func() {

		err := db.Close()

		if err != nil {
			slog.Error("Failed to close database connection", "error", err)
		}
	}()

	to_prune := make([]database_sql.Table, 0)

	if geojson || all {

		gt, err := tables.NewGeoJSONTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create '%s' table because %s", tables.GEOJSON_TABLE_NAME, err)
		}

		to_prune = append(to_prune, gt)
	}

	if supersedes || all {

		t, err := tables.NewSupersedesTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create '%s' table because %s", tables.SUPERSEDES_TABLE_NAME, err)
		}

		to_prune = append(to_prune, t)
	}

	if rtree || all {

		gt, err := tables.NewRTreeTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create 'rtree' table because %s", err)
		}

		to_prune = append(to_prune, gt)
	}

	if properties || all {

		gt, err := tables.NewPropertiesTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create 'properties' table because %s", err)
		}

		to_prune = append(to_prune, gt)
	}

	if spr || all {

		st, err := tables.NewSPRTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create '%s' table because %s", tables.SPR_TABLE_NAME, err)
		}

		to_prune = append(to_prune, st)
	}

	if spelunker || all {

		st, err := tables.NewSpelunkerTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create '%s' table because %s", tables.SPELUNKER_TABLE_NAME, err)
		}

		to_prune = append(to_prune, st)
	}

	if names || all {

		nm, err := tables.NewNamesTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create '%s' table because %s", tables.NAMES_TABLE_NAME, err)
		}

		to_prune = append(to_prune, nm)
	}

	if ancestors || all {

		an, err := tables.NewAncestorsTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create '%s' table because %s", tables.ANCESTORS_TABLE_NAME, err)
		}

		to_prune = append(to_prune, an)
	}

	if concordances || all {

		cn, err := tables.NewConcordancesTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create '%s' table because %s", tables.CONCORDANCES_TABLE_NAME, err)
		}

		to_prune = append(to_prune, cn)
	}

	if geometries {

		gm, err := tables.NewGeometriesTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create '%s' table because %v", tables.CONCORDANCES_TABLE_NAME, err)
		}

		to_prune = append(to_prune, gm)
	}

	if search {

		st, err := tables.NewSearchTableWithDatabase(ctx, db)

		if err != nil {
			return fmt.Errorf("failed to create 'search' table because %v", err)
		}

		to_prune = append(to_prune, st)
	}

	if len(to_prune) == 0 {
		return fmt.Errorf("You forgot to specify which (any) tables to prune")
	}

	err = prune.PruneTables(ctx, db, to_prune...)
	
	if err != nil {
		return fmt.Errorf("Failed to prune tables, %w", err)
	}

	return nil
}
