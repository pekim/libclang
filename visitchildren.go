package libclang

/*
#include <clang-c/Index.h>

enum CXChildVisitResult visitChildrenCallback(CXCursor cursor, CXCursor parent, CXClientData client_data);

static void visitChildren(CXCursor parent, CXClientData client_data) {
	clang_visitChildren(parent, visitChildrenCallback, client_data);
}

*/
import "C"

import (
	"runtime"
	"sync"
)

type ChildVisitResult C.enum_CXChildVisitResult

const ChildVisit_Break = ChildVisitResult(C.CXChildVisit_Break)
const ChildVisit_Continue = ChildVisitResult(C.CXChildVisit_Continue)
const ChildVisit_Recurse = ChildVisitResult(C.CXChildVisit_Recurse)

type childVisitCallback func(cursor C.CXCursor, parent C.CXCursor) C.enum_CXChildVisitResult
type ChildVisitCallback func(cursor Cursor, parent Cursor) ChildVisitResult

var childVisitCallbackId int
var childVisitCallbackInstanceMap = make(map[int]childVisitCallback)
var childVisitCallbackLock sync.Mutex

func getChildVisitCallbackForId(data C.CXClientData) childVisitCallback {
	childVisitCallbackLock.Lock()
	defer childVisitCallbackLock.Unlock()

	return childVisitCallbackInstanceMap[int(uintptr(data))]
}

func getChildVisitCallbackId(cb childVisitCallback) C.CXClientData {
	childVisitCallbackLock.Lock()
	defer childVisitCallbackLock.Unlock()

	childVisitCallbackId++
	childVisitCallbackInstanceMap[childVisitCallbackId] = cb

	return intToCXClientData(childVisitCallbackId)
}

//export visitChildrenCallback
func visitChildrenCallback(cursor C.CXCursor, parent C.CXCursor, data C.CXClientData) C.enum_CXChildVisitResult {
	var pinner runtime.Pinner
	pinner.Pin(&cursor)
	pinner.Pin(&parent)
	defer pinner.Unpin()

	cb := getChildVisitCallbackForId(data)
	return cb(cursor, parent)
}

func visitChildren(cursor C.CXCursor, cb childVisitCallback) {
	C.visitChildren(cursor, getChildVisitCallbackId(cb))
}

func VisitChildren(cursor Cursor, cb ChildVisitCallback) {
	visitChildren(cursor.c(), func(cursor, parent C.CXCursor) C.enum_CXChildVisitResult {
		var pinner runtime.Pinner
		pinner.Pin(&cursor)
		pinner.Pin(&parent)
		defer pinner.Unpin()

		// Ignore children that are in system headers.
		// May want to make this configurable in the future.
		loc := C.clang_getCursorLocation(cursor)
		inSystemHeader := C.clang_Location_isInSystemHeader(loc) != 0
		if inSystemHeader {
			return C.CXChildVisit_Continue
		}

		res := cb(Cursor(cursor), Cursor(parent))
		return C.enum_CXChildVisitResult(res)
	})
}
