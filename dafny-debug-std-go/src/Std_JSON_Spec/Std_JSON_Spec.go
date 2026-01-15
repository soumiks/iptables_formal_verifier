// Package Std_JSON_Spec
// Dafny module Std_JSON_Spec compiled into Go

package Std_JSON_Spec

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
	return "Std_JSON_Spec.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) EscapeUnicode(c uint16) _dafny.Sequence {
	var _0_sStr _dafny.Sequence = m_Std_Strings_HexConversion.Companion_Default___.OfNat(_dafny.IntOfUint16(c))
	_ = _0_sStr
	var _1_s _dafny.Sequence = m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF16(_0_sStr)
	_ = _1_s
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate(((_dafny.IntOfInt64(4)).Minus(_dafny.IntOfUint32((_1_s).Cardinality()))).Uint32(), func(coer34 func(_dafny.Int) uint16) func(_dafny.Int) interface{} {
		return func(arg35 _dafny.Int) interface{} {
			return coer34(arg35)
		}
	}(func(_2___v8 _dafny.Int) uint16 {
		return uint16(_dafny.CodePoint('0'))
	})), _1_s)
}
func (_static *CompanionStruct_Default___) Escape(str _dafny.Sequence, start _dafny.Int) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (start).Cmp(_dafny.IntOfUint32((str).Cardinality())) >= 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, func() _dafny.Sequence {
			var _source0 uint16 = (str).Select((start).Uint32()).(uint16)
			_ = _source0
			{
				if (_source0) == (uint16(34)) {
					return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF16(_dafny.UnicodeSeqOfUtf8Bytes("\\\""))
				}
			}
			{
				if (_source0) == (uint16(92)) {
					return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF16(_dafny.UnicodeSeqOfUtf8Bytes("\\\\"))
				}
			}
			{
				if (_source0) == (uint16(8)) {
					return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF16(_dafny.UnicodeSeqOfUtf8Bytes("\\b"))
				}
			}
			{
				if (_source0) == (uint16(12)) {
					return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF16(_dafny.UnicodeSeqOfUtf8Bytes("\\f"))
				}
			}
			{
				if (_source0) == (uint16(10)) {
					return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF16(_dafny.UnicodeSeqOfUtf8Bytes("\\n"))
				}
			}
			{
				if (_source0) == (uint16(13)) {
					return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF16(_dafny.UnicodeSeqOfUtf8Bytes("\\r"))
				}
			}
			{
				if (_source0) == (uint16(9)) {
					return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF16(_dafny.UnicodeSeqOfUtf8Bytes("\\t"))
				}
			}
			{
				var _1_c uint16 = _source0
				_ = _1_c
				if (_1_c) < (uint16(31)) {
					return _dafny.Companion_Sequence_.Concatenate(m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF16(_dafny.UnicodeSeqOfUtf8Bytes("\\u")), Companion_Default___.EscapeUnicode(_1_c))
				} else {
					return _dafny.SeqOf((str).Select((start).Uint32()).(uint16))
				}
			}
		}())
		var _in0 _dafny.Sequence = str
		_ = _in0
		var _in1 _dafny.Int = (start).Plus(_dafny.One)
		_ = _in1
		str = _in0
		start = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) EscapeToUTF8(str _dafny.Sequence, start _dafny.Int) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = (m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ToUTF16Checked(str)).ToResult(m_Std_JSON_Errors.Companion_SerializationError_.Create_InvalidUnicode_())
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_utf16 _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _1_utf16
		var _2_escaped _dafny.Sequence = Companion_Default___.Escape(_1_utf16, _dafny.Zero)
		_ = _2_escaped
		var _3_valueOrError1 m_Std_Wrappers.Result = ((m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.FromUTF16Checked(_2_escaped)).ToOption()).ToResult(m_Std_JSON_Errors.Companion_SerializationError_.Create_InvalidUnicode_())
		_ = _3_valueOrError1
		if (_3_valueOrError1).IsFailure() {
			return (_3_valueOrError1).PropagateFailure()
		} else {
			var _4_utf32 _dafny.Sequence = (_3_valueOrError1).Extract().(_dafny.Sequence)
			_ = _4_utf32
			return (m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ToUTF8Checked(_4_utf32)).ToResult(m_Std_JSON_Errors.Companion_SerializationError_.Create_InvalidUnicode_())
		}
	}
}
func (_static *CompanionStruct_Default___) String_(str _dafny.Sequence) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.EscapeToUTF8(str, _dafny.Zero)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_inBytes _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _1_inBytes
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("\"")), _1_inBytes), m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("\""))))
	}
}
func (_static *CompanionStruct_Default___) IntToBytes(n _dafny.Int) _dafny.Sequence {
	var _0_s _dafny.Sequence = m_Std_Strings.Companion_Default___.OfInt(n)
	_ = _0_s
	return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_0_s)
}
func (_static *CompanionStruct_Default___) Number(dec m_Std_JSON_Values.Decimal) m_Std_Wrappers.Result {
	return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.IntToBytes((dec).Dtor_n()), (func() _dafny.Sequence {
		if ((dec).Dtor_e10()).Sign() == 0 {
			return _dafny.SeqOf()
		}
		return _dafny.Companion_Sequence_.Concatenate(m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("e")), Companion_Default___.IntToBytes((dec).Dtor_e10()))
	})()))
}
func (_static *CompanionStruct_Default___) KeyValue(kv _dafny.Tuple) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.String_((*(kv).IndexInt(0)).(_dafny.Sequence))
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_key _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _1_key
		var _2_valueOrError1 m_Std_Wrappers.Result = Companion_Default___.JSON((*(kv).IndexInt(1)).(m_Std_JSON_Values.JSON))
		_ = _2_valueOrError1
		if (_2_valueOrError1).IsFailure() {
			return (_2_valueOrError1).PropagateFailure()
		} else {
			var _3_value _dafny.Sequence = (_2_valueOrError1).Extract().(_dafny.Sequence)
			_ = _3_value
			return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1_key, m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes(":"))), _3_value))
		}
	}
}
func (_static *CompanionStruct_Default___) Join(sep _dafny.Sequence, items _dafny.Sequence) m_Std_Wrappers.Result {
	if (_dafny.IntOfUint32((items).Cardinality())).Sign() == 0 {
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.SeqOf())
	} else {
		var _0_valueOrError0 m_Std_Wrappers.Result = (items).Select(0).(m_Std_Wrappers.Result)
		_ = _0_valueOrError0
		if (_0_valueOrError0).IsFailure() {
			return (_0_valueOrError0).PropagateFailure()
		} else {
			var _1_first _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
			_ = _1_first
			if (_dafny.IntOfUint32((items).Cardinality())).Cmp(_dafny.One) == 0 {
				return m_Std_Wrappers.Companion_Result_.Create_Success_(_1_first)
			} else {
				var _2_valueOrError1 m_Std_Wrappers.Result = Companion_Default___.Join(sep, (items).Drop(1))
				_ = _2_valueOrError1
				if (_2_valueOrError1).IsFailure() {
					return (_2_valueOrError1).PropagateFailure()
				} else {
					var _3_rest _dafny.Sequence = (_2_valueOrError1).Extract().(_dafny.Sequence)
					_ = _3_rest
					return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1_first, sep), _3_rest))
				}
			}
		}
	}
}
func (_static *CompanionStruct_Default___) Object(obj _dafny.Sequence) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.Join(m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes(",")), _dafny.SeqCreate((_dafny.IntOfUint32((obj).Cardinality())).Uint32(), func(coer35 func(_dafny.Int) m_Std_Wrappers.Result) func(_dafny.Int) interface{} {
		return func(arg36 _dafny.Int) interface{} {
			return coer35(arg36)
		}
	}((func(_1_obj _dafny.Sequence) func(_dafny.Int) m_Std_Wrappers.Result {
		return func(_2_i _dafny.Int) m_Std_Wrappers.Result {
			return Companion_Default___.KeyValue((_1_obj).Select((_2_i).Uint32()).(_dafny.Tuple))
		}
	})(obj))))
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _3_middle _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _3_middle
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("{")), _3_middle), m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("}"))))
	}
}
func (_static *CompanionStruct_Default___) Array(arr _dafny.Sequence) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.Join(m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes(",")), _dafny.SeqCreate((_dafny.IntOfUint32((arr).Cardinality())).Uint32(), func(coer36 func(_dafny.Int) m_Std_Wrappers.Result) func(_dafny.Int) interface{} {
		return func(arg37 _dafny.Int) interface{} {
			return coer36(arg37)
		}
	}((func(_1_arr _dafny.Sequence) func(_dafny.Int) m_Std_Wrappers.Result {
		return func(_2_i _dafny.Int) m_Std_Wrappers.Result {
			return Companion_Default___.JSON((_1_arr).Select((_2_i).Uint32()).(m_Std_JSON_Values.JSON))
		}
	})(arr))))
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _3_middle _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _3_middle
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("[")), _3_middle), m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("]"))))
	}
}
func (_static *CompanionStruct_Default___) JSON(js m_Std_JSON_Values.JSON) m_Std_Wrappers.Result {
	var _source0 m_Std_JSON_Values.JSON = js
	_ = _source0
	{
		if _source0.Is_Null() {
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("null")))
		}
	}
	{
		if _source0.Is_Bool() {
			var _0_b bool = _source0.Get_().(m_Std_JSON_Values.JSON_Bool).B
			_ = _0_b
			return m_Std_Wrappers.Companion_Result_.Create_Success_((func() _dafny.Sequence {
				if _0_b {
					return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("true"))
				}
				return m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ASCIIToUTF8(_dafny.UnicodeSeqOfUtf8Bytes("false"))
			})())
		}
	}
	{
		if _source0.Is_String() {
			var _1_str _dafny.Sequence = _source0.Get_().(m_Std_JSON_Values.JSON_String).Str
			_ = _1_str
			return Companion_Default___.String_(_1_str)
		}
	}
	{
		if _source0.Is_Number() {
			var _2_dec m_Std_JSON_Values.Decimal = _source0.Get_().(m_Std_JSON_Values.JSON_Number).Num
			_ = _2_dec
			return Companion_Default___.Number(_2_dec)
		}
	}
	{
		if _source0.Is_Object() {
			var _3_obj _dafny.Sequence = _source0.Get_().(m_Std_JSON_Values.JSON_Object).Obj
			_ = _3_obj
			return Companion_Default___.Object(_3_obj)
		}
	}
	{
		var _4_arr _dafny.Sequence = _source0.Get_().(m_Std_JSON_Values.JSON_Array).Arr
		_ = _4_arr
		return Companion_Default___.Array(_4_arr)
	}
}

// End of class Default__
