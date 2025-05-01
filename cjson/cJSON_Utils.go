package cjson

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
// llgo:link (*CJSON).GetPointer C.cJSONUtils_GetPointer
func (p *CJSON) GetPointer(pointer *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).GetPointerCaseSensitive C.cJSONUtils_GetPointerCaseSensitive
func (p *CJSON) GetPointerCaseSensitive(pointer *int8) *CJSON {
	return nil
}
// llgo:link (*CJSON).GeneratePatches C.cJSONUtils_GeneratePatches
func (p *CJSON) GeneratePatches(to *CJSON) *CJSON {
	return nil
}
// llgo:link (*CJSON).GeneratePatchesCaseSensitive C.cJSONUtils_GeneratePatchesCaseSensitive
func (p *CJSON) GeneratePatchesCaseSensitive(to *CJSON) *CJSON {
	return nil
}
// llgo:link (*CJSON).AddPatchToArray C.cJSONUtils_AddPatchToArray
func (p *CJSON) AddPatchToArray(operation *int8, path *int8, value *CJSON) {
}
// llgo:link (*CJSON).ApplyPatches C.cJSONUtils_ApplyPatches
func (p *CJSON) ApplyPatches(patches *CJSON) c.Int {
	return 0
}
// llgo:link (*CJSON).ApplyPatchesCaseSensitive C.cJSONUtils_ApplyPatchesCaseSensitive
func (p *CJSON) ApplyPatchesCaseSensitive(patches *CJSON) c.Int {
	return 0
}
// llgo:link (*CJSON).MergePatch C.cJSONUtils_MergePatch
func (p *CJSON) MergePatch(patch *CJSON) *CJSON {
	return nil
}
// llgo:link (*CJSON).MergePatchCaseSensitive C.cJSONUtils_MergePatchCaseSensitive
func (p *CJSON) MergePatchCaseSensitive(patch *CJSON) *CJSON {
	return nil
}
// llgo:link (*CJSON).GenerateMergePatch C.cJSONUtils_GenerateMergePatch
func (p *CJSON) GenerateMergePatch(to *CJSON) *CJSON {
	return nil
}
// llgo:link (*CJSON).GenerateMergePatchCaseSensitive C.cJSONUtils_GenerateMergePatchCaseSensitive
func (p *CJSON) GenerateMergePatchCaseSensitive(to *CJSON) *CJSON {
	return nil
}
// llgo:link (*CJSON).FindPointerFromObjectTo C.cJSONUtils_FindPointerFromObjectTo
func (p *CJSON) FindPointerFromObjectTo(target *CJSON) *int8 {
	return nil
}
// llgo:link (*CJSON).SortObject C.cJSONUtils_SortObject
func (p *CJSON) SortObject() {
}
// llgo:link (*CJSON).SortObjectCaseSensitive C.cJSONUtils_SortObjectCaseSensitive
func (p *CJSON) SortObjectCaseSensitive() {
}
