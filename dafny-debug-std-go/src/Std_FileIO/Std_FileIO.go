// Package Std_FileIO
// Dafny module Std_FileIO compiled into Go

package Std_FileIO

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
	return "Std_FileIO.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ReadBytesFromFile(path _dafny.Sequence) m_Std_Wrappers.Result {
	var res m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Default(_dafny.EmptySeq)
	_ = res
	var _0_isError bool
	_ = _0_isError
	var _1_bytesRead _dafny.Sequence
	_ = _1_bytesRead
	var _2_errorMsg _dafny.Sequence
	_ = _2_errorMsg
	var _out0 bool
	_ = _out0
	var _out1 _dafny.Sequence
	_ = _out1
	var _out2 _dafny.Sequence
	_ = _out2
	_out0, _out1, _out2 = m_Std_FileIOInternalExterns.Companion_Default___.INTERNAL__ReadBytesFromFile(path)
	_0_isError = _out0
	_1_bytesRead = _out1
	_2_errorMsg = _out2
	if _0_isError {
		res = m_Std_Wrappers.Companion_Result_.Create_Failure_(_2_errorMsg)
	} else {
		res = m_Std_Wrappers.Companion_Result_.Create_Success_(_1_bytesRead)
	}
	return res
	return res
}
func (_static *CompanionStruct_Default___) WriteBytesToFile(path _dafny.Sequence, bytes _dafny.Sequence) m_Std_Wrappers.Result {
	var res m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Default(_dafny.TupleOf())
	_ = res
	var _0_isError bool
	_ = _0_isError
	var _1_errorMsg _dafny.Sequence
	_ = _1_errorMsg
	var _out0 bool
	_ = _out0
	var _out1 _dafny.Sequence
	_ = _out1
	_out0, _out1 = m_Std_FileIOInternalExterns.Companion_Default___.INTERNAL__WriteBytesToFile(path, bytes)
	_0_isError = _out0
	_1_errorMsg = _out1
	if _0_isError {
		res = m_Std_Wrappers.Companion_Result_.Create_Failure_(_1_errorMsg)
	} else {
		res = m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.TupleOf())
	}
	return res
	return res
}
func (_static *CompanionStruct_Default___) ReadUTF8FromFile(fileName _dafny.Sequence) m_Std_Wrappers.Result {
	var r m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Default(_dafny.EmptySeq)
	_ = r
	var _0_valueOrError0 m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Default(_dafny.EmptySeq)
	_ = _0_valueOrError0
	var _out0 m_Std_Wrappers.Result
	_ = _out0
	_out0 = Companion_Default___.ReadBytesFromFile(fileName)
	_0_valueOrError0 = _out0
	if (_0_valueOrError0).IsFailure() {
		r = (_0_valueOrError0).PropagateFailure()
		return r
	}
	var _1_bytes _dafny.Sequence
	_ = _1_bytes
	_1_bytes = (_0_valueOrError0).Extract().(_dafny.Sequence)
	r = m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.FromUTF8Checked(_dafny.SeqCreate((_dafny.IntOfUint32((_1_bytes).Cardinality())).Uint32(), func(coer23 func(_dafny.Int) uint8) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer23(arg24)
		}
	}((func(_2_bytes _dafny.Sequence) func(_dafny.Int) uint8 {
		return func(_3_i _dafny.Int) uint8 {
			return uint8((_2_bytes).Select((_3_i).Uint32()).(uint8))
		}
	})(_1_bytes))))
	return r
	return r
}
func (_static *CompanionStruct_Default___) WriteUTF8ToFile(fileName _dafny.Sequence, content _dafny.Sequence) m_Std_Wrappers.Outcome {
	var r m_Std_Wrappers.Outcome = m_Std_Wrappers.Companion_Outcome_.Default()
	_ = r
	var _0_bytes _dafny.Sequence
	_ = _0_bytes
	_0_bytes = (m_Std_Unicode_UnicodeStringsWithUnicodeChar.Companion_Default___.ToUTF8Checked(content)).Dtor_value().(_dafny.Sequence)
	var _1_writeResult m_Std_Wrappers.Result
	_ = _1_writeResult
	var _out0 m_Std_Wrappers.Result
	_ = _out0
	_out0 = Companion_Default___.WriteBytesToFile(fileName, _dafny.SeqCreate((_dafny.IntOfUint32((_0_bytes).Cardinality())).Uint32(), func(coer24 func(_dafny.Int) uint8) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer24(arg25)
		}
	}((func(_2_bytes _dafny.Sequence) func(_dafny.Int) uint8 {
		return func(_3_i _dafny.Int) uint8 {
			return uint8((_2_bytes).Select((_3_i).Uint32()).(uint8))
		}
	})(_0_bytes))))
	_1_writeResult = _out0
	if (_1_writeResult).IsFailure() {
		r = m_Std_Wrappers.Companion_Outcome_.Create_Fail_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("Failed to write to file '"), fileName), _dafny.UnicodeSeqOfUtf8Bytes("': ")), (_1_writeResult).Dtor_error().(_dafny.Sequence)))
		return r
	}
	r = m_Std_Wrappers.Companion_Outcome_.Create_Pass_()
	return r
	return r
}

// End of class Default__
