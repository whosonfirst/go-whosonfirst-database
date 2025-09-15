package main

// This tool does NOT load any `database/sql` drivers. It is provided as an example of code that might do so (load `database/sql` drivers) and use the `app/sql/index` package to index database records.

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst-database/app/sql/tables/prune"
)

func main() {

	ctx := context.Background()
	err := prune.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to prune, %v", err)
	}
}
