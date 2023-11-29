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
	@rm -rf ../registrar-release
	@mkdir ../registrar-release
	@mkdir ../registrar-release/logs 
	@mkdir ../registrar-release/tempDir
	@mkdir ../registrar-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized.exe ../registrar-release
	@mv ../registrar-release/.release.env ../registrar-release/.env
	@echo "zipping release directory..."
	@zip -r ../registrar-digitized.zip ../registrar-release
	@echo "zip complete. zip file saved in go/ directory"

linux-zip: clean build
	@echo "initializing release directory..."
	@rm -rf ../registrar-release 
	@mkdir ../registrar-release
	@mkdir ../registrar-release/logs 
	@mkdir ../registrar-release/tempDir
	@mkdir ../registrar-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized ../registrar-release
	@mv ../registrar-release/.release.env ../registrar-release/.env
	@echo "zipping release directory..."
	@zip -r ../registrar-digitized.zip ../registrar-release
	@echo "zip complete. zip file saved in go/ directory"

win-release: clean win
	@echo "initializing release directory..."
	@rm -rf ../registrar-release
	@mkdir ../registrar-release
	@mkdir ../registrar-release/logs 
	@mkdir ../registrar-release/tempDir
	@mkdir ../registrar-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized.exe ../registrar-release
	@mv ../registrar-release/.release.env ../registrar-release/.env
	@echo "compiled release version. release files saved in go/ directory"

linux-release: clean build 
	@echo "initializing release directory..."
	@rm -rf ../registrar-release 
	@mkdir ../registrar-release
	@mkdir ../registrar-release/logs 
	@mkdir ../registrar-release/tempDir
	@mkdir ../registrar-release/nas
	@echo "copying files to release directory..."
	@cp -r .release.env assets/ views/ webfonts/ registrar-digitized ../registrar-release
	@mv ../registrar-release/.release.env ../registrar-release/.env
	@echo "compiled release version. release files saved in go/ directory"

clean:
	@echo "cleaning up..."
	@rm -rf registrar-digitized
	@rm -rf registrar-digitized.exe
	@echo "clean complete"
