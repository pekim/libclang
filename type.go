package libclang

// #include <clang-c/Index.h>
import "C"
import "strings"

type Type C.CXType

func (t Type) c() C.CXType {
	return C.CXType(t)
}

func (t Type) Spelling() string {
	return cxString(C.clang_getTypeSpelling(t.c()))
}

func (t Type) IsConst() bool {
	return C.clang_isConstQualifiedType(t.c()) > 0 ||
		strings.HasPrefix(t.Spelling(), "const ")
}
