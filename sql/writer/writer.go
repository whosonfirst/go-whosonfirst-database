package writer

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/url"
	"strconv"

	database_sql "github.com/sfomuseum/go-database/sql"
	wof_tables "github.com/whosonfirst/go-whosonfirst-database/sql/tables"
	wof_writer "github.com/whosonfirst/go-writer/v3"
)

func init() {
	ctx := context.Background()
	wof_writer.RegisterWriter(ctx, "wof-sql", NewSQLWriter)
}

// SQLWriter implements the `whosonfirst/go-writer/v3.Writer` interface for `database/sql` compatible Who's On First databases.
type SQLWriter struct {
	wof_writer.Writer
	db     *sql.DB
	tables []database_sql.Table
}

func NewSQLWriter(ctx context.Context, uri string) (wof_writer.Writer, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	db, err := database_sql.OpenWithURI(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create database, %w", err)
	}

	init_opts := new(wof_tables.InitTablesOptions)

	params := []string{
		"rtree",
		"geojson",
		"properties",
		"spr",
		"spelunker",
		"concordances",
		"ancestors",
		"search",
		"names",
		"supersedes",
		"geometries",
		"spelunker-tables",
		"spatial-tables",
		"all",
		"strict-alt-files",
	}

	for _, p := range params {

		if !q.Has(p) {
			continue
		}
		if q.Get(p) == "" {
			continue
		}

		switch p {
		case "index_alt":
			init_opts.IndexAlt = q[p]
		default:

			v, err := strconv.ParseBool(q.Get(p))

			if err != nil {
				return nil, fmt.Errorf("Failed to parse %s= parameter, %w", p, err)
			}

			switch p {
			case "rtree":
				init_opts.RTree = v
			case "geojson":
				init_opts.GeoJSON = v
			case "properties":
				init_opts.Properties = v
			case "spr":
				init_opts.SPR = v
			case "spelunker":
				init_opts.Spelunker = v
			case "concordances":
				init_opts.Concordances = v
			case "ancestors":
				init_opts.Ancestors = v
			case "search":
				init_opts.Search = v
			case "names":
				init_opts.Names = v
			case "supersedes":
				init_opts.Supersedes = v
			case "geometries":
				init_opts.Geometries = v
			case "spelunker-tables":
				init_opts.SpelunkerTables = v
			case "spatial-tables":
				init_opts.SpatialTables = v
			case "all":
				init_opts.All = v
			case "strict-alt-files":
				init_opts.StrictAltFiles = v
			default:
				slog.Warn("Invalid or unsupported parameter", "parameter", p)
			}

		}

	}

	to_index, err := wof_tables.InitTables(ctx, db, init_opts)

	if err != nil {
		return nil, err
	}

	wr := &SQLWriter{
		db:     db,
		tables: to_index,
	}

	return wr, nil
}

func (wr *SQLWriter) Write(ctx context.Context, path string, r io.ReadSeeker) (int64, error) {

	body, err := io.ReadAll(r)

	if err != nil {
		return 0, fmt.Errorf("Failed to read document, %w", err)
	}

	err = database_sql.IndexRecord(ctx, wr.db, body, wr.tables...)

	if err != nil {
		return 0, fmt.Errorf("Failed to index record, %w", err)
	}

	return 0, nil
}

func (wr *SQLWriter) WriterURI(ctx context.Context, uri string) string {
	return uri
}

func (wr *SQLWriter) Flush(ctx context.Context) error {
	return nil
}

func (wr *SQLWriter) Close(ctx context.Context) error {
	return nil
}

func (wr *SQLWriter) SetLogger(ctx context.Context, logger *log.Logger) error {
	slog.Debug("SQLWriter no longer supports SetLogger. Please use log/slog methods instead.")
	return nil
}
