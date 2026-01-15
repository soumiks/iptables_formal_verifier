// Package Std_JSON_Utils_Lexers_Core
// Dafny module Std_JSON_Utils_Lexers_Core compiled into Go

package Std_JSON_Utils_Lexers_Core

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

type Dummy__ struct{}

// Definition of datatype LexerResult
type LexerResult struct {
	Data_LexerResult_
}

func (_this LexerResult) Get_() Data_LexerResult_ {
	return _this.Data_LexerResult_
}

type Data_LexerResult_ interface {
	isLexerResult()
}

type CompanionStruct_LexerResult_ struct {
}

var Companion_LexerResult_ = CompanionStruct_LexerResult_{}

type LexerResult_Accept struct {
}

func (LexerResult_Accept) isLexerResult() {}

func (CompanionStruct_LexerResult_) Create_Accept_() LexerResult {
	return LexerResult{LexerResult_Accept{}}
}

func (_this LexerResult) Is_Accept() bool {
	_, ok := _this.Get_().(LexerResult_Accept)
	return ok
}

type LexerResult_Reject struct {
	Err interface{}
}

func (LexerResult_Reject) isLexerResult() {}

func (CompanionStruct_LexerResult_) Create_Reject_(Err interface{}) LexerResult {
	return LexerResult{LexerResult_Reject{Err}}
}

func (_this LexerResult) Is_Reject() bool {
	_, ok := _this.Get_().(LexerResult_Reject)
	return ok
}

type LexerResult_Partial struct {
	St interface{}
}

func (LexerResult_Partial) isLexerResult() {}

func (CompanionStruct_LexerResult_) Create_Partial_(St interface{}) LexerResult {
	return LexerResult{LexerResult_Partial{St}}
}

func (_this LexerResult) Is_Partial() bool {
	_, ok := _this.Get_().(LexerResult_Partial)
	return ok
}

func (CompanionStruct_LexerResult_) Default() LexerResult {
	return Companion_LexerResult_.Create_Accept_()
}

func (_this LexerResult) Dtor_err() interface{} {
	return _this.Get_().(LexerResult_Reject).Err
}

func (_this LexerResult) Dtor_st() interface{} {
	return _this.Get_().(LexerResult_Partial).St
}

func (_this LexerResult) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case LexerResult_Accept:
		{
			return "Core.LexerResult.Accept"
		}
	case LexerResult_Reject:
		{
			return "Core.LexerResult.Reject" + "(" + _dafny.String(data.Err) + ")"
		}
	case LexerResult_Partial:
		{
			return "Core.LexerResult.Partial" + "(" + _dafny.String(data.St) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this LexerResult) Equals(other LexerResult) bool {
	switch data1 := _this.Get_().(type) {
	case LexerResult_Accept:
		{
			_, ok := other.Get_().(LexerResult_Accept)
			return ok
		}
	case LexerResult_Reject:
		{
			data2, ok := other.Get_().(LexerResult_Reject)
			return ok && _dafny.AreEqual(data1.Err, data2.Err)
		}
	case LexerResult_Partial:
		{
			data2, ok := other.Get_().(LexerResult_Partial)
			return ok && _dafny.AreEqual(data1.St, data2.St)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this LexerResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(LexerResult)
	return ok && _this.Equals(typed)
}

func Type_LexerResult_() _dafny.TypeDescriptor {
	return type_LexerResult_{}
}

type type_LexerResult_ struct {
}

func (_this type_LexerResult_) Default() interface{} {
	return Companion_LexerResult_.Default()
}

func (_this type_LexerResult_) String() string {
	return "Std_JSON_Utils_Lexers_Core.LexerResult"
}
func (_this LexerResult) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = LexerResult{}

// End of datatype LexerResult
