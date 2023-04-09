package config

import (
	"encoding/json"
	"time"
)

// Value represents a value of any type
type Value interface {
	Bool() bool
	BoolDef(def bool) bool
	Int() int
	IntDef(def int) int
	StringDef(def string) string
	String(def string) string
	Float64(def float64) float64
	Float64Def(def float64) float64
	Duration() time.Duration
	DurationDef(def time.Duration) time.Duration
	StringSlice() []string
	StringSliceDef(def []string) []string
	StringMap() map[string]string
	StringMapDef(def map[string]string) map[string]string
	Scan(val any) error
	Bytes() []byte
}

type value struct {
	v any
}

func valueDef[T any](v *value, def T) T {
	if vs, ok := v.v.(T); ok {
		return vs
	}
	return def
}

func (v *value) Bool() bool {
	return v.BoolDef(false)
}

func (v *value) BoolDef(def bool) bool {
	return valueDef(v, def)
}

func (v value) Int() int {
	return v.IntDef(0)
}

func (v *value) IntDef(def int) int {
	return valueDef(v, def)
}

func (v *value) String(def string) string {
	return valueDef(v, def)
}

func (v *value) StringDef(def string) string {
	return v.StringDef(def)
}

func (v *value) Float64(def float64) float64 {
	return v.Float64(0)
}

func (v *value) Scan(val any) error {
	return json.Unmarshal(valueDef(v, []byte{}), &val)
}

func (v *value) Bytes() []byte {
	return valueDef(v, []byte{})
}

func (v *value) Float64Def(def float64) float64 {
	return valueDef(v, def)
}

func (v *value) Duration() time.Duration {
	return v.DurationDef(0)
}

func (v *value) DurationDef(def time.Duration) time.Duration {
	return valueDef(v, def)
}

func (v *value) StringSlice() []string {
	return v.StringSliceDef([]string{})
}

func (v *value) StringSliceDef(def []string) []string {
	return valueDef(v, def)
}

func (v *value) StringMap() map[string]string {
	return v.StringMapDef(map[string]string{})
}

func (v *value) StringMapDef(def map[string]string) map[string]string {
	return valueDef(v, def)
}
