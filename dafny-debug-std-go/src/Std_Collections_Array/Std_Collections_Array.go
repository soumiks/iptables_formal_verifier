// Package Std_Collections_Array
// Dafny module Std_Collections_Array compiled into Go

package Std_Collections_Array

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
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
	return "Std_Collections_Array.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) BinarySearch(a _dafny.Array, key interface{}, less func(interface{}, interface{}) bool) m_Std_Wrappers.Option {
	var r m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
	_ = r
	var _0_lo _dafny.Int
	_ = _0_lo
	var _1_hi _dafny.Int
	_ = _1_hi
	var _rhs0 _dafny.Int = _dafny.Zero
	_ = _rhs0
	var _rhs1 _dafny.Int = _dafny.ArrayLen((a), 0)
	_ = _rhs1
	_0_lo = _rhs0
	_1_hi = _rhs1
	for (_0_lo).Cmp(_1_hi) < 0 {
		var _2_mid _dafny.Int
		_ = _2_mid
		_2_mid = ((_0_lo).Plus(_1_hi)).DivBy(_dafny.IntOfInt64(2))
		if (less)(key, (a).ArrayGet1((_2_mid).Int())) {
			_1_hi = _2_mid
		} else if (less)((a).ArrayGet1((_2_mid).Int()), key) {
			_0_lo = (_2_mid).Plus(_dafny.One)
		} else {
			r = m_Std_Wrappers.Companion_Option_.Create_Some_(_2_mid)
			return r
		}
	}
	r = m_Std_Wrappers.Companion_Option_.Create_None_()
	return r
	return r
}

// End of class Default__
