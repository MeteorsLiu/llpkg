package libtool

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const LT_DLLOADER_H = 1

type LtDlloader unsafe.Pointer
type LtModule unsafe.Pointer
type LtUserData unsafe.Pointer

type LtAdvise struct {
	Unused [8]uint8
}
type LtDladvise *LtAdvise

// llgo:type C
type LtModuleOpen func(LtUserData, *int8, LtDladvise) LtModule

// llgo:type C
type LtModuleClose func(LtUserData, LtModule) c.Int

// llgo:type C
type LtFindSym func(LtUserData, LtModule, *int8) unsafe.Pointer

// llgo:type C
type LtDlloaderInit func(LtUserData) c.Int

// llgo:type C
type LtDlloaderExit func(LtUserData) c.Int
type LtDlloaderPriority c.Int

const (
	LTDLLOADERPREPEND LtDlloaderPriority = 0
	LTDLLOADERAPPEND  LtDlloaderPriority = 1
)

/*
This structure defines a module loader, as populated by the get_vtable

	entry point of each loader.
*/
type LtDlvtable struct {
	Name         *int8
	SymPrefix    *int8
	ModuleOpen   *unsafe.Pointer
	ModuleClose  *unsafe.Pointer
	FindSym      *unsafe.Pointer
	DlloaderInit *unsafe.Pointer
	DlloaderExit *unsafe.Pointer
	DlloaderData LtUserData
	Priority     LtDlloaderPriority
}

// llgo:link (*LtDlvtable).LtDlloaderAdd C.lt_dlloader_add
func (recv_ *LtDlvtable) LtDlloaderAdd() c.Int {
	return 0
}

//go:linkname LtDlloaderNext C.lt_dlloader_next
func LtDlloaderNext(loader LtDlloader) LtDlloader

//go:linkname LtDlloaderRemove C.lt_dlloader_remove
func LtDlloaderRemove(name *int8) *LtDlvtable

//go:linkname LtDlloaderFind C.lt_dlloader_find
func LtDlloaderFind(name *int8) *LtDlvtable

//go:linkname LtDlloaderGet C.lt_dlloader_get
func LtDlloaderGet(loader LtDlloader) *LtDlvtable

// llgo:type C
type LtGetVtable func(LtUserData) *LtDlvtable
