package routes

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/pxgo/GoFM/modules"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWS(ctx echo.Context) error {
	conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		modules.Logger.Error(err)
		return err
	}

	modules.WS.SaveClient(conn)
	modules.WS.SendMusicInfoToTargetClient(conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				modules.WS.RemoveClient(conn)
				modules.Logger.Error(err)
				break
			}
			modules.Logger.Error(err)
			return err
		}
	}
	return nil
}
