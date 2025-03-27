package libxml2

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const LIBXML_DOTTED_VERSION = "2.11.6"
const LIBXML_VERSION = 21106
const LIBXML_VERSION_STRING = "21106"
const LIBXML_VERSION_EXTRA = ""
const LIBXML_MODULE_EXTENSION = ".so"

/*
 * use those to be sure nothing nasty will happen if
 * your library and includes mismatch
 */
//go:linkname CheckVersion C.xmlCheckVersion
func CheckVersion(version c.Int)
