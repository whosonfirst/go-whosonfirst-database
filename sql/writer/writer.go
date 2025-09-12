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

	sfom_sql "github.com/sfomuseum/go-database/sql"
	wof_tables "github.com/whosonfirst/go-whosonfirst-database/sql/tables"
	wof_writer "github.com/whosonfirst/go-writer/v3"
)

func init() {
	ctx := context.Background()
	wof_writer.RegisterWriter(ctx, "sql", NewMySQLWriter)
}

type MySQLWriter struct {
	wof_writer.Writer
	db     *sql.DB
	tables []sfom_sql.Table
}

func NewMySQLWriter(ctx context.Context, uri string) (wof_writer.Writer, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	db, err := sfom_sql.OpenWithURI(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create database, %w", err)
	}

	index_geojson := true
	index_whosonfirst := true

	if q.Get("geojson") != "" {

		index, err := strconv.ParseBool(q.Get("geojson"))

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?geojson= parameter, %w", err)
		}

		index_geojson = index
	}

	if q.Get("whosonfirst") != "" {

		index, err := strconv.ParseBool(q.Get("whosonfirst"))

		if err != nil {
			return nil, fmt.Errorf("Failed to parse ?whosonfirst= parameter, %w", err)
		}

		index_whosonfirst = index
	}

	to_index := make([]sfom_sql.Table, 0)

	if index_geojson {

		t, err := wof_tables.NewGeoJSONTableWithDatabase(ctx, db)

		if err != nil {
			return nil, fmt.Errorf("Failed to create GeoJSON table, %w", err)
		}

		to_index = append(to_index, t)
	}

	if index_whosonfirst {

		t, err := wof_tables.NewWhosonfirstTableWithDatabase(ctx, db)

		if err != nil {
			return nil, fmt.Errorf("Failed to create Whosonfirst table, %w", err)
		}

		to_index = append(to_index, t)
	}

	wr := &MySQLWriter{
		db:     db,
		tables: to_index,
	}

	return wr, nil
}

func (wr *MySQLWriter) Write(ctx context.Context, path string, r io.ReadSeeker) (int64, error) {

	body, err := io.ReadAll(r)

	if err != nil {
		return 0, fmt.Errorf("Failed to read document, %w", err)
	}

	for _, t := range wr.tables {

		err = t.IndexRecord(ctx, wr.db, body)

		if err != nil {
			return 0, fmt.Errorf("Failed to index %s table for %s, %w", t.Name(), path, err)
		}
	}

	return 0, nil
}

func (wr *MySQLWriter) WriterURI(ctx context.Context, uri string) string {
	return uri
}

func (wr *MySQLWriter) Flush(ctx context.Context) error {
	return nil
}

func (wr *MySQLWriter) Close(ctx context.Context) error {
	return nil
}

func (wr *MySQLWriter) SetLogger(ctx context.Context, logger *log.Logger) error {
	slog.Debug("MySQLWriter no longer supports SetLogger. Please use log/slog methods instead.")
	return nil
}
