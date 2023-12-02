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
	@rm -rf ./release/windows-release
	@mkdir ./release/windows-release
	@mkdir ./release/windows-release/tempDir
	@mkdir ./release/windows-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized.exe ./release/windows-release
	@mv ./release/windows-release/.release.env ./release/windows-release/.env
	@echo "zipping release directory..."
	@cd release/ && zip -r registrar-digitized-windows.zip windows-release
	@echo "zip complete. zip file saved in go/ directory"

linux-zip: clean build
	@echo "initializing release directory..."
	@rm -rf ./release/linux-release
	@mkdir ./release/linux-release
	@mkdir ./release/linux-release/tempDir
	@mkdir ./release/linux-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized ./release/linux-release
	@mv ./release/linux-release/.release.env ./release/linux-release/.env
	@echo "zipping release directory..."
	@cd release/ && zip -r registrar-digitized-linux.zip linux-release
	@echo "zip complete. zip file saved in go/ directory"

win-release: clean win
	@echo "initializing release directory..."
	@rm -rf ./release/windows-release
	@mkdir ./release/windows-release
	@mkdir ./release/windows-release/tempDir
	@mkdir ./release/windows-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized.exe ./release/windows-release
	@mv ./release/windows-release/.release.env ./release/windows-release/.env
	@echo "compiled release version. release files saved in go/ directory"

linux-release: clean build 
	@echo "initializing release directory..."
	@rm -rf ./release/linux-release
	@mkdir ./release/linux-release
	@mkdir ./release/linux-release/tempDir
	@mkdir ./release/linux-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized ./release/linux-release
	@mv ./release/linux-release/.release.env ./release/linux-release/.env
	@echo "compiled release version. release files saved in go/ directory"

clean:
	@echo "cleaning up..."
	@rm -rf registrar-digitized*
	@rm -rf release/*
	@echo "clean complete"

test:
	@go test -v ./...
