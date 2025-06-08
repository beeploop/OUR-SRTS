include .env .db.env

all: templates tailwind build

build:
	@echo "building..."
	@go build -o bin/our-srts cmd/main.go
	@echo "built binary in /bin"

run: 
	@go run cmd/main.go

test:
	@grc go test -v -failfast -cover ./...

clean:
	@rm -rf bin

server-watch:
	@air --build.cmd "go build -o tmp/main cmd/main.go" \
		--build.bin "tmp/main" --build.delay "100" \
		--build.exclude_dir [] --build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true

templates:
	@templ generate

templates-watch:
	@templ generate --watch --proxy="http://localhost:${PORT}" -v

tailwind:
	@pnpm tailwindcss -i ./web/assets/styles/tailwind.css -o ./web/assets/styles/style.css

tailwind-watch:
	@pnpm tailwindcss -i ./web/assets/styles/tailwind.css -o ./web/assets/styles/style.css --watch

migrate-up:
	@goose mysql ${DSN} -dir migrations up

migrate-reset:
	@goose mysql ${DSN} -dir migrations reset

watch:
	make -j3 templates-watch tailwind-watch server-watch

.PHONY: all build run test clean server-watch templates templates-watch tailwind tailwind-watch migrate-up migrate-reset
