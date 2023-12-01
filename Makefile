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

win-zip: clean win
	@echo "initializing release directory..."
	@rm -rf ./release/registrar-windows-release
	@mkdir ./release/registrar-windows-release
	@mkdir ./release/registrar-windows-release/tempDir
	@mkdir ./release/registrar-windows-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized.exe ./release/registrar-windows-release
	@mv ./release/registrar-windows-release/.release.env ./release/registrar-windows-release/.env
	@echo "zipping release directory..."
	@cd release/ && zip -r registrar-digitized-windows.zip registrar-windows-release
	@echo "zip complete. zip file saved in go/ directory"

linux-zip: clean build
	@echo "initializing release directory..."
	@rm -rf ./release/registrar-linux-release
	@mkdir ./release/registrar-linux-release
	@mkdir ./release/registrar-linux-release/tempDir
	@mkdir ./release/registrar-linux-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized ./release/registrar-linux-release
	@mv ./release/registrar-linux-release/.release.env ./release/registrar-linux-release/.env
	@echo "zipping release directory..."
	@cd release/ && zip -r registrar-digitized-linux.zip registrar-linux-release
	@echo "zip complete. zip file saved in go/ directory"

win-release: clean win
	@echo "initializing release directory..."
	@rm -rf ./release/registrar-windows-release
	@mkdir ./release/registrar-windows-release
	@mkdir ./release/registrar-windows-release/tempDir
	@mkdir ./release/registrar-windows-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized.exe ./release/registrar-windows-release
	@mv ./release/registrar-windows-release/.release.env ./release/registrar-windows-release/.env
	@echo "compiled release version. release files saved in go/ directory"

linux-release: clean build 
	@echo "initializing release directory..."
	@rm -rf ./release/registrar-linux-release
	@mkdir ./release/registrar-linux-release
	@mkdir ./release/registrar-linux-release/tempDir
	@mkdir ./release/registrar-linux-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized ./release/registrar-linux-release
	@mv ./release/registrar-linux-release/.release.env ./release/registrar-linux-release/.env
	@echo "compiled release version. release files saved in go/ directory"

clean:
	@echo "cleaning up..."
	@rm -rf registrar-digitized*
	@echo "clean complete"

test:
	@go test -v ./...
