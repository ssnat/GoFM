package modules

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func InitServer() {

	addr := fmt.Sprintf("%s:%d", Config.Host, Config.Port)

	http.HandleFunc("/", FMHandle)

	go func() {
		log.Fatal(http.ListenAndServe(addr, nil))
	}()

	Logger.Info(fmt.Sprintf("Server is running at %s.", addr))
}

func FMHandle(w http.ResponseWriter, r *http.Request) {

	ip := GetRealIP(r)

	Logger.Info(fmt.Sprintf("Client %s connected", ip))

	store := MusicReader.GetStoreData()
	if store == nil {
		_, err := io.WriteString(w, "Oops, it seems like the FM hasn't started up.")
		if err != nil {
			Logger.Error(err)
		}
		return
	}

	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "audio/mpeg")

	init := false

	for {
		var targetBuffer []byte

		store := MusicReader.GetStoreData()

		if !init {
			init = true
			targetBuffer = store.InitialBuffer[:]
		} else {
			targetBuffer = store.UnitBuffer[:]
		}

		var timeout = store.Timeout

		_, err := w.Write(targetBuffer)
		if err != nil {
			Logger.Error(err)
			return
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
