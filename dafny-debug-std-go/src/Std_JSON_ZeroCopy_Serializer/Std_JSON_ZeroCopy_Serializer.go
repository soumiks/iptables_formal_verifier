// Package Std_JSON_ZeroCopy_Serializer
// Dafny module Std_JSON_ZeroCopy_Serializer compiled into Go

package Std_JSON_ZeroCopy_Serializer

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
	return "Std_JSON_ZeroCopy_Serializer.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Serialize(js m_Std_JSON_Grammar.Structural) m_Std_Wrappers.Result {
	var rbs m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Default(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
	_ = rbs
	var _0_writer m_Std_JSON_Utils_Views_Writers.Writer__
	_ = _0_writer
	_0_writer = Companion_Default___.Text(js)
	var _1_valueOrError0 m_Std_Wrappers.OutcomeResult = m_Std_Wrappers.Companion_OutcomeResult_.Default()
	_ = _1_valueOrError0
	_1_valueOrError0 = m_Std_Wrappers.Companion_Default___.Need((_0_writer).Unsaturated_q(), m_Std_JSON_Errors.Companion_SerializationError_.Create_OutOfMemory_())
	if (_1_valueOrError0).IsFailure() {
		rbs = (_1_valueOrError0).PropagateFailure()
		return rbs
	}
	var _2_bs _dafny.Array
	_ = _2_bs
	var _out0 _dafny.Array
	_ = _out0
	_out0 = (_0_writer).ToArray()
	_2_bs = _out0
	rbs = m_Std_Wrappers.Companion_Result_.Create_Success_(_2_bs)
	return rbs
	return rbs
}
func (_static *CompanionStruct_Default___) SerializeTo(js m_Std_JSON_Grammar.Structural, dest _dafny.Array) m_Std_Wrappers.Result {
	var len_ m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Default(uint32(0))
	_ = len_
	var _0_writer m_Std_JSON_Utils_Views_Writers.Writer__
	_ = _0_writer
	_0_writer = Companion_Default___.Text(js)
	var _1_valueOrError0 m_Std_Wrappers.OutcomeResult = m_Std_Wrappers.Companion_OutcomeResult_.Default()
	_ = _1_valueOrError0
	_1_valueOrError0 = m_Std_Wrappers.Companion_Default___.Need((_0_writer).Unsaturated_q(), m_Std_JSON_Errors.Companion_SerializationError_.Create_OutOfMemory_())
	if (_1_valueOrError0).IsFailure() {
		len_ = (_1_valueOrError0).PropagateFailure()
		return len_
	}
	var _2_valueOrError1 m_Std_Wrappers.OutcomeResult = m_Std_Wrappers.Companion_OutcomeResult_.Default()
	_ = _2_valueOrError1
	_2_valueOrError1 = m_Std_Wrappers.Companion_Default___.Need((_dafny.IntOfUint32((_0_writer).Dtor_length())).Cmp(_dafny.ArrayLen((dest), 0)) <= 0, m_Std_JSON_Errors.Companion_SerializationError_.Create_OutOfMemory_())
	if (_2_valueOrError1).IsFailure() {
		len_ = (_2_valueOrError1).PropagateFailure()
		return len_
	}
	(_0_writer).CopyTo(dest)
	len_ = m_Std_Wrappers.Companion_Result_.Create_Success_((_0_writer).Dtor_length())
	return len_
	return len_
}
func (_static *CompanionStruct_Default___) Text(js m_Std_JSON_Grammar.Structural) m_Std_JSON_Utils_Views_Writers.Writer__ {
	return Companion_Default___.JSON(js, m_Std_JSON_Utils_Views_Writers.Companion_Writer___.Empty())
}
func (_static *CompanionStruct_Default___) JSON(js m_Std_JSON_Grammar.Structural, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	return (((writer).Append((js).Dtor_before())).Then((func(_0_js m_Std_JSON_Grammar.Structural) func(m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
		return func(_1_wr m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
			return Companion_Default___.Value((_0_js).Dtor_t().(m_Std_JSON_Grammar.Value), _1_wr)
		}
	})(js))).Append((js).Dtor_after())
}
func (_static *CompanionStruct_Default___) Value(v m_Std_JSON_Grammar.Value, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var _source0 m_Std_JSON_Grammar.Value = v
	_ = _source0
	{
		if _source0.Is_Null() {
			var _0_n m_Std_JSON_Utils_Views_Core.View__ = _source0.Get_().(m_Std_JSON_Grammar.Value_Null).N
			_ = _0_n
			var _1_wr m_Std_JSON_Utils_Views_Writers.Writer__ = (writer).Append(_0_n)
			_ = _1_wr
			return _1_wr
		}
	}
	{
		if _source0.Is_Bool() {
			var _2_b m_Std_JSON_Utils_Views_Core.View__ = _source0.Get_().(m_Std_JSON_Grammar.Value_Bool).B
			_ = _2_b
			var _3_wr m_Std_JSON_Utils_Views_Writers.Writer__ = (writer).Append(_2_b)
			_ = _3_wr
			return _3_wr
		}
	}
	{
		if _source0.Is_String() {
			var _4_str m_Std_JSON_Grammar.Jstring = _source0.Get_().(m_Std_JSON_Grammar.Value_String).Str
			_ = _4_str
			var _5_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.String_(_4_str, writer)
			_ = _5_wr
			return _5_wr
		}
	}
	{
		if _source0.Is_Number() {
			var _6_num m_Std_JSON_Grammar.Jnumber = _source0.Get_().(m_Std_JSON_Grammar.Value_Number).Num
			_ = _6_num
			var _7_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.Number(_6_num, writer)
			_ = _7_wr
			return _7_wr
		}
	}
	{
		if _source0.Is_Object() {
			var _8_obj m_Std_JSON_Grammar.Bracketed = _source0.Get_().(m_Std_JSON_Grammar.Value_Object).Obj
			_ = _8_obj
			var _9_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.Object(_8_obj, writer)
			_ = _9_wr
			return _9_wr
		}
	}
	{
		var _10_arr m_Std_JSON_Grammar.Bracketed = _source0.Get_().(m_Std_JSON_Grammar.Value_Array).Arr
		_ = _10_arr
		var _11_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.Array(_10_arr, writer)
		_ = _11_wr
		return _11_wr
	}
}
func (_static *CompanionStruct_Default___) String_(str m_Std_JSON_Grammar.Jstring, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	return (((writer).Append((str).Dtor_lq())).Append((str).Dtor_contents())).Append((str).Dtor_rq())
}
func (_static *CompanionStruct_Default___) Number(num m_Std_JSON_Grammar.Jnumber, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var _0_wr1 m_Std_JSON_Utils_Views_Writers.Writer__ = ((writer).Append((num).Dtor_minus())).Append((num).Dtor_num())
	_ = _0_wr1
	var _1_wr2 m_Std_JSON_Utils_Views_Writers.Writer__ = (func() m_Std_JSON_Utils_Views_Writers.Writer__ {
		if ((num).Dtor_frac()).Is_NonEmpty() {
			return ((_0_wr1).Append((((num).Dtor_frac()).Dtor_t().(m_Std_JSON_Grammar.Jfrac)).Dtor_period())).Append((((num).Dtor_frac()).Dtor_t().(m_Std_JSON_Grammar.Jfrac)).Dtor_num())
		}
		return _0_wr1
	})()
	_ = _1_wr2
	var _2_wr3 m_Std_JSON_Utils_Views_Writers.Writer__ = (func() m_Std_JSON_Utils_Views_Writers.Writer__ {
		if ((num).Dtor_exp()).Is_NonEmpty() {
			return (((_1_wr2).Append((((num).Dtor_exp()).Dtor_t().(m_Std_JSON_Grammar.Jexp)).Dtor_e())).Append((((num).Dtor_exp()).Dtor_t().(m_Std_JSON_Grammar.Jexp)).Dtor_sign())).Append((((num).Dtor_exp()).Dtor_t().(m_Std_JSON_Grammar.Jexp)).Dtor_num())
		}
		return _1_wr2
	})()
	_ = _2_wr3
	var _3_wr m_Std_JSON_Utils_Views_Writers.Writer__ = _2_wr3
	_ = _3_wr
	return _3_wr
}
func (_static *CompanionStruct_Default___) StructuralView(st m_Std_JSON_Grammar.Structural, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	return (((writer).Append((st).Dtor_before())).Append((st).Dtor_t().(m_Std_JSON_Utils_Views_Core.View__))).Append((st).Dtor_after())
}
func (_static *CompanionStruct_Default___) Object(obj m_Std_JSON_Grammar.Bracketed, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var _0_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.StructuralView((obj).Dtor_l(), writer)
	_ = _0_wr
	var _1_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.Members(obj, _0_wr)
	_ = _1_wr
	var _2_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.StructuralView((obj).Dtor_r(), _1_wr)
	_ = _2_wr
	return _2_wr
}
func (_static *CompanionStruct_Default___) Array(arr m_Std_JSON_Grammar.Bracketed, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var _0_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.StructuralView((arr).Dtor_l(), writer)
	_ = _0_wr
	var _1_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.Items(arr, _0_wr)
	_ = _1_wr
	var _2_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.StructuralView((arr).Dtor_r(), _1_wr)
	_ = _2_wr
	return _2_wr
}
func (_static *CompanionStruct_Default___) Members(obj m_Std_JSON_Grammar.Bracketed, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var wr m_Std_JSON_Utils_Views_Writers.Writer__ = m_Std_JSON_Utils_Views_Writers.Companion_Writer_.Witness()
	_ = wr
	var _out0 m_Std_JSON_Utils_Views_Writers.Writer__
	_ = _out0
	_out0 = Companion_Default___.MembersImpl(obj, writer)
	wr = _out0
	return wr
}
func (_static *CompanionStruct_Default___) Items(arr m_Std_JSON_Grammar.Bracketed, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var wr m_Std_JSON_Utils_Views_Writers.Writer__ = m_Std_JSON_Utils_Views_Writers.Companion_Writer_.Witness()
	_ = wr
	var _out0 m_Std_JSON_Utils_Views_Writers.Writer__
	_ = _out0
	_out0 = Companion_Default___.ItemsImpl(arr, writer)
	wr = _out0
	return wr
}
func (_static *CompanionStruct_Default___) MembersImpl(obj m_Std_JSON_Grammar.Bracketed, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var wr m_Std_JSON_Utils_Views_Writers.Writer__ = m_Std_JSON_Utils_Views_Writers.Companion_Writer_.Witness()
	_ = wr
	wr = writer
	var _0_members _dafny.Sequence
	_ = _0_members
	_0_members = (obj).Dtor_data()
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_0_members).Cardinality())
	_ = _hi0
	for _1_i := _dafny.Zero; _1_i.Cmp(_hi0) < 0; _1_i = _1_i.Plus(_dafny.One) {
		wr = Companion_Default___.Member((_0_members).Select((_1_i).Uint32()).(m_Std_JSON_Grammar.Suffixed), wr)
	}
	return wr
}
func (_static *CompanionStruct_Default___) ItemsImpl(arr m_Std_JSON_Grammar.Bracketed, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var wr m_Std_JSON_Utils_Views_Writers.Writer__ = m_Std_JSON_Utils_Views_Writers.Companion_Writer_.Witness()
	_ = wr
	wr = writer
	var _0_items _dafny.Sequence
	_ = _0_items
	_0_items = (arr).Dtor_data()
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_0_items).Cardinality())
	_ = _hi0
	for _1_i := _dafny.Zero; _1_i.Cmp(_hi0) < 0; _1_i = _1_i.Plus(_dafny.One) {
		wr = Companion_Default___.Item((_0_items).Select((_1_i).Uint32()).(m_Std_JSON_Grammar.Suffixed), wr)
	}
	return wr
}
func (_static *CompanionStruct_Default___) Member(m m_Std_JSON_Grammar.Suffixed, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var _0_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.String_(((m).Dtor_t().(m_Std_JSON_Grammar.JKeyValue)).Dtor_k(), writer)
	_ = _0_wr
	var _1_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.StructuralView(((m).Dtor_t().(m_Std_JSON_Grammar.JKeyValue)).Dtor_colon(), _0_wr)
	_ = _1_wr
	var _2_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.Value(((m).Dtor_t().(m_Std_JSON_Grammar.JKeyValue)).Dtor_v(), _1_wr)
	_ = _2_wr
	var _3_wr m_Std_JSON_Utils_Views_Writers.Writer__ = (func() m_Std_JSON_Utils_Views_Writers.Writer__ {
		if ((m).Dtor_suffix()).Is_Empty() {
			return _2_wr
		}
		return Companion_Default___.StructuralView(((m).Dtor_suffix()).Dtor_t().(m_Std_JSON_Grammar.Structural), _2_wr)
	})()
	_ = _3_wr
	return _3_wr
}
func (_static *CompanionStruct_Default___) Item(m m_Std_JSON_Grammar.Suffixed, writer m_Std_JSON_Utils_Views_Writers.Writer__) m_Std_JSON_Utils_Views_Writers.Writer__ {
	var _0_wr m_Std_JSON_Utils_Views_Writers.Writer__ = Companion_Default___.Value((m).Dtor_t().(m_Std_JSON_Grammar.Value), writer)
	_ = _0_wr
	var _1_wr m_Std_JSON_Utils_Views_Writers.Writer__ = (func() m_Std_JSON_Utils_Views_Writers.Writer__ {
		if ((m).Dtor_suffix()).Is_Empty() {
			return _0_wr
		}
		return Companion_Default___.StructuralView(((m).Dtor_suffix()).Dtor_t().(m_Std_JSON_Grammar.Structural), _0_wr)
	})()
	_ = _1_wr
	return _1_wr
}

// End of class Default__
