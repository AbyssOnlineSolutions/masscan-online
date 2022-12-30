package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var data MASSCAN

var mu sync.Mutex

// 接続されるクライアント
var clients = make(map[*websocket.Conn]bool)

// メッセージブロードキャストチャネル
var broadcast = make(chan Message)

// アップグレーダ
var upgrader = websocket.Upgrader{}

func main() {
	// ファイルサーバー
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// WebSocket
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// 送られてきたGETリクエストをWebSocketにアップグレード
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	// クライアントを登録
	clients[ws] = true

	mu.Lock()
	ws.WriteJSON(data)
	mu.Unlock()
	for {
		var message Message
		// 新しいメッセージをJSONとして読み込み、Message構造体にマッピング
		err := ws.ReadJSON(&message)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// 受け取ったメッセージをbroadcastチャネルに送る
		broadcast <- message
	}
}

func handleMessages() {
	for {
		// broadcastチャネルからメッセージを受け取る
		message := <-broadcast
		go StartScan(message.Cmd)
	}
}
