package os

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

//func FileExtension(path string) string {
//	return filepath.Ext(path)
//}

func callerStepFile(skip int) string {
	_, file, _, ok := runtime.Caller(skip)
	if !ok {
		panic(`Can not get current file info`)
	}

	return file
}

// RunFile Get current running file path
func RunFile() string {
	// 0 is current file, so except
	// 1 is current file, so except
	return callerStepFile(2)
}

func RunPath(paths ...string) string {
	dir := path.Dir(callerStepFile(2))
	if len(paths) == 0 {
		return dir
	}

	return filepath.Join(dir, filepath.Join(paths...))
}

// RunDir Get current running directory path
func RunDir() string {
	return path.Dir(callerStepFile(2))
}

// RunRelative Get current directory relative path
func RunRelative(rpath string) string {
	rpath, _ = filepath.Abs(filepath.Join(path.Dir(callerStepFile(2)), rpath))
	return rpath
}

func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}

func Mkdir(path string, m os.FileMode) (err error) {
	err = os.MkdirAll(path, m)
	return
}

func MkdirDefault(path string) (err error) {
	err = os.MkdirAll(path, 0555)
	return
}
