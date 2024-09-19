run: ./bin/cmd
	go run github.com/joho/godotenv/cmd/godotenv@v1.5.1 \
		./bin/cmd

dev/cmd:
	go run github.com/joho/godotenv/cmd/godotenv@v1.5.1 \
		go run cmd/main.go

build: ./bin/cmd

./bin/cmd:
	go build cmd/main.go -o bin/cmd
