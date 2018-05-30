package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"live"
	"os"
	"strings"
	"time"
	"websock"
)

func main() {
	var content io.Reader
	for {
		hf, ctt, err := live.HasFootball()
		if err != nil {
			fmt.Fprintf(os.Stderr, "HasFootball error: %v\n", err)
			os.Exit(1)
		}
		if hf {
			content = ctt
			break
		} else {
			fmt.Println("no football now")
		}

		time.Sleep(30 * time.Second)
	}
	initData, reactData, err := live.Parse(content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(initData, reactData, err)
	conn, resp, err := websock.CreateWebsocketConn(websock.WebsocketUrl(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create websocket failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read faield: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read faield: %v\n", err)
		os.Exit(1)
	}
	go func() {
		for {
			time.Sleep(25 * time.Second)
			conn.WriteMessage(websocket.TextMessage, []byte(`["\n"]`))
		}
	}()

	var msg string

	for {
		_, message, _ := conn.ReadMessage()
		switch message[0] {
		case 'o': // 连接打开
			fmt.Println(string(message))
			conn.WriteMessage(websocket.TextMessage, []byte(`["CONNECT\naccept-version:1.1,1.0\nheart-beat:10000,10000\n\n\u0000"]`))
		case 'h':
			fmt.Println(string(message))
		default:
			msg = string(message)
			fmt.Println(msg)
			if strings.Contains(msg, "CONNECTED") { // 连接建立
			}
		}
	}
}

func subscribe(conn websocket.Conn, stompQueue string, matchid int) {
	s := fmt.Sprintf(`["SUBSCRIBE\nx-queue-name:%s\nexclusive:true%s\nid:sub-0\ndestination:/exchange/live.animation.actions/%d\n\n\u0000"]`, stompQueue, matchid)
	conn.WriteMessage(websocket.TextMessage, []byte(s))
}
