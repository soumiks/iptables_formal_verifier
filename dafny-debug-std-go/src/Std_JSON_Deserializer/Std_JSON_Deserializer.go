// Package Std_JSON_Deserializer
// Dafny module Std_JSON_Deserializer compiled into Go

package Std_JSON_Deserializer

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
	return "Std_JSON_Deserializer.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Bool(js m_Std_JSON_Utils_Views_Core.View__) bool {
	return ((js).At(uint32(0))) == (uint8(_dafny.CodePoint('t')))
}
func (_static *CompanionStruct_Default___) UnsupportedEscape16(code _dafny.Sequence) m_Std_JSON_Errors.DeserializationError {
	return m_Std_JSON_Errors.Companion_DeserializationError_.Create_UnsupportedEscape_(((m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.FromUTF16Checked(code)).ToOption()).GetOr(_dafny.UnicodeSeqOfUtf8Bytes("Couldn't decode UTF-16")).(_dafny.Sequence))
}
func (_static *CompanionStruct_Default___) ToNat16(str _dafny.Sequence) uint16 {
	var _0_hd _dafny.Int = m_Std_JSON_Deserializer_Uint16StrConversion.Companion_Default___.ToNat(str)
	_ = _0_hd
	return (_0_hd).Uint16()
}
func (_static *CompanionStruct_Default___) Unescape(str _dafny.Sequence, start _dafny.Int, prefix _dafny.Sequence) m_Std_Wrappers.Result {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (start).Cmp(_dafny.IntOfUint32((str).Cardinality())) >= 0 {
		return m_Std_Wrappers.Companion_Result_.Create_Success_(prefix)
	} else if ((str).Select((start).Uint32()).(uint16)) == (uint16(_dafny.CodePoint('\\'))) {
		if (_dafny.IntOfUint32((str).Cardinality())).Cmp((start).Plus(_dafny.One)) == 0 {
			return m_Std_Wrappers.Companion_Result_.Create_Failure_(m_Std_JSON_Errors.Companion_DeserializationError_.Create_EscapeAtEOS_())
		} else {
			var _0_c uint16 = (str).Select(((start).Plus(_dafny.One)).Uint32()).(uint16)
			_ = _0_c
			if (_0_c) == (uint16(_dafny.CodePoint('u'))) {
				if (_dafny.IntOfUint32((str).Cardinality())).Cmp((start).Plus(_dafny.IntOfInt64(6))) < 0 {
					return m_Std_Wrappers.Companion_Result_.Create_Failure_(m_Std_JSON_Errors.Companion_DeserializationError_.Create_EscapeAtEOS_())
				} else {
					var _1_code _dafny.Sequence = (str).Subsequence(((start).Plus(_dafny.IntOfInt64(2))).Uint32(), ((start).Plus(_dafny.IntOfInt64(6))).Uint32())
					_ = _1_code
					if _dafny.Quantifier((_1_code).UniqueElements(), false, func(_exists_var_0 uint16) bool {
						var _2_c uint16
						_2_c = interface{}(_exists_var_0).(uint16)
						if true {
							return (_dafny.Companion_Sequence_.Contains(_1_code, _2_c)) && (!(Companion_Default___.HEX__TABLE__16()).Contains(_2_c))
						} else {
							return false
						}
					}) {
						return m_Std_Wrappers.Companion_Result_.Create_Failure_(Companion_Default___.UnsupportedEscape16(_1_code))
					} else {
						var _3_hd uint16 = Companion_Default___.ToNat16(_1_code)
						_ = _3_hd
						var _in0 _dafny.Sequence = str
						_ = _in0
						var _in1 _dafny.Int = (start).Plus(_dafny.IntOfInt64(6))
						_ = _in1
						var _in2 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(prefix, _dafny.SeqOf(_3_hd))
						_ = _in2
						str = _in0
						start = _in1
						prefix = _in2
						goto TAIL_CALL_START
					}
				}
			} else {
				var _4_unescaped uint16 = func() uint16 {
					var _source0 uint16 = _0_c
					_ = _source0
					{
						if (_source0) == (uint16(34)) {
							return uint16(34)
						}
					}
					{
						if (_source0) == (uint16(92)) {
							return uint16(92)
						}
					}
					{
						if (_source0) == (uint16(98)) {
							return uint16(8)
						}
					}
					{
						if (_source0) == (uint16(102)) {
							return uint16(12)
						}
					}
					{
						if (_source0) == (uint16(110)) {
							return uint16(10)
						}
					}
					{
						if (_source0) == (uint16(114)) {
							return uint16(13)
						}
					}
					{
						if (_source0) == (uint16(116)) {
							return uint16(9)
						}
					}
					{
						return uint16(0)
					}
				}()
				_ = _4_unescaped
				if (_dafny.IntOfUint16(_4_unescaped)).Sign() == 0 {
					return m_Std_Wrappers.Companion_Result_.Create_Failure_(Companion_Default___.UnsupportedEscape16((str).Subsequence((start).Uint32(), ((start).Plus(_dafny.IntOfInt64(2))).Uint32())))
				} else {
					var _in3 _dafny.Sequence = str
					_ = _in3
					var _in4 _dafny.Int = (start).Plus(_dafny.IntOfInt64(2))
					_ = _in4
					var _in5 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(prefix, _dafny.SeqOf(_4_unescaped))
					_ = _in5
					str = _in3
					start = _in4
					prefix = _in5
					goto TAIL_CALL_START
				}
			}
		}
	} else {
		var _in6 _dafny.Sequence = str
		_ = _in6
		var _in7 _dafny.Int = (start).Plus(_dafny.One)
		_ = _in7
		var _in8 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(prefix, _dafny.SeqOf((str).Select((start).Uint32()).(uint16)))
		_ = _in8
		str = _in6
		start = _in7
		prefix = _in8
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) String_(js m_Std_JSON_Grammar.Jstring) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = (m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.FromUTF8Checked(((js).Dtor_contents()).Bytes())).MapFailure(func(coer43 func(_dafny.Sequence) m_Std_JSON_Errors.DeserializationError) func(interface{}) interface{} {
		return func(arg44 interface{}) interface{} {
			return coer43(arg44.(_dafny.Sequence))
		}
	}(func(_1_error _dafny.Sequence) m_Std_JSON_Errors.DeserializationError {
		return m_Std_JSON_Errors.Companion_DeserializationError_.Create_InvalidUnicode_(_1_error)
	}))
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _2_asUtf32 _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _2_asUtf32
		var _3_valueOrError1 m_Std_Wrappers.Result = (m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ToUTF16Checked(_2_asUtf32)).ToResult(m_Std_JSON_Errors.Companion_DeserializationError_.Create_InvalidUnicode_(_dafny.UnicodeSeqOfUtf8Bytes("")))
		_ = _3_valueOrError1
		if (_3_valueOrError1).IsFailure() {
			return (_3_valueOrError1).PropagateFailure()
		} else {
			var _4_asUint16 _dafny.Sequence = (_3_valueOrError1).Extract().(_dafny.Sequence)
			_ = _4_asUint16
			var _5_valueOrError2 m_Std_Wrappers.Result = Companion_Default___.Unescape(_4_asUint16, _dafny.Zero, _dafny.SeqOf())
			_ = _5_valueOrError2
			if (_5_valueOrError2).IsFailure() {
				return (_5_valueOrError2).PropagateFailure()
			} else {
				var _6_unescaped _dafny.Sequence = (_5_valueOrError2).Extract().(_dafny.Sequence)
				_ = _6_unescaped
				return (m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.FromUTF16Checked(_6_unescaped)).MapFailure(func(coer44 func(_dafny.Sequence) m_Std_JSON_Errors.DeserializationError) func(interface{}) interface{} {
					return func(arg45 interface{}) interface{} {
						return coer44(arg45.(_dafny.Sequence))
					}
				}(func(_7_error _dafny.Sequence) m_Std_JSON_Errors.DeserializationError {
					return m_Std_JSON_Errors.Companion_DeserializationError_.Create_InvalidUnicode_(_7_error)
				}))
			}
		}
	}
}
func (_static *CompanionStruct_Default___) ToInt(sign m_Std_JSON_Utils_Views_Core.View__, n m_Std_JSON_Utils_Views_Core.View__) m_Std_Wrappers.Result {
	var _0_n _dafny.Int = m_Std_JSON_ByteStrConversion.Companion_Default___.ToNat((n).Bytes())
	_ = _0_n
	return m_Std_Wrappers.Companion_Result_.Create_Success_((func() _dafny.Int {
		if (sign).Char_q(_dafny.CodePoint('-')) {
			return (_dafny.Zero).Minus(_0_n)
		}
		return _0_n
	})())
}
func (_static *CompanionStruct_Default___) Number(js m_Std_JSON_Grammar.Jnumber) m_Std_Wrappers.Result {
	var _let_tmp_rhs0 m_Std_JSON_Grammar.Jnumber = js
	_ = _let_tmp_rhs0
	var _0_minus m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Grammar.Jnumber_JNumber).Minus
	_ = _0_minus
	var _1_num m_Std_JSON_Utils_Views_Core.View__ = _let_tmp_rhs0.Get_().(m_Std_JSON_Grammar.Jnumber_JNumber).Num
	_ = _1_num
	var _2_frac m_Std_JSON_Grammar.Maybe = _let_tmp_rhs0.Get_().(m_Std_JSON_Grammar.Jnumber_JNumber).Frac
	_ = _2_frac
	var _3_exp m_Std_JSON_Grammar.Maybe = _let_tmp_rhs0.Get_().(m_Std_JSON_Grammar.Jnumber_JNumber).Exp
	_ = _3_exp
	var _4_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.ToInt(_0_minus, _1_num)
	_ = _4_valueOrError0
	if (_4_valueOrError0).IsFailure() {
		return (_4_valueOrError0).PropagateFailure()
	} else {
		var _5_n _dafny.Int = (_4_valueOrError0).Extract().(_dafny.Int)
		_ = _5_n
		var _6_valueOrError1 m_Std_Wrappers.Result = func() m_Std_Wrappers.Result {
			var _source0 m_Std_JSON_Grammar.Maybe = _3_exp
			_ = _source0
			{
				if _source0.Is_Empty() {
					return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.Zero)
				}
			}
			{
				var t0 m_Std_JSON_Grammar.Jexp = _source0.Get_().(m_Std_JSON_Grammar.Maybe_NonEmpty).T.(m_Std_JSON_Grammar.Jexp)
				_ = t0
				var _7_sign m_Std_JSON_Utils_Views_Core.View__ = t0.Get_().(m_Std_JSON_Grammar.Jexp_JExp).Sign
				_ = _7_sign
				var _8_num m_Std_JSON_Utils_Views_Core.View__ = t0.Get_().(m_Std_JSON_Grammar.Jexp_JExp).Num
				_ = _8_num
				return Companion_Default___.ToInt(_7_sign, _8_num)
			}
		}()
		_ = _6_valueOrError1
		if (_6_valueOrError1).IsFailure() {
			return (_6_valueOrError1).PropagateFailure()
		} else {
			var _9_e10 _dafny.Int = (_6_valueOrError1).Extract().(_dafny.Int)
			_ = _9_e10
			var _source1 m_Std_JSON_Grammar.Maybe = _2_frac
			_ = _source1
			{
				if _source1.Is_Empty() {
					return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Values.Companion_Decimal_.Create_Decimal_(_5_n, _9_e10))
				}
			}
			{
				var t1 m_Std_JSON_Grammar.Jfrac = _source1.Get_().(m_Std_JSON_Grammar.Maybe_NonEmpty).T.(m_Std_JSON_Grammar.Jfrac)
				_ = t1
				var _10_num m_Std_JSON_Utils_Views_Core.View__ = t1.Get_().(m_Std_JSON_Grammar.Jfrac_JFrac).Num
				_ = _10_num
				var _11_pow10 _dafny.Int = _dafny.IntOfUint32((_10_num).Length())
				_ = _11_pow10
				var _12_valueOrError2 m_Std_Wrappers.Result = Companion_Default___.ToInt(_0_minus, _10_num)
				_ = _12_valueOrError2
				if (_12_valueOrError2).IsFailure() {
					return (_12_valueOrError2).PropagateFailure()
				} else {
					var _13_frac _dafny.Int = (_12_valueOrError2).Extract().(_dafny.Int)
					_ = _13_frac
					return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Values.Companion_Decimal_.Create_Decimal_(((_5_n).Times(m_Std_Arithmetic_Power.Companion_Default___.Pow(_dafny.IntOfInt64(10), _11_pow10))).Plus(_13_frac), (_9_e10).Minus(_11_pow10)))
				}
			}
		}
	}
}
func (_static *CompanionStruct_Default___) KeyValue(js m_Std_JSON_Grammar.JKeyValue) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.String_((js).Dtor_k())
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_k _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _1_k
		var _2_valueOrError1 m_Std_Wrappers.Result = Companion_Default___.Value((js).Dtor_v())
		_ = _2_valueOrError1
		if (_2_valueOrError1).IsFailure() {
			return (_2_valueOrError1).PropagateFailure()
		} else {
			var _3_v m_Std_JSON_Values.JSON = (_2_valueOrError1).Extract().(m_Std_JSON_Values.JSON)
			_ = _3_v
			return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.TupleOf(_1_k, _3_v))
		}
	}
}
func (_static *CompanionStruct_Default___) Object(js m_Std_JSON_Grammar.Bracketed) m_Std_Wrappers.Result {
	var _0_f func(m_Std_JSON_Grammar.Suffixed) m_Std_Wrappers.Result = (func(_1_js m_Std_JSON_Grammar.Bracketed) func(m_Std_JSON_Grammar.Suffixed) m_Std_Wrappers.Result {
		return func(_2_d m_Std_JSON_Grammar.Suffixed) m_Std_Wrappers.Result {
			return Companion_Default___.KeyValue((_2_d).Dtor_t().(m_Std_JSON_Grammar.JKeyValue))
		}
	})(js)
	_ = _0_f
	return m_Std_Collections_Seq.Companion_Default___.MapWithResult(func(coer45 func(m_Std_JSON_Grammar.Suffixed) m_Std_Wrappers.Result) func(interface{}) m_Std_Wrappers.Result {
		return func(arg46 interface{}) m_Std_Wrappers.Result {
			return coer45(arg46.(m_Std_JSON_Grammar.Suffixed))
		}
	}(_0_f), (js).Dtor_data())
}
func (_static *CompanionStruct_Default___) Array(js m_Std_JSON_Grammar.Bracketed) m_Std_Wrappers.Result {
	var _0_f func(m_Std_JSON_Grammar.Suffixed) m_Std_Wrappers.Result = (func(_1_js m_Std_JSON_Grammar.Bracketed) func(m_Std_JSON_Grammar.Suffixed) m_Std_Wrappers.Result {
		return func(_2_d m_Std_JSON_Grammar.Suffixed) m_Std_Wrappers.Result {
			return Companion_Default___.Value((_2_d).Dtor_t().(m_Std_JSON_Grammar.Value))
		}
	})(js)
	_ = _0_f
	return m_Std_Collections_Seq.Companion_Default___.MapWithResult(func(coer46 func(m_Std_JSON_Grammar.Suffixed) m_Std_Wrappers.Result) func(interface{}) m_Std_Wrappers.Result {
		return func(arg47 interface{}) m_Std_Wrappers.Result {
			return coer46(arg47.(m_Std_JSON_Grammar.Suffixed))
		}
	}(_0_f), (js).Dtor_data())
}
func (_static *CompanionStruct_Default___) Value(js m_Std_JSON_Grammar.Value) m_Std_Wrappers.Result {
	var _source0 m_Std_JSON_Grammar.Value = js
	_ = _source0
	{
		if _source0.Is_Null() {
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Values.Companion_JSON_.Create_Null_())
		}
	}
	{
		if _source0.Is_Bool() {
			var _0_b m_Std_JSON_Utils_Views_Core.View__ = _source0.Get_().(m_Std_JSON_Grammar.Value_Bool).B
			_ = _0_b
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Values.Companion_JSON_.Create_Bool_(Companion_Default___.Bool(_0_b)))
		}
	}
	{
		if _source0.Is_String() {
			var _1_str m_Std_JSON_Grammar.Jstring = _source0.Get_().(m_Std_JSON_Grammar.Value_String).Str
			_ = _1_str
			var _2_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.String_(_1_str)
			_ = _2_valueOrError0
			if (_2_valueOrError0).IsFailure() {
				return (_2_valueOrError0).PropagateFailure()
			} else {
				var _3_s _dafny.Sequence = (_2_valueOrError0).Extract().(_dafny.Sequence)
				_ = _3_s
				return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Values.Companion_JSON_.Create_String_(_3_s))
			}
		}
	}
	{
		if _source0.Is_Number() {
			var _4_dec m_Std_JSON_Grammar.Jnumber = _source0.Get_().(m_Std_JSON_Grammar.Value_Number).Num
			_ = _4_dec
			var _5_valueOrError1 m_Std_Wrappers.Result = Companion_Default___.Number(_4_dec)
			_ = _5_valueOrError1
			if (_5_valueOrError1).IsFailure() {
				return (_5_valueOrError1).PropagateFailure()
			} else {
				var _6_n m_Std_JSON_Values.Decimal = (_5_valueOrError1).Extract().(m_Std_JSON_Values.Decimal)
				_ = _6_n
				return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Values.Companion_JSON_.Create_Number_(_6_n))
			}
		}
	}
	{
		if _source0.Is_Object() {
			var _7_obj m_Std_JSON_Grammar.Bracketed = _source0.Get_().(m_Std_JSON_Grammar.Value_Object).Obj
			_ = _7_obj
			var _8_valueOrError2 m_Std_Wrappers.Result = Companion_Default___.Object(_7_obj)
			_ = _8_valueOrError2
			if (_8_valueOrError2).IsFailure() {
				return (_8_valueOrError2).PropagateFailure()
			} else {
				var _9_o _dafny.Sequence = (_8_valueOrError2).Extract().(_dafny.Sequence)
				_ = _9_o
				return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Values.Companion_JSON_.Create_Object_(_9_o))
			}
		}
	}
	{
		var _10_arr m_Std_JSON_Grammar.Bracketed = _source0.Get_().(m_Std_JSON_Grammar.Value_Array).Arr
		_ = _10_arr
		var _11_valueOrError3 m_Std_Wrappers.Result = Companion_Default___.Array(_10_arr)
		_ = _11_valueOrError3
		if (_11_valueOrError3).IsFailure() {
			return (_11_valueOrError3).PropagateFailure()
		} else {
			var _12_a _dafny.Sequence = (_11_valueOrError3).Extract().(_dafny.Sequence)
			_ = _12_a
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Values.Companion_JSON_.Create_Array_(_12_a))
		}
	}
}
func (_static *CompanionStruct_Default___) JSON(js m_Std_JSON_Grammar.Structural) m_Std_Wrappers.Result {
	return Companion_Default___.Value((js).Dtor_t().(m_Std_JSON_Grammar.Value))
}
func (_static *CompanionStruct_Default___) HEX__TABLE__16() _dafny.Map {
	return m_Std_JSON_Deserializer_Uint16StrConversion.Companion_Default___.CharToDigit()
}
func (_static *CompanionStruct_Default___) DIGITS() _dafny.Map {
	return m_Std_JSON_ByteStrConversion.Companion_Default___.CharToDigit()
}
func (_static *CompanionStruct_Default___) MINUS() uint8 {
	return uint8(_dafny.CodePoint('-'))
}

// End of class Default__
