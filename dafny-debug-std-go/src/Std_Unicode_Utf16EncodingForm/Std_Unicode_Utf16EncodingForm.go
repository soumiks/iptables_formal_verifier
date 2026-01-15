// Package Std_Unicode_Utf16EncodingForm
// Dafny module Std_Unicode_Utf16EncodingForm compiled into Go

package Std_Unicode_Utf16EncodingForm

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
	m_Std_BoundedInts "Std_BoundedInts"
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
	m_Std_FileIOInternalExterns "Std_FileIOInternalExterns"
	m_Std_Frames "Std_Frames"
	m_Std_Functions "Std_Functions"
	m_Std_GenericActions "Std_GenericActions"
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
	m_Std_Unicode_Utf8EncodingForm "Std_Unicode_Utf8EncodingForm"
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
	return "Std_Unicode_Utf16EncodingForm.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) IsMinimalWellFormedCodeUnitSubsequence(s _dafny.Sequence) bool {
	if (_dafny.IntOfUint32((s).Cardinality())).Cmp(_dafny.One) == 0 {
		return Companion_Default___.IsWellFormedSingleCodeUnitSequence(s)
	} else if (_dafny.IntOfUint32((s).Cardinality())).Cmp(_dafny.IntOfInt64(2)) == 0 {
		var _0_b bool = Companion_Default___.IsWellFormedDoubleCodeUnitSequence(s)
		_ = _0_b
		return _0_b
	} else {
		return false
	}
}
func (_static *CompanionStruct_Default___) IsWellFormedSingleCodeUnitSequence(s _dafny.Sequence) bool {
	var _0_firstWord uint16 = (s).Select(0).(uint16)
	_ = _0_firstWord
	return (((uint16(0)) <= (_0_firstWord)) && ((_0_firstWord) <= (uint16(55295)))) || (((uint16(57344)) <= (_0_firstWord)) && ((_0_firstWord) <= (uint16(65535))))
}
func (_static *CompanionStruct_Default___) IsWellFormedDoubleCodeUnitSequence(s _dafny.Sequence) bool {
	var _0_firstWord uint16 = (s).Select(0).(uint16)
	_ = _0_firstWord
	var _1_secondWord uint16 = (s).Select(1).(uint16)
	_ = _1_secondWord
	return (((uint16(55296)) <= (_0_firstWord)) && ((_0_firstWord) <= (uint16(56319)))) && (((uint16(56320)) <= (_1_secondWord)) && ((_1_secondWord) <= (uint16(57343))))
}
func (_static *CompanionStruct_Default___) SplitPrefixMinimalWellFormedCodeUnitSubsequence(s _dafny.Sequence) m_Std_Wrappers.Option {
	if ((_dafny.IntOfUint32((s).Cardinality())).Cmp(_dafny.One) >= 0) && (Companion_Default___.IsWellFormedSingleCodeUnitSequence((s).Take(1))) {
		return m_Std_Wrappers.Companion_Option_.Create_Some_((s).Take(1))
	} else if ((_dafny.IntOfUint32((s).Cardinality())).Cmp(_dafny.IntOfInt64(2)) >= 0) && (Companion_Default___.IsWellFormedDoubleCodeUnitSequence((s).Take(2))) {
		return m_Std_Wrappers.Companion_Option_.Create_Some_((s).Take(2))
	} else {
		return m_Std_Wrappers.Companion_Option_.Create_None_()
	}
}
func (_static *CompanionStruct_Default___) EncodeScalarValue(v uint32) _dafny.Sequence {
	if (((uint32(0)) <= (v)) && ((v) <= (uint32(55295)))) || (((uint32(57344)) <= (v)) && ((v) <= (uint32(65535)))) {
		return Companion_Default___.EncodeScalarValueSingleWord(v)
	} else {
		return Companion_Default___.EncodeScalarValueDoubleWord(v)
	}
}
func (_static *CompanionStruct_Default___) EncodeScalarValueSingleWord(v uint32) _dafny.Sequence {
	var _0_firstWord uint16 = uint16(v)
	_ = _0_firstWord
	return _dafny.SeqOf(_0_firstWord)
}
func (_static *CompanionStruct_Default___) EncodeScalarValueDoubleWord(v uint32) _dafny.Sequence {
	var _0_x2 uint16 = uint16((v) & (uint32(1023)))
	_ = _0_x2
	var _1_x1 uint8 = uint8(((v) & (uint32(64512))) >> (uint8(10)))
	_ = _1_x1
	var _2_u uint8 = uint8(((v) & (uint32(2031616))) >> (uint8(16)))
	_ = _2_u
	var _3_w uint8 = uint8((((_2_u) - (func() uint8 { return (uint8(1)) })()) & 0x1F))
	_ = _3_w
	var _4_firstWord uint16 = ((uint16(55296)) | ((uint16(_3_w)) << (uint8(6)))) | (uint16(_1_x1))
	_ = _4_firstWord
	var _5_secondWord uint16 = (uint16(56320)) | (uint16(_0_x2))
	_ = _5_secondWord
	return _dafny.SeqOf(_4_firstWord, _5_secondWord)
}
func (_static *CompanionStruct_Default___) DecodeMinimalWellFormedCodeUnitSubsequence(m _dafny.Sequence) uint32 {
	if (_dafny.IntOfUint32((m).Cardinality())).Cmp(_dafny.One) == 0 {
		return Companion_Default___.DecodeMinimalWellFormedCodeUnitSubsequenceSingleWord(m)
	} else {
		return Companion_Default___.DecodeMinimalWellFormedCodeUnitSubsequenceDoubleWord(m)
	}
}
func (_static *CompanionStruct_Default___) DecodeMinimalWellFormedCodeUnitSubsequenceSingleWord(m _dafny.Sequence) uint32 {
	var _0_firstWord uint16 = (m).Select(0).(uint16)
	_ = _0_firstWord
	var _1_x uint16 = _0_firstWord
	_ = _1_x
	return uint32(_1_x)
}
func (_static *CompanionStruct_Default___) DecodeMinimalWellFormedCodeUnitSubsequenceDoubleWord(m _dafny.Sequence) uint32 {
	var _0_firstWord uint16 = (m).Select(0).(uint16)
	_ = _0_firstWord
	var _1_secondWord uint16 = (m).Select(1).(uint16)
	_ = _1_secondWord
	var _2_x2 uint32 = uint32((_1_secondWord) & (uint16(1023)))
	_ = _2_x2
	var _3_x1 uint32 = uint32((_0_firstWord) & (uint16(63)))
	_ = _3_x1
	var _4_w uint32 = uint32(((_0_firstWord) & (uint16(960))) >> (uint8(6)))
	_ = _4_w
	var _5_u uint32 = (((_4_w) + (uint32(1))) & 0xFFFFFF)
	_ = _5_u
	var _6_v uint32 = ((((_5_u) << (uint8(16))) & 0xFFFFFF) | (((_3_x1) << (uint8(10))) & 0xFFFFFF)) | (_2_x2)
	_ = _6_v
	return _6_v
}
func (_static *CompanionStruct_Default___) PartitionCodeUnitSequenceChecked(s _dafny.Sequence) m_Std_Wrappers.Result {
	var maybeParts m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Default(_dafny.EmptySeq)
	_ = maybeParts
	if _dafny.Companion_Sequence_.Equal(s, _dafny.SeqOf()) {
		maybeParts = m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.SeqOf())
		return maybeParts
	}
	var _0_result _dafny.Sequence
	_ = _0_result
	_0_result = _dafny.SeqOf()
	var _1_rest _dafny.Sequence
	_ = _1_rest
	_1_rest = s
	for (_dafny.IntOfUint32((_1_rest).Cardinality())).Sign() == 1 {
		var _2_valueOrError0 m_Std_Wrappers.Result = m_Std_Wrappers.Result{}
		_ = _2_valueOrError0
		_2_valueOrError0 = (Companion_Default___.SplitPrefixMinimalWellFormedCodeUnitSubsequence(_1_rest)).ToResult(_1_rest)
		if (_2_valueOrError0).IsFailure() {
			maybeParts = (_2_valueOrError0).PropagateFailure()
			return maybeParts
		}
		var _3_prefix _dafny.Sequence
		_ = _3_prefix
		_3_prefix = (_2_valueOrError0).Extract().(_dafny.Sequence)
		_0_result = _dafny.Companion_Sequence_.Concatenate(_0_result, _dafny.SeqOf(_3_prefix))
		_1_rest = (_1_rest).Drop((_dafny.IntOfUint32((_3_prefix).Cardinality())).Uint32())
	}
	maybeParts = m_Std_Wrappers.Companion_Result_.Create_Success_(_0_result)
	return maybeParts
	return maybeParts
}
func (_static *CompanionStruct_Default___) PartitionCodeUnitSequence(s _dafny.Sequence) _dafny.Sequence {
	return (Companion_Default___.PartitionCodeUnitSequenceChecked(s)).Extract().(_dafny.Sequence)
}
func (_static *CompanionStruct_Default___) IsWellFormedCodeUnitSequence(s _dafny.Sequence) bool {
	return (Companion_Default___.PartitionCodeUnitSequenceChecked(s)).Is_Success()
}
func (_static *CompanionStruct_Default___) EncodeScalarSequence(vs _dafny.Sequence) _dafny.Sequence {
	var s _dafny.Sequence = Companion_WellFormedCodeUnitSeq_.Witness()
	_ = s
	s = _dafny.SeqOf()
	var _lo0 _dafny.Int = _dafny.Zero
	_ = _lo0
	for _0_i := _dafny.IntOfUint32((vs).Cardinality()); _lo0.Cmp(_0_i) < 0; {
		_0_i = _0_i.Minus(_dafny.One)
		var _1_next _dafny.Sequence
		_ = _1_next
		_1_next = Companion_Default___.EncodeScalarValue((vs).Select((_0_i).Uint32()).(uint32))
		s = _dafny.Companion_Sequence_.Concatenate(_1_next, s)
	}
	return s
}
func (_static *CompanionStruct_Default___) DecodeCodeUnitSequence(s _dafny.Sequence) _dafny.Sequence {
	var _0_parts _dafny.Sequence = Companion_Default___.PartitionCodeUnitSequence(s)
	_ = _0_parts
	var _1_vs _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.MapPartialFunction(func(coer9 func(_dafny.Sequence) uint32) func(interface{}) interface{} {
		return func(arg10 interface{}) interface{} {
			return coer9(arg10.(_dafny.Sequence))
		}
	}(Companion_Default___.DecodeMinimalWellFormedCodeUnitSubsequence), _0_parts)
	_ = _1_vs
	return _1_vs
}
func (_static *CompanionStruct_Default___) DecodeErrorMessage(index _dafny.Int) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("Could not decode byte at index "), m_Std_Strings.Companion_Default___.OfInt(index))
}
func (_static *CompanionStruct_Default___) DecodeCodeUnitSequenceChecked(s _dafny.Sequence) m_Std_Wrappers.Result {
	var resultVs m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Default(_dafny.EmptySeq)
	_ = resultVs
	var _0_maybeParts m_Std_Wrappers.Result
	_ = _0_maybeParts
	_0_maybeParts = Companion_Default___.PartitionCodeUnitSequenceChecked(s)
	if (_0_maybeParts).Is_Failure() {
		resultVs = m_Std_Wrappers.Companion_Result_.Create_Failure_(Companion_Default___.DecodeErrorMessage((_dafny.IntOfUint32((s).Cardinality())).Minus(_dafny.IntOfUint32(((_0_maybeParts).Dtor_error().(_dafny.Sequence)).Cardinality()))))
		return resultVs
	}
	var _1_parts _dafny.Sequence
	_ = _1_parts
	_1_parts = (_0_maybeParts).Dtor_value().(_dafny.Sequence)
	var _2_vs _dafny.Sequence
	_ = _2_vs
	_2_vs = m_Std_Collections_Seq.Companion_Default___.MapPartialFunction(func(coer10 func(_dafny.Sequence) uint32) func(interface{}) interface{} {
		return func(arg11 interface{}) interface{} {
			return coer10(arg11.(_dafny.Sequence))
		}
	}(Companion_Default___.DecodeMinimalWellFormedCodeUnitSubsequence), _1_parts)
	resultVs = m_Std_Wrappers.Companion_Result_.Create_Success_(_2_vs)
	return resultVs
	return resultVs
}

// End of class Default__

// Definition of class WellFormedCodeUnitSeq
type WellFormedCodeUnitSeq struct {
}

func New_WellFormedCodeUnitSeq_() *WellFormedCodeUnitSeq {
	_this := WellFormedCodeUnitSeq{}

	return &_this
}

type CompanionStruct_WellFormedCodeUnitSeq_ struct {
}

var Companion_WellFormedCodeUnitSeq_ = CompanionStruct_WellFormedCodeUnitSeq_{}

func (*WellFormedCodeUnitSeq) String() string {
	return "Std_Unicode_Utf16EncodingForm.WellFormedCodeUnitSeq"
}
func (_this *CompanionStruct_WellFormedCodeUnitSeq_) Witness() _dafny.Sequence {
	return _dafny.SeqOf()
}

// End of class WellFormedCodeUnitSeq

func Type_WellFormedCodeUnitSeq_() _dafny.TypeDescriptor {
	return type_WellFormedCodeUnitSeq_{}
}

type type_WellFormedCodeUnitSeq_ struct {
}

func (_this type_WellFormedCodeUnitSeq_) Default() interface{} {
	return Companion_WellFormedCodeUnitSeq_.Witness()
}

func (_this type_WellFormedCodeUnitSeq_) String() string {
	return "Std_Unicode_Utf16EncodingForm.WellFormedCodeUnitSeq"
}
func (_this *CompanionStruct_WellFormedCodeUnitSeq_) Is_(__source _dafny.Sequence) bool {
	var _3_s _dafny.Sequence = (__source)
	_ = _3_s
	return Companion_Default___.IsWellFormedCodeUnitSequence(_3_s)
}

// Definition of class MinimalWellFormedCodeUnitSeq
type MinimalWellFormedCodeUnitSeq struct {
}

func New_MinimalWellFormedCodeUnitSeq_() *MinimalWellFormedCodeUnitSeq {
	_this := MinimalWellFormedCodeUnitSeq{}

	return &_this
}

type CompanionStruct_MinimalWellFormedCodeUnitSeq_ struct {
}

var Companion_MinimalWellFormedCodeUnitSeq_ = CompanionStruct_MinimalWellFormedCodeUnitSeq_{}

func (*MinimalWellFormedCodeUnitSeq) String() string {
	return "Std_Unicode_Utf16EncodingForm.MinimalWellFormedCodeUnitSeq"
}

// End of class MinimalWellFormedCodeUnitSeq

func Type_MinimalWellFormedCodeUnitSeq_() _dafny.TypeDescriptor {
	return type_MinimalWellFormedCodeUnitSeq_{}
}

type type_MinimalWellFormedCodeUnitSeq_ struct {
}

func (_this type_MinimalWellFormedCodeUnitSeq_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_MinimalWellFormedCodeUnitSeq_) String() string {
	return "Std_Unicode_Utf16EncodingForm.MinimalWellFormedCodeUnitSeq"
}
func (_this *CompanionStruct_MinimalWellFormedCodeUnitSeq_) Is_(__source _dafny.Sequence) bool {
	var _4_s _dafny.Sequence = (__source)
	_ = _4_s
	return Companion_Default___.IsMinimalWellFormedCodeUnitSubsequence(_4_s)
}
