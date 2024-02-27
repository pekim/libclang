package libclang

// #include <clang-c/Index.h>
import "C"

type CursorKind C.enum_CXCursorKind

func (k CursorKind) String() string {
	return cursorKindStrings[k]
}

const Cursor_UnexposedDecl = CursorKind(C.CXCursor_UnexposedDecl)
const Cursor_StructDecl = CursorKind(C.CXCursor_StructDecl)
const Cursor_UnionDecl = CursorKind(C.CXCursor_UnionDecl)
const Cursor_ClassDecl = CursorKind(C.CXCursor_ClassDecl)
const Cursor_EnumDecl = CursorKind(C.CXCursor_EnumDecl)
const Cursor_FieldDecl = CursorKind(C.CXCursor_FieldDecl)
const Cursor_EnumConstantDecl = CursorKind(C.CXCursor_EnumConstantDecl)
const Cursor_FunctionDecl = CursorKind(C.CXCursor_FunctionDecl)
const Cursor_VarDecl = CursorKind(C.CXCursor_VarDecl)
const Cursor_ParmDecl = CursorKind(C.CXCursor_ParmDecl)
const Cursor_ObjCInterfaceDecl = CursorKind(C.CXCursor_ObjCInterfaceDecl)
const Cursor_ObjCCategoryDecl = CursorKind(C.CXCursor_ObjCCategoryDecl)
const Cursor_ObjCProtocolDecl = CursorKind(C.CXCursor_ObjCProtocolDecl)
const Cursor_ObjCPropertyDecl = CursorKind(C.CXCursor_ObjCPropertyDecl)
const Cursor_ObjCIvarDecl = CursorKind(C.CXCursor_ObjCIvarDecl)
const Cursor_ObjCInstanceMethodDecl = CursorKind(C.CXCursor_ObjCInstanceMethodDecl)
const Cursor_ObjCClassMethodDecl = CursorKind(C.CXCursor_ObjCClassMethodDecl)
const Cursor_ObjCImplementationDecl = CursorKind(C.CXCursor_ObjCImplementationDecl)
const Cursor_ObjCCategoryImplDecl = CursorKind(C.CXCursor_ObjCCategoryImplDecl)
const Cursor_TypedefDecl = CursorKind(C.CXCursor_TypedefDecl)
const Cursor_CXXMethod = CursorKind(C.CXCursor_CXXMethod)
const Cursor_Namespace = CursorKind(C.CXCursor_Namespace)
const Cursor_LinkageSpec = CursorKind(C.CXCursor_LinkageSpec)
const Cursor_Constructor = CursorKind(C.CXCursor_Constructor)
const Cursor_Destructor = CursorKind(C.CXCursor_Destructor)
const Cursor_ConversionFunction = CursorKind(C.CXCursor_ConversionFunction)
const Cursor_TemplateTypeParameter = CursorKind(C.CXCursor_TemplateTypeParameter)
const Cursor_NonTypeTemplateParameter = CursorKind(C.CXCursor_NonTypeTemplateParameter)
const Cursor_TemplateTemplateParameter = CursorKind(C.CXCursor_TemplateTemplateParameter)
const Cursor_FunctionTemplate = CursorKind(C.CXCursor_FunctionTemplate)
const Cursor_ClassTemplate = CursorKind(C.CXCursor_ClassTemplate)
const Cursor_ClassTemplatePartialSpecialization = CursorKind(C.CXCursor_ClassTemplatePartialSpecialization)
const Cursor_NamespaceAlias = CursorKind(C.CXCursor_NamespaceAlias)
const Cursor_UsingDirective = CursorKind(C.CXCursor_UsingDirective)
const Cursor_UsingDeclaration = CursorKind(C.CXCursor_UsingDeclaration)
const Cursor_TypeAliasDecl = CursorKind(C.CXCursor_TypeAliasDecl)
const Cursor_ObjCSynthesizeDecl = CursorKind(C.CXCursor_ObjCSynthesizeDecl)
const Cursor_ObjCDynamicDecl = CursorKind(C.CXCursor_ObjCDynamicDecl)
const Cursor_CXXAccessSpecifier = CursorKind(C.CXCursor_CXXAccessSpecifier)
const Cursor_FirstDecl = CursorKind(C.CXCursor_FirstDecl)
const Cursor_LastDecl = CursorKind(C.CXCursor_LastDecl)
const Cursor_FirstRef = CursorKind(C.CXCursor_FirstRef)
const Cursor_ObjCSuperClassRef = CursorKind(C.CXCursor_ObjCSuperClassRef)
const Cursor_ObjCProtocolRef = CursorKind(C.CXCursor_ObjCProtocolRef)
const Cursor_ObjCClassRef = CursorKind(C.CXCursor_ObjCClassRef)
const Cursor_TypeRef = CursorKind(C.CXCursor_TypeRef)
const Cursor_CXXBaseSpecifier = CursorKind(C.CXCursor_CXXBaseSpecifier)
const Cursor_TemplateRef = CursorKind(C.CXCursor_TemplateRef)
const Cursor_NamespaceRef = CursorKind(C.CXCursor_NamespaceRef)
const Cursor_MemberRef = CursorKind(C.CXCursor_MemberRef)
const Cursor_LabelRef = CursorKind(C.CXCursor_LabelRef)
const Cursor_OverloadedDeclRef = CursorKind(C.CXCursor_OverloadedDeclRef)
const Cursor_VariableRef = CursorKind(C.CXCursor_VariableRef)
const Cursor_LastRef = CursorKind(C.CXCursor_LastRef)
const Cursor_FirstInvalid = CursorKind(C.CXCursor_FirstInvalid)
const Cursor_InvalidFile = CursorKind(C.CXCursor_InvalidFile)
const Cursor_NoDeclFound = CursorKind(C.CXCursor_NoDeclFound)
const Cursor_NotImplemented = CursorKind(C.CXCursor_NotImplemented)
const Cursor_InvalidCode = CursorKind(C.CXCursor_InvalidCode)
const Cursor_LastInvalid = CursorKind(C.CXCursor_LastInvalid)
const Cursor_FirstExpr = CursorKind(C.CXCursor_FirstExpr)
const Cursor_UnexposedExpr = CursorKind(C.CXCursor_UnexposedExpr)
const Cursor_DeclRefExpr = CursorKind(C.CXCursor_DeclRefExpr)
const Cursor_MemberRefExpr = CursorKind(C.CXCursor_MemberRefExpr)
const Cursor_CallExpr = CursorKind(C.CXCursor_CallExpr)
const Cursor_ObjCMessageExpr = CursorKind(C.CXCursor_ObjCMessageExpr)
const Cursor_BlockExpr = CursorKind(C.CXCursor_BlockExpr)
const Cursor_IntegerLiteral = CursorKind(C.CXCursor_IntegerLiteral)
const Cursor_FloatingLiteral = CursorKind(C.CXCursor_FloatingLiteral)
const Cursor_ImaginaryLiteral = CursorKind(C.CXCursor_ImaginaryLiteral)
const Cursor_StringLiteral = CursorKind(C.CXCursor_StringLiteral)
const Cursor_CharacterLiteral = CursorKind(C.CXCursor_CharacterLiteral)
const Cursor_ParenExpr = CursorKind(C.CXCursor_ParenExpr)
const Cursor_UnaryOperator = CursorKind(C.CXCursor_UnaryOperator)
const Cursor_ArraySubscriptExpr = CursorKind(C.CXCursor_ArraySubscriptExpr)
const Cursor_BinaryOperator = CursorKind(C.CXCursor_BinaryOperator)
const Cursor_CompoundAssignOperator = CursorKind(C.CXCursor_CompoundAssignOperator)
const Cursor_ConditionalOperator = CursorKind(C.CXCursor_ConditionalOperator)
const Cursor_CStyleCastExpr = CursorKind(C.CXCursor_CStyleCastExpr)
const Cursor_CompoundLiteralExpr = CursorKind(C.CXCursor_CompoundLiteralExpr)
const Cursor_InitListExpr = CursorKind(C.CXCursor_InitListExpr)
const Cursor_AddrLabelExpr = CursorKind(C.CXCursor_AddrLabelExpr)
const Cursor_StmtExpr = CursorKind(C.CXCursor_StmtExpr)
const Cursor_GenericSelectionExpr = CursorKind(C.CXCursor_GenericSelectionExpr)
const Cursor_GNUNullExpr = CursorKind(C.CXCursor_GNUNullExpr)
const Cursor_CXXStaticCastExpr = CursorKind(C.CXCursor_CXXStaticCastExpr)
const Cursor_CXXDynamicCastExpr = CursorKind(C.CXCursor_CXXDynamicCastExpr)
const Cursor_CXXReinterpretCastExpr = CursorKind(C.CXCursor_CXXReinterpretCastExpr)
const Cursor_CXXConstCastExpr = CursorKind(C.CXCursor_CXXConstCastExpr)
const Cursor_CXXFunctionalCastExpr = CursorKind(C.CXCursor_CXXFunctionalCastExpr)
const Cursor_CXXTypeidExpr = CursorKind(C.CXCursor_CXXTypeidExpr)
const Cursor_CXXBoolLiteralExpr = CursorKind(C.CXCursor_CXXBoolLiteralExpr)
const Cursor_CXXNullPtrLiteralExpr = CursorKind(C.CXCursor_CXXNullPtrLiteralExpr)
const Cursor_CXXThisExpr = CursorKind(C.CXCursor_CXXThisExpr)
const Cursor_CXXThrowExpr = CursorKind(C.CXCursor_CXXThrowExpr)
const Cursor_CXXNewExpr = CursorKind(C.CXCursor_CXXNewExpr)
const Cursor_CXXDeleteExpr = CursorKind(C.CXCursor_CXXDeleteExpr)
const Cursor_UnaryExpr = CursorKind(C.CXCursor_UnaryExpr)
const Cursor_ObjCStringLiteral = CursorKind(C.CXCursor_ObjCStringLiteral)
const Cursor_ObjCEncodeExpr = CursorKind(C.CXCursor_ObjCEncodeExpr)
const Cursor_ObjCSelectorExpr = CursorKind(C.CXCursor_ObjCSelectorExpr)
const Cursor_ObjCProtocolExpr = CursorKind(C.CXCursor_ObjCProtocolExpr)
const Cursor_ObjCBridgedCastExpr = CursorKind(C.CXCursor_ObjCBridgedCastExpr)
const Cursor_PackExpansionExpr = CursorKind(C.CXCursor_PackExpansionExpr)
const Cursor_SizeOfPackExpr = CursorKind(C.CXCursor_SizeOfPackExpr)
const Cursor_LambdaExpr = CursorKind(C.CXCursor_LambdaExpr)
const Cursor_ObjCBoolLiteralExpr = CursorKind(C.CXCursor_ObjCBoolLiteralExpr)
const Cursor_ObjCSelfExpr = CursorKind(C.CXCursor_ObjCSelfExpr)
const Cursor_OMPArraySectionExpr = CursorKind(C.CXCursor_OMPArraySectionExpr)
const Cursor_ObjCAvailabilityCheckExpr = CursorKind(C.CXCursor_ObjCAvailabilityCheckExpr)
const Cursor_FixedPointLiteral = CursorKind(C.CXCursor_FixedPointLiteral)
const Cursor_OMPArrayShapingExpr = CursorKind(C.CXCursor_OMPArrayShapingExpr)
const Cursor_OMPIteratorExpr = CursorKind(C.CXCursor_OMPIteratorExpr)
const Cursor_CXXAddrspaceCastExpr = CursorKind(C.CXCursor_CXXAddrspaceCastExpr)
const Cursor_ConceptSpecializationExpr = CursorKind(C.CXCursor_ConceptSpecializationExpr)
const Cursor_RequiresExpr = CursorKind(C.CXCursor_RequiresExpr)
const Cursor_CXXParenListInitExpr = CursorKind(C.CXCursor_CXXParenListInitExpr)

// const Cursor_PackIndexingExpr = CursorKind(C.CXCursor_PackIndexingExpr)
const Cursor_LastExpr = CursorKind(C.CXCursor_LastExpr)
const Cursor_FirstStmt = CursorKind(C.CXCursor_FirstStmt)
const Cursor_UnexposedStmt = CursorKind(C.CXCursor_UnexposedStmt)
const Cursor_LabelStmt = CursorKind(C.CXCursor_LabelStmt)
const Cursor_CompoundStmt = CursorKind(C.CXCursor_CompoundStmt)
const Cursor_CaseStmt = CursorKind(C.CXCursor_CaseStmt)
const Cursor_DefaultStmt = CursorKind(C.CXCursor_DefaultStmt)
const Cursor_IfStmt = CursorKind(C.CXCursor_IfStmt)
const Cursor_SwitchStmt = CursorKind(C.CXCursor_SwitchStmt)
const Cursor_WhileStmt = CursorKind(C.CXCursor_WhileStmt)
const Cursor_DoStmt = CursorKind(C.CXCursor_DoStmt)
const Cursor_ForStmt = CursorKind(C.CXCursor_ForStmt)
const Cursor_GotoStmt = CursorKind(C.CXCursor_GotoStmt)
const Cursor_IndirectGotoStmt = CursorKind(C.CXCursor_IndirectGotoStmt)
const Cursor_ContinueStmt = CursorKind(C.CXCursor_ContinueStmt)
const Cursor_BreakStmt = CursorKind(C.CXCursor_BreakStmt)
const Cursor_ReturnStmt = CursorKind(C.CXCursor_ReturnStmt)
const Cursor_GCCAsmStmt = CursorKind(C.CXCursor_GCCAsmStmt)
const Cursor_AsmStmt = CursorKind(C.CXCursor_AsmStmt)
const Cursor_ObjCAtTryStmt = CursorKind(C.CXCursor_ObjCAtTryStmt)
const Cursor_ObjCAtCatchStmt = CursorKind(C.CXCursor_ObjCAtCatchStmt)
const Cursor_ObjCAtFinallyStmt = CursorKind(C.CXCursor_ObjCAtFinallyStmt)
const Cursor_ObjCAtThrowStmt = CursorKind(C.CXCursor_ObjCAtThrowStmt)
const Cursor_ObjCAtSynchronizedStmt = CursorKind(C.CXCursor_ObjCAtSynchronizedStmt)
const Cursor_ObjCAutoreleasePoolStmt = CursorKind(C.CXCursor_ObjCAutoreleasePoolStmt)
const Cursor_ObjCForCollectionStmt = CursorKind(C.CXCursor_ObjCForCollectionStmt)
const Cursor_CXXCatchStmt = CursorKind(C.CXCursor_CXXCatchStmt)
const Cursor_CXXTryStmt = CursorKind(C.CXCursor_CXXTryStmt)
const Cursor_CXXForRangeStmt = CursorKind(C.CXCursor_CXXForRangeStmt)
const Cursor_SEHTryStmt = CursorKind(C.CXCursor_SEHTryStmt)
const Cursor_SEHExceptStmt = CursorKind(C.CXCursor_SEHExceptStmt)
const Cursor_SEHFinallyStmt = CursorKind(C.CXCursor_SEHFinallyStmt)
const Cursor_MSAsmStmt = CursorKind(C.CXCursor_MSAsmStmt)
const Cursor_NullStmt = CursorKind(C.CXCursor_NullStmt)
const Cursor_DeclStmt = CursorKind(C.CXCursor_DeclStmt)
const Cursor_OMPParallelDirective = CursorKind(C.CXCursor_OMPParallelDirective)
const Cursor_OMPSimdDirective = CursorKind(C.CXCursor_OMPSimdDirective)
const Cursor_OMPForDirective = CursorKind(C.CXCursor_OMPForDirective)
const Cursor_OMPSectionsDirective = CursorKind(C.CXCursor_OMPSectionsDirective)
const Cursor_OMPSectionDirective = CursorKind(C.CXCursor_OMPSectionDirective)
const Cursor_OMPSingleDirective = CursorKind(C.CXCursor_OMPSingleDirective)
const Cursor_OMPParallelForDirective = CursorKind(C.CXCursor_OMPParallelForDirective)
const Cursor_OMPParallelSectionsDirective = CursorKind(C.CXCursor_OMPParallelSectionsDirective)
const Cursor_OMPTaskDirective = CursorKind(C.CXCursor_OMPTaskDirective)
const Cursor_OMPMasterDirective = CursorKind(C.CXCursor_OMPMasterDirective)
const Cursor_OMPCriticalDirective = CursorKind(C.CXCursor_OMPCriticalDirective)
const Cursor_OMPTaskyieldDirective = CursorKind(C.CXCursor_OMPTaskyieldDirective)
const Cursor_OMPBarrierDirective = CursorKind(C.CXCursor_OMPBarrierDirective)
const Cursor_OMPTaskwaitDirective = CursorKind(C.CXCursor_OMPTaskwaitDirective)
const Cursor_OMPFlushDirective = CursorKind(C.CXCursor_OMPFlushDirective)
const Cursor_SEHLeaveStmt = CursorKind(C.CXCursor_SEHLeaveStmt)
const Cursor_OMPOrderedDirective = CursorKind(C.CXCursor_OMPOrderedDirective)
const Cursor_OMPAtomicDirective = CursorKind(C.CXCursor_OMPAtomicDirective)
const Cursor_OMPForSimdDirective = CursorKind(C.CXCursor_OMPForSimdDirective)
const Cursor_OMPParallelForSimdDirective = CursorKind(C.CXCursor_OMPParallelForSimdDirective)
const Cursor_OMPTargetDirective = CursorKind(C.CXCursor_OMPTargetDirective)
const Cursor_OMPTeamsDirective = CursorKind(C.CXCursor_OMPTeamsDirective)
const Cursor_OMPTaskgroupDirective = CursorKind(C.CXCursor_OMPTaskgroupDirective)
const Cursor_OMPCancellationPointDirective = CursorKind(C.CXCursor_OMPCancellationPointDirective)
const Cursor_OMPCancelDirective = CursorKind(C.CXCursor_OMPCancelDirective)
const Cursor_OMPTargetDataDirective = CursorKind(C.CXCursor_OMPTargetDataDirective)
const Cursor_OMPTaskLoopDirective = CursorKind(C.CXCursor_OMPTaskLoopDirective)
const Cursor_OMPTaskLoopSimdDirective = CursorKind(C.CXCursor_OMPTaskLoopSimdDirective)
const Cursor_OMPDistributeDirective = CursorKind(C.CXCursor_OMPDistributeDirective)
const Cursor_OMPTargetEnterDataDirective = CursorKind(C.CXCursor_OMPTargetEnterDataDirective)
const Cursor_OMPTargetExitDataDirective = CursorKind(C.CXCursor_OMPTargetExitDataDirective)
const Cursor_OMPTargetParallelDirective = CursorKind(C.CXCursor_OMPTargetParallelDirective)
const Cursor_OMPTargetParallelForDirective = CursorKind(C.CXCursor_OMPTargetParallelForDirective)
const Cursor_OMPTargetUpdateDirective = CursorKind(C.CXCursor_OMPTargetUpdateDirective)
const Cursor_OMPDistributeParallelForDirective = CursorKind(C.CXCursor_OMPDistributeParallelForDirective)
const Cursor_OMPDistributeParallelForSimdDirective = CursorKind(C.CXCursor_OMPDistributeParallelForSimdDirective)
const Cursor_OMPDistributeSimdDirective = CursorKind(C.CXCursor_OMPDistributeSimdDirective)
const Cursor_OMPTargetParallelForSimdDirective = CursorKind(C.CXCursor_OMPTargetParallelForSimdDirective)
const Cursor_OMPTargetSimdDirective = CursorKind(C.CXCursor_OMPTargetSimdDirective)
const Cursor_OMPTeamsDistributeDirective = CursorKind(C.CXCursor_OMPTeamsDistributeDirective)
const Cursor_OMPTeamsDistributeSimdDirective = CursorKind(C.CXCursor_OMPTeamsDistributeSimdDirective)
const Cursor_OMPTeamsDistributeParallelForSimdDirective = CursorKind(C.CXCursor_OMPTeamsDistributeParallelForSimdDirective)
const Cursor_OMPTeamsDistributeParallelForDirective = CursorKind(C.CXCursor_OMPTeamsDistributeParallelForDirective)
const Cursor_OMPTargetTeamsDirective = CursorKind(C.CXCursor_OMPTargetTeamsDirective)
const Cursor_OMPTargetTeamsDistributeDirective = CursorKind(C.CXCursor_OMPTargetTeamsDistributeDirective)
const Cursor_OMPTargetTeamsDistributeParallelForDirective = CursorKind(C.CXCursor_OMPTargetTeamsDistributeParallelForDirective)
const Cursor_OMPTargetTeamsDistributeParallelForSimdDirective = CursorKind(C.CXCursor_OMPTargetTeamsDistributeParallelForSimdDirective)
const Cursor_OMPTargetTeamsDistributeSimdDirective = CursorKind(C.CXCursor_OMPTargetTeamsDistributeSimdDirective)
const Cursor_BuiltinBitCastExpr = CursorKind(C.CXCursor_BuiltinBitCastExpr)
const Cursor_OMPMasterTaskLoopDirective = CursorKind(C.CXCursor_OMPMasterTaskLoopDirective)
const Cursor_OMPParallelMasterTaskLoopDirective = CursorKind(C.CXCursor_OMPParallelMasterTaskLoopDirective)
const Cursor_OMPMasterTaskLoopSimdDirective = CursorKind(C.CXCursor_OMPMasterTaskLoopSimdDirective)
const Cursor_OMPParallelMasterTaskLoopSimdDirective = CursorKind(C.CXCursor_OMPParallelMasterTaskLoopSimdDirective)
const Cursor_OMPParallelMasterDirective = CursorKind(C.CXCursor_OMPParallelMasterDirective)
const Cursor_OMPDepobjDirective = CursorKind(C.CXCursor_OMPDepobjDirective)
const Cursor_OMPScanDirective = CursorKind(C.CXCursor_OMPScanDirective)
const Cursor_OMPTileDirective = CursorKind(C.CXCursor_OMPTileDirective)
const Cursor_OMPCanonicalLoop = CursorKind(C.CXCursor_OMPCanonicalLoop)
const Cursor_OMPInteropDirective = CursorKind(C.CXCursor_OMPInteropDirective)
const Cursor_OMPDispatchDirective = CursorKind(C.CXCursor_OMPDispatchDirective)
const Cursor_OMPMaskedDirective = CursorKind(C.CXCursor_OMPMaskedDirective)
const Cursor_OMPUnrollDirective = CursorKind(C.CXCursor_OMPUnrollDirective)
const Cursor_OMPMetaDirective = CursorKind(C.CXCursor_OMPMetaDirective)
const Cursor_OMPGenericLoopDirective = CursorKind(C.CXCursor_OMPGenericLoopDirective)
const Cursor_OMPTeamsGenericLoopDirective = CursorKind(C.CXCursor_OMPTeamsGenericLoopDirective)
const Cursor_OMPTargetTeamsGenericLoopDirective = CursorKind(C.CXCursor_OMPTargetTeamsGenericLoopDirective)
const Cursor_OMPParallelGenericLoopDirective = CursorKind(C.CXCursor_OMPParallelGenericLoopDirective)
const Cursor_OMPTargetParallelGenericLoopDirective = CursorKind(C.CXCursor_OMPTargetParallelGenericLoopDirective)
const Cursor_OMPParallelMaskedDirective = CursorKind(C.CXCursor_OMPParallelMaskedDirective)
const Cursor_OMPMaskedTaskLoopDirective = CursorKind(C.CXCursor_OMPMaskedTaskLoopDirective)
const Cursor_OMPMaskedTaskLoopSimdDirective = CursorKind(C.CXCursor_OMPMaskedTaskLoopSimdDirective)
const Cursor_OMPParallelMaskedTaskLoopDirective = CursorKind(C.CXCursor_OMPParallelMaskedTaskLoopDirective)
const Cursor_OMPParallelMaskedTaskLoopSimdDirective = CursorKind(C.CXCursor_OMPParallelMaskedTaskLoopSimdDirective)
const Cursor_OMPErrorDirective = CursorKind(C.CXCursor_OMPErrorDirective)

// const Cursor_OMPScopeDirective = CursorKind(C.CXCursor_OMPScopeDirective)
// const Cursor_OpenACCComputeConstruct = CursorKind(C.CXCursor_OpenACCComputeConstruct)
const Cursor_LastStmt = CursorKind(C.CXCursor_LastStmt)
const Cursor_TranslationUnit = CursorKind(C.CXCursor_TranslationUnit)
const Cursor_FirstAttr = CursorKind(C.CXCursor_FirstAttr)
const Cursor_UnexposedAttr = CursorKind(C.CXCursor_UnexposedAttr)
const Cursor_IBActionAttr = CursorKind(C.CXCursor_IBActionAttr)
const Cursor_IBOutletAttr = CursorKind(C.CXCursor_IBOutletAttr)
const Cursor_IBOutletCollectionAttr = CursorKind(C.CXCursor_IBOutletCollectionAttr)
const Cursor_CXXFinalAttr = CursorKind(C.CXCursor_CXXFinalAttr)
const Cursor_CXXOverrideAttr = CursorKind(C.CXCursor_CXXOverrideAttr)
const Cursor_AnnotateAttr = CursorKind(C.CXCursor_AnnotateAttr)
const Cursor_AsmLabelAttr = CursorKind(C.CXCursor_AsmLabelAttr)
const Cursor_PackedAttr = CursorKind(C.CXCursor_PackedAttr)
const Cursor_PureAttr = CursorKind(C.CXCursor_PureAttr)
const Cursor_ConstAttr = CursorKind(C.CXCursor_ConstAttr)
const Cursor_NoDuplicateAttr = CursorKind(C.CXCursor_NoDuplicateAttr)
const Cursor_CUDAConstantAttr = CursorKind(C.CXCursor_CUDAConstantAttr)
const Cursor_CUDADeviceAttr = CursorKind(C.CXCursor_CUDADeviceAttr)
const Cursor_CUDAGlobalAttr = CursorKind(C.CXCursor_CUDAGlobalAttr)
const Cursor_CUDAHostAttr = CursorKind(C.CXCursor_CUDAHostAttr)
const Cursor_CUDASharedAttr = CursorKind(C.CXCursor_CUDASharedAttr)
const Cursor_VisibilityAttr = CursorKind(C.CXCursor_VisibilityAttr)
const Cursor_DLLExport = CursorKind(C.CXCursor_DLLExport)
const Cursor_DLLImport = CursorKind(C.CXCursor_DLLImport)
const Cursor_NSReturnsRetained = CursorKind(C.CXCursor_NSReturnsRetained)
const Cursor_NSReturnsNotRetained = CursorKind(C.CXCursor_NSReturnsNotRetained)
const Cursor_NSReturnsAutoreleased = CursorKind(C.CXCursor_NSReturnsAutoreleased)
const Cursor_NSConsumesSelf = CursorKind(C.CXCursor_NSConsumesSelf)
const Cursor_NSConsumed = CursorKind(C.CXCursor_NSConsumed)
const Cursor_ObjCException = CursorKind(C.CXCursor_ObjCException)
const Cursor_ObjCNSObject = CursorKind(C.CXCursor_ObjCNSObject)
const Cursor_ObjCIndependentClass = CursorKind(C.CXCursor_ObjCIndependentClass)
const Cursor_ObjCPreciseLifetime = CursorKind(C.CXCursor_ObjCPreciseLifetime)
const Cursor_ObjCReturnsInnerPointer = CursorKind(C.CXCursor_ObjCReturnsInnerPointer)
const Cursor_ObjCRequiresSuper = CursorKind(C.CXCursor_ObjCRequiresSuper)
const Cursor_ObjCRootClass = CursorKind(C.CXCursor_ObjCRootClass)
const Cursor_ObjCSubclassingRestricted = CursorKind(C.CXCursor_ObjCSubclassingRestricted)
const Cursor_ObjCExplicitProtocolImpl = CursorKind(C.CXCursor_ObjCExplicitProtocolImpl)
const Cursor_ObjCDesignatedInitializer = CursorKind(C.CXCursor_ObjCDesignatedInitializer)
const Cursor_ObjCRuntimeVisible = CursorKind(C.CXCursor_ObjCRuntimeVisible)
const Cursor_ObjCBoxable = CursorKind(C.CXCursor_ObjCBoxable)
const Cursor_FlagEnum = CursorKind(C.CXCursor_FlagEnum)
const Cursor_ConvergentAttr = CursorKind(C.CXCursor_ConvergentAttr)
const Cursor_WarnUnusedAttr = CursorKind(C.CXCursor_WarnUnusedAttr)
const Cursor_WarnUnusedResultAttr = CursorKind(C.CXCursor_WarnUnusedResultAttr)
const Cursor_AlignedAttr = CursorKind(C.CXCursor_AlignedAttr)
const Cursor_LastAttr = CursorKind(C.CXCursor_LastAttr)
const Cursor_PreprocessingDirective = CursorKind(C.CXCursor_PreprocessingDirective)
const Cursor_MacroDefinition = CursorKind(C.CXCursor_MacroDefinition)
const Cursor_MacroExpansion = CursorKind(C.CXCursor_MacroExpansion)
const Cursor_MacroInstantiation = CursorKind(C.CXCursor_MacroInstantiation)
const Cursor_InclusionDirective = CursorKind(C.CXCursor_InclusionDirective)
const Cursor_FirstPreprocessing = CursorKind(C.CXCursor_FirstPreprocessing)
const Cursor_LastPreprocessing = CursorKind(C.CXCursor_LastPreprocessing)
const Cursor_ModuleImportDecl = CursorKind(C.CXCursor_ModuleImportDecl)
const Cursor_TypeAliasTemplateDecl = CursorKind(C.CXCursor_TypeAliasTemplateDecl)
const Cursor_StaticAssert = CursorKind(C.CXCursor_StaticAssert)
const Cursor_FriendDecl = CursorKind(C.CXCursor_FriendDecl)
const Cursor_ConceptDecl = CursorKind(C.CXCursor_ConceptDecl)
const Cursor_FirstExtraDecl = CursorKind(C.CXCursor_FirstExtraDecl)
const Cursor_LastExtraDecl = CursorKind(C.CXCursor_LastExtraDecl)
const Cursor_OverloadCandidate = CursorKind(C.CXCursor_OverloadCandidate)

var cursorKindStrings = map[CursorKind]string{
	Cursor_UnexposedDecl:                      "UnexposedDecl",
	Cursor_StructDecl:                         "StructDecl",
	Cursor_UnionDecl:                          "UnionDecl",
	Cursor_ClassDecl:                          "ClassDecl",
	Cursor_EnumDecl:                           "EnumDecl",
	Cursor_FieldDecl:                          "FieldDecl",
	Cursor_EnumConstantDecl:                   "EnumConstantDecl",
	Cursor_FunctionDecl:                       "FunctionDecl",
	Cursor_VarDecl:                            "VarDecl",
	Cursor_ParmDecl:                           "ParmDecl",
	Cursor_ObjCInterfaceDecl:                  "ObjCInterfaceDecl",
	Cursor_ObjCCategoryDecl:                   "ObjCCategoryDecl",
	Cursor_ObjCProtocolDecl:                   "ObjCProtocolDecl",
	Cursor_ObjCPropertyDecl:                   "ObjCPropertyDecl",
	Cursor_ObjCIvarDecl:                       "ObjCIvarDecl",
	Cursor_ObjCInstanceMethodDecl:             "ObjCInstanceMethodDecl",
	Cursor_ObjCClassMethodDecl:                "ObjCClassMethodDecl",
	Cursor_ObjCImplementationDecl:             "ObjCImplementationDecl",
	Cursor_ObjCCategoryImplDecl:               "ObjCCategoryImplDecl",
	Cursor_TypedefDecl:                        "TypedefDecl",
	Cursor_CXXMethod:                          "CXXMethod",
	Cursor_Namespace:                          "Namespace",
	Cursor_LinkageSpec:                        "LinkageSpec",
	Cursor_Constructor:                        "Constructor",
	Cursor_Destructor:                         "Destructor",
	Cursor_ConversionFunction:                 "ConversionFunction",
	Cursor_TemplateTypeParameter:              "TemplateTypeParameter",
	Cursor_NonTypeTemplateParameter:           "NonTypeTemplateParameter",
	Cursor_TemplateTemplateParameter:          "TemplateTemplateParameter",
	Cursor_FunctionTemplate:                   "FunctionTemplate",
	Cursor_ClassTemplate:                      "ClassTemplate",
	Cursor_ClassTemplatePartialSpecialization: "ClassTemplatePartialSpecialization",
	Cursor_NamespaceAlias:                     "NamespaceAlias",
	Cursor_UsingDirective:                     "UsingDirective",
	Cursor_UsingDeclaration:                   "UsingDeclaration",
	Cursor_TypeAliasDecl:                      "TypeAliasDecl",
	Cursor_ObjCSynthesizeDecl:                 "ObjCSynthesizeDecl",
	Cursor_ObjCDynamicDecl:                    "ObjCDynamicDecl",
	Cursor_CXXAccessSpecifier:                 "CXXAccessSpecifier",
	// Cursor_FirstDecl:                                        "FirstDecl",
	// Cursor_LastDecl:                                         "LastDecl",
	// Cursor_FirstRef:                                         "FirstRef",
	Cursor_ObjCSuperClassRef: "ObjCSuperClassRef",
	Cursor_ObjCProtocolRef:   "ObjCProtocolRef",
	Cursor_ObjCClassRef:      "ObjCClassRef",
	Cursor_TypeRef:           "TypeRef",
	Cursor_CXXBaseSpecifier:  "CXXBaseSpecifier",
	Cursor_TemplateRef:       "TemplateRef",
	Cursor_NamespaceRef:      "NamespaceRef",
	Cursor_MemberRef:         "MemberRef",
	Cursor_LabelRef:          "LabelRef",
	Cursor_OverloadedDeclRef: "OverloadedDeclRef",
	Cursor_VariableRef:       "VariableRef",
	// Cursor_LastRef:                                          "LastRef",
	// Cursor_FirstInvalid:                                     "FirstInvalid",
	Cursor_InvalidFile:    "InvalidFile",
	Cursor_NoDeclFound:    "NoDeclFound",
	Cursor_NotImplemented: "NotImplemented",
	Cursor_InvalidCode:    "InvalidCode",
	// Cursor_LastInvalid:                                      "LastInvalid",
	// Cursor_FirstExpr:                                        "FirstExpr",
	Cursor_UnexposedExpr:             "UnexposedExpr",
	Cursor_DeclRefExpr:               "DeclRefExpr",
	Cursor_MemberRefExpr:             "MemberRefExpr",
	Cursor_CallExpr:                  "CallExpr",
	Cursor_ObjCMessageExpr:           "ObjCMessageExpr",
	Cursor_BlockExpr:                 "BlockExpr",
	Cursor_IntegerLiteral:            "IntegerLiteral",
	Cursor_FloatingLiteral:           "FloatingLiteral",
	Cursor_ImaginaryLiteral:          "ImaginaryLiteral",
	Cursor_StringLiteral:             "StringLiteral",
	Cursor_CharacterLiteral:          "CharacterLiteral",
	Cursor_ParenExpr:                 "ParenExpr",
	Cursor_UnaryOperator:             "UnaryOperator",
	Cursor_ArraySubscriptExpr:        "ArraySubscriptExpr",
	Cursor_BinaryOperator:            "BinaryOperator",
	Cursor_CompoundAssignOperator:    "CompoundAssignOperator",
	Cursor_ConditionalOperator:       "ConditionalOperator",
	Cursor_CStyleCastExpr:            "CStyleCastExpr",
	Cursor_CompoundLiteralExpr:       "CompoundLiteralExpr",
	Cursor_InitListExpr:              "InitListExpr",
	Cursor_AddrLabelExpr:             "AddrLabelExpr",
	Cursor_StmtExpr:                  "StmtExpr",
	Cursor_GenericSelectionExpr:      "GenericSelectionExpr",
	Cursor_GNUNullExpr:               "GNUNullExpr",
	Cursor_CXXStaticCastExpr:         "CXXStaticCastExpr",
	Cursor_CXXDynamicCastExpr:        "CXXDynamicCastExpr",
	Cursor_CXXReinterpretCastExpr:    "CXXReinterpretCastExpr",
	Cursor_CXXConstCastExpr:          "CXXConstCastExpr",
	Cursor_CXXFunctionalCastExpr:     "CXXFunctionalCastExpr",
	Cursor_CXXTypeidExpr:             "CXXTypeidExpr",
	Cursor_CXXBoolLiteralExpr:        "CXXBoolLiteralExpr",
	Cursor_CXXNullPtrLiteralExpr:     "CXXNullPtrLiteralExpr",
	Cursor_CXXThisExpr:               "CXXThisExpr",
	Cursor_CXXThrowExpr:              "CXXThrowExpr",
	Cursor_CXXNewExpr:                "CXXNewExpr",
	Cursor_CXXDeleteExpr:             "CXXDeleteExpr",
	Cursor_UnaryExpr:                 "UnaryExpr",
	Cursor_ObjCStringLiteral:         "ObjCStringLiteral",
	Cursor_ObjCEncodeExpr:            "ObjCEncodeExpr",
	Cursor_ObjCSelectorExpr:          "ObjCSelectorExpr",
	Cursor_ObjCProtocolExpr:          "ObjCProtocolExpr",
	Cursor_ObjCBridgedCastExpr:       "ObjCBridgedCastExpr",
	Cursor_PackExpansionExpr:         "PackExpansionExpr",
	Cursor_SizeOfPackExpr:            "SizeOfPackExpr",
	Cursor_LambdaExpr:                "LambdaExpr",
	Cursor_ObjCBoolLiteralExpr:       "ObjCBoolLiteralExpr",
	Cursor_ObjCSelfExpr:              "ObjCSelfExpr",
	Cursor_OMPArraySectionExpr:       "OMPArraySectionExpr",
	Cursor_ObjCAvailabilityCheckExpr: "ObjCAvailabilityCheckExpr",
	Cursor_FixedPointLiteral:         "FixedPointLiteral",
	Cursor_OMPArrayShapingExpr:       "OMPArrayShapingExpr",
	Cursor_OMPIteratorExpr:           "OMPIteratorExpr",
	Cursor_CXXAddrspaceCastExpr:      "CXXAddrspaceCastExpr",
	Cursor_ConceptSpecializationExpr: "ConceptSpecializationExpr",
	Cursor_RequiresExpr:              "RequiresExpr",
	Cursor_CXXParenListInitExpr:      "CXXParenListInitExpr",
	// Cursor_PackIndexingExpr:          "PackIndexingExpr",
	// Cursor_LastExpr:                                         "LastExpr",
	// Cursor_FirstStmt:                                        "FirstStmt",
	Cursor_UnexposedStmt:    "UnexposedStmt",
	Cursor_LabelStmt:        "LabelStmt",
	Cursor_CompoundStmt:     "CompoundStmt",
	Cursor_CaseStmt:         "CaseStmt",
	Cursor_DefaultStmt:      "DefaultStmt",
	Cursor_IfStmt:           "IfStmt",
	Cursor_SwitchStmt:       "SwitchStmt",
	Cursor_WhileStmt:        "WhileStmt",
	Cursor_DoStmt:           "DoStmt",
	Cursor_ForStmt:          "ForStmt",
	Cursor_GotoStmt:         "GotoStmt",
	Cursor_IndirectGotoStmt: "IndirectGotoStmt",
	Cursor_ContinueStmt:     "ContinueStmt",
	Cursor_BreakStmt:        "BreakStmt",
	Cursor_ReturnStmt:       "ReturnStmt",
	// Cursor_GCCAsmStmt:                                       "GCCAsmStmt",
	Cursor_AsmStmt:                                          "AsmStmt",
	Cursor_ObjCAtTryStmt:                                    "ObjCAtTryStmt",
	Cursor_ObjCAtCatchStmt:                                  "ObjCAtCatchStmt",
	Cursor_ObjCAtFinallyStmt:                                "ObjCAtFinallyStmt",
	Cursor_ObjCAtThrowStmt:                                  "ObjCAtThrowStmt",
	Cursor_ObjCAtSynchronizedStmt:                           "ObjCAtSynchronizedStmt",
	Cursor_ObjCAutoreleasePoolStmt:                          "ObjCAutoreleasePoolStmt",
	Cursor_ObjCForCollectionStmt:                            "ObjCForCollectionStmt",
	Cursor_CXXCatchStmt:                                     "CXXCatchStmt",
	Cursor_CXXTryStmt:                                       "CXXTryStmt",
	Cursor_CXXForRangeStmt:                                  "CXXForRangeStmt",
	Cursor_SEHTryStmt:                                       "SEHTryStmt",
	Cursor_SEHExceptStmt:                                    "SEHExceptStmt",
	Cursor_SEHFinallyStmt:                                   "SEHFinallyStmt",
	Cursor_MSAsmStmt:                                        "MSAsmStmt",
	Cursor_NullStmt:                                         "NullStmt",
	Cursor_DeclStmt:                                         "DeclStmt",
	Cursor_OMPParallelDirective:                             "OMPParallelDirective",
	Cursor_OMPSimdDirective:                                 "OMPSimdDirective",
	Cursor_OMPForDirective:                                  "OMPForDirective",
	Cursor_OMPSectionsDirective:                             "OMPSectionsDirective",
	Cursor_OMPSectionDirective:                              "OMPSectionDirective",
	Cursor_OMPSingleDirective:                               "OMPSingleDirective",
	Cursor_OMPParallelForDirective:                          "OMPParallelForDirective",
	Cursor_OMPParallelSectionsDirective:                     "OMPParallelSectionsDirective",
	Cursor_OMPTaskDirective:                                 "OMPTaskDirective",
	Cursor_OMPMasterDirective:                               "OMPMasterDirective",
	Cursor_OMPCriticalDirective:                             "OMPCriticalDirective",
	Cursor_OMPTaskyieldDirective:                            "OMPTaskyieldDirective",
	Cursor_OMPBarrierDirective:                              "OMPBarrierDirective",
	Cursor_OMPTaskwaitDirective:                             "OMPTaskwaitDirective",
	Cursor_OMPFlushDirective:                                "OMPFlushDirective",
	Cursor_SEHLeaveStmt:                                     "SEHLeaveStmt",
	Cursor_OMPOrderedDirective:                              "OMPOrderedDirective",
	Cursor_OMPAtomicDirective:                               "OMPAtomicDirective",
	Cursor_OMPForSimdDirective:                              "OMPForSimdDirective",
	Cursor_OMPParallelForSimdDirective:                      "OMPParallelForSimdDirective",
	Cursor_OMPTargetDirective:                               "OMPTargetDirective",
	Cursor_OMPTeamsDirective:                                "OMPTeamsDirective",
	Cursor_OMPTaskgroupDirective:                            "OMPTaskgroupDirective",
	Cursor_OMPCancellationPointDirective:                    "OMPCancellationPointDirective",
	Cursor_OMPCancelDirective:                               "OMPCancelDirective",
	Cursor_OMPTargetDataDirective:                           "OMPTargetDataDirective",
	Cursor_OMPTaskLoopDirective:                             "OMPTaskLoopDirective",
	Cursor_OMPTaskLoopSimdDirective:                         "OMPTaskLoopSimdDirective",
	Cursor_OMPDistributeDirective:                           "OMPDistributeDirective",
	Cursor_OMPTargetEnterDataDirective:                      "OMPTargetEnterDataDirective",
	Cursor_OMPTargetExitDataDirective:                       "OMPTargetExitDataDirective",
	Cursor_OMPTargetParallelDirective:                       "OMPTargetParallelDirective",
	Cursor_OMPTargetParallelForDirective:                    "OMPTargetParallelForDirective",
	Cursor_OMPTargetUpdateDirective:                         "OMPTargetUpdateDirective",
	Cursor_OMPDistributeParallelForDirective:                "OMPDistributeParallelForDirective",
	Cursor_OMPDistributeParallelForSimdDirective:            "OMPDistributeParallelForSimdDirective",
	Cursor_OMPDistributeSimdDirective:                       "OMPDistributeSimdDirective",
	Cursor_OMPTargetParallelForSimdDirective:                "OMPTargetParallelForSimdDirective",
	Cursor_OMPTargetSimdDirective:                           "OMPTargetSimdDirective",
	Cursor_OMPTeamsDistributeDirective:                      "OMPTeamsDistributeDirective",
	Cursor_OMPTeamsDistributeSimdDirective:                  "OMPTeamsDistributeSimdDirective",
	Cursor_OMPTeamsDistributeParallelForSimdDirective:       "OMPTeamsDistributeParallelForSimdDirective",
	Cursor_OMPTeamsDistributeParallelForDirective:           "OMPTeamsDistributeParallelForDirective",
	Cursor_OMPTargetTeamsDirective:                          "OMPTargetTeamsDirective",
	Cursor_OMPTargetTeamsDistributeDirective:                "OMPTargetTeamsDistributeDirective",
	Cursor_OMPTargetTeamsDistributeParallelForDirective:     "OMPTargetTeamsDistributeParallelForDirective",
	Cursor_OMPTargetTeamsDistributeParallelForSimdDirective: "OMPTargetTeamsDistributeParallelForSimdDirective",
	Cursor_OMPTargetTeamsDistributeSimdDirective:            "OMPTargetTeamsDistributeSimdDirective",
	Cursor_BuiltinBitCastExpr:                               "BuiltinBitCastExpr",
	Cursor_OMPMasterTaskLoopDirective:                       "OMPMasterTaskLoopDirective",
	Cursor_OMPParallelMasterTaskLoopDirective:               "OMPParallelMasterTaskLoopDirective",
	Cursor_OMPMasterTaskLoopSimdDirective:                   "OMPMasterTaskLoopSimdDirective",
	Cursor_OMPParallelMasterTaskLoopSimdDirective:           "OMPParallelMasterTaskLoopSimdDirective",
	Cursor_OMPParallelMasterDirective:                       "OMPParallelMasterDirective",
	Cursor_OMPDepobjDirective:                               "OMPDepobjDirective",
	Cursor_OMPScanDirective:                                 "OMPScanDirective",
	Cursor_OMPTileDirective:                                 "OMPTileDirective",
	Cursor_OMPCanonicalLoop:                                 "OMPCanonicalLoop",
	Cursor_OMPInteropDirective:                              "OMPInteropDirective",
	Cursor_OMPDispatchDirective:                             "OMPDispatchDirective",
	Cursor_OMPMaskedDirective:                               "OMPMaskedDirective",
	Cursor_OMPUnrollDirective:                               "OMPUnrollDirective",
	Cursor_OMPMetaDirective:                                 "OMPMetaDirective",
	Cursor_OMPGenericLoopDirective:                          "OMPGenericLoopDirective",
	Cursor_OMPTeamsGenericLoopDirective:                     "OMPTeamsGenericLoopDirective",
	Cursor_OMPTargetTeamsGenericLoopDirective:               "OMPTargetTeamsGenericLoopDirective",
	Cursor_OMPParallelGenericLoopDirective:                  "OMPParallelGenericLoopDirective",
	Cursor_OMPTargetParallelGenericLoopDirective:            "OMPTargetParallelGenericLoopDirective",
	Cursor_OMPParallelMaskedDirective:                       "OMPParallelMaskedDirective",
	Cursor_OMPMaskedTaskLoopDirective:                       "OMPMaskedTaskLoopDirective",
	Cursor_OMPMaskedTaskLoopSimdDirective:                   "OMPMaskedTaskLoopSimdDirective",
	Cursor_OMPParallelMaskedTaskLoopDirective:               "OMPParallelMaskedTaskLoopDirective",
	Cursor_OMPParallelMaskedTaskLoopSimdDirective:           "OMPParallelMaskedTaskLoopSimdDirective",
	Cursor_OMPErrorDirective:                                "OMPErrorDirective",
	// Cursor_OMPScopeDirective:                                "OMPScopeDirective",
	// Cursor_OpenACCComputeConstruct:                          "OpenACCComputeConstruct",
	// Cursor_LastStmt:                                         "LastStmt",
	Cursor_TranslationUnit: "TranslationUnit",
	// Cursor_FirstAttr:                                        "FirstAttr",
	Cursor_UnexposedAttr:             "UnexposedAttr",
	Cursor_IBActionAttr:              "IBActionAttr",
	Cursor_IBOutletAttr:              "IBOutletAttr",
	Cursor_IBOutletCollectionAttr:    "IBOutletCollectionAttr",
	Cursor_CXXFinalAttr:              "CXXFinalAttr",
	Cursor_CXXOverrideAttr:           "CXXOverrideAttr",
	Cursor_AnnotateAttr:              "AnnotateAttr",
	Cursor_AsmLabelAttr:              "AsmLabelAttr",
	Cursor_PackedAttr:                "PackedAttr",
	Cursor_PureAttr:                  "PureAttr",
	Cursor_ConstAttr:                 "ConstAttr",
	Cursor_NoDuplicateAttr:           "NoDuplicateAttr",
	Cursor_CUDAConstantAttr:          "CUDAConstantAttr",
	Cursor_CUDADeviceAttr:            "CUDADeviceAttr",
	Cursor_CUDAGlobalAttr:            "CUDAGlobalAttr",
	Cursor_CUDAHostAttr:              "CUDAHostAttr",
	Cursor_CUDASharedAttr:            "CUDASharedAttr",
	Cursor_VisibilityAttr:            "VisibilityAttr",
	Cursor_DLLExport:                 "DLLExport",
	Cursor_DLLImport:                 "DLLImport",
	Cursor_NSReturnsRetained:         "NSReturnsRetained",
	Cursor_NSReturnsNotRetained:      "NSReturnsNotRetained",
	Cursor_NSReturnsAutoreleased:     "NSReturnsAutoreleased",
	Cursor_NSConsumesSelf:            "NSConsumesSelf",
	Cursor_NSConsumed:                "NSConsumed",
	Cursor_ObjCException:             "ObjCException",
	Cursor_ObjCNSObject:              "ObjCNSObject",
	Cursor_ObjCIndependentClass:      "ObjCIndependentClass",
	Cursor_ObjCPreciseLifetime:       "ObjCPreciseLifetime",
	Cursor_ObjCReturnsInnerPointer:   "ObjCReturnsInnerPointer",
	Cursor_ObjCRequiresSuper:         "ObjCRequiresSuper",
	Cursor_ObjCRootClass:             "ObjCRootClass",
	Cursor_ObjCSubclassingRestricted: "ObjCSubclassingRestricted",
	Cursor_ObjCExplicitProtocolImpl:  "ObjCExplicitProtocolImpl",
	Cursor_ObjCDesignatedInitializer: "ObjCDesignatedInitializer",
	Cursor_ObjCRuntimeVisible:        "ObjCRuntimeVisible",
	Cursor_ObjCBoxable:               "ObjCBoxable",
	Cursor_FlagEnum:                  "FlagEnum",
	Cursor_ConvergentAttr:            "ConvergentAttr",
	Cursor_WarnUnusedAttr:            "WarnUnusedAttr",
	Cursor_WarnUnusedResultAttr:      "WarnUnusedResultAttr",
	Cursor_AlignedAttr:               "AlignedAttr",
	// Cursor_LastAttr:                  "LastAttr",
	Cursor_PreprocessingDirective: "PreprocessingDirective",
	Cursor_MacroDefinition:        "MacroDefinition",
	Cursor_MacroExpansion:         "MacroExpansion",
	// Cursor_MacroInstantiation:     "MacroInstantiation",
	Cursor_InclusionDirective: "InclusionDirective",
	// Cursor_FirstPreprocessing:        "FirstPreprocessing",
	// Cursor_LastPreprocessing:         "LastPreprocessing",
	Cursor_ModuleImportDecl:      "ModuleImportDecl",
	Cursor_TypeAliasTemplateDecl: "TypeAliasTemplateDecl",
	Cursor_StaticAssert:          "StaticAssert",
	Cursor_FriendDecl:            "FriendDecl",
	Cursor_ConceptDecl:           "ConceptDecl",
	// Cursor_FirstExtraDecl:            "FirstExtraDecl",
	// Cursor_LastExtraDecl:             "LastExtraDecl",
	Cursor_OverloadCandidate: "OverloadCandidate",
}
