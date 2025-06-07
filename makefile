build:
	@go build -o bin/our-srts cmd/main.go

run: 
	@go run cmd/main.go

test:
	@grc go test -v -failfast -cover ./...

clean:
	@rm -rf bin

watch:
	@air --build.cmd "go build -o tmp/main cmd/main.go" \
		--build.bin "tmp/main" --build.delay "100" \
		--build.exclude_dir [] --build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true
