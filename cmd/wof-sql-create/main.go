package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst-database/app/sql/tables/create"
)

func main() {

	ctx := context.Background()
	err := create.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to create tables, %v", err)
	}
}
