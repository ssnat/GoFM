package routes

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	e.GET("/info", GetServerInfo)
	e.GET("/ws", HandleWS)
	e.GET("/fm", GetFMStream)
	e.GET("/fm/info", GetMusicInfo)
	e.GET("/fm/info/cover", GetMusicCover)
	e.GET("/favicon.ico", GetMusicCover)
}
