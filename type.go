package libclang

// #include <clang-c/Index.h>
import "C"

type Type C.CXType

func (t Type) c() C.CXType {
	return C.CXType(t)
}

func (t Type) Spelling() string {
	return cxString(C.clang_getTypeSpelling(t.c()))
}
