build:
	@echo "building application..."
	@go build -o Registrar-SRTS
	@echo "build complete"

migrate-up:
	@migrate -database mysql://root:Password_1@/registrar -path db/migrations up

migrate-down:
	@migrate -database mysql://root:Password_1@/registrar -path db/migrations down

watch-css:
	@tailwindcss -i ./views/styles/tailwind.css -o ./views/styles/output.css --watch

build-css:
	@echo "building css..."
	@tailwindcss -i ./views/styles/tailwind.css -o ./views/styles/output.css
	@echo "build complete"

run: build 
	@./Registrar-SRTS

win:
	@echo "building windows application..."
	@env GOOS=windows GOARCH=amd64 go build -o Registrar-SRTS.exe
	@echo "build complete"

run-win: win
	@./Registrar-SRTS.exe

win-zip: clean win
	@echo "initializing release directory..."
	@rm -rf ./release/SRTS-windows-release
	@mkdir ./release/SRTS-windows-release
	@mkdir ./release/SRTS-windows-release/tempDir
	@mkdir ./release/SRTS-windows-release/nas
	@echo "copying files to release directory..."
	@cp -r .example.env assets/ views/ webfonts/ Registrar-SRTS.exe ./release/SRTS-windows-release
	@mv ./release/SRTS-windows-release/.example.env ./release/SRTS-windows-release/.env
	@echo "zipping release directory..."
	@cd release/ && zip -r Registrar-SRTS-windows.zip SRTS-windows-release
	@echo "zip complete. zip file saved in go/ directory"

linux-zip: clean build
	@echo "initializing release directory..."
	@rm -rf ./release/SRTS-linux-release
	@mkdir ./release/SRTS-linux-release
	@mkdir ./release/SRTS-linux-release/tempDir
	@mkdir ./release/SRTS-linux-release/nas
	@echo "copying files to release directory..."
	@cp -r .example.env assets/ views/ webfonts/ Registrar-SRTS ./release/SRTS-linux-release
	@mv ./release/SRTS-linux-release/.example.env ./release/SRTS-linux-release/.env
	@echo "zipping release directory..."
	@cd release/ && zip -r Registrar-SRTS-linux.zip SRTS-linux-release
	@echo "zip complete. zip file saved in go/ directory"

win-release: clean win
	@echo "initializing release directory..."
	@rm -rf ./release/SRTS-windows-release
	@mkdir ./release/SRTS-windows-release
	@mkdir ./release/SRTS-windows-release/tempDir
	@mkdir ./release/SRTS-windows-release/nas
	@echo "copying files to release directory..."
	@cp -r .example.env assets/ views/ webfonts/ Registrar-SRTS.exe ./release/SRTS-windows-release
	@mv ./release/SRTS-windows-release/.example.env ./release/SRTS-windows-release/.env
	@echo "compiled release version. release files saved in go/ directory"

linux-release: clean build 
	@echo "initializing release directory..."
	@rm -rf ./release/SRTS-linux-release
	@mkdir ./release/SRTS-linux-release
	@mkdir ./release/SRTS-linux-release/tempDir
	@mkdir ./release/SRTS-linux-release/nas
	@echo "copying files to release directory..."
	@cp -r .example.env assets/ views/ webfonts/ Registrar-SRTS ./release/SRTS-linux-release
	@mv ./release/SRTS-linux-release/.example.env ./release/SRTS-linux-release/.env
	@echo "compiled release version. release files saved in go/ directory"

clean:
	@echo "cleaning up..."
	@rm -rf Registrar-SRTS*
	@rm -rf release/*
	@rm -rf *.log
	@echo "clean complete"

test:
	@go test -v ./...
