// Package Std_Unicode_UnicodeStringsWithUnicodeChar
// Dafny module Std_Unicode_UnicodeStringsWithUnicodeChar compiled into Go

package Std_Unicode_UnicodeStringsWithUnicodeChar

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
	m_Std_Unicode_Utf16EncodingForm "Std_Unicode_Utf16EncodingForm"
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
var _ m_Std_Unicode_Utf16EncodingForm.Dummy__

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
	return "Std_Unicode_UnicodeStringsWithUnicodeChar.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) CharAsUnicodeScalarValue(c _dafny.CodePoint) uint32 {
	return (_dafny.IntOfInt32(rune(c))).Uint32()
}
func (_static *CompanionStruct_Default___) CharFromUnicodeScalarValue(sv uint32) _dafny.CodePoint {
	return _dafny.CodePoint((_dafny.IntOfUint32(sv)).Int32())
}
func (_static *CompanionStruct_Default___) ToUTF8Checked(s _dafny.Sequence) m_Std_Wrappers.Option {
	var _0_asCodeUnits _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.Map(func(coer11 func(_dafny.CodePoint) uint32) func(interface{}) interface{} {
		return func(arg12 interface{}) interface{} {
			return coer11(arg12.(_dafny.CodePoint))
		}
	}(Companion_Default___.CharAsUnicodeScalarValue), s)
	_ = _0_asCodeUnits
	var _1_asUtf8CodeUnits _dafny.Sequence = m_Std_Unicode_Utf8EncodingForm.Companion_Default___.EncodeScalarSequence(_0_asCodeUnits)
	_ = _1_asUtf8CodeUnits
	var _2_asBytes _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.Map(func(coer12 func(uint8) uint8) func(interface{}) interface{} {
		return func(arg13 interface{}) interface{} {
			return coer12(arg13.(uint8))
		}
	}(func(_3_cu uint8) uint8 {
		return uint8(_3_cu)
	}), _1_asUtf8CodeUnits)
	_ = _2_asBytes
	return m_Std_Wrappers.Companion_Option_.Create_Some_(_2_asBytes)
}
func (_static *CompanionStruct_Default___) FromUTF8Checked(bs _dafny.Sequence) m_Std_Wrappers.Result {
	var _0_asCodeUnits _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.Map(func(coer13 func(uint8) uint8) func(interface{}) interface{} {
		return func(arg14 interface{}) interface{} {
			return coer13(arg14.(uint8))
		}
	}(func(_1_c uint8) uint8 {
		return uint8(_1_c)
	}), bs)
	_ = _0_asCodeUnits
	var _2_valueOrError0 m_Std_Wrappers.Result = m_Std_Unicode_Utf8EncodingForm.Companion_Default___.DecodeCodeUnitSequenceChecked(_0_asCodeUnits)
	_ = _2_valueOrError0
	if (_2_valueOrError0).IsFailure() {
		return (_2_valueOrError0).PropagateFailure()
	} else {
		var _3_utf32 _dafny.Sequence = (_2_valueOrError0).Extract().(_dafny.Sequence)
		_ = _3_utf32
		var _4_asChars _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.Map(func(coer14 func(uint32) _dafny.CodePoint) func(interface{}) interface{} {
			return func(arg15 interface{}) interface{} {
				return coer14(arg15.(uint32))
			}
		}(Companion_Default___.CharFromUnicodeScalarValue), _3_utf32)
		_ = _4_asChars
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_4_asChars)
	}
}
func (_static *CompanionStruct_Default___) ToUTF16Checked(s _dafny.Sequence) m_Std_Wrappers.Option {
	var _0_asCodeUnits _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.Map(func(coer15 func(_dafny.CodePoint) uint32) func(interface{}) interface{} {
		return func(arg16 interface{}) interface{} {
			return coer15(arg16.(_dafny.CodePoint))
		}
	}(Companion_Default___.CharAsUnicodeScalarValue), s)
	_ = _0_asCodeUnits
	var _1_asUtf16CodeUnits _dafny.Sequence = m_Std_Unicode_Utf16EncodingForm.Companion_Default___.EncodeScalarSequence(_0_asCodeUnits)
	_ = _1_asUtf16CodeUnits
	var _2_asBytes _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.Map(func(coer16 func(uint16) uint16) func(interface{}) interface{} {
		return func(arg17 interface{}) interface{} {
			return coer16(arg17.(uint16))
		}
	}(func(_3_cu uint16) uint16 {
		return uint16(_3_cu)
	}), _1_asUtf16CodeUnits)
	_ = _2_asBytes
	return m_Std_Wrappers.Companion_Option_.Create_Some_(_2_asBytes)
}
func (_static *CompanionStruct_Default___) FromUTF16Checked(bs _dafny.Sequence) m_Std_Wrappers.Result {
	var _0_asCodeUnits _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.Map(func(coer17 func(uint16) uint16) func(interface{}) interface{} {
		return func(arg18 interface{}) interface{} {
			return coer17(arg18.(uint16))
		}
	}(func(_1_c uint16) uint16 {
		return uint16(_1_c)
	}), bs)
	_ = _0_asCodeUnits
	var _2_valueOrError0 m_Std_Wrappers.Result = m_Std_Unicode_Utf16EncodingForm.Companion_Default___.DecodeCodeUnitSequenceChecked(_0_asCodeUnits)
	_ = _2_valueOrError0
	if (_2_valueOrError0).IsFailure() {
		return (_2_valueOrError0).PropagateFailure()
	} else {
		var _3_utf32 _dafny.Sequence = (_2_valueOrError0).Extract().(_dafny.Sequence)
		_ = _3_utf32
		var _4_asChars _dafny.Sequence = m_Std_Collections_Seq.Companion_Default___.Map(func(coer18 func(uint32) _dafny.CodePoint) func(interface{}) interface{} {
			return func(arg19 interface{}) interface{} {
				return coer18(arg19.(uint32))
			}
		}(Companion_Default___.CharFromUnicodeScalarValue), _3_utf32)
		_ = _4_asChars
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_4_asChars)
	}
}
func (_static *CompanionStruct_Default___) ASCIIToUTF8(s _dafny.Sequence) _dafny.Sequence {
	return m_Std_Collections_Seq.Companion_Default___.Map(func(coer19 func(_dafny.CodePoint) uint8) func(interface{}) interface{} {
		return func(arg20 interface{}) interface{} {
			return coer19(arg20.(_dafny.CodePoint))
		}
	}(func(_0_c _dafny.CodePoint) uint8 {
		return uint8(_0_c)
	}), s)
}
func (_static *CompanionStruct_Default___) ASCIIToUTF16(s _dafny.Sequence) _dafny.Sequence {
	return m_Std_Collections_Seq.Companion_Default___.Map(func(coer20 func(_dafny.CodePoint) uint16) func(interface{}) interface{} {
		return func(arg21 interface{}) interface{} {
			return coer20(arg21.(_dafny.CodePoint))
		}
	}(func(_0_c _dafny.CodePoint) uint16 {
		return uint16(_0_c)
	}), s)
}

// End of class Default__
