package routes

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pxgo/GoFM/modules"
	"net/http"
	"strings"
	"time"
)

func GetFMStream(ctx echo.Context) error {

	ip := GetRealIP(ctx.Request())

	modules.Logger.Info(fmt.Sprintf("Client %s connected", ip))

	res := ctx.Response()

	store := modules.MusicReader.GetBufferStoreData()
	if store == nil {
		err := errors.New("oops, it seems like the FM hasn't started up")
		modules.Logger.Error(err)
		return err
	}

	res.Header().Set("Connection", "Keep-Alive")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.Header().Set("Transfer-Encoding", "chunked")
	res.Header().Set("Content-Type", "audio/mpeg")

	init := false
	order := 0

	for {
		var targetBuffer []byte

		store := modules.MusicReader.GetBufferStoreData()

		if store.Order == order {
			time.Sleep(time.Millisecond * 100)
			continue
		}

		order = store.Order

		if !init {
			init = true
			targetBuffer = store.InitialBuffer[:]
		} else {
			targetBuffer = store.UnitBuffer[:]
		}

		var timeout = store.Timeout

		_, err := res.Write(targetBuffer)
		if err != nil {
			modules.Logger.Error(err)
			return err
		}

		time.Sleep(time.Millisecond * time.Duration(timeout))
	}
}

func GetRealIP(r *http.Request) string {
	forwardedFor := r.Header.Get("X-Forwarded-For")

	if forwardedFor == "" {
		return strings.Split(r.RemoteAddr, ":")[0]
	}

	ips := strings.Split(forwardedFor, ", ")
	return ips[len(ips)-1]
}
