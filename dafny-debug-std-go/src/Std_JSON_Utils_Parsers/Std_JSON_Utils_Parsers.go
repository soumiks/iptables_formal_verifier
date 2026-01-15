// Package Std_JSON_Utils_Parsers
// Dafny module Std_JSON_Utils_Parsers compiled into Go

package Std_JSON_Utils_Parsers

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
	m_Std_Actions "Std_Actions"
	m_Std_ActionsExterns "Std_ActionsExterns"
	m_Std_Arithmetic_DivInternals "Std_Arithmetic_DivInternals"
	m_Std_Arithmetic_DivInternalsNonlinear "Std_Arithmetic_DivInternalsNonlinear"
	m_Std_Arithmetic_DivMod "Std_Arithmetic_DivMod"
	m_Std_Arithmetic_GeneralInternals "Std_Arithmetic_GeneralInternals"
	m_Std_Arithmetic_Logarithm "Std_Arithmetic_Logarithm"
	m_Std_Arithmetic_ModInternals "Std_Arithmetic_ModInternals"
	m_Std_Arithmetic_ModInternalsNonlinear "Std_Arithmetic_ModInternalsNonlinear"
	m_Std_Arithmetic_Mul "Std_Arithmetic_Mul"
	m_Std_Arithmetic_MulInternals "Std_Arithmetic_MulInternals"
	m_Std_Arithmetic_MulInternalsNonlinear "Std_Arithmetic_MulInternalsNonlinear"
	m_Std_Arithmetic_Power "Std_Arithmetic_Power"
	m_Std_Arithmetic_Power2 "Std_Arithmetic_Power2"
	m_Std_Base64 "Std_Base64"
	m_Std_BoundedInts "Std_BoundedInts"
	m_Std_BulkActions "Std_BulkActions"
	m_Std_Collections_Array "Std_Collections_Array"
	m_Std_Collections_Imap "Std_Collections_Imap"
	m_Std_Collections_Iset "Std_Collections_Iset"
	m_Std_Collections_Map "Std_Collections_Map"
	m_Std_Collections_Multiset "Std_Collections_Multiset"
	m_Std_Collections_Seq "Std_Collections_Seq"
	m_Std_Collections_Set "Std_Collections_Set"
	m_Std_Collections_Tuple "Std_Collections_Tuple"
	m_Std_Concurrent "Std_Concurrent"
	m_Std_Consumers "Std_Consumers"
	m_Std_DynamicArray "Std_DynamicArray"
	m_Std_FileIO "Std_FileIO"
	m_Std_FileIOInternalExterns "Std_FileIOInternalExterns"
	m_Std_Frames "Std_Frames"
	m_Std_Functions "Std_Functions"
	m_Std_GenericActions "Std_GenericActions"
	m_Std_JSON_Errors "Std_JSON_Errors"
	m_Std_JSON_Spec "Std_JSON_Spec"
	m_Std_JSON_Utils_Cursors "Std_JSON_Utils_Cursors"
	m_Std_JSON_Utils_Lexers_Core "Std_JSON_Utils_Lexers_Core"
	m_Std_JSON_Utils_Lexers_Strings "Std_JSON_Utils_Lexers_Strings"
	m_Std_JSON_Utils_Views_Core "Std_JSON_Utils_Views_Core"
	m_Std_JSON_Utils_Views_Writers "Std_JSON_Utils_Views_Writers"
	m_Std_JSON_Values "Std_JSON_Values"
	m_Std_Math "Std_Math"
	m_Std_Ordinal "Std_Ordinal"
	m_Std_Producers "Std_Producers"
	m_Std_Relations "Std_Relations"
	m_Std_Strings "Std_Strings"
	m_Std_Strings_CharStrEscaping "Std_Strings_CharStrEscaping"
	m_Std_Strings_DecimalConversion "Std_Strings_DecimalConversion"
	m_Std_Strings_HexConversion "Std_Strings_HexConversion"
	m_Std_Termination "Std_Termination"
	m_Std_Unicode_Base "Std_Unicode_Base"
	m_Std_Unicode_UnicodeStringsWithUnicodeChar "Std_Unicode_UnicodeStringsWithUnicodeChar"
	m_Std_Unicode_Utf16EncodingForm "Std_Unicode_Utf16EncodingForm"
	m_Std_Unicode_Utf8EncodingForm "Std_Unicode_Utf8EncodingForm"
	m_Std_Unicode_Utf8EncodingScheme "Std_Unicode_Utf8EncodingScheme"
	m_Std_Wrappers "Std_Wrappers"
	m__System "System_"
	_dafny "dafny"
	os "os"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_IptablesToSmt.Dummy__
var _ m_Analysis.Dummy__
var _ m_Program.Dummy__
var _ m_Std_Wrappers.Dummy__
var _ m_Std_Relations.Dummy__
var _ m_Std_Math.Dummy__
var _ m_Std_Collections_Seq.Dummy__
var _ m_Std_Collections_Array.Dummy__
var _ m_Std_Collections_Imap.Dummy__
var _ m_Std_Functions.Dummy__
var _ m_Std_Collections_Iset.Dummy__
var _ m_Std_Collections_Map.Dummy__
var _ m_Std_Collections_Multiset.Dummy__
var _ m_Std_Collections_Set.Dummy__
var _ m_Std_Collections_Tuple.Dummy__
var _ m_Std_Ordinal.Dummy__
var _ m_Std_Termination.Dummy__
var _ m_Std_Frames.Dummy__
var _ m_Std_GenericActions.Dummy__
var _ m_Std_BoundedInts.Dummy__
var _ m_Std_DynamicArray.Dummy__
var _ m_Std_Actions.Dummy__
var _ m_Std_Consumers.Dummy__
var _ m_Std_Producers.Dummy__
var _ m_Std_ActionsExterns.Dummy__
var _ m_Std_ActionsExterns.Dummy__
var _ m_Std_Concurrent.Dummy__
var _ m_Std_Concurrent.Dummy__
var _ m_Std_FileIOInternalExterns.Dummy__
var _ m_Std_FileIOInternalExterns.Dummy__
var _ m_Std_Unicode_Base.Dummy__
var _ m_Std_Arithmetic_GeneralInternals.Dummy__
var _ m_Std_Arithmetic_MulInternalsNonlinear.Dummy__
var _ m_Std_Arithmetic_MulInternals.Dummy__
var _ m_Std_Arithmetic_Mul.Dummy__
var _ m_Std_Arithmetic_ModInternalsNonlinear.Dummy__
var _ m_Std_Arithmetic_DivInternalsNonlinear.Dummy__
var _ m_Std_Arithmetic_ModInternals.Dummy__
var _ m_Std_Arithmetic_DivInternals.Dummy__
var _ m_Std_Arithmetic_DivMod.Dummy__
var _ m_Std_Arithmetic_Power.Dummy__
var _ m_Std_Arithmetic_Logarithm.Dummy__
var _ m_Std_Arithmetic_Power2.Dummy__
var _ m_Std_Strings_HexConversion.Dummy__
var _ m_Std_Strings_DecimalConversion.Dummy__
var _ m_Std_Strings_CharStrEscaping.Dummy__
var _ m_Std_Strings.Dummy__
var _ m_Std_Unicode_Utf8EncodingForm.Dummy__
var _ m_Std_Unicode_Utf16EncodingForm.Dummy__
var _ m_Std_Unicode_UnicodeStringsWithUnicodeChar.Dummy__
var _ m_Std_Unicode_Utf8EncodingScheme.Dummy__
var _ m_Std_FileIO.Dummy__
var _ m_Std_BulkActions.Dummy__
var _ m_Std_Base64.Dummy__
var _ m_Std_JSON_Values.Dummy__
var _ m_Std_JSON_Errors.Dummy__
var _ m_Std_JSON_Spec.Dummy__
var _ m_Std_JSON_Utils_Views_Core.Dummy__
var _ m_Std_JSON_Utils_Views_Writers.Dummy__
var _ m_Std_JSON_Utils_Lexers_Core.Dummy__
var _ m_Std_JSON_Utils_Lexers_Strings.Dummy__
var _ m_Std_JSON_Utils_Cursors.Dummy__

type Dummy__ struct{}

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "Std_JSON_Utils_Parsers.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ParserWitness() func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	return func(_0___v9 m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
		return m_Std_Wrappers.Companion_Result_.Create_Failure_(m_Std_JSON_Utils_Cursors.Companion_CursorError_.Create_EOF_())
	}
}
func (_static *CompanionStruct_Default___) SubParserWitness() func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	return func(_0_cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
		return m_Std_Wrappers.Companion_Result_.Create_Failure_(m_Std_JSON_Utils_Cursors.Companion_CursorError_.Create_EOF_())
	}
}

// End of class Default__

// Definition of class Parser
type Parser struct {
}

func New_Parser_() *Parser {
	_this := Parser{}

	return &_this
}

type CompanionStruct_Parser_ struct {
}

var Companion_Parser_ = CompanionStruct_Parser_{}

func (*Parser) String() string {
	return "Std_JSON_Utils_Parsers.Parser"
}
func (_this *CompanionStruct_Parser_) Witness() func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	return func(coer39 func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result) func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
		return func(arg40 m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
			return coer39(arg40)
		}
	}(Companion_Default___.ParserWitness())
}

// End of class Parser

func Type_Parser_(Type_T_ _dafny.TypeDescriptor, Type_R_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_Parser_{Type_T_, Type_R_}
}

type type_Parser_ struct {
	Type_T_ _dafny.TypeDescriptor
	Type_R_ _dafny.TypeDescriptor
}

func (_this type_Parser_) Default() interface{} {
	Type_T_ := _this.Type_T_
	_ = Type_T_
	Type_R_ := _this.Type_R_
	_ = Type_R_
	return Companion_Parser_.Witness()
}

func (_this type_Parser_) String() string {
	return "Std_JSON_Utils_Parsers.Parser"
}

// Definition of datatype Parser__
type Parser__ struct {
	Data_Parser___
}

func (_this Parser__) Get_() Data_Parser___ {
	return _this.Data_Parser___
}

type Data_Parser___ interface {
	isParser__()
}

type CompanionStruct_Parser___ struct {
}

var Companion_Parser___ = CompanionStruct_Parser___{}

type Parser___Parser struct {
	Fn func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result
}

func (Parser___Parser) isParser__() {}

func (CompanionStruct_Parser___) Create_Parser_(Fn func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result) Parser__ {
	return Parser__{Parser___Parser{Fn}}
}

func (_this Parser__) Is_Parser() bool {
	_, ok := _this.Get_().(Parser___Parser)
	return ok
}

func (CompanionStruct_Parser___) Default(_default_T interface{}) func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	return func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
		return m_Std_Wrappers.Companion_Result_.Default(m_Std_JSON_Utils_Cursors.Companion_Split_.Default(_default_T))
	}
}

func (_this Parser__) Dtor_fn() func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	return _this.Get_().(Parser___Parser).Fn
}

func (_this Parser__) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Parser___Parser:
		{
			return "Parsers.Parser_.Parser" + "(" + _dafny.String(data.Fn) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Parser__) Equals(other Parser__) bool {
	switch data1 := _this.Get_().(type) {
	case Parser___Parser:
		{
			data2, ok := other.Get_().(Parser___Parser)
			return ok && _dafny.AreEqual(data1.Fn, data2.Fn)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Parser__) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Parser__)
	return ok && _this.Equals(typed)
}

func Type_Parser___(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_Parser___{Type_T_}
}

type type_Parser___ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_Parser___) Default() interface{} {
	Type_T_ := _this.Type_T_
	_ = Type_T_
	return Companion_Parser___.Default(Type_T_.Default())
}

func (_this type_Parser___) String() string {
	return "Std_JSON_Utils_Parsers.Parser__"
}
func (_this Parser__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Parser__{}

// End of datatype Parser__

// Definition of datatype SubParser__
type SubParser__ struct {
	Data_SubParser___
}

func (_this SubParser__) Get_() Data_SubParser___ {
	return _this.Data_SubParser___
}

type Data_SubParser___ interface {
	isSubParser__()
}

type CompanionStruct_SubParser___ struct {
}

var Companion_SubParser___ = CompanionStruct_SubParser___{}

type SubParser___SubParser struct {
	Fn func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result
}

func (SubParser___SubParser) isSubParser__() {}

func (CompanionStruct_SubParser___) Create_SubParser_(Fn func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result) SubParser__ {
	return SubParser__{SubParser___SubParser{Fn}}
}

func (_this SubParser__) Is_SubParser() bool {
	_, ok := _this.Get_().(SubParser___SubParser)
	return ok
}

func (CompanionStruct_SubParser___) Default() func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	return (func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result)(nil)
}

func (_this SubParser__) Dtor_fn() func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	return _this.Get_().(SubParser___SubParser).Fn
}

func (_this SubParser__) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SubParser___SubParser:
		{
			return "Parsers.SubParser_.SubParser" + "(" + _dafny.String(data.Fn) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SubParser__) Equals(other SubParser__) bool {
	switch data1 := _this.Get_().(type) {
	case SubParser___SubParser:
		{
			data2, ok := other.Get_().(SubParser___SubParser)
			return ok && _dafny.AreEqual(data1.Fn, data2.Fn)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SubParser__) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SubParser__)
	return ok && _this.Equals(typed)
}

func Type_SubParser___() _dafny.TypeDescriptor {
	return type_SubParser___{}
}

type type_SubParser___ struct {
}

func (_this type_SubParser___) Default() interface{} {
	return Companion_SubParser___.Default()
}

func (_this type_SubParser___) String() string {
	return "Std_JSON_Utils_Parsers.SubParser__"
}
func (_this SubParser__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SubParser__{}

// End of datatype SubParser__

// Definition of class SubParser
type SubParser struct {
}

func New_SubParser_() *SubParser {
	_this := SubParser{}

	return &_this
}

type CompanionStruct_SubParser_ struct {
}

var Companion_SubParser_ = CompanionStruct_SubParser_{}

func (*SubParser) String() string {
	return "Std_JSON_Utils_Parsers.SubParser"
}
func (_this *CompanionStruct_SubParser_) Witness() func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	return func(coer40 func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result) func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
		return func(arg41 m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
			return coer40(arg41)
		}
	}(Companion_Default___.SubParserWitness())
}

// End of class SubParser

func Type_SubParser_(Type_T_ _dafny.TypeDescriptor, Type_R_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_SubParser_{Type_T_, Type_R_}
}

type type_SubParser_ struct {
	Type_T_ _dafny.TypeDescriptor
	Type_R_ _dafny.TypeDescriptor
}

func (_this type_SubParser_) Default() interface{} {
	Type_T_ := _this.Type_T_
	_ = Type_T_
	Type_R_ := _this.Type_R_
	_ = Type_R_
	return Companion_SubParser_.Witness()
}

func (_this type_SubParser_) String() string {
	return "Std_JSON_Utils_Parsers.SubParser"
}
