package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llpkg/libxml2"
	"unsafe"
)

type Locale unsafe.Pointer
type LocaleChar libxml2.Char

//go:linkname NewLocale C.xsltNewLocale
func NewLocale(langName *libxml2.Char) Locale

//go:linkname FreeLocale C.xsltFreeLocale
func FreeLocale(locale Locale)

//go:linkname Strxfrm C.xsltStrxfrm
func Strxfrm(locale Locale, string *libxml2.Char) *LocaleChar

//go:linkname LocaleStrcmp C.xsltLocaleStrcmp
func LocaleStrcmp(locale Locale, str1 *LocaleChar, str2 *LocaleChar) c.Int

//go:linkname FreeLocales C.xsltFreeLocales
func FreeLocales()
