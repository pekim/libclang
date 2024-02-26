package libclang

// #include <clang-c/Index.h>
import "C"

type Cursor C.CXCursor

func (c Cursor) c() C.CXCursor {
	return C.CXCursor(c)
}

func (c Cursor) Spelling() string {
	return cxString(C.clang_getCursorSpelling(c.c()))
}

func (c Cursor) Type() Type {
	return Type(C.clang_getCursorType(c.c()))
}
