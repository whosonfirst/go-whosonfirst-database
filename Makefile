GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

TAGS=sqlite3,fts5

cli:
	@make cli-sql
	@make cli-opensearch

cli-sql:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-sql-create cmd/wof-sql-create/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-sql-index cmd/wof-sql-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-sql-prune cmd/wof-sql-prune/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -tags $(TAGS) -o bin/wof-sql-iterwriter cmd/wof-sql-iterwriter/main.go

cli-opensearch:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-index cmd/wof-opensearch-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-query cmd/wof-opensearch-query/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-create-index cmd/wof-opensearch-create-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-delete-index cmd/wof-opensearch-delete-index/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-put-mapping cmd/wof-opensearch-put-mapping/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-get-mapping cmd/wof-opensearch-get-mapping/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-put-settings cmd/wof-opensearch-put-settings/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-list-indices cmd/wof-opensearch-list-indices/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-opensearch-indices-stats cmd/wof-opensearch-indices-stats/main.go
