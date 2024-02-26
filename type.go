package libclang

// #include <clang-c/Index.h>
import "C"

import (
	"strings"
)

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

func (t Type) CanonicalType() Type {
	return Type(C.clang_getCanonicalType(t.c()))
}

func (t Type) PointeeType() Type {
	return Type(C.clang_getPointeeType(t.c()))
}

func (t Type) Kind() TypeKind {
	return TypeKind(t.kind)
}

func (t Type) IsFunctionTypeVariadic() bool {
	return C.clang_isFunctionTypeVariadic(t.c()) == 1
}

func (t Type) ResultType() Type {
	return Type(C.clang_getResultType(t.c()))
}
