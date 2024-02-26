package libclang

// #include <stdlib.h>
// #include <clang-c/Index.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func parseUnit(sourceFilepath string, parseArgs []*C.char) error {
	cFilepath := C.CString(sourceFilepath)
	defer C.free(unsafe.Pointer(cFilepath))

	index := C.clang_createIndex(0, 1)
	var unit C.CXTranslationUnit

	parseRes := C.clang_parseTranslationUnit2(
		index, cFilepath, &parseArgs[0], C.int(len(parseArgs)), nil, 0,
		C.CXTranslationUnit_SkipFunctionBodies, &unit,
	)
	if parseRes != C.CXError_Success {
		return fmt.Errorf("non-zero error %d from clang_parseTranslationUnit2 for %s\n",
			parseRes, sourceFilepath)

	}

	cursor := C.clang_getTranslationUnitCursor(unit)
	visitChildren(cursor,
		func(cursor, parent C.CXCursor) C.enum_CXChildVisitResult {
			loc := C.clang_getCursorLocation(cursor)
			inSystemHeader := C.clang_Location_isInSystemHeader(loc) != 0

			if !inSystemHeader {
				fmt.Println(getCursorSpelling(cursor))
			}

			return C.CXChildVisit_Continue
		},
	)

	// cleanup
	C.clang_disposeIndex(index)
	C.clang_disposeTranslationUnit(unit)

	return nil
}
