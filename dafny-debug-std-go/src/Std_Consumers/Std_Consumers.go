// Package Std_Consumers
// Dafny module Std_Consumers compiled into Go

package Std_Consumers

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
	m_Std_Actions "Std_Actions"
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
var _ m_Std_Actions.Dummy__

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
	return "Std_Consumers.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) IsTrue(b bool) bool {
	return (b) == (true)
}
func (_static *CompanionStruct_Default___) IsFalse(b bool) bool {
	return (b) == (false)
}
func (_static *CompanionStruct_Default___) WasConsumed(pair _dafny.Tuple) bool {
	return (*(pair).IndexInt(1)).(bool)
}
func (_static *CompanionStruct_Default___) WasNotConsumed(pair _dafny.Tuple) bool {
	return !((*(pair).IndexInt(1)).(bool))
}

// End of class Default__

// Definition of trait IConsumer
type IConsumer interface {
	String() string
	Invoke(i interface{}) interface{}
	Accept(t interface{})
}

func (_static *CompanionStruct_IConsumer_) Accept(_this IConsumer, t interface{}) {
	{
		var _0_r _dafny.Tuple
		_ = _0_r
		var _out0 interface{}
		_ = _out0
		_out0 = (_this).Invoke(t)
		_0_r = _out0.(_dafny.Tuple)
	}
}

type CompanionStruct_IConsumer_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_IConsumer_ = CompanionStruct_IConsumer_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_IConsumer_) CastTo_(x interface{}) IConsumer {
	var t IConsumer
	t, _ = x.(IConsumer)
	return t
}

// End of trait IConsumer

// Definition of datatype ConsumerState
type ConsumerState struct {
	Data_ConsumerState_
}

func (_this ConsumerState) Get_() Data_ConsumerState_ {
	return _this.Data_ConsumerState_
}

type Data_ConsumerState_ interface {
	isConsumerState()
}

type CompanionStruct_ConsumerState_ struct {
}

var Companion_ConsumerState_ = CompanionStruct_ConsumerState_{}

type ConsumerState_ConsumerState struct {
	Consumer Consumer
	Capacity m_Std_Wrappers.Option
	History  _dafny.Sequence
}

func (ConsumerState_ConsumerState) isConsumerState() {}

func (CompanionStruct_ConsumerState_) Create_ConsumerState_(Consumer Consumer, Capacity m_Std_Wrappers.Option, History _dafny.Sequence) ConsumerState {
	return ConsumerState{ConsumerState_ConsumerState{Consumer, Capacity, History}}
}

func (_this ConsumerState) Is_ConsumerState() bool {
	_, ok := _this.Get_().(ConsumerState_ConsumerState)
	return ok
}

func (CompanionStruct_ConsumerState_) Default() ConsumerState {
	return Companion_ConsumerState_.Create_ConsumerState_((Consumer)(nil), m_Std_Wrappers.Companion_Option_.Default(), _dafny.EmptySeq)
}

func (_this ConsumerState) Dtor_consumer() Consumer {
	return _this.Get_().(ConsumerState_ConsumerState).Consumer
}

func (_this ConsumerState) Dtor_capacity() m_Std_Wrappers.Option {
	return _this.Get_().(ConsumerState_ConsumerState).Capacity
}

func (_this ConsumerState) Dtor_history() _dafny.Sequence {
	return _this.Get_().(ConsumerState_ConsumerState).History
}

func (_this ConsumerState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ConsumerState_ConsumerState:
		{
			return "Consumers.ConsumerState.ConsumerState" + "(" + _dafny.String(data.Consumer) + ", " + _dafny.String(data.Capacity) + ", " + _dafny.String(data.History) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ConsumerState) Equals(other ConsumerState) bool {
	switch data1 := _this.Get_().(type) {
	case ConsumerState_ConsumerState:
		{
			data2, ok := other.Get_().(ConsumerState_ConsumerState)
			return ok && _dafny.AreEqual(data1.Consumer, data2.Consumer) && data1.Capacity.Equals(data2.Capacity) && data1.History.Equals(data2.History)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ConsumerState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ConsumerState)
	return ok && _this.Equals(typed)
}

func Type_ConsumerState_() _dafny.TypeDescriptor {
	return type_ConsumerState_{}
}

type type_ConsumerState_ struct {
}

func (_this type_ConsumerState_) Default() interface{} {
	return Companion_ConsumerState_.Default()
}

func (_this type_ConsumerState_) String() string {
	return "Std_Consumers.ConsumerState"
}
func (_this ConsumerState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ConsumerState{}

// End of datatype ConsumerState

// Definition of trait Consumer
type Consumer interface {
	String() string
	Capacity() m_Std_Wrappers.Option
	Invoke(i interface{}) interface{}
	Accept(t interface{}) bool
}

func (_static *CompanionStruct_Consumer_) Accept(_this Consumer, t interface{}) bool {
	{
		var o bool = false
		_ = o
		var _out0 interface{}
		_ = _out0
		_out0 = (_this).Invoke(t)
		o = _out0.(bool)
		return o
	}
}

type CompanionStruct_Consumer_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_Consumer_ = CompanionStruct_Consumer_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_Consumer_) CastTo_(x interface{}) Consumer {
	var t Consumer
	t, _ = x.(Consumer)
	return t
}

// End of trait Consumer

// Definition of class IgnoreNConsumer
type IgnoreNConsumer struct {
	ConsumedCount _dafny.Int
	_n            _dafny.Int
}

func New_IgnoreNConsumer_() *IgnoreNConsumer {
	_this := IgnoreNConsumer{}

	_this.ConsumedCount = _dafny.Zero
	_this._n = _dafny.Zero
	return &_this
}

type CompanionStruct_IgnoreNConsumer_ struct {
}

var Companion_IgnoreNConsumer_ = CompanionStruct_IgnoreNConsumer_{}

func (_this *IgnoreNConsumer) Equals(other *IgnoreNConsumer) bool {
	return _this == other
}

func (_this *IgnoreNConsumer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*IgnoreNConsumer)
	return ok && _this.Equals(other)
}

func (*IgnoreNConsumer) String() string {
	return "Std_Consumers.IgnoreNConsumer"
}

func Type_IgnoreNConsumer_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_IgnoreNConsumer_{Type_T_}
}

type type_IgnoreNConsumer_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_IgnoreNConsumer_) Default() interface{} {
	return (*IgnoreNConsumer)(nil)
}

func (_this type_IgnoreNConsumer_) String() string {
	return "Std_Consumers.IgnoreNConsumer"
}
func (_this *IgnoreNConsumer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Consumer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Consumer = &IgnoreNConsumer{}
var _ m_Std_Actions.Action = &IgnoreNConsumer{}
var _ m_Std_GenericActions.GenericAction = &IgnoreNConsumer{}
var _ m_Std_Frames.Validatable = &IgnoreNConsumer{}
var _ _dafny.TraitOffspring = &IgnoreNConsumer{}

func (_this *IgnoreNConsumer) Accept(t interface{}) bool {
	var _out1 bool
	_ = _out1
	_out1 = Companion_Consumer_.Accept(_this, t)
	return _out1
}
func (_this *IgnoreNConsumer) Ctor__(n _dafny.Int) {
	{
		(_this)._n = n
		(_this).ConsumedCount = _dafny.Zero
	}
}
func (_this *IgnoreNConsumer) Capacity() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_Some_(((_this).N()).Minus(_this.ConsumedCount))
	}
}
func (_this *IgnoreNConsumer) Invoke(t interface{}) interface{} {
	{
		var t interface{} = t
		_ = t
		var r bool = false
		_ = r
		if (_this.ConsumedCount).Cmp((_this).N()) == 0 {
			r = false
		} else {
			r = true
			(_this).ConsumedCount = (_this.ConsumedCount).Plus(_dafny.One)
		}
		return r
	}
}
func (_this *IgnoreNConsumer) N() _dafny.Int {
	{
		return _this._n
	}
}

// End of class IgnoreNConsumer

// Definition of class ArrayWriter
type ArrayWriter struct {
	Size     _dafny.Int
	_storage _dafny.Array
}

func New_ArrayWriter_() *ArrayWriter {
	_this := ArrayWriter{}

	_this.Size = _dafny.Zero
	_this._storage = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	return &_this
}

type CompanionStruct_ArrayWriter_ struct {
}

var Companion_ArrayWriter_ = CompanionStruct_ArrayWriter_{}

func (_this *ArrayWriter) Equals(other *ArrayWriter) bool {
	return _this == other
}

func (_this *ArrayWriter) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*ArrayWriter)
	return ok && _this.Equals(other)
}

func (*ArrayWriter) String() string {
	return "Std_Consumers.ArrayWriter"
}

func Type_ArrayWriter_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_ArrayWriter_{Type_T_}
}

type type_ArrayWriter_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_ArrayWriter_) Default() interface{} {
	return (*ArrayWriter)(nil)
}

func (_this type_ArrayWriter_) String() string {
	return "Std_Consumers.ArrayWriter"
}
func (_this *ArrayWriter) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Consumer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Consumer = &ArrayWriter{}
var _ m_Std_Actions.Action = &ArrayWriter{}
var _ m_Std_GenericActions.GenericAction = &ArrayWriter{}
var _ m_Std_Frames.Validatable = &ArrayWriter{}
var _ _dafny.TraitOffspring = &ArrayWriter{}

func (_this *ArrayWriter) Accept(t interface{}) bool {
	var _out2 bool
	_ = _out2
	_out2 = Companion_Consumer_.Accept(_this, t)
	return _out2
}
func (_this *ArrayWriter) Ctor__(storage _dafny.Array) {
	{
		(_this)._storage = storage
		(_this).Size = _dafny.Zero
	}
}
func (_this *ArrayWriter) Capacity() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_Some_((_dafny.ArrayLen(((_this).Storage()), 0)).Minus(_this.Size))
	}
}
func (_this *ArrayWriter) Invoke(t interface{}) interface{} {
	{
		var t interface{} = t
		_ = t
		var r bool = false
		_ = r
		if (_this.Size).Cmp(_dafny.ArrayLen(((_this).Storage()), 0)) == 0 {
			r = false
		} else {
			var _index0 _dafny.Int = _this.Size
			_ = _index0
			((_this).Storage()).ArraySet1(t, (_index0).Int())
			(_this).Size = (_this.Size).Plus(_dafny.One)
			r = true
		}
		return r
	}
}
func (_this *ArrayWriter) Storage() _dafny.Array {
	{
		return _this._storage
	}
}

// End of class ArrayWriter

// Definition of class DynamicArrayWriter
type DynamicArrayWriter struct {
	Storage *m_Std_DynamicArray.DynamicArray
}

func New_DynamicArrayWriter_() *DynamicArrayWriter {
	_this := DynamicArrayWriter{}

	_this.Storage = (*m_Std_DynamicArray.DynamicArray)(nil)
	return &_this
}

type CompanionStruct_DynamicArrayWriter_ struct {
}

var Companion_DynamicArrayWriter_ = CompanionStruct_DynamicArrayWriter_{}

func (_this *DynamicArrayWriter) Equals(other *DynamicArrayWriter) bool {
	return _this == other
}

func (_this *DynamicArrayWriter) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*DynamicArrayWriter)
	return ok && _this.Equals(other)
}

func (*DynamicArrayWriter) String() string {
	return "Std_Consumers.DynamicArrayWriter"
}

func Type_DynamicArrayWriter_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_DynamicArrayWriter_{Type_T_}
}

type type_DynamicArrayWriter_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_DynamicArrayWriter_) Default() interface{} {
	return (*DynamicArrayWriter)(nil)
}

func (_this type_DynamicArrayWriter_) String() string {
	return "Std_Consumers.DynamicArrayWriter"
}
func (_this *DynamicArrayWriter) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_IConsumer_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ IConsumer = &DynamicArrayWriter{}
var _ m_Std_Actions.TotalActionProof = &DynamicArrayWriter{}
var _ m_Std_Actions.Action = &DynamicArrayWriter{}
var _ m_Std_GenericActions.GenericAction = &DynamicArrayWriter{}
var _ m_Std_Frames.Validatable = &DynamicArrayWriter{}
var _ _dafny.TraitOffspring = &DynamicArrayWriter{}

func (_this *DynamicArrayWriter) Accept(t interface{}) {
	Companion_IConsumer_.Accept(_this, t)
	return
}
func (_this *DynamicArrayWriter) Ctor__() {
	{
		var _0_a *m_Std_DynamicArray.DynamicArray
		_ = _0_a
		var _nw0 *m_Std_DynamicArray.DynamicArray = m_Std_DynamicArray.New_DynamicArray_()
		_ = _nw0
		_nw0.Ctor__()
		_0_a = _nw0
		(_this).Storage = _0_a
	}
}
func (_this *DynamicArrayWriter) Invoke(t interface{}) interface{} {
	{
		var t interface{} = t
		_ = t
		var r _dafny.Tuple = _dafny.TupleOf()
		_ = r
		(_this.Storage).Push(t)
		r = _dafny.TupleOf()
		return r
	}
}

// End of class DynamicArrayWriter

// Definition of class FoldingConsumer
type FoldingConsumer struct {
	Value interface{}
	_f    func(interface{}, interface{}) interface{}
}

func New_FoldingConsumer_() *FoldingConsumer {
	_this := FoldingConsumer{}

	_this.Value = (interface{})(nil)
	_this._f = func(interface{}, interface{}) interface{} { return (interface{})(nil) }
	return &_this
}

type CompanionStruct_FoldingConsumer_ struct {
}

var Companion_FoldingConsumer_ = CompanionStruct_FoldingConsumer_{}

func (_this *FoldingConsumer) Equals(other *FoldingConsumer) bool {
	return _this == other
}

func (_this *FoldingConsumer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*FoldingConsumer)
	return ok && _this.Equals(other)
}

func (*FoldingConsumer) String() string {
	return "Std_Consumers.FoldingConsumer"
}

func Type_FoldingConsumer_(Type_T_ _dafny.TypeDescriptor, Type_R_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_FoldingConsumer_{Type_T_, Type_R_}
}

type type_FoldingConsumer_ struct {
	Type_T_ _dafny.TypeDescriptor
	Type_R_ _dafny.TypeDescriptor
}

func (_this type_FoldingConsumer_) Default() interface{} {
	return (*FoldingConsumer)(nil)
}

func (_this type_FoldingConsumer_) String() string {
	return "Std_Consumers.FoldingConsumer"
}
func (_this *FoldingConsumer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_IConsumer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ IConsumer = &FoldingConsumer{}
var _ m_Std_Actions.Action = &FoldingConsumer{}
var _ m_Std_GenericActions.GenericAction = &FoldingConsumer{}
var _ m_Std_Frames.Validatable = &FoldingConsumer{}
var _ _dafny.TraitOffspring = &FoldingConsumer{}

func (_this *FoldingConsumer) Accept(t interface{}) {
	Companion_IConsumer_.Accept(_this, t)
	return
}
func (_this *FoldingConsumer) Ctor__(init interface{}, f func(interface{}, interface{}) interface{}) {
	{
		(_this)._f = f
		(_this).Value = init
	}
}
func (_this *FoldingConsumer) Invoke(t interface{}) interface{} {
	{
		var t interface{} = t
		_ = t
		var r _dafny.Tuple = _dafny.TupleOf()
		_ = r
		(_this).Value = ((_this).F())(_this.Value, t)
		r = _dafny.TupleOf()
		return r
	}
}
func (_this *FoldingConsumer) F() func(interface{}, interface{}) interface{} {
	{
		return _this._f
	}
}

// End of class FoldingConsumer

// Definition of class FoldingConsumerTotalActionProof
type FoldingConsumerTotalActionProof struct {
	dummy byte
}

func New_FoldingConsumerTotalActionProof_() *FoldingConsumerTotalActionProof {
	_this := FoldingConsumerTotalActionProof{}

	return &_this
}

type CompanionStruct_FoldingConsumerTotalActionProof_ struct {
}

var Companion_FoldingConsumerTotalActionProof_ = CompanionStruct_FoldingConsumerTotalActionProof_{}

func (_this *FoldingConsumerTotalActionProof) Equals(other *FoldingConsumerTotalActionProof) bool {
	return _this == other
}

func (_this *FoldingConsumerTotalActionProof) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*FoldingConsumerTotalActionProof)
	return ok && _this.Equals(other)
}

func (*FoldingConsumerTotalActionProof) String() string {
	return "Std_Consumers.FoldingConsumerTotalActionProof"
}

func Type_FoldingConsumerTotalActionProof_(Type_I_ _dafny.TypeDescriptor, Type_O_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_FoldingConsumerTotalActionProof_{Type_I_, Type_O_}
}

type type_FoldingConsumerTotalActionProof_ struct {
	Type_I_ _dafny.TypeDescriptor
	Type_O_ _dafny.TypeDescriptor
}

func (_this type_FoldingConsumerTotalActionProof_) Default() interface{} {
	return (*FoldingConsumerTotalActionProof)(nil)
}

func (_this type_FoldingConsumerTotalActionProof_) String() string {
	return "Std_Consumers.FoldingConsumerTotalActionProof"
}
func (_this *FoldingConsumerTotalActionProof) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ m_Std_Actions.TotalActionProof = &FoldingConsumerTotalActionProof{}
var _ m_Std_Frames.Validatable = &FoldingConsumerTotalActionProof{}
var _ _dafny.TraitOffspring = &FoldingConsumerTotalActionProof{}

func (_this *FoldingConsumerTotalActionProof) Ctor__(action *FoldingConsumer) {
	{
	}
}

// End of class FoldingConsumerTotalActionProof

// Definition of class SeqWriter
type SeqWriter struct {
	Values _dafny.Sequence
}

func New_SeqWriter_() *SeqWriter {
	_this := SeqWriter{}

	_this.Values = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_SeqWriter_ struct {
}

var Companion_SeqWriter_ = CompanionStruct_SeqWriter_{}

func (_this *SeqWriter) Equals(other *SeqWriter) bool {
	return _this == other
}

func (_this *SeqWriter) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*SeqWriter)
	return ok && _this.Equals(other)
}

func (*SeqWriter) String() string {
	return "Std_Consumers.SeqWriter"
}

func Type_SeqWriter_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_SeqWriter_{Type_T_}
}

type type_SeqWriter_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_SeqWriter_) Default() interface{} {
	return (*SeqWriter)(nil)
}

func (_this type_SeqWriter_) String() string {
	return "Std_Consumers.SeqWriter"
}
func (_this *SeqWriter) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_IConsumer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ IConsumer = &SeqWriter{}
var _ m_Std_Actions.Action = &SeqWriter{}
var _ m_Std_GenericActions.GenericAction = &SeqWriter{}
var _ m_Std_Frames.Validatable = &SeqWriter{}
var _ _dafny.TraitOffspring = &SeqWriter{}

func (_this *SeqWriter) Accept(t interface{}) {
	Companion_IConsumer_.Accept(_this, t)
	return
}
func (_this *SeqWriter) Ctor__() {
	{
		(_this).Values = _dafny.SeqOf()
	}
}
func (_this *SeqWriter) Invoke(t interface{}) interface{} {
	{
		var t interface{} = t
		_ = t
		var r _dafny.Tuple = _dafny.TupleOf()
		_ = r
		(_this).Values = _dafny.Companion_Sequence_.Concatenate(_this.Values, _dafny.SeqOf(t))
		r = _dafny.TupleOf()
		return r
	}
}

// End of class SeqWriter

// Definition of class SeqWriterTotalActionProof
type SeqWriterTotalActionProof struct {
	dummy byte
}

func New_SeqWriterTotalActionProof_() *SeqWriterTotalActionProof {
	_this := SeqWriterTotalActionProof{}

	return &_this
}

type CompanionStruct_SeqWriterTotalActionProof_ struct {
}

var Companion_SeqWriterTotalActionProof_ = CompanionStruct_SeqWriterTotalActionProof_{}

func (_this *SeqWriterTotalActionProof) Equals(other *SeqWriterTotalActionProof) bool {
	return _this == other
}

func (_this *SeqWriterTotalActionProof) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*SeqWriterTotalActionProof)
	return ok && _this.Equals(other)
}

func (*SeqWriterTotalActionProof) String() string {
	return "Std_Consumers.SeqWriterTotalActionProof"
}

func Type_SeqWriterTotalActionProof_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_SeqWriterTotalActionProof_{Type_T_}
}

type type_SeqWriterTotalActionProof_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_SeqWriterTotalActionProof_) Default() interface{} {
	return (*SeqWriterTotalActionProof)(nil)
}

func (_this type_SeqWriterTotalActionProof_) String() string {
	return "Std_Consumers.SeqWriterTotalActionProof"
}
func (_this *SeqWriterTotalActionProof) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ m_Std_Actions.TotalActionProof = &SeqWriterTotalActionProof{}
var _ m_Std_Frames.Validatable = &SeqWriterTotalActionProof{}
var _ _dafny.TraitOffspring = &SeqWriterTotalActionProof{}

// End of class SeqWriterTotalActionProof
