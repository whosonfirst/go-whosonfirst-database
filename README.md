# go-whosonfirst-database

Go package implementing common properties and methods for working with Who's On First databases.

## Documentation

Documentation is incomplete at this time.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -tags sqlite3 -o bin/wof-sql-create cmd/wof-sql-create/main.go
go build -mod vendor -ldflags="-s -w" -tags sqlite3 -o bin/wof-sql-index cmd/wof-sql-index/main.go
go build -mod vendor -ldflags="-s -w" -tags sqlite3 -o bin/wof-sql-prune cmd/wof-sql-prune/main.go
```

Database support is enabled through tags. The following tags are supported:

| Tag | Package | Notes |
| --- | --- | --- |
| sqlite3 | [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) | |
| mysql | [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) | |


### wof-sql-create

Create, but do not index, one or more tables in a `database/sql` compatiable database.

```
$> ./bin/wof-sql-create -h
  -all
    	Index all tables (except the 'search' and 'geometries' tables which you need to specify explicitly)
  -ancestors
    	Index the 'ancestors' tables
  -concordances
    	Index the 'concordances' tables
  -database-uri string
    	...
  -geojson
    	Index the 'geojson' table
  -geometries
    	Index the 'geometries' table (requires that libspatialite already be installed)
  -names
    	Index the 'names' table
  -properties
    	Index the 'properties' table
  -rtree
    	Index the 'rtree' table
  -search
    	Index the 'search' table (using SQLite FTS4 full-text indexer)
  -spatial-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spatial-sqlite package.
  -spelunker
    	Index the 'spelunker' table
  -spelunker-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spelunker packages
  -spr
    	Index the 'spr' table
  -supersedes
    	Index the 'supersedes' table
  -verbose
    	Enable verbose (debug) logging
```

### wof-sql-index

Index one or more tables in a `database/sql` compatiable database.

```
$> ./bin/wof-sql-index -h
  -all
    	Index all tables (except the 'search' and 'geometries' tables which you need to specify explicitly)
  -ancestors
    	Index the 'ancestors' tables
  -concordances
    	Index the 'concordances' tables
  -database-uri string
    	A URI in the form of 'sql://{DATABASE_SQL_ENGINE}?dsn={DATABASE_SQL_DSN}'. For example: sql://sqlite3?dsn=test.db
  -geojson
    	Index the 'geojson' table
  -geometries
    	Index the 'geometries' table (requires that libspatialite already be installed)
  -index-alt value
    	Zero or more table names where alt geometry files should be indexed.
  -index-relations
    	Index the records related to a feature, specifically wof:belongsto, wof:depicts and wof:involves. Alt files for relations are not indexed at this time.
  -index-relations-reader-uri string
    	A valid go-reader.Reader URI from which to read data for a relations candidate.
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v3.Iterator URI. Supported iterator URI schemes are: cwd://,directory://,featurecollection://,file://,filelist://,geojsonl://,null://,repo:// (default "repo://")
  -names
    	Index the 'names' table
  -optimize
    	Attempt to optimize the database before closing connection (default true)
  -processes int
    	The number of concurrent processes to index data with (default 20)
  -properties
    	Index the 'properties' table
  -rtree
    	Index the 'rtree' table
  -search
    	Index the 'search' table (using SQLite FTS4 full-text indexer)
  -spatial-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spatial-sqlite package.
  -spelunker
    	Index the 'spelunker' table
  -spelunker-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spelunker packages
  -spr
    	Index the 'spr' table
  -strict-alt-files
    	Be strict when indexing alt geometries (default true)
  -supersedes
    	Index the 'supersedes' table
  -timings
    	Display timings during and after indexing
  -verbose
    	Enable verbose (debug) logging
```

#### For example:

```
$> ./bin/wof-sql-index \
	-spatial-tables \
	-timings \
	-database-uri 'sql://sqlite3?dsn=test2.db' \
	/usr/local/data/sfomuseum-data-whosonfirst
	
2025/09/12 12:46:31 INFO Iterator stats elapsed=27.444911792s seen=1604 allocated="1.6 MB" "total allocated"="10 GB" sys="284 MB" numgc=2650
```

And then to use that database with, for example, the [whosonfirst/go-whosonfirst-spatial-sqlite`](#) package:

```
$> cd /usr/local/whosonfirst/go-whosonfirst-spatial-sqlite`](#) package:
$> ./bin/pip \
	-spatial-database-uri 'sqlite://sqlite3?dsn=/usr/local/whosonfirst/go-whosonfirst-database/test2.db' \
	-latitude 37.616951 \
	-longitude -122.383747 \
| jq -r '.places[]["wof:name"]'

Earth
North America
United States
California
San Mateo
San Francisco International Airport
94128
```

### wof-sql-prune

Remove all the records from one or more tables in a `database/sql` compatible Who's On First database.

```
$> ./bin/wof-sql-prune -h
  -all
    	Index all tables (except the 'search' and 'geometries' tables which you need to specify explicitly)
  -ancestors
    	Index the 'ancestors' tables
  -concordances
    	Index the 'concordances' tables
  -database-uri string
    	...
  -geojson
    	Index the 'geojson' table
  -geometries
    	Index the 'geometries' table (requires that libspatialite already be installed)
  -names
    	Index the 'names' table
  -properties
    	Index the 'properties' table
  -rtree
    	Index the 'rtree' table
  -search
    	Index the 'search' table (using SQLite FTS4 full-text indexer)
  -spatial-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spatial-sqlite package.
  -spelunker
    	Index the 'spelunker' table
  -spelunker-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spelunker packages
  -spr
    	Index the 'spr' table
  -supersedes
    	Index the 'supersedes' table
  -verbose
    	Enable verbose (debug) logging
```

## See also

* https://github.com/sfomuseum/go-database/