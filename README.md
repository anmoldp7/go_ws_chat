# Go Websocket Chat

Basic implementation of chat using websockets from [Go Blueprints](https://www.oreilly.com/library/view/go-programming-blueprints/9781786468949/).

## Requirements

Install [Golang](https://golang.org/doc/install). Install websocket dependency by executing the following:

```bash
go get "github.com/gorilla/websocket"
```

## Usage

Execute:

```bash
go build -o ws_chat
./ws_chat
```

While the program is running, access [localhost:8080](http://localhost:8080) in 2 or more tabs, fill the text area and submit.
