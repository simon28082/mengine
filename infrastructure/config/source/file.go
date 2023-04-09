package source

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"io"
	"os"
)

type file struct {
	path   string
	viper  *viper.Viper
	notify chan []byte
}

type PathString string

var WireFileSourceSet = wire.NewSet(NewFile, wire.Value(`abc`))

func NewFile(path string) *file {
	vp := viper.New()
	vp.SetConfigFile(string(path))
	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		//f.config.Reload()
	})
	vp.WatchConfig()
	return &file{
		path:   string(path),
		viper:  vp,
		notify: make(chan []byte, 0),
	}
}

func (f *file) Read() (io.Reader, error) {
	if err := f.viper.ReadInConfig(); err != nil {
		return nil, err
	}

	contentBytes, err := os.ReadFile(f.viper.ConfigFileUsed())
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(contentBytes), nil
}

func (f *file) Close() error {
	close(f.notify)
	return nil
}

func (f *file) Notify() (<-chan []byte, error) {
	return f.notify, nil
}
