BINARY_NAME=slack-file-bot
.DEFAULT_GOAL=run
build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 main.go
	mkdir -p env
	echo "SLACK_BOT_TOKEN=xoxb-5219707682822-5988051349155-FxBuqSxkOtWEkqdljGWOCCRQ" > env/file.env
	echo "CHANNEL_ID=C0570U39R7B" >> env/file.env

run: build
	./bin/${BINARY_NAME}-linux-amd64
 
clean:
	go clean
	rm -rf bin/${BINARY_NAME}-linux-amd64
	rm -rf env/file.env