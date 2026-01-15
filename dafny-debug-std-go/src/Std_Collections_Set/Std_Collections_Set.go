// Package Std_Collections_Set
// Dafny module Std_Collections_Set compiled into Go

package Std_Collections_Set

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
var _ m_Std_Collections_Map.Dummy__
var _ m_Std_Collections_Multiset.Dummy__

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
	return "Std_Collections_Set.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ExtractFromSingleton(s _dafny.Set) interface{} {
	return func(_let_dummy_0 int) interface{} {
		var _0_x interface{} = (interface{})(nil)
		_ = _0_x
		{
			for _iter4 := _dafny.Iterate((s).Elements()); ; {
				_assign_such_that_0, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				_0_x = interface{}(_assign_such_that_0).(interface{})
				if (s).Contains(_0_x) {
					goto L_ASSIGN_SUCH_THAT_0
				}
			}
			panic("assign-such-that search produced no value")
			goto L_ASSIGN_SUCH_THAT_0
		}
	L_ASSIGN_SUCH_THAT_0:
		return _0_x
	}(0)
}
func (_static *CompanionStruct_Default___) Map(f func(interface{}) interface{}, xs _dafny.Set) _dafny.Set {
	var _0_ys _dafny.Set = func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter5 := _dafny.Iterate((xs).Elements()); ; {
			_compr_0, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _1_x interface{}
			_1_x = interface{}(_compr_0).(interface{})
			if (xs).Contains(_1_x) {
				_coll0.Add((f)(_1_x))
			}
		}
		return _coll0.ToSet()
	}()
	_ = _0_ys
	return _0_ys
}
func (_static *CompanionStruct_Default___) Filter(f func(interface{}) bool, xs _dafny.Set) _dafny.Set {
	var _0_ys _dafny.Set = func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter6 := _dafny.Iterate((xs).Elements()); ; {
			_compr_0, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _1_x interface{}
			_1_x = interface{}(_compr_0).(interface{})
			if ((xs).Contains(_1_x)) && ((f)(_1_x)) {
				_coll0.Add(_1_x)
			}
		}
		return _coll0.ToSet()
	}()
	_ = _0_ys
	return _0_ys
}
func (_static *CompanionStruct_Default___) SetRange(a _dafny.Int, b _dafny.Int) _dafny.Set {
	var _0___accumulator _dafny.Set = _dafny.SetOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (a).Cmp(b) == 0 {
		return (_dafny.SetOf()).Union(_0___accumulator)
	} else {
		_0___accumulator = (_0___accumulator).Union(_dafny.SetOf(a))
		var _in0 _dafny.Int = (a).Plus(_dafny.One)
		_ = _in0
		var _in1 _dafny.Int = b
		_ = _in1
		a = _in0
		b = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) SetRangeZeroBound(n _dafny.Int) _dafny.Set {
	return Companion_Default___.SetRange(_dafny.Zero, n)
}

// End of class Default__
