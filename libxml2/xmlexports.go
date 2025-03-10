package libxml2

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

//go:linkname CheckVersion C.xmlCheckVersion
func CheckVersion(version c.Int)
