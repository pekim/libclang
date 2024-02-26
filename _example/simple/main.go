package main

import (
	"fmt"
	"log"

	lc "github.com/pekim/libclang"
)

func main() {
	unit, err := lc.ParseUnit("/usr/include/clang-c/Index.h")
	if err != nil {
		log.Fatal(err)
	}

	lc.VisitChildren(unit.Cursor(), func(cursor, parent lc.Cursor) lc.ChildVisitResult {
		typ := cursor.Type()

		if typ.Kind() == lc.Type_Enum {
			fmt.Println(cursor.Spelling(), "==>", typ.Spelling())
			lc.VisitChildren(cursor, func(cursor, parent lc.Cursor) lc.ChildVisitResult {
				fmt.Println("  ", cursor.Spelling(), "==>", cursor.EnumConstantDeclValue())
				return lc.ChildVisit_Continue
			})
		}

		// fmt.Println("  ", cursor.IsStatic(), typ.IsConst(), typ.Kind())
		// filepath, line, column := cursor.Location()
		// fmt.Println("  ", filepath, line, column)

		return lc.ChildVisit_Continue
	})

	unit.Destroy()
}
