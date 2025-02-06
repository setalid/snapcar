.PHONY: test
test:
	cd api && go test ./...

.PHONY: dev
dev:
	cd api && go run .
