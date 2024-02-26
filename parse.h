#ifndef PARSE_H
#define PARSE_H

#include "unit.h"
#include <clang-c/Index.h>
#include <stdbool.h>

// utilities
bool hasPrefix(const char *pre, const char *str);
char *getCString(CXString string);
char *getCursorSpelling(CXCursor cursor);
char *getTypeSpelling(CXType type);
bool isCursorStatic(CXCursor cursor);
bool isConstType(CXType type);

#endif
