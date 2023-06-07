package routes

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	e.GET("/api/info", GetServerInfo)
	e.GET("/api/ws", HandleWS)
	e.GET("/api/fm", GetFMStream)
	e.GET("/api/fm/info", GetMusicInfo)
	e.GET("/api/fm/info/cover", GetMusicCover)
	e.GET("/favicon.ico", GetMusicCover)
}
