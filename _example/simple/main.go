package main

import (
	"fmt"
	"log"

	"github.com/pekim/libclang"
	lc "github.com/pekim/libclang"
)

func main() {
	unit, err := lc.ParseUnit("/usr/include/clang-c/Index.h")
	if err != nil {
		log.Fatal(err)
	}

	lc.VisitChildren(unit.Cursor(), func(cursor, parent libclang.Cursor) libclang.ChildVisitResult {
		fmt.Println(cursor.Spelling(), "==>", cursor.Type().Spelling())
		return lc.ChildVisit_Continue
	})

	unit.Destroy()
}
