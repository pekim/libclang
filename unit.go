package libclang

// #include <stdlib.h>
// #include <clang-c/Index.h>
// #include "parse.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func parseUnit(sourceFilepath string, parseArgs []*C.char) error {
	cFilepath := C.CString(sourceFilepath)
	defer C.free(unsafe.Pointer(cFilepath))

	index := C.clang_createIndex(0, 1)
	unit := (*C.Unit)(C.calloc(1, C.sizeof_Unit))

	parseRes := C.clang_parseTranslationUnit2(
		index, cFilepath, &parseArgs[0], C.int(len(parseArgs)), nil, 0,
		C.CXTranslationUnit_SkipFunctionBodies, &unit.unit,
	)
	if parseRes != C.CXError_Success {
		return fmt.Errorf("non-zero error %d from clang_parseTranslationUnit2 for %s\n",
			parseRes, sourceFilepath)

	}

	// CXCursor cursor = clang_getTranslationUnitCursor(unit->unit);
	// clang_visitChildren(cursor, visitUnitCursors, unit);

	// cleanup
	C.clang_disposeIndex(index)
	C.clang_disposeTranslationUnit(unit.unit)
	C.free(unsafe.Pointer(unit))

	return nil
}
