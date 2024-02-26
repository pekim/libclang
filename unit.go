package libclang

// #include <stdlib.h>
// #include <stdio.h>
// #include <string.h>
// #include <clang-c/Index.h>
/*

enum CXChildVisitResult visitCallback(CXCursor cursor, CXCursor parent, CXClientData client_data);

static void visitChildren(CXCursor parent, CXClientData client_data) {
	clang_visitChildren(parent, visitCallback, client_data);
}

*/
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
	C.visitChildren(cursor, intToCXClientData(42))

	// cleanup
	C.clang_disposeIndex(index)
	C.clang_disposeTranslationUnit(unit)

	return nil
}

//export visitCallback
func visitCallback(cursor C.CXCursor, parent C.CXCursor, data C.CXClientData) C.enum_CXChildVisitResult {
	fmt.Println("visit", data, getCursorSpelling(cursor))
	return C.CXChildVisit_Continue
}

func cxString(str C.CXString) string {
	cString := C.strdup(C.clang_getCString(str))
	C.clang_disposeString(str)
	return C.GoString(cString)
}

func getCursorSpelling(cursor C.CXCursor) string {
	return cxString(C.clang_getCursorSpelling(cursor))
}
