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

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}

}

func handleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}