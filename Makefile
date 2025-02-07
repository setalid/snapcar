.PHONY: test
test:
	cd api && go test ./...

.PHONY: dev-web
web:
	cd web && npm run dev

.PHONY: dev-api
api:
	cd api && go run .
