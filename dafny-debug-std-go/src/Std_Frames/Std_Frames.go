// Package Std_Frames
// Dafny module Std_Frames compiled into Go

package Std_Frames

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
	m_Std_Collections_Array "Std_Collections_Array"
	m_Std_Collections_Imap "Std_Collections_Imap"
	m_Std_Collections_Iset "Std_Collections_Iset"
	m_Std_Collections_Map "Std_Collections_Map"
	m_Std_Collections_Multiset "Std_Collections_Multiset"
	m_Std_Collections_Seq "Std_Collections_Seq"
	m_Std_Collections_Set "Std_Collections_Set"
	m_Std_Collections_Tuple "Std_Collections_Tuple"
	m_Std_Functions "Std_Functions"
	m_Std_Math "Std_Math"
	m_Std_Ordinal "Std_Ordinal"
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

type Dummy__ struct{}

// Definition of trait Validatable
type Validatable interface {
	String() string
}
type CompanionStruct_Validatable_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_Validatable_ = CompanionStruct_Validatable_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_Validatable_) CastTo_(x interface{}) Validatable {
	var t Validatable
	t, _ = x.(Validatable)
	return t
}

// End of trait Validatable

// Definition of class GhostBox
type GhostBox struct {
	dummy byte
}

func New_GhostBox_() *GhostBox {
	_this := GhostBox{}

	return &_this
}

type CompanionStruct_GhostBox_ struct {
}

var Companion_GhostBox_ = CompanionStruct_GhostBox_{}

func (_this *GhostBox) Equals(other *GhostBox) bool {
	return _this == other
}

func (_this *GhostBox) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*GhostBox)
	return ok && _this.Equals(other)
}

func (*GhostBox) String() string {
	return "Std_Frames.GhostBox"
}

func Type_GhostBox_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_GhostBox_{Type_T_}
}

type type_GhostBox_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_GhostBox_) Default() interface{} {
	return (*GhostBox)(nil)
}

func (_this type_GhostBox_) String() string {
	return "Std_Frames.GhostBox"
}
func (_this *GhostBox) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &GhostBox{}

// End of class GhostBox

// Definition of class Box
type Box struct {
	Value interface{}
}

func New_Box_() *Box {
	_this := Box{}

	_this.Value = (interface{})(nil)
	return &_this
}

type CompanionStruct_Box_ struct {
}

var Companion_Box_ = CompanionStruct_Box_{}

func (_this *Box) Equals(other *Box) bool {
	return _this == other
}

func (_this *Box) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Box)
	return ok && _this.Equals(other)
}

func (*Box) String() string {
	return "Std_Frames.Box"
}

func Type_Box_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_Box_{Type_T_}
}

type type_Box_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_Box_) Default() interface{} {
	return (*Box)(nil)
}

func (_this type_Box_) String() string {
	return "Std_Frames.Box"
}
func (_this *Box) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Box{}

func (_this *Box) Ctor__(value interface{}) {
	{
		(_this).Value = value
	}
}

// End of class Box
