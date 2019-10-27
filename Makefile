.PHONY: config build clean deploy

build:
	mkdir -p bin 
	env GOOS=linux go build -ldflags="-s -w" -o bin/litebot main.go
	zip -r -j -D bin/litebot.zip bin/litebot

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
