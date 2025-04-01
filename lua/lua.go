package lua

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const VERSION_MAJOR = "5"
const VERSION_MINOR = "4"
const VERSION_RELEASE = "7"
const VERSION_NUM = 504
const AUTHORS = "R. Ierusalimschy, L. H. de Figueiredo, W. Celes"
const SIGNATURE = "\x1bLua"
const OK = 0
const YIELD = 1
const ERRRUN = 2
const ERRSYNTAX = 3
const ERRMEM = 4
const ERRERR = 5
const NIL = 0
const BOOLEAN = 1
const LIGHTUSERDATA = 2
const NUMBER = 3
const STRING = 4
const TABLE = 5
const FUNCTION = 6
const USERDATA = 7
const THREAD = 8
const NUMTYPES = 9
const MINSTACK = 20
const RIDX_MAINTHREAD = 1
const RIDX_GLOBALS = 2
const OPADD = 0
const OPSUB = 1
const OPMUL = 2
const OPMOD = 3
const OPPOW = 4
const OPDIV = 5
const OPIDIV = 6
const OPBAND = 7
const OPBOR = 8
const OPBXOR = 9
const OPSHL = 10
const OPSHR = 11
const OPUNM = 12
const OPBNOT = 13
const OPEQ = 0
const OPLT = 1
const OPLE = 2
const GCSTOP = 0
const GCRESTART = 1
const GCCOLLECT = 2
const GCCOUNT = 3
const GCCOUNTB = 4
const GCSTEP = 5
const GCSETPAUSE = 6
const GCSETSTEPMUL = 7
const GCISRUNNING = 9
const GCGEN = 10
const GCINC = 11
const HOOKCALL = 0
const HOOKRET = 1
const HOOKLINE = 2
const HOOKCOUNT = 3
const HOOKTAILCALL = 4

type State struct {
	Unused [8]uint8
}
type Number float64
type Integer c.LongLong
type Unsigned c.UlongLong
type KContext uintptr

// llgo:type C
type CFunction func(*State) c.Int

// llgo:type C
type KFunction func(*State, c.Int, KContext) c.Int

// llgo:type C
type Reader func(*State, unsafe.Pointer, *uintptr) *int8

// llgo:type C
type Writer func(*State, unsafe.Pointer, uintptr, unsafe.Pointer) c.Int

// llgo:type C
type Alloc func(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr) unsafe.Pointer

// llgo:type C
type WarnFunction func(unsafe.Pointer, *int8, c.Int)

type Debug struct {
	Event           c.Int
	Name            *int8
	Namewhat        *int8
	What            *int8
	Source          *int8
	Srclen          uintptr
	Currentline     c.Int
	Linedefined     c.Int
	Lastlinedefined c.Int
	Nups            int8
	Nparams         int8
	Isvararg        int8
	Istailcall      int8
	Ftransfer       uint16
	Ntransfer       uint16
	ShortSrc        [60]int8
	ICi             *CallInfo
}

// llgo:type C
type Hook func(*State, *Debug)

/*
** state manipulation
 */
//go:linkname Newstate C.lua_newstate
func Newstate(f Alloc, ud unsafe.Pointer) *State

// llgo:link (*State).Close C.lua_close
func (recv_ *State) Close() {
}

// llgo:link (*State).Newthread C.lua_newthread
func (recv_ *State) Newthread() *State {
	return nil
}

// llgo:link (*State).Closethread C.lua_closethread
func (recv_ *State) Closethread(from *State) c.Int {
	return 0
}

// llgo:link (*State).Resetthread C.lua_resetthread
func (recv_ *State) Resetthread() c.Int {
	return 0
}

// llgo:link (*State).Atpanic C.lua_atpanic
func (recv_ *State) Atpanic(panicf CFunction) CFunction {
	return nil
}

// llgo:link (*State).Version C.lua_version
func (recv_ *State) Version() Number {
	return 0
}

/*
** basic stack manipulation
 */
// llgo:link (*State).Absindex C.lua_absindex
func (recv_ *State) Absindex(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Gettop C.lua_gettop
func (recv_ *State) Gettop() c.Int {
	return 0
}

// llgo:link (*State).Settop C.lua_settop
func (recv_ *State) Settop(idx c.Int) {
}

// llgo:link (*State).Pushvalue C.lua_pushvalue
func (recv_ *State) Pushvalue(idx c.Int) {
}

// llgo:link (*State).Rotate C.lua_rotate
func (recv_ *State) Rotate(idx c.Int, n c.Int) {
}

// llgo:link (*State).Copy C.lua_copy
func (recv_ *State) Copy(fromidx c.Int, toidx c.Int) {
}

// llgo:link (*State).Checkstack C.lua_checkstack
func (recv_ *State) Checkstack(n c.Int) c.Int {
	return 0
}

// llgo:link (*State).Xmove C.lua_xmove
func (recv_ *State) Xmove(to *State, n c.Int) {
}

/*
** access functions (stack -> C)
 */
// llgo:link (*State).Isnumber C.lua_isnumber
func (recv_ *State) Isnumber(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Isstring C.lua_isstring
func (recv_ *State) Isstring(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Iscfunction C.lua_iscfunction
func (recv_ *State) Iscfunction(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Isinteger C.lua_isinteger
func (recv_ *State) Isinteger(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Isuserdata C.lua_isuserdata
func (recv_ *State) Isuserdata(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Type C.lua_type
func (recv_ *State) Type(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Typename C.lua_typename
func (recv_ *State) Typename(tp c.Int) *int8 {
	return nil
}

// llgo:link (*State).Tonumberx C.lua_tonumberx
func (recv_ *State) Tonumberx(idx c.Int, isnum *c.Int) Number {
	return 0
}

// llgo:link (*State).Tointegerx C.lua_tointegerx
func (recv_ *State) Tointegerx(idx c.Int, isnum *c.Int) Integer {
	return 0
}

// llgo:link (*State).Toboolean C.lua_toboolean
func (recv_ *State) Toboolean(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Tolstring C.lua_tolstring
func (recv_ *State) Tolstring(idx c.Int, len *uintptr) *int8 {
	return nil
}

// llgo:link (*State).Rawlen C.lua_rawlen
func (recv_ *State) Rawlen(idx c.Int) Unsigned {
	return 0
}

// llgo:link (*State).Tocfunction C.lua_tocfunction
func (recv_ *State) Tocfunction(idx c.Int) CFunction {
	return nil
}

// llgo:link (*State).Touserdata C.lua_touserdata
func (recv_ *State) Touserdata(idx c.Int) unsafe.Pointer {
	return nil
}

// llgo:link (*State).Tothread C.lua_tothread
func (recv_ *State) Tothread(idx c.Int) *State {
	return nil
}

// llgo:link (*State).Topointer C.lua_topointer
func (recv_ *State) Topointer(idx c.Int) unsafe.Pointer {
	return nil
}

// llgo:link (*State).Arith C.lua_arith
func (recv_ *State) Arith(op c.Int) {
}

// llgo:link (*State).Rawequal C.lua_rawequal
func (recv_ *State) Rawequal(idx1 c.Int, idx2 c.Int) c.Int {
	return 0
}

// llgo:link (*State).Compare C.lua_compare
func (recv_ *State) Compare(idx1 c.Int, idx2 c.Int, op c.Int) c.Int {
	return 0
}

/*
** push functions (C -> stack)
 */
// llgo:link (*State).Pushnil C.lua_pushnil
func (recv_ *State) Pushnil() {
}

// llgo:link (*State).Pushnumber C.lua_pushnumber
func (recv_ *State) Pushnumber(n Number) {
}

// llgo:link (*State).Pushinteger C.lua_pushinteger
func (recv_ *State) Pushinteger(n Integer) {
}

// llgo:link (*State).Pushlstring C.lua_pushlstring
func (recv_ *State) Pushlstring(s *int8, len uintptr) *int8 {
	return nil
}

// llgo:link (*State).Pushstring C.lua_pushstring
func (recv_ *State) Pushstring(s *int8) *int8 {
	return nil
}

// llgo:link (*State).Pushvfstring C.lua_pushvfstring
func (recv_ *State) Pushvfstring(fmt *int8, argp unsafe.Pointer) *int8 {
	return nil
}

// llgo:link (*State).Pushfstring C.lua_pushfstring
func (recv_ *State) Pushfstring(fmt *int8, __llgo_va_list ...interface{}) *int8 {
	return nil
}

// llgo:link (*State).Pushcclosure C.lua_pushcclosure
func (recv_ *State) Pushcclosure(fn CFunction, n c.Int) {
}

// llgo:link (*State).Pushboolean C.lua_pushboolean
func (recv_ *State) Pushboolean(b c.Int) {
}

// llgo:link (*State).Pushlightuserdata C.lua_pushlightuserdata
func (recv_ *State) Pushlightuserdata(p unsafe.Pointer) {
}

// llgo:link (*State).Pushthread C.lua_pushthread
func (recv_ *State) Pushthread() c.Int {
	return 0
}

/*
** get functions (Lua -> stack)
 */
// llgo:link (*State).Getglobal C.lua_getglobal
func (recv_ *State) Getglobal(name *int8) c.Int {
	return 0
}

// llgo:link (*State).Gettable C.lua_gettable
func (recv_ *State) Gettable(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Getfield C.lua_getfield
func (recv_ *State) Getfield(idx c.Int, k *int8) c.Int {
	return 0
}

// llgo:link (*State).Geti C.lua_geti
func (recv_ *State) Geti(idx c.Int, n Integer) c.Int {
	return 0
}

// llgo:link (*State).Rawget C.lua_rawget
func (recv_ *State) Rawget(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Rawgeti C.lua_rawgeti
func (recv_ *State) Rawgeti(idx c.Int, n Integer) c.Int {
	return 0
}

// llgo:link (*State).Rawgetp C.lua_rawgetp
func (recv_ *State) Rawgetp(idx c.Int, p unsafe.Pointer) c.Int {
	return 0
}

// llgo:link (*State).Createtable C.lua_createtable
func (recv_ *State) Createtable(narr c.Int, nrec c.Int) {
}

// llgo:link (*State).Newuserdatauv C.lua_newuserdatauv
func (recv_ *State) Newuserdatauv(sz uintptr, nuvalue c.Int) unsafe.Pointer {
	return nil
}

// llgo:link (*State).Getmetatable C.lua_getmetatable
func (recv_ *State) Getmetatable(objindex c.Int) c.Int {
	return 0
}

// llgo:link (*State).Getiuservalue C.lua_getiuservalue
func (recv_ *State) Getiuservalue(idx c.Int, n c.Int) c.Int {
	return 0
}

/*
** set functions (stack -> Lua)
 */
// llgo:link (*State).Setglobal C.lua_setglobal
func (recv_ *State) Setglobal(name *int8) {
}

// llgo:link (*State).Settable C.lua_settable
func (recv_ *State) Settable(idx c.Int) {
}

// llgo:link (*State).Setfield C.lua_setfield
func (recv_ *State) Setfield(idx c.Int, k *int8) {
}

// llgo:link (*State).Seti C.lua_seti
func (recv_ *State) Seti(idx c.Int, n Integer) {
}

// llgo:link (*State).Rawset C.lua_rawset
func (recv_ *State) Rawset(idx c.Int) {
}

// llgo:link (*State).Rawseti C.lua_rawseti
func (recv_ *State) Rawseti(idx c.Int, n Integer) {
}

// llgo:link (*State).Rawsetp C.lua_rawsetp
func (recv_ *State) Rawsetp(idx c.Int, p unsafe.Pointer) {
}

// llgo:link (*State).Setmetatable C.lua_setmetatable
func (recv_ *State) Setmetatable(objindex c.Int) c.Int {
	return 0
}

// llgo:link (*State).Setiuservalue C.lua_setiuservalue
func (recv_ *State) Setiuservalue(idx c.Int, n c.Int) c.Int {
	return 0
}

/*
** 'load' and 'call' functions (load and run Lua code)
 */
// llgo:link (*State).Callk C.lua_callk
func (recv_ *State) Callk(nargs c.Int, nresults c.Int, ctx KContext, k KFunction) {
}

// llgo:link (*State).Pcallk C.lua_pcallk
func (recv_ *State) Pcallk(nargs c.Int, nresults c.Int, errfunc c.Int, ctx KContext, k KFunction) c.Int {
	return 0
}

// llgo:link (*State).Load C.lua_load
func (recv_ *State) Load(reader Reader, dt unsafe.Pointer, chunkname *int8, mode *int8) c.Int {
	return 0
}

// llgo:link (*State).Dump C.lua_dump
func (recv_ *State) Dump(writer Writer, data unsafe.Pointer, strip c.Int) c.Int {
	return 0
}

/*
** coroutine functions
 */
// llgo:link (*State).Yieldk C.lua_yieldk
func (recv_ *State) Yieldk(nresults c.Int, ctx KContext, k KFunction) c.Int {
	return 0
}

// llgo:link (*State).Resume C.lua_resume
func (recv_ *State) Resume(from *State, narg c.Int, nres *c.Int) c.Int {
	return 0
}

// llgo:link (*State).Status C.lua_status
func (recv_ *State) Status() c.Int {
	return 0
}

// llgo:link (*State).Isyieldable C.lua_isyieldable
func (recv_ *State) Isyieldable() c.Int {
	return 0
}

/*
** Warning-related functions
 */
// llgo:link (*State).Setwarnf C.lua_setwarnf
func (recv_ *State) Setwarnf(f WarnFunction, ud unsafe.Pointer) {
}

// llgo:link (*State).Warning C.lua_warning
func (recv_ *State) Warning(msg *int8, tocont c.Int) {
}

// llgo:link (*State).Gc C.lua_gc
func (recv_ *State) Gc(what c.Int, __llgo_va_list ...interface{}) c.Int {
	return 0
}

/*
** miscellaneous functions
 */
// llgo:link (*State).Error C.lua_error
func (recv_ *State) Error() c.Int {
	return 0
}

// llgo:link (*State).Next C.lua_next
func (recv_ *State) Next(idx c.Int) c.Int {
	return 0
}

// llgo:link (*State).Concat C.lua_concat
func (recv_ *State) Concat(n c.Int) {
}

// llgo:link (*State).Len C.lua_len
func (recv_ *State) Len(idx c.Int) {
}

// llgo:link (*State).Stringtonumber C.lua_stringtonumber
func (recv_ *State) Stringtonumber(s *int8) uintptr {
	return 0
}

// llgo:link (*State).Getallocf C.lua_getallocf
func (recv_ *State) Getallocf(ud *unsafe.Pointer) Alloc {
	return nil
}

// llgo:link (*State).Setallocf C.lua_setallocf
func (recv_ *State) Setallocf(f Alloc, ud unsafe.Pointer) {
}

// llgo:link (*State).Toclose C.lua_toclose
func (recv_ *State) Toclose(idx c.Int) {
}

// llgo:link (*State).Closeslot C.lua_closeslot
func (recv_ *State) Closeslot(idx c.Int) {
}

// llgo:link (*State).Getstack C.lua_getstack
func (recv_ *State) Getstack(level c.Int, ar *Debug) c.Int {
	return 0
}

// llgo:link (*State).Getinfo C.lua_getinfo
func (recv_ *State) Getinfo(what *int8, ar *Debug) c.Int {
	return 0
}

// llgo:link (*State).Getlocal C.lua_getlocal
func (recv_ *State) Getlocal(ar *Debug, n c.Int) *int8 {
	return nil
}

// llgo:link (*State).Setlocal C.lua_setlocal
func (recv_ *State) Setlocal(ar *Debug, n c.Int) *int8 {
	return nil
}

// llgo:link (*State).Getupvalue C.lua_getupvalue
func (recv_ *State) Getupvalue(funcindex c.Int, n c.Int) *int8 {
	return nil
}

// llgo:link (*State).Setupvalue C.lua_setupvalue
func (recv_ *State) Setupvalue(funcindex c.Int, n c.Int) *int8 {
	return nil
}

// llgo:link (*State).Upvalueid C.lua_upvalueid
func (recv_ *State) Upvalueid(fidx c.Int, n c.Int) unsafe.Pointer {
	return nil
}

// llgo:link (*State).Upvaluejoin C.lua_upvaluejoin
func (recv_ *State) Upvaluejoin(fidx1 c.Int, n1 c.Int, fidx2 c.Int, n2 c.Int) {
}

// llgo:link (*State).Sethook C.lua_sethook
func (recv_ *State) Sethook(func_ Hook, mask c.Int, count c.Int) {
}

// llgo:link (*State).Gethook C.lua_gethook
func (recv_ *State) Gethook() Hook {
	return nil
}

// llgo:link (*State).Gethookmask C.lua_gethookmask
func (recv_ *State) Gethookmask() c.Int {
	return 0
}

// llgo:link (*State).Gethookcount C.lua_gethookcount
func (recv_ *State) Gethookcount() c.Int {
	return 0
}

// llgo:link (*State).Setcstacklimit C.lua_setcstacklimit
func (recv_ *State) Setcstacklimit(limit c.Uint) c.Int {
	return 0
}

type CallInfo struct {
	Unused [8]uint8
}
