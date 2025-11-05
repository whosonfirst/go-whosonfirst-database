package create

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var iterator_uri string

var database_uri string

var all bool
var ancestors bool
var concordances bool
var geojson bool
var spelunker bool
var geometries bool
var names bool
var rtree bool
var properties bool
var search bool
var spr bool
var supersedes bool

var spatial_tables bool
var spelunker_tables bool

var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("index")

	fs.StringVar(&database_uri, "database-uri", "", "...")

	fs.BoolVar(&all, "all", false, "Index all tables (except the 'search' and 'geometries' tables which you need to specify explicitly)")
	fs.BoolVar(&ancestors, "ancestors", false, "Index the 'ancestors' tables")
	fs.BoolVar(&concordances, "concordances", false, "Index the 'concordances' tables")
	fs.BoolVar(&geojson, "geojson", false, "Index the 'geojson' table")
	fs.BoolVar(&spelunker, "spelunker", false, "Index the 'spelunker' table")
	fs.BoolVar(&geometries, "geometries", false, "Index the 'geometries' table (requires that libspatialite already be installed)")
	fs.BoolVar(&names, "names", false, "Index the 'names' table")
	fs.BoolVar(&rtree, "rtree", false, "Index the 'rtree' table")
	fs.BoolVar(&properties, "properties", false, "Index the 'properties' table")
	fs.BoolVar(&search, "search", false, "Index the 'search' table (using SQLite FTS5 full-text indexer)")
	fs.BoolVar(&spr, "spr", false, "Index the 'spr' table")
	fs.BoolVar(&supersedes, "supersedes", false, "Index the 'supersedes' table")

	fs.BoolVar(&spatial_tables, "spatial-tables", false, "If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spatial-sqlite package.")
	fs.BoolVar(&spelunker_tables, "spelunker-tables", false, "If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spelunker packages")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging")
	return fs
}
