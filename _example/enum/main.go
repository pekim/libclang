package main

import (
	"fmt"
	"log"
	"runtime"

	lc "github.com/pekim/libclang"
)

func main() {
	unit, err := lc.ParseUnit("/usr/include/clang-c/Index.h")
	if err != nil {
		log.Fatal(err)
	}

	lc.VisitChildren(unit.Cursor(), func(cursor, parent lc.Cursor) lc.ChildVisitResult {
		var pinner runtime.Pinner

		typ := cursor.Type()
		pinner.Pin(&typ)
		defer pinner.Unpin()

		if typ.Kind() == lc.Type_Enum {
			fmt.Println(cursor.Spelling())
			lc.VisitChildren(cursor, func(cursor, parent lc.Cursor) lc.ChildVisitResult {
				fmt.Println("  ", cursor.Spelling(), "==>", cursor.EnumConstantDeclValue())
				return lc.ChildVisit_Continue
			})
		}

		return lc.ChildVisit_Continue
	})

	unit.Destroy()
}
