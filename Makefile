build:
	@go build

css:
	@tailwindcss -i ./views/styles/tailwind.css -o ./views/styles/output.css --watch

run: build 
	@./registrar


