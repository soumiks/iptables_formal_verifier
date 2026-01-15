// Package Std_JSON_ConcreteSyntax_Spec
// Dafny module Std_JSON_ConcreteSyntax_Spec compiled into Go

package Std_JSON_ConcreteSyntax_Spec

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
	return "Std_JSON_ConcreteSyntax_Spec.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) View(v m_Std_JSON_Utils_Views_Core.View__) _dafny.Sequence {
	return (v).Bytes()
}
func (_static *CompanionStruct_Default___) Structural(self m_Std_JSON_Grammar.Structural, fT func(interface{}) _dafny.Sequence) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.View((self).Dtor_before()), (fT)((self).Dtor_t())), Companion_Default___.View((self).Dtor_after()))
}
func (_static *CompanionStruct_Default___) StructuralView(self m_Std_JSON_Grammar.Structural) _dafny.Sequence {
	return Companion_Default___.Structural(self, func(coer47 func(m_Std_JSON_Utils_Views_Core.View__) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg48 interface{}) _dafny.Sequence {
			return coer47(arg48.(m_Std_JSON_Utils_Views_Core.View__))
		}
	}(Companion_Default___.View))
}
func (_static *CompanionStruct_Default___) Maybe(self m_Std_JSON_Grammar.Maybe, fT func(interface{}) _dafny.Sequence) _dafny.Sequence {
	if (self).Is_Empty() {
		return _dafny.SeqOf()
	} else {
		return (fT)((self).Dtor_t())
	}
}
func (_static *CompanionStruct_Default___) ConcatBytes(ts _dafny.Sequence, fT func(interface{}) _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((ts).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (fT)((ts).Select(0).(interface{})))
		var _in0 _dafny.Sequence = (ts).Drop(1)
		_ = _in0
		var _in1 func(interface{}) _dafny.Sequence = fT
		_ = _in1
		ts = _in0
		fT = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Bracketed(self m_Std_JSON_Grammar.Bracketed, fDatum func(m_Std_JSON_Grammar.Suffixed) _dafny.Sequence) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.StructuralView((self).Dtor_l()), Companion_Default___.ConcatBytes((self).Dtor_data(), func(coer48 func(m_Std_JSON_Grammar.Suffixed) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg49 interface{}) _dafny.Sequence {
			return coer48(arg49.(m_Std_JSON_Grammar.Suffixed))
		}
	}(fDatum))), Companion_Default___.StructuralView((self).Dtor_r()))
}
func (_static *CompanionStruct_Default___) KeyValue(self m_Std_JSON_Grammar.JKeyValue) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.String_((self).Dtor_k()), Companion_Default___.StructuralView((self).Dtor_colon())), Companion_Default___.Value((self).Dtor_v()))
}
func (_static *CompanionStruct_Default___) Frac(self m_Std_JSON_Grammar.Jfrac) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.View((self).Dtor_period()), Companion_Default___.View((self).Dtor_num()))
}
func (_static *CompanionStruct_Default___) Exp(self m_Std_JSON_Grammar.Jexp) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.View((self).Dtor_e()), Companion_Default___.View((self).Dtor_sign())), Companion_Default___.View((self).Dtor_num()))
}
func (_static *CompanionStruct_Default___) Number(self m_Std_JSON_Grammar.Jnumber) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.View((self).Dtor_minus()), Companion_Default___.View((self).Dtor_num())), Companion_Default___.Maybe((self).Dtor_frac(), func(coer49 func(m_Std_JSON_Grammar.Jfrac) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg50 interface{}) _dafny.Sequence {
			return coer49(arg50.(m_Std_JSON_Grammar.Jfrac))
		}
	}(Companion_Default___.Frac))), Companion_Default___.Maybe((self).Dtor_exp(), func(coer50 func(m_Std_JSON_Grammar.Jexp) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg51 interface{}) _dafny.Sequence {
			return coer50(arg51.(m_Std_JSON_Grammar.Jexp))
		}
	}(Companion_Default___.Exp)))
}
func (_static *CompanionStruct_Default___) String_(self m_Std_JSON_Grammar.Jstring) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.View((self).Dtor_lq()), Companion_Default___.View((self).Dtor_contents())), Companion_Default___.View((self).Dtor_rq()))
}
func (_static *CompanionStruct_Default___) CommaSuffix(c m_Std_JSON_Grammar.Maybe) _dafny.Sequence {
	return Companion_Default___.Maybe(c, func(coer51 func(m_Std_JSON_Grammar.Structural) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg52 interface{}) _dafny.Sequence {
			return coer51(arg52.(m_Std_JSON_Grammar.Structural))
		}
	}(Companion_Default___.StructuralView))
}
func (_static *CompanionStruct_Default___) Member(self m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.KeyValue((self).Dtor_t().(m_Std_JSON_Grammar.JKeyValue)), Companion_Default___.CommaSuffix((self).Dtor_suffix()))
}
func (_static *CompanionStruct_Default___) Item(self m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Value((self).Dtor_t().(m_Std_JSON_Grammar.Value)), Companion_Default___.CommaSuffix((self).Dtor_suffix()))
}
func (_static *CompanionStruct_Default___) Object(obj m_Std_JSON_Grammar.Bracketed) _dafny.Sequence {
	return Companion_Default___.Bracketed(obj, func(coer52 func(m_Std_JSON_Grammar.Suffixed) _dafny.Sequence) func(m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
		return func(arg53 m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
			return coer52(arg53)
		}
	}((func(_0_obj m_Std_JSON_Grammar.Bracketed) func(m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
		return func(_1_d m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
			return Companion_Default___.Member(_1_d)
		}
	})(obj)))
}
func (_static *CompanionStruct_Default___) Array(arr m_Std_JSON_Grammar.Bracketed) _dafny.Sequence {
	return Companion_Default___.Bracketed(arr, func(coer53 func(m_Std_JSON_Grammar.Suffixed) _dafny.Sequence) func(m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
		return func(arg54 m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
			return coer53(arg54)
		}
	}((func(_0_arr m_Std_JSON_Grammar.Bracketed) func(m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
		return func(_1_d m_Std_JSON_Grammar.Suffixed) _dafny.Sequence {
			return Companion_Default___.Item(_1_d)
		}
	})(arr)))
}
func (_static *CompanionStruct_Default___) Value(self m_Std_JSON_Grammar.Value) _dafny.Sequence {
	var _source0 m_Std_JSON_Grammar.Value = self
	_ = _source0
	{
		if _source0.Is_Null() {
			var _0_n m_Std_JSON_Utils_Views_Core.View__ = _source0.Get_().(m_Std_JSON_Grammar.Value_Null).N
			_ = _0_n
			return Companion_Default___.View(_0_n)
		}
	}
	{
		if _source0.Is_Bool() {
			var _1_b m_Std_JSON_Utils_Views_Core.View__ = _source0.Get_().(m_Std_JSON_Grammar.Value_Bool).B
			_ = _1_b
			return Companion_Default___.View(_1_b)
		}
	}
	{
		if _source0.Is_String() {
			var _2_str m_Std_JSON_Grammar.Jstring = _source0.Get_().(m_Std_JSON_Grammar.Value_String).Str
			_ = _2_str
			return Companion_Default___.String_(_2_str)
		}
	}
	{
		if _source0.Is_Number() {
			var _3_num m_Std_JSON_Grammar.Jnumber = _source0.Get_().(m_Std_JSON_Grammar.Value_Number).Num
			_ = _3_num
			return Companion_Default___.Number(_3_num)
		}
	}
	{
		if _source0.Is_Object() {
			var _4_obj m_Std_JSON_Grammar.Bracketed = _source0.Get_().(m_Std_JSON_Grammar.Value_Object).Obj
			_ = _4_obj
			return Companion_Default___.Object(_4_obj)
		}
	}
	{
		var _5_arr m_Std_JSON_Grammar.Bracketed = _source0.Get_().(m_Std_JSON_Grammar.Value_Array).Arr
		_ = _5_arr
		return Companion_Default___.Array(_5_arr)
	}
}
func (_static *CompanionStruct_Default___) JSON(js m_Std_JSON_Grammar.Structural) _dafny.Sequence {
	return Companion_Default___.Structural(js, func(coer54 func(m_Std_JSON_Grammar.Value) _dafny.Sequence) func(interface{}) _dafny.Sequence {
		return func(arg55 interface{}) _dafny.Sequence {
			return coer54(arg55.(m_Std_JSON_Grammar.Value))
		}
	}(Companion_Default___.Value))
}

// End of class Default__
