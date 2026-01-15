// Package Std_JSON_ZeroCopy_Deserializer_Core
// Dafny module Std_JSON_ZeroCopy_Deserializer_Core compiled into Go

package Std_JSON_ZeroCopy_Deserializer_Core

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
	m_Std_JSON_ByteStrConversion "Std_JSON_ByteStrConversion"
	m_Std_JSON_ConcreteSyntax_Spec "Std_JSON_ConcreteSyntax_Spec"
	m_Std_JSON_ConcreteSyntax_SpecProperties "Std_JSON_ConcreteSyntax_SpecProperties"
	m_Std_JSON_Deserializer "Std_JSON_Deserializer"
	m_Std_JSON_Deserializer_Uint16StrConversion "Std_JSON_Deserializer_Uint16StrConversion"
	m_Std_JSON_Errors "Std_JSON_Errors"
	m_Std_JSON_Grammar "Std_JSON_Grammar"
	m_Std_JSON_Serializer "Std_JSON_Serializer"
	m_Std_JSON_Spec "Std_JSON_Spec"
	m_Std_JSON_Utils_Cursors "Std_JSON_Utils_Cursors"
	m_Std_JSON_Utils_Lexers_Core "Std_JSON_Utils_Lexers_Core"
	m_Std_JSON_Utils_Lexers_Strings "Std_JSON_Utils_Lexers_Strings"
	m_Std_JSON_Utils_Parsers "Std_JSON_Utils_Parsers"
	m_Std_JSON_Utils_Views_Core "Std_JSON_Utils_Views_Core"
	m_Std_JSON_Utils_Views_Writers "Std_JSON_Utils_Views_Writers"
	m_Std_JSON_Values "Std_JSON_Values"
	m_Std_JSON_ZeroCopy_Serializer "Std_JSON_ZeroCopy_Serializer"
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
var _ m_Std_JSON_Utils_Parsers.Dummy__
var _ m_Std_JSON_Grammar.Dummy__
var _ m_Std_JSON_ByteStrConversion.Dummy__
var _ m_Std_JSON_Serializer.Dummy__
var _ m_Std_JSON_Deserializer_Uint16StrConversion.Dummy__
var _ m_Std_JSON_Deserializer.Dummy__
var _ m_Std_JSON_ConcreteSyntax_Spec.Dummy__
var _ m_Std_JSON_ConcreteSyntax_SpecProperties.Dummy__
var _ m_Std_JSON_ZeroCopy_Serializer.Dummy__

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
	return "Std_JSON_ZeroCopy_Deserializer_Core.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Get(cs m_Std_JSON_Utils_Cursors.Cursor__, err m_Std_JSON_Errors.DeserializationError) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = (cs).Get(err)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_cs m_Std_JSON_Utils_Cursors.Cursor__ = (_0_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Cursor__)
		_ = _1_cs
		return m_Std_Wrappers.Companion_Result_.Create_Success_((_1_cs).Split())
	}
}
func (_static *CompanionStruct_Default___) WS(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_JSON_Utils_Cursors.Split {
	var sp m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Default(m_Std_JSON_Grammar.Companion_Jblanks_.Witness())
	_ = sp
	var _0_point_k uint32
	_ = _0_point_k
	_0_point_k = (cs).Dtor_point()
	var _1_end uint32
	_ = _1_end
	_1_end = (cs).Dtor_end()
	for ((_0_point_k) < (_1_end)) && (m_Std_JSON_Grammar.Companion_Default___.Blank_q(((cs).Dtor_s()).Select(uint32(_0_point_k)).(uint8))) {
		_0_point_k = (_0_point_k) + (uint32(1))
	}
	sp = (m_Std_JSON_Utils_Cursors.Companion_Cursor___.Create_Cursor_((cs).Dtor_s(), (cs).Dtor_beg(), _0_point_k, (cs).Dtor_end())).Split()
	return sp
	return sp
}
func (_static *CompanionStruct_Default___) Structural(cs m_Std_JSON_Utils_Cursors.Cursor__, parser func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result) m_Std_Wrappers.Result {
	var _let_tmp_rhs0 m_Std_JSON_Utils_Cursors.Split = Companion_Default___.WS(cs)
	_ = _let_tmp_rhs0
	var _0_before m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
	_ = _0_before
	var _1_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
	_ = _1_cs
	var _2_valueOrError0 m_Std_Wrappers.Result = (parser)(_1_cs)
	_ = _2_valueOrError0
	if (_2_valueOrError0).IsFailure() {
		return (_2_valueOrError0).PropagateFailure()
	} else {
		var _let_tmp_rhs1 m_Std_JSON_Utils_Cursors.Split = (_2_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Split)
		_ = _let_tmp_rhs1
		var _3_val interface{} = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T
		_ = _3_val
		var _4_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
		_ = _4_cs
		var _let_tmp_rhs2 m_Std_JSON_Utils_Cursors.Split = Companion_Default___.WS(_4_cs)
		_ = _let_tmp_rhs2
		var _5_after m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
		_ = _5_after
		var _6_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
		_ = _6_cs
		return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Structural_.Create_Structural_(_0_before, _3_val, _5_after), _6_cs))
	}
}
func (_static *CompanionStruct_Default___) TryStructural(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_JSON_Utils_Cursors.Split {
	var _let_tmp_rhs0 m_Std_JSON_Utils_Cursors.Split = Companion_Default___.WS(cs)
	_ = _let_tmp_rhs0
	var _0_before m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
	_ = _0_before
	var _1_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
	_ = _1_cs
	var _let_tmp_rhs1 m_Std_JSON_Utils_Cursors.Split = ((_1_cs).SkipByte()).Split()
	_ = _let_tmp_rhs1
	var _2_val m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
	_ = _2_val
	var _3_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
	_ = _3_cs
	var _let_tmp_rhs2 m_Std_JSON_Utils_Cursors.Split = Companion_Default___.WS(_3_cs)
	_ = _let_tmp_rhs2
	var _4_after m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
	_ = _4_after
	var _5_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
	_ = _5_cs
	return m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Structural_.Create_Structural_(_0_before, _2_val, _4_after), _5_cs)
}
func (_static *CompanionStruct_Default___) SpecView() func(m_Std_JSON_Utils_Views_Core.View__) _dafny.Sequence {
	return func(_0_v m_Std_JSON_Utils_Views_Core.View__) _dafny.Sequence {
		return m_Std_JSON_ConcreteSyntax_Spec.Companion_Default___.View(_0_v)
	}
}

// End of class Default__

// Definition of class Jopt
type Jopt struct {
}

func New_Jopt_() *Jopt {
	_this := Jopt{}

	return &_this
}

type CompanionStruct_Jopt_ struct {
}

var Companion_Jopt_ = CompanionStruct_Jopt_{}

func (*Jopt) String() string {
	return "Std_JSON_ZeroCopy_Deserializer_Core.Jopt"
}
func (_this *CompanionStruct_Jopt_) Witness() m_Std_JSON_Utils_Views_Core.View__ {
	return m_Std_JSON_Utils_Views_Core.Companion_View___.OfBytes(_dafny.SeqOf())
}

// End of class Jopt

func Type_Jopt_() _dafny.TypeDescriptor {
	return type_Jopt_{}
}

type type_Jopt_ struct {
}

func (_this type_Jopt_) Default() interface{} {
	return Companion_Jopt_.Witness()
}

func (_this type_Jopt_) String() string {
	return "Std_JSON_ZeroCopy_Deserializer_Core.Jopt"
}

// Definition of class ValueParser
type ValueParser struct {
}

func New_ValueParser_() *ValueParser {
	_this := ValueParser{}

	return &_this
}

type CompanionStruct_ValueParser_ struct {
}

var Companion_ValueParser_ = CompanionStruct_ValueParser_{}

func (*ValueParser) String() string {
	return "Std_JSON_ZeroCopy_Deserializer_Core.ValueParser"
}

// End of class ValueParser

func Type_ValueParser_() _dafny.TypeDescriptor {
	return type_ValueParser_{}
}

type type_ValueParser_ struct {
}

func (_this type_ValueParser_) Default() interface{} {
	return m_Std_JSON_Utils_Parsers.Companion_SubParser_.Witness()
}

func (_this type_ValueParser_) String() string {
	return "Std_JSON_ZeroCopy_Deserializer_Core.ValueParser"
}
