package libclang

// #include <clang-c/Index.h>
import "C"

type TypeKind C.enum_CXTypeKind

func (k TypeKind) String() string {
	return typeKindStrings[k]
}

const Type_Invalid = TypeKind(C.CXType_Invalid)
const Type_Unexposed = TypeKind(C.CXType_Unexposed)
const Type_Void = TypeKind(C.CXType_Void)
const Type_Bool = TypeKind(C.CXType_Bool)
const Type_Char_U = TypeKind(C.CXType_Char_U)
const Type_UChar = TypeKind(C.CXType_UChar)
const Type_Char16 = TypeKind(C.CXType_Char16)
const Type_Char32 = TypeKind(C.CXType_Char32)
const Type_UShort = TypeKind(C.CXType_UShort)
const Type_UInt = TypeKind(C.CXType_UInt)
const Type_ULong = TypeKind(C.CXType_ULong)
const Type_ULongLong = TypeKind(C.CXType_ULongLong)
const Type_UInt128 = TypeKind(C.CXType_UInt128)
const Type_Char_S = TypeKind(C.CXType_Char_S)
const Type_SChar = TypeKind(C.CXType_SChar)
const Type_WChar = TypeKind(C.CXType_WChar)
const Type_Short = TypeKind(C.CXType_Short)
const Type_Int = TypeKind(C.CXType_Int)
const Type_Long = TypeKind(C.CXType_Long)
const Type_LongLong = TypeKind(C.CXType_LongLong)
const Type_Int128 = TypeKind(C.CXType_Int128)
const Type_Float = TypeKind(C.CXType_Float)
const Type_Double = TypeKind(C.CXType_Double)
const Type_LongDouble = TypeKind(C.CXType_LongDouble)
const Type_NullPtr = TypeKind(C.CXType_NullPtr)
const Type_Overload = TypeKind(C.CXType_Overload)
const Type_Dependent = TypeKind(C.CXType_Dependent)
const Type_ObjCId = TypeKind(C.CXType_ObjCId)
const Type_ObjCClass = TypeKind(C.CXType_ObjCClass)
const Type_ObjCSel = TypeKind(C.CXType_ObjCSel)
const Type_Float128 = TypeKind(C.CXType_Float128)
const Type_Half = TypeKind(C.CXType_Half)
const Type_Float16 = TypeKind(C.CXType_Float16)
const Type_ShortAccum = TypeKind(C.CXType_ShortAccum)
const Type_Accum = TypeKind(C.CXType_Accum)
const Type_LongAccum = TypeKind(C.CXType_LongAccum)
const Type_UShortAccum = TypeKind(C.CXType_UShortAccum)
const Type_UAccum = TypeKind(C.CXType_UAccum)
const Type_ULongAccum = TypeKind(C.CXType_ULongAccum)
const Type_BFloat16 = TypeKind(C.CXType_BFloat16)
const Type_Ibm128 = TypeKind(C.CXType_Ibm128)
const Type_FirstBuiltin = TypeKind(C.CXType_FirstBuiltin)
const Type_LastBuiltin = TypeKind(C.CXType_LastBuiltin)
const Type_Complex = TypeKind(C.CXType_Complex)
const Type_Pointer = TypeKind(C.CXType_Pointer)
const Type_BlockPointer = TypeKind(C.CXType_BlockPointer)
const Type_LValueReference = TypeKind(C.CXType_LValueReference)
const Type_RValueReference = TypeKind(C.CXType_RValueReference)
const Type_Record = TypeKind(C.CXType_Record)
const Type_Enum = TypeKind(C.CXType_Enum)
const Type_Typedef = TypeKind(C.CXType_Typedef)
const Type_ObjCInterface = TypeKind(C.CXType_ObjCInterface)
const Type_ObjCObjectPointer = TypeKind(C.CXType_ObjCObjectPointer)
const Type_FunctionNoProto = TypeKind(C.CXType_FunctionNoProto)
const Type_FunctionProto = TypeKind(C.CXType_FunctionProto)
const Type_ConstantArray = TypeKind(C.CXType_ConstantArray)
const Type_Vector = TypeKind(C.CXType_Vector)
const Type_IncompleteArray = TypeKind(C.CXType_IncompleteArray)
const Type_VariableArray = TypeKind(C.CXType_VariableArray)
const Type_DependentSizedArray = TypeKind(C.CXType_DependentSizedArray)
const Type_MemberPointer = TypeKind(C.CXType_MemberPointer)
const Type_Auto = TypeKind(C.CXType_Auto)
const Type_Elaborated = TypeKind(C.CXType_Elaborated)
const Type_Pipe = TypeKind(C.CXType_Pipe)
const Type_OCLImage1dRO = TypeKind(C.CXType_OCLImage1dRO)
const Type_OCLImage1dArrayRO = TypeKind(C.CXType_OCLImage1dArrayRO)
const Type_OCLImage1dBufferRO = TypeKind(C.CXType_OCLImage1dBufferRO)
const Type_OCLImage2dRO = TypeKind(C.CXType_OCLImage2dRO)
const Type_OCLImage2dArrayRO = TypeKind(C.CXType_OCLImage2dArrayRO)
const Type_OCLImage2dDepthRO = TypeKind(C.CXType_OCLImage2dDepthRO)
const Type_OCLImage2dArrayDepthRO = TypeKind(C.CXType_OCLImage2dArrayDepthRO)
const Type_OCLImage2dMSAARO = TypeKind(C.CXType_OCLImage2dMSAARO)
const Type_OCLImage2dArrayMSAARO = TypeKind(C.CXType_OCLImage2dArrayMSAARO)
const Type_OCLImage2dMSAADepthRO = TypeKind(C.CXType_OCLImage2dMSAADepthRO)
const Type_OCLImage2dArrayMSAADepthRO = TypeKind(C.CXType_OCLImage2dArrayMSAADepthRO)
const Type_OCLImage3dRO = TypeKind(C.CXType_OCLImage3dRO)
const Type_OCLImage1dWO = TypeKind(C.CXType_OCLImage1dWO)
const Type_OCLImage1dArrayWO = TypeKind(C.CXType_OCLImage1dArrayWO)
const Type_OCLImage1dBufferWO = TypeKind(C.CXType_OCLImage1dBufferWO)
const Type_OCLImage2dWO = TypeKind(C.CXType_OCLImage2dWO)
const Type_OCLImage2dArrayWO = TypeKind(C.CXType_OCLImage2dArrayWO)
const Type_OCLImage2dDepthWO = TypeKind(C.CXType_OCLImage2dDepthWO)
const Type_OCLImage2dArrayDepthWO = TypeKind(C.CXType_OCLImage2dArrayDepthWO)
const Type_OCLImage2dMSAAWO = TypeKind(C.CXType_OCLImage2dMSAAWO)
const Type_OCLImage2dArrayMSAAWO = TypeKind(C.CXType_OCLImage2dArrayMSAAWO)
const Type_OCLImage2dMSAADepthWO = TypeKind(C.CXType_OCLImage2dMSAADepthWO)
const Type_OCLImage2dArrayMSAADepthWO = TypeKind(C.CXType_OCLImage2dArrayMSAADepthWO)
const Type_OCLImage3dWO = TypeKind(C.CXType_OCLImage3dWO)
const Type_OCLImage1dRW = TypeKind(C.CXType_OCLImage1dRW)
const Type_OCLImage1dArrayRW = TypeKind(C.CXType_OCLImage1dArrayRW)
const Type_OCLImage1dBufferRW = TypeKind(C.CXType_OCLImage1dBufferRW)
const Type_OCLImage2dRW = TypeKind(C.CXType_OCLImage2dRW)
const Type_OCLImage2dArrayRW = TypeKind(C.CXType_OCLImage2dArrayRW)
const Type_OCLImage2dDepthRW = TypeKind(C.CXType_OCLImage2dDepthRW)
const Type_OCLImage2dArrayDepthRW = TypeKind(C.CXType_OCLImage2dArrayDepthRW)
const Type_OCLImage2dMSAARW = TypeKind(C.CXType_OCLImage2dMSAARW)
const Type_OCLImage2dArrayMSAARW = TypeKind(C.CXType_OCLImage2dArrayMSAARW)
const Type_OCLImage2dMSAADepthRW = TypeKind(C.CXType_OCLImage2dMSAADepthRW)
const Type_OCLImage2dArrayMSAADepthRW = TypeKind(C.CXType_OCLImage2dArrayMSAADepthRW)
const Type_OCLImage3dRW = TypeKind(C.CXType_OCLImage3dRW)
const Type_OCLSampler = TypeKind(C.CXType_OCLSampler)
const Type_OCLEvent = TypeKind(C.CXType_OCLEvent)
const Type_OCLQueue = TypeKind(C.CXType_OCLQueue)
const Type_OCLReserveID = TypeKind(C.CXType_OCLReserveID)
const Type_ObjCObject = TypeKind(C.CXType_ObjCObject)
const Type_ObjCTypeParam = TypeKind(C.CXType_ObjCTypeParam)
const Type_Attributed = TypeKind(C.CXType_Attributed)
const Type_OCLIntelSubgroupAVCMcePayload = TypeKind(C.CXType_OCLIntelSubgroupAVCMcePayload)
const Type_OCLIntelSubgroupAVCImePayload = TypeKind(C.CXType_OCLIntelSubgroupAVCImePayload)
const Type_OCLIntelSubgroupAVCRefPayload = TypeKind(C.CXType_OCLIntelSubgroupAVCRefPayload)
const Type_OCLIntelSubgroupAVCSicPayload = TypeKind(C.CXType_OCLIntelSubgroupAVCSicPayload)
const Type_OCLIntelSubgroupAVCMceResult = TypeKind(C.CXType_OCLIntelSubgroupAVCMceResult)
const Type_OCLIntelSubgroupAVCImeResult = TypeKind(C.CXType_OCLIntelSubgroupAVCImeResult)
const Type_OCLIntelSubgroupAVCRefResult = TypeKind(C.CXType_OCLIntelSubgroupAVCRefResult)
const Type_OCLIntelSubgroupAVCSicResult = TypeKind(C.CXType_OCLIntelSubgroupAVCSicResult)
const Type_OCLIntelSubgroupAVCImeResultSingleReferenceStreamout = TypeKind(C.CXType_OCLIntelSubgroupAVCImeResultSingleReferenceStreamout)
const Type_OCLIntelSubgroupAVCImeResultDualReferenceStreamout = TypeKind(C.CXType_OCLIntelSubgroupAVCImeResultDualReferenceStreamout)
const Type_OCLIntelSubgroupAVCImeSingleReferenceStreamin = TypeKind(C.CXType_OCLIntelSubgroupAVCImeSingleReferenceStreamin)
const Type_OCLIntelSubgroupAVCImeDualReferenceStreamin = TypeKind(C.CXType_OCLIntelSubgroupAVCImeDualReferenceStreamin)
const Type_OCLIntelSubgroupAVCImeResultSingleRefStreamout = TypeKind(C.CXType_OCLIntelSubgroupAVCImeResultSingleRefStreamout)
const Type_OCLIntelSubgroupAVCImeResultDualRefStreamout = TypeKind(C.CXType_OCLIntelSubgroupAVCImeResultDualRefStreamout)
const Type_OCLIntelSubgroupAVCImeSingleRefStreamin = TypeKind(C.CXType_OCLIntelSubgroupAVCImeSingleRefStreamin)
const Type_OCLIntelSubgroupAVCImeDualRefStreamin = TypeKind(C.CXType_OCLIntelSubgroupAVCImeDualRefStreamin)
const Type_ExtVector = TypeKind(C.CXType_ExtVector)
const Type_Atomic = TypeKind(C.CXType_Atomic)
const Type_BTFTagAttributed = TypeKind(C.CXType_BTFTagAttributed)

var typeKindStrings = map[TypeKind]string{
	Type_Invalid:     "Invalid",
	Type_Unexposed:   "Unexposed",
	Type_Void:        "Void",
	Type_Bool:        "Bool",
	Type_Char_U:      "Char_U",
	Type_UChar:       "UChar",
	Type_Char16:      "Char16",
	Type_Char32:      "Char32",
	Type_UShort:      "UShort",
	Type_UInt:        "UInt",
	Type_ULong:       "ULong",
	Type_ULongLong:   "ULongLong",
	Type_UInt128:     "UInt128",
	Type_Char_S:      "Char_S",
	Type_SChar:       "SChar",
	Type_WChar:       "WChar",
	Type_Short:       "Short",
	Type_Int:         "Int",
	Type_Long:        "Long",
	Type_LongLong:    "LongLong",
	Type_Int128:      "Int128",
	Type_Float:       "Float",
	Type_Double:      "Double",
	Type_LongDouble:  "LongDouble",
	Type_NullPtr:     "NullPtr",
	Type_Overload:    "Overload",
	Type_Dependent:   "Dependent",
	Type_ObjCId:      "ObjCId",
	Type_ObjCClass:   "ObjCClass",
	Type_ObjCSel:     "ObjCSel",
	Type_Float128:    "Float128",
	Type_Half:        "Half",
	Type_Float16:     "Float16",
	Type_ShortAccum:  "ShortAccum",
	Type_Accum:       "Accum",
	Type_LongAccum:   "LongAccum",
	Type_UShortAccum: "UShortAccum",
	Type_UAccum:      "UAccum",
	Type_ULongAccum:  "ULongAccum",
	Type_BFloat16:    "BFloat16",
	Type_Ibm128:      "Ibm128",
	// Type_FirstBuiltin:                  "FirstBuiltin",
	// Type_LastBuiltin:                   "LastBuiltin",
	Type_Complex:                       "Complex",
	Type_Pointer:                       "Pointer",
	Type_BlockPointer:                  "BlockPointer",
	Type_LValueReference:               "LValueReference",
	Type_RValueReference:               "RValueReference",
	Type_Record:                        "Record",
	Type_Enum:                          "Enum",
	Type_Typedef:                       "Typedef",
	Type_ObjCInterface:                 "ObjCInterface",
	Type_ObjCObjectPointer:             "ObjCObjectPointer",
	Type_FunctionNoProto:               "FunctionNoProto",
	Type_FunctionProto:                 "FunctionProto",
	Type_ConstantArray:                 "ConstantArray",
	Type_Vector:                        "Vector",
	Type_IncompleteArray:               "IncompleteArray",
	Type_VariableArray:                 "VariableArray",
	Type_DependentSizedArray:           "DependentSizedArray",
	Type_MemberPointer:                 "MemberPointer",
	Type_Auto:                          "Auto",
	Type_Elaborated:                    "Elaborated",
	Type_Pipe:                          "Pipe",
	Type_OCLImage1dRO:                  "OCLImage1dRO",
	Type_OCLImage1dArrayRO:             "OCLImage1dArrayRO",
	Type_OCLImage1dBufferRO:            "OCLImage1dBufferRO",
	Type_OCLImage2dRO:                  "OCLImage2dRO",
	Type_OCLImage2dArrayRO:             "OCLImage2dArrayRO",
	Type_OCLImage2dDepthRO:             "OCLImage2dDepthRO",
	Type_OCLImage2dArrayDepthRO:        "OCLImage2dArrayDepthRO",
	Type_OCLImage2dMSAARO:              "OCLImage2dMSAARO",
	Type_OCLImage2dArrayMSAARO:         "OCLImage2dArrayMSAARO",
	Type_OCLImage2dMSAADepthRO:         "OCLImage2dMSAADepthRO",
	Type_OCLImage2dArrayMSAADepthRO:    "OCLImage2dArrayMSAADepthRO",
	Type_OCLImage3dRO:                  "OCLImage3dRO",
	Type_OCLImage1dWO:                  "OCLImage1dWO",
	Type_OCLImage1dArrayWO:             "OCLImage1dArrayWO",
	Type_OCLImage1dBufferWO:            "OCLImage1dBufferWO",
	Type_OCLImage2dWO:                  "OCLImage2dWO",
	Type_OCLImage2dArrayWO:             "OCLImage2dArrayWO",
	Type_OCLImage2dDepthWO:             "OCLImage2dDepthWO",
	Type_OCLImage2dArrayDepthWO:        "OCLImage2dArrayDepthWO",
	Type_OCLImage2dMSAAWO:              "OCLImage2dMSAAWO",
	Type_OCLImage2dArrayMSAAWO:         "OCLImage2dArrayMSAAWO",
	Type_OCLImage2dMSAADepthWO:         "OCLImage2dMSAADepthWO",
	Type_OCLImage2dArrayMSAADepthWO:    "OCLImage2dArrayMSAADepthWO",
	Type_OCLImage3dWO:                  "OCLImage3dWO",
	Type_OCLImage1dRW:                  "OCLImage1dRW",
	Type_OCLImage1dArrayRW:             "OCLImage1dArrayRW",
	Type_OCLImage1dBufferRW:            "OCLImage1dBufferRW",
	Type_OCLImage2dRW:                  "OCLImage2dRW",
	Type_OCLImage2dArrayRW:             "OCLImage2dArrayRW",
	Type_OCLImage2dDepthRW:             "OCLImage2dDepthRW",
	Type_OCLImage2dArrayDepthRW:        "OCLImage2dArrayDepthRW",
	Type_OCLImage2dMSAARW:              "OCLImage2dMSAARW",
	Type_OCLImage2dArrayMSAARW:         "OCLImage2dArrayMSAARW",
	Type_OCLImage2dMSAADepthRW:         "OCLImage2dMSAADepthRW",
	Type_OCLImage2dArrayMSAADepthRW:    "OCLImage2dArrayMSAADepthRW",
	Type_OCLImage3dRW:                  "OCLImage3dRW",
	Type_OCLSampler:                    "OCLSampler",
	Type_OCLEvent:                      "OCLEvent",
	Type_OCLQueue:                      "OCLQueue",
	Type_OCLReserveID:                  "OCLReserveID",
	Type_ObjCObject:                    "ObjCObject",
	Type_ObjCTypeParam:                 "ObjCTypeParam",
	Type_Attributed:                    "Attributed",
	Type_OCLIntelSubgroupAVCMcePayload: "OCLIntelSubgroupAVCMcePayload",
	Type_OCLIntelSubgroupAVCImePayload: "OCLIntelSubgroupAVCImePayload",
	Type_OCLIntelSubgroupAVCRefPayload: "OCLIntelSubgroupAVCRefPayload",
	Type_OCLIntelSubgroupAVCSicPayload: "OCLIntelSubgroupAVCSicPayload",
	Type_OCLIntelSubgroupAVCMceResult:  "OCLIntelSubgroupAVCMceResult",
	Type_OCLIntelSubgroupAVCImeResult:  "OCLIntelSubgroupAVCImeResult",
	Type_OCLIntelSubgroupAVCRefResult:  "OCLIntelSubgroupAVCRefResult",
	Type_OCLIntelSubgroupAVCSicResult:  "OCLIntelSubgroupAVCSicResult",
	Type_OCLIntelSubgroupAVCImeResultSingleReferenceStreamout: "OCLIntelSubgroupAVCImeResultSingleReferenceStreamout",
	Type_OCLIntelSubgroupAVCImeResultDualReferenceStreamout:   "OCLIntelSubgroupAVCImeResultDualReferenceStreamout",
	Type_OCLIntelSubgroupAVCImeSingleReferenceStreamin:        "OCLIntelSubgroupAVCImeSingleReferenceStreamin",
	Type_OCLIntelSubgroupAVCImeDualReferenceStreamin:          "OCLIntelSubgroupAVCImeDualReferenceStreamin",
	// Type_OCLIntelSubgroupAVCImeResultSingleRefStreamout:       "OCLIntelSubgroupAVCImeResultSingleRefStreamout",
	// Type_OCLIntelSubgroupAVCImeResultDualRefStreamout:         "OCLIntelSubgroupAVCImeResultDualRefStreamout",
	// Type_OCLIntelSubgroupAVCImeSingleRefStreamin:              "OCLIntelSubgroupAVCImeSingleRefStreamin",
	// Type_OCLIntelSubgroupAVCImeDualRefStreamin:                "OCLIntelSubgroupAVCImeDualRefStreamin",
	Type_ExtVector:        "ExtVector",
	Type_Atomic:           "Atomic",
	Type_BTFTagAttributed: "BTFTagAttributed",
}
