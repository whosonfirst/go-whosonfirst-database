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

	init_opts := &tables.InitTablesOptions{
		RTree:           rtree,
		GeoJSON:         geojson,
		Properties:      properties,
		SPR:             spr,
		Spelunker:       spelunker,
		Concordances:    concordances,
		Ancestors:       ancestors,
		Search:          search,
		Names:           names,
		Supersedes:      supersedes,
		SpatialTables:   spatial_tables,
		SpelunkerTables: spelunker_tables,
		All:             all,
	}

	to_prune, err := tables.InitTables(ctx, db, init_opts)

	if err != nil {
		return err
	}

	err = prune.PruneTables(ctx, db, to_prune...)

	if err != nil {
		return fmt.Errorf("Failed to prune tables, %w", err)
	}

	return nil
}
