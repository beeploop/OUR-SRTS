build:
	@echo "building application..."
	@go build
	@echo "build complete"

watch-css:
	@tailwindcss -i ./views/styles/tailwind.css -o ./views/styles/output.css --watch

build-css:
	@echo "building css..."
	@tailwindcss -i ./views/styles/tailwind.css -o ./views/styles/output.css
	@echo "build complete"

run: build 
	@./registrar-digitized

win:
	@echo "building windows application..."
	@env GOOS=windows GOARCH=amd64 go build 
	@echo "build complete"

run-win: win
	@./registrar-digitized.exe

clean:
	@echo "cleaning up..."
	@rm -rf registrar-digitized
	@rm -rf registrar-digitized.exe
	@rm logs/*
	@rm documents/*
	@echo "clean complete"
