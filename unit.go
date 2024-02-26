package libclang

// #include <stdlib.h>
// #include <clang-c/Index.h>
import "C"

import (
	"fmt"
	"path"
	"unsafe"
)

type Unit struct {
	c C.CXTranslationUnit
}

func ParseUnit(sourceFilepath string) (Unit, error) {
	var unit Unit

	cFilepath := C.CString(sourceFilepath)
	defer C.free(unsafe.Pointer(cFilepath))

	resourceDir, err := clangResourceDir()
	if err != nil {
		return unit, err
	}

	parseArgs := []*C.char{
		C.CString("-I"),
		C.CString(path.Join(resourceDir, "include")),
	}
	defer func() {
		for _, cString := range parseArgs {
			C.free(unsafe.Pointer(cString))
		}
	}()

	index := C.clang_createIndex(0, 1)
	defer C.clang_disposeIndex(index)

	parseRes := C.clang_parseTranslationUnit2(
		index, cFilepath, &parseArgs[0], C.int(len(parseArgs)), nil, 0,
		C.CXTranslationUnit_SkipFunctionBodies, &unit.c,
	)
	if parseRes != C.CXError_Success {
		return unit, fmt.Errorf("non-zero error %d from clang_parseTranslationUnit2 for %s\n",
			parseRes, sourceFilepath)

	}

	return unit, nil
}

func (u Unit) Cursor() Cursor {
	return Cursor(C.clang_getTranslationUnitCursor(u.c))
}

func (u Unit) Destroy() {
	C.clang_disposeTranslationUnit(u.c)
}
