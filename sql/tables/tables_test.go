package tables

import (
	"database/sql"
	"testing"
)

func TestLoadSchema(t *testing.T) {

	t.Skip()

	engines := []string{
		// "sqlite",
	}

	table_names := []string{
		ANCESTORS_TABLE_NAME,
		CONCORDANCES_TABLE_NAME,
		GEOJSON_TABLE_NAME,
		GEOMETRIES_TABLE_NAME,
		NAMES_TABLE_NAME,
		PROPERTIES_TABLE_NAME,
		RTREE_TABLE_NAME,
		SEARCH_TABLE_NAME,
		SPR_TABLE_NAME,
		SUPERSEDES_TABLE_NAME,
	}

	for _, e := range engines {

		var db *sql.DB

		for _, n := range table_names {

			_, err := LoadSchema(db, n)

			if err != nil {
				t.Fatalf("Failed to load %s table for %s database engine, %v", n, e, err)
			}
		}
	}
}
