// Package Std_BulkActions
// Dafny module Std_BulkActions compiled into Go

package Std_BulkActions

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
	m_Std_Actions "Std_Actions"
	m_Std_ActionsExterns "Std_ActionsExterns"
	m_Std_Arithmetic_DivInternals "Std_Arithmetic_DivInternals"
	m_Std_Arithmetic_DivInternalsNonlinear "Std_Arithmetic_DivInternalsNonlinear"
	m_Std_Arithmetic_DivMod "Std_Arithmetic_DivMod"
	m_Std_Arithmetic_GeneralInternals "Std_Arithmetic_GeneralInternals"
	m_Std_Arithmetic_Logarithm "Std_Arithmetic_Logarithm"
	m_Std_Arithmetic_ModInternals "Std_Arithmetic_ModInternals"
	m_Std_Arithmetic_ModInternalsNonlinear "Std_Arithmetic_ModInternalsNonlinear"
	m_Std_Arithmetic_Mul "Std_Arithmetic_Mul"
	m_Std_Arithmetic_MulInternals "Std_Arithmetic_MulInternals"
	m_Std_Arithmetic_MulInternalsNonlinear "Std_Arithmetic_MulInternalsNonlinear"
	m_Std_Arithmetic_Power "Std_Arithmetic_Power"
	m_Std_Arithmetic_Power2 "Std_Arithmetic_Power2"
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
	m_Std_FileIO "Std_FileIO"
	m_Std_FileIOInternalExterns "Std_FileIOInternalExterns"
	m_Std_Frames "Std_Frames"
	m_Std_Functions "Std_Functions"
	m_Std_GenericActions "Std_GenericActions"
	m_Std_Math "Std_Math"
	m_Std_Ordinal "Std_Ordinal"
	m_Std_Producers "Std_Producers"
	m_Std_Relations "Std_Relations"
	m_Std_Strings "Std_Strings"
	m_Std_Strings_CharStrEscaping "Std_Strings_CharStrEscaping"
	m_Std_Strings_DecimalConversion "Std_Strings_DecimalConversion"
	m_Std_Strings_HexConversion "Std_Strings_HexConversion"
	m_Std_Termination "Std_Termination"
	m_Std_Unicode_Base "Std_Unicode_Base"
	m_Std_Unicode_UnicodeStringsWithUnicodeChar "Std_Unicode_UnicodeStringsWithUnicodeChar"
	m_Std_Unicode_Utf16EncodingForm "Std_Unicode_Utf16EncodingForm"
	m_Std_Unicode_Utf8EncodingForm "Std_Unicode_Utf8EncodingForm"
	m_Std_Unicode_Utf8EncodingScheme "Std_Unicode_Utf8EncodingScheme"
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
var _ m_Std_Arithmetic_ModInternals.Dummy__
var _ m_Std_Arithmetic_DivInternals.Dummy__
var _ m_Std_Arithmetic_DivMod.Dummy__
var _ m_Std_Arithmetic_Power.Dummy__
var _ m_Std_Arithmetic_Logarithm.Dummy__
var _ m_Std_Arithmetic_Power2.Dummy__
var _ m_Std_Strings_HexConversion.Dummy__
var _ m_Std_Strings_DecimalConversion.Dummy__
var _ m_Std_Strings_CharStrEscaping.Dummy__
var _ m_Std_Strings.Dummy__
var _ m_Std_Unicode_Utf8EncodingForm.Dummy__
var _ m_Std_Unicode_Utf16EncodingForm.Dummy__
var _ m_Std_Unicode_UnicodeStringsWithUnicodeChar.Dummy__
var _ m_Std_Unicode_Utf8EncodingScheme.Dummy__
var _ m_Std_FileIO.Dummy__

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
	return "Std_BulkActions.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ToBatched(t interface{}) Batched {
	return Companion_Batched_.Create_BatchValue_(t)
}
func (_static *CompanionStruct_Default___) ToBatchedProducer(values _dafny.Sequence) m_Std_Producers.Producer {
	var result m_Std_Producers.Producer = (m_Std_Producers.Producer)(nil)
	_ = result
	var _0_chunkProducer *m_Std_Producers.SeqReader
	_ = _0_chunkProducer
	var _nw0 *m_Std_Producers.SeqReader = m_Std_Producers.New_SeqReader_()
	_ = _nw0
	_nw0.Ctor__(values)
	_0_chunkProducer = _nw0
	var _1_mapping *m_Std_Actions.FunctionAction
	_ = _1_mapping
	var _nw1 *m_Std_Actions.FunctionAction = m_Std_Actions.New_FunctionAction_()
	_ = _nw1
	_nw1.Ctor__(func(coer25 func(interface{}) Batched) func(interface{}) interface{} {
		return func(arg26 interface{}) interface{} {
			return coer25(arg26)
		}
	}(func(coer26 func(interface{}) Batched) func(interface{}) Batched {
		return func(arg27 interface{}) Batched {
			return coer26(arg27)
		}
	}(Companion_Default___.ToBatched)))
	_1_mapping = _nw1
	var _nw2 *m_Std_Producers.MappedProducer = m_Std_Producers.New_MappedProducer_()
	_ = _nw2
	_nw2.Ctor__(_0_chunkProducer, _1_mapping)
	result = _nw2
	return result
}

// End of class Default__

// Definition of datatype Batched
type Batched struct {
	Data_Batched_
}

func (_this Batched) Get_() Data_Batched_ {
	return _this.Data_Batched_
}

type Data_Batched_ interface {
	isBatched()
}

type CompanionStruct_Batched_ struct {
}

var Companion_Batched_ = CompanionStruct_Batched_{}

type Batched_BatchValue struct {
	Value interface{}
}

func (Batched_BatchValue) isBatched() {}

func (CompanionStruct_Batched_) Create_BatchValue_(Value interface{}) Batched {
	return Batched{Batched_BatchValue{Value}}
}

func (_this Batched) Is_BatchValue() bool {
	_, ok := _this.Get_().(Batched_BatchValue)
	return ok
}

type Batched_BatchError struct {
	Error interface{}
}

func (Batched_BatchError) isBatched() {}

func (CompanionStruct_Batched_) Create_BatchError_(Error interface{}) Batched {
	return Batched{Batched_BatchError{Error}}
}

func (_this Batched) Is_BatchError() bool {
	_, ok := _this.Get_().(Batched_BatchError)
	return ok
}

type Batched_EndOfInput struct {
}

func (Batched_EndOfInput) isBatched() {}

func (CompanionStruct_Batched_) Create_EndOfInput_() Batched {
	return Batched{Batched_EndOfInput{}}
}

func (_this Batched) Is_EndOfInput() bool {
	_, ok := _this.Get_().(Batched_EndOfInput)
	return ok
}

func (CompanionStruct_Batched_) Default() Batched {
	return Companion_Batched_.Create_EndOfInput_()
}

func (_this Batched) Dtor_value() interface{} {
	return _this.Get_().(Batched_BatchValue).Value
}

func (_this Batched) Dtor_error() interface{} {
	return _this.Get_().(Batched_BatchError).Error
}

func (_this Batched) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Batched_BatchValue:
		{
			return "BulkActions.Batched.BatchValue" + "(" + _dafny.String(data.Value) + ")"
		}
	case Batched_BatchError:
		{
			return "BulkActions.Batched.BatchError" + "(" + _dafny.String(data.Error) + ")"
		}
	case Batched_EndOfInput:
		{
			return "BulkActions.Batched.EndOfInput"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Batched) Equals(other Batched) bool {
	switch data1 := _this.Get_().(type) {
	case Batched_BatchValue:
		{
			data2, ok := other.Get_().(Batched_BatchValue)
			return ok && _dafny.AreEqual(data1.Value, data2.Value)
		}
	case Batched_BatchError:
		{
			data2, ok := other.Get_().(Batched_BatchError)
			return ok && _dafny.AreEqual(data1.Error, data2.Error)
		}
	case Batched_EndOfInput:
		{
			_, ok := other.Get_().(Batched_EndOfInput)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Batched) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Batched)
	return ok && _this.Equals(typed)
}

func Type_Batched_() _dafny.TypeDescriptor {
	return type_Batched_{}
}

type type_Batched_ struct {
}

func (_this type_Batched_) Default() interface{} {
	return Companion_Batched_.Default()
}

func (_this type_Batched_) String() string {
	return "Std_BulkActions.Batched"
}
func (_this Batched) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Batched{}

// End of datatype Batched

// Definition of trait BulkAction
type BulkAction interface {
	String() string
	Invoke(i interface{}) interface{}
	Map(input m_Std_Producers.Producer, output m_Std_Consumers.IConsumer)
}
type CompanionStruct_BulkAction_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_BulkAction_ = CompanionStruct_BulkAction_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_BulkAction_) CastTo_(x interface{}) BulkAction {
	var t BulkAction
	t, _ = x.(BulkAction)
	return t
}

// End of trait BulkAction

// Definition of class BatchReader
type BatchReader struct {
	Index     _dafny.Int
	_elements _dafny.Sequence
}

func New_BatchReader_() *BatchReader {
	_this := BatchReader{}

	_this.Index = _dafny.Zero
	_this._elements = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_BatchReader_ struct {
}

var Companion_BatchReader_ = CompanionStruct_BatchReader_{}

func (_this *BatchReader) Equals(other *BatchReader) bool {
	return _this == other
}

func (_this *BatchReader) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*BatchReader)
	return ok && _this.Equals(other)
}

func (*BatchReader) String() string {
	return "Std_BulkActions.BatchReader"
}

func Type_BatchReader_(Type_T_ _dafny.TypeDescriptor, Type_E_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_BatchReader_{Type_T_, Type_E_}
}

type type_BatchReader_ struct {
	Type_T_ _dafny.TypeDescriptor
	Type_E_ _dafny.TypeDescriptor
}

func (_this type_BatchReader_) Default() interface{} {
	return (*BatchReader)(nil)
}

func (_this type_BatchReader_) String() string {
	return "Std_BulkActions.BatchReader"
}
func (_this *BatchReader) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){m_Std_Producers.Companion_Producer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ m_Std_Producers.Producer = &BatchReader{}
var _ m_Std_Actions.Action = &BatchReader{}
var _ m_Std_Actions.TotalActionProof = &BatchReader{}
var _ m_Std_GenericActions.GenericAction = &BatchReader{}
var _ m_Std_Frames.Validatable = &BatchReader{}
var _ _dafny.TraitOffspring = &BatchReader{}

func (_this *BatchReader) Next() m_Std_Wrappers.Option {
	var _out10 m_Std_Wrappers.Option
	_ = _out10
	_out10 = m_Std_Producers.Companion_Producer_.Next(_this)
	return _out10
}
func (_this *BatchReader) Ctor__(elements _dafny.Sequence) {
	{
		(_this)._elements = elements
		(_this).Index = _dafny.Zero
	}
}
func (_this *BatchReader) ProducedCount() _dafny.Int {
	{
		return _this.Index
	}
}
func (_this *BatchReader) Remaining() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_Some_((_dafny.IntOfUint32(((_this).Elements()).Cardinality())).Minus(_this.Index))
	}
}
func (_this *BatchReader) Invoke(t interface{}) interface{} {
	{
		var t _dafny.Tuple = t.(_dafny.Tuple)
		_ = t
		var value m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Default()
		_ = value
		if (_dafny.IntOfUint32(((_this).Elements()).Cardinality())).Cmp(_this.Index) == 0 {
			value = m_Std_Wrappers.Companion_Option_.Create_None_()
		} else {
			value = m_Std_Wrappers.Companion_Option_.Create_Some_(Companion_Batched_.Create_BatchValue_(((_this).Elements()).Select((_this.Index).Uint32()).(interface{})))
			(_this).Index = (_this.Index).Plus(_dafny.One)
		}
		return value
	}
}
func (_this *BatchReader) ForEach(consumer m_Std_Consumers.IConsumer) {
	{
		var consumer m_Std_Consumers.IConsumer = consumer
		_ = consumer
		if func(_is_1 m_Std_Consumers.IConsumer) bool {
			return _dafny.InstanceOf(_is_1, (*BatchSeqWriter)(nil))
		}(consumer) {
			var _0_writer *BatchSeqWriter
			_ = _0_writer
			_0_writer = consumer.(*BatchSeqWriter)
			var _1_s _dafny.Sequence
			_ = _1_s
			var _out0 _dafny.Sequence
			_ = _out0
			_out0 = (_this).Read()
			_1_s = _out0
			(_0_writer).Elements = _dafny.Companion_Sequence_.Concatenate(_0_writer.Elements, _1_s)
			return
		}
		m_Std_Producers.Companion_Default___.DefaultForEach(_this, consumer)
	}
}
func (_this *BatchReader) Fill(consumer m_Std_Consumers.Consumer) {
	{
		var consumer m_Std_Consumers.Consumer = consumer
		_ = consumer
		m_Std_Producers.Companion_Default___.DefaultFill(_this, consumer)
	}
}
func (_this *BatchReader) Read() _dafny.Sequence {
	{
		var s _dafny.Sequence = _dafny.EmptySeq
		_ = s
		if (_this.Index).Sign() == 0 {
			s = (_this).Elements()
		} else {
			s = ((_this).Elements()).Drop((_this.Index).Uint32())
		}
		(_this).Index = _dafny.IntOfUint32(((_this).Elements()).Cardinality())
		var _0_produced _dafny.Sequence
		_ = _0_produced
		_0_produced = m_Std_Collections_Seq.Companion_Default___.Map(func(coer27 func(interface{}) Batched) func(interface{}) interface{} {
			return func(arg28 interface{}) interface{} {
				return coer27(arg28)
			}
		}(func(coer28 func(interface{}) Batched) func(interface{}) Batched {
			return func(arg29 interface{}) Batched {
				return coer28(arg29)
			}
		}(Companion_Default___.ToBatched)), s)
		return s
	}
}
func (_this *BatchReader) Elements() _dafny.Sequence {
	{
		return _this._elements
	}
}

// End of class BatchReader

// Definition of class BatchSeqWriter
type BatchSeqWriter struct {
	Elements _dafny.Sequence
	State    m_Std_Wrappers.Result
}

func New_BatchSeqWriter_() *BatchSeqWriter {
	_this := BatchSeqWriter{}

	_this.Elements = _dafny.EmptySeq
	_this.State = m_Std_Wrappers.Companion_Result_.Default(false)
	return &_this
}

type CompanionStruct_BatchSeqWriter_ struct {
}

var Companion_BatchSeqWriter_ = CompanionStruct_BatchSeqWriter_{}

func (_this *BatchSeqWriter) Equals(other *BatchSeqWriter) bool {
	return _this == other
}

func (_this *BatchSeqWriter) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*BatchSeqWriter)
	return ok && _this.Equals(other)
}

func (*BatchSeqWriter) String() string {
	return "Std_BulkActions.BatchSeqWriter"
}

func Type_BatchSeqWriter_(Type_T_ _dafny.TypeDescriptor, Type_E_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_BatchSeqWriter_{Type_T_, Type_E_}
}

type type_BatchSeqWriter_ struct {
	Type_T_ _dafny.TypeDescriptor
	Type_E_ _dafny.TypeDescriptor
}

func (_this type_BatchSeqWriter_) Default() interface{} {
	return (*BatchSeqWriter)(nil)
}

func (_this type_BatchSeqWriter_) String() string {
	return "Std_BulkActions.BatchSeqWriter"
}
func (_this *BatchSeqWriter) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){m_Std_Consumers.Companion_IConsumer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ m_Std_Consumers.IConsumer = &BatchSeqWriter{}
var _ m_Std_Actions.Action = &BatchSeqWriter{}
var _ m_Std_GenericActions.GenericAction = &BatchSeqWriter{}
var _ m_Std_Frames.Validatable = &BatchSeqWriter{}
var _ _dafny.TraitOffspring = &BatchSeqWriter{}

func (_this *BatchSeqWriter) Accept(t interface{}) {
	m_Std_Consumers.Companion_IConsumer_.Accept(_this, t)
	return
}
func (_this *BatchSeqWriter) Ctor__() {
	{
		(_this).Elements = _dafny.SeqOf()
		(_this).State = m_Std_Wrappers.Companion_Result_.Create_Success_(true)
	}
}
func (_this *BatchSeqWriter) Invoke(t interface{}) interface{} {
	{
		var t Batched = t.(Batched)
		_ = t
		var r _dafny.Tuple = _dafny.TupleOf()
		_ = r
		var _source0 Batched = t
		_ = _source0
		{
			{
				if _source0.Is_BatchValue() {
					var _0_t interface{} = _source0.Get_().(Batched_BatchValue).Value
					_ = _0_t
					(_this).Elements = _dafny.Companion_Sequence_.Concatenate(_this.Elements, _dafny.SeqOf(_0_t))
					goto Lmatch0
				}
			}
			{
				if _source0.Is_BatchError() {
					var _1_e interface{} = _source0.Get_().(Batched_BatchError).Error
					_ = _1_e
					(_this).State = m_Std_Wrappers.Companion_Result_.Create_Failure_(_1_e)
					goto Lmatch0
				}
			}
			{
				(_this).State = m_Std_Wrappers.Companion_Result_.Create_Success_(false)
			}
			goto Lmatch0
		}
	Lmatch0:
		r = _dafny.TupleOf()
		return r
	}
}
func (_this *BatchSeqWriter) Values() _dafny.Sequence {
	{
		return _this.Elements
	}
}

// End of class BatchSeqWriter

// Definition of class BatchSeqWriterTotalProof
type BatchSeqWriterTotalProof struct {
	dummy byte
}

func New_BatchSeqWriterTotalProof_() *BatchSeqWriterTotalProof {
	_this := BatchSeqWriterTotalProof{}

	return &_this
}

type CompanionStruct_BatchSeqWriterTotalProof_ struct {
}

var Companion_BatchSeqWriterTotalProof_ = CompanionStruct_BatchSeqWriterTotalProof_{}

func (_this *BatchSeqWriterTotalProof) Equals(other *BatchSeqWriterTotalProof) bool {
	return _this == other
}

func (_this *BatchSeqWriterTotalProof) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*BatchSeqWriterTotalProof)
	return ok && _this.Equals(other)
}

func (*BatchSeqWriterTotalProof) String() string {
	return "Std_BulkActions.BatchSeqWriterTotalProof"
}

func Type_BatchSeqWriterTotalProof_(Type_T_ _dafny.TypeDescriptor, Type_E_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_BatchSeqWriterTotalProof_{Type_T_, Type_E_}
}

type type_BatchSeqWriterTotalProof_ struct {
	Type_T_ _dafny.TypeDescriptor
	Type_E_ _dafny.TypeDescriptor
}

func (_this type_BatchSeqWriterTotalProof_) Default() interface{} {
	return (*BatchSeqWriterTotalProof)(nil)
}

func (_this type_BatchSeqWriterTotalProof_) String() string {
	return "Std_BulkActions.BatchSeqWriterTotalProof"
}
func (_this *BatchSeqWriterTotalProof) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ m_Std_Actions.TotalActionProof = &BatchSeqWriterTotalProof{}
var _ m_Std_Frames.Validatable = &BatchSeqWriterTotalProof{}
var _ _dafny.TraitOffspring = &BatchSeqWriterTotalProof{}

// End of class BatchSeqWriterTotalProof

// Definition of class BatchArrayWriter
type BatchArrayWriter struct {
	Storage     _dafny.Array
	Size        _dafny.Int
	OtherInputs _dafny.Int
	State       m_Std_Wrappers.Result
}

func New_BatchArrayWriter_() *BatchArrayWriter {
	_this := BatchArrayWriter{}

	_this.Storage = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.Size = _dafny.Zero
	_this.OtherInputs = _dafny.Zero
	_this.State = m_Std_Wrappers.Companion_Result_.Default(false)
	return &_this
}

type CompanionStruct_BatchArrayWriter_ struct {
}

var Companion_BatchArrayWriter_ = CompanionStruct_BatchArrayWriter_{}

func (_this *BatchArrayWriter) Equals(other *BatchArrayWriter) bool {
	return _this == other
}

func (_this *BatchArrayWriter) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*BatchArrayWriter)
	return ok && _this.Equals(other)
}

func (*BatchArrayWriter) String() string {
	return "Std_BulkActions.BatchArrayWriter"
}

func Type_BatchArrayWriter_(Type_T_ _dafny.TypeDescriptor, Type_E_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_BatchArrayWriter_{Type_T_, Type_E_}
}

type type_BatchArrayWriter_ struct {
	Type_T_ _dafny.TypeDescriptor
	Type_E_ _dafny.TypeDescriptor
}

func (_this type_BatchArrayWriter_) Default() interface{} {
	return (*BatchArrayWriter)(nil)
}

func (_this type_BatchArrayWriter_) String() string {
	return "Std_BulkActions.BatchArrayWriter"
}
func (_this *BatchArrayWriter) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){m_Std_Consumers.Companion_Consumer_.TraitID_, m_Std_Actions.Companion_Action_.TraitID_, m_Std_GenericActions.Companion_GenericAction_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ m_Std_Consumers.Consumer = &BatchArrayWriter{}
var _ m_Std_Actions.Action = &BatchArrayWriter{}
var _ m_Std_GenericActions.GenericAction = &BatchArrayWriter{}
var _ m_Std_Frames.Validatable = &BatchArrayWriter{}
var _ _dafny.TraitOffspring = &BatchArrayWriter{}

func (_this *BatchArrayWriter) Accept(t interface{}) bool {
	var _out3 bool
	_ = _out3
	_out3 = m_Std_Consumers.Companion_Consumer_.Accept(_this, t)
	return _out3
}
func (_this *BatchArrayWriter) Ctor__(storage _dafny.Array) {
	{
		(_this).Storage = storage
		(_this).Size = _dafny.Zero
		(_this).OtherInputs = _dafny.Zero
		(_this).State = m_Std_Wrappers.Companion_Result_.Create_Success_(true)
	}
}
func (_this *BatchArrayWriter) Capacity() m_Std_Wrappers.Option {
	{
		return m_Std_Wrappers.Companion_Option_.Create_Some_(((_dafny.ArrayLen((_this.Storage), 0)).Minus(_this.Size)).Minus(_this.OtherInputs))
	}
}
func (_this *BatchArrayWriter) Invoke(t interface{}) interface{} {
	{
		var t Batched = t.(Batched)
		_ = t
		var r bool = false
		_ = r
		if ((_this.Size).Plus(_this.OtherInputs)).Cmp(_dafny.ArrayLen((_this.Storage), 0)) == 0 {
			r = false
		} else {
			var _source0 Batched = t
			_ = _source0
			{
				{
					if _source0.Is_BatchValue() {
						var _0_value interface{} = _source0.Get_().(Batched_BatchValue).Value
						_ = _0_value
						var _arr0 _dafny.Array = _this.Storage
						_ = _arr0
						var _index0 _dafny.Int = _this.Size
						_ = _index0
						(_arr0).ArraySet1(_0_value, (_index0).Int())
						(_this).Size = (_this.Size).Plus(_dafny.One)
						goto Lmatch0
					}
				}
				{
					if _source0.Is_BatchError() {
						var _1_e interface{} = _source0.Get_().(Batched_BatchError).Error
						_ = _1_e
						(_this).State = m_Std_Wrappers.Companion_Result_.Create_Failure_(_1_e)
						(_this).OtherInputs = (_this.OtherInputs).Plus(_dafny.One)
						goto Lmatch0
					}
				}
				{
					(_this).State = m_Std_Wrappers.Companion_Result_.Create_Success_(false)
					(_this).OtherInputs = (_this.OtherInputs).Plus(_dafny.One)
				}
				goto Lmatch0
			}
		Lmatch0:
			r = true
		}
		return r
	}
}
func (_this *BatchArrayWriter) Values() _dafny.Sequence {
	{
		return _dafny.ArrayRangeToSeq(_this.Storage, _dafny.NilInt, (_this.Size))
	}
}

// End of class BatchArrayWriter

// Definition of class BatchArrayWriterTotalProof
type BatchArrayWriterTotalProof struct {
	dummy byte
}

func New_BatchArrayWriterTotalProof_() *BatchArrayWriterTotalProof {
	_this := BatchArrayWriterTotalProof{}

	return &_this
}

type CompanionStruct_BatchArrayWriterTotalProof_ struct {
}

var Companion_BatchArrayWriterTotalProof_ = CompanionStruct_BatchArrayWriterTotalProof_{}

func (_this *BatchArrayWriterTotalProof) Equals(other *BatchArrayWriterTotalProof) bool {
	return _this == other
}

func (_this *BatchArrayWriterTotalProof) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*BatchArrayWriterTotalProof)
	return ok && _this.Equals(other)
}

func (*BatchArrayWriterTotalProof) String() string {
	return "Std_BulkActions.BatchArrayWriterTotalProof"
}

func Type_BatchArrayWriterTotalProof_(Type_T_ _dafny.TypeDescriptor, Type_E_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_BatchArrayWriterTotalProof_{Type_T_, Type_E_}
}

type type_BatchArrayWriterTotalProof_ struct {
	Type_T_ _dafny.TypeDescriptor
	Type_E_ _dafny.TypeDescriptor
}

func (_this type_BatchArrayWriterTotalProof_) Default() interface{} {
	return (*BatchArrayWriterTotalProof)(nil)
}

func (_this type_BatchArrayWriterTotalProof_) String() string {
	return "Std_BulkActions.BatchArrayWriterTotalProof"
}
func (_this *BatchArrayWriterTotalProof) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){m_Std_Actions.Companion_TotalActionProof_.TraitID_, m_Std_Frames.Companion_Validatable_.TraitID_}
}

var _ m_Std_Actions.TotalActionProof = &BatchArrayWriterTotalProof{}
var _ m_Std_Frames.Validatable = &BatchArrayWriterTotalProof{}
var _ _dafny.TraitOffspring = &BatchArrayWriterTotalProof{}

// End of class BatchArrayWriterTotalProof
