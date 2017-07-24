# GoChat

GoChat is a simple chat application, which mimics a basic chatroom where anyone can join the conversation with an email id. The application utilises the Websockets protocol to achieve real-time messaging. The backend is written in Go and allows you to log the chat to a MongoDB Document. The frontend is a simple single page app built with VueJS which allows the user to join and send messages(also supports emoji). 

## Usage
* Fork and clone the repository to your computer.
* In the `main.go` file in the `src` directory change the MongoDB URI to point to your own database.
* While in the `src` directory, execute `go run main.go` to start the server.

## Libraries
* [mgo.v2](https://gopkg.in/mgo.v2)
* [gorilla/websockets](https://github.com/gorilla/websocket)
* [VueJS](https://vuejs.org)
* [jQuery](https://jquery.com/)
* [MaterializeCSS](https://materializecss.com/)
* [EmojiOne](https://www.emojione.com/)

## License
This project is licensed under the [MIT License](https://github.com/GoChat/Cosmo/blob/master/LICENSE.md)
