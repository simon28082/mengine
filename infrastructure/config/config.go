package config

import (
	"github.com/spf13/viper"
	"time"
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	Get(path string) Value
	Set(path string, val any)
	Delete(path string) error
}

// Value represents a value of any type
type Value interface {
	Exists() bool
	Bool() bool
	BoolDef(def bool) bool
	Int(def int) int
	String(def string) string
	Float64(def float64) float64
	Duration(def time.Duration) time.Duration
	StringSlice(def []string) []string
	StringMap(def map[string]string) map[string]string
	Scan(val interface{}) error
	Bytes() []byte
}

func a() {
	//mapstructure.DecodeHookFuncType()
	viper.SetDefault("a.b.c", "content")
	viper.ReadInConfig()
	//viper.GetViper().Set()
	//viper.UnsupportedRemoteProviderError()
	//viper.Get
}
