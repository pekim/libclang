package libclang

// #include <clang-c/Index.h>
// #include <string.h>
import "C"

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func cxString(str C.CXString) string {
	cString := C.strdup(C.clang_getCString(str))
	C.clang_disposeString(str)
	return C.GoString(cString)
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

func clangResourceDir() (string, error) {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	if err != nil {
		log.Fatal(err)
	}

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		return "", errors.New("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		return "", fmt.Errorf("expected clang resource dir to start with '/', but it %s", resDir)
	}

	return resDir, nil
}
