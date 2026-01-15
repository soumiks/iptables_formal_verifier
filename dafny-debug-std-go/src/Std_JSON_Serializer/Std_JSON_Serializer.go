// Package Std_JSON_Serializer
// Dafny module Std_JSON_Serializer compiled into Go

package Std_JSON_Serializer

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
	m_Std_JSON_Errors "Std_JSON_Errors"
	m_Std_JSON_Grammar "Std_JSON_Grammar"
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
	return "Std_JSON_Serializer.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Bool(b bool) m_Std_JSON_Utils_Views_Core.View__ {
	return m_Std_JSON_Utils_Views_Core.Companion_View___.OfBytes((func() _dafny.Sequence {
		if b {
			return m_Std_JSON_Grammar.Companion_Default___.TRUE()
		}
		return m_Std_JSON_Grammar.Companion_Default___.FALSE()
	})())
}
func (_static *CompanionStruct_Default___) CheckLength(s _dafny.Sequence, err m_Std_JSON_Errors.SerializationError) m_Std_Wrappers.Outcome {
	return m_Std_Wrappers.Companion_Outcome_.Need((_dafny.IntOfUint32((s).Cardinality())).Cmp(m_Std_BoundedInts.Companion_Default___.TWO__TO__THE__32()) < 0, err)
}
func (_static *CompanionStruct_Default___) String_(str _dafny.Sequence) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = m_Std_JSON_Spec.Companion_Default___.EscapeToUTF8(str, _dafny.Zero)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_bs _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _1_bs
		var _2_o m_Std_Wrappers.Outcome = Companion_Default___.CheckLength(_1_bs, m_Std_JSON_Errors.Companion_SerializationError_.Create_StringTooLong_(str))
		_ = _2_o
		if (_2_o).Is_Pass() {
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Jstring_.Create_JString_(m_Std_JSON_Grammar.Companion_Default___.DOUBLEQUOTE(), m_Std_JSON_Utils_Views_Core.Companion_View___.OfBytes(_1_bs), m_Std_JSON_Grammar.Companion_Default___.DOUBLEQUOTE()))
		} else {
			return m_Std_Wrappers.Companion_Result_.Create_Failure_((_2_o).Dtor_error().(m_Std_JSON_Errors.SerializationError))
		}
	}
}
func (_static *CompanionStruct_Default___) Sign(n _dafny.Int) m_Std_JSON_Utils_Views_Core.View__ {
	return m_Std_JSON_Utils_Views_Core.Companion_View___.OfBytes((func() _dafny.Sequence {
		if (n).Sign() == -1 {
			return _dafny.SeqOf(uint8(_dafny.CodePoint('-')))
		}
		return _dafny.SeqOf()
	})())
}
func (_static *CompanionStruct_Default___) Int_k(n _dafny.Int) _dafny.Sequence {
	return m_Std_JSON_ByteStrConversion.Companion_Default___.OfInt(n, Companion_Default___.MINUS())
}
func (_static *CompanionStruct_Default___) Int(n _dafny.Int) m_Std_Wrappers.Result {
	var _0_bs _dafny.Sequence = Companion_Default___.Int_k(n)
	_ = _0_bs
	var _1_o m_Std_Wrappers.Outcome = Companion_Default___.CheckLength(_0_bs, m_Std_JSON_Errors.Companion_SerializationError_.Create_IntTooLarge_(n))
	_ = _1_o
	if (_1_o).Is_Pass() {
		return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Utils_Views_Core.Companion_View___.OfBytes(_0_bs))
	} else {
		return m_Std_Wrappers.Companion_Result_.Create_Failure_((_1_o).Dtor_error().(m_Std_JSON_Errors.SerializationError))
	}
}
func (_static *CompanionStruct_Default___) Number(dec m_Std_JSON_Values.Decimal) m_Std_Wrappers.Result {
	var _pat_let_tv0 = dec
	_ = _pat_let_tv0
	var _pat_let_tv1 = dec
	_ = _pat_let_tv1
	var _0_minus m_Std_JSON_Utils_Views_Core.View__ = Companion_Default___.Sign((dec).Dtor_n())
	_ = _0_minus
	var _1_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.Int(m_Std_Math.Companion_Default___.Abs((dec).Dtor_n()))
	_ = _1_valueOrError0
	if (_1_valueOrError0).IsFailure() {
		return (_1_valueOrError0).PropagateFailure()
	} else {
		var _2_num m_Std_JSON_Utils_Views_Core.View__ = (_1_valueOrError0).Extract().(m_Std_JSON_Utils_Views_Core.View__)
		_ = _2_num
		var _3_frac m_Std_JSON_Grammar.Maybe = m_Std_JSON_Grammar.Companion_Maybe_.Create_Empty_()
		_ = _3_frac
		var _4_valueOrError1 m_Std_Wrappers.Result = (func() m_Std_Wrappers.Result {
			if ((dec).Dtor_e10()).Sign() == 0 {
				return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Maybe_.Create_Empty_())
			}
			return func(_pat_let2_0 m_Std_JSON_Utils_Views_Core.View__) m_Std_Wrappers.Result {
				return func(_5_e m_Std_JSON_Utils_Views_Core.View__) m_Std_Wrappers.Result {
					return func(_pat_let3_0 m_Std_JSON_Utils_Views_Core.View__) m_Std_Wrappers.Result {
						return func(_6_sign m_Std_JSON_Utils_Views_Core.View__) m_Std_Wrappers.Result {
							return func(_pat_let4_0 m_Std_Wrappers.Result) m_Std_Wrappers.Result {
								return func(_7_valueOrError2 m_Std_Wrappers.Result) m_Std_Wrappers.Result {
									return (func() m_Std_Wrappers.Result {
										if (_7_valueOrError2).IsFailure() {
											return (_7_valueOrError2).PropagateFailure()
										}
										return func(_pat_let5_0 m_Std_JSON_Utils_Views_Core.View__) m_Std_Wrappers.Result {
											return func(_8_num m_Std_JSON_Utils_Views_Core.View__) m_Std_Wrappers.Result {
												return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Maybe_.Create_NonEmpty_(m_Std_JSON_Grammar.Companion_Jexp_.Create_JExp_(_5_e, _6_sign, _8_num)))
											}(_pat_let5_0)
										}((_7_valueOrError2).Extract().(m_Std_JSON_Utils_Views_Core.View__))
									})()
								}(_pat_let4_0)
							}(Companion_Default___.Int(m_Std_Math.Companion_Default___.Abs((_pat_let_tv1).Dtor_e10())))
						}(_pat_let3_0)
					}(Companion_Default___.Sign((_pat_let_tv0).Dtor_e10()))
				}(_pat_let2_0)
			}(m_Std_JSON_Utils_Views_Core.Companion_View___.OfBytes(_dafny.SeqOf(uint8(_dafny.CodePoint('e')))))
		})()
		_ = _4_valueOrError1
		if (_4_valueOrError1).IsFailure() {
			return (_4_valueOrError1).PropagateFailure()
		} else {
			var _9_exp m_Std_JSON_Grammar.Maybe = (_4_valueOrError1).Extract().(m_Std_JSON_Grammar.Maybe)
			_ = _9_exp
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Jnumber_.Create_JNumber_(_0_minus, _2_num, m_Std_JSON_Grammar.Companion_Maybe_.Create_Empty_(), _9_exp))
		}
	}
}
func (_static *CompanionStruct_Default___) MkStructural(v interface{}) m_Std_JSON_Grammar.Structural {
	return m_Std_JSON_Grammar.Companion_Structural_.Create_Structural_(m_Std_JSON_Grammar.Companion_Default___.EMPTY(), v, m_Std_JSON_Grammar.Companion_Default___.EMPTY())
}
func (_static *CompanionStruct_Default___) KeyValue(kv _dafny.Tuple) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.String_((*(kv).IndexInt(0)).(_dafny.Sequence))
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_k m_Std_JSON_Grammar.Jstring = (_0_valueOrError0).Extract().(m_Std_JSON_Grammar.Jstring)
		_ = _1_k
		var _2_valueOrError1 m_Std_Wrappers.Result = Companion_Default___.Value((*(kv).IndexInt(1)).(m_Std_JSON_Values.JSON))
		_ = _2_valueOrError1
		if (_2_valueOrError1).IsFailure() {
			return (_2_valueOrError1).PropagateFailure()
		} else {
			var _3_v m_Std_JSON_Grammar.Value = (_2_valueOrError1).Extract().(m_Std_JSON_Grammar.Value)
			_ = _3_v
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_JKeyValue_.Create_KeyValue_(_1_k, Companion_Default___.COLON(), _3_v))
		}
	}
}
func (_static *CompanionStruct_Default___) MkSuffixedSequence(ds _dafny.Sequence, suffix m_Std_JSON_Grammar.Structural, start _dafny.Int) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (start).Cmp(_dafny.IntOfUint32((ds).Cardinality())) >= 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else if (start).Cmp((_dafny.IntOfUint32((ds).Cardinality())).Minus(_dafny.One)) == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(m_Std_JSON_Grammar.Companion_Suffixed_.Create_Suffixed_((ds).Select((start).Uint32()).(interface{}), m_Std_JSON_Grammar.Companion_Maybe_.Create_Empty_())))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(m_Std_JSON_Grammar.Companion_Suffixed_.Create_Suffixed_((ds).Select((start).Uint32()).(interface{}), m_Std_JSON_Grammar.Companion_Maybe_.Create_NonEmpty_(suffix))))
		var _in0 _dafny.Sequence = ds
		_ = _in0
		var _in1 m_Std_JSON_Grammar.Structural = suffix
		_ = _in1
		var _in2 _dafny.Int = (start).Plus(_dafny.One)
		_ = _in2
		ds = _in0
		suffix = _in1
		start = _in2
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Object(obj _dafny.Sequence) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = m_Std_Collections_Seq.Companion_Default___.MapWithResult(func(coer41 func(_dafny.Tuple) m_Std_Wrappers.Result) func(interface{}) m_Std_Wrappers.Result {
		return func(arg42 interface{}) m_Std_Wrappers.Result {
			return coer41(arg42.(_dafny.Tuple))
		}
	}((func(_1_obj _dafny.Sequence) func(_dafny.Tuple) m_Std_Wrappers.Result {
		return func(_2_v _dafny.Tuple) m_Std_Wrappers.Result {
			return Companion_Default___.KeyValue(_2_v)
		}
	})(obj)), obj)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _3_items _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _3_items
		return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Bracketed_.Create_Bracketed_(Companion_Default___.MkStructural(m_Std_JSON_Grammar.Companion_Default___.LBRACE()), Companion_Default___.MkSuffixedSequence(_3_items, Companion_Default___.COMMA(), _dafny.Zero), Companion_Default___.MkStructural(m_Std_JSON_Grammar.Companion_Default___.RBRACE())))
	}
}
func (_static *CompanionStruct_Default___) Array(arr _dafny.Sequence) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = m_Std_Collections_Seq.Companion_Default___.MapWithResult(func(coer42 func(m_Std_JSON_Values.JSON) m_Std_Wrappers.Result) func(interface{}) m_Std_Wrappers.Result {
		return func(arg43 interface{}) m_Std_Wrappers.Result {
			return coer42(arg43.(m_Std_JSON_Values.JSON))
		}
	}((func(_1_arr _dafny.Sequence) func(m_Std_JSON_Values.JSON) m_Std_Wrappers.Result {
		return func(_2_v m_Std_JSON_Values.JSON) m_Std_Wrappers.Result {
			return Companion_Default___.Value(_2_v)
		}
	})(arr)), arr)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _3_items _dafny.Sequence = (_0_valueOrError0).Extract().(_dafny.Sequence)
		_ = _3_items
		return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Bracketed_.Create_Bracketed_(Companion_Default___.MkStructural(m_Std_JSON_Grammar.Companion_Default___.LBRACKET()), Companion_Default___.MkSuffixedSequence(_3_items, Companion_Default___.COMMA(), _dafny.Zero), Companion_Default___.MkStructural(m_Std_JSON_Grammar.Companion_Default___.RBRACKET())))
	}
}
func (_static *CompanionStruct_Default___) Value(js m_Std_JSON_Values.JSON) m_Std_Wrappers.Result {
	var _source0 m_Std_JSON_Values.JSON = js
	_ = _source0
	{
		if _source0.Is_Null() {
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Value_.Create_Null_(m_Std_JSON_Utils_Views_Core.Companion_View___.OfBytes(m_Std_JSON_Grammar.Companion_Default___.NULL())))
		}
	}
	{
		if _source0.Is_Bool() {
			var _0_b bool = _source0.Get_().(m_Std_JSON_Values.JSON_Bool).B
			_ = _0_b
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Value_.Create_Bool_(Companion_Default___.Bool(_0_b)))
		}
	}
	{
		if _source0.Is_String() {
			var _1_str _dafny.Sequence = _source0.Get_().(m_Std_JSON_Values.JSON_String).Str
			_ = _1_str
			var _2_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.String_(_1_str)
			_ = _2_valueOrError0
			if (_2_valueOrError0).IsFailure() {
				return (_2_valueOrError0).PropagateFailure()
			} else {
				var _3_s m_Std_JSON_Grammar.Jstring = (_2_valueOrError0).Extract().(m_Std_JSON_Grammar.Jstring)
				_ = _3_s
				return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Value_.Create_String_(_3_s))
			}
		}
	}
	{
		if _source0.Is_Number() {
			var _4_dec m_Std_JSON_Values.Decimal = _source0.Get_().(m_Std_JSON_Values.JSON_Number).Num
			_ = _4_dec
			var _5_valueOrError1 m_Std_Wrappers.Result = Companion_Default___.Number(_4_dec)
			_ = _5_valueOrError1
			if (_5_valueOrError1).IsFailure() {
				return (_5_valueOrError1).PropagateFailure()
			} else {
				var _6_n m_Std_JSON_Grammar.Jnumber = (_5_valueOrError1).Extract().(m_Std_JSON_Grammar.Jnumber)
				_ = _6_n
				return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Value_.Create_Number_(_6_n))
			}
		}
	}
	{
		if _source0.Is_Object() {
			var _7_obj _dafny.Sequence = _source0.Get_().(m_Std_JSON_Values.JSON_Object).Obj
			_ = _7_obj
			var _8_valueOrError2 m_Std_Wrappers.Result = Companion_Default___.Object(_7_obj)
			_ = _8_valueOrError2
			if (_8_valueOrError2).IsFailure() {
				return (_8_valueOrError2).PropagateFailure()
			} else {
				var _9_o m_Std_JSON_Grammar.Bracketed = (_8_valueOrError2).Extract().(m_Std_JSON_Grammar.Bracketed)
				_ = _9_o
				return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Value_.Create_Object_(_9_o))
			}
		}
	}
	{
		var _10_arr _dafny.Sequence = _source0.Get_().(m_Std_JSON_Values.JSON_Array).Arr
		_ = _10_arr
		var _11_valueOrError3 m_Std_Wrappers.Result = Companion_Default___.Array(_10_arr)
		_ = _11_valueOrError3
		if (_11_valueOrError3).IsFailure() {
			return (_11_valueOrError3).PropagateFailure()
		} else {
			var _12_a m_Std_JSON_Grammar.Bracketed = (_11_valueOrError3).Extract().(m_Std_JSON_Grammar.Bracketed)
			_ = _12_a
			return m_Std_Wrappers.Companion_Result_.Create_Success_(m_Std_JSON_Grammar.Companion_Value_.Create_Array_(_12_a))
		}
	}
}
func (_static *CompanionStruct_Default___) JSON(js m_Std_JSON_Values.JSON) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.Value(js)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_val m_Std_JSON_Grammar.Value = (_0_valueOrError0).Extract().(m_Std_JSON_Grammar.Value)
		_ = _1_val
		return m_Std_Wrappers.Companion_Result_.Create_Success_(Companion_Default___.MkStructural(_1_val))
	}
}
func (_static *CompanionStruct_Default___) DIGITS() _dafny.Sequence {
	return m_Std_JSON_ByteStrConversion.Companion_Default___.Chars()
}
func (_static *CompanionStruct_Default___) MINUS() uint8 {
	return uint8(_dafny.CodePoint('-'))
}
func (_static *CompanionStruct_Default___) COLON() m_Std_JSON_Grammar.Structural {
	return Companion_Default___.MkStructural(m_Std_JSON_Grammar.Companion_Default___.COLON())
}
func (_static *CompanionStruct_Default___) COMMA() m_Std_JSON_Grammar.Structural {
	return Companion_Default___.MkStructural(m_Std_JSON_Grammar.Companion_Default___.COMMA())
}

// End of class Default__

// Definition of class Bytes32
type Bytes32 struct {
}

func New_Bytes32_() *Bytes32 {
	_this := Bytes32{}

	return &_this
}

type CompanionStruct_Bytes32_ struct {
}

var Companion_Bytes32_ = CompanionStruct_Bytes32_{}

func (*Bytes32) String() string {
	return "Std_JSON_Serializer.Bytes32"
}

// End of class Bytes32

func Type_Bytes32_() _dafny.TypeDescriptor {
	return type_Bytes32_{}
}

type type_Bytes32_ struct {
}

func (_this type_Bytes32_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_Bytes32_) String() string {
	return "Std_JSON_Serializer.Bytes32"
}
func (_this *CompanionStruct_Bytes32_) Is_(__source _dafny.Sequence) bool {
	var _0_bs _dafny.Sequence = (__source)
	_ = _0_bs
	return (_dafny.IntOfUint32((_0_bs).Cardinality())).Cmp(m_Std_BoundedInts.Companion_Default___.TWO__TO__THE__32()) < 0
}

// Definition of class String32
type String32 struct {
}

func New_String32_() *String32 {
	_this := String32{}

	return &_this
}

type CompanionStruct_String32_ struct {
}

var Companion_String32_ = CompanionStruct_String32_{}

func (*String32) String() string {
	return "Std_JSON_Serializer.String32"
}

// End of class String32

func Type_String32_() _dafny.TypeDescriptor {
	return type_String32_{}
}

type type_String32_ struct {
}

func (_this type_String32_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_String32_) String() string {
	return "Std_JSON_Serializer.String32"
}
func (_this *CompanionStruct_String32_) Is_(__source _dafny.Sequence) bool {
	var _1_s _dafny.Sequence = (__source)
	_ = _1_s
	return (_dafny.IntOfUint32((_1_s).Cardinality())).Cmp(m_Std_BoundedInts.Companion_Default___.TWO__TO__THE__32()) < 0
}
