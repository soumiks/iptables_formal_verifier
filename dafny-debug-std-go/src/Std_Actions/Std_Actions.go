// Package Std_Actions
// Dafny module Std_Actions compiled into Go

package Std_Actions

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
	m_Std_DynamicArray "Std_DynamicArray"
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
var _ m_Std_DynamicArray.Dummy__

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
	return "Std_Actions.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) InputsOf(history _dafny.Sequence) _dafny.Sequence {
	return m_Std_Collections_Seq.Companion_Default___.Map(func(coer5 func(_dafny.Tuple) interface{}) func(interface{}) interface{} {
		return func(arg6 interface{}) interface{} {
			return coer5(arg6.(_dafny.Tuple))
		}
	}(func(_0_e _dafny.Tuple) interface{} {
		return (*(_0_e).IndexInt(0))
	}), history)
}
func (_static *CompanionStruct_Default___) OutputsOf(history _dafny.Sequence) _dafny.Sequence {
	return m_Std_Collections_Seq.Companion_Default___.Map(func(coer6 func(_dafny.Tuple) interface{}) func(interface{}) interface{} {
		return func(arg7 interface{}) interface{} {
			return coer6(arg7.(_dafny.Tuple))
		}
	}(func(_0_e _dafny.Tuple) interface{} {
		return (*(_0_e).IndexInt(1))
	}), history)
}

// End of class Default__

// Definition of trait Action
type Action interface {
	String() string
	Invoke(i interface{}) interface{}
}
type CompanionStruct_Action_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_Action_ = CompanionStruct_Action_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_Action_) CastTo_(x interface{}) Action {
	var t Action
	t, _ = x.(Action)
	return t
}

// End of trait Action

// Definition of trait TotalActionProof
type TotalActionProof interface {
	String() string
}
type CompanionStruct_TotalActionProof_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_TotalActionProof_ = CompanionStruct_TotalActionProof_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_TotalActionProof_) CastTo_(x interface{}) TotalActionProof {
	var t TotalActionProof
	t, _ = x.(TotalActionProof)
	return t
}

// End of trait TotalActionProof

// Definition of class DefaultTotalActionProof
type DefaultTotalActionProof struct {
	_action Action
}

func New_DefaultTotalActionProof_() *DefaultTotalActionProof {
	_this := DefaultTotalActionProof{}

	_this._action = (Action)(nil)
	return &_this
}

type CompanionStruct_DefaultTotalActionProof_ struct {
}

var Companion_DefaultTotalActionProof_ = CompanionStruct_DefaultTotalActionProof_{}

func (_this *DefaultTotalActionProof) Equals(other *DefaultTotalActionProof) bool {
	return _this == other
}

func (_this *DefaultTotalActionProof) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*DefaultTotalActionProof)
	return ok && _this.Equals(other)
}

func (*DefaultTotalActionProof) String() string {
	return "Std_Actions.DefaultTotalActionProof"
}

func Type_DefaultTotalActionProof_(Type_I_ _dafny.TypeDescriptor, Type_O_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_DefaultTotalActionProof_{Type_I_, Type_O_}
}

type type_DefaultTotalActionProof_ struct {
	Type_I_ _dafny.TypeDescriptor
	Type_O_ _dafny.TypeDescriptor
}

func (_this type_DefaultTotalActionProof_) Default() interface{} {
	return (*DefaultTotalActionProof)(nil)
}

func (_this type_DefaultTotalActionProof_) String() string {
	return "Std_Actions.DefaultTotalActionProof"
}
func (_this *DefaultTotalActionProof) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_TotalActionProof_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ TotalActionProof = &DefaultTotalActionProof{}
var _ m_Std_Frames.Validatable = &DefaultTotalActionProof{}
var _ _dafny.TraitOffspring = &DefaultTotalActionProof{}

func (_this *DefaultTotalActionProof) Go__action() Action {
	{
		return _this._action
	}
}

// End of class DefaultTotalActionProof

// Definition of class FunctionAction
type FunctionAction struct {
	_f func(interface{}) interface{}
}

func New_FunctionAction_() *FunctionAction {
	_this := FunctionAction{}

	_this._f = (func(interface{}) interface{})(nil)
	return &_this
}

type CompanionStruct_FunctionAction_ struct {
}

var Companion_FunctionAction_ = CompanionStruct_FunctionAction_{}

func (_this *FunctionAction) Equals(other *FunctionAction) bool {
	return _this == other
}

func (_this *FunctionAction) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*FunctionAction)
	return ok && _this.Equals(other)
}

func (*FunctionAction) String() string {
	return "Std_Actions.FunctionAction"
}

func Type_FunctionAction_(Type_I_ _dafny.TypeDescriptor, Type_O_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_FunctionAction_{Type_I_, Type_O_}
}

type type_FunctionAction_ struct {
	Type_I_ _dafny.TypeDescriptor
	Type_O_ _dafny.TypeDescriptor
}

func (_this type_FunctionAction_) Default() interface{} {
	return (*FunctionAction)(nil)
}

func (_this type_FunctionAction_) String() string {
	return "Std_Actions.FunctionAction"
}
func (_this *FunctionAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Action = &FunctionAction{}
var _ m_Std_GenericActions.GenericAction = &FunctionAction{}
var _ m_Std_Frames.Validatable = &FunctionAction{}
var _ _dafny.TraitOffspring = &FunctionAction{}

func (_this *FunctionAction) Ctor__(f func(interface{}) interface{}) {
	{
		(_this)._f = f
	}
}
func (_this *FunctionAction) Invoke(i interface{}) interface{} {
	{
		var i interface{} = i
		_ = i
		var o interface{} = (interface{})(nil)
		_ = o
		o = ((_this).F())(i)
		return o
	}
}
func (_this *FunctionAction) F() func(interface{}) interface{} {
	{
		return _this._f
	}
}

// End of class FunctionAction

// Definition of class TotalFunctionActionProof
type TotalFunctionActionProof struct {
	dummy byte
}

func New_TotalFunctionActionProof_() *TotalFunctionActionProof {
	_this := TotalFunctionActionProof{}

	return &_this
}

type CompanionStruct_TotalFunctionActionProof_ struct {
}

var Companion_TotalFunctionActionProof_ = CompanionStruct_TotalFunctionActionProof_{}

func (_this *TotalFunctionActionProof) Equals(other *TotalFunctionActionProof) bool {
	return _this == other
}

func (_this *TotalFunctionActionProof) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*TotalFunctionActionProof)
	return ok && _this.Equals(other)
}

func (*TotalFunctionActionProof) String() string {
	return "Std_Actions.TotalFunctionActionProof"
}

func Type_TotalFunctionActionProof_(Type_I_ _dafny.TypeDescriptor, Type_O_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_TotalFunctionActionProof_{Type_I_, Type_O_}
}

type type_TotalFunctionActionProof_ struct {
	Type_I_ _dafny.TypeDescriptor
	Type_O_ _dafny.TypeDescriptor
}

func (_this type_TotalFunctionActionProof_) Default() interface{} {
	return (*TotalFunctionActionProof)(nil)
}

func (_this type_TotalFunctionActionProof_) String() string {
	return "Std_Actions.TotalFunctionActionProof"
}
func (_this *TotalFunctionActionProof) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_TotalActionProof_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ TotalActionProof = &TotalFunctionActionProof{}
var _ m_Std_Frames.Validatable = &TotalFunctionActionProof{}
var _ _dafny.TraitOffspring = &TotalFunctionActionProof{}

// End of class TotalFunctionActionProof

// Definition of class ComposedAction
type ComposedAction struct {
	_first  Action
	_second Action
}

func New_ComposedAction_() *ComposedAction {
	_this := ComposedAction{}

	_this._first = (Action)(nil)
	_this._second = (Action)(nil)
	return &_this
}

type CompanionStruct_ComposedAction_ struct {
}

var Companion_ComposedAction_ = CompanionStruct_ComposedAction_{}

func (_this *ComposedAction) Equals(other *ComposedAction) bool {
	return _this == other
}

func (_this *ComposedAction) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*ComposedAction)
	return ok && _this.Equals(other)
}

func (*ComposedAction) String() string {
	return "Std_Actions.ComposedAction"
}

func Type_ComposedAction_(Type_I_ _dafny.TypeDescriptor, Type_M_ _dafny.TypeDescriptor, Type_O_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_ComposedAction_{Type_I_, Type_M_, Type_O_}
}

type type_ComposedAction_ struct {
	Type_I_ _dafny.TypeDescriptor
	Type_M_ _dafny.TypeDescriptor
	Type_O_ _dafny.TypeDescriptor
}

func (_this type_ComposedAction_) Default() interface{} {
	return (*ComposedAction)(nil)
}

func (_this type_ComposedAction_) String() string {
	return "Std_Actions.ComposedAction"
}
func (_this *ComposedAction) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Action = &ComposedAction{}
var _ m_Std_GenericActions.GenericAction = &ComposedAction{}
var _ m_Std_Frames.Validatable = &ComposedAction{}
var _ _dafny.TraitOffspring = &ComposedAction{}

func (_this *ComposedAction) Ctor__(first Action, second Action) {
	{
		(_this)._first = first
		(_this)._second = second
	}
}
func (_this *ComposedAction) Invoke(i interface{}) interface{} {
	{
		var i interface{} = i
		_ = i
		var o interface{} = (interface{})(nil)
		_ = o
		var _0_m interface{}
		_ = _0_m
		var _out0 interface{}
		_ = _out0
		_out0 = ((_this).First()).Invoke(i)
		_0_m = _out0
		var _out1 interface{}
		_ = _out1
		_out1 = ((_this).Second()).Invoke(_0_m)
		o = _out1
		return o
	}
}
func (_this *ComposedAction) First() Action {
	{
		return _this._first
	}
}
func (_this *ComposedAction) Second() Action {
	{
		return _this._second
	}
}

// End of class ComposedAction

// Definition of trait ActionCompositionProof
type ActionCompositionProof interface {
	String() string
}
type CompanionStruct_ActionCompositionProof_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_ActionCompositionProof_ = CompanionStruct_ActionCompositionProof_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_ActionCompositionProof_) CastTo_(x interface{}) ActionCompositionProof {
	var t ActionCompositionProof
	t, _ = x.(ActionCompositionProof)
	return t
}

// End of trait ActionCompositionProof
