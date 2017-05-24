package main

import(
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)


type Message struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{}

func main() {
	var PORT=8000

	fs := http.FileServer(http.Dir('../public'))
	http.Handle('/', fs)

	http.HandleFunc('/ws', handleConnections)

	go handleMessages()

	log.Println('Server started on port' + PORT)
	err := http.ListenAndServe(':'+PORT, nil)

	if err != nil {
		log.Fatal('Error while setting up server', err)
	}
} 