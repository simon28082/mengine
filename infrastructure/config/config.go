package config

import (
	"bytes"
	"github.com/simon28082/mengine/infrastructure/config/source"
	"github.com/sourcegraph/conc/panics"
	"github.com/spf13/viper"
)

type Config interface {
	Get(path string) Value

	Set(path string, val any)

	Delete(path string) error
}

//type Reloader interface {
//	Reload(contents []byte) error
//}

type config struct {
	viper  *viper.Viper
	source source.Source
}

//var WireConfigFileSourceSet = wire.NewSet(NewConfig, source.WireFileSourceSet)

func NewConfig(source source.Source) (*config, error) {
	reader, err := source.Read()
	if err != nil {
		return nil, err
	}
	vp := viper.New()
	err1 := vp.ReadConfig(reader)
	if err1 != nil {
		return nil, err1
	}

	c := &config{
		viper:  vp,
		source: source,
	}
	if err := c.watch(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *config) Get(path string) Value {
	v := c.viper.Get(path)
	return &value{v: v}
}

func (c *config) Set(path string, val any) {
	c.viper.Set(path, val)
}

func (c *config) Delete(path string) error {
	c.viper.Set(path, nil)
	return nil
}

func (c *config) watch() error {
	notifier, ok := c.source.(source.Notifier)
	if !ok {
		return nil
	}

	notification, err := notifier.Notify()
	if err != nil {
		return err
	}

	var pc panics.Catcher
	go pc.Try(func() {
		for v := range notification {
			if err != nil {
				return
			}
			c.viper.ReadConfig(bytes.NewBuffer(v))
		}
	})

	return pc.Recovered().AsError()
}

//func (c *config) Reload(contents []byte) error {
//	return c.viper.ReadConfig(bytes.NewBuffer(contents))
//}
