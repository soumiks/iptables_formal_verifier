// Package Std_JSON_ZeroCopy_Deserializer_Strings
// Dafny module Std_JSON_ZeroCopy_Deserializer_Strings compiled into Go

package Std_JSON_ZeroCopy_Deserializer_Strings

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
	m_Std_JSON_ZeroCopy_Deserializer_Core "Std_JSON_ZeroCopy_Deserializer_Core"
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
var _ m_Std_JSON_ZeroCopy_Deserializer_Core.Dummy__

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
	return "Std_JSON_ZeroCopy_Deserializer_Strings.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) StringBody(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var pr m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Default(m_Std_JSON_Utils_Cursors.Companion_Cursor_.Witness())
	_ = pr
	var _0_escaped bool
	_ = _0_escaped
	_0_escaped = false
	var _hi0 uint32 = (cs).Dtor_end()
	_ = _hi0
	for _1_point_k := (cs).Dtor_point(); _1_point_k < _hi0; _1_point_k++ {
		var _2_byte uint8
		_ = _2_byte
		_2_byte = ((cs).Dtor_s()).Select(uint32(_1_point_k)).(uint8)
		if ((_2_byte) == (uint8(_dafny.CodePoint('"')))) && (!(_0_escaped)) {
			pr = m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Cursor___.Create_Cursor_((cs).Dtor_s(), (cs).Dtor_beg(), _1_point_k, (cs).Dtor_end()))
			return pr
		} else if (_2_byte) == (uint8(_dafny.CodePoint('\\'))) {
			_0_escaped = !(_0_escaped)
		} else {
			_0_escaped = false
		}
	}
	pr = m_Std_Wrappers.Companion_Result_.Create_Failure_(m_Std_JSON_Utils_Cursors.Companion_CursorError_.Create_EOF_())
	return pr
	return pr
}
func (_static *CompanionStruct_Default___) Quote(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = (cs).AssertChar(_dafny.CodePoint('"'))
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_cs m_Std_JSON_Utils_Cursors.Cursor__ = (_0_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Cursor__)
		_ = _1_cs
		return m_Std_Wrappers.Companion_Result_.Create_Success_((_1_cs).Split())
	}
}
func (_static *CompanionStruct_Default___) String_(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _0_origCs m_Std_JSON_Utils_Cursors.Cursor__ = cs
	_ = _0_origCs
	var _1_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.Quote(cs)
	_ = _1_valueOrError0
	if (_1_valueOrError0).IsFailure() {
		return (_1_valueOrError0).PropagateFailure()
	} else {
		var _let_tmp_rhs0 m_Std_JSON_Utils_Cursors.Split = (_1_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Split)
		_ = _let_tmp_rhs0
		var _2_lq m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
		_ = _2_lq
		var _3_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
		_ = _3_cs
		var _4_valueOrError1 m_Std_Wrappers.Result = Companion_Default___.StringBody(_3_cs)
		_ = _4_valueOrError1
		if (_4_valueOrError1).IsFailure() {
			return (_4_valueOrError1).PropagateFailure()
		} else {
			var _5_contents m_Std_JSON_Utils_Cursors.Cursor__ = (_4_valueOrError1).Extract().(m_Std_JSON_Utils_Cursors.Cursor__)
			_ = _5_contents
			var _let_tmp_rhs1 m_Std_JSON_Utils_Cursors.Split = (_5_contents).Split()
			_ = _let_tmp_rhs1
			var _6_contents m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
			_ = _6_contents
			var _7_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _7_cs
			var _8_valueOrError2 m_Std_Wrappers.Result = Companion_Default___.Quote(_7_cs)
			_ = _8_valueOrError2
			if (_8_valueOrError2).IsFailure() {
				return (_8_valueOrError2).PropagateFailure()
			} else {
				var _let_tmp_rhs2 m_Std_JSON_Utils_Cursors.Split = (_8_valueOrError2).Extract().(m_Std_JSON_Utils_Cursors.Split)
				_ = _let_tmp_rhs2
				var _9_rq m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
				_ = _9_rq
				var _10_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
				_ = _10_cs
				var _11_result m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Jstring_.Create_JString_(_2_lq, _6_contents, _9_rq), _10_cs)
				_ = _11_result
				return m_Std_Wrappers.Companion_Result_.Create_Success_(_11_result)
			}
		}
	}
}

// End of class Default__
