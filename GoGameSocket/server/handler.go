package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	clientConnections   = make(map[uuid.UUID]*ClientConnection)
	clientConnectionsMu sync.Mutex
	queueMessages       []CommandMessage
	waitingRoom         []uuid.UUID
	gameRooms           []GameRoom
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	connectionId := uuid.New()
	// when the connection is closed.
	conn.SetCloseHandler(func(code int, text string) error {
		clientConnectionsMu.Lock()
		delete(clientConnections, connectionId)
		log.Printf("WebSocket closed: code=%d, text=%s", code, text)
		clientConnectionsMu.Unlock()
		return nil
	})
	// When our method is destroyd.
	defer func() {
		conn.Close()
	}()
	// Add Client To the memory.
	clientConnectionsMu.Lock()
	clientConnections[connectionId] = &ClientConnection{
		ConnectionId: connectionId,
		Conn:         conn,
	}
	clientConnectionsMu.Unlock()
	otherClients := getOtherClients(connectionId)
	sendMessage(otherClients, CommandMessage{CommandType: "textMessage", Data: "new client is connected with connectionId:" + (connectionId.String())})

	for {
		// message from client.
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var commandMessage CommandMessage
		parseErr := json.Unmarshal(message, &commandMessage)
		if parseErr != nil {
			log.Println(parseErr)
			return
		}
		commandMessage.ConnectionId = connectionId
		queueMessages = append(queueMessages, commandMessage)
		fmt.Println(string(message))
		/// put my message inside queue.

		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
	// end of the code defer.
}

// this function is will return other users.
func getOtherClients(clientId uuid.UUID) []*ClientConnection {
	var clients = []*ClientConnection{}
	for connctionId := range clientConnections {
		if connctionId == clientId {
			continue
		}
		client := clientConnections[connctionId]
		clients = append(clients, client)
	}
	return clients
}

func sendMessage(connections []*ClientConnection, commandMessage CommandMessage) {
	for _, clientConnection := range connections {
		clientConnection.Conn.WriteJSON(commandMessage)
	}
}

func QueueMessageTask() {
	for {
		if len(queueMessages) > 0 {
			pickedMessage := queueMessages[0]
			waitingRoom = append(waitingRoom, pickedMessage.ConnectionId)
			queueMessages = queueMessages[1:]
		}
		time.Sleep(2 * time.Second) // simulate periodic work
	}
}

func WaitingRoomTask() {
	for {
		if len(waitingRoom) != 0 && len(waitingRoom)%2 == 0 {
			newRoom := GameRoom{
				GameId:            uuid.New(),
				ClientConnections: []uuid.UUID{},
			}
			gameRooms = append(gameRooms, newRoom)
		}
		time.Sleep(5 * time.Second) // simulate periodic work
	}
}

// background task to handle queue.
// go backgroundTask()
// time.Sleep(5 * time.Second) // simulate periodic work
