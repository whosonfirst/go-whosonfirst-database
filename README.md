# go-whosonfirst-database

Go package implementing common properties and methods for working with Who's On First databases.

## Documentation

Documentation is incomplete at this time.

## sfomuseum/go-database

Under the hood this package makes extensive use of the [sfomuseum/go-database](https://github.com/sfomuseum/go-database/) package to provide common interfaces for (SQL) database tables and to provide a variety of helper methods.

## go-whosonfirst-database/sql

This package provides the core functionality for indexing one or more Who's On First sources (repos, etc.) in to a `database/sql` compatible database.

However, it does NOT load any specific `database/sql` drivers by default. This is assumed to happen in separate `go-whosonfirst-database-{DATABASE}` packages. This package simply provides common code used to index Who's On First documents in a `database/sql` compatible database.

### SQLite

* https://github.com/whosonfirst/go-whosonfirst-database-sqlite

## See also

* https://github.com/sfomuseum/go-database/
* https://github.com/whosonfirst/go-whosonfirst-database-sqlite