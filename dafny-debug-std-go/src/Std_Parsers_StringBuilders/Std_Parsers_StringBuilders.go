// Package Std_Parsers_StringBuilders
// Dafny module Std_Parsers_StringBuilders compiled into Go

package Std_Parsers_StringBuilders

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
	m_Std_Base64 "Std_Base64"
	m_Std_BoundedInts "Std_BoundedInts"
	m_Std_BulkActions "Std_BulkActions"
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
	m_Std_JSON_API "Std_JSON_API"
	m_Std_JSON_ByteStrConversion "Std_JSON_ByteStrConversion"
	m_Std_JSON_ConcreteSyntax_Spec "Std_JSON_ConcreteSyntax_Spec"
	m_Std_JSON_ConcreteSyntax_SpecProperties "Std_JSON_ConcreteSyntax_SpecProperties"
	m_Std_JSON_Deserializer "Std_JSON_Deserializer"
	m_Std_JSON_Deserializer_Uint16StrConversion "Std_JSON_Deserializer_Uint16StrConversion"
	m_Std_JSON_Errors "Std_JSON_Errors"
	m_Std_JSON_Grammar "Std_JSON_Grammar"
	m_Std_JSON_Serializer "Std_JSON_Serializer"
	m_Std_JSON_Spec "Std_JSON_Spec"
	m_Std_JSON_Utils_Cursors "Std_JSON_Utils_Cursors"
	m_Std_JSON_Utils_Lexers_Core "Std_JSON_Utils_Lexers_Core"
	m_Std_JSON_Utils_Lexers_Strings "Std_JSON_Utils_Lexers_Strings"
	m_Std_JSON_Utils_Parsers "Std_JSON_Utils_Parsers"
	m_Std_JSON_Utils_Views_Core "Std_JSON_Utils_Views_Core"
	m_Std_JSON_Utils_Views_Writers "Std_JSON_Utils_Views_Writers"
	m_Std_JSON_Values "Std_JSON_Values"
	m_Std_JSON_ZeroCopy_API "Std_JSON_ZeroCopy_API"
	m_Std_JSON_ZeroCopy_Deserializer "Std_JSON_ZeroCopy_Deserializer"
	m_Std_JSON_ZeroCopy_Deserializer_API "Std_JSON_ZeroCopy_Deserializer_API"
	m_Std_JSON_ZeroCopy_Deserializer_ArrayParams "Std_JSON_ZeroCopy_Deserializer_ArrayParams"
	m_Std_JSON_ZeroCopy_Deserializer_Arrays "Std_JSON_ZeroCopy_Deserializer_Arrays"
	m_Std_JSON_ZeroCopy_Deserializer_Constants "Std_JSON_ZeroCopy_Deserializer_Constants"
	m_Std_JSON_ZeroCopy_Deserializer_Core "Std_JSON_ZeroCopy_Deserializer_Core"
	m_Std_JSON_ZeroCopy_Deserializer_Numbers "Std_JSON_ZeroCopy_Deserializer_Numbers"
	m_Std_JSON_ZeroCopy_Deserializer_ObjectParams "Std_JSON_ZeroCopy_Deserializer_ObjectParams"
	m_Std_JSON_ZeroCopy_Deserializer_Objects "Std_JSON_ZeroCopy_Deserializer_Objects"
	m_Std_JSON_ZeroCopy_Deserializer_Strings "Std_JSON_ZeroCopy_Deserializer_Strings"
	m_Std_JSON_ZeroCopy_Deserializer_Values "Std_JSON_ZeroCopy_Deserializer_Values"
	m_Std_JSON_ZeroCopy_Serializer "Std_JSON_ZeroCopy_Serializer"
	m_Std_Math "Std_Math"
	m_Std_Ordinal "Std_Ordinal"
	m_Std_Parsers_InputString "Std_Parsers_InputString"
	m_Std_Parsers_StringParsers "Std_Parsers_StringParsers"
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
var _ m_Std_BulkActions.Dummy__
var _ m_Std_Base64.Dummy__
var _ m_Std_JSON_Values.Dummy__
var _ m_Std_JSON_Errors.Dummy__
var _ m_Std_JSON_Spec.Dummy__
var _ m_Std_JSON_Utils_Views_Core.Dummy__
var _ m_Std_JSON_Utils_Views_Writers.Dummy__
var _ m_Std_JSON_Utils_Lexers_Core.Dummy__
var _ m_Std_JSON_Utils_Lexers_Strings.Dummy__
var _ m_Std_JSON_Utils_Cursors.Dummy__
var _ m_Std_JSON_Utils_Parsers.Dummy__
var _ m_Std_JSON_Grammar.Dummy__
var _ m_Std_JSON_ByteStrConversion.Dummy__
var _ m_Std_JSON_Serializer.Dummy__
var _ m_Std_JSON_Deserializer_Uint16StrConversion.Dummy__
var _ m_Std_JSON_Deserializer.Dummy__
var _ m_Std_JSON_ConcreteSyntax_Spec.Dummy__
var _ m_Std_JSON_ConcreteSyntax_SpecProperties.Dummy__
var _ m_Std_JSON_ZeroCopy_Serializer.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Core.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Strings.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Numbers.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_ObjectParams.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Objects.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_ArrayParams.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Arrays.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Constants.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_Values.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer_API.Dummy__
var _ m_Std_JSON_ZeroCopy_Deserializer.Dummy__
var _ m_Std_JSON_ZeroCopy_API.Dummy__
var _ m_Std_JSON_API.Dummy__
var _ m_Std_Parsers_InputString.Dummy__
var _ m_Std_Parsers_StringParsers.Dummy__

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
	return "Std_Parsers_StringBuilders.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) ToInput(other _dafny.Sequence) m_Std_Collections_Seq.Slice {
	return m_Std_Collections_Seq.Companion_Slice_.Create_Slice_(other, _dafny.Zero, _dafny.IntOfUint32((other).Cardinality()))
}
func (_static *CompanionStruct_Default___) ToInputEnd(other _dafny.Sequence, fromEnd _dafny.Int) m_Std_Collections_Seq.Slice {
	return m_Std_Collections_Seq.Companion_Slice_.Create_Slice_(other, (_dafny.IntOfUint32((other).Cardinality())).Minus(fromEnd), _dafny.IntOfUint32((other).Cardinality()))
}
func (_static *CompanionStruct_Default___) S(s _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.String_(s)
}
func (_static *CompanionStruct_Default___) String_(s _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return Companion_Default___.S(s)
}
func (_static *CompanionStruct_Default___) Except(s _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer177 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg187 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer177(arg187)
		}
	}(Companion_B_.Rep(Companion_Default___.CharTest((func(_0_s _dafny.Sequence) func(_dafny.CodePoint) bool {
		return func(_1_c _dafny.CodePoint) bool {
			return !_dafny.Companion_Sequence_.Contains(_0_s, _1_c)
		}
	})(s), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("Not '"), s), _dafny.UnicodeSeqOfUtf8Bytes("'")))))
}
func (_static *CompanionStruct_Default___) DebugSummaryInput(name _dafny.Sequence, input _dafny.Sequence) _dafny.Sequence {
	return m_Std_Parsers_StringParsers.Companion_Default___.DebugSummaryInput(name, input)
}
func (_static *CompanionStruct_Default___) PrintDebugSummaryOutput(name _dafny.Sequence, input _dafny.Sequence, result m_Std_Parsers_StringParsers.ParseResult) {
	m_Std_Parsers_StringParsers.Companion_Default___.PrintDebugSummaryOutput(name, input, result)
}
func (_static *CompanionStruct_Default___) FailureToString(input _dafny.Sequence, result m_Std_Parsers_StringParsers.ParseResult) _dafny.Sequence {
	return m_Std_Parsers_StringParsers.Companion_Default___.FailureToString(input, result, _dafny.IntOfInt64(-1))
}
func (_static *CompanionStruct_Default___) Apply(parser func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, input _dafny.Sequence) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.Apply(func(coer178 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg188 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer178(arg188)
		}
	}((parser)), input)
}
func (_static *CompanionStruct_Default___) InputToString(input m_Std_Collections_Seq.Slice) _dafny.Sequence {
	return m_Std_Parsers_InputString.Companion_Default___.View(input)
}
func (_static *CompanionStruct_Default___) SucceedWith(t interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer179 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg189 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer179(arg189)
		}
	}(m_Std_Parsers_StringParsers.Companion_Default___.SucceedWith(t))
}
func (_static *CompanionStruct_Default___) FailWith(message _dafny.Sequence, level m_Std_Parsers_StringParsers.FailureLevel) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer180 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg190 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer180(arg190)
		}
	}(m_Std_Parsers_StringParsers.Companion_Default___.FailWith(message, level))
}
func (_static *CompanionStruct_Default___) ResultWith(result m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer181 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg191 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer181(arg191)
		}
	}(m_Std_Parsers_StringParsers.Companion_Default___.ResultWith(result))
}
func (_static *CompanionStruct_Default___) MId(r interface{}) interface{} {
	return r
}
func (_static *CompanionStruct_Default___) MapIdentity(r interface{}) interface{} {
	return r
}
func (_static *CompanionStruct_Default___) O(alternatives _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	if (_dafny.IntOfUint32((alternatives).Cardinality())).Sign() == 0 {
		return func(coer182 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg192 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer182(arg192)
			}
		}(Companion_Default___.FailWith(_dafny.UnicodeSeqOfUtf8Bytes("no alternative"), m_Std_Parsers_StringParsers.Companion_FailureLevel_.Create_Recoverable_()))
	} else if (_dafny.IntOfUint32((alternatives).Cardinality())).Cmp(_dafny.One) == 0 {
		return (alternatives).Select(0).(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult)
	} else {
		return func(coer183 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg193 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer183(arg193)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Or(func(coer184 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg194 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer184(arg194)
			}
		}(((alternatives).Select(0).(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult))), func(coer185 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg195 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer185(arg195)
			}
		}((Companion_Default___.O((alternatives).Drop(1))))))
	}
}
func (_static *CompanionStruct_Default___) Or(alternatives _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer186 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg196 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer186(arg196)
		}
	}(Companion_Default___.O(alternatives))
}
func (_static *CompanionStruct_Default___) CharTest(test func(_dafny.CodePoint) bool, name _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.CharTest(test, name)
}
func (_static *CompanionStruct_Default___) Rec(underlying func(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer187 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg197 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer187(arg197)
		}
	}(m_Std_Parsers_StringParsers.Companion_Default___.Recursive(func(coer188 func(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg198 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(coer189 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(arg199 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return coer189(arg199)
				}
			}(coer188(func(coer190 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(arg200 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return coer190(arg200)
				}
			}(arg198)))
		}
	}((func(_0_underlying func(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(_1_p func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return ((_0_underlying)(_1_p))
		}
	})(underlying))))
}
func (_static *CompanionStruct_Default___) Recursive(underlying func(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer191 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg201 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer191(arg201)
		}
	}(Companion_Default___.Rec(func(coer192 func(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg202 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(coer193 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(arg203 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return coer193(arg203)
				}
			}(coer192(func(coer194 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(arg204 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return coer194(arg204)
				}
			}(arg202)))
		}
	}(underlying)))
}
func (_static *CompanionStruct_Default___) InputLength(input m_Std_Collections_Seq.Slice) _dafny.Int {
	return m_Std_Parsers_InputString.Companion_Default___.Length(input)
}
func (_static *CompanionStruct_Default___) NonProgressing(input1 m_Std_Collections_Seq.Slice, input2 m_Std_Collections_Seq.Slice) bool {
	return (Companion_Default___.InputLength(input1)).Cmp(Companion_Default___.InputLength(input2)) <= 0
}
func (_static *CompanionStruct_Default___) RecursiveProgressError(name _dafny.Sequence, input1 m_Std_Collections_Seq.Slice, input2 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.RecursiveProgressError(name, input1, input2)
}
func (_static *CompanionStruct_Default___) RecNoStack(underlying func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return (func(_0_underlying func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return Companion_Default___.RecNoStack__(func(coer195 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(arg205 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return coer195(arg205)
				}
			}(_0_underlying), func(coer196 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(arg206 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return coer196(arg206)
				}
			}(_0_underlying), _1_input, _1_input, _dafny.SeqOf())
		}
	})(underlying)
}
func (_static *CompanionStruct_Default___) RecursiveNoStack(underlying func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer197 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg207 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer197(arg207)
		}
	}(Companion_Default___.RecNoStack(func(coer198 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg208 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer198(arg208)
		}
	}(underlying)))
}
func (_static *CompanionStruct_Default___) RecNoStack__(continuation func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, underlying func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, input m_Std_Collections_Seq.Slice, previousInput m_Std_Collections_Seq.Slice, callbacks _dafny.Sequence) m_Std_Parsers_StringParsers.ParseResult {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _0_continuationResult m_Std_Parsers_StringParsers.ParseResult = (continuation)(input)
	_ = _0_continuationResult
	var _1_remaining m_Std_Collections_Seq.Slice = (_0_continuationResult).Remaining()
	_ = _1_remaining
	if ((_0_continuationResult).IsFailure()) || (((*((_0_continuationResult).Extract()).IndexInt(0)).(RecNoStackResult)).Is_RecReturn()) {
		var _2_parseResult m_Std_Parsers_StringParsers.ParseResult = (func() m_Std_Parsers_StringParsers.ParseResult {
			if (_0_continuationResult).IsFailure() {
				return (_0_continuationResult).PropagateFailure()
			}
			return m_Std_Parsers_StringParsers.Companion_ParseResult_.Create_ParseSuccess_(((*((_0_continuationResult).Extract()).IndexInt(0)).(RecNoStackResult)).Dtor_toReturn(), _1_remaining)
		})()
		_ = _2_parseResult
		if (_dafny.IntOfUint32((callbacks).Cardinality())).Sign() == 0 {
			return _2_parseResult
		} else {
			var _3_toCompute func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult = ((callbacks).Select(0).(func(m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult))(_2_parseResult)
			_ = _3_toCompute
			if (m_Std_Parsers_InputString.Companion_Default___.Length(input)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(_1_remaining)) < 0 {
				return Companion_Default___.RecursiveProgressError(_dafny.UnicodeSeqOfUtf8Bytes("Parsers.RecNoStack[internal]"), input, _1_remaining)
			} else if (m_Std_Parsers_InputString.Companion_Default___.Length(previousInput)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(input)) < 0 {
				return Companion_Default___.RecursiveProgressError(_dafny.UnicodeSeqOfUtf8Bytes("Parsers.RecNoStack[internal]"), previousInput, input)
			} else {
				var _in0 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult = _3_toCompute
				_ = _in0
				var _in1 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult = underlying
				_ = _in1
				var _in2 m_Std_Collections_Seq.Slice = _1_remaining
				_ = _in2
				var _in3 m_Std_Collections_Seq.Slice = input
				_ = _in3
				var _in4 _dafny.Sequence = (callbacks).Drop(1)
				_ = _in4
				continuation = _in0
				underlying = _in1
				input = _in2
				previousInput = _in3
				callbacks = _in4
				goto TAIL_CALL_START
			}
		}
	} else {
		var _4_recursor RecNoStackResult = (*((_0_continuationResult).Extract()).IndexInt(0)).(RecNoStackResult)
		_ = _4_recursor
		var _5_rToNewParserOfRecursiveNoStackResultOfR func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult = func(coer199 func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg209 interface{}, arg210 m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(coer200 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return func(arg211 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
						return coer200(arg211)
					}
				}(coer199(arg209, arg210))
			}
		}((_4_recursor).Dtor_toContinue())
		_ = _5_rToNewParserOfRecursiveNoStackResultOfR
		if (m_Std_Parsers_InputString.Companion_Default___.Length(_1_remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(input)) < 0 {
			var _in5 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult = underlying
			_ = _in5
			var _in6 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult = underlying
			_ = _in6
			var _in7 m_Std_Collections_Seq.Slice = _1_remaining
			_ = _in7
			var _in8 m_Std_Collections_Seq.Slice = _1_remaining
			_ = _in8
			var _in9 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((func(_6_rToNewParserOfRecursiveNoStackResultOfR func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(_7_p m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return (func() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
						if (_7_p).IsFailure() {
							return func(coer201 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
								return func(arg212 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
									return coer201(arg212)
								}
							}(m_Std_Parsers_StringParsers.Companion_Default___.ResultWith((_7_p).PropagateFailure()))
						}
						return func(_pat_let48_0 _dafny.Tuple) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
							return func(_8_r interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
								return func(_9_remaining2 m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
									return (_6_rToNewParserOfRecursiveNoStackResultOfR)(_8_r, _9_remaining2)
								}((*(_pat_let48_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
							}((*(_pat_let48_0).IndexInt(0)).(interface{}))
						}((_7_p).Extract())
					})()
				}
			})(_5_rToNewParserOfRecursiveNoStackResultOfR)), callbacks)
			_ = _in9
			continuation = _in5
			underlying = _in6
			input = _in7
			previousInput = _in8
			callbacks = _in9
			goto TAIL_CALL_START
		} else if ((m_Std_Parsers_InputString.Companion_Default___.Length(_1_remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(input)) == 0) && ((m_Std_Parsers_InputString.Companion_Default___.Length(_1_remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(previousInput)) < 0) {
			var _in10 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult = underlying
			_ = _in10
			var _in11 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult = underlying
			_ = _in11
			var _in12 m_Std_Collections_Seq.Slice = _1_remaining
			_ = _in12
			var _in13 m_Std_Collections_Seq.Slice = _1_remaining
			_ = _in13
			var _in14 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((func(_10_rToNewParserOfRecursiveNoStackResultOfR func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(_11_p m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return (func() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
						if (_11_p).IsFailure() {
							return func(coer202 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
								return func(arg213 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
									return coer202(arg213)
								}
							}(m_Std_Parsers_StringParsers.Companion_Default___.ResultWith((_11_p).PropagateFailure()))
						}
						return func(_pat_let49_0 _dafny.Tuple) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
							return func(_12_r interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
								return func(_13_remaining2 m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
									return (_10_rToNewParserOfRecursiveNoStackResultOfR)(_12_r, _13_remaining2)
								}((*(_pat_let49_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
							}((*(_pat_let49_0).IndexInt(0)).(interface{}))
						}((_11_p).Extract())
					})()
				}
			})(_5_rToNewParserOfRecursiveNoStackResultOfR)), callbacks)
			_ = _in14
			continuation = _in10
			underlying = _in11
			input = _in12
			previousInput = _in13
			callbacks = _in14
			goto TAIL_CALL_START
		} else {
			return Companion_Default___.RecursiveProgressError(_dafny.UnicodeSeqOfUtf8Bytes("ParserBuilders.RecNoStack"), input, _1_remaining)
		}
	}
}
func (_static *CompanionStruct_Default___) RecMap(underlying _dafny.Map, startFun _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer203 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg214 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer203(arg214)
		}
	}(m_Std_Parsers_StringParsers.Companion_Default___.RecursiveMap(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter8 := _dafny.Iterate((underlying).Keys().Elements()); ; {
			_compr_0, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _0_k _dafny.Sequence
			_0_k = interface{}(_compr_0).(_dafny.Sequence)
			if (underlying).Contains(_0_k) {
				_coll0.Add(_0_k, m_Std_Parsers_StringParsers.Companion_RecursiveDef_.Create_RecursiveDef_(((underlying).Get(_0_k).(RecMapDef)).Dtor_order(), func(coer204 func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return func(arg215 func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
						return func(coer205 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
							return func(arg216 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
								return coer205(arg216)
							}
						}(coer204(func(coer206 func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
							return func(arg217 _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
								return func(coer207 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
									return func(arg218 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
										return coer207(arg218)
									}
								}(coer206(arg217))
							}
						}(arg215)))
					}
				}((func(_1_underlying _dafny.Map, _2_k _dafny.Sequence) func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return func(_3_selector func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
						return ((func(coer208 func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
							return func(arg219 func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
								return func(coer209 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
									return func(arg220 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
										return coer209(arg220)
									}
								}(coer208(func(coer210 func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
									return func(arg221 _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
										return func(coer211 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
											return func(arg222 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
												return coer211(arg222)
											}
										}(coer210(arg221))
									}
								}(arg219)))
							}
						}(((_1_underlying).Get(_2_k).(RecMapDef)).Dtor_definition()))((func(_4_selector func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
							return func(_5_name _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
								return (_4_selector)(_5_name)
							}
						})(_3_selector)))
					}
				})(underlying, _0_k))))
			}
		}
		return _coll0.ToMap()
	}(), startFun))
}
func (_static *CompanionStruct_Default___) RecursiveMap(underlying _dafny.Map, startFun _dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(coer212 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(arg223 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return coer212(arg223)
		}
	}(Companion_Default___.RecMap(underlying, startFun))
}
func (_static *CompanionStruct_Default___) Int() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.Int()
}
func (_static *CompanionStruct_Default___) Nat() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.Nat()
}
func (_static *CompanionStruct_Default___) Digit() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.Digit()
}
func (_static *CompanionStruct_Default___) DigitNumber() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.DigitNumber()
}
func (_static *CompanionStruct_Default___) WS() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.WS()
}
func (_static *CompanionStruct_Default___) Whitespace() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return Companion_Default___.WS()
}
func (_static *CompanionStruct_Default___) Nothing() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.Epsilon()
}
func (_static *CompanionStruct_Default___) EOS() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return m_Std_Parsers_StringParsers.Companion_Default___.EndOfString()
}
func (_static *CompanionStruct_Default___) EndOfString() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return Companion_Default___.EOS()
}

// End of class Default__

// Definition of datatype B
type B struct {
	Data_B_
}

func (_this B) Get_() Data_B_ {
	return _this.Data_B_
}

type Data_B_ interface {
	isB()
}

type CompanionStruct_B_ struct {
}

var Companion_B_ = CompanionStruct_B_{}

type B_B struct {
	Apply func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult
}

func (B_B) isB() {}

func (CompanionStruct_B_) Create_B_(Apply func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) B {
	return B{B_B{Apply}}
}

func (_this B) Is_B() bool {
	_, ok := _this.Get_().(B_B)
	return ok
}

func (CompanionStruct_B_) Default() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return m_Std_Parsers_StringParsers.Companion_ParseResult_.Default()
	}
}

func (_this B) Dtor_apply() func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return _this.Get_().(B_B).Apply
}

func (_this B) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case B_B:
		{
			return "StringBuilders.B.B" + "(" + _dafny.String(data.Apply) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this B) Equals(other B) bool {
	switch data1 := _this.Get_().(type) {
	case B_B:
		{
			data2, ok := other.Get_().(B_B)
			return ok && _dafny.AreEqual(data1.Apply, data2.Apply)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this B) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(B)
	return ok && _this.Equals(typed)
}

func Type_B_() _dafny.TypeDescriptor {
	return type_B_{}
}

type type_B_ struct {
}

func (_this type_B_) Default() interface{} {
	return Companion_B_.Default()
}

func (_this type_B_) String() string {
	return "Std_Parsers_StringBuilders.B"
}
func (_this B) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = B{}

func (_static CompanionStruct_B_) Apply(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, input _dafny.Sequence) m_Std_Parsers_StringParsers.ParseResult {
	{
		return (_this)(Companion_Default___.ToInput(input))
	}
}
func (_static CompanionStruct_B_) Q_(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer213 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg224 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer213(arg224)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Maybe(func(coer214 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg225 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer214(arg225)
			}
		}((_this))))
	}
}
func (_static CompanionStruct_B_) Option(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer215 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg226 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer215(arg226)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Maybe(func(coer216 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg227 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer216(arg227)
			}
		}((_this))))
	}
}
func (_static CompanionStruct_B_) Q_q_(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer217 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg228 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer217(arg228)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Q_(func(coer218 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg229 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer218(arg229)
			}
		}((_this))))
	}
}
func (_static CompanionStruct_B_) FailureResetsInput(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer219 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg230 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer219(arg230)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Q_(func(coer220 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg231 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer220(arg231)
			}
		}((_this))))
	}
}
func (_static CompanionStruct_B_) E__I(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, other func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer221 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg232 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer221(arg232)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.ConcatKeepRight(func(coer222 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg233 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer222(arg233)
			}
		}((_this)), func(coer223 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg234 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer223(arg234)
			}
		}((other))))
	}
}
func (_static CompanionStruct_B_) ConcatKeepRight(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, other func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer224 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg235 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer224(arg235)
			}
		}(Companion_B_.E__I(_this, func(coer225 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg236 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer225(arg236)
			}
		}(other)))
	}
}
func (_static CompanionStruct_B_) I__e(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, other func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer226 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg237 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer226(arg237)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.ConcatKeepLeft(func(coer227 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg238 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer227(arg238)
			}
		}((_this)), func(coer228 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg239 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer228(arg239)
			}
		}((other))))
	}
}
func (_static CompanionStruct_B_) ConcatKeepLeft(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, other func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return Companion_B_.I__e(_this, func(coer229 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg240 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer229(arg240)
			}
		}(other))
	}
}
func (_static CompanionStruct_B_) I__I(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, other func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer230 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg241 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer230(arg241)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Concat(func(coer231 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg242 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer231(arg242)
			}
		}((_this)), func(coer232 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg243 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer232(arg243)
			}
		}((other))))
	}
}
func (_static CompanionStruct_B_) Concat(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, other func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer233 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg244 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer233(arg244)
			}
		}(Companion_B_.I__I(_this, func(coer234 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg245 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer234(arg245)
			}
		}(other)))
	}
}
func (_static CompanionStruct_B_) If(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, cond func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer235 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg246 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer235(arg246)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.If(func(coer236 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg247 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer236(arg247)
			}
		}((cond)), func(coer237 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg248 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer237(arg248)
			}
		}((_this))))
	}
}
func (_static CompanionStruct_B_) M(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, mappingFunc func(interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer238 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg249 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer238(arg249)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Map(func(coer239 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg250 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer239(arg250)
			}
		}((_this)), func(coer240 func(interface{}) interface{}) func(interface{}) interface{} {
			return func(arg251 interface{}) interface{} {
				return coer240(arg251)
			}
		}(mappingFunc)))
	}
}
func (_static CompanionStruct_B_) Map(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, mappingFunc func(interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer241 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg252 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer241(arg252)
			}
		}(Companion_B_.M(_this, func(coer242 func(interface{}) interface{}) func(interface{}) interface{} {
			return func(arg253 interface{}) interface{} {
				return coer242(arg253)
			}
		}(mappingFunc)))
	}
}
func (_static CompanionStruct_B_) M2(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, unfolder func(interface{}) _dafny.Tuple, mappingFunc func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer243 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg254 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer243(arg254)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Map(func(coer244 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg255 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer244(arg255)
			}
		}((_this)), func(coer245 func(interface{}) interface{}) func(interface{}) interface{} {
			return func(arg256 interface{}) interface{} {
				return coer245(arg256)
			}
		}((func(_0_unfolder func(interface{}) _dafny.Tuple, _1_mappingFunc func(interface{}, interface{}) interface{}) func(interface{}) interface{} {
			return func(_2_x interface{}) interface{} {
				return func(_pat_let50_0 _dafny.Tuple) interface{} {
					return func(_3_x _dafny.Tuple) interface{} {
						return (_1_mappingFunc)((*(_3_x).IndexInt(0)), (*(_3_x).IndexInt(1)))
					}(_pat_let50_0)
				}((_0_unfolder)(_2_x))
			}
		})(unfolder, mappingFunc))))
	}
}
func (_static CompanionStruct_B_) Map2(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, unfolder func(interface{}) _dafny.Tuple, mappingFunc func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer246 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg257 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer246(arg257)
			}
		}(Companion_B_.M2(_this, func(coer247 func(interface{}) _dafny.Tuple) func(interface{}) _dafny.Tuple {
			return func(arg258 interface{}) _dafny.Tuple {
				return coer247(arg258)
			}
		}(unfolder), func(coer248 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
			return func(arg259 interface{}, arg260 interface{}) interface{} {
				return coer248(arg259, arg260)
			}
		}(mappingFunc)))
	}
}
func (_static CompanionStruct_B_) M3(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, unfolder func(interface{}) _dafny.Tuple, mappingFunc func(interface{}, interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer249 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg261 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer249(arg261)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Map(func(coer250 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg262 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer250(arg262)
			}
		}((_this)), func(coer251 func(interface{}) interface{}) func(interface{}) interface{} {
			return func(arg263 interface{}) interface{} {
				return coer251(arg263)
			}
		}((func(_0_unfolder func(interface{}) _dafny.Tuple, _1_mappingFunc func(interface{}, interface{}, interface{}) interface{}) func(interface{}) interface{} {
			return func(_2_x interface{}) interface{} {
				return func(_pat_let51_0 _dafny.Tuple) interface{} {
					return func(_3_x _dafny.Tuple) interface{} {
						return (_1_mappingFunc)((*(_3_x).IndexInt(0)), (*(_3_x).IndexInt(1)), (*(_3_x).IndexInt(2)))
					}(_pat_let51_0)
				}((_0_unfolder)(_2_x))
			}
		})(unfolder, mappingFunc))))
	}
}
func (_static CompanionStruct_B_) Map3(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, unfolder func(interface{}) _dafny.Tuple, mappingFunc func(interface{}, interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer252 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg264 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer252(arg264)
			}
		}(Companion_B_.M3(_this, func(coer253 func(interface{}) _dafny.Tuple) func(interface{}) _dafny.Tuple {
			return func(arg265 interface{}) _dafny.Tuple {
				return coer253(arg265)
			}
		}(unfolder), func(coer254 func(interface{}, interface{}, interface{}) interface{}) func(interface{}, interface{}, interface{}) interface{} {
			return func(arg266 interface{}, arg267 interface{}, arg268 interface{}) interface{} {
				return coer254(arg266, arg267, arg268)
			}
		}(mappingFunc)))
	}
}
func (_static CompanionStruct_B_) Then(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, other func(interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer255 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg269 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer255(arg269)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Bind(func(coer256 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg270 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer256(arg270)
			}
		}((_this)), func(coer257 func(interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg271 interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(coer258 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return func(arg272 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
						return coer258(arg272)
					}
				}(coer257(arg271))
			}
		}((func(_0_other func(interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(_1_result interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return ((_0_other)(_1_result))
			}
		})(other))))
	}
}
func (_static CompanionStruct_B_) ThenWithRemaining(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, other func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer259 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg273 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer259(arg273)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.BindSucceeds(func(coer260 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg274 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer260(arg274)
			}
		}((_this)), func(coer261 func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg275 interface{}, arg276 m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(coer262 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return func(arg277 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
						return coer262(arg277)
					}
				}(coer261(arg275, arg276))
			}
		}((func(_0_other func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(_1_result interface{}, _2_input m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return ((_0_other)(_1_result, _2_input))
			}
		})(other))))
	}
}
func (_static CompanionStruct_B_) Bind(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, other func(m_Std_Parsers_StringParsers.ParseResult, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer263 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg278 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer263(arg278)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.BindResult(func(coer264 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg279 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer264(arg279)
			}
		}((_this)), func(coer265 func(m_Std_Parsers_StringParsers.ParseResult, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Parsers_StringParsers.ParseResult, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg280 m_Std_Parsers_StringParsers.ParseResult, arg281 m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return func(coer266 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
					return func(arg282 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
						return coer266(arg282)
					}
				}(coer265(arg280, arg281))
			}
		}((func(_0_other func(m_Std_Parsers_StringParsers.ParseResult, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Parsers_StringParsers.ParseResult, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(_1_result m_Std_Parsers_StringParsers.ParseResult, _2_input m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return ((_0_other)(_1_result, _2_input))
			}
		})(other))))
	}
}
func (_static CompanionStruct_B_) Debug(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, name _dafny.Sequence, onEnter func(_dafny.Sequence, m_Std_Collections_Seq.Slice) interface{}, onExit func(_dafny.Sequence, interface{}, m_Std_Parsers_StringParsers.ParseResult) _dafny.Tuple) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer267 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg283 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer267(arg283)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Debug(func(coer268 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg284 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer268(arg284)
			}
		}((_this)), name, func(coer269 func(_dafny.Sequence, m_Std_Collections_Seq.Slice) interface{}) func(_dafny.Sequence, m_Std_Collections_Seq.Slice) interface{} {
			return func(arg285 _dafny.Sequence, arg286 m_Std_Collections_Seq.Slice) interface{} {
				return coer269(arg285, arg286)
			}
		}(onEnter), func(coer270 func(_dafny.Sequence, interface{}, m_Std_Parsers_StringParsers.ParseResult) _dafny.Tuple) func(_dafny.Sequence, interface{}, m_Std_Parsers_StringParsers.ParseResult) _dafny.Tuple {
			return func(arg287 _dafny.Sequence, arg288 interface{}, arg289 m_Std_Parsers_StringParsers.ParseResult) _dafny.Tuple {
				return coer270(arg287, arg288, arg289)
			}
		}(onExit)))
	}
}
func (_static CompanionStruct_B_) RepFold(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, init interface{}, combine func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer271 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg290 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer271(arg290)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.Rep(func(coer272 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg291 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer272(arg291)
			}
		}((_this)), func(coer273 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
			return func(arg292 interface{}, arg293 interface{}) interface{} {
				return coer273(arg292, arg293)
			}
		}(combine), init))
	}
}
func (_static CompanionStruct_B_) RepeatFold(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, init interface{}, combine func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer274 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg294 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer274(arg294)
			}
		}(Companion_B_.RepFold(_this, init, func(coer275 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
			return func(arg295 interface{}, arg296 interface{}) interface{} {
				return coer275(arg295, arg296)
			}
		}(combine)))
	}
}
func (_static CompanionStruct_B_) RepSep(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, separator func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer276 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg297 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer276(arg297)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.RepSep(func(coer277 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg298 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer277(arg298)
			}
		}((_this)), func(coer278 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg299 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer278(arg299)
			}
		}((separator))))
	}
}
func (_static CompanionStruct_B_) RepeatSeparator(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, separator func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return Companion_B_.RepSep(_this, func(coer279 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg300 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer279(arg300)
			}
		}(separator))
	}
}
func (_static CompanionStruct_B_) RepMerge(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, merger func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer280 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg301 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer280(arg301)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.RepMerge(func(coer281 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg302 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer281(arg302)
			}
		}((_this)), func(coer282 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
			return func(arg303 interface{}, arg304 interface{}) interface{} {
				return coer282(arg303, arg304)
			}
		}(merger)))
	}
}
func (_static CompanionStruct_B_) RepeatMerge(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, merger func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return Companion_B_.RepMerge(_this, merger)
	}
}
func (_static CompanionStruct_B_) RepSepMerge(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, separator func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, merger func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer283 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg305 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer283(arg305)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.RepSepMerge(func(coer284 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg306 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer284(arg306)
			}
		}((_this)), func(coer285 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg307 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer285(arg307)
			}
		}((separator)), func(coer286 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
			return func(arg308 interface{}, arg309 interface{}) interface{} {
				return coer286(arg308, arg309)
			}
		}(merger)))
	}
}
func (_static CompanionStruct_B_) RepeatSeparatorMerge(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, separator func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult, merger func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return Companion_B_.RepSepMerge(_this, func(coer287 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg310 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer287(arg310)
			}
		}(separator), merger)
	}
}
func (_static CompanionStruct_B_) Rep(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer288 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg311 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer288(arg311)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.ZeroOrMore(func(coer289 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg312 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer289(arg312)
			}
		}((_this))))
	}
}
func (_static CompanionStruct_B_) Repeat(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return Companion_B_.Rep(_this)
	}
}
func (_static CompanionStruct_B_) Rep1(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return func(coer290 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg313 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer290(arg313)
			}
		}(m_Std_Parsers_StringParsers.Companion_Default___.OneOrMore(func(coer291 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg314 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer291(arg314)
			}
		}((_this))))
	}
}
func (_static CompanionStruct_B_) RepeatAtLeastOnce(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return Companion_B_.Rep1(_this)
	}
}
func (_static CompanionStruct_B_) End(_this func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	{
		return Companion_B_.I__e(_this, func(coer292 func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return func(arg315 m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
				return coer292(arg315)
			}
		}(Companion_Default___.EOS()))
	}
}

// End of datatype B

// Definition of datatype RecNoStackResult
type RecNoStackResult struct {
	Data_RecNoStackResult_
}

func (_this RecNoStackResult) Get_() Data_RecNoStackResult_ {
	return _this.Data_RecNoStackResult_
}

type Data_RecNoStackResult_ interface {
	isRecNoStackResult()
}

type CompanionStruct_RecNoStackResult_ struct {
}

var Companion_RecNoStackResult_ = CompanionStruct_RecNoStackResult_{}

type RecNoStackResult_RecReturn struct {
	ToReturn interface{}
}

func (RecNoStackResult_RecReturn) isRecNoStackResult() {}

func (CompanionStruct_RecNoStackResult_) Create_RecReturn_(ToReturn interface{}) RecNoStackResult {
	return RecNoStackResult{RecNoStackResult_RecReturn{ToReturn}}
}

func (_this RecNoStackResult) Is_RecReturn() bool {
	_, ok := _this.Get_().(RecNoStackResult_RecReturn)
	return ok
}

type RecNoStackResult_RecContinue struct {
	ToContinue func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult
}

func (RecNoStackResult_RecContinue) isRecNoStackResult() {}

func (CompanionStruct_RecNoStackResult_) Create_RecContinue_(ToContinue func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) RecNoStackResult {
	return RecNoStackResult{RecNoStackResult_RecContinue{ToContinue}}
}

func (_this RecNoStackResult) Is_RecContinue() bool {
	_, ok := _this.Get_().(RecNoStackResult_RecContinue)
	return ok
}

func (CompanionStruct_RecNoStackResult_) Default() RecNoStackResult {
	return Companion_RecNoStackResult_.Create_RecContinue_(func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return m_Std_Parsers_StringParsers.Companion_ParseResult_.Default()
		}
	})
}

func (_this RecNoStackResult) Dtor_toReturn() interface{} {
	return _this.Get_().(RecNoStackResult_RecReturn).ToReturn
}

func (_this RecNoStackResult) Dtor_toContinue() func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return _this.Get_().(RecNoStackResult_RecContinue).ToContinue
}

func (_this RecNoStackResult) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RecNoStackResult_RecReturn:
		{
			return "StringBuilders.RecNoStackResult.RecReturn" + "(" + _dafny.String(data.ToReturn) + ")"
		}
	case RecNoStackResult_RecContinue:
		{
			return "StringBuilders.RecNoStackResult.RecContinue" + "(" + _dafny.String(data.ToContinue) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RecNoStackResult) Equals(other RecNoStackResult) bool {
	switch data1 := _this.Get_().(type) {
	case RecNoStackResult_RecReturn:
		{
			data2, ok := other.Get_().(RecNoStackResult_RecReturn)
			return ok && _dafny.AreEqual(data1.ToReturn, data2.ToReturn)
		}
	case RecNoStackResult_RecContinue:
		{
			data2, ok := other.Get_().(RecNoStackResult_RecContinue)
			return ok && _dafny.AreEqual(data1.ToContinue, data2.ToContinue)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RecNoStackResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RecNoStackResult)
	return ok && _this.Equals(typed)
}

func Type_RecNoStackResult_() _dafny.TypeDescriptor {
	return type_RecNoStackResult_{}
}

type type_RecNoStackResult_ struct {
}

func (_this type_RecNoStackResult_) Default() interface{} {
	return Companion_RecNoStackResult_.Default()
}

func (_this type_RecNoStackResult_) String() string {
	return "Std_Parsers_StringBuilders.RecNoStackResult"
}
func (_this RecNoStackResult) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RecNoStackResult{}

// End of datatype RecNoStackResult

// Definition of datatype RecMapDef
type RecMapDef struct {
	Data_RecMapDef_
}

func (_this RecMapDef) Get_() Data_RecMapDef_ {
	return _this.Data_RecMapDef_
}

type Data_RecMapDef_ interface {
	isRecMapDef()
}

type CompanionStruct_RecMapDef_ struct {
}

var Companion_RecMapDef_ = CompanionStruct_RecMapDef_{}

type RecMapDef_RecMapDef struct {
	Order      _dafny.Int
	Definition func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult
}

func (RecMapDef_RecMapDef) isRecMapDef() {}

func (CompanionStruct_RecMapDef_) Create_RecMapDef_(Order _dafny.Int, Definition func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) RecMapDef {
	return RecMapDef{RecMapDef_RecMapDef{Order, Definition}}
}

func (_this RecMapDef) Is_RecMapDef() bool {
	_, ok := _this.Get_().(RecMapDef_RecMapDef)
	return ok
}

func (CompanionStruct_RecMapDef_) Default() RecMapDef {
	return Companion_RecMapDef_.Create_RecMapDef_(_dafny.Zero, func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
		return func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
			return m_Std_Parsers_StringParsers.Companion_ParseResult_.Default()
		}
	})
}

func (_this RecMapDef) Dtor_order() _dafny.Int {
	return _this.Get_().(RecMapDef_RecMapDef).Order
}

func (_this RecMapDef) Dtor_definition() func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult) func(m_Std_Collections_Seq.Slice) m_Std_Parsers_StringParsers.ParseResult {
	return _this.Get_().(RecMapDef_RecMapDef).Definition
}

func (_this RecMapDef) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RecMapDef_RecMapDef:
		{
			return "StringBuilders.RecMapDef.RecMapDef" + "(" + _dafny.String(data.Order) + ", " + _dafny.String(data.Definition) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RecMapDef) Equals(other RecMapDef) bool {
	switch data1 := _this.Get_().(type) {
	case RecMapDef_RecMapDef:
		{
			data2, ok := other.Get_().(RecMapDef_RecMapDef)
			return ok && data1.Order.Cmp(data2.Order) == 0 && _dafny.AreEqual(data1.Definition, data2.Definition)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RecMapDef) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RecMapDef)
	return ok && _this.Equals(typed)
}

func Type_RecMapDef_() _dafny.TypeDescriptor {
	return type_RecMapDef_{}
}

type type_RecMapDef_ struct {
}

func (_this type_RecMapDef_) Default() interface{} {
	return Companion_RecMapDef_.Default()
}

func (_this type_RecMapDef_) String() string {
	return "Std_Parsers_StringBuilders.RecMapDef"
}
func (_this RecMapDef) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RecMapDef{}

// End of datatype RecMapDef
