package libclang

// #include <clang-c/Index.h>
import "C"

import (
	"unsafe"
)

// intToCXClientData converts an int to CXClientData.
// The typical use is when providing an int as the data argument to a clang function.
//
// The conversion is performed  without triggering the unsafeptr vet warning ("possible misuse of unsafe.pointer").
func intToCXClientData(i int) C.CXClientData {
	notReallyAnAddress := uintptr(i)
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&notReallyAnAddress))
	return C.CXClientData(ptr)
}
