package main

type room struct {
	// forward は他のクライアントに転送するためのメッセージを保持するチャネルです。
	forward chan []byte
	// join はチャットルームに参加しようとしているクライアントの為のチャネルです。
	join chan *client
	// leave はチャットルームから退室しようとしているクライアントの為のチャネルです。
	leave chan *client
	// client には在室しているすべてのクライアントが保持されます。
	clients map[*client]bool
}
