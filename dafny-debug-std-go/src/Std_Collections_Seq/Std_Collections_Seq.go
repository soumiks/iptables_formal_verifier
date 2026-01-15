// Package Std_Collections_Seq
// Dafny module Std_Collections_Seq compiled into Go

package Std_Collections_Seq

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
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
	return "Std_Collections_Seq.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) First(xs _dafny.Sequence) interface{} {
	return (xs).Select(0).(interface{})
}
func (_static *CompanionStruct_Default___) DropFirst(xs _dafny.Sequence) _dafny.Sequence {
	return (xs).Drop(1)
}
func (_static *CompanionStruct_Default___) Last(xs _dafny.Sequence) interface{} {
	return (xs).Select(((_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)).Uint32()).(interface{})
}
func (_static *CompanionStruct_Default___) DropLast(xs _dafny.Sequence) _dafny.Sequence {
	return (xs).Take(((_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)).Uint32())
}
func (_static *CompanionStruct_Default___) ToArray(xs _dafny.Sequence) _dafny.Array {
	var a _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = a
	var _len0_0 _dafny.Int = _dafny.IntOfUint32((xs).Cardinality())
	_ = _len0_0
	var _nw0 _dafny.Array
	_ = _nw0
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw0 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) interface{} = (func(_0_xs _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(_1_i _dafny.Int) interface{} {
				return (_0_xs).Select((_1_i).Uint32()).(interface{})
			}
		})(xs)
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
	a = _nw0
	return a
}
func (_static *CompanionStruct_Default___) ToSet(xs _dafny.Sequence) _dafny.Set {
	return func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((xs).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_x interface{}
			_0_x = interface{}(_compr_0).(interface{})
			if _dafny.Companion_Sequence_.Contains(xs, _0_x) {
				_coll0.Add(_0_x)
			}
		}
		return _coll0.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) IndexOf(xs _dafny.Sequence, v interface{}) _dafny.Int {
	var _0___accumulator _dafny.Int = _dafny.Zero
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if _dafny.AreEqual((xs).Select(0).(interface{}), v) {
		return (_dafny.Zero).Plus(_0___accumulator)
	} else {
		_0___accumulator = (_0___accumulator).Plus(_dafny.One)
		var _in0 _dafny.Sequence = (xs).Drop(1)
		_ = _in0
		var _in1 interface{} = v
		_ = _in1
		xs = _in0
		v = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) IndexOfOption(xs _dafny.Sequence, v interface{}) m_Std_Wrappers.Option {
	return Companion_Default___.IndexByOption(xs, func(coer0 func(interface{}) bool) func(interface{}) bool {
		return func(arg0 interface{}) bool {
			return coer0(arg0)
		}
	}((func(_0_v interface{}) func(interface{}) bool {
		return func(_1_x interface{}) bool {
			return _dafny.AreEqual(_1_x, _0_v)
		}
	})(v)))
}
func (_static *CompanionStruct_Default___) IndexByOption(xs _dafny.Sequence, p func(interface{}) bool) m_Std_Wrappers.Option {
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return m_Std_Wrappers.Companion_Option_.Create_None_()
	} else if (p)((xs).Select(0).(interface{})) {
		return m_Std_Wrappers.Companion_Option_.Create_Some_(_dafny.Zero)
	} else {
		var _0_o_k m_Std_Wrappers.Option = Companion_Default___.IndexByOption((xs).Drop(1), p)
		_ = _0_o_k
		if (_0_o_k).Is_Some() {
			return m_Std_Wrappers.Companion_Option_.Create_Some_(((_0_o_k).Dtor_value().(_dafny.Int)).Plus(_dafny.One))
		} else {
			return m_Std_Wrappers.Companion_Option_.Create_None_()
		}
	}
}
func (_static *CompanionStruct_Default___) LastIndexOf(xs _dafny.Sequence, v interface{}) _dafny.Int {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if _dafny.AreEqual((xs).Select(((_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)).Uint32()).(interface{}), v) {
		return (_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)
	} else {
		var _in0 _dafny.Sequence = (xs).Take(((_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)).Uint32())
		_ = _in0
		var _in1 interface{} = v
		_ = _in1
		xs = _in0
		v = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) LastIndexOfOption(xs _dafny.Sequence, v interface{}) m_Std_Wrappers.Option {
	return Companion_Default___.LastIndexByOption(xs, func(coer1 func(interface{}) bool) func(interface{}) bool {
		return func(arg1 interface{}) bool {
			return coer1(arg1)
		}
	}((func(_0_v interface{}) func(interface{}) bool {
		return func(_1_x interface{}) bool {
			return _dafny.AreEqual(_1_x, _0_v)
		}
	})(v)))
}
func (_static *CompanionStruct_Default___) LastIndexByOption(xs _dafny.Sequence, p func(interface{}) bool) m_Std_Wrappers.Option {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return m_Std_Wrappers.Companion_Option_.Create_None_()
	} else if (p)((xs).Select(((_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)).Uint32()).(interface{})) {
		return m_Std_Wrappers.Companion_Option_.Create_Some_((_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One))
	} else {
		var _in0 _dafny.Sequence = (xs).Take(((_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)).Uint32())
		_ = _in0
		var _in1 func(interface{}) bool = p
		_ = _in1
		xs = _in0
		p = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Remove(xs _dafny.Sequence, pos _dafny.Int) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((xs).Take((pos).Uint32()), (xs).Drop(((pos).Plus(_dafny.One)).Uint32()))
}
func (_static *CompanionStruct_Default___) RemoveValue(xs _dafny.Sequence, v interface{}) _dafny.Sequence {
	if !_dafny.Companion_Sequence_.Contains(xs, v) {
		return xs
	} else {
		var _0_i _dafny.Int = Companion_Default___.IndexOf(xs, v)
		_ = _0_i
		return _dafny.Companion_Sequence_.Concatenate((xs).Take((_0_i).Uint32()), (xs).Drop(((_0_i).Plus(_dafny.One)).Uint32()))
	}
}
func (_static *CompanionStruct_Default___) Insert(xs _dafny.Sequence, a interface{}, pos _dafny.Int) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((xs).Take((pos).Uint32()), _dafny.SeqOf(a)), (xs).Drop((pos).Uint32()))
}
func (_static *CompanionStruct_Default___) Reverse(xs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if _dafny.Companion_Sequence_.Equal(xs, _dafny.SeqOf()) {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((xs).Select(((_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)).Uint32()).(interface{})))
		var _in0 _dafny.Sequence = (xs).Subsequence(0, ((_dafny.IntOfUint32((xs).Cardinality())).Minus(_dafny.One)).Uint32())
		_ = _in0
		xs = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Repeat(v interface{}, length _dafny.Int) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (length).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(v))
		var _in0 interface{} = v
		_ = _in0
		var _in1 _dafny.Int = (length).Minus(_dafny.One)
		_ = _in1
		v = _in0
		length = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Unzip(xs _dafny.Sequence) _dafny.Tuple {
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.TupleOf(_dafny.SeqOf(), _dafny.SeqOf())
	} else {
		var _let_tmp_rhs0 _dafny.Tuple = Companion_Default___.Unzip(Companion_Default___.DropLast(xs))
		_ = _let_tmp_rhs0
		var _0_a _dafny.Sequence = (*(_let_tmp_rhs0).IndexInt(0)).(_dafny.Sequence)
		_ = _0_a
		var _1_b _dafny.Sequence = (*(_let_tmp_rhs0).IndexInt(1)).(_dafny.Sequence)
		_ = _1_b
		return _dafny.TupleOf(_dafny.Companion_Sequence_.Concatenate(_0_a, _dafny.SeqOf((*(Companion_Default___.Last(xs).(_dafny.Tuple)).IndexInt(0)))), _dafny.Companion_Sequence_.Concatenate(_1_b, _dafny.SeqOf((*(Companion_Default___.Last(xs).(_dafny.Tuple)).IndexInt(1)))))
	}
}
func (_static *CompanionStruct_Default___) Zip(xs _dafny.Sequence, ys _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(), _0___accumulator)
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.TupleOf(Companion_Default___.Last(xs), Companion_Default___.Last(ys))), _0___accumulator)
		var _in0 _dafny.Sequence = Companion_Default___.DropLast(xs)
		_ = _in0
		var _in1 _dafny.Sequence = Companion_Default___.DropLast(ys)
		_ = _in1
		xs = _in0
		ys = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Max(xs _dafny.Sequence) _dafny.Int {
	if (_dafny.IntOfUint32((xs).Cardinality())).Cmp(_dafny.One) == 0 {
		return (xs).Select(0).(_dafny.Int)
	} else {
		return m_Std_Math.Companion_Default___.Max((xs).Select(0).(_dafny.Int), Companion_Default___.Max((xs).Drop(1)))
	}
}
func (_static *CompanionStruct_Default___) Min(xs _dafny.Sequence) _dafny.Int {
	if (_dafny.IntOfUint32((xs).Cardinality())).Cmp(_dafny.One) == 0 {
		return (xs).Select(0).(_dafny.Int)
	} else {
		return m_Std_Math.Companion_Default___.Min((xs).Select(0).(_dafny.Int), Companion_Default___.Min((xs).Drop(1)))
	}
}
func (_static *CompanionStruct_Default___) Flatten(xs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (xs).Select(0).(_dafny.Sequence))
		var _in0 _dafny.Sequence = (xs).Drop(1)
		_ = _in0
		xs = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) FlattenReverse(xs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(), _0___accumulator)
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Last(xs).(_dafny.Sequence), _0___accumulator)
		var _in0 _dafny.Sequence = Companion_Default___.DropLast(xs)
		_ = _in0
		xs = _in0
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Join(seqs _dafny.Sequence, separator _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((seqs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else if (_dafny.IntOfUint32((seqs).Cardinality())).Cmp(_dafny.One) == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (seqs).Select(0).(_dafny.Sequence))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.Companion_Sequence_.Concatenate((seqs).Select(0).(_dafny.Sequence), separator))
		var _in0 _dafny.Sequence = (seqs).Drop(1)
		_ = _in0
		var _in1 _dafny.Sequence = separator
		_ = _in1
		seqs = _in0
		separator = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) Split(s _dafny.Sequence, delim interface{}) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _1_i m_Std_Wrappers.Option = Companion_Default___.IndexOfOption(s, delim)
	_ = _1_i
	if (_1_i).Is_Some() {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((s).Take(((_1_i).Dtor_value().(_dafny.Int)).Uint32())))
		var _in0 _dafny.Sequence = (s).Drop((((_1_i).Dtor_value().(_dafny.Int)).Plus(_dafny.One)).Uint32())
		_ = _in0
		var _in1 interface{} = delim
		_ = _in1
		s = _in0
		delim = _in1
		goto TAIL_CALL_START
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf(s))
	}
}
func (_static *CompanionStruct_Default___) SplitOnce(s _dafny.Sequence, delim interface{}) _dafny.Tuple {
	var _0_i m_Std_Wrappers.Option = Companion_Default___.IndexOfOption(s, delim)
	_ = _0_i
	return _dafny.TupleOf((s).Take(((_0_i).Dtor_value().(_dafny.Int)).Uint32()), (s).Drop((((_0_i).Dtor_value().(_dafny.Int)).Plus(_dafny.One)).Uint32()))
}
func (_static *CompanionStruct_Default___) SplitOnceOption(s _dafny.Sequence, delim interface{}) m_Std_Wrappers.Option {
	var _0_valueOrError0 m_Std_Wrappers.Option = Companion_Default___.IndexOfOption(s, delim)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_i _dafny.Int = (_0_valueOrError0).Extract().(_dafny.Int)
		_ = _1_i
		return m_Std_Wrappers.Companion_Option_.Create_Some_(_dafny.TupleOf((s).Take((_1_i).Uint32()), (s).Drop(((_1_i).Plus(_dafny.One)).Uint32())))
	}
}
func (_static *CompanionStruct_Default___) Map(f func(interface{}) interface{}, xs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((f)((xs).Select(0).(interface{}))))
		var _in0 func(interface{}) interface{} = f
		_ = _in0
		var _in1 _dafny.Sequence = (xs).Drop(1)
		_ = _in1
		f = _in0
		xs = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) MapPartialFunction(f func(interface{}) interface{}, xs _dafny.Sequence) _dafny.Sequence {
	return Companion_Default___.Map(func(coer2 func(interface{}) interface{}) func(interface{}) interface{} {
		return func(arg2 interface{}) interface{} {
			return coer2(arg2)
		}
	}(f), xs)
}
func (_static *CompanionStruct_Default___) MapWithResult(f func(interface{}) m_Std_Wrappers.Result, xs _dafny.Sequence) m_Std_Wrappers.Result {
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.SeqOf())
	} else {
		var _0_valueOrError0 m_Std_Wrappers.Result = (f)((xs).Select(0).(interface{}))
		_ = _0_valueOrError0
		if (_0_valueOrError0).IsFailure() {
			return (_0_valueOrError0).PropagateFailure()
		} else {
			var _1_head interface{} = (_0_valueOrError0).Extract()
			_ = _1_head
			var _2_valueOrError1 m_Std_Wrappers.Result = Companion_Default___.MapWithResult(f, (xs).Drop(1))
			_ = _2_valueOrError1
			if (_2_valueOrError1).IsFailure() {
				return (_2_valueOrError1).PropagateFailure()
			} else {
				var _3_tail _dafny.Sequence = (_2_valueOrError1).Extract().(_dafny.Sequence)
				_ = _3_tail
				return m_Std_Wrappers.Companion_Result_.Create_Success_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1_head), _3_tail))
			}
		}
	}
}
func (_static *CompanionStruct_Default___) Filter(f func(interface{}) bool, xs _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf())
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (func() _dafny.Sequence {
			if (f)((xs).Select(0).(interface{})) {
				return _dafny.SeqOf((xs).Select(0).(interface{}))
			}
			return _dafny.SeqOf()
		})())
		var _in0 func(interface{}) bool = f
		_ = _in0
		var _in1 _dafny.Sequence = (xs).Drop(1)
		_ = _in1
		f = _in0
		xs = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) FoldLeft(f func(interface{}, interface{}) interface{}, init interface{}, xs _dafny.Sequence) interface{} {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return init
	} else {
		var _in0 func(interface{}, interface{}) interface{} = f
		_ = _in0
		var _in1 interface{} = (f)(init, (xs).Select(0).(interface{}))
		_ = _in1
		var _in2 _dafny.Sequence = (xs).Drop(1)
		_ = _in2
		f = _in0
		init = _in1
		xs = _in2
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) FoldRight(f func(interface{}, interface{}) interface{}, xs _dafny.Sequence, init interface{}) interface{} {
	if (_dafny.IntOfUint32((xs).Cardinality())).Sign() == 0 {
		return init
	} else {
		return (f)((xs).Select(0).(interface{}), Companion_Default___.FoldRight(f, (xs).Drop(1), init))
	}
}
func (_static *CompanionStruct_Default___) MergeSortBy(lessThanOrEq func(interface{}, interface{}) bool, a _dafny.Sequence) _dafny.Sequence {
	if (_dafny.IntOfUint32((a).Cardinality())).Cmp(_dafny.One) <= 0 {
		return a
	} else {
		var _0_splitIndex _dafny.Int = (_dafny.IntOfUint32((a).Cardinality())).DivBy(_dafny.IntOfInt64(2))
		_ = _0_splitIndex
		var _1_left _dafny.Sequence = (a).Take((_0_splitIndex).Uint32())
		_ = _1_left
		var _2_right _dafny.Sequence = (a).Drop((_0_splitIndex).Uint32())
		_ = _2_right
		var _3_leftSorted _dafny.Sequence = Companion_Default___.MergeSortBy(lessThanOrEq, _1_left)
		_ = _3_leftSorted
		var _4_rightSorted _dafny.Sequence = Companion_Default___.MergeSortBy(lessThanOrEq, _2_right)
		_ = _4_rightSorted
		return Companion_Default___.MergeSortedWith(_3_leftSorted, _4_rightSorted, func(coer3 func(interface{}, interface{}) bool) func(interface{}, interface{}) bool {
			return func(arg3 interface{}, arg4 interface{}) bool {
				return coer3(arg3, arg4)
			}
		}(lessThanOrEq))
	}
}
func (_static *CompanionStruct_Default___) MergeSortedWith(left _dafny.Sequence, right _dafny.Sequence, lessThanOrEq func(interface{}, interface{}) bool) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((left).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, right)
	} else if (_dafny.IntOfUint32((right).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, left)
	} else if (lessThanOrEq)((left).Select(0).(interface{}), (right).Select(0).(interface{})) {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((left).Select(0).(interface{})))
		var _in0 _dafny.Sequence = (left).Drop(1)
		_ = _in0
		var _in1 _dafny.Sequence = right
		_ = _in1
		var _in2 func(interface{}, interface{}) bool = lessThanOrEq
		_ = _in2
		left = _in0
		right = _in1
		lessThanOrEq = _in2
		goto TAIL_CALL_START
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.SeqOf((right).Select(0).(interface{})))
		var _in3 _dafny.Sequence = left
		_ = _in3
		var _in4 _dafny.Sequence = (right).Drop(1)
		_ = _in4
		var _in5 func(interface{}, interface{}) bool = lessThanOrEq
		_ = _in5
		left = _in3
		right = _in4
		lessThanOrEq = _in5
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) All(s _dafny.Sequence, p func(interface{}) bool) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((s).Cardinality())), true, func(_forall_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_forall_var_0).(_dafny.Int)
		return !(((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0)) || ((p)((s).Select((_0_i).Uint32()).(interface{})))
	})
}
func (_static *CompanionStruct_Default___) AllNot(s _dafny.Sequence, p func(interface{}) bool) bool {
	return _dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((s).Cardinality())), true, func(_forall_var_0 _dafny.Int) bool {
		var _0_i _dafny.Int
		_0_i = interface{}(_forall_var_0).(_dafny.Int)
		return !(((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((s).Cardinality())) < 0)) || (!((p)((s).Select((_0_i).Uint32()).(interface{}))))
	})
}
func (_static *CompanionStruct_Default___) Partitioned(s _dafny.Sequence, p func(interface{}) bool) bool {
	goto TAIL_CALL_START
TAIL_CALL_START:
	if _dafny.Companion_Sequence_.Equal(s, _dafny.SeqOf()) {
		return true
	} else if (p)((s).Select(0).(interface{})) {
		var _in0 _dafny.Sequence = (s).Drop(1)
		_ = _in0
		var _in1 func(interface{}) bool = p
		_ = _in1
		s = _in0
		p = _in1
		goto TAIL_CALL_START
	} else {
		return Companion_Default___.AllNot((s).Drop(1), func(coer4 func(interface{}) bool) func(interface{}) bool {
			return func(arg5 interface{}) bool {
				return coer4(arg5)
			}
		}(p))
	}
}

// End of class Default__

// Definition of datatype Slice
type Slice struct {
	Data_Slice_
}

func (_this Slice) Get_() Data_Slice_ {
	return _this.Data_Slice_
}

type Data_Slice_ interface {
	isSlice()
}

type CompanionStruct_Slice_ struct {
}

var Companion_Slice_ = CompanionStruct_Slice_{}

type Slice_Slice struct {
	Data  _dafny.Sequence
	Start _dafny.Int
	End   _dafny.Int
}

func (Slice_Slice) isSlice() {}

func (CompanionStruct_Slice_) Create_Slice_(Data _dafny.Sequence, Start _dafny.Int, End _dafny.Int) Slice {
	return Slice{Slice_Slice{Data, Start, End}}
}

func (_this Slice) Is_Slice() bool {
	_, ok := _this.Get_().(Slice_Slice)
	return ok
}

func (CompanionStruct_Slice_) Default() Slice {
	return Companion_Slice_.Create_Slice_(_dafny.EmptySeq, _dafny.Zero, _dafny.Zero)
}

func (_this Slice) Dtor_data() _dafny.Sequence {
	return _this.Get_().(Slice_Slice).Data
}

func (_this Slice) Dtor_start() _dafny.Int {
	return _this.Get_().(Slice_Slice).Start
}

func (_this Slice) Dtor_end() _dafny.Int {
	return _this.Get_().(Slice_Slice).End
}

func (_this Slice) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Slice_Slice:
		{
			return "Seq.Slice.Slice" + "(" + _dafny.String(data.Data) + ", " + _dafny.String(data.Start) + ", " + _dafny.String(data.End) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Slice) Equals(other Slice) bool {
	switch data1 := _this.Get_().(type) {
	case Slice_Slice:
		{
			data2, ok := other.Get_().(Slice_Slice)
			return ok && data1.Data.Equals(data2.Data) && data1.Start.Cmp(data2.Start) == 0 && data1.End.Cmp(data2.End) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Slice) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Slice)
	return ok && _this.Equals(typed)
}

func Type_Slice_() _dafny.TypeDescriptor {
	return type_Slice_{}
}

type type_Slice_ struct {
}

func (_this type_Slice_) Default() interface{} {
	return Companion_Slice_.Default()
}

func (_this type_Slice_) String() string {
	return "Std_Collections_Seq.Slice"
}
func (_this Slice) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Slice{}

func (_this Slice) View() _dafny.Sequence {
	{
		return ((_this).Dtor_data()).Subsequence(((_this).Dtor_start()).Uint32(), ((_this).Dtor_end()).Uint32())
	}
}
func (_this Slice) Length() _dafny.Int {
	{
		return ((_this).Dtor_end()).Minus((_this).Dtor_start())
	}
}
func (_this Slice) At(i _dafny.Int) interface{} {
	{
		return ((_this).Dtor_data()).Select((((_this).Dtor_start()).Plus(i)).Uint32()).(interface{})
	}
}
func (_this Slice) Drop(firstIncludedIndex _dafny.Int) Slice {
	{
		return Companion_Slice_.Create_Slice_((_this).Dtor_data(), ((_this).Dtor_start()).Plus(firstIncludedIndex), (_this).Dtor_end())
	}
}
func (_this Slice) Sub(firstIncludedIndex _dafny.Int, lastExcludedIndex _dafny.Int) Slice {
	{
		return Companion_Slice_.Create_Slice_((_this).Dtor_data(), ((_this).Dtor_start()).Plus(firstIncludedIndex), ((_this).Dtor_start()).Plus(lastExcludedIndex))
	}
}

// End of datatype Slice
