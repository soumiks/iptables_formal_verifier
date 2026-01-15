// Package Std_Termination
// Dafny module Std_Termination compiled into Go

package Std_Termination

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
	m_Std_Collections_Set "Std_Collections_Set"
	m_Std_Collections_Tuple "Std_Collections_Tuple"
	m_Std_Functions "Std_Functions"
	m_Std_Math "Std_Math"
	m_Std_Ordinal "Std_Ordinal"
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
var _ m_Std_Collections_Set.Dummy__
var _ m_Std_Collections_Tuple.Dummy__
var _ m_Std_Ordinal.Dummy__

type Dummy__ struct{}

// Definition of datatype TerminationMetric
type TerminationMetric struct {
	Data_TerminationMetric_
}

func (_this TerminationMetric) Get_() Data_TerminationMetric_ {
	return _this.Data_TerminationMetric_
}

type Data_TerminationMetric_ interface {
	isTerminationMetric()
}

type CompanionStruct_TerminationMetric_ struct {
}

var Companion_TerminationMetric_ = CompanionStruct_TerminationMetric_{}

type TerminationMetric_TMBool struct {
	BoolValue bool
}

func (TerminationMetric_TMBool) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMBool_(BoolValue bool) TerminationMetric {
	return TerminationMetric{TerminationMetric_TMBool{BoolValue}}
}

func (_this TerminationMetric) Is_TMBool() bool {
	_, ok := _this.Get_().(TerminationMetric_TMBool)
	return ok
}

type TerminationMetric_TMNat struct {
	NatValue _dafny.Int
}

func (TerminationMetric_TMNat) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMNat_(NatValue _dafny.Int) TerminationMetric {
	return TerminationMetric{TerminationMetric_TMNat{NatValue}}
}

func (_this TerminationMetric) Is_TMNat() bool {
	_, ok := _this.Get_().(TerminationMetric_TMNat)
	return ok
}

type TerminationMetric_TMChar struct {
	CharValue _dafny.Int
}

func (TerminationMetric_TMChar) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMChar_(CharValue _dafny.Int) TerminationMetric {
	return TerminationMetric{TerminationMetric_TMChar{CharValue}}
}

func (_this TerminationMetric) Is_TMChar() bool {
	_, ok := _this.Get_().(TerminationMetric_TMChar)
	return ok
}

type TerminationMetric_TMOrdinal struct {
	OrdinalValue _dafny.Ord
}

func (TerminationMetric_TMOrdinal) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMOrdinal_(OrdinalValue _dafny.Ord) TerminationMetric {
	return TerminationMetric{TerminationMetric_TMOrdinal{OrdinalValue}}
}

func (_this TerminationMetric) Is_TMOrdinal() bool {
	_, ok := _this.Get_().(TerminationMetric_TMOrdinal)
	return ok
}

type TerminationMetric_TMObject struct {
	ObjectValue interface{}
}

func (TerminationMetric_TMObject) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMObject_(ObjectValue interface{}) TerminationMetric {
	return TerminationMetric{TerminationMetric_TMObject{ObjectValue}}
}

func (_this TerminationMetric) Is_TMObject() bool {
	_, ok := _this.Get_().(TerminationMetric_TMObject)
	return ok
}

type TerminationMetric_TMSeq struct {
	SeqValue _dafny.Sequence
}

func (TerminationMetric_TMSeq) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMSeq_(SeqValue _dafny.Sequence) TerminationMetric {
	return TerminationMetric{TerminationMetric_TMSeq{SeqValue}}
}

func (_this TerminationMetric) Is_TMSeq() bool {
	_, ok := _this.Get_().(TerminationMetric_TMSeq)
	return ok
}

type TerminationMetric_TMSet struct {
	SetValue _dafny.Set
}

func (TerminationMetric_TMSet) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMSet_(SetValue _dafny.Set) TerminationMetric {
	return TerminationMetric{TerminationMetric_TMSet{SetValue}}
}

func (_this TerminationMetric) Is_TMSet() bool {
	_, ok := _this.Get_().(TerminationMetric_TMSet)
	return ok
}

type TerminationMetric_TMTuple struct {
	Base   TerminationMetric
	First  TerminationMetric
	Second TerminationMetric
}

func (TerminationMetric_TMTuple) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMTuple_(Base TerminationMetric, First TerminationMetric, Second TerminationMetric) TerminationMetric {
	return TerminationMetric{TerminationMetric_TMTuple{Base, First, Second}}
}

func (_this TerminationMetric) Is_TMTuple() bool {
	_, ok := _this.Get_().(TerminationMetric_TMTuple)
	return ok
}

type TerminationMetric_TMTop struct {
}

func (TerminationMetric_TMTop) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMTop_() TerminationMetric {
	return TerminationMetric{TerminationMetric_TMTop{}}
}

func (_this TerminationMetric) Is_TMTop() bool {
	_, ok := _this.Get_().(TerminationMetric_TMTop)
	return ok
}

type TerminationMetric_TMSucc struct {
	Original TerminationMetric
}

func (TerminationMetric_TMSucc) isTerminationMetric() {}

func (CompanionStruct_TerminationMetric_) Create_TMSucc_(Original TerminationMetric) TerminationMetric {
	return TerminationMetric{TerminationMetric_TMSucc{Original}}
}

func (_this TerminationMetric) Is_TMSucc() bool {
	_, ok := _this.Get_().(TerminationMetric_TMSucc)
	return ok
}

func (CompanionStruct_TerminationMetric_) Default() TerminationMetric {
	return Companion_TerminationMetric_.Create_TMBool_(false)
}

func (_this TerminationMetric) Dtor_boolValue() bool {
	return _this.Get_().(TerminationMetric_TMBool).BoolValue
}

func (_this TerminationMetric) Dtor_natValue() _dafny.Int {
	return _this.Get_().(TerminationMetric_TMNat).NatValue
}

func (_this TerminationMetric) Dtor_charValue() _dafny.Int {
	return _this.Get_().(TerminationMetric_TMChar).CharValue
}

func (_this TerminationMetric) Dtor_ordinalValue() _dafny.Ord {
	return _this.Get_().(TerminationMetric_TMOrdinal).OrdinalValue
}

func (_this TerminationMetric) Dtor_objectValue() interface{} {
	return _this.Get_().(TerminationMetric_TMObject).ObjectValue
}

func (_this TerminationMetric) Dtor_seqValue() _dafny.Sequence {
	return _this.Get_().(TerminationMetric_TMSeq).SeqValue
}

func (_this TerminationMetric) Dtor_setValue() _dafny.Set {
	return _this.Get_().(TerminationMetric_TMSet).SetValue
}

func (_this TerminationMetric) Dtor_base() TerminationMetric {
	return _this.Get_().(TerminationMetric_TMTuple).Base
}

func (_this TerminationMetric) Dtor_first() TerminationMetric {
	return _this.Get_().(TerminationMetric_TMTuple).First
}

func (_this TerminationMetric) Dtor_second() TerminationMetric {
	return _this.Get_().(TerminationMetric_TMTuple).Second
}

func (_this TerminationMetric) Dtor_original() TerminationMetric {
	return _this.Get_().(TerminationMetric_TMSucc).Original
}

func (_this TerminationMetric) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case TerminationMetric_TMBool:
		{
			return "Termination.TerminationMetric.TMBool" + "(" + _dafny.String(data.BoolValue) + ")"
		}
	case TerminationMetric_TMNat:
		{
			return "Termination.TerminationMetric.TMNat" + "(" + _dafny.String(data.NatValue) + ")"
		}
	case TerminationMetric_TMChar:
		{
			return "Termination.TerminationMetric.TMChar" + "(" + _dafny.String(data.CharValue) + ")"
		}
	case TerminationMetric_TMOrdinal:
		{
			return "Termination.TerminationMetric.TMOrdinal" + "(" + _dafny.String(data.OrdinalValue) + ")"
		}
	case TerminationMetric_TMObject:
		{
			return "Termination.TerminationMetric.TMObject" + "(" + _dafny.String(data.ObjectValue) + ")"
		}
	case TerminationMetric_TMSeq:
		{
			return "Termination.TerminationMetric.TMSeq" + "(" + _dafny.String(data.SeqValue) + ")"
		}
	case TerminationMetric_TMSet:
		{
			return "Termination.TerminationMetric.TMSet" + "(" + _dafny.String(data.SetValue) + ")"
		}
	case TerminationMetric_TMTuple:
		{
			return "Termination.TerminationMetric.TMTuple" + "(" + _dafny.String(data.Base) + ", " + _dafny.String(data.First) + ", " + _dafny.String(data.Second) + ")"
		}
	case TerminationMetric_TMTop:
		{
			return "Termination.TerminationMetric.TMTop"
		}
	case TerminationMetric_TMSucc:
		{
			return "Termination.TerminationMetric.TMSucc" + "(" + _dafny.String(data.Original) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this TerminationMetric) Equals(other TerminationMetric) bool {
	switch data1 := _this.Get_().(type) {
	case TerminationMetric_TMBool:
		{
			data2, ok := other.Get_().(TerminationMetric_TMBool)
			return ok && data1.BoolValue == data2.BoolValue
		}
	case TerminationMetric_TMNat:
		{
			data2, ok := other.Get_().(TerminationMetric_TMNat)
			return ok && data1.NatValue.Cmp(data2.NatValue) == 0
		}
	case TerminationMetric_TMChar:
		{
			data2, ok := other.Get_().(TerminationMetric_TMChar)
			return ok && data1.CharValue.Cmp(data2.CharValue) == 0
		}
	case TerminationMetric_TMOrdinal:
		{
			data2, ok := other.Get_().(TerminationMetric_TMOrdinal)
			return ok && data1.OrdinalValue.Cmp(data2.OrdinalValue) == 0
		}
	case TerminationMetric_TMObject:
		{
			data2, ok := other.Get_().(TerminationMetric_TMObject)
			return ok && _dafny.AreEqual(data1.ObjectValue, data2.ObjectValue)
		}
	case TerminationMetric_TMSeq:
		{
			data2, ok := other.Get_().(TerminationMetric_TMSeq)
			return ok && data1.SeqValue.Equals(data2.SeqValue)
		}
	case TerminationMetric_TMSet:
		{
			data2, ok := other.Get_().(TerminationMetric_TMSet)
			return ok && data1.SetValue.Equals(data2.SetValue)
		}
	case TerminationMetric_TMTuple:
		{
			data2, ok := other.Get_().(TerminationMetric_TMTuple)
			return ok && data1.Base.Equals(data2.Base) && data1.First.Equals(data2.First) && data1.Second.Equals(data2.Second)
		}
	case TerminationMetric_TMTop:
		{
			_, ok := other.Get_().(TerminationMetric_TMTop)
			return ok
		}
	case TerminationMetric_TMSucc:
		{
			data2, ok := other.Get_().(TerminationMetric_TMSucc)
			return ok && data1.Original.Equals(data2.Original)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this TerminationMetric) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(TerminationMetric)
	return ok && _this.Equals(typed)
}

func Type_TerminationMetric_() _dafny.TypeDescriptor {
	return type_TerminationMetric_{}
}

type type_TerminationMetric_ struct {
}

func (_this type_TerminationMetric_) Default() interface{} {
	return Companion_TerminationMetric_.Default()
}

func (_this type_TerminationMetric_) String() string {
	return "Std_Termination.TerminationMetric"
}
func (_this TerminationMetric) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = TerminationMetric{}

// End of datatype TerminationMetric
