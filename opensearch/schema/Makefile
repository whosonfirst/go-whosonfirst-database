# This is for debugging. Do not change this at your own risk.
# (That means you should change this.)
OSPSWD=dkjfhsjdkfkjdjhksfhskd98475kjHkzjxckj

FIELD_LIMIT=7000

spelunker-local:
	@make spelunker-local-index
	@make spelunker-local-fieldlimit

spelunker-local-index:
	cat 2.x/mappings.spelunker.json | \
		curl -k \
		-H 'Content-Type: application/json' \
		-X PUT \
		https://admin:$(OSPSWD)@localhost:9200/spelunker \
		-d @-

spelunker-local-fieldlimit:
	curl -k \
		-H 'Content-type:application/json' \
		-XPUT https://admin:$(OSPSWD)@localhost:9200/spelunker/_settings \
		-d '{"index.mapping.total_fields.limit": $(FIELD_LIMIT)}'
