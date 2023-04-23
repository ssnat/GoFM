package modules

import (
	"fmt"
	"github.com/dmulholl/mp3lib"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type IMusicReader struct {
	InitialFrame int
	UnitFrame    int

	Index int
	File  *os.File

	InitialBuffer []byte
	UnitBuffer    []byte
	Timeout       int

	Lock sync.RWMutex
}

var MusicReader = IMusicReader{
	InitialFrame: 500,
	UnitFrame:    50,

	Index: 0,
	File:  nil,

	InitialBuffer: nil,
	UnitBuffer:    nil,
	Timeout:       0,
}

func (musicReader *IMusicReader) SelectNextMusic() {
	mp3FilePaths, err := GetMp3FilePaths()
	if err != nil {
		Logger.Error(err)
		return
	}
	if Config.Random {
		MusicReader.Index = rand.Intn(len(mp3FilePaths))
	} else {
		MusicReader.Index += 1
		if MusicReader.Index >= len(mp3FilePaths) {
			MusicReader.Index = 0
		}
	}

	filePath := mp3FilePaths[MusicReader.Index]
	file, err := os.Open(filePath)
	if err != nil {
		Logger.Error(err)
		musicReader.SelectNextMusic()
		return
	}
	MusicReader.File = file
}

func (musicReader *IMusicReader) CloseFile() {
	if musicReader.File != nil {
		err := musicReader.File.Close()
		if err != nil {
			Logger.Error(err)
		}
		musicReader.File = nil
	}
}

func (musicReader *IMusicReader) NoFile() bool {
	return musicReader.File == nil
}

func (musicReader *IMusicReader) Sleep() {
	time.Sleep(time.Millisecond * time.Duration(musicReader.Timeout))
}

func (musicReader *IMusicReader) SetInitialBuffer() {
	var initBuffer []byte
	var unitBuffer []byte

	var timeout = 0

	for i := 0; i < musicReader.InitialFrame; i++ {
		frame := mp3lib.NextFrame(musicReader.File)
		if frame == nil {
			musicReader.CloseFile()
			continue
		}
		initBuffer = append(initBuffer, frame.RawBytes...)

		if i >= musicReader.InitialFrame-musicReader.UnitFrame {

			unitBuffer = append(unitBuffer, frame.RawBytes...)

			timeout += 1000 * frame.SampleCount / frame.SamplingRate
		}
	}

	musicReader.Lock.Lock()
	musicReader.InitialBuffer = initBuffer
	musicReader.UnitBuffer = unitBuffer
	musicReader.Timeout = timeout
	musicReader.Lock.Unlock()
}

func (musicReader *IMusicReader) SetUnitBuffer() {
	var unitBuffer []byte

	var timeout = 0

	for i := 0; i < musicReader.UnitFrame; i++ {
		frame := mp3lib.NextFrame(musicReader.File)
		if frame == nil {
			musicReader.CloseFile()
			continue
		}
		unitBuffer = append(unitBuffer, frame.RawBytes...)
		timeout += 1000 * frame.SampleCount / frame.SamplingRate
	}

	initialBuffer := musicReader.InitialBuffer[:]
	initialBuffer = initialBuffer[len(unitBuffer):]
	initialBuffer = append(initialBuffer, unitBuffer...)

	musicReader.Lock.Lock()
	musicReader.InitialBuffer = initialBuffer
	musicReader.UnitBuffer = unitBuffer
	musicReader.Timeout = timeout
	musicReader.Lock.Unlock()
}

func (musicReader *IMusicReader) StartLoop() {
	for {
		if musicReader.NoFile() {
			musicReader.SelectNextMusic()
		}
		if musicReader.InitialBuffer == nil {
			musicReader.SetInitialBuffer()
		} else {
			musicReader.SetUnitBuffer()
		}
		musicReader.Sleep()
	}
}

func GetMp3FilePaths() ([]string, error) {
	var mp3Files []string
	err := filepath.Walk(Config.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".mp3") {
			mp3Files = append(mp3Files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(mp3Files) == 0 {
		Logger.Error(fmt.Sprintf("There are no MP3 files in the music directory."))
		time.Sleep(time.Second * 5)
		return GetMp3FilePaths()
	}
	return mp3Files, nil
}

func InitReader() {
	go func() {
		MusicReader.StartLoop()
	}()
	Logger.Info(fmt.Sprintf("Music directory is %s.", Config.Directory))
}
