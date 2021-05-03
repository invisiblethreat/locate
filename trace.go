package locate

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Location returns the current file and line
func Location() string {
	pc, file, _, ok := runtime.Caller(1)
	if !ok {
		return "lost"
	}
	fn := runtime.FuncForPC(pc)
	return fmt.Sprintf("%s:%s", fn.Name(), filepath.Base(file))
}

func FileLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "lost"
	}

	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

func FullFileLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "lost"
	}

	return fmt.Sprintf("%s:%d", file, line)
}

type Formatter struct {
	Package  bool
	Function bool
	FullPath bool
	File     bool
	Line     bool
}
