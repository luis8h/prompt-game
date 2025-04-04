.PHONY: tailwind-watch
tailwind-watch:
	npx tailwindcss -i ./static/css/tailwind/input.css -o ./static/css/tailwind/output.css --watch

.PHONY: tailwind-build
tailwind-build:
	npx tailwindcss -i ./static/css/tailwind/input.css -o ./static/css/tailwind/output.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...
