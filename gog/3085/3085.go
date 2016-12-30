package main

import "github.com/gorilla/websocket"

type demo struct{
	socket *websocket.Conn
}

func (d *demo) Stuff() {
	defer d.socket.Close()
}

func main() {
	d := &demo{}
	defer d.socket.Close()
}
