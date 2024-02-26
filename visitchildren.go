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
	"sync"
)

type childVisitCallback func(cursor C.CXCursor, parent C.CXCursor) C.enum_CXChildVisitResult

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

func visitChildren(cursor C.CXCursor, cb childVisitCallback) {
	C.visitChildren(cursor, getChildVisitCallbackId(cb))
}

//export visitChildrenCallback
func visitChildrenCallback(cursor C.CXCursor, parent C.CXCursor, data C.CXClientData) C.enum_CXChildVisitResult {
	cb := getChildVisitCallbackForId(data)
	return cb(cursor, parent)
}
