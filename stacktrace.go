package errors

import (
	"runtime"
	"strings"
)

type frame struct {
	runtime.Frame
}

func (f frame) file() string {
	return f.Frame.File
}

func (f frame) line() int {
	return f.Frame.Line
}

func (f frame) function() string {
	name := f.Frame.Function
	i := strings.LastIndexByte(name, '.')
	return name[i+1:]
}

func stackLevel(skip int) (f frame) {
	var frame [3]uintptr
	runtime.Callers(skip+2, frame[:])
	frames := runtime.CallersFrames(frame[:])
	f.Frame, _ = frames.Next()
	return
}
