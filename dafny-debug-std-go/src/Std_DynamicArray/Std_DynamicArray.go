// Package Std_DynamicArray
// Dafny module Std_DynamicArray compiled into Go

package Std_DynamicArray

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
	m_Std_BoundedInts "Std_BoundedInts"
	m_Std_Collections_Array "Std_Collections_Array"
	m_Std_Collections_Imap "Std_Collections_Imap"
	m_Std_Collections_Iset "Std_Collections_Iset"
	m_Std_Collections_Map "Std_Collections_Map"
	m_Std_Collections_Multiset "Std_Collections_Multiset"
	m_Std_Collections_Seq "Std_Collections_Seq"
	m_Std_Collections_Set "Std_Collections_Set"
	m_Std_Collections_Tuple "Std_Collections_Tuple"
	m_Std_Frames "Std_Frames"
	m_Std_Functions "Std_Functions"
	m_Std_GenericActions "Std_GenericActions"
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
var _ m_Std_Frames.Dummy__
var _ m_Std_GenericActions.Dummy__
var _ m_Std_BoundedInts.Dummy__

type Dummy__ struct{}

// Definition of class DynamicArray
type DynamicArray struct {
	Size     _dafny.Int
	Capacity _dafny.Int
	Data     _dafny.Array
}

func New_DynamicArray_() *DynamicArray {
	_this := DynamicArray{}

	_this.Size = _dafny.Zero
	_this.Capacity = _dafny.Zero
	_this.Data = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	return &_this
}

type CompanionStruct_DynamicArray_ struct {
}

var Companion_DynamicArray_ = CompanionStruct_DynamicArray_{}

func (_this *DynamicArray) Equals(other *DynamicArray) bool {
	return _this == other
}

func (_this *DynamicArray) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*DynamicArray)
	return ok && _this.Equals(other)
}

func (*DynamicArray) String() string {
	return "Std_DynamicArray.DynamicArray"
}

func Type_DynamicArray_(Type_A_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_DynamicArray_{Type_A_}
}

type type_DynamicArray_ struct {
	Type_A_ _dafny.TypeDescriptor
}

func (_this type_DynamicArray_) Default() interface{} {
	return (*DynamicArray)(nil)
}

func (_this type_DynamicArray_) String() string {
	return "Std_DynamicArray.DynamicArray"
}
func (_this *DynamicArray) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &DynamicArray{}

func (_this *DynamicArray) Ctor__() {
	{
		(_this).Size = _dafny.Zero
		(_this).Capacity = _dafny.Zero
		var _nw0 _dafny.Array = _dafny.NewArray(_dafny.Zero)
		_ = _nw0
		(_this).Data = _nw0
	}
}
func (_this *DynamicArray) At(index _dafny.Int) interface{} {
	{
		return (_this.Data).ArrayGet1((index).Int())
	}
}
func (_this *DynamicArray) Put(index _dafny.Int, element interface{}) {
	{
		var _arr0 _dafny.Array = _this.Data
		_ = _arr0
		(_arr0).ArraySet1(element, (index).Int())
	}
}
func (_this *DynamicArray) Ensure(reserved _dafny.Int, defaultValue interface{}) {
	{
		var _0_newCapacity _dafny.Int
		_ = _0_newCapacity
		_0_newCapacity = _this.Capacity
		for (reserved).Cmp((_0_newCapacity).Minus(_this.Size)) > 0 {
			_0_newCapacity = (_this).DefaultNewCapacity(_0_newCapacity)
		}
		if (_0_newCapacity).Cmp(_this.Capacity) > 0 {
			(_this).Realloc(defaultValue, _0_newCapacity)
		}
	}
}
func (_this *DynamicArray) PopFast() {
	{
		(_this).Size = (_this.Size).Minus(_dafny.One)
	}
}
func (_this *DynamicArray) PushFast(element interface{}) {
	{
		var _arr0 _dafny.Array = _this.Data
		_ = _arr0
		var _index0 _dafny.Int = _this.Size
		_ = _index0
		(_arr0).ArraySet1(element, (_index0).Int())
		(_this).Size = (_this.Size).Plus(_dafny.One)
	}
}
func (_this *DynamicArray) Push(element interface{}) {
	{
		if (_this.Size).Cmp(_this.Capacity) == 0 {
			(_this).ReallocDefault(element)
		}
		(_this).PushFast(element)
	}
}
func (_this *DynamicArray) Realloc(defaultValue interface{}, newCapacity _dafny.Int) {
	{
		var _0_oldData _dafny.Array
		_ = _0_oldData
		var _1_oldCapacity _dafny.Int
		_ = _1_oldCapacity
		var _rhs0 _dafny.Array = _this.Data
		_ = _rhs0
		var _rhs1 _dafny.Int = _this.Capacity
		_ = _rhs1
		_0_oldData = _rhs0
		_1_oldCapacity = _rhs1
		var _len0_0 _dafny.Int = newCapacity
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) interface{} = (func(_2_defaultValue interface{}) func(_dafny.Int) interface{} {
				return func(_3___v0 _dafny.Int) interface{} {
					return _2_defaultValue
				}
			})(defaultValue)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw0).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw0).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		var _rhs2 _dafny.Array = _nw0
		_ = _rhs2
		var _rhs3 _dafny.Int = newCapacity
		_ = _rhs3
		var _lhs0 *DynamicArray = _this
		_ = _lhs0
		var _lhs1 *DynamicArray = _this
		_ = _lhs1
		_lhs0.Data = _rhs2
		_lhs1.Capacity = _rhs3
		(_this).CopyFrom(_0_oldData, _1_oldCapacity)
	}
}
func (_this *DynamicArray) DefaultNewCapacity(capacity _dafny.Int) _dafny.Int {
	{
		if (capacity).Sign() == 0 {
			return _dafny.IntOfInt64(8)
		} else {
			return (_dafny.IntOfInt64(2)).Times(capacity)
		}
	}
}
func (_this *DynamicArray) ReallocDefault(defaultValue interface{}) {
	{
		(_this).Realloc(defaultValue, (_this).DefaultNewCapacity(_this.Capacity))
	}
}
func (_this *DynamicArray) CopyFrom(newData _dafny.Array, count _dafny.Int) {
	{
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, count)); ; {
			_guard_loop_0, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _0_index _dafny.Int
			_0_index = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_0_index).Sign() != -1) && ((_0_index).Cmp(count) < 0)) {
				var _arr0 _dafny.Array = _this.Data
				_ = _arr0
				(_arr0).ArraySet1((newData).ArrayGet1((_0_index).Int()), (_0_index).Int())
			}
		}
	}
}

// End of class DynamicArray
