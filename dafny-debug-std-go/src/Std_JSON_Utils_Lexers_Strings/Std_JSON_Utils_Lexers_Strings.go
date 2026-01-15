// Package Std_JSON_Utils_Lexers_Strings
// Dafny module Std_JSON_Utils_Lexers_Strings compiled into Go

package Std_JSON_Utils_Lexers_Strings

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
	m_Std_JSON_Utils_Lexers_Core "Std_JSON_Utils_Lexers_Core"
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
	return "Std_JSON_Utils_Lexers_Strings.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) StringBody(escaped bool, byte_ int16) m_Std_JSON_Utils_Lexers_Core.LexerResult {
	if (byte_) == (int16(_dafny.CodePoint('\\'))) {
		return m_Std_JSON_Utils_Lexers_Core.Companion_LexerResult_.Create_Partial_(!(escaped))
	} else if ((byte_) == (int16(_dafny.CodePoint('"')))) && (!(escaped)) {
		return m_Std_JSON_Utils_Lexers_Core.Companion_LexerResult_.Create_Accept_()
	} else {
		return m_Std_JSON_Utils_Lexers_Core.Companion_LexerResult_.Create_Partial_(false)
	}
}
func (_static *CompanionStruct_Default___) String_(st StringLexerState, byte_ int16) m_Std_JSON_Utils_Lexers_Core.LexerResult {
	var _source0 StringLexerState = st
	_ = _source0
	{
		if _source0.Is_Start() {
			if (byte_) == (int16(_dafny.CodePoint('"'))) {
				return m_Std_JSON_Utils_Lexers_Core.Companion_LexerResult_.Create_Partial_(Companion_StringLexerState_.Create_Body_(false))
			} else {
				return m_Std_JSON_Utils_Lexers_Core.Companion_LexerResult_.Create_Reject_(_dafny.UnicodeSeqOfUtf8Bytes("String must start with double quote"))
			}
		}
	}
	{
		if _source0.Is_End() {
			return m_Std_JSON_Utils_Lexers_Core.Companion_LexerResult_.Create_Accept_()
		}
	}
	{
		var _0_escaped bool = _source0.Get_().(StringLexerState_Body).Escaped
		_ = _0_escaped
		if (byte_) == (int16(_dafny.CodePoint('\\'))) {
			return m_Std_JSON_Utils_Lexers_Core.Companion_LexerResult_.Create_Partial_(Companion_StringLexerState_.Create_Body_(!(_0_escaped)))
		} else if ((byte_) == (int16(_dafny.CodePoint('"')))) && (!(_0_escaped)) {
			return m_Std_JSON_Utils_Lexers_Core.Companion_LexerResult_.Create_Partial_(Companion_StringLexerState_.Create_End_())
		} else {
			return m_Std_JSON_Utils_Lexers_Core.Companion_LexerResult_.Create_Partial_(Companion_StringLexerState_.Create_Body_(false))
		}
	}
}
func (_static *CompanionStruct_Default___) StringBodyLexerStart() bool {
	return false
}
func (_static *CompanionStruct_Default___) StringLexerStart() StringLexerState {
	return Companion_StringLexerState_.Create_Start_()
}

// End of class Default__

// Definition of datatype StringLexerState
type StringLexerState struct {
	Data_StringLexerState_
}

func (_this StringLexerState) Get_() Data_StringLexerState_ {
	return _this.Data_StringLexerState_
}

type Data_StringLexerState_ interface {
	isStringLexerState()
}

type CompanionStruct_StringLexerState_ struct {
}

var Companion_StringLexerState_ = CompanionStruct_StringLexerState_{}

type StringLexerState_Start struct {
}

func (StringLexerState_Start) isStringLexerState() {}

func (CompanionStruct_StringLexerState_) Create_Start_() StringLexerState {
	return StringLexerState{StringLexerState_Start{}}
}

func (_this StringLexerState) Is_Start() bool {
	_, ok := _this.Get_().(StringLexerState_Start)
	return ok
}

type StringLexerState_Body struct {
	Escaped bool
}

func (StringLexerState_Body) isStringLexerState() {}

func (CompanionStruct_StringLexerState_) Create_Body_(Escaped bool) StringLexerState {
	return StringLexerState{StringLexerState_Body{Escaped}}
}

func (_this StringLexerState) Is_Body() bool {
	_, ok := _this.Get_().(StringLexerState_Body)
	return ok
}

type StringLexerState_End struct {
}

func (StringLexerState_End) isStringLexerState() {}

func (CompanionStruct_StringLexerState_) Create_End_() StringLexerState {
	return StringLexerState{StringLexerState_End{}}
}

func (_this StringLexerState) Is_End() bool {
	_, ok := _this.Get_().(StringLexerState_End)
	return ok
}

func (CompanionStruct_StringLexerState_) Default() StringLexerState {
	return Companion_StringLexerState_.Create_Start_()
}

func (_this StringLexerState) Dtor_escaped() bool {
	return _this.Get_().(StringLexerState_Body).Escaped
}

func (_this StringLexerState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case StringLexerState_Start:
		{
			return "Strings.StringLexerState.Start"
		}
	case StringLexerState_Body:
		{
			return "Strings.StringLexerState.Body" + "(" + _dafny.String(data.Escaped) + ")"
		}
	case StringLexerState_End:
		{
			return "Strings.StringLexerState.End"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this StringLexerState) Equals(other StringLexerState) bool {
	switch data1 := _this.Get_().(type) {
	case StringLexerState_Start:
		{
			_, ok := other.Get_().(StringLexerState_Start)
			return ok
		}
	case StringLexerState_Body:
		{
			data2, ok := other.Get_().(StringLexerState_Body)
			return ok && data1.Escaped == data2.Escaped
		}
	case StringLexerState_End:
		{
			_, ok := other.Get_().(StringLexerState_End)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this StringLexerState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(StringLexerState)
	return ok && _this.Equals(typed)
}

func Type_StringLexerState_() _dafny.TypeDescriptor {
	return type_StringLexerState_{}
}

type type_StringLexerState_ struct {
}

func (_this type_StringLexerState_) Default() interface{} {
	return Companion_StringLexerState_.Default()
}

func (_this type_StringLexerState_) String() string {
	return "Std_JSON_Utils_Lexers_Strings.StringLexerState"
}
func (_this StringLexerState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = StringLexerState{}

// End of datatype StringLexerState
