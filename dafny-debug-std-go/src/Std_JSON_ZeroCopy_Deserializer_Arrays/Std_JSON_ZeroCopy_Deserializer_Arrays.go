// Package Std_JSON_ZeroCopy_Deserializer_Arrays
// Dafny module Std_JSON_ZeroCopy_Deserializer_Arrays compiled into Go

package Std_JSON_ZeroCopy_Deserializer_Arrays

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
	m_Std_JSON_ZeroCopy_Deserializer_ArrayParams "Std_JSON_ZeroCopy_Deserializer_ArrayParams"
	m_Std_JSON_ZeroCopy_Deserializer_Core "Std_JSON_ZeroCopy_Deserializer_Core"
	m_Std_JSON_ZeroCopy_Deserializer_Numbers "Std_JSON_ZeroCopy_Deserializer_Numbers"
	m_Std_JSON_ZeroCopy_Deserializer_ObjectParams "Std_JSON_ZeroCopy_Deserializer_ObjectParams"
	m_Std_JSON_ZeroCopy_Deserializer_Objects "Std_JSON_ZeroCopy_Deserializer_Objects"
	m_Std_JSON_ZeroCopy_Deserializer_Strings "Std_JSON_ZeroCopy_Deserializer_Strings"
	m_Std_JSON_ZeroCopy_Serializer "Std_JSON_ZeroCopy_Serializer"
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
var _ m_Std_JSON_ZeroCopy_Serializer.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Core.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Strings.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Numbers.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_ObjectParams.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Objects.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Dummy__

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
	return "Std_JSON_ZeroCopy_Deserializer_Arrays.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Array(cs m_Std_JSON_Utils_Cursors.Cursor__, json func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = Companion_Default___.Bracketed(cs, json)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_sp m_Std_JSON_Utils_Cursors.Split = (_0_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Split)
		_ = _1_sp
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_1_sp)
	}
}
func (_static *CompanionStruct_Default___) Open(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = (cs).AssertByte(m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Companion_Default___.OPEN())
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_cs m_Std_JSON_Utils_Cursors.Cursor__ = (_0_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Cursor__)
		_ = _1_cs
		return m_Std_Wrappers.Companion_Result_.Create_Success_((_1_cs).Split())
	}
}
func (_static *CompanionStruct_Default___) Close(cs m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = (cs).AssertByte(m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Companion_Default___.CLOSE())
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_cs m_Std_JSON_Utils_Cursors.Cursor__ = (_0_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Cursor__)
		_ = _1_cs
		return m_Std_Wrappers.Companion_Result_.Create_Success_((_1_cs).Split())
	}
}
func (_static *CompanionStruct_Default___) BracketedFromParts(open m_Std_JSON_Utils_Cursors.Split, elems m_Std_JSON_Utils_Cursors.Split, close_ m_Std_JSON_Utils_Cursors.Split) m_Std_JSON_Utils_Cursors.Split {
	var _0_sp m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_Std_JSON_Grammar.Companion_Bracketed_.Create_Bracketed_((open).Dtor_t().(m_Std_JSON_Grammar.Structural), (elems).Dtor_t().(_dafny.Sequence), (close_).Dtor_t().(m_Std_JSON_Grammar.Structural)), (close_).Dtor_cs())
	_ = _0_sp
	return _0_sp
}
func (_static *CompanionStruct_Default___) AppendWithSuffix(elems m_Std_JSON_Utils_Cursors.Split, elem m_Std_JSON_Utils_Cursors.Split, sep m_Std_JSON_Utils_Cursors.Split) m_Std_JSON_Utils_Cursors.Split {
	var _0_suffixed m_Std_JSON_Grammar.Suffixed = m_Std_JSON_Grammar.Companion_Suffixed_.Create_Suffixed_((elem).Dtor_t().(m_Std_JSON_Grammar.Value), m_Std_JSON_Grammar.Companion_Maybe_.Create_NonEmpty_((sep).Dtor_t().(m_Std_JSON_Grammar.Structural)))
	_ = _0_suffixed
	var _1_elems_k m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(_dafny.Companion_Sequence_.Concatenate((elems).Dtor_t().(_dafny.Sequence), _dafny.SeqOf(_0_suffixed)), (sep).Dtor_cs())
	_ = _1_elems_k
	return _1_elems_k
}
func (_static *CompanionStruct_Default___) AppendLast(elems m_Std_JSON_Utils_Cursors.Split, elem m_Std_JSON_Utils_Cursors.Split, sep m_Std_JSON_Utils_Cursors.Split) m_Std_JSON_Utils_Cursors.Split {
	var _0_suffixed m_Std_JSON_Grammar.Suffixed = m_Std_JSON_Grammar.Companion_Suffixed_.Create_Suffixed_((elem).Dtor_t().(m_Std_JSON_Grammar.Value), m_Std_JSON_Grammar.Companion_Maybe_.Create_Empty_())
	_ = _0_suffixed
	var _1_elems_k m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(_dafny.Companion_Sequence_.Concatenate((elems).Dtor_t().(_dafny.Sequence), _dafny.SeqOf(_0_suffixed)), (elem).Dtor_cs())
	_ = _1_elems_k
	return _1_elems_k
}
func (_static *CompanionStruct_Default___) Elements(json func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result, open m_Std_JSON_Utils_Cursors.Split, elems m_Std_JSON_Utils_Cursors.Split) m_Std_Wrappers.Result {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _0_valueOrError0 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Companion_Default___.Element((elems).Dtor_cs(), json)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_elem m_Std_JSON_Utils_Cursors.Split = (_0_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Split)
		_ = _1_elem
		if ((_1_elem).Dtor_cs()).EOF_q() {
			return m_Std_Wrappers.Companion_Result_.Create_Failure_(m_Std_JSON_Utils_Cursors.Companion_CursorError_.Create_EOF_())
		} else {
			var _2_sep m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_ZeroCopy_Deserializer_Core.Companion_Default___.TryStructural((_1_elem).Dtor_cs())
			_ = _2_sep
			var _3_s0 int16 = (((_2_sep).Dtor_t().(m_Std_JSON_Grammar.Structural)).Dtor_t().(m_Std_JSON_Utils_Views_Core.View__)).Peek()
			_ = _3_s0
			if ((_3_s0) == (int16(Companion_Default___.SEPARATOR()))) && (((((_2_sep).Dtor_t().(m_Std_JSON_Grammar.Structural)).Dtor_t().(m_Std_JSON_Utils_Views_Core.View__)).Length()) == (uint32(1))) {
				var _4_sep m_Std_JSON_Utils_Cursors.Split = _2_sep
				_ = _4_sep
				var _5_elems m_Std_JSON_Utils_Cursors.Split = Companion_Default___.AppendWithSuffix(elems, _1_elem, _4_sep)
				_ = _5_elems
				var _in0 func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result = json
				_ = _in0
				var _in1 m_Std_JSON_Utils_Cursors.Split = open
				_ = _in1
				var _in2 m_Std_JSON_Utils_Cursors.Split = _5_elems
				_ = _in2
				json = _in0
				open = _in1
				elems = _in2
				goto TAIL_CALL_START
			} else if ((_3_s0) == (int16(m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Companion_Default___.CLOSE()))) && (((((_2_sep).Dtor_t().(m_Std_JSON_Grammar.Structural)).Dtor_t().(m_Std_JSON_Utils_Views_Core.View__)).Length()) == (uint32(1))) {
				var _6_sep m_Std_JSON_Utils_Cursors.Split = _2_sep
				_ = _6_sep
				var _7_elems_k m_Std_JSON_Utils_Cursors.Split = Companion_Default___.AppendLast(elems, _1_elem, _6_sep)
				_ = _7_elems_k
				var _8_bracketed m_Std_JSON_Utils_Cursors.Split = Companion_Default___.BracketedFromParts(open, _7_elems_k, _6_sep)
				_ = _8_bracketed
				return m_Std_Wrappers.Companion_Result_.Create_Success_(_8_bracketed)
			} else {
				var _9_separator uint8 = Companion_Default___.SEPARATOR()
				_ = _9_separator
				var _10_pr m_Std_Wrappers.Result = m_Std_Wrappers.Companion_Result_.Create_Failure_(m_Std_JSON_Utils_Cursors.Companion_CursorError_.Create_ExpectingAnyByte_(_dafny.SeqOf(m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Companion_Default___.CLOSE(), _9_separator), _3_s0))
				_ = _10_pr
				return _10_pr
			}
		}
	}
}
func (_static *CompanionStruct_Default___) Bracketed(cs m_Std_JSON_Utils_Cursors.Cursor__, json func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result) m_Std_Wrappers.Result {
	var _0_valueOrError0 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_Core.Companion_Default___.Structural(cs, func(coer58 func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result) func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
		return func(arg59 m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
			return coer58(arg59)
		}
	}(Companion_Default___.Open))
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_open m_Std_JSON_Utils_Cursors.Split = (_0_valueOrError0).Extract().(m_Std_JSON_Utils_Cursors.Split)
		_ = _1_open
		var _2_elems m_Std_JSON_Utils_Cursors.Split = m_Std_JSON_Utils_Cursors.Companion_Split_.Create_SP_(_dafny.SeqOf(), (_1_open).Dtor_cs())
		_ = _2_elems
		if (((_1_open).Dtor_cs()).Peek()) == (int16(m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Companion_Default___.CLOSE())) {
			var _3_p func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result = Companion_Default___.Close
			_ = _3_p
			var _4_valueOrError1 m_Std_Wrappers.Result = m_Std_JSON_ZeroCopy_Deserializer_Core.Companion_Default___.Structural((_1_open).Dtor_cs(), func(coer59 func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result) func(m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
				return func(arg60 m_Std_JSON_Utils_Cursors.Cursor__) m_Std_Wrappers.Result {
					return coer59(arg60)
				}
			}(_3_p))
			_ = _4_valueOrError1
			if (_4_valueOrError1).IsFailure() {
				return (_4_valueOrError1).PropagateFailure()
			} else {
				var _5_close m_Std_JSON_Utils_Cursors.Split = (_4_valueOrError1).Extract().(m_Std_JSON_Utils_Cursors.Split)
				_ = _5_close
				return m_Std_Wrappers.Companion_Result_.Create_Success_(Companion_Default___.BracketedFromParts(_1_open, _2_elems, _5_close))
			}
		} else {
			return Companion_Default___.Elements(json, _1_open, _2_elems)
		}
	}
}
func (_static *CompanionStruct_Default___) SpecViewOpen() func(m_Std_JSON_Utils_Views_Core.View__) _dafny.Sequence {
	return m_Std_JSON_ZeroCopy_Deserializer_Core.Companion_Default___.SpecView()
}
func (_static *CompanionStruct_Default___) SpecViewClose() func(m_Std_JSON_Utils_Views_Core.View__) _dafny.Sequence {
	return m_Std_JSON_ZeroCopy_Deserializer_Core.Companion_Default___.SpecView()
}
func (_static *CompanionStruct_Default___) SEPARATOR() uint8 {
	return uint8(_dafny.CodePoint(','))
}

// End of class Default__

// Definition of class Jopen
type Jopen struct {
}

func New_Jopen_() *Jopen {
	_this := Jopen{}

	return &_this
}

type CompanionStruct_Jopen_ struct {
}

var Companion_Jopen_ = CompanionStruct_Jopen_{}

func (*Jopen) String() string {
	return "Std_JSON_ZeroCopy_Deserializer_Arrays.Jopen"
}
func (_this *CompanionStruct_Jopen_) Witness() m_Std_JSON_Utils_Views_Core.View__ {
	return m_Std_JSON_Utils_Views_Core.Companion_View___.OfBytes(_dafny.SeqOf(m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Companion_Default___.OPEN()))
}

// End of class Jopen

func Type_Jopen_() _dafny.TypeDescriptor {
	return type_Jopen_{}
}

type type_Jopen_ struct {
}

func (_this type_Jopen_) Default() interface{} {
	return Companion_Jopen_.Witness()
}

func (_this type_Jopen_) String() string {
	return "Std_JSON_ZeroCopy_Deserializer_Arrays.Jopen"
}

// Definition of class Jclose
type Jclose struct {
}

func New_Jclose_() *Jclose {
	_this := Jclose{}

	return &_this
}

type CompanionStruct_Jclose_ struct {
}

var Companion_Jclose_ = CompanionStruct_Jclose_{}

func (*Jclose) String() string {
	return "Std_JSON_ZeroCopy_Deserializer_Arrays.Jclose"
}
func (_this *CompanionStruct_Jclose_) Witness() m_Std_JSON_Utils_Views_Core.View__ {
	return m_Std_JSON_Utils_Views_Core.Companion_View___.OfBytes(_dafny.SeqOf(m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Companion_Default___.CLOSE()))
}

// End of class Jclose

func Type_Jclose_() _dafny.TypeDescriptor {
	return type_Jclose_{}
}

type type_Jclose_ struct {
}

func (_this type_Jclose_) Default() interface{} {
	return Companion_Jclose_.Witness()
}

func (_this type_Jclose_) String() string {
	return "Std_JSON_ZeroCopy_Deserializer_Arrays.Jclose"
}
