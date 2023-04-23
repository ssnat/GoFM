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

	Store    *sync.Map
	StoreKey string

	Lock sync.RWMutex
}

var MusicReader = IMusicReader{
	InitialFrame: 500,
	UnitFrame:    50,

	Index: 0,
	File:  nil,

	Store:    &sync.Map{},
	StoreKey: "Store",
}

type IMusicReaderStoreData struct {
	InitialBuffer []byte
	UnitBuffer    []byte
	Timeout       int
	Order         int
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

func (musicReader *IMusicReader) GetStoreData() *IMusicReaderStoreData {
	store, ok := musicReader.Store.Load(musicReader.StoreKey)
	if !ok {
		return nil
	}
	data := store.(IMusicReaderStoreData)
	return &data
}

func (musicReader *IMusicReader) SetStoreData(data IMusicReaderStoreData) {
	musicReader.Store.Store(musicReader.StoreKey, data)
}

func (musicReader *IMusicReader) Sleep() {
	store := musicReader.GetStoreData()
	if store != nil {
		time.Sleep(time.Millisecond * time.Duration(store.Timeout))
	}
}

func (musicReader *IMusicReader) SetInitialBuffer() {
	var initialBuffer []byte
	var unitBuffer []byte

	var timeout = 0

	for i := 0; i < musicReader.InitialFrame; i++ {
		frame := mp3lib.NextFrame(musicReader.File)
		if frame == nil {
			musicReader.CloseFile()
			continue
		}
		initialBuffer = append(initialBuffer, frame.RawBytes...)

		if i >= musicReader.InitialFrame-musicReader.UnitFrame {

			unitBuffer = append(unitBuffer, frame.RawBytes...)

			timeout += 1000 * frame.SampleCount / frame.SamplingRate
		}
	}

	musicReader.SetStoreData(IMusicReaderStoreData{
		InitialBuffer: initialBuffer,
		UnitBuffer:    unitBuffer,
		Timeout:       timeout,
		Order:         1,
	})
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

	store := musicReader.GetStoreData()

	initialBuffer := store.InitialBuffer[:]
	initialBuffer = initialBuffer[len(unitBuffer):]
	initialBuffer = append(initialBuffer, unitBuffer...)

	musicReader.SetStoreData(IMusicReaderStoreData{
		InitialBuffer: initialBuffer,
		UnitBuffer:    unitBuffer,
		Timeout:       timeout,
		Order:         store.Order + 1,
	})
}

func (musicReader *IMusicReader) StartLoop() {
	for {
		if musicReader.NoFile() {
			musicReader.SelectNextMusic()
		}
		store := musicReader.GetStoreData()
		if store == nil {
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
