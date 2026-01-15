// Package Std_JSON_ZeroCopy_Deserializer_Values
// Dafny module Std_JSON_ZeroCopy_Deserializer_Values compiled into Go

package Std_JSON_ZeroCopy_Deserializer_Values

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
	m_Std_JSON_ZeroCopy_Deserializer_ArrayParams "Std_JSON_ZeroCopy_Deserializer_ArrayParams"
	m_Std_JSON_ZeroCopy_Deserializer_Arrays "Std_JSON_ZeroCopy_Deserializer_Arrays"
	m_Std_JSON_ZeroCopy_Deserializer_Constants "Std_JSON_ZeroCopy_Deserializer_Constants"
	m_Std_JSON_ZeroCopy_Deserializer_Core "Std_JSON_ZeroCopy_Deserializer_Core"
	m_Std_JSON_ZeroCopy_Deserializer_Numbers "Std_JSON_ZeroCopy_Deserializer_Numbers"
	m_Std_JSON_ZeroCopy_Deserializer_ObjectParams "Std_JSON_ZeroCopy_Deserializer_ObjectParams"
	m_Std_JSON_ZeroCopy_Deserializer_Objects "Std_JSON_ZeroCopy_Deserializer_Objects"
	m_Std_JSON_ZeroCopy_Deserializer_Strings "Std_JSON_ZeroCopy_Deserializer_Strings"
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
var _ m_Std_JSON_ZeroCopy_Deserializer_Strings.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Numbers.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_ObjectParams.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Objects.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Arrays.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Constants.Dummy__

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
	return "Std_JSON_ZeroCopy_Deserializer_Values.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Value(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _0_c int16 = (cs).Peek()
	_ = _0_c
	if (_0_c) == (int16(_dafny.CodePoint('{'))) {
		var _1_valueOrError0 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_Objects.Companion_Default___.Object(cs, Companion_Default___.ValueParser(cs))
		_ = _1_valueOrError0
		if (_1_valueOrError0).IsFailure() {
			return (_1_valueOrError0).PropagateFailure()
		} else {
			var _let_tmp_rhs0 m_Std_JSON_Utils_Cursors.Split = (_1_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _let_tmp_rhs0
			var _2_obj m_Std_JSON_Grammar.Bracketed = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Grammar.Bracketed)
			_ = _2_obj
			var _3_cs_k m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _3_cs_k
			var _4_v m_Std_JSON_Grammar.Value = m_Std_JSON_Grammar.Companion_Value_.Create_Object_(_2_obj)
			_ = _4_v
			var _5_sp m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(_4_v, _3_cs_k)
			_ = _5_sp
			return m_Std_Wrappers.Companion_Result_.Create_Success_(_5_sp)
		}
	} else if (_0_c) == (int16(_dafny.CodePoint('['))) {
		var _6_valueOrError1 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_Arrays.Companion_Default___.Array(cs, Companion_Default___.ValueParser(cs))
		_ = _6_valueOrError1
		if (_6_valueOrError1).IsFailure() {
			return (_6_valueOrError1).PropagateFailure()
		} else {
			var _let_tmp_rhs1 m_Std_JSON_Utils_Cursors.Split = (_6_valueOrError1).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _let_tmp_rhs1
			var _7_arr m_Std_JSON_Grammar.Bracketed = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Grammar.Bracketed)
			_ = _7_arr
			var _8_cs_k m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _8_cs_k
			var _9_v m_Std_JSON_Grammar.Value = m_Std_JSON_Grammar.Companion_Value_.Create_Array_(_7_arr)
			_ = _9_v
			var _10_sp m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(_9_v, _8_cs_k)
			_ = _10_sp
			return m_Std_Wrappers.Companion_Result_.Create_Success_(_10_sp)
		}
	} else if (_0_c) == (int16(_dafny.CodePoint('"'))) {
		var _11_valueOrError2 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_Strings.Companion_Default___.String_(cs)
		_ = _11_valueOrError2
		if (_11_valueOrError2).IsFailure() {
			return (_11_valueOrError2).PropagateFailure()
		} else {
			var _let_tmp_rhs2 m_Std_JSON_Utils_Cursors.Split = (_11_valueOrError2).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _let_tmp_rhs2
			var _12_str m_Std_JSON_Grammar.Jstring = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Grammar.Jstring)
			_ = _12_str
			var _13_cs_k m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _13_cs_k
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Value_.Create_String_(_12_str), _13_cs_k))
		}
	} else if (_0_c) == (int16(_dafny.CodePoint('t'))) {
		var _14_valueOrError3 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_Constants.Companion_Default___.Constant(cs, m_Std_JSON_Grammar.Companion_Default___.TRUE())
		_ = _14_valueOrError3
		if (_14_valueOrError3).IsFailure() {
			return (_14_valueOrError3).PropagateFailure()
		} else {
			var _let_tmp_rhs3 m_Std_JSON_Utils_Cursors.Split = (_14_valueOrError3).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _let_tmp_rhs3
			var _15_cst m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs3.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
			_ = _15_cst
			var _16_cs_k m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs3.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _16_cs_k
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Value_.Create_Bool_(_15_cst), _16_cs_k))
		}
	} else if (_0_c) == (int16(_dafny.CodePoint('f'))) {
		var _17_valueOrError4 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_Constants.Companion_Default___.Constant(cs, m_Std_JSON_Grammar.Companion_Default___.FALSE())
		_ = _17_valueOrError4
		if (_17_valueOrError4).IsFailure() {
			return (_17_valueOrError4).PropagateFailure()
		} else {
			var _let_tmp_rhs4 m_Std_JSON_Utils_Cursors.Split = (_17_valueOrError4).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _let_tmp_rhs4
			var _18_cst m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs4.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
			_ = _18_cst
			var _19_cs_k m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs4.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _19_cs_k
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Value_.Create_Bool_(_18_cst), _19_cs_k))
		}
	} else if (_0_c) == (int16(_dafny.CodePoint('n'))) {
		var _20_valueOrError5 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_Constants.Companion_Default___.Constant(cs, m_Std_JSON_Grammar.Companion_Default___.NULL())
		_ = _20_valueOrError5
		if (_20_valueOrError5).IsFailure() {
			return (_20_valueOrError5).PropagateFailure()
		} else {
			var _let_tmp_rhs5 m_Std_JSON_Utils_Cursors.Split = (_20_valueOrError5).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _let_tmp_rhs5
			var _21_cst m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs5.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
			_ = _21_cst
			var _22_cs_k m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs5.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _22_cs_k
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Value_.Create_Null_(_21_cst), _22_cs_k))
		}
	} else {
		var _23_valueOrError6 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_Numbers.Companion_Default___.Number(cs)
		_ = _23_valueOrError6
		if (_23_valueOrError6).IsFailure() {
			return (_23_valueOrError6).PropagateFailure()
		} else {
			var _let_tmp_rhs6 m_Std_JSON_Utils_Cursors.Split = (_23_valueOrError6).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _let_tmp_rhs6
			var _24_num m_Std_JSON_Grammar.Jnumber = _let_tmp_rhs6.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Grammar.Jnumber)
			_ = _24_num
			var _25_cs_k m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs6.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _25_cs_k
			var _26_v m_Std_JSON_Grammar.Value = m_Std_JSON_Grammar.Companion_Value_.Create_Number_(_24_num)
			_ = _26_v
			var _27_sp m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(_26_v, _25_cs_k)
			_ = _27_sp
			return m_Std_Wrappers.Companion_Result_.Create_Success_(_27_sp)
		}
	}
}
func (_static *CompanionStruct_Default___) ValueParser(cs m_Std_JSON_Utils_Cursors.Cursor__) func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _0_pre func(m_Std_JSON_Utils_Cursors.Cursor__) bool = (func(_1_cs m_Std_JSON_Utils_Cursors.Cursor__) func(m_Std_JSON_Utils_Cursors.Cursor__) bool {
		return func(_2_ps_k m_Std_JSON_Utils_Cursors.Cursor__) bool {
			return ((_2_ps_k).Length()) < ((_1_cs).Length())
		}
	})(cs)
	_ = _0_pre
	var _3_fn func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result = (func(_4_pre func(m_Std_JSON_Utils_Cursors.Cursor__) bool) func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
		return func(_5_ps_k m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
			return Companion_Default___.Value(_5_ps_k)
		}
	})(_0_pre)
	_ = _3_fn
	return _3_fn
}

// End of class Default__
