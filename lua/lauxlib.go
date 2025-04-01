package lua

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const GNAME = "_G"
const LOADED_TABLE = "_LOADED"
const PRELOAD_TABLE = "_PRELOAD"
const FILEHANDLE = "FILE*"

type Buffer struct {
	B    *int8
	Size uintptr
	N    uintptr
	L    *State
	Init struct {
		B [1024]int8
	}
}

type Reg struct {
	Name *int8
	Func unsafe.Pointer
}

// llgo:link (*State).Checkversion_ C.luaL_checkversion_
func (recv_ *State) Checkversion_(ver Number, sz uintptr) {
}

// llgo:link (*State).Getmetafield C.luaL_getmetafield
func (recv_ *State) Getmetafield(obj c.Int, e *int8) c.Int {
	return 0
}

// llgo:link (*State).Callmeta C.luaL_callmeta
func (recv_ *State) Callmeta(obj c.Int, e *int8) c.Int {
	return 0
}

// llgo:link (*State).Tolstring__1 C.luaL_tolstring
func (recv_ *State) Tolstring__1(idx c.Int, len *uintptr) *int8 {
	return nil
}

// llgo:link (*State).Argerror C.luaL_argerror
func (recv_ *State) Argerror(arg c.Int, extramsg *int8) c.Int {
	return 0
}

// llgo:link (*State).Typeerror C.luaL_typeerror
func (recv_ *State) Typeerror(arg c.Int, tname *int8) c.Int {
	return 0
}

// llgo:link (*State).Checklstring C.luaL_checklstring
func (recv_ *State) Checklstring(arg c.Int, l *uintptr) *int8 {
	return nil
}

// llgo:link (*State).Optlstring C.luaL_optlstring
func (recv_ *State) Optlstring(arg c.Int, def *int8, l *uintptr) *int8 {
	return nil
}

// llgo:link (*State).Checknumber C.luaL_checknumber
func (recv_ *State) Checknumber(arg c.Int) Number {
	return 0
}

// llgo:link (*State).Optnumber C.luaL_optnumber
func (recv_ *State) Optnumber(arg c.Int, def Number) Number {
	return 0
}

// llgo:link (*State).Checkinteger C.luaL_checkinteger
func (recv_ *State) Checkinteger(arg c.Int) Integer {
	return 0
}

// llgo:link (*State).Optinteger C.luaL_optinteger
func (recv_ *State) Optinteger(arg c.Int, def Integer) Integer {
	return 0
}

// llgo:link (*State).Checkstack__1 C.luaL_checkstack
func (recv_ *State) Checkstack__1(sz c.Int, msg *int8) {
}

// llgo:link (*State).Checktype C.luaL_checktype
func (recv_ *State) Checktype(arg c.Int, t c.Int) {
}

// llgo:link (*State).Checkany C.luaL_checkany
func (recv_ *State) Checkany(arg c.Int) {
}

// llgo:link (*State).Newmetatable C.luaL_newmetatable
func (recv_ *State) Newmetatable(tname *int8) c.Int {
	return 0
}

// llgo:link (*State).Setmetatable__1 C.luaL_setmetatable
func (recv_ *State) Setmetatable__1(tname *int8) {
}

// llgo:link (*State).Testudata C.luaL_testudata
func (recv_ *State) Testudata(ud c.Int, tname *int8) unsafe.Pointer {
	return nil
}

// llgo:link (*State).Checkudata C.luaL_checkudata
func (recv_ *State) Checkudata(ud c.Int, tname *int8) unsafe.Pointer {
	return nil
}

// llgo:link (*State).Where C.luaL_where
func (recv_ *State) Where(lvl c.Int) {
}

// llgo:link (*State).Error__1 C.luaL_error
func (recv_ *State) Error__1(fmt *int8, __llgo_va_list ...interface{}) c.Int {
	return 0
}

// llgo:link (*State).Checkoption C.luaL_checkoption
func (recv_ *State) Checkoption(arg c.Int, def *int8, lst **int8) c.Int {
	return 0
}

// llgo:link (*State).Fileresult C.luaL_fileresult
func (recv_ *State) Fileresult(stat c.Int, fname *int8) c.Int {
	return 0
}

// llgo:link (*State).Execresult C.luaL_execresult
func (recv_ *State) Execresult(stat c.Int) c.Int {
	return 0
}

// llgo:link (*State).Ref C.luaL_ref
func (recv_ *State) Ref(t c.Int) c.Int {
	return 0
}

// llgo:link (*State).Unref C.luaL_unref
func (recv_ *State) Unref(t c.Int, ref c.Int) {
}

// llgo:link (*State).Loadfilex C.luaL_loadfilex
func (recv_ *State) Loadfilex(filename *int8, mode *int8) c.Int {
	return 0
}

// llgo:link (*State).Loadbufferx C.luaL_loadbufferx
func (recv_ *State) Loadbufferx(buff *int8, sz uintptr, name *int8, mode *int8) c.Int {
	return 0
}

// llgo:link (*State).Loadstring C.luaL_loadstring
func (recv_ *State) Loadstring(s *int8) c.Int {
	return 0
}

//go:linkname Newstate__1 C.luaL_newstate
func Newstate__1() *State

// llgo:link (*State).Len__1 C.luaL_len
func (recv_ *State) Len__1(idx c.Int) Integer {
	return 0
}

// llgo:link (*Buffer).Addgsub C.luaL_addgsub
func (recv_ *Buffer) Addgsub(s *int8, p *int8, r *int8) {
}

// llgo:link (*State).Gsub C.luaL_gsub
func (recv_ *State) Gsub(s *int8, p *int8, r *int8) *int8 {
	return nil
}

// llgo:link (*State).Setfuncs C.luaL_setfuncs
func (recv_ *State) Setfuncs(l *Reg, nup c.Int) {
}

// llgo:link (*State).Getsubtable C.luaL_getsubtable
func (recv_ *State) Getsubtable(idx c.Int, fname *int8) c.Int {
	return 0
}

// llgo:link (*State).Traceback C.luaL_traceback
func (recv_ *State) Traceback(L1 *State, msg *int8, level c.Int) {
}

// llgo:link (*State).Requiref C.luaL_requiref
func (recv_ *State) Requiref(modname *int8, openf CFunction, glb c.Int) {
}

// llgo:link (*State).Buffinit C.luaL_buffinit
func (recv_ *State) Buffinit(B *Buffer) {
}

// llgo:link (*Buffer).Prepbuffsize C.luaL_prepbuffsize
func (recv_ *Buffer) Prepbuffsize(sz uintptr) *int8 {
	return nil
}

// llgo:link (*Buffer).Addlstring C.luaL_addlstring
func (recv_ *Buffer) Addlstring(s *int8, l uintptr) {
}

// llgo:link (*Buffer).Addstring C.luaL_addstring
func (recv_ *Buffer) Addstring(s *int8) {
}

// llgo:link (*Buffer).Addvalue C.luaL_addvalue
func (recv_ *Buffer) Addvalue() {
}

// llgo:link (*Buffer).Pushresult C.luaL_pushresult
func (recv_ *Buffer) Pushresult() {
}

// llgo:link (*Buffer).Pushresultsize C.luaL_pushresultsize
func (recv_ *Buffer) Pushresultsize(sz uintptr) {
}

// llgo:link (*State).Buffinitsize C.luaL_buffinitsize
func (recv_ *State) Buffinitsize(B *Buffer, sz uintptr) *int8 {
	return nil
}

type Stream struct {
	F      *c.FILE
	Closef unsafe.Pointer
}
