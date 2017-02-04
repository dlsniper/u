package main

type (
	client struct {
		message chan string
	}
	clientList struct {
		m (map[string]*client)
	}
)

func main() {
	clientList := clientList{m: make(map[string]*client)}
	message := ""
	for _, c := range clientList.m {
		c.message <- message
	}
}
