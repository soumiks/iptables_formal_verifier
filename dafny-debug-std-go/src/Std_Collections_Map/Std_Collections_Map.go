// Package Std_Collections_Map
// Dafny module Std_Collections_Map compiled into Go

package Std_Collections_Map

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
	m_Std_Collections_Array "Std_Collections_Array"
	m_Std_Collections_Imap "Std_Collections_Imap"
	m_Std_Collections_Iset "Std_Collections_Iset"
	m_Std_Collections_Seq "Std_Collections_Seq"
	m_Std_Functions "Std_Functions"
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
var _ m_Std_Collections_Imap.Dummy__
var _ m_Std_Functions.Dummy__
var _ m_Std_Collections_Iset.Dummy__

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
	return "Std_Collections_Map.Default__"
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
func (_static *CompanionStruct_Default___) ToImap(m _dafny.Map) _dafny.Map {
	return func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter1 := _dafny.Iterate((m).Keys().Elements()); ; {
			_compr_0, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _0_x interface{}
			_0_x = interface{}(_compr_0).(interface{})
			if (m).Contains(_0_x) {
				_coll0.Add(_0_x, (m).Get(_0_x).(interface{}))
			}
		}
		return _coll0.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) RemoveKeys(m _dafny.Map, xs _dafny.Set) _dafny.Map {
	return (m).Subtract(xs)
}
func (_static *CompanionStruct_Default___) Remove(m _dafny.Map, x interface{}) _dafny.Map {
	var _0_m_k _dafny.Map = func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter2 := _dafny.Iterate((m).Keys().Elements()); ; {
			_compr_0, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _1_x_k interface{}
			_1_x_k = interface{}(_compr_0).(interface{})
			if ((m).Contains(_1_x_k)) && (!_dafny.AreEqual(_1_x_k, x)) {
				_coll0.Add(_1_x_k, (m).Get(_1_x_k).(interface{}))
			}
		}
		return _coll0.ToMap()
	}()
	_ = _0_m_k
	return _0_m_k
}
func (_static *CompanionStruct_Default___) Restrict(m _dafny.Map, xs _dafny.Set) _dafny.Map {
	return func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter3 := _dafny.Iterate((xs).Elements()); ; {
			_compr_0, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _0_x interface{}
			_0_x = interface{}(_compr_0).(interface{})
			if ((xs).Contains(_0_x)) && ((m).Contains(_0_x)) {
				_coll0.Add(_0_x, (m).Get(_0_x).(interface{}))
			}
		}
		return _coll0.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Union(m _dafny.Map, m_k _dafny.Map) _dafny.Map {
	return (m).Merge(m_k)
}

// End of class Default__
