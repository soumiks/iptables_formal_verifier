// Package Std_Base64
// Dafny module Std_Base64 compiled into Go

package Std_Base64

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
	return "Std_Base64.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) IsBase64Char(c _dafny.CodePoint) bool {
	return (((((c) == (_dafny.CodePoint('+'))) || ((c) == (_dafny.CodePoint('/')))) || (((_dafny.CodePoint('0')) <= (c)) && ((c) <= (_dafny.CodePoint('9'))))) || (((_dafny.CodePoint('A')) <= (c)) && ((c) <= (_dafny.CodePoint('Z'))))) || (((_dafny.CodePoint('a')) <= (c)) && ((c) <= (_dafny.CodePoint('z'))))
}
func (_static *CompanionStruct_Default___) IsUnpaddedBase64String(s _dafny.Sequence) bool {
	return (((_dafny.IntOfUint32((s).Cardinality())).Modulo(_dafny.IntOfInt64(4))).Sign() == 0) && (_dafny.Quantifier((s).UniqueElements(), true, func(_forall_var_0 _dafny.CodePoint) bool {
		var _0_k _dafny.CodePoint
		_0_k = interface{}(_forall_var_0).(_dafny.CodePoint)
		return !(_dafny.Companion_Sequence_.Contains(s, _0_k)) || (Companion_Default___.IsBase64Char(_0_k))
	}))
}
func (_static *CompanionStruct_Default___) IndexToChar(i uint8) _dafny.CodePoint {
	if (i) == (uint8(63)) {
		return _dafny.CodePoint('/')
	} else if (i) == (uint8(62)) {
		return _dafny.CodePoint('+')
	} else if ((uint8(52)) <= (i)) && ((i) <= (uint8(61))) {
		return _dafny.CodePoint((_dafny.IntOfUint8((((i) - (func() uint8 { return (uint8(4)) })()) & 0x3F))).Int32())
	} else if ((uint8(26)) <= (i)) && ((i) <= (uint8(51))) {
		return (_dafny.CodePoint((_dafny.IntOfUint8(i)).Int32())) + (_dafny.CodePoint((_dafny.IntOfInt64(71)).Int32()))
	} else {
		return (_dafny.CodePoint((_dafny.IntOfUint8(i)).Int32())) + (_dafny.CodePoint((_dafny.IntOfInt64(65)).Int32()))
	}
}
func (_static *CompanionStruct_Default___) CharToIndex(c _dafny.CodePoint) uint8 {
	if (c) == (_dafny.CodePoint('/')) {
		return uint8(63)
	} else if (c) == (_dafny.CodePoint('+')) {
		return uint8(62)
	} else if ((_dafny.CodePoint('0')) <= (c)) && ((c) <= (_dafny.CodePoint('9'))) {
		return (_dafny.IntOfInt32(rune((c) + (_dafny.CodePoint((_dafny.IntOfInt64(4)).Int32()))))).Uint8()
	} else if ((_dafny.CodePoint('a')) <= (c)) && ((c) <= (_dafny.CodePoint('z'))) {
		return (_dafny.IntOfInt32(rune((c) - (_dafny.CodePoint((_dafny.IntOfInt64(71)).Int32()))))).Uint8()
	} else {
		return (_dafny.IntOfInt32(rune((c) - (_dafny.CodePoint((_dafny.IntOfInt64(65)).Int32()))))).Uint8()
	}
}
func (_static *CompanionStruct_Default___) BV24ToSeq(x uint32) _dafny.Sequence {
	var _0_b0 uint8 = uint8(((x) >> (uint8(16))) & (uint32(255)))
	_ = _0_b0
	var _1_b1 uint8 = uint8(((x) >> (uint8(8))) & (uint32(255)))
	_ = _1_b1
	var _2_b2 uint8 = uint8((x) & (uint32(255)))
	_ = _2_b2
	return _dafny.SeqOf(_0_b0, _1_b1, _2_b2)
}
func (_static *CompanionStruct_Default___) SeqToBV24(x _dafny.Sequence) uint32 {
	return ((((uint32((x).Select(0).(uint8))) << (uint8(16))) & 0xFFFFFF) | (((uint32((x).Select(1).(uint8))) << (uint8(8))) & 0xFFFFFF)) | (uint32((x).Select(2).(uint8)))
}
func (_static *CompanionStruct_Default___) BV24ToIndexSeq(x uint32) _dafny.Sequence {
	var _0_b0 uint8 = uint8(((x) >> (uint8(18))) & (uint32(63)))
	_ = _0_b0
	var _1_b1 uint8 = uint8(((x) >> (uint8(12))) & (uint32(63)))
	_ = _1_b1
	var _2_b2 uint8 = uint8(((x) >> (uint8(6))) & (uint32(63)))
	_ = _2_b2
	var _3_b3 uint8 = uint8((x) & (uint32(63)))
	_ = _3_b3
	return _dafny.SeqOf(_0_b0, _1_b1, _2_b2, _3_b3)
}
func (_static *CompanionStruct_Default___) IndexSeqToBV24(x _dafny.Sequence) uint32 {
	return (((((uint32((x).Select(0).(uint8))) << (uint8(18))) & 0xFFFFFF) | (((uint32((x).Select(1).(uint8))) << (uint8(12))) & 0xFFFFFF)) | (((uint32((x).Select(2).(uint8))) << (uint8(6))) & 0xFFFFFF)) | (uint32((x).Select(3).(uint8)))
}
func (_static *CompanionStruct_Default___) DecodeBlock(s _dafny.Sequence) _dafny.Sequence {
	return Companion_Default___.BV24ToSeq(Companion_Default___.IndexSeqToBV24(s))
}
func (_static *CompanionStruct_Default___) EncodeBlock(s _dafny.Sequence) _dafny.Sequence {
	return Companion_Default___.BV24ToIndexSeq(Companion_Default___.SeqToBV24(s))
}
func (_static *CompanionStruct_Default___) DecodeRecursively(s _dafny.Sequence) _dafny.Sequence {
	var b _dafny.Sequence = _dafny.EmptySeq
	_ = b
	var _0_resultLength _dafny.Int
	_ = _0_resultLength
	_0_resultLength = ((_dafny.IntOfUint32((s).Cardinality())).DivBy(_dafny.IntOfInt64(4))).Times(_dafny.IntOfInt64(3))
	var _1_result _dafny.Array
	_ = _1_result
	var _len0_0 _dafny.Int = _0_resultLength
	_ = _len0_0
	var _nw0 _dafny.Array
	_ = _nw0
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw0 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) uint8 = func(_2_i _dafny.Int) uint8 {
			return uint8(0)
		}
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw0).ArraySet1Byte(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw0).ArraySet1Byte(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_1_result = _nw0
	var _3_i _dafny.Int
	_ = _3_i
	_3_i = _dafny.IntOfUint32((s).Cardinality())
	var _4_j _dafny.Int
	_ = _4_j
	_4_j = _0_resultLength
	for (_3_i).Sign() == 1 {
		_3_i = (_3_i).Minus(_dafny.IntOfInt64(4))
		_4_j = (_4_j).Minus(_dafny.IntOfInt64(3))
		var _5_block _dafny.Sequence
		_ = _5_block
		_5_block = Companion_Default___.DecodeBlock((s).Subsequence((_3_i).Uint32(), ((_3_i).Plus(_dafny.IntOfInt64(4))).Uint32()))
		(_1_result).ArraySet1Byte((_5_block).Select(0).(uint8), (_4_j).Int())
		var _index0 _dafny.Int = (_4_j).Plus(_dafny.One)
		_ = _index0
		(_1_result).ArraySet1Byte((_5_block).Select(1).(uint8), (_index0).Int())
		var _index1 _dafny.Int = (_4_j).Plus(_dafny.IntOfInt64(2))
		_ = _index1
		(_1_result).ArraySet1Byte((_5_block).Select(2).(uint8), (_index1).Int())
	}
	b = _dafny.ArrayRangeToSeq(_1_result, _dafny.NilInt, _dafny.NilInt)
	return b
}
func (_static *CompanionStruct_Default___) EncodeRecursively(b _dafny.Sequence) _dafny.Sequence {
	var s _dafny.Sequence = _dafny.EmptySeq
	_ = s
	var _0_resultLength _dafny.Int
	_ = _0_resultLength
	_0_resultLength = ((_dafny.IntOfUint32((b).Cardinality())).DivBy(_dafny.IntOfInt64(3))).Times(_dafny.IntOfInt64(4))
	var _1_result _dafny.Array
	_ = _1_result
	var _len0_0 _dafny.Int = _0_resultLength
	_ = _len0_0
	var _nw0 _dafny.Array
	_ = _nw0
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw0 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) uint8 = func(_2_i _dafny.Int) uint8 {
			return uint8(0)
		}
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw0).ArraySet1Byte(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw0).ArraySet1Byte(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_1_result = _nw0
	var _3_i _dafny.Int
	_ = _3_i
	_3_i = _dafny.IntOfUint32((b).Cardinality())
	var _4_j _dafny.Int
	_ = _4_j
	_4_j = _0_resultLength
	for (_3_i).Sign() == 1 {
		_3_i = (_3_i).Minus(_dafny.IntOfInt64(3))
		_4_j = (_4_j).Minus(_dafny.IntOfInt64(4))
		var _5_block _dafny.Sequence
		_ = _5_block
		_5_block = Companion_Default___.EncodeBlock((b).Subsequence((_3_i).Uint32(), ((_3_i).Plus(_dafny.IntOfInt64(3))).Uint32()))
		(_1_result).ArraySet1Byte((_5_block).Select(0).(uint8), (_4_j).Int())
		var _index0 _dafny.Int = (_4_j).Plus(_dafny.One)
		_ = _index0
		(_1_result).ArraySet1Byte((_5_block).Select(1).(uint8), (_index0).Int())
		var _index1 _dafny.Int = (_4_j).Plus(_dafny.IntOfInt64(2))
		_ = _index1
		(_1_result).ArraySet1Byte((_5_block).Select(2).(uint8), (_index1).Int())
		var _index2 _dafny.Int = (_4_j).Plus(_dafny.IntOfInt64(3))
		_ = _index2
		(_1_result).ArraySet1Byte((_5_block).Select(3).(uint8), (_index2).Int())
	}
	s = _dafny.ArrayRangeToSeq(_1_result, _dafny.NilInt, _dafny.NilInt)
	return s
}
func (_static *CompanionStruct_Default___) FromCharsToIndices(s _dafny.Sequence) _dafny.Sequence {
	return _dafny.SeqCreate((_dafny.IntOfUint32((s).Cardinality())).Uint32(), func(coer29 func(_dafny.Int) uint8) func(_dafny.Int) interface{} {
		return func(arg30 _dafny.Int) interface{} {
			return coer29(arg30)
		}
	}((func(_0_s _dafny.Sequence) func(_dafny.Int) uint8 {
		return func(_1_i _dafny.Int) uint8 {
			return Companion_Default___.CharToIndex((_0_s).Select((_1_i).Uint32()).(_dafny.CodePoint))
		}
	})(s)))
}
func (_static *CompanionStruct_Default___) FromIndicesToChars(b _dafny.Sequence) _dafny.Sequence {
	return _dafny.SeqCreate((_dafny.IntOfUint32((b).Cardinality())).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg31 _dafny.Int) interface{} {
			return coer30(arg31)
		}
	}((func(_0_b _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
		return func(_1_i _dafny.Int) _dafny.CodePoint {
			return Companion_Default___.IndexToChar((_0_b).Select((_1_i).Uint32()).(uint8))
		}
	})(b)))
}
func (_static *CompanionStruct_Default___) DecodeUnpadded(s _dafny.Sequence) _dafny.Sequence {
	return Companion_Default___.DecodeRecursively(Companion_Default___.FromCharsToIndices(s))
}
func (_static *CompanionStruct_Default___) EncodeUnpadded(b _dafny.Sequence) _dafny.Sequence {
	return Companion_Default___.FromIndicesToChars(Companion_Default___.EncodeRecursively(b))
}
func (_static *CompanionStruct_Default___) Is1Padding(s _dafny.Sequence) bool {
	return ((((((_dafny.IntOfUint32((s).Cardinality())).Cmp(_dafny.IntOfInt64(4)) == 0) && (Companion_Default___.IsBase64Char((s).Select(0).(_dafny.CodePoint)))) && (Companion_Default___.IsBase64Char((s).Select(1).(_dafny.CodePoint)))) && (Companion_Default___.IsBase64Char((s).Select(2).(_dafny.CodePoint)))) && (((Companion_Default___.CharToIndex((s).Select(2).(_dafny.CodePoint))) & (uint8(3))) == (uint8(0)))) && (((s).Select(3).(_dafny.CodePoint)) == (_dafny.CodePoint('=')))
}
func (_static *CompanionStruct_Default___) Decode1Padding(s _dafny.Sequence) _dafny.Sequence {
	var _0_d _dafny.Sequence = Companion_Default___.DecodeBlock(_dafny.SeqOf(Companion_Default___.CharToIndex((s).Select(0).(_dafny.CodePoint)), Companion_Default___.CharToIndex((s).Select(1).(_dafny.CodePoint)), Companion_Default___.CharToIndex((s).Select(2).(_dafny.CodePoint)), uint8(0)))
	_ = _0_d
	return _dafny.SeqOf((_0_d).Select(0).(uint8), (_0_d).Select(1).(uint8))
}
func (_static *CompanionStruct_Default___) Encode1Padding(b _dafny.Sequence) _dafny.Sequence {
	var _0_e _dafny.Sequence = Companion_Default___.EncodeBlock(_dafny.SeqOf((b).Select(0).(uint8), (b).Select(1).(uint8), uint8(0)))
	_ = _0_e
	return _dafny.SeqOf(Companion_Default___.IndexToChar((_0_e).Select(0).(uint8)), Companion_Default___.IndexToChar((_0_e).Select(1).(uint8)), Companion_Default___.IndexToChar((_0_e).Select(2).(uint8)), _dafny.CodePoint('='))
}
func (_static *CompanionStruct_Default___) Is2Padding(s _dafny.Sequence) bool {
	return ((((((_dafny.IntOfUint32((s).Cardinality())).Cmp(_dafny.IntOfInt64(4)) == 0) && (Companion_Default___.IsBase64Char((s).Select(0).(_dafny.CodePoint)))) && (Companion_Default___.IsBase64Char((s).Select(1).(_dafny.CodePoint)))) && (((Companion_Default___.CharToIndex((s).Select(1).(_dafny.CodePoint))) % (uint8(16))) == (uint8(0)))) && (((s).Select(2).(_dafny.CodePoint)) == (_dafny.CodePoint('=')))) && (((s).Select(3).(_dafny.CodePoint)) == (_dafny.CodePoint('=')))
}
func (_static *CompanionStruct_Default___) Decode2Padding(s _dafny.Sequence) _dafny.Sequence {
	var _0_d _dafny.Sequence = Companion_Default___.DecodeBlock(_dafny.SeqOf(Companion_Default___.CharToIndex((s).Select(0).(_dafny.CodePoint)), Companion_Default___.CharToIndex((s).Select(1).(_dafny.CodePoint)), uint8(0), uint8(0)))
	_ = _0_d
	return _dafny.SeqOf((_0_d).Select(0).(uint8))
}
func (_static *CompanionStruct_Default___) Encode2Padding(b _dafny.Sequence) _dafny.Sequence {
	var _0_e _dafny.Sequence = Companion_Default___.EncodeBlock(_dafny.SeqOf((b).Select(0).(uint8), uint8(0), uint8(0)))
	_ = _0_e
	return _dafny.SeqOf(Companion_Default___.IndexToChar((_0_e).Select(0).(uint8)), Companion_Default___.IndexToChar((_0_e).Select(1).(uint8)), _dafny.CodePoint('='), _dafny.CodePoint('='))
}
func (_static *CompanionStruct_Default___) IsBase64String(s _dafny.Sequence) bool {
	var _0_finalBlockStart _dafny.Int = (_dafny.IntOfUint32((s).Cardinality())).Minus(_dafny.IntOfInt64(4))
	_ = _0_finalBlockStart
	return (((_dafny.IntOfUint32((s).Cardinality())).Modulo(_dafny.IntOfInt64(4))).Sign() == 0) && ((Companion_Default___.IsUnpaddedBase64String(s)) || ((Companion_Default___.IsUnpaddedBase64String((s).Take((_0_finalBlockStart).Uint32()))) && ((Companion_Default___.Is1Padding((s).Drop((_0_finalBlockStart).Uint32()))) || (Companion_Default___.Is2Padding((s).Drop((_0_finalBlockStart).Uint32()))))))
}
func (_static *CompanionStruct_Default___) DecodeValid(s _dafny.Sequence) _dafny.Sequence {
	if _dafny.Companion_Sequence_.Equal(s, _dafny.SeqOf()) {
		return _dafny.SeqOf()
	} else {
		var _0_finalBlockStart _dafny.Int = (_dafny.IntOfUint32((s).Cardinality())).Minus(_dafny.IntOfInt64(4))
		_ = _0_finalBlockStart
		var _1_prefix _dafny.Sequence = (s).Take((_0_finalBlockStart).Uint32())
		_ = _1_prefix
		var _2_suffix _dafny.Sequence = (s).Drop((_0_finalBlockStart).Uint32())
		_ = _2_suffix
		if Companion_Default___.Is1Padding(_2_suffix) {
			return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.DecodeUnpadded(_1_prefix), Companion_Default___.Decode1Padding(_2_suffix))
		} else if Companion_Default___.Is2Padding(_2_suffix) {
			return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.DecodeUnpadded(_1_prefix), Companion_Default___.Decode2Padding(_2_suffix))
		} else {
			return Companion_Default___.DecodeUnpadded(s)
		}
	}
}
func (_static *CompanionStruct_Default___) DecodeBV(s _dafny.Sequence) m_Std_Wrappers.Result {
	if Companion_Default___.IsBase64String(s) {
		return m_Std_Wrappers.Companion_Result_.Create_Success_(Companion_Default___.DecodeValid(s))
	} else {
		return m_Std_Wrappers.Companion_Result_.Create_Failure_(_dafny.UnicodeSeqOfUtf8Bytes("The encoding is malformed"))
	}
}
func (_static *CompanionStruct_Default___) EncodeBV(b _dafny.Sequence) _dafny.Sequence {
	if ((_dafny.IntOfUint32((b).Cardinality())).Modulo(_dafny.IntOfInt64(3))).Sign() == 0 {
		return Companion_Default___.EncodeUnpadded(b)
	} else if ((_dafny.IntOfUint32((b).Cardinality())).Modulo(_dafny.IntOfInt64(3))).Cmp(_dafny.One) == 0 {
		var _0_s1 _dafny.Sequence = Companion_Default___.EncodeUnpadded((b).Take(((_dafny.IntOfUint32((b).Cardinality())).Minus(_dafny.One)).Uint32()))
		_ = _0_s1
		var _1_s2 _dafny.Sequence = Companion_Default___.Encode2Padding((b).Drop(((_dafny.IntOfUint32((b).Cardinality())).Minus(_dafny.One)).Uint32()))
		_ = _1_s2
		return _dafny.Companion_Sequence_.Concatenate(_0_s1, _1_s2)
	} else {
		var _2_s1 _dafny.Sequence = Companion_Default___.EncodeUnpadded((b).Take(((_dafny.IntOfUint32((b).Cardinality())).Minus(_dafny.IntOfInt64(2))).Uint32()))
		_ = _2_s1
		var _3_s2 _dafny.Sequence = Companion_Default___.Encode1Padding((b).Drop(((_dafny.IntOfUint32((b).Cardinality())).Minus(_dafny.IntOfInt64(2))).Uint32()))
		_ = _3_s2
		return _dafny.Companion_Sequence_.Concatenate(_2_s1, _3_s2)
	}
}
func (_static *CompanionStruct_Default___) UInt8sToBVs(u _dafny.Sequence) _dafny.Sequence {
	return _dafny.SeqCreate((_dafny.IntOfUint32((u).Cardinality())).Uint32(), func(coer31 func(_dafny.Int) uint8) func(_dafny.Int) interface{} {
		return func(arg32 _dafny.Int) interface{} {
			return coer31(arg32)
		}
	}((func(_0_u _dafny.Sequence) func(_dafny.Int) uint8 {
		return func(_1_i _dafny.Int) uint8 {
			return uint8((_0_u).Select((_1_i).Uint32()).(uint8))
		}
	})(u)))
}
func (_static *CompanionStruct_Default___) BVsToUInt8s(b _dafny.Sequence) _dafny.Sequence {
	return _dafny.SeqCreate((_dafny.IntOfUint32((b).Cardinality())).Uint32(), func(coer32 func(_dafny.Int) uint8) func(_dafny.Int) interface{} {
		return func(arg33 _dafny.Int) interface{} {
			return coer32(arg33)
		}
	}((func(_0_b _dafny.Sequence) func(_dafny.Int) uint8 {
		return func(_1_i _dafny.Int) uint8 {
			return uint8((_0_b).Select((_1_i).Uint32()).(uint8))
		}
	})(b)))
}
func (_static *CompanionStruct_Default___) Encode(u _dafny.Sequence) _dafny.Sequence {
	return Companion_Default___.EncodeBV(Companion_Default___.UInt8sToBVs(u))
}
func (_static *CompanionStruct_Default___) Decode(s _dafny.Sequence) m_Std_Wrappers.Result {
	if Companion_Default___.IsBase64String(s) {
		var _0_b _dafny.Sequence = Companion_Default___.DecodeValid(s)
		_ = _0_b
		return m_Std_Wrappers.Companion_Result_.Create_Success_(Companion_Default___.BVsToUInt8s(_0_b))
	} else {
		return m_Std_Wrappers.Companion_Result_.Create_Failure_(_dafny.UnicodeSeqOfUtf8Bytes("The encoding is malformed"))
	}
}

// End of class Default__
