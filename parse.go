package libclang

// #include <stdlib.h>
import "C"

import (
	"log"
	"os/exec"
	"path"
	"strings"
	"unsafe"
)

func Parse() {
	parseArgs := []*C.char{
		C.CString("-I"),
		C.CString(path.Join(clangResourceDir(), "include")),
	}
	defer func() {
		for _, cString := range parseArgs {
			C.free(unsafe.Pointer(cString))
		}
	}()

	parseUnit("/usr/include/clang-c/Index.h", parseArgs)
}

func clangResourceDir() string {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	if err != nil {
		log.Fatal(err)
	}

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		log.Fatal("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		log.Fatalf("expected clang resource dir to start with '/', but it %s", resDir)
	}

	return resDir
}
