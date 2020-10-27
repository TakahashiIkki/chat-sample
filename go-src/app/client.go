package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行っている1人のユーザーを表します
type client struct {
	// socket はこのクライアントの為のWebSocketです
	socket *websocket.Conn
	// send はメッセージが送られるチャネルです
	send chan []byte
	// room はこのクライアントが参加しているチャットルームです
	room *room
}

func (c *client) read() {
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
	c.socket.Close()
}