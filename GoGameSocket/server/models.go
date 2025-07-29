package server

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ClientConnection struct {
	ConnectionId uuid.UUID
	Conn         *websocket.Conn
	/// Room Id
}

/// Type room with list of connected clients and room id.

// { "commandType" : "textMessage" , "data": "hello im new client" }
// commandTypes joinWaitingRoom
type CommandMessage struct {
	CommandType  string `json:"commandType"`
	Data         string `json:"data"`
	ConnectionId uuid.UUID
}

type GameRoom struct {
	GameId            uuid.UUID
	ClientConnections []uuid.UUID
}
