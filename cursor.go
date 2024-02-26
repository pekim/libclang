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

func (c Cursor) IsStatic() bool {
	return C.clang_Cursor_getStorageClass(c.c()) == C.CX_SC_Static
}

func (c Cursor) Location() (string, int, int) {
	loc := C.clang_getCursorLocation(c.c())
	var file C.CXFile
	var line C.unsigned
	var column C.unsigned
	C.clang_getSpellingLocation(loc, &file, &line, &column, nil)
	filename := C.clang_getFileName(file)
	return cxString(filename), int(line), int(column)
}

func (c Cursor) EnumConstantDeclValue() int64 {
	return int64(C.clang_getEnumConstantDeclValue(c.c()))
}

func (c Cursor) EnumConstantDeclUnsignedValue() uint64 {
	return uint64(C.clang_getEnumConstantDeclUnsignedValue(c.c()))
}

func (c Cursor) Availability() AvailabilityKind {
	return AvailabilityKind(C.clang_getCursorAvailability(c.c()))
}

func (c Cursor) IsDeprecated() bool {
	return c.Availability() == Availability_Deprecated
}
