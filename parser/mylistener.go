package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)

type MyListener struct {
	*BaseCListener // embedded
}

// constructor
func NewMyListener() *MyListener {
	return &MyListener{
		BaseCListener: new(BaseCListener),
	}
}

// override

// VisitTerminal is called when a terminal node is visited.
func (s *MyListener) VisitTerminal(node antlr.TerminalNode) {
	log.Print("test") // "test" を visitterminal という風に関数名にするとどこで呼ばれているかタイミングがわかります
}

// VisitErrorNode is called when an error node is visited.
func (s *MyListener) VisitErrorNode(node antlr.ErrorNode) {
	log.Print("test")
}

// EnterEveryRule is called when any rule is entered.
func (s *MyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	log.Print("test")
}

// ExitEveryRule is called when any rule is exited.
func (s *MyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	log.Print("test")
}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *MyListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {
	log.Print("test")
}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *MyListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {
	log.Print("test")
}

// EnterGenericSelection is called when production genericSelection is entered.
func (s *MyListener) EnterGenericSelection(ctx *GenericSelectionContext) {
	log.Print("test")
}

// ExitGenericSelection is called when production genericSelection is exited.
func (s *MyListener) ExitGenericSelection(ctx *GenericSelectionContext) {
	log.Print("test")
}

// EnterGenericAssocList is called when production genericAssocList is entered.
func (s *MyListener) EnterGenericAssocList(ctx *GenericAssocListContext) {
	log.Print("test")
}

// ExitGenericAssocList is called when production genericAssocList is exited.
func (s *MyListener) ExitGenericAssocList(ctx *GenericAssocListContext) {
	log.Print("test")
}

// EnterGenericAssociation is called when production genericAssociation is entered.
func (s *MyListener) EnterGenericAssociation(ctx *GenericAssociationContext) {
	log.Print("test")
}

// ExitGenericAssociation is called when production genericAssociation is exited.
func (s *MyListener) ExitGenericAssociation(ctx *GenericAssociationContext) {
	log.Print("test")
}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *MyListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {
	log.Print("test")
}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *MyListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {
	log.Print("test")
}

// EnterArgumentExpressionList is called when production argumentExpressionList is entered.
func (s *MyListener) EnterArgumentExpressionList(ctx *ArgumentExpressionListContext) {
	log.Print("test")
}

// ExitArgumentExpressionList is called when production argumentExpressionList is exited.
func (s *MyListener) ExitArgumentExpressionList(ctx *ArgumentExpressionListContext) {
	log.Print("test")
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *MyListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {
	log.Print("test")
}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *MyListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {
	log.Print("test")
}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *MyListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {
	log.Print("test")
}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *MyListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {
	log.Print("test")
}

// EnterCastExpression is called when production castExpression is entered.
func (s *MyListener) EnterCastExpression(ctx *CastExpressionContext) {
	log.Print("test")
}

// ExitCastExpression is called when production castExpression is exited.
func (s *MyListener) ExitCastExpression(ctx *CastExpressionContext) {
	log.Print("test")
}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *MyListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
	log.Print("test")
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *MyListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
	log.Print("test")
}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *MyListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {
	log.Print("test")
}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *MyListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {
	log.Print("test")
}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *MyListener) EnterShiftExpression(ctx *ShiftExpressionContext) {
	log.Print("test")
}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *MyListener) ExitShiftExpression(ctx *ShiftExpressionContext) {
	log.Print("test")
}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *MyListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {
	log.Print("test")
}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *MyListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {
	log.Print("test")
}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *MyListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {
	log.Print("test")
}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *MyListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {
	log.Print("test")
}

// EnterAndExpression is called when production andExpression is entered.
func (s *MyListener) EnterAndExpression(ctx *AndExpressionContext) {
	log.Print("test")
}

// ExitAndExpression is called when production andExpression is exited.
func (s *MyListener) ExitAndExpression(ctx *AndExpressionContext) {
	log.Print("test")
}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *MyListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {
	log.Print("test")
}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *MyListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {
	log.Print("test")
}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *MyListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {
	log.Print("test")
}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *MyListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {
	log.Print("test")
}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *MyListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {
	log.Print("test")
}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *MyListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {
	log.Print("test")
}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *MyListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {
	log.Print("test")
}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *MyListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {
	log.Print("test")
}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *MyListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {
	log.Print("test")
}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *MyListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {
	log.Print("test")
}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *MyListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {
	log.Print("test")
}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *MyListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {
	log.Print("test")
}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *MyListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {
	log.Print("test")
}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *MyListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {
	log.Print("test")
}

// EnterExpression is called when production expression is entered.
func (s *MyListener) EnterExpression(ctx *ExpressionContext) {
	log.Print("test")
}

// ExitExpression is called when production expression is exited.
func (s *MyListener) ExitExpression(ctx *ExpressionContext) {
	log.Print("test")
}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *MyListener) EnterConstantExpression(ctx *ConstantExpressionContext) {
	log.Print("test")
}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *MyListener) ExitConstantExpression(ctx *ConstantExpressionContext) {
	log.Print("test")
}

// EnterDeclaration is called when production declaration is entered.
func (s *MyListener) EnterDeclaration(ctx *DeclarationContext) {
	log.Print("test")
}

// ExitDeclaration is called when production declaration is exited.
func (s *MyListener) ExitDeclaration(ctx *DeclarationContext) {
	log.Print("test")
}

// EnterDeclarationSpecifiers is called when production declarationSpecifiers is entered.
func (s *MyListener) EnterDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {
	log.Print("test")
}

// ExitDeclarationSpecifiers is called when production declarationSpecifiers is exited.
func (s *MyListener) ExitDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {
	log.Print("test")
}

// EnterDeclarationSpecifiers2 is called when production declarationSpecifiers2 is entered.
func (s *MyListener) EnterDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) {
	log.Print("test")
}

// ExitDeclarationSpecifiers2 is called when production declarationSpecifiers2 is exited.
func (s *MyListener) ExitDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) {
	log.Print("test")
}

// EnterDeclarationSpecifier is called when production declarationSpecifier is entered.
func (s *MyListener) EnterDeclarationSpecifier(ctx *DeclarationSpecifierContext) {
	log.Print("test")
}

// ExitDeclarationSpecifier is called when production declarationSpecifier is exited.
func (s *MyListener) ExitDeclarationSpecifier(ctx *DeclarationSpecifierContext) {
	log.Print("test")
}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *MyListener) EnterInitDeclaratorList(ctx *InitDeclaratorListContext) {
	log.Print("test")
}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *MyListener) ExitInitDeclaratorList(ctx *InitDeclaratorListContext) {
	log.Print("test")
}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *MyListener) EnterInitDeclarator(ctx *InitDeclaratorContext) {
	log.Print("test")
}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *MyListener) ExitInitDeclarator(ctx *InitDeclaratorContext) {
	log.Print("test")
}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *MyListener) EnterStorageClassSpecifier(ctx *StorageClassSpecifierContext) {
	log.Print("test")
}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *MyListener) ExitStorageClassSpecifier(ctx *StorageClassSpecifierContext) {
	log.Print("test")
}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *MyListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {
	log.Print("test")
}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *MyListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {
	log.Print("test")
}

// EnterStructOrUnionSpecifier is called when production structOrUnionSpecifier is entered.
func (s *MyListener) EnterStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {
	log.Print("test")
}

// ExitStructOrUnionSpecifier is called when production structOrUnionSpecifier is exited.
func (s *MyListener) ExitStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {
	log.Print("test")
}

// EnterStructOrUnion is called when production structOrUnion is entered.
func (s *MyListener) EnterStructOrUnion(ctx *StructOrUnionContext) {
	log.Print("test")
}

// ExitStructOrUnion is called when production structOrUnion is exited.
func (s *MyListener) ExitStructOrUnion(ctx *StructOrUnionContext) {
	log.Print("test")
}

// EnterStructDeclarationList is called when production structDeclarationList is entered.
func (s *MyListener) EnterStructDeclarationList(ctx *StructDeclarationListContext) {
	log.Print("test")
}

// ExitStructDeclarationList is called when production structDeclarationList is exited.
func (s *MyListener) ExitStructDeclarationList(ctx *StructDeclarationListContext) {
	log.Print("test")
}

// EnterStructDeclaration is called when production structDeclaration is entered.
func (s *MyListener) EnterStructDeclaration(ctx *StructDeclarationContext) {
	log.Print("test")
}

// ExitStructDeclaration is called when production structDeclaration is exited.
func (s *MyListener) ExitStructDeclaration(ctx *StructDeclarationContext) {
	log.Print("test")
}

// EnterSpecifierQualifierList is called when production specifierQualifierList is entered.
func (s *MyListener) EnterSpecifierQualifierList(ctx *SpecifierQualifierListContext) {
	log.Print("test")
}

// ExitSpecifierQualifierList is called when production specifierQualifierList is exited.
func (s *MyListener) ExitSpecifierQualifierList(ctx *SpecifierQualifierListContext) {
	log.Print("test")
}

// EnterStructDeclaratorList is called when production structDeclaratorList is entered.
func (s *MyListener) EnterStructDeclaratorList(ctx *StructDeclaratorListContext) {
	log.Print("test")
}

// ExitStructDeclaratorList is called when production structDeclaratorList is exited.
func (s *MyListener) ExitStructDeclaratorList(ctx *StructDeclaratorListContext) {
	log.Print("test")
}

// EnterStructDeclarator is called when production structDeclarator is entered.
func (s *MyListener) EnterStructDeclarator(ctx *StructDeclaratorContext) {
	log.Print("test")
}

// ExitStructDeclarator is called when production structDeclarator is exited.
func (s *MyListener) ExitStructDeclarator(ctx *StructDeclaratorContext) {
	log.Print("test")
}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *MyListener) EnterEnumSpecifier(ctx *EnumSpecifierContext) {
	log.Print("test")
}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *MyListener) ExitEnumSpecifier(ctx *EnumSpecifierContext) {
	log.Print("test")
}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *MyListener) EnterEnumeratorList(ctx *EnumeratorListContext) {
	log.Print("test")
}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *MyListener) ExitEnumeratorList(ctx *EnumeratorListContext) {
	log.Print("test")
}

// EnterEnumerator is called when production enumerator is entered.
func (s *MyListener) EnterEnumerator(ctx *EnumeratorContext) {
	log.Print("test")
}

// ExitEnumerator is called when production enumerator is exited.
func (s *MyListener) ExitEnumerator(ctx *EnumeratorContext) {
	log.Print("test")
}

// EnterEnumerationConstant is called when production enumerationConstant is entered.
func (s *MyListener) EnterEnumerationConstant(ctx *EnumerationConstantContext) {
	log.Print("test")
}

// ExitEnumerationConstant is called when production enumerationConstant is exited.
func (s *MyListener) ExitEnumerationConstant(ctx *EnumerationConstantContext) {
	log.Print("test")
}

// EnterAtomicTypeSpecifier is called when production atomicTypeSpecifier is entered.
func (s *MyListener) EnterAtomicTypeSpecifier(ctx *AtomicTypeSpecifierContext) {
	log.Print("test")
}

// ExitAtomicTypeSpecifier is called when production atomicTypeSpecifier is exited.
func (s *MyListener) ExitAtomicTypeSpecifier(ctx *AtomicTypeSpecifierContext) {
	log.Print("test")
}

// EnterTypeQualifier is called when production typeQualifier is entered.
func (s *MyListener) EnterTypeQualifier(ctx *TypeQualifierContext) {
	log.Print("test")
}

// ExitTypeQualifier is called when production typeQualifier is exited.
func (s *MyListener) ExitTypeQualifier(ctx *TypeQualifierContext) {
	log.Print("test")
}

// EnterFunctionSpecifier is called when production functionSpecifier is entered.
func (s *MyListener) EnterFunctionSpecifier(ctx *FunctionSpecifierContext) {
	log.Print("test")
}

// ExitFunctionSpecifier is called when production functionSpecifier is exited.
func (s *MyListener) ExitFunctionSpecifier(ctx *FunctionSpecifierContext) {
	log.Print("test")
}

// EnterAlignmentSpecifier is called when production alignmentSpecifier is entered.
func (s *MyListener) EnterAlignmentSpecifier(ctx *AlignmentSpecifierContext) {
	log.Print("test")
}

// ExitAlignmentSpecifier is called when production alignmentSpecifier is exited.
func (s *MyListener) ExitAlignmentSpecifier(ctx *AlignmentSpecifierContext) {
	log.Print("test")
}

// EnterDeclarator is called when production declarator is entered.
func (s *MyListener) EnterDeclarator(ctx *DeclaratorContext) {
	log.Print("test")
}

// ExitDeclarator is called when production declarator is exited.
func (s *MyListener) ExitDeclarator(ctx *DeclaratorContext) {
	log.Print("test")
}

// EnterDirectDeclarator is called when production directDeclarator is entered.
func (s *MyListener) EnterDirectDeclarator(ctx *DirectDeclaratorContext) {
	log.Print("test")
}

// ExitDirectDeclarator is called when production directDeclarator is exited.
func (s *MyListener) ExitDirectDeclarator(ctx *DirectDeclaratorContext) {
	log.Print("test")
}

// EnterGccDeclaratorExtension is called when production gccDeclaratorExtension is entered.
func (s *MyListener) EnterGccDeclaratorExtension(ctx *GccDeclaratorExtensionContext) {
	log.Print("test")
}

// ExitGccDeclaratorExtension is called when production gccDeclaratorExtension is exited.
func (s *MyListener) ExitGccDeclaratorExtension(ctx *GccDeclaratorExtensionContext) {
	log.Print("test")
}

// EnterGccAttributeSpecifier is called when production gccAttributeSpecifier is entered.
func (s *MyListener) EnterGccAttributeSpecifier(ctx *GccAttributeSpecifierContext) {
	log.Print("test")
}

// ExitGccAttributeSpecifier is called when production gccAttributeSpecifier is exited.
func (s *MyListener) ExitGccAttributeSpecifier(ctx *GccAttributeSpecifierContext) {
	log.Print("test")
}

// EnterGccAttributeList is called when production gccAttributeList is entered.
func (s *MyListener) EnterGccAttributeList(ctx *GccAttributeListContext) {
	log.Print("test")
}

// ExitGccAttributeList is called when production gccAttributeList is exited.
func (s *MyListener) ExitGccAttributeList(ctx *GccAttributeListContext) {
	log.Print("test")
}

// EnterGccAttribute is called when production gccAttribute is entered.
func (s *MyListener) EnterGccAttribute(ctx *GccAttributeContext) {
	log.Print("test")
}

// ExitGccAttribute is called when production gccAttribute is exited.
func (s *MyListener) ExitGccAttribute(ctx *GccAttributeContext) {
	log.Print("test")
}

// EnterNestedParenthesesBlock is called when production nestedParenthesesBlock is entered.
func (s *MyListener) EnterNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) {
	log.Print("test")
}

// ExitNestedParenthesesBlock is called when production nestedParenthesesBlock is exited.
func (s *MyListener) ExitNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) {
	log.Print("test")
}

// EnterPointer is called when production pointer is entered.
func (s *MyListener) EnterPointer(ctx *PointerContext) {
	log.Print("test")
}

// ExitPointer is called when production pointer is exited.
func (s *MyListener) ExitPointer(ctx *PointerContext) {
	log.Print("test")
}

// EnterTypeQualifierList is called when production typeQualifierList is entered.
func (s *MyListener) EnterTypeQualifierList(ctx *TypeQualifierListContext) {
	log.Print("test")
}

// ExitTypeQualifierList is called when production typeQualifierList is exited.
func (s *MyListener) ExitTypeQualifierList(ctx *TypeQualifierListContext) {
	log.Print("test")
}

// EnterParameterTypeList is called when production parameterTypeList is entered.
func (s *MyListener) EnterParameterTypeList(ctx *ParameterTypeListContext) {
	log.Print("test")
}

// ExitParameterTypeList is called when production parameterTypeList is exited.
func (s *MyListener) ExitParameterTypeList(ctx *ParameterTypeListContext) {
	log.Print("test")
}

// EnterParameterList is called when production parameterList is entered.
func (s *MyListener) EnterParameterList(ctx *ParameterListContext) {
	log.Print("test")
}

// ExitParameterList is called when production parameterList is exited.
func (s *MyListener) ExitParameterList(ctx *ParameterListContext) {
	log.Print("test")
}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *MyListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {
	log.Print("test")
}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *MyListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {
	log.Print("test")
}

// EnterIdentifierList is called when production identifierList is entered.
func (s *MyListener) EnterIdentifierList(ctx *IdentifierListContext) {
	log.Print("test")
}

// ExitIdentifierList is called when production identifierList is exited.
func (s *MyListener) ExitIdentifierList(ctx *IdentifierListContext) {
	log.Print("test")
}

// EnterTypeName is called when production typeName is entered.
func (s *MyListener) EnterTypeName(ctx *TypeNameContext) {
	log.Print("test")
}

// ExitTypeName is called when production typeName is exited.
func (s *MyListener) ExitTypeName(ctx *TypeNameContext) {
	log.Print("test")
}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *MyListener) EnterAbstractDeclarator(ctx *AbstractDeclaratorContext) {
	log.Print("test")
}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *MyListener) ExitAbstractDeclarator(ctx *AbstractDeclaratorContext) {
	log.Print("test")
}

// EnterDirectAbstractDeclarator is called when production directAbstractDeclarator is entered.
func (s *MyListener) EnterDirectAbstractDeclarator(ctx *DirectAbstractDeclaratorContext) {
	log.Print("test")
}

// ExitDirectAbstractDeclarator is called when production directAbstractDeclarator is exited.
func (s *MyListener) ExitDirectAbstractDeclarator(ctx *DirectAbstractDeclaratorContext) {
	log.Print("test")
}

// EnterTypedefName is called when production typedefName is entered.
func (s *MyListener) EnterTypedefName(ctx *TypedefNameContext) {
	log.Print("test")
}

// ExitTypedefName is called when production typedefName is exited.
func (s *MyListener) ExitTypedefName(ctx *TypedefNameContext) {
	log.Print("test")
}

// EnterInitializer is called when production initializer is entered.
func (s *MyListener) EnterInitializer(ctx *InitializerContext) {
	log.Print("test")
}

// ExitInitializer is called when production initializer is exited.
func (s *MyListener) ExitInitializer(ctx *InitializerContext) {
	log.Print("test")
}

// EnterInitializerList is called when production initializerList is entered.
func (s *MyListener) EnterInitializerList(ctx *InitializerListContext) {
	log.Print("test")
}

// ExitInitializerList is called when production initializerList is exited.
func (s *MyListener) ExitInitializerList(ctx *InitializerListContext) {
	log.Print("test")
}

// EnterDesignation is called when production designation is entered.
func (s *MyListener) EnterDesignation(ctx *DesignationContext) {
	log.Print("test")
}

// ExitDesignation is called when production designation is exited.
func (s *MyListener) ExitDesignation(ctx *DesignationContext) {
	log.Print("test")
}

// EnterDesignatorList is called when production designatorList is entered.
func (s *MyListener) EnterDesignatorList(ctx *DesignatorListContext) {
	log.Print("test")
}

// ExitDesignatorList is called when production designatorList is exited.
func (s *MyListener) ExitDesignatorList(ctx *DesignatorListContext) {
	log.Print("test")
}

// EnterDesignator is called when production designator is entered.
func (s *MyListener) EnterDesignator(ctx *DesignatorContext) {
	log.Print("test")
}

// ExitDesignator is called when production designator is exited.
func (s *MyListener) ExitDesignator(ctx *DesignatorContext) {
	log.Print("test")
}

// EnterStaticAssertDeclaration is called when production staticAssertDeclaration is entered.
func (s *MyListener) EnterStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {
	log.Print("test")
}

// ExitStaticAssertDeclaration is called when production staticAssertDeclaration is exited.
func (s *MyListener) ExitStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {
	log.Print("test")
}

// EnterStatement is called when production statement is entered.
func (s *MyListener) EnterStatement(ctx *StatementContext) {
	log.Print("test")
}

// ExitStatement is called when production statement is exited.
func (s *MyListener) ExitStatement(ctx *StatementContext) {
	log.Print("test")
}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *MyListener) EnterLabeledStatement(ctx *LabeledStatementContext) {
	log.Print("test")
}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *MyListener) ExitLabeledStatement(ctx *LabeledStatementContext) {
	log.Print("test")
}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *MyListener) EnterCompoundStatement(ctx *CompoundStatementContext) {
	log.Print("test")
}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *MyListener) ExitCompoundStatement(ctx *CompoundStatementContext) {
	log.Print("test")
}

// EnterBlockItemList is called when production blockItemList is entered.
func (s *MyListener) EnterBlockItemList(ctx *BlockItemListContext) {
	log.Print("test")
}

// ExitBlockItemList is called when production blockItemList is exited.
func (s *MyListener) ExitBlockItemList(ctx *BlockItemListContext) {
	log.Print("test")
}

// EnterBlockItem is called when production blockItem is entered.
func (s *MyListener) EnterBlockItem(ctx *BlockItemContext) {
	log.Print("test")
}

// ExitBlockItem is called when production blockItem is exited.
func (s *MyListener) ExitBlockItem(ctx *BlockItemContext) {
	log.Print("test")
}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *MyListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {
	log.Print("test")
}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *MyListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {
	log.Print("test")
}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *MyListener) EnterSelectionStatement(ctx *SelectionStatementContext) {
	log.Print("test")
}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *MyListener) ExitSelectionStatement(ctx *SelectionStatementContext) {
	log.Print("test")
}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *MyListener) EnterIterationStatement(ctx *IterationStatementContext) {
	log.Print("test")
}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *MyListener) ExitIterationStatement(ctx *IterationStatementContext) {
	log.Print("test")
}

// EnterForCondition is called when production forCondition is entered.
func (s *MyListener) EnterForCondition(ctx *ForConditionContext) {
	log.Print("test")
}

// ExitForCondition is called when production forCondition is exited.
func (s *MyListener) ExitForCondition(ctx *ForConditionContext) {
	log.Print("test")
}

// EnterForDeclaration is called when production forDeclaration is entered.
func (s *MyListener) EnterForDeclaration(ctx *ForDeclarationContext) {
	log.Print("test")
}

// ExitForDeclaration is called when production forDeclaration is exited.
func (s *MyListener) ExitForDeclaration(ctx *ForDeclarationContext) {
	log.Print("test")
}

// EnterForExpression is called when production forExpression is entered.
func (s *MyListener) EnterForExpression(ctx *ForExpressionContext) {
	log.Print("test")
}

// ExitForExpression is called when production forExpression is exited.
func (s *MyListener) ExitForExpression(ctx *ForExpressionContext) {
	log.Print("test")
}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *MyListener) EnterJumpStatement(ctx *JumpStatementContext) {
	log.Print("test")
}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *MyListener) ExitJumpStatement(ctx *JumpStatementContext) {
	log.Print("test")
}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *MyListener) EnterCompilationUnit(ctx *CompilationUnitContext) {
	log.Print("test")
}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *MyListener) ExitCompilationUnit(ctx *CompilationUnitContext) {
	log.Print("test")
}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *MyListener) EnterTranslationUnit(ctx *TranslationUnitContext) {
	log.Print("test")
}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *MyListener) ExitTranslationUnit(ctx *TranslationUnitContext) {
	log.Print("test")
}

// EnterExternalDeclaration is called when production externalDeclaration is entered.
func (s *MyListener) EnterExternalDeclaration(ctx *ExternalDeclarationContext) {
	log.Print("test")
}

// ExitExternalDeclaration is called when production externalDeclaration is exited.
func (s *MyListener) ExitExternalDeclaration(ctx *ExternalDeclarationContext) {
	log.Print("test")
}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *MyListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {
	log.Print("test")
}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *MyListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {
	log.Print("test")
}

// EnterDeclarationList is called when production declarationList is entered.
func (s *MyListener) EnterDeclarationList(ctx *DeclarationListContext) {
	log.Print("test")
}

// ExitDeclarationList is called when production declarationList is exited.
func (s *MyListener) ExitDeclarationList(ctx *DeclarationListContext) {
	log.Print("test")
}
