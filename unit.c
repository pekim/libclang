#include <clang-c/Index.h>
#include <stdio.h>
#include <stdlib.h>

#include "parse.h"

enum CXChildVisitResult visitUnitCursors(CXCursor cursor, CXCursor parent,
                                         CXClientData client_data)
{
    Unit *unit = client_data;

    CXSourceLocation loc = clang_getCursorLocation(cursor);
    if (clang_Location_isInSystemHeader(loc) != 0)
    {
        return CXChildVisit_Continue;
    }

    switch (clang_getCursorKind(cursor))
    {
    case CXCursor_StructDecl:
    {
        CXType type = clang_getCursorType(cursor);
        char *name = getCursorSpelling(cursor);

        printf("struct : %s\n", name);
    }
    break;
    case CXCursor_FunctionDecl:
    {
        CXType type = clang_getCursorType(cursor);
        char *name = getCursorSpelling(cursor);

        printf("function : %s\n", name);
    }
    break;
    case CXCursor_EnumDecl:
    {
        printf("enum\n");
    }
    break;
    default:
        break;
    }

    return CXChildVisit_Continue;
}

Unit *parseUnit(char *srcFilepath, const char *const *args, int num_args)
{
    CXIndex index = clang_createIndex(0, 1);
    Unit *unit = calloc(1, sizeof(Unit));

    enum CXErrorCode err = clang_parseTranslationUnit2(
        index, srcFilepath, args, num_args, NULL, 0,
        CXTranslationUnit_SkipFunctionBodies, &unit->unit);

    if (err != 0)
    {
        printf("non-zero error %d from clang_parseTranslationUnit2 for %s\n",
               err, srcFilepath);
        exit(1);
    }

    CXCursor cursor = clang_getTranslationUnitCursor(unit->unit);
    clang_visitChildren(cursor, visitUnitCursors, unit);

    clang_disposeIndex(index);

    return unit;
}

void freeUnit(Unit *unit)
{
    clang_disposeTranslationUnit(unit->unit);

    free(unit);
}
