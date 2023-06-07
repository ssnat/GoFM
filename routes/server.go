package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pxgo/GoFM/modules"
	"github.com/pxgo/GoFM/tools"
	"net/http"
)

func GetServerInfo(ctx echo.Context) error {
	musicInfo := modules.MusicReader.GetMusicInfo()
	err := ctx.JSON(http.StatusOK, tools.Response.GetResponseBody(struct {
		Name    string              `json:"name"`
		Version string              `json:"version"`
		Time    int64               `json:"time"`
		FMInfo  *modules.IMusicInfo `json:"FMInfo"`
	}{
		Name:    modules.Config.Name,
		Version: modules.Config.Version,
		Time:    modules.Config.Time,
		FMInfo:  musicInfo,
	}))
	if err != nil {
		modules.Logger.Error(err)
		return err
	}
	return nil
}
