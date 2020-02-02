package errors

import (
	"runtime"
	"strings"
)

// frame wraps a runtime.Frame to provide some helper functions while still allowing access to
// the original runtime.Frame
type frame struct {
	runtime.Frame
}

// file is the runtime.Frame.File stripped down to just the filename
func (f frame) file() string {
	return f.Frame.File
}

// line is the line of the runtime.Frame and exposed for convenience.
func (f frame) line() int {
	return f.Frame.Line
}

// function is the runtime.Frame.Function stripped down to just the function name
func (f frame) function() string {
	name := f.Frame.Function
	i := strings.LastIndexByte(name, '.')
	return name[i+1:]
}

// stack returns a stack Frame
func stack() frame {
	return stackLevel(1)
}

// stackLevel returns a stack Frame skipping the number of supplied frames.
// This is primarily used by other libraries who use this package
// internally as the additional.
func stackLevel(skip int) (f frame) {
	var frame [3]uintptr
	runtime.Callers(skip+2, frame[:])
	frames := runtime.CallersFrames(frame[:])
	f.Frame, _ = frames.Next()
	return
}
