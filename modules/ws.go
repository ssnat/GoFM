package modules

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

type IWS struct {
	Clients []*websocket.Conn
}

var WS = IWS{}

func (ws *IWS) SaveClient(conn *websocket.Conn) {
	ws.Clients = append(ws.Clients, conn)
}

func (ws *IWS) RemoveClient(conn *websocket.Conn) {
	var clients []*websocket.Conn
	for _, client := range ws.Clients {
		if client != conn {
			clients = append(clients, client)
		}
	}
	ws.Clients = clients
}

func (ws *IWS) SendMusicInfoToTargetClient(conn *websocket.Conn) {
	musicInfoJson, err := json.Marshal(MusicReader.GetMusicInfo())
	if err != nil {
		Logger.Error(err)
	} else {
		err := conn.WriteMessage(websocket.TextMessage, musicInfoJson)
		if err != nil {
			Logger.Error(err)
		}
	}
}

func (ws *IWS) SendMusicInfoToAllClient() {
	for _, client := range ws.Clients {
		ws.SendMusicInfoToTargetClient(client)
	}
}
