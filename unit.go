package libclang

// #include <stdlib.h>
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

/*
   CXSourceLocation loc = clang_getCursorLocation(cursor);
   if (clang_Location_isInSystemHeader(loc) != 0)
   {
       return CXChildVisit_Continue;
   }

   switch (clang_getCursorKind(cursor))
   {
   case CXCursor_StructDecl:
   {
       CXType type = clang_getCursorType(cursor);
       char *name = getCursorSpelling(cursor);

       printf("struct : %s\n", name);
   }
   break;
   case CXCursor_FunctionDecl:
   {
       CXType type = clang_getCursorType(cursor);
       char *name = getCursorSpelling(cursor);

       printf("function : %s\n", name);
   }
*/
