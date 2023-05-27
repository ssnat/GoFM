package main

import (
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pxgo/GoFM/middlewares"
	"github.com/pxgo/GoFM/modules"
	"github.com/pxgo/GoFM/routes"
	"io/fs"
	"net/http"
	"os"
)

//go:embed web/dist/*
var publicFiles embed.FS

func main() {

	modules.InitReader()

	e := echo.New()

	e.HideBanner = true
	e.HTTPErrorHandler = middlewares.CustomHTTPErrorHandler
	e.Use(middlewares.LoggerIn)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	var fileSystem http.FileSystem
	if modules.Config.Debug {
		fileSystem = http.FS(os.DirFS("web/dist"))
	} else {
		fsys, err := fs.Sub(publicFiles, "web/dist")
		if err != nil {
			modules.Logger.Error(err)
			panic(err)
		}

		fileSystem = http.FS(fsys)
	}

	assetHandler := http.FileServer(fileSystem)

	routes.InitRoutes(e)

	e.GET("/*", echo.WrapHandler(assetHandler))

	err := e.Start(fmt.Sprintf("%s:%d", modules.Config.Host, modules.Config.Port))
	if err != nil {
		modules.Logger.Error(err)
	}
}
