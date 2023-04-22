package modules

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type IConfig struct {
	Port      int
	Host      string
	Directory string
	Random    bool
	Debug     bool
}

var Config *IConfig

func init() {

	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var port int
	var host string
	var random bool
	var directory string
	var debug bool
	var help bool

	flag.IntVar(&port, "p", 8090, "server port number")
	flag.StringVar(&host, "host", "0.0.0.0", "server host address")
	flag.BoolVar(&random, "r", false, "enable random playback mode")
	flag.BoolVar(&debug, "debug", false, "enable debug mode for server")
	flag.StringVar(&directory, "d", root, "directory to play")
	flag.BoolVar(&help, "h", false, "show help information")

	flag.Parse()

	if help {
		fmt.Println("Usage: GoFM [options]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	directory, err = filepath.Abs(directory)

	if err != nil {
		log.Fatal(err)
	}

	Config = &IConfig{
		Port:      port,
		Host:      host,
		Random:    random,
		Directory: directory,
		Debug:     debug,
	}
}

func GetConfig() *IConfig {
	return Config
}
