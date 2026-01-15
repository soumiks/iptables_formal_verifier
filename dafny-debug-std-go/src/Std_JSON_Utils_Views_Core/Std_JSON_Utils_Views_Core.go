// Package Std_JSON_Utils_Views_Core
// Dafny module Std_JSON_Utils_Views_Core compiled into Go

package Std_JSON_Utils_Views_Core

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
	return "Std_JSON_Utils_Views_Core.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Adjacent(lv View__, rv View__) bool {
	return (((lv).Dtor_end()) == ((rv).Dtor_beg())) && (_dafny.Companion_Sequence_.Equal((lv).Dtor_s(), (rv).Dtor_s()))
}
func (_static *CompanionStruct_Default___) Merge(lv View__, rv View__) View__ {
	var _0_dt__update__tmp_h0 View__ = lv
	_ = _0_dt__update__tmp_h0
	var _1_dt__update_hend_h0 uint32 = (rv).Dtor_end()
	_ = _1_dt__update_hend_h0
	return Companion_View___.Create_View_((_0_dt__update__tmp_h0).Dtor_s(), (_0_dt__update__tmp_h0).Dtor_beg(), _1_dt__update_hend_h0)
}

// End of class Default__

// Definition of class View
type View struct {
}

func New_View_() *View {
	_this := View{}

	return &_this
}

type CompanionStruct_View_ struct {
}

var Companion_View_ = CompanionStruct_View_{}

func (*View) String() string {
	return "Std_JSON_Utils_Views_Core.View"
}
func (_this *CompanionStruct_View_) Witness() View__ {
	return Companion_View___.Create_View_(_dafny.SeqOf(), uint32(0), uint32(0))
}

// End of class View

func Type_View_() _dafny.TypeDescriptor {
	return type_View_{}
}

type type_View_ struct {
}

func (_this type_View_) Default() interface{} {
	return Companion_View_.Witness()
}

func (_this type_View_) String() string {
	return "Std_JSON_Utils_Views_Core.View"
}

// Definition of datatype View__
type View__ struct {
	Data_View___
}

func (_this View__) Get_() Data_View___ {
	return _this.Data_View___
}

type Data_View___ interface {
	isView__()
}

type CompanionStruct_View___ struct {
}

var Companion_View___ = CompanionStruct_View___{}

type View___View struct {
	S   _dafny.Sequence
	Beg uint32
	End uint32
}

func (View___View) isView__() {}

func (CompanionStruct_View___) Create_View_(S _dafny.Sequence, Beg uint32, End uint32) View__ {
	return View__{View___View{S, Beg, End}}
}

func (_this View__) Is_View() bool {
	_, ok := _this.Get_().(View___View)
	return ok
}

func (CompanionStruct_View___) Default() View__ {
	return Companion_View___.Create_View_(_dafny.EmptySeq, uint32(0), uint32(0))
}

func (_this View__) Dtor_s() _dafny.Sequence {
	return _this.Get_().(View___View).S
}

func (_this View__) Dtor_beg() uint32 {
	return _this.Get_().(View___View).Beg
}

func (_this View__) Dtor_end() uint32 {
	return _this.Get_().(View___View).End
}

func (_this View__) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case View___View:
		{
			return "Core.View_.View" + "(" + _dafny.String(data.S) + ", " + _dafny.String(data.Beg) + ", " + _dafny.String(data.End) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this View__) Equals(other View__) bool {
	switch data1 := _this.Get_().(type) {
	case View___View:
		{
			data2, ok := other.Get_().(View___View)
			return ok && data1.S.Equals(data2.S) && data1.Beg == data2.Beg && data1.End == data2.End
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this View__) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(View__)
	return ok && _this.Equals(typed)
}

func Type_View___() _dafny.TypeDescriptor {
	return type_View___{}
}

type type_View___ struct {
}

func (_this type_View___) Default() interface{} {
	return Companion_View___.Default()
}

func (_this type_View___) String() string {
	return "Std_JSON_Utils_Views_Core.View__"
}
func (_this View__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = View__{}

func (_this View__) Length() uint32 {
	{
		return ((_this).Dtor_end()) - (func() uint32 { return ((_this).Dtor_beg()) })()
	}
}
func (_this View__) Bytes() _dafny.Sequence {
	{
		return ((_this).Dtor_s()).Subsequence(uint32((_this).Dtor_beg()), uint32((_this).Dtor_end()))
	}
}
func (_static CompanionStruct_View___) OfBytes(bs _dafny.Sequence) View__ {
	return Companion_View___.Create_View_(bs, uint32(0), uint32((bs).Cardinality()))
}
func (_static CompanionStruct_View___) OfString(s _dafny.Sequence) _dafny.Sequence {
	return _dafny.SeqCreate((_dafny.IntOfUint32((s).Cardinality())).Uint32(), func(coer37 func(_dafny.Int) uint8) func(_dafny.Int) interface{} {
		return func(arg38 _dafny.Int) interface{} {
			return coer37(arg38)
		}
	}((func(_0_s _dafny.Sequence) func(_dafny.Int) uint8 {
		return func(_1_i _dafny.Int) uint8 {
			return uint8((_0_s).Select((_1_i).Uint32()).(_dafny.CodePoint))
		}
	})(s)))
}
func (_this View__) Byte_q(c uint8) bool {
	{
		var _hresult bool = false
		_ = _hresult
		_hresult = (((_this).Length()) == (uint32(1))) && (((_this).At(uint32(0))) == (c))
		return _hresult
		return _hresult
	}
}
func (_this View__) Char_q(c _dafny.CodePoint) bool {
	{
		return (_this).Byte_q(uint8(c))
	}
}
func (_this View__) At(idx uint32) uint8 {
	{
		return ((_this).Dtor_s()).Select(uint32(((_this).Dtor_beg()) + (idx))).(uint8)
	}
}
func (_this View__) Peek() int16 {
	{
		if (_this).Empty_q() {
			return int16(-1)
		} else {
			return int16((_this).At(uint32(0)))
		}
	}
}
func (_this View__) CopyTo(dest _dafny.Array, start uint32) {
	{
		var _hi0 uint32 = (_this).Length()
		_ = _hi0
		for _0_idx := uint32(0); _0_idx < _hi0; _0_idx++ {
			var _index0 uint32 = (start) + (_0_idx)
			_ = _index0
			(dest).ArraySet1Byte(((_this).Dtor_s()).Select(uint32(((_this).Dtor_beg())+(_0_idx))).(uint8), int(_index0))
		}
	}
}
func (_static CompanionStruct_View___) Empty() View__ {
	return Companion_View___.Create_View_(_dafny.SeqOf(), uint32(0), uint32(0))
}
func (_this View__) Empty_q() bool {
	{
		return ((_this).Dtor_beg()) == ((_this).Dtor_end())
	}
}

// End of datatype View__
