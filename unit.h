#ifndef UNIT_H
#define UNIT_H

#include "parse.h"
#include <clang-c/Index.h>

typedef struct
{
    CXTranslationUnit unit;
} Unit;

Unit *parseUnit(char *srcFilepath, const char *const *args, int num_args);
void freeUnit(Unit *unit);

#endif
