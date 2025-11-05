package create

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	_ "github.com/whosonfirst/go-whosonfirst-database/sql"

	database_sql "github.com/sfomuseum/go-database/sql"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-database/sql/tables"
)

const index_alt_all string = "*"

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

	logger := slog.Default()

	db, err := database_sql.OpenWithURI(ctx, database_uri)

	if err != nil {
		return fmt.Errorf("Failed to create database connection, %w", err)
	}

	defer func() {

		err := db.Close()

		if err != nil {
			logger.Error("Failed to close database connection", "error", err)
		}
	}()

	init_opts := &tables.InitTablesOptions{
		RTree:           rtree,
		GeoJSON:         geojson,
		Geometries:      geometries,
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

	to_create, err := tables.InitTables(ctx, db, init_opts)

	if err != nil {
		return fmt.Errorf("Failed to initialize tables, %w", err)
	}

	db_opts := database_sql.DefaultConfigureDatabaseOptions()
	db_opts.CreateTablesIfNecessary = true
	db_opts.Tables = to_create

	return database_sql.ConfigureDatabase(ctx, db, db_opts)

}
