package websocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id     uuid.UUID
	socket *websocket.Conn
	send   chan EventResponse
	//Used for filtering to which client response will be sent
	Apparatus     Apparatus
	CompetitionId string
}

// Reads incoming messages and sends it to server
func (client *Client) read(server *Server) {
	defer func() {
		server.unregister <- client
		client.socket.Close()
	}()

	for {
		var message EventMessage
		err := client.socket.ReadJSON(&message)
		if err != nil {
			break
		}
		server.broadcast <- message
	}
}

func (client *Client) write() {
	defer func() {
		client.socket.Close()
	}()

	for {
		select {
		//ok in this context checks if channel is still open and there is message available
		case message, ok := <-client.send:
			if !ok {
				client.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.socket.WriteJSON(message)
		}
	}
}
