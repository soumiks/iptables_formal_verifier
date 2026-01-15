// Package Std_Collections_Imap
// Dafny module Std_Collections_Imap compiled into Go

package Std_Collections_Imap

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
	m_Std_Collections_Array "Std_Collections_Array"
	m_Std_Collections_Seq "Std_Collections_Seq"
	m_Std_Math "Std_Math"
	m_Std_Relations "Std_Relations"
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
	return "Std_Collections_Imap.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Get(m _dafny.Map, x interface{}) m_Std_Wrappers.Option {
	if (m).Contains(x) {
		return m_Std_Wrappers.Companion_Option_.Create_Some_((m).Get(x).(interface{}))
	} else {
		return m_Std_Wrappers.Companion_Option_.Create_None_()
	}
}

// End of class Default__
