/*
Package libclang provides a means to parse a compilation unit with libclang,
and walk the unit's Cursors.

It's mostly a fairly thin wrapper around some of the C api provided by libclang.
The C api is documented at https://clang.llvm.org/doxygen/group__CINDEX.html.

Note that only a small very small subset of libclang's api is currently supported.
It's enough to parse C and C++ header files, and extract some basic information about
functions, constants, enums, structs, and so on.

# pre-requisites

The clang dev package (with header files) needs to be installed.

	[sudo] dnf install clang-devel

or

	[sudo] apt install clang-dev

# example

A simple example, that just visits a unit's top level cursors.

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
			fmt.Println(cursor.Spelling(), "==>", typ.Spelling())

			return lc.ChildVisit_Continue
		})

		unit.Destroy()
	}
*/
package libclang
