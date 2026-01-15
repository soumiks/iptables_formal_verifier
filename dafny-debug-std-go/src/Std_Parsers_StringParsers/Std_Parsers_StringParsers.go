// Package Std_Parsers_StringParsers
// Dafny module Std_Parsers_StringParsers compiled into Go

package Std_Parsers_StringParsers

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
	return "Std_Parsers_StringParsers.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Char(expectedChar _dafny.CodePoint) func(m_Std_Collections_Seq.Slice) ParseResult {
	return Companion_Default___.CharTest((func(_0_expectedChar _dafny.CodePoint) func(_dafny.CodePoint) bool {
		return func(_1_c _dafny.CodePoint) bool {
			return (_1_c) == (_0_expectedChar)
		}
	})(expectedChar), _dafny.SeqOf(expectedChar))
}
func (_static *CompanionStruct_Default___) Space() func(m_Std_Collections_Seq.Slice) ParseResult {
	return Companion_Default___.CharTest(func(_0_c _dafny.CodePoint) bool {
		return _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes(" \t\r\n               　"), _0_c)
	}, _dafny.UnicodeSeqOfUtf8Bytes("space"))
}
func (_static *CompanionStruct_Default___) WS() func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer62 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg63 m_Std_Collections_Seq.Slice) ParseResult {
			return coer62(arg63)
		}
	}(Companion_Default___.ZeroOrMore(func(coer63 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg64 m_Std_Collections_Seq.Slice) ParseResult {
			return coer63(arg64)
		}
	}(Companion_Default___.Space())))
}
func (_static *CompanionStruct_Default___) Digit() func(m_Std_Collections_Seq.Slice) ParseResult {
	return Companion_Default___.CharTest(func(_0_c _dafny.CodePoint) bool {
		return _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("0123456789"), _0_c)
	}, _dafny.UnicodeSeqOfUtf8Bytes("digit"))
}
func (_static *CompanionStruct_Default___) DigitNumber() func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer64 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg65 m_Std_Collections_Seq.Slice) ParseResult {
			return coer64(arg65)
		}
	}(Companion_Default___.Map(func(coer65 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg66 m_Std_Collections_Seq.Slice) ParseResult {
			return coer65(arg66)
		}
	}(Companion_Default___.Digit()), func(coer66 func(_dafny.CodePoint) _dafny.Int) func(interface{}) interface{} {
		return func(arg67 interface{}) interface{} {
			return coer66(arg67.(_dafny.CodePoint))
		}
	}(func(_0_c _dafny.CodePoint) _dafny.Int {
		return func(_pat_let6_0 _dafny.Int) _dafny.Int {
			return func(_1_d _dafny.Int) _dafny.Int {
				return func(_pat_let7_0 _dafny.Int) _dafny.Int {
					return func(_2_n _dafny.Int) _dafny.Int {
						return _2_n
					}(_pat_let7_0)
				}((func() _dafny.Int {
					if (_1_d).Sign() != -1 {
						return _1_d
					}
					return _dafny.Zero
				})())
			}(_pat_let6_0)
		}(Companion_Default___.DigitToInt(_0_c))
	})))
}
func (_static *CompanionStruct_Default___) Nat() func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer67 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg68 m_Std_Collections_Seq.Slice) ParseResult {
			return coer67(arg68)
		}
	}(Companion_Default___.Bind(func(coer68 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg69 m_Std_Collections_Seq.Slice) ParseResult {
			return coer68(arg69)
		}
	}(Companion_Default___.DigitNumber()), func(coer69 func(_dafny.Int) func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg70 interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(coer70 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg71 m_Std_Collections_Seq.Slice) ParseResult {
					return coer70(arg71)
				}
			}(coer69(arg70.(_dafny.Int)))
		}
	}(func(_0_result _dafny.Int) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(coer71 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(arg72 m_Std_Collections_Seq.Slice) ParseResult {
				return coer71(arg72)
			}
		}(Companion_Default___.Rep(func(coer72 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(arg73 m_Std_Collections_Seq.Slice) ParseResult {
				return coer72(arg73)
			}
		}(Companion_Default___.DigitNumber()), func(coer73 func(_dafny.Int, _dafny.Int) _dafny.Int) func(interface{}, interface{}) interface{} {
			return func(arg74 interface{}, arg75 interface{}) interface{} {
				return coer73(arg74.(_dafny.Int), arg75.(_dafny.Int))
			}
		}(func(_1_previous _dafny.Int, _2_c _dafny.Int) _dafny.Int {
			return func(_pat_let8_0 _dafny.Int) _dafny.Int {
				return func(_3_r _dafny.Int) _dafny.Int {
					return _3_r
				}(_pat_let8_0)
			}(((_1_previous).Times(_dafny.IntOfInt64(10))).Plus(_2_c))
		}), _0_result))
	})))
}
func (_static *CompanionStruct_Default___) Int() func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer74 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg76 m_Std_Collections_Seq.Slice) ParseResult {
			return coer74(arg76)
		}
	}(Companion_Default___.Bind(func(coer75 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg77 m_Std_Collections_Seq.Slice) ParseResult {
			return coer75(arg77)
		}
	}(func(coer76 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg78 m_Std_Collections_Seq.Slice) ParseResult {
			return coer76(arg78)
		}
	}(Companion_Default___.Maybe(func(coer77 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg79 m_Std_Collections_Seq.Slice) ParseResult {
			return coer77(arg79)
		}
	}(Companion_Default___.Char(_dafny.CodePoint('-')))))), func(coer78 func(m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg80 interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(coer79 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg81 m_Std_Collections_Seq.Slice) ParseResult {
					return coer79(arg81)
				}
			}(coer78(arg80.(m_Std_Wrappers.Option)))
		}
	}(func(_0_minusSign m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(coer80 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(arg82 m_Std_Collections_Seq.Slice) ParseResult {
				return coer80(arg82)
			}
		}(Companion_Default___.Map(func(coer81 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(arg83 m_Std_Collections_Seq.Slice) ParseResult {
				return coer81(arg83)
			}
		}(Companion_Default___.Nat()), func(coer82 func(_dafny.Int) _dafny.Int) func(interface{}) interface{} {
			return func(arg84 interface{}) interface{} {
				return coer82(arg84.(_dafny.Int))
			}
		}((func(_1_minusSign m_Std_Wrappers.Option) func(_dafny.Int) _dafny.Int {
			return func(_2_result _dafny.Int) _dafny.Int {
				return (func() _dafny.Int {
					if (_1_minusSign).Is_Some() {
						return (_dafny.Zero).Minus(_2_result)
					}
					return _2_result
				})()
			}
		})(_0_minusSign))))
	})))
}
func (_static *CompanionStruct_Default___) String_(expected _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_expected _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) ParseResult {
			return (func() ParseResult {
				if ((_dafny.IntOfUint32((_0_expected).Cardinality())).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(_1_input)) <= 0) && (_dafny.Companion_Sequence_.Equal((m_Std_Parsers_InputString.Companion_Default___.Slice(_1_input, _dafny.Zero, _dafny.IntOfUint32((_0_expected).Cardinality()))).View(), _0_expected)) {
					return Companion_ParseResult_.Create_ParseSuccess_(_0_expected, m_Std_Parsers_InputString.Companion_Default___.Drop(_1_input, _dafny.IntOfUint32((_0_expected).Cardinality())))
				}
				return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Recoverable_(), Companion_FailureData_.Create_FailureData_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("expected '"), _0_expected), _dafny.UnicodeSeqOfUtf8Bytes("'")), _1_input, m_Std_Wrappers.Companion_Option_.Create_None_()))
			})()
		}
	})(expected)
}
func (_static *CompanionStruct_Default___) ExtractLineCol(input _dafny.Sequence, pos _dafny.Int) CodeLocation {
	var output CodeLocation = Companion_CodeLocation_.Default()
	_ = output
	var _0_lineNumber _dafny.Int = _dafny.Zero
	_ = _0_lineNumber
	var _1_colNumber _dafny.Int = _dafny.Zero
	_ = _1_colNumber
	var _2_lineStr _dafny.Sequence = _dafny.EmptySeq
	_ = _2_lineStr
	_0_lineNumber = _dafny.One
	var _3_startLinePos _dafny.Int
	_ = _3_startLinePos
	_3_startLinePos = _dafny.Zero
	_1_colNumber = _dafny.Zero
	var _4_i _dafny.Int
	_ = _4_i
	_4_i = _dafny.Zero
	for ((_4_i).Cmp(_dafny.IntOfUint32((input).Cardinality())) < 0) && ((_4_i).Cmp(pos) != 0) {
		_1_colNumber = (_1_colNumber).Plus(_dafny.One)
		if ((((input).Select((_4_i).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('\r'))) && (((_4_i).Plus(_dafny.One)).Cmp(_dafny.IntOfUint32((input).Cardinality())) < 0)) && (((input).Select(((_4_i).Plus(_dafny.One)).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('\n'))) {
			_0_lineNumber = (_0_lineNumber).Plus(_dafny.One)
			_1_colNumber = _dafny.Zero
			_4_i = (_4_i).Plus(_dafny.One)
			_3_startLinePos = (_4_i).Plus(_dafny.One)
		} else if _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("\r\n"), (input).Select((_4_i).Uint32()).(_dafny.CodePoint)) {
			_0_lineNumber = (_0_lineNumber).Plus(_dafny.One)
			_1_colNumber = _dafny.Zero
			_3_startLinePos = (_4_i).Plus(_dafny.One)
		}
		_4_i = (_4_i).Plus(_dafny.One)
	}
	for ((_4_i).Cmp(_dafny.IntOfUint32((input).Cardinality())) < 0) && (!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("\r\n"), (input).Select((_4_i).Uint32()).(_dafny.CodePoint))) {
		_4_i = (_4_i).Plus(_dafny.One)
	}
	_2_lineStr = (input).Subsequence((_3_startLinePos).Uint32(), (_4_i).Uint32())
	output = Companion_CodeLocation_.Create_CodeLocation_(_0_lineNumber, _1_colNumber, _2_lineStr)
	return output
}
func (_static *CompanionStruct_Default___) DebugSummary(input _dafny.Sequence) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if (_dafny.IntOfUint32((input).Cardinality())).Sign() == 1 {
			return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("'"), func() _dafny.Sequence {
				var _source0 _dafny.CodePoint = (input).Select(0).(_dafny.CodePoint)
				_ = _source0
				{
					if (_source0) == (_dafny.CodePoint('\n')) {
						return _dafny.UnicodeSeqOfUtf8Bytes("\\n")
					}
				}
				{
					if (_source0) == (_dafny.CodePoint('\r')) {
						return _dafny.UnicodeSeqOfUtf8Bytes("\\r")
					}
				}
				{
					if (_source0) == (_dafny.CodePoint('\t')) {
						return _dafny.UnicodeSeqOfUtf8Bytes("\\t")
					}
				}
				{
					var _0_c _dafny.CodePoint = _source0
					_ = _0_c
					return _dafny.SeqOf(_0_c)
				}
			}()), (func() _dafny.Sequence {
				if (_dafny.IntOfUint32((input).Cardinality())).Cmp(_dafny.One) == 0 {
					return _dafny.UnicodeSeqOfUtf8Bytes("' and end of string")
				}
				return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("'"), _dafny.UnicodeSeqOfUtf8Bytes(" and ")), m_Std_Strings.Companion_Default___.OfInt((_dafny.IntOfUint32((input).Cardinality())).Minus(_dafny.One))), _dafny.UnicodeSeqOfUtf8Bytes(" char")), (func() _dafny.Sequence {
					if (_dafny.IntOfUint32((input).Cardinality())).Cmp(_dafny.IntOfInt64(2)) == 0 {
						return _dafny.UnicodeSeqOfUtf8Bytes("")
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("s")
				})()), _dafny.UnicodeSeqOfUtf8Bytes(" remaining"))
			})())
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("'' (end of string)")
	})(), _dafny.UnicodeSeqOfUtf8Bytes("\n"))
}
func (_static *CompanionStruct_Default___) DebugNameSummary(name _dafny.Sequence, input _dafny.Sequence) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("["), name), _dafny.UnicodeSeqOfUtf8Bytes("] ")), Companion_Default___.DebugSummary(input))
}
func (_static *CompanionStruct_Default___) DebugSummaryInput(name _dafny.Sequence, input _dafny.Sequence) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("> "), Companion_Default___.DebugNameSummary(name, input))
}
func (_static *CompanionStruct_Default___) PrintDebugSummaryInput(name _dafny.Sequence, input _dafny.Sequence) {
	_dafny.Print(Companion_Default___.DebugSummaryInput(name, input).VerbatimString(false))
}
func (_static *CompanionStruct_Default___) NewIndent(input _dafny.Sequence, indent _dafny.Sequence) _dafny.Sequence {
	var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
	_ = _0___accumulator
	goto TAIL_CALL_START
TAIL_CALL_START:
	if (_dafny.IntOfUint32((input).Cardinality())).Sign() == 0 {
		return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.UnicodeSeqOfUtf8Bytes(""))
	} else {
		_0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (func() _dafny.Sequence {
			if ((input).Select(0).(_dafny.CodePoint)) == (_dafny.CodePoint('\n')) {
				return _dafny.Companion_Sequence_.Concatenate((input).Take(1), indent)
			}
			return (input).Take(1)
		})())
		var _in0 _dafny.Sequence = (input).Drop(1)
		_ = _in0
		var _in1 _dafny.Sequence = indent
		_ = _in1
		input = _in0
		indent = _in1
		goto TAIL_CALL_START
	}
}
func (_static *CompanionStruct_Default___) PrintDebugSummaryOutput(name _dafny.Sequence, input _dafny.Sequence, result ParseResult) {
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("< ").VerbatimString(false))
	_dafny.Print(Companion_Default___.DebugNameSummary(name, input).VerbatimString(false))
	if (result).Is_ParseFailure() {
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("| Unparsed: ").VerbatimString(false))
		_dafny.Print(Companion_Default___.DebugSummary(m_Std_Parsers_InputString.Companion_Default___.View((result).Remaining())).VerbatimString(false))
		if (m_Std_Parsers_InputString.Companion_Default___.Length((result).Remaining())).Cmp(_dafny.IntOfUint32((input).Cardinality())) < 0 {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("| Was committed\n").VerbatimString(false))
		}
		_dafny.Print(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("| "), Companion_Default___.NewIndent(Companion_Default___.FailureToString(input, result, _dafny.IntOfInt64(-1)), _dafny.UnicodeSeqOfUtf8Bytes("| "))).VerbatimString(false))
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	} else {
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("| Success: ").VerbatimString(false))
		_dafny.Print((result).Dtor_result())
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(", ").VerbatimString(false))
		_dafny.Print(Companion_Default___.DebugSummary(m_Std_Parsers_InputString.Companion_Default___.View((result).Remaining())).VerbatimString(false))
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	}
}
func (_static *CompanionStruct_Default___) FailureToString(input _dafny.Sequence, result ParseResult, printPos _dafny.Int) _dafny.Sequence {
	var _0_failure _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("")
	_ = _0_failure
	var _1_failure _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_0_failure, (func() _dafny.Sequence {
		if (printPos).Cmp(_dafny.IntOfInt64(-1)) == 0 {
			return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if ((result).Dtor_level()).Equals(Companion_FailureLevel_.Create_Fatal_()) {
					return _dafny.UnicodeSeqOfUtf8Bytes("Fatal error")
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("Error")
			})(), _dafny.UnicodeSeqOfUtf8Bytes(":\n"))
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("")
	})())
	_ = _1_failure
	var _2_pos _dafny.Int = (_dafny.IntOfUint32((input).Cardinality())).Minus(m_Std_Parsers_InputString.Companion_Default___.Length(((result).Dtor_data()).Dtor_remaining()))
	_ = _2_pos
	var _3_pos _dafny.Int = (func() _dafny.Int {
		if (_2_pos).Sign() == -1 {
			return _dafny.Zero
		}
		return _2_pos
	})()
	_ = _3_pos
	var _4_failure _dafny.Sequence = (func() _dafny.Sequence {
		if (printPos).Cmp(_3_pos) == 0 {
			return _1_failure
		}
		return func(_pat_let9_0 CodeLocation) _dafny.Sequence {
			return func(_5_output CodeLocation) _dafny.Sequence {
				return func(_pat_let10_0 CodeLocation) _dafny.Sequence {
					return func(_6_line _dafny.Int) _dafny.Sequence {
						return func(_7_col _dafny.Int) _dafny.Sequence {
							return func(_8_lineStr _dafny.Sequence) _dafny.Sequence {
								return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1_failure, m_Std_Strings.Companion_Default___.OfInt(_6_line)), _dafny.UnicodeSeqOfUtf8Bytes(": ")), _8_lineStr), _dafny.UnicodeSeqOfUtf8Bytes("\n")), m_Std_Collections_Seq.Companion_Default___.Repeat(_dafny.CodePoint(' '), ((_7_col).Plus(_dafny.IntOfInt64(2))).Plus(_dafny.IntOfUint32((m_Std_Strings.Companion_Default___.OfInt(_6_line)).Cardinality())))), _dafny.UnicodeSeqOfUtf8Bytes("^")), _dafny.UnicodeSeqOfUtf8Bytes("\n"))
							}(_pat_let10_0.Get_().(CodeLocation_CodeLocation).LineStr)
						}(_pat_let10_0.Get_().(CodeLocation_CodeLocation).ColNumber)
					}(_pat_let10_0.Get_().(CodeLocation_CodeLocation).LineNumber)
				}(_5_output)
			}(_pat_let9_0)
		}(Companion_Default___.ExtractLineCol(input, _3_pos))
	})()
	_ = _4_failure
	var _9_failure _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_4_failure, ((result).Dtor_data()).Dtor_message())
	_ = _9_failure
	if (((result).Dtor_data()).Dtor_next()).Is_Some() {
		var _10_failure _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_9_failure, _dafny.UnicodeSeqOfUtf8Bytes(", or\n"))
		_ = _10_failure
		var _11_subFailure _dafny.Sequence = Companion_Default___.FailureToString(input, Companion_ParseResult_.Create_ParseFailure_((result).Dtor_level(), (((result).Dtor_data()).Dtor_next()).Dtor_value().(FailureData)), _3_pos)
		_ = _11_subFailure
		var _12_failure _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_10_failure, _11_subFailure)
		_ = _12_failure
		return _12_failure
	} else {
		var _13_failure _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_9_failure, _dafny.UnicodeSeqOfUtf8Bytes("\n"))
		_ = _13_failure
		return _13_failure
	}
}
func (_static *CompanionStruct_Default___) Apply(parser func(m_Std_Collections_Seq.Slice) ParseResult, input _dafny.Sequence) ParseResult {
	return (parser)(Companion_Default___.ToInput(input))
}
func (_static *CompanionStruct_Default___) ToInput(input _dafny.Sequence) m_Std_Collections_Seq.Slice {
	return m_Std_Collections_Seq.Companion_Slice_.Create_Slice_(input, _dafny.Zero, _dafny.IntOfUint32((input).Cardinality()))
}
func (_static *CompanionStruct_Default___) IsRemaining(input m_Std_Collections_Seq.Slice, remaining m_Std_Collections_Seq.Slice) bool {
	return ((m_Std_Parsers_InputString.Companion_Default___.Length(remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(input)) <= 0) && ((m_Std_Parsers_InputString.Companion_Default___.Drop(input, (m_Std_Parsers_InputString.Companion_Default___.Length(input)).Minus(m_Std_Parsers_InputString.Companion_Default___.Length(remaining)))).Equals(remaining))
}
func (_static *CompanionStruct_Default___) SucceedWith(result interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_result interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) ParseResult {
			return Companion_ParseResult_.Create_ParseSuccess_(_0_result, _1_input)
		}
	})(result)
}
func (_static *CompanionStruct_Default___) Epsilon() func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer83 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg85 m_Std_Collections_Seq.Slice) ParseResult {
			return coer83(arg85)
		}
	}(Companion_Default___.SucceedWith(_dafny.TupleOf()))
}
func (_static *CompanionStruct_Default___) FailWith(message _dafny.Sequence, level FailureLevel) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_level FailureLevel, _1_message _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return Companion_ParseResult_.Create_ParseFailure_(_0_level, Companion_FailureData_.Create_FailureData_(_1_message, _2_input, m_Std_Wrappers.Companion_Option_.Create_None_()))
		}
	})(level, message)
}
func (_static *CompanionStruct_Default___) ResultWith(result ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_result ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) ParseResult {
			return _0_result
		}
	})(result)
}
func (_static *CompanionStruct_Default___) EndOfString() func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(_0_input m_Std_Collections_Seq.Slice) ParseResult {
		return (func() ParseResult {
			if (m_Std_Parsers_InputString.Companion_Default___.Length(_0_input)).Sign() == 0 {
				return Companion_ParseResult_.Create_ParseSuccess_(_dafny.TupleOf(), _0_input)
			}
			return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Recoverable_(), Companion_FailureData_.Create_FailureData_(_dafny.UnicodeSeqOfUtf8Bytes("expected end of string"), _0_input, m_Std_Wrappers.Companion_Option_.Create_None_()))
		})()
	}
}
func (_static *CompanionStruct_Default___) Bind(left func(m_Std_Collections_Seq.Slice) ParseResult, right func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_left func(m_Std_Collections_Seq.Slice) ParseResult, _1_right func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let11_0 ParseResult) ParseResult {
				return func(_3_valueOrError0 ParseResult) ParseResult {
					return (func() ParseResult {
						if (_3_valueOrError0).IsFailure() {
							return (_3_valueOrError0).PropagateFailure()
						}
						return func(_pat_let12_0 _dafny.Tuple) ParseResult {
							return func(_4_leftResult interface{}) ParseResult {
								return func(_5_remaining m_Std_Collections_Seq.Slice) ParseResult {
									return ((_1_right)(_4_leftResult))(_5_remaining)
								}((*(_pat_let12_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
							}((*(_pat_let12_0).IndexInt(0)).(interface{}))
						}((_3_valueOrError0).Extract())
					})()
				}(_pat_let11_0)
			}((_0_left)(_2_input))
		}
	})(left, right)
}
func (_static *CompanionStruct_Default___) BindSucceeds(left func(m_Std_Collections_Seq.Slice) ParseResult, right func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_left func(m_Std_Collections_Seq.Slice) ParseResult, _1_right func(interface{}, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let13_0 ParseResult) ParseResult {
				return func(_3_valueOrError0 ParseResult) ParseResult {
					return (func() ParseResult {
						if (_3_valueOrError0).IsFailure() {
							return (_3_valueOrError0).PropagateFailure()
						}
						return func(_pat_let14_0 _dafny.Tuple) ParseResult {
							return func(_4_leftResult interface{}) ParseResult {
								return func(_5_remaining m_Std_Collections_Seq.Slice) ParseResult {
									return ((_1_right)(_4_leftResult, _5_remaining))(_5_remaining)
								}((*(_pat_let14_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
							}((*(_pat_let14_0).IndexInt(0)).(interface{}))
						}((_3_valueOrError0).Extract())
					})()
				}(_pat_let13_0)
			}((_0_left)(_2_input))
		}
	})(left, right)
}
func (_static *CompanionStruct_Default___) BindResult(left func(m_Std_Collections_Seq.Slice) ParseResult, right func(ParseResult, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_right func(ParseResult, m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) ParseResult, _1_left func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return ((_0_right)((_1_left)(_2_input), _2_input))(_2_input)
		}
	})(right, left)
}
func (_static *CompanionStruct_Default___) Map(underlying func(m_Std_Collections_Seq.Slice) ParseResult, mappingFunc func(interface{}) interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_underlying func(m_Std_Collections_Seq.Slice) ParseResult, _1_mappingFunc func(interface{}) interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let15_0 ParseResult) ParseResult {
				return func(_3_valueOrError0 ParseResult) ParseResult {
					return (func() ParseResult {
						if (_3_valueOrError0).IsFailure() {
							return (_3_valueOrError0).PropagateFailure()
						}
						return func(_pat_let16_0 _dafny.Tuple) ParseResult {
							return func(_4_result interface{}) ParseResult {
								return func(_5_remaining m_Std_Collections_Seq.Slice) ParseResult {
									return func(_pat_let17_0 interface{}) ParseResult {
										return func(_6_u interface{}) ParseResult {
											return Companion_ParseResult_.Create_ParseSuccess_(_6_u, _5_remaining)
										}(_pat_let17_0)
									}((_1_mappingFunc)(_4_result))
								}((*(_pat_let16_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
							}((*(_pat_let16_0).IndexInt(0)).(interface{}))
						}((_3_valueOrError0).Extract())
					})()
				}(_pat_let15_0)
			}((_0_underlying)(_2_input))
		}
	})(underlying, mappingFunc)
}
func (_static *CompanionStruct_Default___) Not(underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let18_0 ParseResult) ParseResult {
				return func(_2_l ParseResult) ParseResult {
					return (func() ParseResult {
						if (_2_l).IsFailure() {
							return (func() ParseResult {
								if (_2_l).IsFatal() {
									return (_2_l).PropagateFailure()
								}
								return Companion_ParseResult_.Create_ParseSuccess_(_dafny.TupleOf(), _1_input)
							})()
						}
						return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Recoverable_(), Companion_FailureData_.Create_FailureData_(_dafny.UnicodeSeqOfUtf8Bytes("not failed"), _1_input, m_Std_Wrappers.Companion_Option_.Create_None_()))
					})()
				}(_pat_let18_0)
			}((_0_underlying)(_1_input))
		}
	})(underlying)
}
func (_static *CompanionStruct_Default___) And(left func(m_Std_Collections_Seq.Slice) ParseResult, right func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_left func(m_Std_Collections_Seq.Slice) ParseResult, _1_right func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let19_0 ParseResult) ParseResult {
				return func(_3_valueOrError0 ParseResult) ParseResult {
					return (func() ParseResult {
						if (_3_valueOrError0).IsFailure() {
							return (_3_valueOrError0).PropagateFailure()
						}
						return func(_pat_let20_0 _dafny.Tuple) ParseResult {
							return func(_4_l interface{}) ParseResult {
								return func(_5_remainingLeft m_Std_Collections_Seq.Slice) ParseResult {
									return func(_pat_let21_0 ParseResult) ParseResult {
										return func(_6_valueOrError1 ParseResult) ParseResult {
											return (func() ParseResult {
												if (_6_valueOrError1).IsFailure() {
													return (_6_valueOrError1).PropagateFailure()
												}
												return func(_pat_let22_0 _dafny.Tuple) ParseResult {
													return func(_7_r interface{}) ParseResult {
														return func(_8_remainingRight m_Std_Collections_Seq.Slice) ParseResult {
															return Companion_ParseResult_.Create_ParseSuccess_(_dafny.TupleOf(_4_l, _7_r), _8_remainingRight)
														}((*(_pat_let22_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
													}((*(_pat_let22_0).IndexInt(0)).(interface{}))
												}((_6_valueOrError1).Extract())
											})()
										}(_pat_let21_0)
									}((_1_right)(_2_input))
								}((*(_pat_let20_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
							}((*(_pat_let20_0).IndexInt(0)).(interface{}))
						}((_3_valueOrError0).Extract())
					})()
				}(_pat_let19_0)
			}((_0_left)(_2_input))
		}
	})(left, right)
}
func (_static *CompanionStruct_Default___) Or(left func(m_Std_Collections_Seq.Slice) ParseResult, right func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_left func(m_Std_Collections_Seq.Slice) ParseResult, _1_right func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let23_0 ParseResult) ParseResult {
				return func(_3_p ParseResult) ParseResult {
					return (func() ParseResult {
						if !((_3_p).NeedsAlternative(_2_input)) {
							return _3_p
						}
						return func(_pat_let24_0 ParseResult) ParseResult {
							return func(_4_p2 ParseResult) ParseResult {
								return (func() ParseResult {
									if !((_4_p2).NeedsAlternative(_2_input)) {
										return _4_p2
									}
									return (_4_p2).MapRecoverableError((func(_5_p ParseResult) func(FailureData) FailureData {
										return func(_6_dataRight FailureData) FailureData {
											return ((_5_p).Dtor_data()).Concat(_6_dataRight)
										}
									})(_3_p))
								})()
							}(_pat_let24_0)
						}((_1_right)(_2_input))
					})()
				}(_pat_let23_0)
			}((_0_left)(_2_input))
		}
	})(left, right)
}
func (_static *CompanionStruct_Default___) OrSeq(alternatives _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
	if (_dafny.IntOfUint32((alternatives).Cardinality())).Sign() == 0 {
		return func(coer84 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(arg86 m_Std_Collections_Seq.Slice) ParseResult {
				return coer84(arg86)
			}
		}(Companion_Default___.FailWith(_dafny.UnicodeSeqOfUtf8Bytes("no alternatives"), Companion_FailureLevel_.Create_Recoverable_()))
	} else if (_dafny.IntOfUint32((alternatives).Cardinality())).Cmp(_dafny.One) == 0 {
		return (alternatives).Select(0).(func(m_Std_Collections_Seq.Slice) ParseResult)
	} else {
		return func(coer85 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(arg87 m_Std_Collections_Seq.Slice) ParseResult {
				return coer85(arg87)
			}
		}(Companion_Default___.Or(func(coer86 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(arg88 m_Std_Collections_Seq.Slice) ParseResult {
				return coer86(arg88)
			}
		}((alternatives).Select(0).(func(m_Std_Collections_Seq.Slice) ParseResult)), func(coer87 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(arg89 m_Std_Collections_Seq.Slice) ParseResult {
				return coer87(arg89)
			}
		}(Companion_Default___.OrSeq((alternatives).Drop(1)))))
	}
}
func (_static *CompanionStruct_Default___) Lookahead(underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let25_0 ParseResult) ParseResult {
				return func(_2_p ParseResult) ParseResult {
					return (func() ParseResult {
						if (_2_p).IsFailure() {
							return (func() ParseResult {
								if (_2_p).IsFatal() {
									return _2_p
								}
								return func(_pat_let26_0 ParseResult) ParseResult {
									return func(_3_dt__update__tmp_h0 ParseResult) ParseResult {
										return func(_pat_let27_0 FailureData) ParseResult {
											return func(_4_dt__update_hdata_h0 FailureData) ParseResult {
												return Companion_ParseResult_.Create_ParseFailure_((_3_dt__update__tmp_h0).Dtor_level(), _4_dt__update_hdata_h0)
											}(_pat_let27_0)
										}(Companion_FailureData_.Create_FailureData_(((_2_p).Dtor_data()).Dtor_message(), _1_input, m_Std_Wrappers.Companion_Option_.Create_None_()))
									}(_pat_let26_0)
								}(_2_p)
							})()
						}
						return func(_pat_let28_0 ParseResult) ParseResult {
							return func(_5_dt__update__tmp_h1 ParseResult) ParseResult {
								return func(_pat_let29_0 m_Std_Collections_Seq.Slice) ParseResult {
									return func(_6_dt__update_hremaining_h0 m_Std_Collections_Seq.Slice) ParseResult {
										return Companion_ParseResult_.Create_ParseSuccess_((_5_dt__update__tmp_h1).Dtor_result(), _6_dt__update_hremaining_h0)
									}(_pat_let29_0)
								}(_1_input)
							}(_pat_let28_0)
						}(_2_p)
					})()
				}(_pat_let25_0)
			}((_0_underlying)(_1_input))
		}
	})(underlying)
}
func (_static *CompanionStruct_Default___) Q_(underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let30_0 ParseResult) ParseResult {
				return func(_2_p ParseResult) ParseResult {
					return (func() ParseResult {
						if (_2_p).IsFailure() {
							return (func() ParseResult {
								if (_2_p).IsFatal() {
									return _2_p
								}
								return func(_pat_let31_0 ParseResult) ParseResult {
									return func(_3_dt__update__tmp_h0 ParseResult) ParseResult {
										return func(_pat_let32_0 FailureData) ParseResult {
											return func(_4_dt__update_hdata_h0 FailureData) ParseResult {
												return Companion_ParseResult_.Create_ParseFailure_((_3_dt__update__tmp_h0).Dtor_level(), _4_dt__update_hdata_h0)
											}(_pat_let32_0)
										}(Companion_FailureData_.Create_FailureData_(((_2_p).Dtor_data()).Dtor_message(), _1_input, m_Std_Wrappers.Companion_Option_.Create_None_()))
									}(_pat_let31_0)
								}(_2_p)
							})()
						}
						return _2_p
					})()
				}(_pat_let30_0)
			}((_0_underlying)(_1_input))
		}
	})(underlying)
}
func (_static *CompanionStruct_Default___) If(condition func(m_Std_Collections_Seq.Slice) ParseResult, succeed func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer88 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg90 m_Std_Collections_Seq.Slice) ParseResult {
			return coer88(arg90)
		}
	}(Companion_Default___.Bind(func(coer89 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg91 m_Std_Collections_Seq.Slice) ParseResult {
			return coer89(arg91)
		}
	}(func(coer90 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg92 m_Std_Collections_Seq.Slice) ParseResult {
			return coer90(arg92)
		}
	}(Companion_Default___.Lookahead(func(coer91 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg93 m_Std_Collections_Seq.Slice) ParseResult {
			return coer91(arg93)
		}
	}(condition)))), func(coer92 func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg94 interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(coer93 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg95 m_Std_Collections_Seq.Slice) ParseResult {
					return coer93(arg95)
				}
			}(coer92(arg94))
		}
	}((func(_0_succeed func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_l interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
			return _0_succeed
		}
	})(succeed))))
}
func (_static *CompanionStruct_Default___) Maybe(underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let33_0 ParseResult) ParseResult {
				return func(_2_u ParseResult) ParseResult {
					return (func() ParseResult {
						if ((_2_u).IsFatalFailure()) || (((_2_u).IsFailure()) && (!((_2_u).NeedsAlternative(_1_input)))) {
							return (_2_u).PropagateFailure()
						}
						return (func() ParseResult {
							if (_2_u).Is_ParseSuccess() {
								return (_2_u).Map(func(coer94 func(interface{}) m_Std_Wrappers.Option) func(interface{}) interface{} {
									return func(arg96 interface{}) interface{} {
										return coer94(arg96)
									}
								}(func(_3_result interface{}) m_Std_Wrappers.Option {
									return m_Std_Wrappers.Companion_Option_.Create_Some_(_3_result)
								}))
							}
							return Companion_ParseResult_.Create_ParseSuccess_(m_Std_Wrappers.Companion_Option_.Create_None_(), _1_input)
						})()
					})()
				}(_pat_let33_0)
			}((_0_underlying)(_1_input))
		}
	})(underlying)
}
func (_static *CompanionStruct_Default___) ConcatMap(left func(m_Std_Collections_Seq.Slice) ParseResult, right func(m_Std_Collections_Seq.Slice) ParseResult, mapper func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_left func(m_Std_Collections_Seq.Slice) ParseResult, _1_right func(m_Std_Collections_Seq.Slice) ParseResult, _2_mapper func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_3_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let34_0 ParseResult) ParseResult {
				return func(_4_valueOrError0 ParseResult) ParseResult {
					return (func() ParseResult {
						if (_4_valueOrError0).IsFailure() {
							return (_4_valueOrError0).PropagateFailure()
						}
						return func(_pat_let35_0 _dafny.Tuple) ParseResult {
							return func(_5_l interface{}) ParseResult {
								return func(_6_remaining m_Std_Collections_Seq.Slice) ParseResult {
									return func(_pat_let36_0 ParseResult) ParseResult {
										return func(_7_valueOrError1 ParseResult) ParseResult {
											return (func() ParseResult {
												if (_7_valueOrError1).IsFailure() {
													return (_7_valueOrError1).PropagateFailure()
												}
												return func(_pat_let37_0 _dafny.Tuple) ParseResult {
													return func(_8_r interface{}) ParseResult {
														return func(_9_remaining2 m_Std_Collections_Seq.Slice) ParseResult {
															return Companion_ParseResult_.Create_ParseSuccess_((_2_mapper)(_5_l, _8_r), _9_remaining2)
														}((*(_pat_let37_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
													}((*(_pat_let37_0).IndexInt(0)).(interface{}))
												}((_7_valueOrError1).Extract())
											})()
										}(_pat_let36_0)
									}((_1_right)(_6_remaining))
								}((*(_pat_let35_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
							}((*(_pat_let35_0).IndexInt(0)).(interface{}))
						}((_4_valueOrError0).Extract())
					})()
				}(_pat_let34_0)
			}((_0_left)(_3_input))
		}
	})(left, right, mapper)
}
func (_static *CompanionStruct_Default___) Concat(left func(m_Std_Collections_Seq.Slice) ParseResult, right func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_left func(m_Std_Collections_Seq.Slice) ParseResult, _1_right func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let38_0 ParseResult) ParseResult {
				return func(_3_valueOrError0 ParseResult) ParseResult {
					return (func() ParseResult {
						if (_3_valueOrError0).IsFailure() {
							return (_3_valueOrError0).PropagateFailure()
						}
						return func(_pat_let39_0 _dafny.Tuple) ParseResult {
							return func(_4_l interface{}) ParseResult {
								return func(_5_remaining m_Std_Collections_Seq.Slice) ParseResult {
									return func(_pat_let40_0 ParseResult) ParseResult {
										return func(_6_valueOrError1 ParseResult) ParseResult {
											return (func() ParseResult {
												if (_6_valueOrError1).IsFailure() {
													return (_6_valueOrError1).PropagateFailure()
												}
												return func(_pat_let41_0 _dafny.Tuple) ParseResult {
													return func(_7_r interface{}) ParseResult {
														return func(_8_remaining2 m_Std_Collections_Seq.Slice) ParseResult {
															return Companion_ParseResult_.Create_ParseSuccess_(_dafny.TupleOf(_4_l, _7_r), _8_remaining2)
														}((*(_pat_let41_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
													}((*(_pat_let41_0).IndexInt(0)).(interface{}))
												}((_6_valueOrError1).Extract())
											})()
										}(_pat_let40_0)
									}((_1_right)(_5_remaining))
								}((*(_pat_let39_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
							}((*(_pat_let39_0).IndexInt(0)).(interface{}))
						}((_3_valueOrError0).Extract())
					})()
				}(_pat_let38_0)
			}((_0_left)(_2_input))
		}
	})(left, right)
}
func (_static *CompanionStruct_Default___) ConcatKeepRight(left func(m_Std_Collections_Seq.Slice) ParseResult, right func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer95 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg97 m_Std_Collections_Seq.Slice) ParseResult {
			return coer95(arg97)
		}
	}(Companion_Default___.ConcatMap(func(coer96 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg98 m_Std_Collections_Seq.Slice) ParseResult {
			return coer96(arg98)
		}
	}(left), func(coer97 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg99 m_Std_Collections_Seq.Slice) ParseResult {
			return coer97(arg99)
		}
	}(right), func(coer98 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
		return func(arg100 interface{}, arg101 interface{}) interface{} {
			return coer98(arg100, arg101)
		}
	}(func(_0_l interface{}, _1_r interface{}) interface{} {
		return _1_r
	})))
}
func (_static *CompanionStruct_Default___) ConcatKeepLeft(left func(m_Std_Collections_Seq.Slice) ParseResult, right func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer99 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg102 m_Std_Collections_Seq.Slice) ParseResult {
			return coer99(arg102)
		}
	}(Companion_Default___.ConcatMap(func(coer100 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg103 m_Std_Collections_Seq.Slice) ParseResult {
			return coer100(arg103)
		}
	}(left), func(coer101 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg104 m_Std_Collections_Seq.Slice) ParseResult {
			return coer101(arg104)
		}
	}(right), func(coer102 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
		return func(arg105 interface{}, arg106 interface{}) interface{} {
			return coer102(arg105, arg106)
		}
	}(func(_0_l interface{}, _1_r interface{}) interface{} {
		return _0_l
	})))
}
func (_static *CompanionStruct_Default___) Debug(underlying func(m_Std_Collections_Seq.Slice) ParseResult, name _dafny.Sequence, onEnter func(_dafny.Sequence, m_Std_Collections_Seq.Slice) interface{}, onExit func(_dafny.Sequence, interface{}, ParseResult) _dafny.Tuple) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_onEnter func(_dafny.Sequence, m_Std_Collections_Seq.Slice) interface{}, _1_name _dafny.Sequence, _2_underlying func(m_Std_Collections_Seq.Slice) ParseResult, _3_onExit func(_dafny.Sequence, interface{}, ParseResult) _dafny.Tuple) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_4_input m_Std_Collections_Seq.Slice) ParseResult {
			return func(_pat_let42_0 interface{}) ParseResult {
				return func(_5_debugData interface{}) ParseResult {
					return func(_pat_let43_0 ParseResult) ParseResult {
						return func(_6_output ParseResult) ParseResult {
							return func(_pat_let44_0 _dafny.Tuple) ParseResult {
								return func(_7___v16 _dafny.Tuple) ParseResult {
									return _6_output
								}(_pat_let44_0)
							}((_3_onExit)(_1_name, _5_debugData, _6_output))
						}(_pat_let43_0)
					}((_2_underlying)(_4_input))
				}(_pat_let42_0)
			}((_0_onEnter)(_1_name, _4_input))
		}
	})(onEnter, name, underlying, onExit)
}
func (_static *CompanionStruct_Default___) ZeroOrMore(underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer103 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg107 m_Std_Collections_Seq.Slice) ParseResult {
			return coer103(arg107)
		}
	}(Companion_Default___.Map(func(coer104 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg108 m_Std_Collections_Seq.Slice) ParseResult {
			return coer104(arg108)
		}
	}(func(coer105 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg109 m_Std_Collections_Seq.Slice) ParseResult {
			return coer105(arg109)
		}
	}(Companion_Default___.Rep(func(coer106 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg110 m_Std_Collections_Seq.Slice) ParseResult {
			return coer106(arg110)
		}
	}(underlying), func(coer107 func(SeqB, interface{}) SeqB) func(interface{}, interface{}) interface{} {
		return func(arg111 interface{}, arg112 interface{}) interface{} {
			return coer107(arg111.(SeqB), arg112)
		}
	}(func(_0_result SeqB, _1_r interface{}) SeqB {
		return (_0_result).Append(_1_r)
	}), Companion_SeqB_.Create_SeqBNil_()))), func(coer108 func(SeqB) _dafny.Sequence) func(interface{}) interface{} {
		return func(arg113 interface{}) interface{} {
			return coer108(arg113.(SeqB))
		}
	}(func(_2_result SeqB) _dafny.Sequence {
		return (_2_result).ToSequence()
	})))
}
func (_static *CompanionStruct_Default___) OneOrMore(underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer109 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg114 m_Std_Collections_Seq.Slice) ParseResult {
			return coer109(arg114)
		}
	}(Companion_Default___.Bind(func(coer110 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg115 m_Std_Collections_Seq.Slice) ParseResult {
			return coer110(arg115)
		}
	}(underlying), func(coer111 func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg116 interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(coer112 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg117 m_Std_Collections_Seq.Slice) ParseResult {
					return coer112(arg117)
				}
			}(coer111(arg116))
		}
	}((func(_0_underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_r interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(coer113 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg118 m_Std_Collections_Seq.Slice) ParseResult {
					return coer113(arg118)
				}
			}(Companion_Default___.Map(func(coer114 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg119 m_Std_Collections_Seq.Slice) ParseResult {
					return coer114(arg119)
				}
			}(func(coer115 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg120 m_Std_Collections_Seq.Slice) ParseResult {
					return coer115(arg120)
				}
			}(Companion_Default___.Rep(func(coer116 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg121 m_Std_Collections_Seq.Slice) ParseResult {
					return coer116(arg121)
				}
			}(_0_underlying), func(coer117 func(SeqB, interface{}) SeqB) func(interface{}, interface{}) interface{} {
				return func(arg122 interface{}, arg123 interface{}) interface{} {
					return coer117(arg122.(SeqB), arg123)
				}
			}(func(_2_result SeqB, _3_r interface{}) SeqB {
				return (_2_result).Append(_3_r)
			}), Companion_SeqB_.Create_SeqBCons_(_1_r, Companion_SeqB_.Create_SeqBNil_())))), func(coer118 func(SeqB) _dafny.Sequence) func(interface{}) interface{} {
				return func(arg124 interface{}) interface{} {
					return coer118(arg124.(SeqB))
				}
			}(func(_4_result SeqB) _dafny.Sequence {
				return (_4_result).ToSequence()
			})))
		}
	})(underlying))))
}
func (_static *CompanionStruct_Default___) Rep(underlying func(m_Std_Collections_Seq.Slice) ParseResult, combine func(interface{}, interface{}) interface{}, acc interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_underlying func(m_Std_Collections_Seq.Slice) ParseResult, _1_combine func(interface{}, interface{}) interface{}, _2_acc interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_3_input m_Std_Collections_Seq.Slice) ParseResult {
			return Companion_Default___.Rep__(func(coer119 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg125 m_Std_Collections_Seq.Slice) ParseResult {
					return coer119(arg125)
				}
			}(_0_underlying), func(coer120 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
				return func(arg126 interface{}, arg127 interface{}) interface{} {
					return coer120(arg126, arg127)
				}
			}(_1_combine), _2_acc, _3_input)
		}
	})(underlying, combine, acc)
}
func (_static *CompanionStruct_Default___) RepSep(underlying func(m_Std_Collections_Seq.Slice) ParseResult, separator func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer121 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg128 m_Std_Collections_Seq.Slice) ParseResult {
			return coer121(arg128)
		}
	}(Companion_Default___.Bind(func(coer122 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg129 m_Std_Collections_Seq.Slice) ParseResult {
			return coer122(arg129)
		}
	}(func(coer123 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg130 m_Std_Collections_Seq.Slice) ParseResult {
			return coer123(arg130)
		}
	}(Companion_Default___.Maybe(func(coer124 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg131 m_Std_Collections_Seq.Slice) ParseResult {
			return coer124(arg131)
		}
	}(underlying)))), func(coer125 func(m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg132 interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(coer126 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg133 m_Std_Collections_Seq.Slice) ParseResult {
					return coer126(arg133)
				}
			}(coer125(arg132.(m_Std_Wrappers.Option)))
		}
	}((func(_0_separator func(m_Std_Collections_Seq.Slice) ParseResult, _1_underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_result m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult {
			return (func() func(m_Std_Collections_Seq.Slice) ParseResult {
				if (_2_result).Is_None() {
					return func(coer127 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
						return func(arg134 m_Std_Collections_Seq.Slice) ParseResult {
							return coer127(arg134)
						}
					}(Companion_Default___.SucceedWith(_dafny.SeqOf()))
				}
				return func(coer128 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg135 m_Std_Collections_Seq.Slice) ParseResult {
						return coer128(arg135)
					}
				}(Companion_Default___.Map(func(coer129 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg136 m_Std_Collections_Seq.Slice) ParseResult {
						return coer129(arg136)
					}
				}(func(coer130 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg137 m_Std_Collections_Seq.Slice) ParseResult {
						return coer130(arg137)
					}
				}(Companion_Default___.Rep(func(coer131 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg138 m_Std_Collections_Seq.Slice) ParseResult {
						return coer131(arg138)
					}
				}(func(coer132 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg139 m_Std_Collections_Seq.Slice) ParseResult {
						return coer132(arg139)
					}
				}(Companion_Default___.ConcatKeepRight(func(coer133 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg140 m_Std_Collections_Seq.Slice) ParseResult {
						return coer133(arg140)
					}
				}(_0_separator), func(coer134 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg141 m_Std_Collections_Seq.Slice) ParseResult {
						return coer134(arg141)
					}
				}(_1_underlying)))), func(coer135 func(SeqB, interface{}) SeqB) func(interface{}, interface{}) interface{} {
					return func(arg142 interface{}, arg143 interface{}) interface{} {
						return coer135(arg142.(SeqB), arg143)
					}
				}(func(_3_acc SeqB, _4_a interface{}) SeqB {
					return (_3_acc).Append(_4_a)
				}), Companion_SeqB_.Create_SeqBCons_((_2_result).Dtor_value(), Companion_SeqB_.Create_SeqBNil_())))), func(coer136 func(SeqB) _dafny.Sequence) func(interface{}) interface{} {
					return func(arg144 interface{}) interface{} {
						return coer136(arg144.(SeqB))
					}
				}(func(_5_result SeqB) _dafny.Sequence {
					return (_5_result).ToSequence()
				})))
			})()
		}
	})(separator, underlying))))
}
func (_static *CompanionStruct_Default___) RepMerge(underlying func(m_Std_Collections_Seq.Slice) ParseResult, merger func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer137 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg145 m_Std_Collections_Seq.Slice) ParseResult {
			return coer137(arg145)
		}
	}(Companion_Default___.Bind(func(coer138 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg146 m_Std_Collections_Seq.Slice) ParseResult {
			return coer138(arg146)
		}
	}(func(coer139 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg147 m_Std_Collections_Seq.Slice) ParseResult {
			return coer139(arg147)
		}
	}(Companion_Default___.Maybe(func(coer140 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg148 m_Std_Collections_Seq.Slice) ParseResult {
			return coer140(arg148)
		}
	}(underlying)))), func(coer141 func(m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg149 interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(coer142 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg150 m_Std_Collections_Seq.Slice) ParseResult {
					return coer142(arg150)
				}
			}(coer141(arg149.(m_Std_Wrappers.Option)))
		}
	}((func(_0_underlying func(m_Std_Collections_Seq.Slice) ParseResult, _1_merger func(interface{}, interface{}) interface{}) func(m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_result m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult {
			return (func() func(m_Std_Collections_Seq.Slice) ParseResult {
				if (_2_result).Is_None() {
					return func(coer143 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
						return func(arg151 m_Std_Collections_Seq.Slice) ParseResult {
							return coer143(arg151)
						}
					}(Companion_Default___.FailWith(_dafny.UnicodeSeqOfUtf8Bytes("No first element in RepMerge"), Companion_FailureLevel_.Create_Recoverable_()))
				}
				return func(coer144 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg152 m_Std_Collections_Seq.Slice) ParseResult {
						return coer144(arg152)
					}
				}(Companion_Default___.Rep(func(coer145 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg153 m_Std_Collections_Seq.Slice) ParseResult {
						return coer145(arg153)
					}
				}(_0_underlying), func(coer146 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
					return func(arg154 interface{}, arg155 interface{}) interface{} {
						return coer146(arg154, arg155)
					}
				}((func(_3_merger func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
					return func(_4_acc interface{}, _5_a interface{}) interface{} {
						return (_3_merger)(_4_acc, _5_a)
					}
				})(_1_merger)), (_2_result).Dtor_value()))
			})()
		}
	})(underlying, merger))))
}
func (_static *CompanionStruct_Default___) RepSepMerge(underlying func(m_Std_Collections_Seq.Slice) ParseResult, separator func(m_Std_Collections_Seq.Slice) ParseResult, merger func(interface{}, interface{}) interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
	return func(coer147 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg156 m_Std_Collections_Seq.Slice) ParseResult {
			return coer147(arg156)
		}
	}(Companion_Default___.Bind(func(coer148 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg157 m_Std_Collections_Seq.Slice) ParseResult {
			return coer148(arg157)
		}
	}(func(coer149 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg158 m_Std_Collections_Seq.Slice) ParseResult {
			return coer149(arg158)
		}
	}(Companion_Default___.Maybe(func(coer150 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg159 m_Std_Collections_Seq.Slice) ParseResult {
			return coer150(arg159)
		}
	}(underlying)))), func(coer151 func(m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(arg160 interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(coer152 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg161 m_Std_Collections_Seq.Slice) ParseResult {
					return coer152(arg161)
				}
			}(coer151(arg160.(m_Std_Wrappers.Option)))
		}
	}((func(_0_separator func(m_Std_Collections_Seq.Slice) ParseResult, _1_underlying func(m_Std_Collections_Seq.Slice) ParseResult, _2_merger func(interface{}, interface{}) interface{}) func(m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_3_result m_Std_Wrappers.Option) func(m_Std_Collections_Seq.Slice) ParseResult {
			return (func() func(m_Std_Collections_Seq.Slice) ParseResult {
				if (_3_result).Is_None() {
					return func(coer153 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
						return func(arg162 m_Std_Collections_Seq.Slice) ParseResult {
							return coer153(arg162)
						}
					}(Companion_Default___.FailWith(_dafny.UnicodeSeqOfUtf8Bytes("No first element in RepSepMerge"), Companion_FailureLevel_.Create_Recoverable_()))
				}
				return func(coer154 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg163 m_Std_Collections_Seq.Slice) ParseResult {
						return coer154(arg163)
					}
				}(Companion_Default___.Rep(func(coer155 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg164 m_Std_Collections_Seq.Slice) ParseResult {
						return coer155(arg164)
					}
				}(func(coer156 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg165 m_Std_Collections_Seq.Slice) ParseResult {
						return coer156(arg165)
					}
				}(Companion_Default___.ConcatKeepRight(func(coer157 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg166 m_Std_Collections_Seq.Slice) ParseResult {
						return coer157(arg166)
					}
				}(_0_separator), func(coer158 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg167 m_Std_Collections_Seq.Slice) ParseResult {
						return coer158(arg167)
					}
				}(_1_underlying)))), func(coer159 func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
					return func(arg168 interface{}, arg169 interface{}) interface{} {
						return coer159(arg168, arg169)
					}
				}((func(_4_merger func(interface{}, interface{}) interface{}) func(interface{}, interface{}) interface{} {
					return func(_5_acc interface{}, _6_a interface{}) interface{} {
						return (_4_merger)(_5_acc, _6_a)
					}
				})(_2_merger)), (_3_result).Dtor_value()))
			})()
		}
	})(separator, underlying, merger))))
}
func (_static *CompanionStruct_Default___) Rep__(underlying func(m_Std_Collections_Seq.Slice) ParseResult, combine func(interface{}, interface{}) interface{}, acc interface{}, input m_Std_Collections_Seq.Slice) ParseResult {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _source0 ParseResult = (underlying)(input)
	_ = _source0
	{
		if _source0.Is_ParseSuccess() {
			var _0_result interface{} = _source0.Get_().(ParseResult_ParseSuccess).Result
			_ = _0_result
			var _1_remaining m_Std_Collections_Seq.Slice = _source0.Get_().(ParseResult_ParseSuccess).Remaining
			_ = _1_remaining
			if (m_Std_Parsers_InputString.Companion_Default___.Length(_1_remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(input)) >= 0 {
				return Companion_ParseResult_.Create_ParseSuccess_(acc, input)
			} else {
				var _in0 func(m_Std_Collections_Seq.Slice) ParseResult = underlying
				_ = _in0
				var _in1 func(interface{}, interface{}) interface{} = combine
				_ = _in1
				var _in2 interface{} = (combine)(acc, _0_result)
				_ = _in2
				var _in3 m_Std_Collections_Seq.Slice = _1_remaining
				_ = _in3
				underlying = _in0
				combine = _in1
				acc = _in2
				input = _in3
				goto TAIL_CALL_START
			}
		}
	}
	{
		var _2_failure ParseResult = _source0
		_ = _2_failure
		if (_2_failure).NeedsAlternative(input) {
			return Companion_ParseResult_.Create_ParseSuccess_(acc, input)
		} else {
			return (_2_failure).PropagateFailure()
		}
	}
}
func (_static *CompanionStruct_Default___) Recursive(underlying func(func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_underlying func(func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) ParseResult {
			return Companion_Default___.Recursive__(func(coer160 func(func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult) func(func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg170 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(coer161 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
						return func(arg171 m_Std_Collections_Seq.Slice) ParseResult {
							return coer161(arg171)
						}
					}(coer160(func(coer162 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
						return func(arg172 m_Std_Collections_Seq.Slice) ParseResult {
							return coer162(arg172)
						}
					}(arg170)))
				}
			}(_0_underlying), _1_input)
		}
	})(underlying)
}
func (_static *CompanionStruct_Default___) RecursiveProgressError(name _dafny.Sequence, input m_Std_Collections_Seq.Slice, remaining m_Std_Collections_Seq.Slice) ParseResult {
	if (m_Std_Parsers_InputString.Companion_Default___.Length(remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(input)) == 0 {
		return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Recoverable_(), Companion_FailureData_.Create_FailureData_(_dafny.Companion_Sequence_.Concatenate(name, _dafny.UnicodeSeqOfUtf8Bytes(" no progress in recursive parser")), remaining, m_Std_Wrappers.Companion_Option_.Create_None_()))
	} else {
		return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Fatal_(), Companion_FailureData_.Create_FailureData_(_dafny.Companion_Sequence_.Concatenate(name, _dafny.UnicodeSeqOfUtf8Bytes("fixpoint called with an increasing remaining sequence")), remaining, m_Std_Wrappers.Companion_Option_.Create_None_()))
	}
}
func (_static *CompanionStruct_Default___) Recursive__(underlying func(func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult, input m_Std_Collections_Seq.Slice) ParseResult {
	var _0_callback func(m_Std_Collections_Seq.Slice) ParseResult = (func(_1_input m_Std_Collections_Seq.Slice, _2_underlying func(func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_3_remaining m_Std_Collections_Seq.Slice) ParseResult {
			return (func() ParseResult {
				if (m_Std_Parsers_InputString.Companion_Default___.Length(_3_remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(_1_input)) < 0 {
					return Companion_Default___.Recursive__(_2_underlying, _3_remaining)
				}
				return Companion_Default___.RecursiveProgressError(_dafny.UnicodeSeqOfUtf8Bytes("Parsers.Recursive"), _1_input, _3_remaining)
			})()
		}
	})(input, underlying)
	_ = _0_callback
	return ((underlying)(_0_callback))(input)
}
func (_static *CompanionStruct_Default___) RecursiveNoStack(underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_underlying func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_1_input m_Std_Collections_Seq.Slice) ParseResult {
			return Companion_Default___.RecursiveNoStack__(func(coer163 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg173 m_Std_Collections_Seq.Slice) ParseResult {
					return coer163(arg173)
				}
			}(_0_underlying), func(coer164 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg174 m_Std_Collections_Seq.Slice) ParseResult {
					return coer164(arg174)
				}
			}(_0_underlying), _1_input, _dafny.SeqOf())
		}
	})(underlying)
}
func (_static *CompanionStruct_Default___) RecursiveNoStack__(continuation func(m_Std_Collections_Seq.Slice) ParseResult, underlying func(m_Std_Collections_Seq.Slice) ParseResult, input m_Std_Collections_Seq.Slice, callbacks _dafny.Sequence) ParseResult {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _0_valueOrError0 ParseResult = (continuation)(input)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _let_tmp_rhs0 _dafny.Tuple = (_0_valueOrError0).Extract()
		_ = _let_tmp_rhs0
		var _1_recursor RecursiveNoStackResult = (*(_let_tmp_rhs0).IndexInt(0)).(RecursiveNoStackResult)
		_ = _1_recursor
		var _2_remaining m_Std_Collections_Seq.Slice = (*(_let_tmp_rhs0).IndexInt(1)).(m_Std_Collections_Seq.Slice)
		_ = _2_remaining
		var _source0 RecursiveNoStackResult = _1_recursor
		_ = _source0
		{
			if _source0.Is_RecursiveReturn() {
				var _3_result interface{} = _source0.Get_().(RecursiveNoStackResult_RecursiveReturn).A0_
				_ = _3_result
				if (_dafny.IntOfUint32((callbacks).Cardinality())).Sign() == 0 {
					return Companion_ParseResult_.Create_ParseSuccess_(_3_result, _2_remaining)
				} else {
					var _4_toCompute func(m_Std_Collections_Seq.Slice) ParseResult = ((callbacks).Select(0).(func(ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult))(Companion_ParseResult_.Create_ParseSuccess_(_3_result, _2_remaining))
					_ = _4_toCompute
					if (m_Std_Parsers_InputString.Companion_Default___.Length(input)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(_2_remaining)) < 0 {
						return Companion_Default___.RecursiveProgressError(_dafny.UnicodeSeqOfUtf8Bytes("Parsers.RecursiveNoStack[internal]"), input, _2_remaining)
					} else {
						var _in0 func(m_Std_Collections_Seq.Slice) ParseResult = _4_toCompute
						_ = _in0
						var _in1 func(m_Std_Collections_Seq.Slice) ParseResult = underlying
						_ = _in1
						var _in2 m_Std_Collections_Seq.Slice = _2_remaining
						_ = _in2
						var _in3 _dafny.Sequence = (callbacks).Drop(1)
						_ = _in3
						continuation = _in0
						underlying = _in1
						input = _in2
						callbacks = _in3
						goto TAIL_CALL_START
					}
				}
			}
		}
		{
			var _5_rToNewParserOfRecursiveNoStackResultOfR func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult = func(coer165 func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult) func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(arg175 interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(coer166 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
						return func(arg176 m_Std_Collections_Seq.Slice) ParseResult {
							return coer166(arg176)
						}
					}(coer165(arg175))
				}
			}(_source0.Get_().(RecursiveNoStackResult_RecursiveContinue).A0_)
			_ = _5_rToNewParserOfRecursiveNoStackResultOfR
			if (m_Std_Parsers_InputString.Companion_Default___.Length(input)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(_2_remaining)) <= 0 {
				return Companion_Default___.RecursiveProgressError(_dafny.UnicodeSeqOfUtf8Bytes("Parsers.RecursiveNoStack"), input, _2_remaining)
			} else {
				var _in4 func(m_Std_Collections_Seq.Slice) ParseResult = underlying
				_ = _in4
				var _in5 func(m_Std_Collections_Seq.Slice) ParseResult = underlying
				_ = _in5
				var _in6 m_Std_Collections_Seq.Slice = _2_remaining
				_ = _in6
				var _in7 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((func(_6_rToNewParserOfRecursiveNoStackResultOfR func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult) func(ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(_7_p ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
						return (func() func(m_Std_Collections_Seq.Slice) ParseResult {
							if (_7_p).IsFailure() {
								return func(coer167 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
									return func(arg177 m_Std_Collections_Seq.Slice) ParseResult {
										return coer167(arg177)
									}
								}(Companion_Default___.ResultWith((_7_p).PropagateFailure()))
							}
							return func(_pat_let45_0 _dafny.Tuple) func(m_Std_Collections_Seq.Slice) ParseResult {
								return func(_8_r interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
									return func(_9_remaining2 m_Std_Collections_Seq.Slice) func(m_Std_Collections_Seq.Slice) ParseResult {
										return (_6_rToNewParserOfRecursiveNoStackResultOfR)(_8_r)
									}((*(_pat_let45_0).IndexInt(1)).(m_Std_Collections_Seq.Slice))
								}((*(_pat_let45_0).IndexInt(0)).(interface{}))
							}((_7_p).Extract())
						})()
					}
				})(_5_rToNewParserOfRecursiveNoStackResultOfR)), callbacks)
				_ = _in7
				continuation = _in4
				underlying = _in5
				input = _in6
				callbacks = _in7
				goto TAIL_CALL_START
			}
		}
	}
}
func (_static *CompanionStruct_Default___) RecursiveMap(underlying _dafny.Map, fun _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_underlying _dafny.Map, _1_fun _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return Companion_Default___.RecursiveMap__(_0_underlying, _1_fun, _2_input)
		}
	})(underlying, fun)
}
func (_static *CompanionStruct_Default___) RecursiveMap__(underlying _dafny.Map, fun _dafny.Sequence, input m_Std_Collections_Seq.Slice) ParseResult {
	if !(underlying).Contains(fun) {
		return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Fatal_(), Companion_FailureData_.Create_FailureData_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("parser '"), fun), _dafny.UnicodeSeqOfUtf8Bytes("' not found")), input, m_Std_Wrappers.Companion_Option_.Create_None_()))
	} else {
		var _let_tmp_rhs0 RecursiveDef = (underlying).Get(fun).(RecursiveDef)
		_ = _let_tmp_rhs0
		var _0_orderFun _dafny.Int = _let_tmp_rhs0.Get_().(RecursiveDef_RecursiveDef).Order
		_ = _0_orderFun
		var _1_definitionFun func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult = func(coer168 func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult) func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(arg178 func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(coer169 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg179 m_Std_Collections_Seq.Slice) ParseResult {
						return coer169(arg179)
					}
				}(coer168(func(coer170 func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(arg180 _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
						return func(coer171 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
							return func(arg181 m_Std_Collections_Seq.Slice) ParseResult {
								return coer171(arg181)
							}
						}(coer170(arg180))
					}
				}(arg178)))
			}
		}(_let_tmp_rhs0.Get_().(RecursiveDef_RecursiveDef).Definition)
		_ = _1_definitionFun
		var _2_callback func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult = (func(_3_underlying _dafny.Map, _4_input m_Std_Collections_Seq.Slice, _5_orderFun _dafny.Int, _6_fun _dafny.Sequence) func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
			return func(_7_fun_k _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
				return func(_pat_let46_0 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
					return func(_17_p func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
						return _17_p
					}(_pat_let46_0)
				}((func() func(m_Std_Collections_Seq.Slice) ParseResult {
					if !((_3_underlying).Keys()).Contains(_7_fun_k) {
						return func(coer172 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
							return func(arg182 m_Std_Collections_Seq.Slice) ParseResult {
								return coer172(arg182)
							}
						}(Companion_Default___.FailWith(_dafny.Companion_Sequence_.Concatenate(_7_fun_k, _dafny.UnicodeSeqOfUtf8Bytes(" not defined")), Companion_FailureLevel_.Create_Fatal_()))
					}
					return func(_pat_let47_0 RecursiveDef) func(m_Std_Collections_Seq.Slice) ParseResult {
						return func(_8_orderFun_k _dafny.Int) func(m_Std_Collections_Seq.Slice) ParseResult {
							return func(_9_definitionFun_k func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
								return (func(_10_input m_Std_Collections_Seq.Slice, _11_orderFun_k _dafny.Int, _12_orderFun _dafny.Int, _13_underlying _dafny.Map, _14_fun_k _dafny.Sequence, _15_fun _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
									return func(_16_remaining m_Std_Collections_Seq.Slice) ParseResult {
										return (func() ParseResult {
											if ((m_Std_Parsers_InputString.Companion_Default___.Length(_16_remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(_10_input)) < 0) || (((m_Std_Parsers_InputString.Companion_Default___.Length(_16_remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(_10_input)) == 0) && ((_11_orderFun_k).Cmp(_12_orderFun) < 0)) {
												return Companion_Default___.RecursiveMap__(_13_underlying, _14_fun_k, _16_remaining)
											}
											return (func() ParseResult {
												if (m_Std_Parsers_InputString.Companion_Default___.Length(_16_remaining)).Cmp(m_Std_Parsers_InputString.Companion_Default___.Length(_10_input)) == 0 {
													return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Recoverable_(), Companion_FailureData_.Create_FailureData_(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("non-progressing recursive call requires that order of '"), _14_fun_k), _dafny.UnicodeSeqOfUtf8Bytes("' (")), m_Std_Strings.Companion_Default___.OfInt(_11_orderFun_k)), _dafny.UnicodeSeqOfUtf8Bytes(") is lower than the order of '")), _15_fun), _dafny.UnicodeSeqOfUtf8Bytes("' (")), m_Std_Strings.Companion_Default___.OfInt(_12_orderFun)), _dafny.UnicodeSeqOfUtf8Bytes(")")), _16_remaining, m_Std_Wrappers.Companion_Option_.Create_None_()))
												}
												return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Fatal_(), Companion_FailureData_.Create_FailureData_(_dafny.UnicodeSeqOfUtf8Bytes("parser did not return a suffix of the input"), _16_remaining, m_Std_Wrappers.Companion_Option_.Create_None_()))
											})()
										})()
									}
								})(_4_input, _8_orderFun_k, _5_orderFun, _3_underlying, _7_fun_k, _6_fun)
							}(func(coer173 func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult) func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
								return func(arg183 func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
									return func(coer174 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
										return func(arg184 m_Std_Collections_Seq.Slice) ParseResult {
											return coer174(arg184)
										}
									}(coer173(func(coer175 func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
										return func(arg185 _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
											return func(coer176 func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
												return func(arg186 m_Std_Collections_Seq.Slice) ParseResult {
													return coer176(arg186)
												}
											}(coer175(arg185))
										}
									}(arg183)))
								}
							}(_pat_let47_0.Get_().(RecursiveDef_RecursiveDef).Definition))
						}(_pat_let47_0.Get_().(RecursiveDef_RecursiveDef).Order)
					}((_3_underlying).Get(_7_fun_k).(RecursiveDef))
				})())
			}
		})(underlying, input, _0_orderFun, fun)
		_ = _2_callback
		return ((_1_definitionFun)(_2_callback))(input)
	}
}
func (_static *CompanionStruct_Default___) CharTest(test func(_dafny.CodePoint) bool, name _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
	return (func(_0_test func(_dafny.CodePoint) bool, _1_name _dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(_2_input m_Std_Collections_Seq.Slice) ParseResult {
			return (func() ParseResult {
				if ((m_Std_Parsers_InputString.Companion_Default___.Length(_2_input)).Sign() == 1) && ((_0_test)(m_Std_Parsers_InputString.Companion_Default___.CharAt(_2_input, _dafny.Zero))) {
					return Companion_ParseResult_.Create_ParseSuccess_(m_Std_Parsers_InputString.Companion_Default___.CharAt(_2_input, _dafny.Zero), m_Std_Parsers_InputString.Companion_Default___.Drop(_2_input, _dafny.One))
				}
				return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Recoverable_(), Companion_FailureData_.Create_FailureData_(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("expected a "), _1_name), _2_input, m_Std_Wrappers.Companion_Option_.Create_None_()))
			})()
		}
	})(test, name)
}
func (_static *CompanionStruct_Default___) DigitToInt(c _dafny.CodePoint) _dafny.Int {
	var _source0 _dafny.CodePoint = c
	_ = _source0
	{
		if (_source0) == (_dafny.CodePoint('0')) {
			return _dafny.Zero
		}
	}
	{
		if (_source0) == (_dafny.CodePoint('1')) {
			return _dafny.One
		}
	}
	{
		if (_source0) == (_dafny.CodePoint('2')) {
			return _dafny.IntOfInt64(2)
		}
	}
	{
		if (_source0) == (_dafny.CodePoint('3')) {
			return _dafny.IntOfInt64(3)
		}
	}
	{
		if (_source0) == (_dafny.CodePoint('4')) {
			return _dafny.IntOfInt64(4)
		}
	}
	{
		if (_source0) == (_dafny.CodePoint('5')) {
			return _dafny.IntOfInt64(5)
		}
	}
	{
		if (_source0) == (_dafny.CodePoint('6')) {
			return _dafny.IntOfInt64(6)
		}
	}
	{
		if (_source0) == (_dafny.CodePoint('7')) {
			return _dafny.IntOfInt64(7)
		}
	}
	{
		if (_source0) == (_dafny.CodePoint('8')) {
			return _dafny.IntOfInt64(8)
		}
	}
	{
		if (_source0) == (_dafny.CodePoint('9')) {
			return _dafny.IntOfInt64(9)
		}
	}
	{
		return _dafny.IntOfInt64(-1)
	}
}
func (_static *CompanionStruct_Default___) StringToInt(s _dafny.Sequence) _dafny.Int {
	if (_dafny.IntOfUint32((s).Cardinality())).Sign() == 0 {
		return _dafny.Zero
	} else if (_dafny.IntOfUint32((s).Cardinality())).Cmp(_dafny.One) == 0 {
		return Companion_Default___.DigitToInt((s).Select(0).(_dafny.CodePoint))
	} else if ((s).Select(0).(_dafny.CodePoint)) == (_dafny.CodePoint('-')) {
		return (_dafny.Zero).Minus(Companion_Default___.StringToInt((s).Drop(1)))
	} else {
		return ((Companion_Default___.StringToInt((s).Subsequence(0, ((_dafny.IntOfUint32((s).Cardinality())).Minus(_dafny.One)).Uint32()))).Times(_dafny.IntOfInt64(10))).Plus(Companion_Default___.StringToInt((s).Subsequence(((_dafny.IntOfUint32((s).Cardinality())).Minus(_dafny.One)).Uint32(), (_dafny.IntOfUint32((s).Cardinality())).Uint32())))
	}
}

// End of class Default__

// Definition of datatype CodeLocation
type CodeLocation struct {
	Data_CodeLocation_
}

func (_this CodeLocation) Get_() Data_CodeLocation_ {
	return _this.Data_CodeLocation_
}

type Data_CodeLocation_ interface {
	isCodeLocation()
}

type CompanionStruct_CodeLocation_ struct {
}

var Companion_CodeLocation_ = CompanionStruct_CodeLocation_{}

type CodeLocation_CodeLocation struct {
	LineNumber _dafny.Int
	ColNumber  _dafny.Int
	LineStr    _dafny.Sequence
}

func (CodeLocation_CodeLocation) isCodeLocation() {}

func (CompanionStruct_CodeLocation_) Create_CodeLocation_(LineNumber _dafny.Int, ColNumber _dafny.Int, LineStr _dafny.Sequence) CodeLocation {
	return CodeLocation{CodeLocation_CodeLocation{LineNumber, ColNumber, LineStr}}
}

func (_this CodeLocation) Is_CodeLocation() bool {
	_, ok := _this.Get_().(CodeLocation_CodeLocation)
	return ok
}

func (CompanionStruct_CodeLocation_) Default() CodeLocation {
	return Companion_CodeLocation_.Create_CodeLocation_(_dafny.Zero, _dafny.Zero, _dafny.EmptySeq)
}

func (_this CodeLocation) Dtor_lineNumber() _dafny.Int {
	return _this.Get_().(CodeLocation_CodeLocation).LineNumber
}

func (_this CodeLocation) Dtor_colNumber() _dafny.Int {
	return _this.Get_().(CodeLocation_CodeLocation).ColNumber
}

func (_this CodeLocation) Dtor_lineStr() _dafny.Sequence {
	return _this.Get_().(CodeLocation_CodeLocation).LineStr
}

func (_this CodeLocation) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case CodeLocation_CodeLocation:
		{
			return "StringParsers.CodeLocation.CodeLocation" + "(" + _dafny.String(data.LineNumber) + ", " + _dafny.String(data.ColNumber) + ", " + data.LineStr.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this CodeLocation) Equals(other CodeLocation) bool {
	switch data1 := _this.Get_().(type) {
	case CodeLocation_CodeLocation:
		{
			data2, ok := other.Get_().(CodeLocation_CodeLocation)
			return ok && data1.LineNumber.Cmp(data2.LineNumber) == 0 && data1.ColNumber.Cmp(data2.ColNumber) == 0 && data1.LineStr.Equals(data2.LineStr)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this CodeLocation) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(CodeLocation)
	return ok && _this.Equals(typed)
}

func Type_CodeLocation_() _dafny.TypeDescriptor {
	return type_CodeLocation_{}
}

type type_CodeLocation_ struct {
}

func (_this type_CodeLocation_) Default() interface{} {
	return Companion_CodeLocation_.Default()
}

func (_this type_CodeLocation_) String() string {
	return "Std_Parsers_StringParsers.CodeLocation"
}
func (_this CodeLocation) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = CodeLocation{}

// End of datatype CodeLocation

// Definition of datatype ExtractLineMutableState
type ExtractLineMutableState struct {
	Data_ExtractLineMutableState_
}

func (_this ExtractLineMutableState) Get_() Data_ExtractLineMutableState_ {
	return _this.Data_ExtractLineMutableState_
}

type Data_ExtractLineMutableState_ interface {
	isExtractLineMutableState()
}

type CompanionStruct_ExtractLineMutableState_ struct {
}

var Companion_ExtractLineMutableState_ = CompanionStruct_ExtractLineMutableState_{}

type ExtractLineMutableState_ExtractLineMutableState struct {
	Input        _dafny.Sequence
	Pos          _dafny.Int
	StartLinePos _dafny.Int
	I            _dafny.Int
	LineNumber   _dafny.Int
	ColNumber    _dafny.Int
}

func (ExtractLineMutableState_ExtractLineMutableState) isExtractLineMutableState() {}

func (CompanionStruct_ExtractLineMutableState_) Create_ExtractLineMutableState_(Input _dafny.Sequence, Pos _dafny.Int, StartLinePos _dafny.Int, I _dafny.Int, LineNumber _dafny.Int, ColNumber _dafny.Int) ExtractLineMutableState {
	return ExtractLineMutableState{ExtractLineMutableState_ExtractLineMutableState{Input, Pos, StartLinePos, I, LineNumber, ColNumber}}
}

func (_this ExtractLineMutableState) Is_ExtractLineMutableState() bool {
	_, ok := _this.Get_().(ExtractLineMutableState_ExtractLineMutableState)
	return ok
}

func (CompanionStruct_ExtractLineMutableState_) Default() ExtractLineMutableState {
	return Companion_ExtractLineMutableState_.Create_ExtractLineMutableState_(_dafny.EmptySeq, _dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this ExtractLineMutableState) Dtor_input() _dafny.Sequence {
	return _this.Get_().(ExtractLineMutableState_ExtractLineMutableState).Input
}

func (_this ExtractLineMutableState) Dtor_pos() _dafny.Int {
	return _this.Get_().(ExtractLineMutableState_ExtractLineMutableState).Pos
}

func (_this ExtractLineMutableState) Dtor_startLinePos() _dafny.Int {
	return _this.Get_().(ExtractLineMutableState_ExtractLineMutableState).StartLinePos
}

func (_this ExtractLineMutableState) Dtor_i() _dafny.Int {
	return _this.Get_().(ExtractLineMutableState_ExtractLineMutableState).I
}

func (_this ExtractLineMutableState) Dtor_lineNumber() _dafny.Int {
	return _this.Get_().(ExtractLineMutableState_ExtractLineMutableState).LineNumber
}

func (_this ExtractLineMutableState) Dtor_colNumber() _dafny.Int {
	return _this.Get_().(ExtractLineMutableState_ExtractLineMutableState).ColNumber
}

func (_this ExtractLineMutableState) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ExtractLineMutableState_ExtractLineMutableState:
		{
			return "StringParsers.ExtractLineMutableState.ExtractLineMutableState" + "(" + data.Input.VerbatimString(true) + ", " + _dafny.String(data.Pos) + ", " + _dafny.String(data.StartLinePos) + ", " + _dafny.String(data.I) + ", " + _dafny.String(data.LineNumber) + ", " + _dafny.String(data.ColNumber) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ExtractLineMutableState) Equals(other ExtractLineMutableState) bool {
	switch data1 := _this.Get_().(type) {
	case ExtractLineMutableState_ExtractLineMutableState:
		{
			data2, ok := other.Get_().(ExtractLineMutableState_ExtractLineMutableState)
			return ok && data1.Input.Equals(data2.Input) && data1.Pos.Cmp(data2.Pos) == 0 && data1.StartLinePos.Cmp(data2.StartLinePos) == 0 && data1.I.Cmp(data2.I) == 0 && data1.LineNumber.Cmp(data2.LineNumber) == 0 && data1.ColNumber.Cmp(data2.ColNumber) == 0
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ExtractLineMutableState) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ExtractLineMutableState)
	return ok && _this.Equals(typed)
}

func Type_ExtractLineMutableState_() _dafny.TypeDescriptor {
	return type_ExtractLineMutableState_{}
}

type type_ExtractLineMutableState_ struct {
}

func (_this type_ExtractLineMutableState_) Default() interface{} {
	return Companion_ExtractLineMutableState_.Default()
}

func (_this type_ExtractLineMutableState_) String() string {
	return "Std_Parsers_StringParsers.ExtractLineMutableState"
}
func (_this ExtractLineMutableState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ExtractLineMutableState{}

// End of datatype ExtractLineMutableState

// Definition of datatype FailureData
type FailureData struct {
	Data_FailureData_
}

func (_this FailureData) Get_() Data_FailureData_ {
	return _this.Data_FailureData_
}

type Data_FailureData_ interface {
	isFailureData()
}

type CompanionStruct_FailureData_ struct {
}

var Companion_FailureData_ = CompanionStruct_FailureData_{}

type FailureData_FailureData struct {
	Message   _dafny.Sequence
	Remaining m_Std_Collections_Seq.Slice
	Next      m_Std_Wrappers.Option
}

func (FailureData_FailureData) isFailureData() {}

func (CompanionStruct_FailureData_) Create_FailureData_(Message _dafny.Sequence, Remaining m_Std_Collections_Seq.Slice, Next m_Std_Wrappers.Option) FailureData {
	return FailureData{FailureData_FailureData{Message, Remaining, Next}}
}

func (_this FailureData) Is_FailureData() bool {
	_, ok := _this.Get_().(FailureData_FailureData)
	return ok
}

func (CompanionStruct_FailureData_) Default() FailureData {
	return Companion_FailureData_.Create_FailureData_(_dafny.EmptySeq, m_Std_Collections_Seq.Companion_Slice_.Default(), m_Std_Wrappers.Companion_Option_.Default())
}

func (_this FailureData) Dtor_message() _dafny.Sequence {
	return _this.Get_().(FailureData_FailureData).Message
}

func (_this FailureData) Dtor_remaining() m_Std_Collections_Seq.Slice {
	return _this.Get_().(FailureData_FailureData).Remaining
}

func (_this FailureData) Dtor_next() m_Std_Wrappers.Option {
	return _this.Get_().(FailureData_FailureData).Next
}

func (_this FailureData) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case FailureData_FailureData:
		{
			return "StringParsers.FailureData.FailureData" + "(" + data.Message.VerbatimString(true) + ", " + _dafny.String(data.Remaining) + ", " + _dafny.String(data.Next) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this FailureData) Equals(other FailureData) bool {
	switch data1 := _this.Get_().(type) {
	case FailureData_FailureData:
		{
			data2, ok := other.Get_().(FailureData_FailureData)
			return ok && data1.Message.Equals(data2.Message) && data1.Remaining.Equals(data2.Remaining) && data1.Next.Equals(data2.Next)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this FailureData) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(FailureData)
	return ok && _this.Equals(typed)
}

func Type_FailureData_() _dafny.TypeDescriptor {
	return type_FailureData_{}
}

type type_FailureData_ struct {
}

func (_this type_FailureData_) Default() interface{} {
	return Companion_FailureData_.Default()
}

func (_this type_FailureData_) String() string {
	return "Std_Parsers_StringParsers.FailureData"
}
func (_this FailureData) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = FailureData{}

func (_this FailureData) Concat(other FailureData) FailureData {
	{
		if ((_this).Dtor_next()).Equals(m_Std_Wrappers.Companion_Option_.Create_None_()) {
			var _0_dt__update__tmp_h0 FailureData = _this
			_ = _0_dt__update__tmp_h0
			var _1_dt__update_hnext_h0 m_Std_Wrappers.Option = m_Std_Wrappers.Companion_Option_.Create_Some_(other)
			_ = _1_dt__update_hnext_h0
			return Companion_FailureData_.Create_FailureData_((_0_dt__update__tmp_h0).Dtor_message(), (_0_dt__update__tmp_h0).Dtor_remaining(), _1_dt__update_hnext_h0)
		} else {
			return Companion_FailureData_.Create_FailureData_((_this).Dtor_message(), (_this).Dtor_remaining(), m_Std_Wrappers.Companion_Option_.Create_Some_((((_this).Dtor_next()).Dtor_value().(FailureData)).Concat(other)))
		}
	}
}

// End of datatype FailureData

// Definition of datatype FailureLevel
type FailureLevel struct {
	Data_FailureLevel_
}

func (_this FailureLevel) Get_() Data_FailureLevel_ {
	return _this.Data_FailureLevel_
}

type Data_FailureLevel_ interface {
	isFailureLevel()
}

type CompanionStruct_FailureLevel_ struct {
}

var Companion_FailureLevel_ = CompanionStruct_FailureLevel_{}

type FailureLevel_Fatal struct {
}

func (FailureLevel_Fatal) isFailureLevel() {}

func (CompanionStruct_FailureLevel_) Create_Fatal_() FailureLevel {
	return FailureLevel{FailureLevel_Fatal{}}
}

func (_this FailureLevel) Is_Fatal() bool {
	_, ok := _this.Get_().(FailureLevel_Fatal)
	return ok
}

type FailureLevel_Recoverable struct {
}

func (FailureLevel_Recoverable) isFailureLevel() {}

func (CompanionStruct_FailureLevel_) Create_Recoverable_() FailureLevel {
	return FailureLevel{FailureLevel_Recoverable{}}
}

func (_this FailureLevel) Is_Recoverable() bool {
	_, ok := _this.Get_().(FailureLevel_Recoverable)
	return ok
}

func (CompanionStruct_FailureLevel_) Default() FailureLevel {
	return Companion_FailureLevel_.Create_Fatal_()
}

func (_ CompanionStruct_FailureLevel_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_FailureLevel_.Create_Fatal_(), true
		case 1:
			return Companion_FailureLevel_.Create_Recoverable_(), true
		default:
			return FailureLevel{}, false
		}
	}
}

func (_this FailureLevel) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case FailureLevel_Fatal:
		{
			return "StringParsers.FailureLevel.Fatal"
		}
	case FailureLevel_Recoverable:
		{
			return "StringParsers.FailureLevel.Recoverable"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this FailureLevel) Equals(other FailureLevel) bool {
	switch _this.Get_().(type) {
	case FailureLevel_Fatal:
		{
			_, ok := other.Get_().(FailureLevel_Fatal)
			return ok
		}
	case FailureLevel_Recoverable:
		{
			_, ok := other.Get_().(FailureLevel_Recoverable)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this FailureLevel) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(FailureLevel)
	return ok && _this.Equals(typed)
}

func Type_FailureLevel_() _dafny.TypeDescriptor {
	return type_FailureLevel_{}
}

type type_FailureLevel_ struct {
}

func (_this type_FailureLevel_) Default() interface{} {
	return Companion_FailureLevel_.Default()
}

func (_this type_FailureLevel_) String() string {
	return "Std_Parsers_StringParsers.FailureLevel"
}
func (_this FailureLevel) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = FailureLevel{}

// End of datatype FailureLevel

// Definition of datatype ParseResult
type ParseResult struct {
	Data_ParseResult_
}

func (_this ParseResult) Get_() Data_ParseResult_ {
	return _this.Data_ParseResult_
}

type Data_ParseResult_ interface {
	isParseResult()
}

type CompanionStruct_ParseResult_ struct {
}

var Companion_ParseResult_ = CompanionStruct_ParseResult_{}

type ParseResult_ParseFailure struct {
	Level FailureLevel
	Data  FailureData
}

func (ParseResult_ParseFailure) isParseResult() {}

func (CompanionStruct_ParseResult_) Create_ParseFailure_(Level FailureLevel, Data FailureData) ParseResult {
	return ParseResult{ParseResult_ParseFailure{Level, Data}}
}

func (_this ParseResult) Is_ParseFailure() bool {
	_, ok := _this.Get_().(ParseResult_ParseFailure)
	return ok
}

type ParseResult_ParseSuccess struct {
	Result    interface{}
	Remaining m_Std_Collections_Seq.Slice
}

func (ParseResult_ParseSuccess) isParseResult() {}

func (CompanionStruct_ParseResult_) Create_ParseSuccess_(Result interface{}, Remaining m_Std_Collections_Seq.Slice) ParseResult {
	return ParseResult{ParseResult_ParseSuccess{Result, Remaining}}
}

func (_this ParseResult) Is_ParseSuccess() bool {
	_, ok := _this.Get_().(ParseResult_ParseSuccess)
	return ok
}

func (CompanionStruct_ParseResult_) Default() ParseResult {
	return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Default(), Companion_FailureData_.Default())
}

func (_this ParseResult) Dtor_level() FailureLevel {
	return _this.Get_().(ParseResult_ParseFailure).Level
}

func (_this ParseResult) Dtor_data() FailureData {
	return _this.Get_().(ParseResult_ParseFailure).Data
}

func (_this ParseResult) Dtor_result() interface{} {
	return _this.Get_().(ParseResult_ParseSuccess).Result
}

func (_this ParseResult) Dtor_remaining() m_Std_Collections_Seq.Slice {
	return _this.Get_().(ParseResult_ParseSuccess).Remaining
}

func (_this ParseResult) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case ParseResult_ParseFailure:
		{
			return "StringParsers.ParseResult.ParseFailure" + "(" + _dafny.String(data.Level) + ", " + _dafny.String(data.Data) + ")"
		}
	case ParseResult_ParseSuccess:
		{
			return "StringParsers.ParseResult.ParseSuccess" + "(" + _dafny.String(data.Result) + ", " + _dafny.String(data.Remaining) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ParseResult) Equals(other ParseResult) bool {
	switch data1 := _this.Get_().(type) {
	case ParseResult_ParseFailure:
		{
			data2, ok := other.Get_().(ParseResult_ParseFailure)
			return ok && data1.Level.Equals(data2.Level) && data1.Data.Equals(data2.Data)
		}
	case ParseResult_ParseSuccess:
		{
			data2, ok := other.Get_().(ParseResult_ParseSuccess)
			return ok && _dafny.AreEqual(data1.Result, data2.Result) && data1.Remaining.Equals(data2.Remaining)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this ParseResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(ParseResult)
	return ok && _this.Equals(typed)
}

func Type_ParseResult_() _dafny.TypeDescriptor {
	return type_ParseResult_{}
}

type type_ParseResult_ struct {
}

func (_this type_ParseResult_) Default() interface{} {
	return Companion_ParseResult_.Default()
}

func (_this type_ParseResult_) String() string {
	return "Std_Parsers_StringParsers.ParseResult"
}
func (_this ParseResult) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = ParseResult{}

func (_this ParseResult) Remaining() m_Std_Collections_Seq.Slice {
	{
		if (_this).Is_ParseSuccess() {
			return (_this).Dtor_remaining()
		} else {
			return ((_this).Dtor_data()).Dtor_remaining()
		}
	}
}
func (_this ParseResult) IsFailure() bool {
	{
		return (_this).Is_ParseFailure()
	}
}
func (_this ParseResult) IsFatalFailure() bool {
	{
		return ((_this).Is_ParseFailure()) && (((_this).Dtor_level()).Equals(Companion_FailureLevel_.Create_Fatal_()))
	}
}
func (_this ParseResult) IsFatal() bool {
	{
		return ((_this).Dtor_level()).Equals(Companion_FailureLevel_.Create_Fatal_())
	}
}
func (_this ParseResult) PropagateFailure() ParseResult {
	{
		return Companion_ParseResult_.Create_ParseFailure_((_this).Dtor_level(), (_this).Dtor_data())
	}
}
func (_this ParseResult) Extract() _dafny.Tuple {
	{
		return _dafny.TupleOf((_this).Dtor_result(), (_this).Dtor_remaining())
	}
}
func (_this ParseResult) Map(f func(interface{}) interface{}) ParseResult {
	{
		var _source0 ParseResult = _this
		_ = _source0
		{
			if _source0.Is_ParseSuccess() {
				var _0_result interface{} = _source0.Get_().(ParseResult_ParseSuccess).Result
				_ = _0_result
				var _1_remaining m_Std_Collections_Seq.Slice = _source0.Get_().(ParseResult_ParseSuccess).Remaining
				_ = _1_remaining
				return Companion_ParseResult_.Create_ParseSuccess_((f)(_0_result), _1_remaining)
			}
		}
		{
			var _2_level FailureLevel = _source0.Get_().(ParseResult_ParseFailure).Level
			_ = _2_level
			var _3_data FailureData = _source0.Get_().(ParseResult_ParseFailure).Data
			_ = _3_data
			return Companion_ParseResult_.Create_ParseFailure_(_2_level, _3_data)
		}
	}
}
func (_this ParseResult) MapRecoverableError(f func(FailureData) FailureData) ParseResult {
	{
		var _source0 ParseResult = _this
		_ = _source0
		{
			if _source0.Is_ParseFailure() {
				var level0 FailureLevel = _source0.Get_().(ParseResult_ParseFailure).Level
				_ = level0
				if level0.Is_Recoverable() {
					var _0_data FailureData = _source0.Get_().(ParseResult_ParseFailure).Data
					_ = _0_data
					return Companion_ParseResult_.Create_ParseFailure_(Companion_FailureLevel_.Create_Recoverable_(), (f)(_0_data))
				}
			}
		}
		{
			return _this
		}
	}
}
func (_this ParseResult) NeedsAlternative(input m_Std_Collections_Seq.Slice) bool {
	{
		return (((_this).Is_ParseFailure()) && (((_this).Dtor_level()).Equals(Companion_FailureLevel_.Create_Recoverable_()))) && ((input).Equals((_this).Remaining()))
	}
}

// End of datatype ParseResult

// Definition of datatype SeqB
type SeqB struct {
	Data_SeqB_
}

func (_this SeqB) Get_() Data_SeqB_ {
	return _this.Data_SeqB_
}

type Data_SeqB_ interface {
	isSeqB()
}

type CompanionStruct_SeqB_ struct {
}

var Companion_SeqB_ = CompanionStruct_SeqB_{}

type SeqB_SeqBCons struct {
	Last interface{}
	Init SeqB
}

func (SeqB_SeqBCons) isSeqB() {}

func (CompanionStruct_SeqB_) Create_SeqBCons_(Last interface{}, Init SeqB) SeqB {
	return SeqB{SeqB_SeqBCons{Last, Init}}
}

func (_this SeqB) Is_SeqBCons() bool {
	_, ok := _this.Get_().(SeqB_SeqBCons)
	return ok
}

type SeqB_SeqBNil struct {
}

func (SeqB_SeqBNil) isSeqB() {}

func (CompanionStruct_SeqB_) Create_SeqBNil_() SeqB {
	return SeqB{SeqB_SeqBNil{}}
}

func (_this SeqB) Is_SeqBNil() bool {
	_, ok := _this.Get_().(SeqB_SeqBNil)
	return ok
}

func (CompanionStruct_SeqB_) Default() SeqB {
	return Companion_SeqB_.Create_SeqBNil_()
}

func (_this SeqB) Dtor_last() interface{} {
	return _this.Get_().(SeqB_SeqBCons).Last
}

func (_this SeqB) Dtor_init() SeqB {
	return _this.Get_().(SeqB_SeqBCons).Init
}

func (_this SeqB) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case SeqB_SeqBCons:
		{
			return "StringParsers.SeqB.SeqBCons" + "(" + _dafny.String(data.Last) + ", " + _dafny.String(data.Init) + ")"
		}
	case SeqB_SeqBNil:
		{
			return "StringParsers.SeqB.SeqBNil"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SeqB) Equals(other SeqB) bool {
	switch data1 := _this.Get_().(type) {
	case SeqB_SeqBCons:
		{
			data2, ok := other.Get_().(SeqB_SeqBCons)
			return ok && _dafny.AreEqual(data1.Last, data2.Last) && data1.Init.Equals(data2.Init)
		}
	case SeqB_SeqBNil:
		{
			_, ok := other.Get_().(SeqB_SeqBNil)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SeqB) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SeqB)
	return ok && _this.Equals(typed)
}

func Type_SeqB_() _dafny.TypeDescriptor {
	return type_SeqB_{}
}

type type_SeqB_ struct {
}

func (_this type_SeqB_) Default() interface{} {
	return Companion_SeqB_.Default()
}

func (_this type_SeqB_) String() string {
	return "Std_Parsers_StringParsers.SeqB"
}
func (_this SeqB) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SeqB{}

func (_this SeqB) Append(elem interface{}) SeqB {
	{
		return Companion_SeqB_.Create_SeqBCons_(elem, _this)
	}
}
func (_this SeqB) Length() _dafny.Int {
	{
		var _0___accumulator _dafny.Int = _dafny.Zero
		_ = _0___accumulator
		goto TAIL_CALL_START
	TAIL_CALL_START:
		if (_this).Is_SeqBNil() {
			return (_dafny.Zero).Plus(_0___accumulator)
		} else {
			_0___accumulator = (_0___accumulator).Plus(_dafny.One)
			var _in0 SeqB = (_this).Dtor_init()
			_ = _in0
			_this = _in0

			goto TAIL_CALL_START
		}
	}
}
func (_this SeqB) ToSequence() _dafny.Sequence {
	{
		var _hresult _dafny.Sequence = _dafny.EmptySeq
		_ = _hresult
		if (_this).Is_SeqBNil() {
			_hresult = _dafny.SeqOf()
			return _hresult
		}
		var _0_defaultElem interface{}
		_ = _0_defaultElem
		_0_defaultElem = (_this).Dtor_last()
		var _1_l _dafny.Int
		_ = _1_l
		_1_l = (_this).Length()
		var _2_elements _dafny.Array
		_ = _2_elements
		var _len0_0 _dafny.Int = _1_l
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) interface{} = (func(_3_defaultElem interface{}) func(_dafny.Int) interface{} {
				return func(_4_i _dafny.Int) interface{} {
					return _3_defaultElem
				}
			})(_0_defaultElem)
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
		_2_elements = _nw0
		var _5_t SeqB
		_ = _5_t
		_5_t = _this
		var _6_i _dafny.Int
		_ = _6_i
		_6_i = _1_l
		for !((_5_t).Is_SeqBNil()) {
			_6_i = (_6_i).Minus(_dafny.One)
			(_2_elements).ArraySet1((_5_t).Dtor_last(), (_6_i).Int())
			_5_t = (_5_t).Dtor_init()
		}
		_hresult = _dafny.ArrayRangeToSeq(_2_elements, _dafny.NilInt, _dafny.NilInt)
		return _hresult
		return _hresult
	}
}

// End of datatype SeqB

// Definition of datatype RecursiveNoStackResult
type RecursiveNoStackResult struct {
	Data_RecursiveNoStackResult_
}

func (_this RecursiveNoStackResult) Get_() Data_RecursiveNoStackResult_ {
	return _this.Data_RecursiveNoStackResult_
}

type Data_RecursiveNoStackResult_ interface {
	isRecursiveNoStackResult()
}

type CompanionStruct_RecursiveNoStackResult_ struct {
}

var Companion_RecursiveNoStackResult_ = CompanionStruct_RecursiveNoStackResult_{}

type RecursiveNoStackResult_RecursiveReturn struct {
	A0_ interface{}
}

func (RecursiveNoStackResult_RecursiveReturn) isRecursiveNoStackResult() {}

func (CompanionStruct_RecursiveNoStackResult_) Create_RecursiveReturn_(A0_ interface{}) RecursiveNoStackResult {
	return RecursiveNoStackResult{RecursiveNoStackResult_RecursiveReturn{A0_}}
}

func (_this RecursiveNoStackResult) Is_RecursiveReturn() bool {
	_, ok := _this.Get_().(RecursiveNoStackResult_RecursiveReturn)
	return ok
}

type RecursiveNoStackResult_RecursiveContinue struct {
	A0_ func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult
}

func (RecursiveNoStackResult_RecursiveContinue) isRecursiveNoStackResult() {}

func (CompanionStruct_RecursiveNoStackResult_) Create_RecursiveContinue_(A0_ func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult) RecursiveNoStackResult {
	return RecursiveNoStackResult{RecursiveNoStackResult_RecursiveContinue{A0_}}
}

func (_this RecursiveNoStackResult) Is_RecursiveContinue() bool {
	_, ok := _this.Get_().(RecursiveNoStackResult_RecursiveContinue)
	return ok
}

func (CompanionStruct_RecursiveNoStackResult_) Default() RecursiveNoStackResult {
	return Companion_RecursiveNoStackResult_.Create_RecursiveContinue_(func(interface{}) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(m_Std_Collections_Seq.Slice) ParseResult { return Companion_ParseResult_.Default() }
	})
}

func (_this RecursiveNoStackResult) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RecursiveNoStackResult_RecursiveReturn:
		{
			return "StringParsers.RecursiveNoStackResult.RecursiveReturn" + "(" + _dafny.String(data.A0_) + ")"
		}
	case RecursiveNoStackResult_RecursiveContinue:
		{
			return "StringParsers.RecursiveNoStackResult.RecursiveContinue" + "(" + _dafny.String(data.A0_) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RecursiveNoStackResult) Equals(other RecursiveNoStackResult) bool {
	switch data1 := _this.Get_().(type) {
	case RecursiveNoStackResult_RecursiveReturn:
		{
			data2, ok := other.Get_().(RecursiveNoStackResult_RecursiveReturn)
			return ok && _dafny.AreEqual(data1.A0_, data2.A0_)
		}
	case RecursiveNoStackResult_RecursiveContinue:
		{
			data2, ok := other.Get_().(RecursiveNoStackResult_RecursiveContinue)
			return ok && _dafny.AreEqual(data1.A0_, data2.A0_)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RecursiveNoStackResult) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RecursiveNoStackResult)
	return ok && _this.Equals(typed)
}

func Type_RecursiveNoStackResult_() _dafny.TypeDescriptor {
	return type_RecursiveNoStackResult_{}
}

type type_RecursiveNoStackResult_ struct {
}

func (_this type_RecursiveNoStackResult_) Default() interface{} {
	return Companion_RecursiveNoStackResult_.Default()
}

func (_this type_RecursiveNoStackResult_) String() string {
	return "Std_Parsers_StringParsers.RecursiveNoStackResult"
}
func (_this RecursiveNoStackResult) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RecursiveNoStackResult{}

// End of datatype RecursiveNoStackResult

// Definition of datatype RecursiveDef
type RecursiveDef struct {
	Data_RecursiveDef_
}

func (_this RecursiveDef) Get_() Data_RecursiveDef_ {
	return _this.Data_RecursiveDef_
}

type Data_RecursiveDef_ interface {
	isRecursiveDef()
}

type CompanionStruct_RecursiveDef_ struct {
}

var Companion_RecursiveDef_ = CompanionStruct_RecursiveDef_{}

type RecursiveDef_RecursiveDef struct {
	Order      _dafny.Int
	Definition func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult
}

func (RecursiveDef_RecursiveDef) isRecursiveDef() {}

func (CompanionStruct_RecursiveDef_) Create_RecursiveDef_(Order _dafny.Int, Definition func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult) RecursiveDef {
	return RecursiveDef{RecursiveDef_RecursiveDef{Order, Definition}}
}

func (_this RecursiveDef) Is_RecursiveDef() bool {
	_, ok := _this.Get_().(RecursiveDef_RecursiveDef)
	return ok
}

func (CompanionStruct_RecursiveDef_) Default() RecursiveDef {
	return Companion_RecursiveDef_.Create_RecursiveDef_(_dafny.Zero, func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
		return func(m_Std_Collections_Seq.Slice) ParseResult { return Companion_ParseResult_.Default() }
	})
}

func (_this RecursiveDef) Dtor_order() _dafny.Int {
	return _this.Get_().(RecursiveDef_RecursiveDef).Order
}

func (_this RecursiveDef) Dtor_definition() func(func(_dafny.Sequence) func(m_Std_Collections_Seq.Slice) ParseResult) func(m_Std_Collections_Seq.Slice) ParseResult {
	return _this.Get_().(RecursiveDef_RecursiveDef).Definition
}

func (_this RecursiveDef) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case RecursiveDef_RecursiveDef:
		{
			return "StringParsers.RecursiveDef.RecursiveDef" + "(" + _dafny.String(data.Order) + ", " + _dafny.String(data.Definition) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this RecursiveDef) Equals(other RecursiveDef) bool {
	switch data1 := _this.Get_().(type) {
	case RecursiveDef_RecursiveDef:
		{
			data2, ok := other.Get_().(RecursiveDef_RecursiveDef)
			return ok && data1.Order.Cmp(data2.Order) == 0 && _dafny.AreEqual(data1.Definition, data2.Definition)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this RecursiveDef) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(RecursiveDef)
	return ok && _this.Equals(typed)
}

func Type_RecursiveDef_() _dafny.TypeDescriptor {
	return type_RecursiveDef_{}
}

type type_RecursiveDef_ struct {
}

func (_this type_RecursiveDef_) Default() interface{} {
	return Companion_RecursiveDef_.Default()
}

func (_this type_RecursiveDef_) String() string {
	return "Std_Parsers_StringParsers.RecursiveDef"
}
func (_this RecursiveDef) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = RecursiveDef{}

// End of datatype RecursiveDef
