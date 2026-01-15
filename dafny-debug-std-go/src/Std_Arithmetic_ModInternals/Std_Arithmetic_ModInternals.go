// Package Std_Arithmetic_ModInternals
// Dafny module Std_Arithmetic_ModInternals compiled into Go

package Std_Arithmetic_ModInternals

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
	m_Std_Actions "Std_Actions"
	m_Std_ActionsExterns "Std_ActionsExterns"
	m_Std_Arithmetic_DivInternalsNonlinear "Std_Arithmetic_DivInternalsNonlinear"
	m_Std_Arithmetic_GeneralInternals "Std_Arithmetic_GeneralInternals"
	m_Std_Arithmetic_ModInternalsNonlinear "Std_Arithmetic_ModInternalsNonlinear"
	m_Std_Arithmetic_Mul "Std_Arithmetic_Mul"
	m_Std_Arithmetic_MulInternals "Std_Arithmetic_MulInternals"
	m_Std_Arithmetic_MulInternalsNonlinear "Std_Arithmetic_MulInternalsNonlinear"
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
	m_Std_Unicode_Base "Std_Unicode_Base"
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
	return "Std_Arithmetic_ModInternals.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ModRecursive(x _dafny.Int, d _dafny.Int) _dafny.Int {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (x).Sign() == -1 {
		var _in0 _dafny.Int = (d).Plus(x)
		_ = _in0
		var _in1 _dafny.Int = d
		_ = _in1
		x = _in0
		d = _in1
		goto TAIL_CALL_START
	} else if (x).Cmp(d) < 0 {
		return x
	} else {
		var _in2 _dafny.Int = (x).Minus(d)
		_ = _in2
		var _in3 _dafny.Int = d
		_ = _in3
		x = _in2
		d = _in3
		goto TAIL_CALL_START
	}
}

// End of class Default__
