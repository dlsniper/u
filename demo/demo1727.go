package main

import "net"

func main() {
	clientList := make(map[int]struct {
		message chan string
		conn    *net.Conn
	})
	message := ""
	for _, c := range clientList {
		select {
		case c.message <- message: // message (not c.message) resolves to clientList.message field definition instead of message := "" declaration
		default:
			(*c.conn).Close() // c is not resolved, change the code to c.conn.Close() and resolving will work
		}
	}
}
