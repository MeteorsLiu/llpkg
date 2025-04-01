package lua

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const COLIBNAME = "coroutine"
const ABLIBNAME = "table"
const IOLIBNAME = "io"
const OSLIBNAME = "os"
const STRLIBNAME = "string"
const UTF8LIBNAME = "utf8"
const MATHLIBNAME = "math"
const DBLIBNAME = "debug"
const LOADLIBNAME = "package"

// llgo:link (*State).Base C.luaopen_base
func (recv_ *State) Base() c.Int {
	return 0
}

// llgo:link (*State).Coroutine C.luaopen_coroutine
func (recv_ *State) Coroutine() c.Int {
	return 0
}

// llgo:link (*State).Table C.luaopen_table
func (recv_ *State) Table() c.Int {
	return 0
}

// llgo:link (*State).Io C.luaopen_io
func (recv_ *State) Io() c.Int {
	return 0
}

// llgo:link (*State).Os C.luaopen_os
func (recv_ *State) Os() c.Int {
	return 0
}

// llgo:link (*State).String C.luaopen_string
func (recv_ *State) String() c.Int {
	return 0
}

// llgo:link (*State).Utf8 C.luaopen_utf8
func (recv_ *State) Utf8() c.Int {
	return 0
}

// llgo:link (*State).Math C.luaopen_math
func (recv_ *State) Math() c.Int {
	return 0
}

// llgo:link (*State).Debug C.luaopen_debug
func (recv_ *State) Debug() c.Int {
	return 0
}

// llgo:link (*State).Package C.luaopen_package
func (recv_ *State) Package() c.Int {
	return 0
}

/* open all previous libraries */
// llgo:link (*State).Openlibs C.luaL_openlibs
func (recv_ *State) Openlibs() {
}
