// Package Std_JSON_ZeroCopy_Deserializer_Numbers
// Dafny module Std_JSON_ZeroCopy_Deserializer_Numbers compiled into Go

package Std_JSON_ZeroCopy_Deserializer_Numbers

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
	return "Std_JSON_ZeroCopy_Deserializer_Numbers.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Digits(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_JSON_Utils_Cursors.Split {
	return ((cs).SkipWhile(m_Std_JSON_Grammar.Companion_Default___.Digit_q)).Split()
}
func (_static *CompanionStruct_Default___) NonEmptyDigits(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _0_sp m_Std_JSON_Utils_Cursors.Split = Companion_Default___.Digits(cs)
	_ = _0_sp
	if ((_0_sp).Dtor_t().(m_Std_JSON_Utils_Views_Core.View__)).Empty_q() {
		return m_Std_Wrappers.Companion_Result_.Create_Failure_(m_Std_JSON_Utils_Cursors.Companion_CursorError_.Create_OtherError_(m_Std_JSON_Errors.Companion_DeserializationError_.Create_EmptyNumber_()))
	} else {
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_0_sp)
	}
}
func (_static *CompanionStruct_Default___) NonZeroInt(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	return Companion_Default___.NonEmptyDigits(cs)
}
func (_static *CompanionStruct_Default___) OptionalMinus(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_JSON_Utils_Cursors.Split {
	return ((cs).SkipIf(func(_0_c uint8) bool {
		return (_0_c) == (uint8(_dafny.CodePoint('-')))
	})).Split()
}
func (_static *CompanionStruct_Default___) OptionalSign(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_JSON_Utils_Cursors.Split {
	return ((cs).SkipIf(func(_0_c uint8) bool {
		return ((_0_c) == (uint8(_dafny.CodePoint('-')))) || ((_0_c) == (uint8(_dafny.CodePoint('+'))))
	})).Split()
}
func (_static *CompanionStruct_Default___) TrimmedInt(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _0_sp m_Std_JSON_Utils_Cursors.Split = ((cs).SkipIf(func(_1_c uint8) bool {
		return (_1_c) == (uint8(_dafny.CodePoint('0')))
	})).Split()
	_ = _0_sp
	if ((_0_sp).Dtor_t().(m_Std_JSON_Utils_Views_Core.View__)).Empty_q() {
		return Companion_Default___.NonZeroInt((_0_sp).Dtor_cs())
	} else {
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_0_sp)
	}
}
func (_static *CompanionStruct_Default___) Exp(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _let_tmp_rhs0 m_Std_JSON_Utils_Cursors.Split = ((cs).SkipIf(func(_0_c uint8) bool {
		return ((_0_c) == (uint8(_dafny.CodePoint('e')))) || ((_0_c) == (uint8(_dafny.CodePoint('E'))))
	})).Split()
	_ = _let_tmp_rhs0
	var _1_e m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
	_ = _1_e
	var _2_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
	_ = _2_cs
	if (_1_e).Empty_q() {
		return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Maybe_.Create_Empty_(), _2_cs))
	} else {
		var _let_tmp_rhs1 m_Std_JSON_Utils_Cursors.Split = Companion_Default___.OptionalSign(_2_cs)
		_ = _let_tmp_rhs1
		var _3_sign m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
		_ = _3_sign
		var _4_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
		_ = _4_cs
		var _5_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.NonEmptyDigits(_4_cs)
		_ = _5_valueOrError0
		if (_5_valueOrError0).IsFailure() {
			return (_5_valueOrError0).PropagateFailure()
		} else {
			var _let_tmp_rhs2 m_Std_JSON_Utils_Cursors.Split = (_5_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _let_tmp_rhs2
			var _6_num m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
			_ = _6_num
			var _7_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs2.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _7_cs
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Maybe_.Create_NonEmpty_(m_Std_JSON_Grammar.Companion_Jexp_.Create_JExp_(_1_e, _3_sign, _6_num)), _7_cs))
		}
	}
}
func (_static *CompanionStruct_Default___) Frac(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _let_tmp_rhs0 m_Std_JSON_Utils_Cursors.Split = ((cs).SkipIf(func(_0_c uint8) bool {
		return (_0_c) == (uint8(_dafny.CodePoint('.')))
	})).Split()
	_ = _let_tmp_rhs0
	var _1_period m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
	_ = _1_period
	var _2_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
	_ = _2_cs
	if (_1_period).Empty_q() {
		return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Maybe_.Create_Empty_(), _2_cs))
	} else {
		var _3_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.NonEmptyDigits(_2_cs)
		_ = _3_valueOrError0
		if (_3_valueOrError0).IsFailure() {
			return (_3_valueOrError0).PropagateFailure()
		} else {
			var _let_tmp_rhs1 m_Std_JSON_Utils_Cursors.Split = (_3_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _let_tmp_rhs1
			var _4_num m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).T.(m_Std_JSON_Utils_Views_Core.View__)
			_ = _4_num
			var _5_cs m_Std_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs1.Get_().(m_Std_JSON_Utils_Cursors.Split_SP).Cs
			_ = _5_cs
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Maybe_.Create_NonEmpty_(m_Std_JSON_Grammar.Companion_Jfrac_.Create_JFrac_(_1_period, _4_num)), _5_cs))
		}
	}
}
func (_static *CompanionStruct_Default___) NumberFromParts(minus m_Std_JSON_Utils_Cursors.Split, num m_Std_JSON_Utils_Cursors.Split, frac m_Std_JSON_Utils_Cursors.Split, exp m_Std_JSON_Utils_Cursors.Split) m_Std_JSON_Utils_Cursors.Split {
	var _0_sp m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Jnumber_.Create_JNumber_((minus).Dtor_t().(m_Std_JSON_Utils_Views_Core.View__), (num).Dtor_t().(m_Std_JSON_Utils_Views_Core.View__), (frac).Dtor_t().(m_Std_JSON_Grammar.Maybe), (exp).Dtor_t().(m_Std_JSON_Grammar.Maybe)), (exp).Dtor_cs())
	_ = _0_sp
	return _0_sp
}
func (_static *CompanionStruct_Default___) Number(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _0_minus m_Std_JSON_Utils_Cursors.Split = Companion_Default___.OptionalMinus(cs)
	_ = _0_minus
	var _1_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.TrimmedInt((_0_minus).Dtor_cs())
	_ = _1_valueOrError0
	if (_1_valueOrError0).IsFailure() {
		return (_1_valueOrError0).PropagateFailure()
	} else {
		var _2_num m_Std_JSON_Utils_Cursors.Split = (_1_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Split)
		_ = _2_num
		var _3_valueOrError1 m_Std_Wrappers.Result = Companion_Default___.Frac((_2_num).Dtor_cs())
		_ = _3_valueOrError1
		if (_3_valueOrError1).IsFailure() {
			return (_3_valueOrError1).PropagateFailure()
		} else {
			var _4_frac m_Std_JSON_Utils_Cursors.Split = (_3_valueOrError1).Extract().(m_Std_JSON_Utils_Cursors.Split)
			_ = _4_frac
			var _5_valueOrError2 m_Std_Wrappers.Result = Companion_Default___.Exp((_4_frac).Dtor_cs())
			_ = _5_valueOrError2
			if (_5_valueOrError2).IsFailure() {
				return (_5_valueOrError2).PropagateFailure()
			} else {
				var _6_exp m_Std_JSON_Utils_Cursors.Split = (_5_valueOrError2).Extract().(m_Std_JSON_Utils_Cursors.Split)
				_ = _6_exp
				return m_Std_Wrappers.Companion_Result_.Create_Success_(Companion_Default___.NumberFromParts(_0_minus, _2_num, _4_frac, _6_exp))
			}
		}
	}
}

// End of class Default__
