// Package Std_JSON_Utils_Views_Writers
// Dafny module Std_JSON_Utils_Views_Writers compiled into Go

package Std_JSON_Utils_Views_Writers

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
	m_Std_JSON_Spec "Std_JSON_Spec"
	m_Std_JSON_Utils_Views_Core "Std_JSON_Utils_Views_Core"
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

type Dummy__ struct{}

// Definition of datatype Chain
type Chain struct {
	Data_Chain_
}

func (_this Chain) Get_() Data_Chain_ {
	return _this.Data_Chain_
}

type Data_Chain_ interface {
	isChain()
}

type CompanionStruct_Chain_ struct {
}

var Companion_Chain_ = CompanionStruct_Chain_{}

type Chain_Empty struct {
}

func (Chain_Empty) isChain() {}

func (CompanionStruct_Chain_) Create_Empty_() Chain {
	return Chain{Chain_Empty{}}
}

func (_this Chain) Is_Empty() bool {
	_, ok := _this.Get_().(Chain_Empty)
	return ok
}

type Chain_Chain struct {
	Previous Chain
	V        m_Std_JSON_Utils_Views_Core.View__
}

func (Chain_Chain) isChain() {}

func (CompanionStruct_Chain_) Create_Chain_(Previous Chain, V m_Std_JSON_Utils_Views_Core.View__) Chain {
	return Chain{Chain_Chain{Previous, V}}
}

func (_this Chain) Is_Chain() bool {
	_, ok := _this.Get_().(Chain_Chain)
	return ok
}

func (CompanionStruct_Chain_) Default() Chain {
	return Companion_Chain_.Create_Empty_()
}

func (_this Chain) Dtor_previous() Chain {
	return _this.Get_().(Chain_Chain).Previous
}

func (_this Chain) Dtor_v() m_Std_JSON_Utils_Views_Core.View__ {
	return _this.Get_().(Chain_Chain).V
}

func (_this Chain) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Chain_Empty:
		{
			return "Writers.Chain.Empty"
		}
	case Chain_Chain:
		{
			return "Writers.Chain.Chain" + "(" + _dafny.String(data.Previous) + ", " + _dafny.String(data.V) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Chain) Equals(other Chain) bool {
	switch data1 := _this.Get_().(type) {
	case Chain_Empty:
		{
			_, ok := other.Get_().(Chain_Empty)
			return ok
		}
	case Chain_Chain:
		{
			data2, ok := other.Get_().(Chain_Chain)
			return ok && data1.Previous.Equals(data2.Previous) && data1.V.Equals(data2.V)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Chain) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Chain)
	return ok && _this.Equals(typed)
}

func Type_Chain_() _dafny.TypeDescriptor {
	return type_Chain_{}
}

type type_Chain_ struct {
}

func (_this type_Chain_) Default() interface{} {
	return Companion_Chain_.Default()
}

func (_this type_Chain_) String() string {
	return "Std_JSON_Utils_Views_Writers.Chain"
}
func (_this Chain) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Chain{}

func (_this Chain) Length() _dafny.Int {
	{
		var _0___accumulator _dafny.Int = _dafny.Zero
		_ = _0___accumulator
		goto TAIL_CALL_START
	TAIL_CALL_START:
		if (_this).Is_Empty() {
			return (_dafny.Zero).Plus(_0___accumulator)
		} else {
			_0___accumulator = (_dafny.IntOfUint32(((_this).Dtor_v()).Length())).Plus(_0___accumulator)
			var _in0 Chain = (_this).Dtor_previous()
			_ = _in0
			_this = _in0

			goto TAIL_CALL_START
		}
	}
}
func (_this Chain) Count() _dafny.Int {
	{
		var _0___accumulator _dafny.Int = _dafny.Zero
		_ = _0___accumulator
		goto TAIL_CALL_START
	TAIL_CALL_START:
		if (_this).Is_Empty() {
			return (_dafny.Zero).Plus(_0___accumulator)
		} else {
			_0___accumulator = (_dafny.One).Plus(_0___accumulator)
			var _in0 Chain = (_this).Dtor_previous()
			_ = _in0
			_this = _in0

			goto TAIL_CALL_START
		}
	}
}
func (_this Chain) Bytes() _dafny.Sequence {
	{
		var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
		_ = _0___accumulator
		goto TAIL_CALL_START
	TAIL_CALL_START:
		if (_this).Is_Empty() {
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(), _0___accumulator)
		} else {
			_0___accumulator = _dafny.Companion_Sequence_.Concatenate(((_this).Dtor_v()).Bytes(), _0___accumulator)
			var _in0 Chain = (_this).Dtor_previous()
			_ = _in0
			_this = _in0

			goto TAIL_CALL_START
		}
	}
}
func (_this Chain) Append(v_k m_Std_JSON_Utils_Views_Core.View__) Chain {
	{
		if ((_this).Is_Chain()) && (m_Std_JSON_Utils_Views_Core.Companion_Default___.Adjacent((_this).Dtor_v(), v_k)) {
			return Companion_Chain_.Create_Chain_((_this).Dtor_previous(), m_Std_JSON_Utils_Views_Core.Companion_Default___.Merge((_this).Dtor_v(), v_k))
		} else {
			return Companion_Chain_.Create_Chain_(_this, v_k)
		}
	}
}
func (_this Chain) CopyTo(dest _dafny.Array, end uint32) {
	{
		goto TAIL_CALL_START
	TAIL_CALL_START:
		if (_this).Is_Chain() {
			var _0_end uint32
			_ = _0_end
			_0_end = (end) - (func() uint32 { return (((_this).Dtor_v()).Length()) })()
			((_this).Dtor_v()).CopyTo(dest, _0_end)
			var _in0 Chain = (_this).Dtor_previous()
			_ = _in0
			var _in1 _dafny.Array = dest
			_ = _in1
			var _in2 uint32 = _0_end
			_ = _in2
			_this = _in0

			dest = _in1
			end = _in2
			goto TAIL_CALL_START
		}
	}
}

// End of datatype Chain

// Definition of class Writer
type Writer struct {
}

func New_Writer_() *Writer {
	_this := Writer{}

	return &_this
}

type CompanionStruct_Writer_ struct {
}

var Companion_Writer_ = CompanionStruct_Writer_{}

func (*Writer) String() string {
	return "Std_JSON_Utils_Views_Writers.Writer"
}
func (_this *CompanionStruct_Writer_) Witness() Writer__ {
	return Companion_Writer___.Create_Writer_(uint32(0), Companion_Chain_.Create_Empty_())
}

// End of class Writer

func Type_Writer_() _dafny.TypeDescriptor {
	return type_Writer_{}
}

type type_Writer_ struct {
}

func (_this type_Writer_) Default() interface{} {
	return Companion_Writer_.Witness()
}

func (_this type_Writer_) String() string {
	return "Std_JSON_Utils_Views_Writers.Writer"
}

// Definition of datatype Writer__
type Writer__ struct {
	Data_Writer___
}

func (_this Writer__) Get_() Data_Writer___ {
	return _this.Data_Writer___
}

type Data_Writer___ interface {
	isWriter__()
}

type CompanionStruct_Writer___ struct {
}

var Companion_Writer___ = CompanionStruct_Writer___{}

type Writer___Writer struct {
	Length uint32
	Chain  Chain
}

func (Writer___Writer) isWriter__() {}

func (CompanionStruct_Writer___) Create_Writer_(Length uint32, Chain Chain) Writer__ {
	return Writer__{Writer___Writer{Length, Chain}}
}

func (_this Writer__) Is_Writer() bool {
	_, ok := _this.Get_().(Writer___Writer)
	return ok
}

func (CompanionStruct_Writer___) Default() Writer__ {
	return Companion_Writer___.Create_Writer_(uint32(0), Companion_Chain_.Default())
}

func (_this Writer__) Dtor_length() uint32 {
	return _this.Get_().(Writer___Writer).Length
}

func (_this Writer__) Dtor_chain() Chain {
	return _this.Get_().(Writer___Writer).Chain
}

func (_this Writer__) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Writer___Writer:
		{
			return "Writers.Writer_.Writer" + "(" + _dafny.String(data.Length) + ", " + _dafny.String(data.Chain) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Writer__) Equals(other Writer__) bool {
	switch data1 := _this.Get_().(type) {
	case Writer___Writer:
		{
			data2, ok := other.Get_().(Writer___Writer)
			return ok && data1.Length == data2.Length && data1.Chain.Equals(data2.Chain)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Writer__) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Writer__)
	return ok && _this.Equals(typed)
}

func Type_Writer___() _dafny.TypeDescriptor {
	return type_Writer___{}
}

type type_Writer___ struct {
}

func (_this type_Writer___) Default() interface{} {
	return Companion_Writer___.Default()
}

func (_this type_Writer___) String() string {
	return "Std_JSON_Utils_Views_Writers.Writer__"
}
func (_this Writer__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Writer__{}

func (_this Writer__) Bytes() _dafny.Sequence {
	{
		return ((_this).Dtor_chain()).Bytes()
	}
}
func (_static CompanionStruct_Writer___) SaturatedAddU32(a uint32, b uint32) uint32 {
	if (a) <= ((m_Std_BoundedInts.Companion_Default___.UINT32__MAX()) - (func() uint32 { return (b) })()) {
		return (a) + (b)
	} else {
		return m_Std_BoundedInts.Companion_Default___.UINT32__MAX()
	}
}
func (_this Writer__) Append(v_k m_Std_JSON_Utils_Views_Core.View__) Writer__ {
	{
		return Companion_Writer___.Create_Writer_(Companion_Writer___.SaturatedAddU32((_this).Dtor_length(), (v_k).Length()), ((_this).Dtor_chain()).Append(v_k))
	}
}
func (_this Writer__) Then(fn func(Writer__) Writer__) Writer__ {
	{
		return (fn)(_this)
	}
}
func (_this Writer__) CopyTo(dest _dafny.Array) {
	{
		((_this).Dtor_chain()).CopyTo(dest, (_this).Dtor_length())
	}
}
func (_this Writer__) ToArray() _dafny.Array {
	{
		var bs _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = bs
		var _len0_0 _dafny.Int = _dafny.IntOfAny((_this).Dtor_length())
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) uint8 = func(_0_i _dafny.Int) uint8 {
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
		bs = _nw0
		(_this).CopyTo(bs)
		return bs
	}
}
func (_static CompanionStruct_Writer___) Empty() Writer__ {
	return Companion_Writer___.Create_Writer_(uint32(0), Companion_Chain_.Create_Empty_())
}
func (_this Writer__) Unsaturated_q() bool {
	{
		return ((_this).Dtor_length()) != (m_Std_BoundedInts.Companion_Default___.UINT32__MAX()) /* dircomp */
	}
}
func (_this Writer__) Empty_q() bool {
	{
		return ((_this).Dtor_chain()).Is_Empty()
	}
}

// End of datatype Writer__
