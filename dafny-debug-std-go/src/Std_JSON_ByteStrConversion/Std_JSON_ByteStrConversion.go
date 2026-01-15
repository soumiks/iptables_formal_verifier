// Package Std_JSON_ByteStrConversion
// Dafny module Std_JSON_ByteStrConversion compiled into Go

package Std_JSON_ByteStrConversion

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
	return "Std_JSON_ByteStrConversion.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) BASE() _dafny.Int {
	return Companion_Default___.Base()
}
func (_static *CompanionStruct_Default___) IsDigitChar(c uint8) bool {
	return (Companion_Default___.CharToDigit()).Contains(c)
}
func (_static *CompanionStruct_Default___) OfDigits(digits _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if _dafny.Companion_Sequence_.Equal(digits, _dafny.SeqOf()) {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(), _0___accumulator)
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((Companion_Default___.Chars()).Select(((digits).Select(0).(_dafny.Int)).Uint32()).(uint8)), _0___accumulator)
		var _in0 _dafny.Sequence = (digits).Drop(1)
		_ = _in0
		digits = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) OfNat(n _dafny.Int) _dafny.Sequence {
	if (n).Sign() == 0 {
		return _dafny.SeqOf((Companion_Default___.Chars()).Select(0).(uint8))
	} else {
		return Companion_Default___.OfDigits(Companion_Default___.FromNat(n))
	}
}
func (_static *CompanionStruct_Default___) IsNumberStr(str _dafny.Sequence, minus uint8) bool {
	return !(!_dafny.Companion_Sequence_.Equal(str, _dafny.SeqOf())) || (((((str).Select(0).(uint8)) == (minus)) || ((Companion_Default___.CharToDigit()).Contains((str).Select(0).(uint8)))) && (_dafny.Quantifier(((str).Drop(1)).UniqueElements(), true, func(_forall_var_0 uint8) bool {
		var _0_c uint8
		_0_c = interface{}(_forall_var_0).(uint8)
		if true {
			return !(_dafny.Companion_Sequence_.Contains((str).Drop(1), _0_c)) || (Companion_Default___.IsDigitChar(_0_c))
		} else {
			return true
		}
	})))
}
func (_static *CompanionStruct_Default___) OfInt(n _dafny.Int, minus uint8) _dafny.Sequence {
	if (n).Sign() != -1 {
		return Companion_Default___.OfNat(n)
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(minus), Companion_Default___.OfNat((_dafny.Zero).Minus(n)))
	}
}
func (_static *CompanionStruct_Default___) ToNat(str _dafny.Sequence) _dafny.Int {
	if _dafny.Companion_Sequence_.Equal(str, _dafny.SeqOf()) {
		return _dafny.Zero
	} else {
		var _0_c uint8 = (str).Select(((_dafny.IntOfUint32((str).Cardinality())).Minus(_dafny.One)).Uint32()).(uint8)
		_ = _0_c
		return ((Companion_Default___.ToNat((str).Take(((_dafny.IntOfUint32((str).Cardinality())).Minus(_dafny.One)).Uint32()))).Times(Companion_Default___.Base())).Plus((Companion_Default___.CharToDigit()).Get(_0_c).(_dafny.Int))
	}
}
func (_static *CompanionStruct_Default___) ToInt(str _dafny.Sequence, minus uint8) _dafny.Int {
	if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(minus), str) {
		return (_dafny.Zero).Minus(Companion_Default___.ToNat((str).Drop(1)))
	} else {
		return Companion_Default___.ToNat(str)
	}
}
func (_static *CompanionStruct_Default___) ToNatRight(xs _dafny.Sequence) _dafny.Int {
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Zero
	} else {
		return ((Companion_Default___.ToNatRight(m_Std_Collections_Seq.Companion_Default___.DropFirst(xs))).Times(Companion_Default___.BASE())).Plus(m_Std_Collections_Seq.Companion_Default___.First(xs).(_dafny.Int))
	}
}
func (_static *CompanionStruct_Default___) ToNatLeft(xs _dafny.Sequence) _dafny.Int {
	var _0___accumulator _dafny.Int = _dafny.Zero
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return (_dafny.Zero).Plus(_0___accumulator)
	} else {
		_0___accumulator = ((m_Std_Collections_Seq.Companion_Default___.Last(xs).(_dafny.Int)).Times(m_Std_Arithmetic_Power.Companion_Default___.Pow(Companion_Default___.BASE(), (_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)))).Plus(_0___accumulator)
		var _in0 _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.DropLast(xs)
		_ = _in0
		xs = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) FromNat(n _dafny.Int) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (n).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((n).Modulo(Companion_Default___.BASE())))
		var _in0 _dafny.Int = (n).DivBy(Companion_Default___.BASE())
		_ = _in0
		n = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) SeqExtend(xs _dafny.Sequence, n _dafny.Int) _dafny.Sequence {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Cmp(n) >= 0 {
		return xs
	} else {
		var _in0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(xs, _dafny.SeqOf(_dafny.Zero))
		_ = _in0
		var _in1 _dafny.Int = n
		_ = _in1
		xs = _in0
		n = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) SeqExtendMultiple(xs _dafny.Sequence, n _dafny.Int) _dafny.Sequence {
	var _0_newLen _dafny.Int = ((_dafny.IntOfUint32((xs).Cardinality())).Plus(n)).Minus((_dafny.IntOfUint32((xs).Cardinality())).Modulo(n))
	_ = _0_newLen
	return Companion_Default___.SeqExtend(xs, _0_newLen)
}
func (_static *CompanionStruct_Default___) FromNatWithLen(n _dafny.Int, len_ _dafny.Int) _dafny.Sequence {
	return Companion_Default___.SeqExtend(Companion_Default___.FromNat(n), len_)
}
func (_static *CompanionStruct_Default___) SeqZero(len_ _dafny.Int) _dafny.Sequence {
	var _0_xs _dafny.Sequence = Companion_Default___.FromNatWithLen(_dafny.Zero, len_)
	_ = _0_xs
	return _0_xs
}
func (_static *CompanionStruct_Default___) SeqAdd(xs _dafny.Sequence, ys _dafny.Sequence) _dafny.Tuple {
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.TupleOf(_dafny.SeqOf(), _dafny.Zero)
	} else {
		var _let_tmp_rhs0 _dafny.Tuple = Companion_Default___.SeqAdd(m_Std_Collections_Seq.Companion_Default___.DropLast(xs), m_Std_Collections_Seq.Companion_Default___.DropLast(ys))
		_ = _let_tmp_rhs0
		var _0_zs_k _dafny.Sequence = (*(_let_tmp_rhs0).IndexInt(0)).(_dafny.Sequence)
		_ = _0_zs_k
		var _1_cin _dafny.Int = (*(_let_tmp_rhs0).IndexInt(1)).(_dafny.Int)
		_ = _1_cin
		var _2_sum _dafny.Int = ((m_Std_Collections_Seq.Companion_Default___.Last(xs).(_dafny.Int)).Plus(m_Std_Collections_Seq.Companion_Default___.Last(ys).(_dafny.Int))).Plus(_1_cin)
		_ = _2_sum
		var _let_tmp_rhs1 _dafny.Tuple = (func() _dafny.Tuple {
			if (_2_sum).Cmp(Companion_Default___.BASE()) < 0 {
				return _dafny.TupleOf(_2_sum, _dafny.Zero)
			}
			return _dafny.TupleOf((_2_sum).Minus(Companion_Default___.BASE()), _dafny.One)
		})()
		_ = _let_tmp_rhs1
		var _3_sum__out _dafny.Int = (*(_let_tmp_rhs1).IndexInt(0)).(_dafny.Int)
		_ = _3_sum__out
		var _4_cout _dafny.Int = (*(_let_tmp_rhs1).IndexInt(1)).(_dafny.Int)
		_ = _4_cout
		return _dafny.TupleOf(_dafny.Companion_Sequence_.Concatenate(_0_zs_k, _dafny.SeqOf(_3_sum__out)), _4_cout)
	}
}
func (_static *CompanionStruct_Default___) SeqSub(xs _dafny.Sequence, ys _dafny.Sequence) _dafny.Tuple {
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.TupleOf(_dafny.SeqOf(), _dafny.Zero)
	} else {
		var _let_tmp_rhs0 _dafny.Tuple = Companion_Default___.SeqSub(m_Std_Collections_Seq.Companion_Default___.DropLast(xs), m_Std_Collections_Seq.Companion_Default___.DropLast(ys))
		_ = _let_tmp_rhs0
		var _0_zs _dafny.Sequence = (*(_let_tmp_rhs0).IndexInt(0)).(_dafny.Sequence)
		_ = _0_zs
		var _1_cin _dafny.Int = (*(_let_tmp_rhs0).IndexInt(1)).(_dafny.Int)
		_ = _1_cin
		var _let_tmp_rhs1 _dafny.Tuple = (func() _dafny.Tuple {
			if (m_Std_Collections_Seq.Companion_Default___.Last(xs).(_dafny.Int)).Cmp((m_Std_Collections_Seq.Companion_Default___.Last(ys).(_dafny.Int)).Plus(_1_cin)) >= 0 {
				return _dafny.TupleOf(((m_Std_Collections_Seq.Companion_Default___.Last(xs).(_dafny.Int)).Minus(m_Std_Collections_Seq.Companion_Default___.Last(ys).(_dafny.Int))).Minus(_1_cin), _dafny.Zero)
			}
			return _dafny.TupleOf((((Companion_Default___.BASE()).Plus(m_Std_Collections_Seq.Companion_Default___.Last(xs).(_dafny.Int))).Minus(m_Std_Collections_Seq.Companion_Default___.Last(ys).(_dafny.Int))).Minus(_1_cin), _dafny.One)
		})()
		_ = _let_tmp_rhs1
		var _2_diff__out _dafny.Int = (*(_let_tmp_rhs1).IndexInt(0)).(_dafny.Int)
		_ = _2_diff__out
		var _3_cout _dafny.Int = (*(_let_tmp_rhs1).IndexInt(1)).(_dafny.Int)
		_ = _3_cout
		return _dafny.TupleOf(_dafny.Companion_Sequence_.Concatenate(_0_zs, _dafny.SeqOf(_2_diff__out)), _3_cout)
	}
}
func (_static *CompanionStruct_Default___) Chars() _dafny.Sequence {
	return _dafny.SeqOf(uint8(_dafny.CodePoint('0')), uint8(_dafny.CodePoint('1')), uint8(_dafny.CodePoint('2')), uint8(_dafny.CodePoint('3')), uint8(_dafny.CodePoint('4')), uint8(_dafny.CodePoint('5')), uint8(_dafny.CodePoint('6')), uint8(_dafny.CodePoint('7')), uint8(_dafny.CodePoint('8')), uint8(_dafny.CodePoint('9')))
}
func (_static *CompanionStruct_Default___) Base() _dafny.Int {
	return _dafny.IntOfUint32((Companion_Default___.Chars()).Cardinality())
}
func (_static *CompanionStruct_Default___) CharToDigit() _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(uint8(_dafny.CodePoint('0')), _dafny.Zero).UpdateUnsafe(uint8(_dafny.CodePoint('1')), _dafny.One).UpdateUnsafe(uint8(_dafny.CodePoint('2')), _dafny.IntOfInt64(2)).UpdateUnsafe(uint8(_dafny.CodePoint('3')), _dafny.IntOfInt64(3)).UpdateUnsafe(uint8(_dafny.CodePoint('4')), _dafny.IntOfInt64(4)).UpdateUnsafe(uint8(_dafny.CodePoint('5')), _dafny.IntOfInt64(5)).UpdateUnsafe(uint8(_dafny.CodePoint('6')), _dafny.IntOfInt64(6)).UpdateUnsafe(uint8(_dafny.CodePoint('7')), _dafny.IntOfInt64(7)).UpdateUnsafe(uint8(_dafny.CodePoint('8')), _dafny.IntOfInt64(8)).UpdateUnsafe(uint8(_dafny.CodePoint('9')), _dafny.IntOfInt64(9))
}

// End of class Default__

// Definition of class CharSeq
type CharSeq struct {
}

func New_CharSeq_() *CharSeq {
	_this := CharSeq{}

	return &_this
}

type CompanionStruct_CharSeq_ struct {
}

var Companion_CharSeq_ = CompanionStruct_CharSeq_{}

func (*CharSeq) String() string {
	return "Std_JSON_ByteStrConversion.CharSeq"
}

// End of class CharSeq

func Type_CharSeq_() _dafny.TypeDescriptor {
	return type_CharSeq_{}
}

type type_CharSeq_ struct {
}

func (_this type_CharSeq_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_CharSeq_) String() string {
	return "Std_JSON_ByteStrConversion.CharSeq"
}
func (_this *CompanionStruct_CharSeq_) Is_(__source _dafny.Sequence) bool {
	var _0_chars _dafny.Sequence = (__source)
	_ = _0_chars
	return (_dafny.IntOfUint32((_0_chars).Cardinality())).Cmp(_dafny.One) > 0
}

// Definition of class Digit
type Digit struct {
}

func New_Digit_() *Digit {
	_this := Digit{}

	return &_this
}

type CompanionStruct_Digit_ struct {
}

var Companion_Digit_ = CompanionStruct_Digit_{}

func (*Digit) String() string {
	return "Std_JSON_ByteStrConversion.Digit"
}

// End of class Digit

func Type_Digit_() _dafny.TypeDescriptor {
	return type_Digit_{}
}

type type_Digit_ struct {
}

func (_this type_Digit_) Default() interface{} {
	return _dafny.Zero
}

func (_this type_Digit_) String() string {
	return "Std_JSON_ByteStrConversion.Digit"
}
func (_this *CompanionStruct_Digit_) Is_(__source _dafny.Int) bool {
	var _1_i _dafny.Int = (__source)
	_ = _1_i
	if m__System.Companion_Nat_.Is_(_1_i) {
		return ((_1_i).Sign() != -1) && ((_1_i).Cmp(Companion_Default___.BASE()) < 0)
	}
	return false
}
