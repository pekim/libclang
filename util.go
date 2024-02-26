package libclang

// #include <clang-c/Index.h>
// #include <string.h>
import "C"
import "strings"

func cxString(str C.CXString) string {
	cString := C.strdup(C.clang_getCString(str))
	C.clang_disposeString(str)
	return C.GoString(cString)
}

func getCursorSpelling(cursor C.CXCursor) string {
	return cxString(C.clang_getCursorSpelling(cursor))
}

func getTypeSpelling(typ C.CXType) string {
	return cxString(C.clang_getTypeSpelling(typ))
}

func isCursorStatic(cursor C.CXCursor) bool {
	return C.clang_Cursor_getStorageClass(cursor) == C.CX_SC_Static
}

func isConstType(typ C.CXType) bool {
	return C.clang_isConstQualifiedType(typ) > 0 ||
		strings.HasPrefix(getTypeSpelling(typ), "const ")
}
