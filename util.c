#include <string.h>

#include "parse.h"

bool hasPrefix(const char *pre, const char *str)
{
    return strncmp(pre, str, strlen(pre)) == 0;
}

char *getCString(CXString string)
{
    char *cString = strdup(clang_getCString(string));
    clang_disposeString(string);

    return cString;
}

char *getCursorSpelling(CXCursor cursor)
{
    return getCString(clang_getCursorSpelling(cursor));
}

bool isCursorStatic(CXCursor cursor)
{
    return clang_Cursor_getStorageClass(cursor) == CX_SC_Static;
}

char *getTypeSpelling(CXType type)
{
    return getCString(clang_getTypeSpelling(type));
}

bool isConstType(CXType type)
{
    return clang_isConstQualifiedType(type) > 0 ||
           hasPrefix("const ", getTypeSpelling(type));
}
