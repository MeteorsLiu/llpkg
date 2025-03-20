package libtool

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const LTDL_H = 1

type LtHandle struct {
	Unused [8]uint8
}
type LtDlhandle *LtHandle

/* Initialisation and finalisation functions for libltdl. */
//go:linkname LtDlinit C.lt_dlinit
func LtDlinit() c.Int

//go:linkname LtDlexit C.lt_dlexit
func LtDlexit() c.Int

/* Module search path manipulation.  */
//go:linkname LtDladdsearchdir C.lt_dladdsearchdir
func LtDladdsearchdir(search_dir *int8) c.Int

//go:linkname LtDlinsertsearchdir C.lt_dlinsertsearchdir
func LtDlinsertsearchdir(before *int8, search_dir *int8) c.Int

//go:linkname LtDlsetsearchpath C.lt_dlsetsearchpath
func LtDlsetsearchpath(search_path *int8) c.Int

//go:linkname LtDlgetsearchpath C.lt_dlgetsearchpath
func LtDlgetsearchpath() *int8

//go:linkname LtDlforeachfile C.lt_dlforeachfile
func LtDlforeachfile(search_path *int8, func_ func(*int8, unsafe.Pointer) c.Int, data unsafe.Pointer) c.Int

/* User module loading advisors.  */
//go:linkname LtDladviseInit C.lt_dladvise_init
func LtDladviseInit(advise *LtDladvise) c.Int

//go:linkname LtDladviseDestroy C.lt_dladvise_destroy
func LtDladviseDestroy(advise *LtDladvise) c.Int

//go:linkname LtDladviseExt C.lt_dladvise_ext
func LtDladviseExt(advise *LtDladvise) c.Int

//go:linkname LtDladviseResident C.lt_dladvise_resident
func LtDladviseResident(advise *LtDladvise) c.Int

//go:linkname LtDladviseLocal C.lt_dladvise_local
func LtDladviseLocal(advise *LtDladvise) c.Int

//go:linkname LtDladviseGlobal C.lt_dladvise_global
func LtDladviseGlobal(advise *LtDladvise) c.Int

//go:linkname LtDladvisePreload C.lt_dladvise_preload
func LtDladvisePreload(advise *LtDladvise) c.Int

/* Portable libltdl versions of the system dlopen() API. */
//go:linkname LtDlopen C.lt_dlopen
func LtDlopen(filename *int8) LtDlhandle

//go:linkname LtDlopenext C.lt_dlopenext
func LtDlopenext(filename *int8) LtDlhandle

//go:linkname LtDlopenadvise C.lt_dlopenadvise
func LtDlopenadvise(filename *int8, advise LtDladvise) LtDlhandle

//go:linkname LtDlsym C.lt_dlsym
func LtDlsym(handle LtDlhandle, name *int8) unsafe.Pointer

//go:linkname LtDlerror C.lt_dlerror
func LtDlerror() *int8

//go:linkname LtDlclose C.lt_dlclose
func LtDlclose(handle LtDlhandle) c.Int

/*
A preopened symbol. Arrays of this type comprise the exported

	symbols for a dlpreopened module.
*/
type LtDlsymlist struct {
	Name    *int8
	Address unsafe.Pointer
}

// llgo:type C
type LtDlpreloadCallbackFunc func(LtDlhandle) c.Int

// llgo:link (*LtDlsymlist).LtDlpreload C.lt_dlpreload
func (recv_ *LtDlsymlist) LtDlpreload() c.Int {
	return 0
}

// llgo:link (*LtDlsymlist).LtDlpreloadDefault C.lt_dlpreload_default
func (recv_ *LtDlsymlist) LtDlpreloadDefault() c.Int {
	return 0
}

//go:linkname LtDlpreloadOpen C.lt_dlpreload_open
func LtDlpreloadOpen(originator *int8, func_ LtDlpreloadCallbackFunc) c.Int

type LtDlinterfaceId unsafe.Pointer

// llgo:type C
type LtDlhandleInterface func(LtDlhandle, *int8) c.Int

//go:linkname LtDlinterfaceRegister C.lt_dlinterface_register
func LtDlinterfaceRegister(id_string *int8, iface LtDlhandleInterface) LtDlinterfaceId

//go:linkname LtDlinterfaceFree C.lt_dlinterface_free
func LtDlinterfaceFree(key LtDlinterfaceId)

//go:linkname LtDlcallerSetData C.lt_dlcaller_set_data
func LtDlcallerSetData(key LtDlinterfaceId, handle LtDlhandle, data unsafe.Pointer) unsafe.Pointer

//go:linkname LtDlcallerGetData C.lt_dlcaller_get_data
func LtDlcallerGetData(key LtDlinterfaceId, handle LtDlhandle) unsafe.Pointer

/* Read only information pertaining to a loaded module. */

type LtDlinfo struct {
	Filename    *int8
	Name        *int8
	RefCount    c.Int
	IsResident  c.Uint
	IsSymglobal c.Uint
	IsSymlocal  c.Uint
}

//go:linkname LtDlgetinfo C.lt_dlgetinfo
func LtDlgetinfo(handle LtDlhandle) *LtDlinfo

//go:linkname LtDlhandleIterate C.lt_dlhandle_iterate
func LtDlhandleIterate(iface LtDlinterfaceId, place LtDlhandle) LtDlhandle

//go:linkname LtDlhandleFetch C.lt_dlhandle_fetch
func LtDlhandleFetch(iface LtDlinterfaceId, module_name *int8) LtDlhandle

//go:linkname LtDlhandleMap C.lt_dlhandle_map
func LtDlhandleMap(iface LtDlinterfaceId, func_ func(LtDlhandle, unsafe.Pointer) c.Int, data unsafe.Pointer) c.Int

/* Deprecated module residency management API. */
//go:linkname LtDlmakeresident C.lt_dlmakeresident
func LtDlmakeresident(handle LtDlhandle) c.Int

//go:linkname LtDlisresident C.lt_dlisresident
func LtDlisresident(handle LtDlhandle) c.Int
