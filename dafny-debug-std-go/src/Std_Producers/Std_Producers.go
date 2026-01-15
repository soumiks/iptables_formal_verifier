// Package Std_Producers
// Dafny module Std_Producers compiled into Go

package Std_Producers

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
	m_Std_Consumers "Std_Consumers"
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
var _ m_Std_Consumers.Dummy__

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
	return "Std_Producers.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) DefaultForEach(producer Producer, consumer m_Std_Consumers.IConsumer) {
	{
		for true {
			{
				var _0_t m_Std_Wrappers.Option
				_ = _0_t
				var _out0 m_Std_Wrappers.Option
				_ = _out0
				_out0 = (producer).Next()
				_0_t = _out0
				if (_0_t).Equals(m_Std_Wrappers.Companion_Option_.Create_None_()) {
					goto L2
				}
				(consumer).Accept((_0_t).Dtor_value())
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
}
func (_static *CompanionStruct_Default___) DefaultFill(producer Producer, consumer m_Std_Consumers.Consumer) {
	{
		var _lo0 _dafny.Int = _dafny.Zero
		_ = _lo0
		for _0_c := ((consumer).Capacity()).Dtor_value().(_dafny.Int); _lo0.Cmp(_0_c) < 0; {
			_0_c = _0_c.Minus(_dafny.One)
			{
				var _1_t m_Std_Wrappers.Option
				_ = _1_t
				var _out0 m_Std_Wrappers.Option
				_ = _out0
				_out0 = (producer).Next()
				_1_t = _out0
				if (_1_t).Equals(m_Std_Wrappers.Companion_Option_.Create_None_()) {
					goto L3
				}
				var _2_accepted bool
				_ = _2_accepted
				var _out1 bool
				_ = _out1
				_out1 = (consumer).Accept((_1_t).Dtor_value())
				_2_accepted = _out1
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
}
func (_static *CompanionStruct_Default___) IsNone(o m_Std_Wrappers.Option) bool {
	return (o).Is_None()
}
func (_static *CompanionStruct_Default___) IsSome(o m_Std_Wrappers.Option) bool {
	return (o).Is_Some()
}
func (_static *CompanionStruct_Default___) ProducedOf(outputs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if ((_dafny.IntOfUint32((outputs).Cardinality())).Sign() == 0) || (((outputs).Select(0).(m_Std_Wrappers.Option)).Is_None()) {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(((outputs).Select(0).(m_Std_Wrappers.Option)).Dtor_value()))
		var _in0 _dafny.Sequence = (outputs).Drop(1)
		_ = _in0
		outputs = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) CollectToSeq(p Producer) _dafny.Sequence {
	var s _dafny.Sequence = _dafny.EmptySeq
	_ = s
	var _0_seqWriter *m_Std_Consumers.SeqWriter
	_ = _0_seqWriter
	var _nw0 *m_Std_Consumers.SeqWriter = m_Std_Consumers.New_SeqWriter_()
	_ = _nw0
	_nw0.Ctor__()
	_0_seqWriter = _nw0
	(p).ForEach(_0_seqWriter)
	s = _0_seqWriter.Values
	return s
	return s
}

// End of class Default__

// Definition of trait IProducer
type IProducer interface {
	String() string
	Invoke(i interface{}) interface{}
	Next() interface{}
}

func (_static *CompanionStruct_IProducer_) Next(_this IProducer) interface{} {
	{
		var r interface{} = (interface{})(nil)
		_ = r
		var _out0 interface{}
		_ = _out0
		_out0 = (_this).Invoke(_dafny.TupleOf())
		r = _out0
		return r
	}
}

type CompanionStruct_IProducer_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_IProducer_ = CompanionStruct_IProducer_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_IProducer_) CastTo_(x interface{}) IProducer {
	var t IProducer
	t, _ = x.(IProducer)
	return t
}

// End of trait IProducer

// Definition of trait ProducesSetProof
type ProducesSetProof interface {
	String() string
}
type CompanionStruct_ProducesSetProof_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_ProducesSetProof_ = CompanionStruct_ProducesSetProof_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_ProducesSetProof_) CastTo_(x interface{}) ProducesSetProof {
	var t ProducesSetProof
	t, _ = x.(ProducesSetProof)
	return t
}

// End of trait ProducesSetProof

// Definition of class FunctionalIProducer
type FunctionalIProducer struct {
	State   interface{}
	_stepFn func(interface{}) _dafny.Tuple
}

func New_FunctionalIProducer_() *FunctionalIProducer {
	_this := FunctionalIProducer{}

	_this.State = (interface{})(nil)
	_this._stepFn = func(interface{}) _dafny.Tuple { return _dafny.Tuple{} }
	return &_this
}

type CompanionStruct_FunctionalIProducer_ struct {
}

var Companion_FunctionalIProducer_ = CompanionStruct_FunctionalIProducer_{}

func (_this *FunctionalIProducer) Equals(other *FunctionalIProducer) bool {
	return _this == other
}

func (_this *FunctionalIProducer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*FunctionalIProducer)
	return ok && _this.Equals(other)
}

func (*FunctionalIProducer) String() string {
	return "Std_Producers.FunctionalIProducer"
}

func Type_FunctionalIProducer_(Type_S_ _dafny.TypeDescriptor, Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_FunctionalIProducer_{Type_S_, Type_T_}
}

type type_FunctionalIProducer_ struct {
	Type_S_ _dafny.TypeDescriptor
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_FunctionalIProducer_) Default() interface{} {
	return (*FunctionalIProducer)(nil)
}

func (_this type_FunctionalIProducer_) String() string {
	return "Std_Producers.FunctionalIProducer"
}
func (_this *FunctionalIProducer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_IProducer_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ IProducer = &FunctionalIProducer{}
var _ m_Std_Actions.TotalActionProof = &FunctionalIProducer{}
var _ m_Std_Actions.Action = &FunctionalIProducer{}
var _ m_Std_GenericActions.GenericAction = &FunctionalIProducer{}
var _ m_Std_Frames.Validatable = &FunctionalIProducer{}
var _ _dafny.TraitOffspring = &FunctionalIProducer{}

func (_this *FunctionalIProducer) Next() interface{} {
	var _out1 interface{}
	_ = _out1
	_out1 = Companion_IProducer_.Next(_this)
	return _out1
}
func (_this *FunctionalIProducer) Ctor__(state interface{}, stepFn func(interface{}) _dafny.Tuple) {
	{
		(_this).State = state
		(_this)._stepFn = stepFn
	}
}
func (_this *FunctionalIProducer) Invoke(i interface{}) interface{} {
	{
		var i _dafny.Tuple = i.(_dafny.Tuple)
		_ = i
		var o interface{} = (interface{})(nil)
		_ = o
		var _let_tmp_rhs0 _dafny.Tuple = ((_this).StepFn())(_this.State)
		_ = _let_tmp_rhs0
		var _0_newState interface{} = (*(_let_tmp_rhs0).IndexInt(0)).(interface{})
		_ = _0_newState
		var _1_result_k interface{} = (*(_let_tmp_rhs0).IndexInt(1)).(interface{})
		_ = _1_result_k
		(_this).State = _0_newState
		o = _1_result_k
		return o
	}
}
func (_this *FunctionalIProducer) StepFn() func(interface{}) _dafny.Tuple {
	{
		return _this._stepFn
	}
}

// End of class FunctionalIProducer

// Definition of datatype ProducerState
type ProducerState struct {
	Data_ProducerState_
}

func (_this ProducerState) Get_() Data_ProducerState_ {
	return _this.Data_ProducerState_
}

type Data_ProducerState_ interface {
	isProducerState()
}

type CompanionStruct_ProducerState_ struct {
}

var Companion_ProducerState_ = CompanionStruct_ProducerState_{}

type ProducerState_ProducerState struct {
	Producer  Producer
	Remaining m_Std_Wrappers.Option
	Outputs   _dafny.Sequence
}

func (ProducerState_ProducerState) isProducerState() {}

func (CompanionStruct_ProducerState_) Create_ProducerState_(Producer Producer, Remaining m_Std_Wrappers.Option, Outputs _dafny.Sequence) ProducerState {
	return ProducerState{ProducerState_ProducerState{Producer, Remaining, Outputs}}
}

func (_this ProducerState) Is_ProducerState() bool {
	_, ok := _this.Get_().(ProducerState_ProducerState)
	return ok
}

func (CompanionStruct_ProducerState_) Default() ProducerState {
	return Companion_ProducerState_.Create_ProducerState_((Producer)(nil), m_Std_Wrappers.Companion_Option_.Default(), _dafny.EmptySeq)
}

func (_this ProducerState) Dtor_producer() Producer {
	return _this.Get_().(ProducerState_ProducerState).Producer
}

func (_this ProducerState) Dtor_remaining() m_Std_Wrappers.Option {
	return _this.Get_().(ProducerState_ProducerState).Remaining
}

func (_this ProducerState) Dtor_outputs() _dafny.Sequence {
	return _this.Get_().(ProducerState_ProducerState).Outputs
}

func (_this ProducerState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ProducerState_ProducerState:
		{
			return "Producers.ProducerState.ProducerState" + "(" + _dafny.String(data.Producer) + ", " + _dafny.String(data.Remaining) + ", " + _dafny.String(data.Outputs) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ProducerState) Equals(other ProducerState) bool {
	switch data1 := _this.Get_().(type) {
	case ProducerState_ProducerState:
		{
			data2, ok := other.Get_().(ProducerState_ProducerState)
			return ok && _dafny.AreEqual(data1.Producer, data2.Producer) && data1.Remaining.Equals(data2.Remaining) && data1.Outputs.Equals(data2.Outputs)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ProducerState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ProducerState)
	return ok && _this.Equals(typed)
}

func Type_ProducerState_() _dafny.TypeDescriptor {
	return type_ProducerState_{}
}

type type_ProducerState_ struct {
}

func (_this type_ProducerState_) Default() interface{} {
	return Companion_ProducerState_.Default()
}

func (_this type_ProducerState_) String() string {
	return "Std_Producers.ProducerState"
}
func (_this ProducerState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ProducerState{}

// End of datatype ProducerState

// Definition of trait Producer
type Producer interface {
	String() string
	ProducedCount() _dafny.Int
	Remaining() m_Std_Wrappers.Option
	Invoke(i interface{}) interface{}
	Next() m_Std_Wrappers.Option
	ForEach(consumer m_Std_Consumers.IConsumer)
	Fill(consumer m_Std_Consumers.Consumer)
}

func (_static *CompanionStruct_Producer_) Next(_this Producer) m_Std_Wrappers.Option {
	{
		var r m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = r
		var _out0 interface{}
		_ = _out0
		_out0 = (_this).Invoke(_dafny.TupleOf())
		r = _out0.(m_Std_Wrappers.Option)
		return r
	}
}

type CompanionStruct_Producer_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_Producer_ = CompanionStruct_Producer_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_Producer_) CastTo_(x interface{}) Producer {
	var t Producer
	t, _ = x.(Producer)
	return t
}

// End of trait Producer

// Definition of class EmptyProducer
type EmptyProducer struct {
	dummy byte
}

func New_EmptyProducer_() *EmptyProducer {
	_this := EmptyProducer{}

	return &_this
}

type CompanionStruct_EmptyProducer_ struct {
}

var Companion_EmptyProducer_ = CompanionStruct_EmptyProducer_{}

func (_this *EmptyProducer) Equals(other *EmptyProducer) bool {
	return _this == other
}

func (_this *EmptyProducer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*EmptyProducer)
	return ok && _this.Equals(other)
}

func (*EmptyProducer) String() string {
	return "Std_Producers.EmptyProducer"
}

func Type_EmptyProducer_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_EmptyProducer_{Type_T_}
}

type type_EmptyProducer_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_EmptyProducer_) Default() interface{} {
	return (*EmptyProducer)(nil)
}

func (_this type_EmptyProducer_) String() string {
	return "Std_Producers.EmptyProducer"
}
func (_this *EmptyProducer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Producer = &EmptyProducer{}
var _ m_Std_Actions.Action = &EmptyProducer{}
var _ m_Std_Actions.TotalActionProof = &EmptyProducer{}
var _ m_Std_GenericActions.GenericAction = &EmptyProducer{}
var _ m_Std_Frames.Validatable = &EmptyProducer{}
var _ _dafny.TraitOffspring = &EmptyProducer{}

func (_this *EmptyProducer) Next() m_Std_Wrappers.Option {
	var _out1 m_Std_Wrappers.Option
	_ = _out1
	_out1 = Companion_Producer_.Next(_this)
	return _out1
}
func (_this *EmptyProducer) Ctor__() {
	{
	}
}
func (_this *EmptyProducer) ProducedCount() _dafny.Int {
	{
		return _dafny.Zero
	}
}
func (_this *EmptyProducer) Remaining() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_Some_(_dafny.Zero)
	}
}
func (_this *EmptyProducer) Invoke(t interface{}) interface{} {
	{
		var t _dafny.Tuple = t.(_dafny.Tuple)
		_ = t
		var value m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = value
		value = m_Std_Wrappers.Companion_Option_.Create_None_()
		return value
	}
}
func (_this *EmptyProducer) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *EmptyProducer) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		Companion_Default___.DefaultFill(_this, consumer)
	}
}

// End of class EmptyProducer

// Definition of class RepeatProducer
type RepeatProducer struct {
	Go__producedCount _dafny.Int
	_n                _dafny.Int
	_t                interface{}
}

func New_RepeatProducer_() *RepeatProducer {
	_this := RepeatProducer{}

	_this.Go__producedCount = _dafny.Zero
	_this._n = _dafny.Zero
	_this._t = (interface{})(nil)
	return &_this
}

type CompanionStruct_RepeatProducer_ struct {
}

var Companion_RepeatProducer_ = CompanionStruct_RepeatProducer_{}

func (_this *RepeatProducer) Equals(other *RepeatProducer) bool {
	return _this == other
}

func (_this *RepeatProducer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*RepeatProducer)
	return ok && _this.Equals(other)
}

func (*RepeatProducer) String() string {
	return "Std_Producers.RepeatProducer"
}

func Type_RepeatProducer_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_RepeatProducer_{Type_T_}
}

type type_RepeatProducer_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_RepeatProducer_) Default() interface{} {
	return (*RepeatProducer)(nil)
}

func (_this type_RepeatProducer_) String() string {
	return "Std_Producers.RepeatProducer"
}
func (_this *RepeatProducer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Producer = &RepeatProducer{}
var _ m_Std_Actions.Action = &RepeatProducer{}
var _ m_Std_Actions.TotalActionProof = &RepeatProducer{}
var _ m_Std_GenericActions.GenericAction = &RepeatProducer{}
var _ m_Std_Frames.Validatable = &RepeatProducer{}
var _ _dafny.TraitOffspring = &RepeatProducer{}

func (_this *RepeatProducer) Next() m_Std_Wrappers.Option {
	var _out2 m_Std_Wrappers.Option
	_ = _out2
	_out2 = Companion_Producer_.Next(_this)
	return _out2
}
func (_this *RepeatProducer) Ctor__(n _dafny.Int, t interface{}) {
	{
		(_this)._n = n
		(_this)._t = t
		(_this).Go__producedCount = _dafny.Zero
	}
}
func (_this *RepeatProducer) ProducedCount() _dafny.Int {
	{
		return _this.Go__producedCount
	}
}
func (_this *RepeatProducer) Remaining() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_Some_(((_this).N()).Minus(_this.Go__producedCount))
	}
}
func (_this *RepeatProducer) Invoke(i interface{}) interface{} {
	{
		var i _dafny.Tuple = i.(_dafny.Tuple)
		_ = i
		var value m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = value
		if (_this.Go__producedCount).Cmp((_this).N()) == 0 {
			value = m_Std_Wrappers.Companion_Option_.Create_None_()
		} else {
			value = m_Std_Wrappers.Companion_Option_.Create_Some_((_this).T())
			(_this).Go__producedCount = (_this.Go__producedCount).Plus(_dafny.One)
		}
		return value
	}
}
func (_this *RepeatProducer) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *RepeatProducer) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		Companion_Default___.DefaultFill(_this, consumer)
	}
}
func (_this *RepeatProducer) N() _dafny.Int {
	{
		return _this._n
	}
}
func (_this *RepeatProducer) T() interface{} {
	{
		return _this._t
	}
}

// End of class RepeatProducer

// Definition of class SeqReader
type SeqReader struct {
	Index     _dafny.Int
	_elements _dafny.Sequence
}

func New_SeqReader_() *SeqReader {
	_this := SeqReader{}

	_this.Index = _dafny.Zero
	_this._elements = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_SeqReader_ struct {
}

var Companion_SeqReader_ = CompanionStruct_SeqReader_{}

func (_this *SeqReader) Equals(other *SeqReader) bool {
	return _this == other
}

func (_this *SeqReader) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*SeqReader)
	return ok && _this.Equals(other)
}

func (*SeqReader) String() string {
	return "Std_Producers.SeqReader"
}

func Type_SeqReader_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_SeqReader_{Type_T_}
}

type type_SeqReader_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_SeqReader_) Default() interface{} {
	return (*SeqReader)(nil)
}

func (_this type_SeqReader_) String() string {
	return "Std_Producers.SeqReader"
}
func (_this *SeqReader) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Producer = &SeqReader{}
var _ m_Std_Actions.Action = &SeqReader{}
var _ m_Std_Actions.TotalActionProof = &SeqReader{}
var _ m_Std_GenericActions.GenericAction = &SeqReader{}
var _ m_Std_Frames.Validatable = &SeqReader{}
var _ _dafny.TraitOffspring = &SeqReader{}

func (_this *SeqReader) Next() m_Std_Wrappers.Option {
	var _out3 m_Std_Wrappers.Option
	_ = _out3
	_out3 = Companion_Producer_.Next(_this)
	return _out3
}
func (_this *SeqReader) Ctor__(elements _dafny.Sequence) {
	{
		(_this)._elements = elements
		(_this).Index = _dafny.Zero
	}
}
func (_this *SeqReader) ProducedCount() _dafny.Int {
	{
		return _this.Index
	}
}
func (_this *SeqReader) Remaining() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_Some_((_dafny.IntOfUint32(((_this).Elements()).Cardinality())).Minus(_this.Index))
	}
}
func (_this *SeqReader) Invoke(t interface{}) interface{} {
	{
		var t _dafny.Tuple = t.(_dafny.Tuple)
		_ = t
		var value m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = value
		if (_dafny.IntOfUint32(((_this).Elements()).Cardinality())).Cmp(_this.Index) == 0 {
			value = m_Std_Wrappers.Companion_Option_.Create_None_()
		} else {
			value = m_Std_Wrappers.Companion_Option_.Create_Some_(((_this).Elements()).Select((_this.Index).Uint32()).(interface{}))
			(_this).Index = (_this.Index).Plus(_dafny.One)
		}
		return value
	}
}
func (_this *SeqReader) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *SeqReader) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		Companion_Default___.DefaultFill(_this, consumer)
	}
}
func (_this *SeqReader) Elements() _dafny.Sequence {
	{
		return _this._elements
	}
}

// End of class SeqReader

// Definition of trait ProducerOfSetProof
type ProducerOfSetProof interface {
	String() string
}
type CompanionStruct_ProducerOfSetProof_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_ProducerOfSetProof_ = CompanionStruct_ProducerOfSetProof_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_ProducerOfSetProof_) CastTo_(x interface{}) ProducerOfSetProof {
	var t ProducerOfSetProof
	t, _ = x.(ProducerOfSetProof)
	return t
}

// End of trait ProducerOfSetProof

// Definition of class LimitedProducer
type LimitedProducer struct {
	Produced  _dafny.Int
	_original IProducer
	_max      _dafny.Int
}

func New_LimitedProducer_() *LimitedProducer {
	_this := LimitedProducer{}

	_this.Produced = _dafny.Zero
	_this._original = (IProducer)(nil)
	_this._max = _dafny.Zero
	return &_this
}

type CompanionStruct_LimitedProducer_ struct {
}

var Companion_LimitedProducer_ = CompanionStruct_LimitedProducer_{}

func (_this *LimitedProducer) Equals(other *LimitedProducer) bool {
	return _this == other
}

func (_this *LimitedProducer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*LimitedProducer)
	return ok && _this.Equals(other)
}

func (*LimitedProducer) String() string {
	return "Std_Producers.LimitedProducer"
}

func Type_LimitedProducer_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_LimitedProducer_{Type_T_}
}

type type_LimitedProducer_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_LimitedProducer_) Default() interface{} {
	return (*LimitedProducer)(nil)
}

func (_this type_LimitedProducer_) String() string {
	return "Std_Producers.LimitedProducer"
}
func (_this *LimitedProducer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Producer = &LimitedProducer{}
var _ m_Std_Actions.Action = &LimitedProducer{}
var _ m_Std_Actions.TotalActionProof = &LimitedProducer{}
var _ m_Std_GenericActions.GenericAction = &LimitedProducer{}
var _ m_Std_Frames.Validatable = &LimitedProducer{}
var _ _dafny.TraitOffspring = &LimitedProducer{}

func (_this *LimitedProducer) Next() m_Std_Wrappers.Option {
	var _out4 m_Std_Wrappers.Option
	_ = _out4
	_out4 = Companion_Producer_.Next(_this)
	return _out4
}
func (_this *LimitedProducer) Ctor__(original IProducer, max _dafny.Int) {
	{
		(_this)._original = original
		(_this)._max = max
		(_this).Produced = _dafny.Zero
	}
}
func (_this *LimitedProducer) ProducedCount() _dafny.Int {
	{
		return _this.Produced
	}
}
func (_this *LimitedProducer) Remaining() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_Some_(((_this).Max()).Minus(_this.Produced))
	}
}
func (_this *LimitedProducer) Invoke(t interface{}) interface{} {
	{
		var t _dafny.Tuple = t.(_dafny.Tuple)
		_ = t
		var value m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = value
		if (_this.Produced).Cmp((_this).Max()) == 0 {
			value = m_Std_Wrappers.Companion_Option_.Create_None_()
		} else {
			var _0_v interface{}
			_ = _0_v
			var _out0 interface{}
			_ = _out0
			_out0 = ((_this).Original()).Invoke(_dafny.TupleOf())
			_0_v = _out0
			value = m_Std_Wrappers.Companion_Option_.Create_Some_(_0_v)
			(_this).Produced = (_this.Produced).Plus(_dafny.One)
		}
		return value
	}
}
func (_this *LimitedProducer) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *LimitedProducer) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		Companion_Default___.DefaultFill(_this, consumer)
	}
}
func (_this *LimitedProducer) Original() IProducer {
	{
		return _this._original
	}
}
func (_this *LimitedProducer) Max() _dafny.Int {
	{
		return _this._max
	}
}

// End of class LimitedProducer

// Definition of class FilteredProducer
type FilteredProducer struct {
	Go__producedCount _dafny.Int
	_source           Producer
	_filter           func(interface{}) bool
}

func New_FilteredProducer_() *FilteredProducer {
	_this := FilteredProducer{}

	_this.Go__producedCount = _dafny.Zero
	_this._source = (Producer)(nil)
	_this._filter = func(interface{}) bool { return false }
	return &_this
}

type CompanionStruct_FilteredProducer_ struct {
}

var Companion_FilteredProducer_ = CompanionStruct_FilteredProducer_{}

func (_this *FilteredProducer) Equals(other *FilteredProducer) bool {
	return _this == other
}

func (_this *FilteredProducer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*FilteredProducer)
	return ok && _this.Equals(other)
}

func (*FilteredProducer) String() string {
	return "Std_Producers.FilteredProducer"
}

func Type_FilteredProducer_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_FilteredProducer_{Type_T_}
}

type type_FilteredProducer_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_FilteredProducer_) Default() interface{} {
	return (*FilteredProducer)(nil)
}

func (_this type_FilteredProducer_) String() string {
	return "Std_Producers.FilteredProducer"
}
func (_this *FilteredProducer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Producer = &FilteredProducer{}
var _ m_Std_Actions.Action = &FilteredProducer{}
var _ m_Std_Actions.TotalActionProof = &FilteredProducer{}
var _ m_Std_GenericActions.GenericAction = &FilteredProducer{}
var _ m_Std_Frames.Validatable = &FilteredProducer{}
var _ _dafny.TraitOffspring = &FilteredProducer{}

func (_this *FilteredProducer) Next() m_Std_Wrappers.Option {
	var _out5 m_Std_Wrappers.Option
	_ = _out5
	_out5 = Companion_Producer_.Next(_this)
	return _out5
}
func (_this *FilteredProducer) Ctor__(source Producer, filter func(interface{}) bool) {
	{
		(_this)._source = source
		(_this)._filter = filter
		(_this).Go__producedCount = _dafny.Zero
	}
}
func (_this *FilteredProducer) ProducedCount() _dafny.Int {
	{
		return _this.Go__producedCount
	}
}
func (_this *FilteredProducer) Remaining() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_None_()
	}
}
func (_this *FilteredProducer) Invoke(t interface{}) interface{} {
	{
		var t _dafny.Tuple = t.(_dafny.Tuple)
		_ = t
		var result m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = result
		result = m_Std_Wrappers.Companion_Option_.Create_None_()
		var _0_notFirstLoop bool
		_ = _0_notFirstLoop
		_0_notFirstLoop = false
		{
			for true {
				{
					_0_notFirstLoop = true
					var _out0 m_Std_Wrappers.Option
					_ = _out0
					_out0 = ((_this).Source()).Next()
					result = _out0
					if ((result).Is_None()) || (((_this).Filter())((result).Dtor_value())) {
						goto L4
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		if (result).Is_Some() {
			(_this).Go__producedCount = (_this.Go__producedCount).Plus(_dafny.One)
		} else {
		}
		return result
	}
}
func (_this *FilteredProducer) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *FilteredProducer) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		Companion_Default___.DefaultFill(_this, consumer)
	}
}
func (_this *FilteredProducer) Source() Producer {
	{
		return _this._source
	}
}
func (_this *FilteredProducer) Filter() func(interface{}) bool {
	{
		return _this._filter
	}
}

// End of class FilteredProducer

// Definition of class ConcatenatedProducer
type ConcatenatedProducer struct {
	_first  Producer
	_second Producer
}

func New_ConcatenatedProducer_() *ConcatenatedProducer {
	_this := ConcatenatedProducer{}

	_this._first = (Producer)(nil)
	_this._second = (Producer)(nil)
	return &_this
}

type CompanionStruct_ConcatenatedProducer_ struct {
}

var Companion_ConcatenatedProducer_ = CompanionStruct_ConcatenatedProducer_{}

func (_this *ConcatenatedProducer) Equals(other *ConcatenatedProducer) bool {
	return _this == other
}

func (_this *ConcatenatedProducer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*ConcatenatedProducer)
	return ok && _this.Equals(other)
}

func (*ConcatenatedProducer) String() string {
	return "Std_Producers.ConcatenatedProducer"
}

func Type_ConcatenatedProducer_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_ConcatenatedProducer_{Type_T_}
}

type type_ConcatenatedProducer_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_ConcatenatedProducer_) Default() interface{} {
	return (*ConcatenatedProducer)(nil)
}

func (_this type_ConcatenatedProducer_) String() string {
	return "Std_Producers.ConcatenatedProducer"
}
func (_this *ConcatenatedProducer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Producer = &ConcatenatedProducer{}
var _ m_Std_Actions.Action = &ConcatenatedProducer{}
var _ m_Std_Actions.TotalActionProof = &ConcatenatedProducer{}
var _ m_Std_GenericActions.GenericAction = &ConcatenatedProducer{}
var _ m_Std_Frames.Validatable = &ConcatenatedProducer{}
var _ _dafny.TraitOffspring = &ConcatenatedProducer{}

func (_this *ConcatenatedProducer) Next() m_Std_Wrappers.Option {
	var _out6 m_Std_Wrappers.Option
	_ = _out6
	_out6 = Companion_Producer_.Next(_this)
	return _out6
}
func (_this *ConcatenatedProducer) Ctor__(first Producer, second Producer) {
	{
		(_this)._first = first
		(_this)._second = second
	}
}
func (_this *ConcatenatedProducer) ProducedCount() _dafny.Int {
	{
		return (((_this).First()).ProducedCount()).Plus(((_this).Second()).ProducedCount())
	}
}
func (_this *ConcatenatedProducer) Remaining() m_Std_Wrappers.Option {
	{
		var _0_left m_Std_Wrappers.Option = ((_this).First()).Remaining()
		_ = _0_left
		var _1_right m_Std_Wrappers.Option = ((_this).Second()).Remaining()
		_ = _1_right
		if ((_0_left).Is_Some()) && ((_1_right).Is_Some()) {
			return m_Std_Wrappers.Companion_Option_.Create_Some_(((_0_left).Dtor_value().(_dafny.Int)).Plus((_1_right).Dtor_value().(_dafny.Int)))
		} else {
			return m_Std_Wrappers.Companion_Option_.Create_None_()
		}
	}
}
func (_this *ConcatenatedProducer) Invoke(t interface{}) interface{} {
	{
		var t _dafny.Tuple = t.(_dafny.Tuple)
		_ = t
		var result m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = result
		var _out0 m_Std_Wrappers.Option
		_ = _out0
		_out0 = ((_this).First()).Next()
		result = _out0
		if (result).Is_Some() {
		} else {
			var _out1 m_Std_Wrappers.Option
			_ = _out1
			_out1 = ((_this).Second()).Next()
			result = _out1
		}
		return result
	}
}
func (_this *ConcatenatedProducer) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *ConcatenatedProducer) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		Companion_Default___.DefaultFill(_this, consumer)
	}
}
func (_this *ConcatenatedProducer) First() Producer {
	{
		return _this._first
	}
}
func (_this *ConcatenatedProducer) Second() Producer {
	{
		return _this._second
	}
}

// End of class ConcatenatedProducer

// Definition of class MappedProducer
type MappedProducer struct {
	_original Producer
	_mapping  m_Std_Actions.Action
}

func New_MappedProducer_() *MappedProducer {
	_this := MappedProducer{}

	_this._original = (Producer)(nil)
	_this._mapping = (m_Std_Actions.Action)(nil)
	return &_this
}

type CompanionStruct_MappedProducer_ struct {
}

var Companion_MappedProducer_ = CompanionStruct_MappedProducer_{}

func (_this *MappedProducer) Equals(other *MappedProducer) bool {
	return _this == other
}

func (_this *MappedProducer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*MappedProducer)
	return ok && _this.Equals(other)
}

func (*MappedProducer) String() string {
	return "Std_Producers.MappedProducer"
}

func Type_MappedProducer_(Type_I_ _dafny.TypeDescriptor, Type_O_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_MappedProducer_{Type_I_, Type_O_}
}

type type_MappedProducer_ struct {
	Type_I_ _dafny.TypeDescriptor
	Type_O_ _dafny.TypeDescriptor
}

func (_this type_MappedProducer_) Default() interface{} {
	return (*MappedProducer)(nil)
}

func (_this type_MappedProducer_) String() string {
	return "Std_Producers.MappedProducer"
}
func (_this *MappedProducer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Producer = &MappedProducer{}
var _ m_Std_Actions.Action = &MappedProducer{}
var _ m_Std_Actions.TotalActionProof = &MappedProducer{}
var _ m_Std_GenericActions.GenericAction = &MappedProducer{}
var _ m_Std_Frames.Validatable = &MappedProducer{}
var _ _dafny.TraitOffspring = &MappedProducer{}

func (_this *MappedProducer) Next() m_Std_Wrappers.Option {
	var _out7 m_Std_Wrappers.Option
	_ = _out7
	_out7 = Companion_Producer_.Next(_this)
	return _out7
}
func (_this *MappedProducer) Ctor__(original Producer, mapping m_Std_Actions.Action) {
	{
		(_this)._original = original
		(_this)._mapping = mapping
	}
}
func (_this *MappedProducer) ProducedCount() _dafny.Int {
	{
		return ((_this).Original()).ProducedCount()
	}
}
func (_this *MappedProducer) Remaining() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_None_()
	}
}
func (_this *MappedProducer) Invoke(t interface{}) interface{} {
	{
		var t _dafny.Tuple = t.(_dafny.Tuple)
		_ = t
		var result m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = result
		var _0_next m_Std_Wrappers.Option
		_ = _0_next
		var _out0 m_Std_Wrappers.Option
		_ = _out0
		_out0 = ((_this).Original()).Next()
		_0_next = _out0
		if (_0_next).Is_Some() {
			var _1_nextValue interface{}
			_ = _1_nextValue
			var _out1 interface{}
			_ = _out1
			_out1 = ((_this).Mapping()).Invoke((_0_next).Dtor_value())
			_1_nextValue = _out1
			result = m_Std_Wrappers.Companion_Option_.Create_Some_(_1_nextValue)
		} else {
			result = m_Std_Wrappers.Companion_Option_.Create_None_()
		}
		return result
	}
}
func (_this *MappedProducer) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *MappedProducer) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		Companion_Default___.DefaultFill(_this, consumer)
	}
}
func (_this *MappedProducer) Original() Producer {
	{
		return _this._original
	}
}
func (_this *MappedProducer) Mapping() m_Std_Actions.Action {
	{
		return _this._mapping
	}
}

// End of class MappedProducer

// Definition of trait ProducerOfNewProducers
type ProducerOfNewProducers interface {
	String() string
	Fill(consumer m_Std_Consumers.Consumer)
	ForEach(consumer m_Std_Consumers.IConsumer)
	Next() m_Std_Wrappers.Option
	ProducedCount() _dafny.Int
	Remaining() m_Std_Wrappers.Option
	Invoke(i interface{}) interface{}
}
type CompanionStruct_ProducerOfNewProducers_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_ProducerOfNewProducers_ = CompanionStruct_ProducerOfNewProducers_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_ProducerOfNewProducers_) CastTo_(x interface{}) ProducerOfNewProducers {
	var t ProducerOfNewProducers
	t, _ = x.(ProducerOfNewProducers)
	return t
}

// End of trait ProducerOfNewProducers

// Definition of trait OutputterOfNewProducers
type OutputterOfNewProducers interface {
	String() string
	Invoke(i interface{}) interface{}
}
type CompanionStruct_OutputterOfNewProducers_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_OutputterOfNewProducers_ = CompanionStruct_OutputterOfNewProducers_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_OutputterOfNewProducers_) CastTo_(x interface{}) OutputterOfNewProducers {
	var t OutputterOfNewProducers
	t, _ = x.(OutputterOfNewProducers)
	return t
}

// End of trait OutputterOfNewProducers

// Definition of class MappedProducerOfNewProducers
type MappedProducerOfNewProducers struct {
	_original Producer
	_mapping  OutputterOfNewProducers
}

func New_MappedProducerOfNewProducers_() *MappedProducerOfNewProducers {
	_this := MappedProducerOfNewProducers{}

	_this._original = (Producer)(nil)
	_this._mapping = (OutputterOfNewProducers)(nil)
	return &_this
}

type CompanionStruct_MappedProducerOfNewProducers_ struct {
}

var Companion_MappedProducerOfNewProducers_ = CompanionStruct_MappedProducerOfNewProducers_{}

func (_this *MappedProducerOfNewProducers) Equals(other *MappedProducerOfNewProducers) bool {
	return _this == other
}

func (_this *MappedProducerOfNewProducers) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*MappedProducerOfNewProducers)
	return ok && _this.Equals(other)
}

func (*MappedProducerOfNewProducers) String() string {
	return "Std_Producers.MappedProducerOfNewProducers"
}

func Type_MappedProducerOfNewProducers_(Type_I_ _dafny.TypeDescriptor, Type_O_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_MappedProducerOfNewProducers_{Type_I_, Type_O_}
}

type type_MappedProducerOfNewProducers_ struct {
	Type_I_ _dafny.TypeDescriptor
	Type_O_ _dafny.TypeDescriptor
}

func (_this type_MappedProducerOfNewProducers_) Default() interface{} {
	return (*MappedProducerOfNewProducers)(nil)
}

func (_this type_MappedProducerOfNewProducers_) String() string {
	return "Std_Producers.MappedProducerOfNewProducers"
}
func (_this *MappedProducerOfNewProducers) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_ProducerOfNewProducers_.TraitID_, Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ ProducerOfNewProducers = &MappedProducerOfNewProducers{}
var _ Producer = &MappedProducerOfNewProducers{}
var _ m_Std_Actions.Action = &MappedProducerOfNewProducers{}
var _ m_Std_Actions.TotalActionProof = &MappedProducerOfNewProducers{}
var _ m_Std_GenericActions.GenericAction = &MappedProducerOfNewProducers{}
var _ m_Std_Frames.Validatable = &MappedProducerOfNewProducers{}
var _ _dafny.TraitOffspring = &MappedProducerOfNewProducers{}

func (_this *MappedProducerOfNewProducers) Next() m_Std_Wrappers.Option {
	var _out8 m_Std_Wrappers.Option
	_ = _out8
	_out8 = Companion_Producer_.Next(_this)
	return _out8
}
func (_this *MappedProducerOfNewProducers) Ctor__(original Producer, mapping OutputterOfNewProducers) {
	{
		(_this)._original = original
		(_this)._mapping = mapping
	}
}
func (_this *MappedProducerOfNewProducers) ProducedCount() _dafny.Int {
	{
		return ((_this).Original()).ProducedCount()
	}
}
func (_this *MappedProducerOfNewProducers) Remaining() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_None_()
	}
}
func (_this *MappedProducerOfNewProducers) Invoke(t interface{}) interface{} {
	{
		var t _dafny.Tuple = t.(_dafny.Tuple)
		_ = t
		var result m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = result
		var _0_next m_Std_Wrappers.Option
		_ = _0_next
		var _out0 m_Std_Wrappers.Option
		_ = _out0
		_out0 = ((_this).Original()).Next()
		_0_next = _out0
		if (_0_next).Is_Some() {
			var _1_nextValue Producer
			_ = _1_nextValue
			var _out1 interface{}
			_ = _out1
			_out1 = ((_this).Mapping()).Invoke((_0_next).Dtor_value())
			_1_nextValue = Companion_Producer_.CastTo_(Companion_Producer_.CastTo_(_out1))
			result = m_Std_Wrappers.Companion_Option_.Create_Some_(_1_nextValue)
		} else {
			result = m_Std_Wrappers.Companion_Option_.Create_None_()
		}
		return result
	}
}
func (_this *MappedProducerOfNewProducers) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *MappedProducerOfNewProducers) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		Companion_Default___.DefaultFill(_this, consumer)
	}
}
func (_this *MappedProducerOfNewProducers) Original() Producer {
	{
		return _this._original
	}
}
func (_this *MappedProducerOfNewProducers) Mapping() OutputterOfNewProducers {
	{
		return _this._mapping
	}
}

// End of class MappedProducerOfNewProducers

// Definition of class FlattenedProducer
type FlattenedProducer struct {
	CurrentInner      m_Std_Wrappers.Option
	Go__producedCount _dafny.Int
	_original         ProducerOfNewProducers
}

func New_FlattenedProducer_() *FlattenedProducer {
	_this := FlattenedProducer{}

	_this.CurrentInner = m_Std_Wrappers.Companion_Option_.Default()
	_this.Go__producedCount = _dafny.Zero
	_this._original = (ProducerOfNewProducers)(nil)
	return &_this
}

type CompanionStruct_FlattenedProducer_ struct {
}

var Companion_FlattenedProducer_ = CompanionStruct_FlattenedProducer_{}

func (_this *FlattenedProducer) Equals(other *FlattenedProducer) bool {
	return _this == other
}

func (_this *FlattenedProducer) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*FlattenedProducer)
	return ok && _this.Equals(other)
}

func (*FlattenedProducer) String() string {
	return "Std_Producers.FlattenedProducer"
}

func Type_FlattenedProducer_(Type_T_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_FlattenedProducer_{Type_T_}
}

type type_FlattenedProducer_ struct {
	Type_T_ _dafny.TypeDescriptor
}

func (_this type_FlattenedProducer_) Default() interface{} {
	return (*FlattenedProducer)(nil)
}

func (_this type_FlattenedProducer_) String() string {
	return "Std_Producers.FlattenedProducer"
}
func (_this *FlattenedProducer) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ Producer = &FlattenedProducer{}
var _ m_Std_Actions.Action = &FlattenedProducer{}
var _ m_Std_Actions.TotalActionProof = &FlattenedProducer{}
var _ m_Std_GenericActions.GenericAction = &FlattenedProducer{}
var _ m_Std_Frames.Validatable = &FlattenedProducer{}
var _ _dafny.TraitOffspring = &FlattenedProducer{}

func (_this *FlattenedProducer) Next() m_Std_Wrappers.Option {
	var _out9 m_Std_Wrappers.Option
	_ = _out9
	_out9 = Companion_Producer_.Next(_this)
	return _out9
}
func (_this *FlattenedProducer) Ctor__(original ProducerOfNewProducers) {
	{
		(_this)._original = original
		(_this).CurrentInner = m_Std_Wrappers.Companion_Option_.Create_None_()
		(_this).Go__producedCount = _dafny.Zero
	}
}
func (_this *FlattenedProducer) ProducedCount() _dafny.Int {
	{
		return _this.Go__producedCount
	}
}
func (_this *FlattenedProducer) Remaining() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_None_()
	}
}
func (_this *FlattenedProducer) Invoke(t interface{}) interface{} {
	{
		var t _dafny.Tuple = t.(_dafny.Tuple)
		_ = t
		var result m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = result
		result = m_Std_Wrappers.Companion_Option_.Create_None_()
		{
			for (result).Is_None() {
				{
					if (_this.CurrentInner).Is_None() {
						var _out0 interface{}
						_ = _out0
						_out0 = ((_this).Original()).Invoke(_dafny.TupleOf())
						(_this).CurrentInner = _out0.(m_Std_Wrappers.Option)
						if (_this.CurrentInner).Is_Some() {
						} else {
							goto L5
						}
					} else {
						var _out1 m_Std_Wrappers.Option
						_ = _out1
						_out1 = (Companion_Producer_.CastTo_((_this.CurrentInner).Dtor_value())).Next()
						result = _out1
						if (result).Is_None() {
							var _0_oldCurrentInner m_Std_Wrappers.Option
							_ = _0_oldCurrentInner
							_0_oldCurrentInner = _this.CurrentInner
							(_this).CurrentInner = m_Std_Wrappers.Companion_Option_.Create_None_()
						} else {
						}
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		if (result).Is_Some() {
			(_this).Go__producedCount = (_this.Go__producedCount).Plus(_dafny.One)
		} else {
		}
		return result
	}
}
func (_this *FlattenedProducer) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *FlattenedProducer) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		Companion_Default___.DefaultFill(_this, consumer)
	}
}
func (_this *FlattenedProducer) Original() ProducerOfNewProducers {
	{
		return _this._original
	}
}

// End of class FlattenedProducer
