// Package Std_Unicode_Base
// Dafny module Std_Unicode_Base compiled into Go

package Std_Unicode_Base

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
	m_Std_Actions "Std_Actions"
	m_Std_ActionsExterns "Std_ActionsExterns"
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
	m_Std_Termination "Std_Termination"
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
	return "Std_Unicode_Base.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) IsInAssignedPlane(i uint32) bool {
	var _0_plane uint8 = uint8((i) >> (uint8(16)))
	_ = _0_plane
	return (Companion_Default___.ASSIGNED__PLANES()).Contains(_0_plane)
}
func (_static *CompanionStruct_Default___) HIGH__SURROGATE__MIN() uint32 {
	return uint32(55296)
}
func (_static *CompanionStruct_Default___) HIGH__SURROGATE__MAX() uint32 {
	return uint32(56319)
}
func (_static *CompanionStruct_Default___) LOW__SURROGATE__MIN() uint32 {
	return uint32(56320)
}
func (_static *CompanionStruct_Default___) LOW__SURROGATE__MAX() uint32 {
	return uint32(57343)
}
func (_static *CompanionStruct_Default___) ASSIGNED__PLANES() _dafny.Set {
	return _dafny.SetOf(uint8(0), uint8(1), uint8(2), uint8(3), uint8(14), uint8(15), uint8(16))
}

// End of class Default__

// Definition of class CodePoint
type CodePoint struct {
}

func New_CodePoint_() *CodePoint {
	_this := CodePoint{}

	return &_this
}

type CompanionStruct_CodePoint_ struct {
}

var Companion_CodePoint_ = CompanionStruct_CodePoint_{}

func (*CodePoint) String() string {
	return "Std_Unicode_Base.CodePoint"
}

// End of class CodePoint

func Type_CodePoint_() _dafny.TypeDescriptor {
	return type_CodePoint_{}
}

type type_CodePoint_ struct {
}

func (_this type_CodePoint_) Default() interface{} {
	return 0
}

func (_this type_CodePoint_) String() string {
	return "Std_Unicode_Base.CodePoint"
}
func (_this *CompanionStruct_CodePoint_) Is_(__source uint32) bool {
	var _0_i uint32 = (__source)
	_ = _0_i
	return ((uint32(0)) <= (_0_i)) && ((_0_i) <= (uint32(1114111)))
}

// Definition of class HighSurrogateCodePoint
type HighSurrogateCodePoint struct {
}

func New_HighSurrogateCodePoint_() *HighSurrogateCodePoint {
	_this := HighSurrogateCodePoint{}

	return &_this
}

type CompanionStruct_HighSurrogateCodePoint_ struct {
}

var Companion_HighSurrogateCodePoint_ = CompanionStruct_HighSurrogateCodePoint_{}

func (*HighSurrogateCodePoint) String() string {
	return "Std_Unicode_Base.HighSurrogateCodePoint"
}
func (_this *CompanionStruct_HighSurrogateCodePoint_) Witness() uint32 {
	return Companion_Default___.HIGH__SURROGATE__MIN()
}

// End of class HighSurrogateCodePoint

func Type_HighSurrogateCodePoint_() _dafny.TypeDescriptor {
	return type_HighSurrogateCodePoint_{}
}

type type_HighSurrogateCodePoint_ struct {
}

func (_this type_HighSurrogateCodePoint_) Default() interface{} {
	return Companion_HighSurrogateCodePoint_.Witness()
}

func (_this type_HighSurrogateCodePoint_) String() string {
	return "Std_Unicode_Base.HighSurrogateCodePoint"
}
func (_this *CompanionStruct_HighSurrogateCodePoint_) Is_(__source uint32) bool {
	var _1_p uint32 = (__source)
	_ = _1_p
	if Companion_CodePoint_.Is_(_1_p) {
		return ((Companion_Default___.HIGH__SURROGATE__MIN()) <= (_1_p)) && ((_1_p) <= (Companion_Default___.HIGH__SURROGATE__MAX()))
	}
	return false
}

// Definition of class LowSurrogateCodePoint
type LowSurrogateCodePoint struct {
}

func New_LowSurrogateCodePoint_() *LowSurrogateCodePoint {
	_this := LowSurrogateCodePoint{}

	return &_this
}

type CompanionStruct_LowSurrogateCodePoint_ struct {
}

var Companion_LowSurrogateCodePoint_ = CompanionStruct_LowSurrogateCodePoint_{}

func (*LowSurrogateCodePoint) String() string {
	return "Std_Unicode_Base.LowSurrogateCodePoint"
}
func (_this *CompanionStruct_LowSurrogateCodePoint_) Witness() uint32 {
	return Companion_Default___.LOW__SURROGATE__MIN()
}

// End of class LowSurrogateCodePoint

func Type_LowSurrogateCodePoint_() _dafny.TypeDescriptor {
	return type_LowSurrogateCodePoint_{}
}

type type_LowSurrogateCodePoint_ struct {
}

func (_this type_LowSurrogateCodePoint_) Default() interface{} {
	return Companion_LowSurrogateCodePoint_.Witness()
}

func (_this type_LowSurrogateCodePoint_) String() string {
	return "Std_Unicode_Base.LowSurrogateCodePoint"
}
func (_this *CompanionStruct_LowSurrogateCodePoint_) Is_(__source uint32) bool {
	var _2_p uint32 = (__source)
	_ = _2_p
	if Companion_CodePoint_.Is_(_2_p) {
		return ((Companion_Default___.LOW__SURROGATE__MIN()) <= (_2_p)) && ((_2_p) <= (Companion_Default___.LOW__SURROGATE__MAX()))
	}
	return false
}

// Definition of class ScalarValue
type ScalarValue struct {
}

func New_ScalarValue_() *ScalarValue {
	_this := ScalarValue{}

	return &_this
}

type CompanionStruct_ScalarValue_ struct {
}

var Companion_ScalarValue_ = CompanionStruct_ScalarValue_{}

func (*ScalarValue) String() string {
	return "Std_Unicode_Base.ScalarValue"
}

// End of class ScalarValue

func Type_ScalarValue_() _dafny.TypeDescriptor {
	return type_ScalarValue_{}
}

type type_ScalarValue_ struct {
}

func (_this type_ScalarValue_) Default() interface{} {
	return 0
}

func (_this type_ScalarValue_) String() string {
	return "Std_Unicode_Base.ScalarValue"
}
func (_this *CompanionStruct_ScalarValue_) Is_(__source uint32) bool {
	var _3_p uint32 = (__source)
	_ = _3_p
	if Companion_CodePoint_.Is_(_3_p) {
		return (((_3_p) < (Companion_Default___.HIGH__SURROGATE__MIN())) || ((_3_p) > (Companion_Default___.HIGH__SURROGATE__MAX()))) && (((_3_p) < (Companion_Default___.LOW__SURROGATE__MIN())) || ((_3_p) > (Companion_Default___.LOW__SURROGATE__MAX())))
	}
	return false
}

// Definition of class AssignedCodePoint
type AssignedCodePoint struct {
}

func New_AssignedCodePoint_() *AssignedCodePoint {
	_this := AssignedCodePoint{}

	return &_this
}

type CompanionStruct_AssignedCodePoint_ struct {
}

var Companion_AssignedCodePoint_ = CompanionStruct_AssignedCodePoint_{}

func (*AssignedCodePoint) String() string {
	return "Std_Unicode_Base.AssignedCodePoint"
}

// End of class AssignedCodePoint

func Type_AssignedCodePoint_() _dafny.TypeDescriptor {
	return type_AssignedCodePoint_{}
}

type type_AssignedCodePoint_ struct {
}

func (_this type_AssignedCodePoint_) Default() interface{} {
	return 0
}

func (_this type_AssignedCodePoint_) String() string {
	return "Std_Unicode_Base.AssignedCodePoint"
}
func (_this *CompanionStruct_AssignedCodePoint_) Is_(__source uint32) bool {
	var _4_p uint32 = (__source)
	_ = _4_p
	if Companion_CodePoint_.Is_(_4_p) {
		return Companion_Default___.IsInAssignedPlane(_4_p)
	}
	return false
}
