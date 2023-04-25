GoFM
=====
GoFM is a cross-platform real-time audio streaming server written in Go. It allows you to stream MP3 audio files from a specified directory over HTTP and synchronize playback across multiple clients.

[![Go Version](https://img.shields.io/badge/Go-v1.16-blue)](https://golang.org/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/PxGo/GoFM)](https://goreportcard.com/report/github.com/PxGo/GoFM)
[![Downloads](https://img.shields.io/github/downloads/PxGo/GoFM/total)](https://github.com/PxGo/GoFM/releases)
[![References](https://img.shields.io/github/forks/PxGo/GoFM?label=references)](https://github.com/PxGo/GoFM/network/members)
[![License](https://img.shields.io/github/license/PxGo/GoFM)](https://github.com/PxGo/GoFM/blob/main/LICENSE)


Usage
-----
To use GoFM, download the latest release from the [Releases page](https://github.com/PxGo/GoFM/releases) and run the server with the following command:

```
./GoFM -d /path/to/your/music/directory
```

By default, GoFM listens on port 8090 and all available network interfaces. You can customize the server's behavior using the following command-line options:

```
-p int
    Specifies the server port number (default 8090).
-host string
    Specifies the server host address (default "0.0.0.0").
-r
    Enables random playback mode.
-debug
    Enables debug mode for the server.
-d string
    Specifies the directory to play audio files from (default "/path/to/your/music/directory").
-h
    Shows help information.
```

For example, to change the server's port number to 8080, use the -p option followed by the desired port number, like this:
```
./GoFM -d /path/to/your/music/directory -p 8080
```
Available Online Services
-----
PxGo FM: [http://fm.stackstore.net:2086](http://fm.stackstore.net:2086)

License
-----
GoFM is released under the [MIT License](https://github.com/PxGo/GoFM/blob/main/LICENSE). Feel free to use, modify, and distribute the software. Contributions are welcome!