package main

import(
	"log"
	"strconv"
	"net/http"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var mongodb = "mongodb://test:test123@ds155191.mlab.com:55191/gochat"

type Message struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{}

func main() {
	var PORT=8000

	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("Server started on port" + strconv.Itoa(PORT))
	httperr := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)

	if httperr != nil {
		log.Fatal("Error while setting up server", httperr)
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
	session, err := mgo.Dial(mongodb)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	chat := session.DB("gochat").C("chats")
	for {
		msg := <-broadcast
		err = chat.Insert(&Message{msg.Email, msg.Username, msg.Message})
		if err != nil {
			log.Fatal(err)
		}

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