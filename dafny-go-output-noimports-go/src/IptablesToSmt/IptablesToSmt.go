// Package IptablesToSmt
// Dafny module IptablesToSmt compiled into Go

package IptablesToSmt

import (
  os "os"
  _dafny "dafny"
  m__System "System_"
)
var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__

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
var Companion_Default___ = CompanionStruct_Default___ {
}

func (_this *Default__) Equals(other *Default__) bool {
  return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
  other, ok := x.(*Default__)
  return ok && _this.Equals(other)
}

func (*Default__) String() string {
  return "IptablesToSmt.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
  return [](*_dafny.TraitID){};
}
var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) GetUsageString() _dafny.Sequence {
  var text _dafny.Sequence = _dafny.EmptySeq
  _ = text
  text = _dafny.UnicodeSeqOfUtf8Bytes("Usage:\n  dafny run src/IptablesToSmt.dfy -- \"$(cat rules.txt)\"\n  ./scripts/iptables_to_smt.sh rules.txt\n")
  return text
}
func (_static *CompanionStruct_Default___) ConvertIptablesToSmt(input _dafny.Sequence) _dafny.Sequence {
  var smt _dafny.Sequence = _dafny.EmptySeq
  _ = smt
  var _0_lines _dafny.Sequence
  _ = _0_lines
  var _out0 _dafny.Sequence
  _ = _out0
  _out0 = Companion_Default___.SplitLines(input)
  _0_lines = _out0
  var _1_rules _dafny.Sequence
  _ = _1_rules
  var _out1 _dafny.Sequence
  _ = _out1
  _out1 = Companion_Default___.ParseRules(_0_lines)
  _1_rules = _out1
  var _out2 _dafny.Sequence
  _ = _out2
  _out2 = Companion_Default___.BuildSmtDocument(_1_rules)
  smt = _out2
  return smt
}
func (_static *CompanionStruct_Default___) ParseRules(lines _dafny.Sequence) _dafny.Sequence {
  var rules _dafny.Sequence = _dafny.EmptySeq
  _ = rules
  var _0_collected _dafny.Sequence
  _ = _0_collected
  _0_collected = _dafny.SeqOf()
  var _1_idx _dafny.Int
  _ = _1_idx
  _1_idx = _dafny.Zero
  {
    for (_1_idx).Cmp(_dafny.IntOfUint32((lines).Cardinality())) < 0 {
      {
        var _2_rawLine _dafny.Sequence
        _ = _2_rawLine
        _2_rawLine = (lines).Select((_1_idx).Uint32()).(_dafny.Sequence)
        var _3_trimmed _dafny.Sequence
        _ = _3_trimmed
        var _out0 _dafny.Sequence
        _ = _out0
        _out0 = Companion_Default___.Trim(_2_rawLine)
        _3_trimmed = _out0
        if (((_dafny.IntOfUint32((_3_trimmed).Cardinality())).Sign() == 1) && (((_3_trimmed).Select(0).(_dafny.CodePoint)) == (_dafny.CodePoint('#')))) {
          _1_idx = (_1_idx).Plus(_dafny.One)
          goto C0;
        }
        if (((_dafny.IntOfUint32((_3_trimmed).Cardinality())).Cmp(_dafny.IntOfInt64(2)) >= 0) && (_dafny.Companion_Sequence_.Equal((_3_trimmed).Subsequence(0, 2), _dafny.UnicodeSeqOfUtf8Bytes("-A")))) {
          var _4_tokens _dafny.Sequence
          _ = _4_tokens
          var _out1 _dafny.Sequence
          _ = _out1
          _out1 = Companion_Default___.SplitOnWhitespace(_3_trimmed)
          _4_tokens = _out1
          var _5_parsedRule Rule
          _ = _5_parsedRule
          var _6_parsedOk bool
          _ = _6_parsedOk
          var _out2 Rule
          _ = _out2
          var _out3 bool
          _ = _out3
          _out2, _out3 = Companion_Default___.ParseRuleTokens(_4_tokens, (_1_idx).Plus(_dafny.One), _2_rawLine)
          _5_parsedRule = _out2
          _6_parsedOk = _out3
          if (_6_parsedOk) {
            _0_collected = _dafny.Companion_Sequence_.Concatenate(_0_collected, _dafny.SeqOf(_5_parsedRule))
          }
        }
        _1_idx = (_1_idx).Plus(_dafny.One)
        goto C0;
      }
    C0:
    }
    goto L0;
  }
L0:
  rules = _0_collected
  return rules
}
func (_static *CompanionStruct_Default___) FindMatch(tokens _dafny.Sequence, flag _dafny.Sequence) FieldMatch {
  goto TAIL_CALL_START
  TAIL_CALL_START:
  if ((_dafny.IntOfUint32((tokens).Cardinality())).Cmp(_dafny.IntOfInt64(2)) < 0) {
    return Companion_FieldMatch_.Create_MatchAny_()
  } else if (_dafny.Companion_Sequence_.Equal((tokens).Select(0).(_dafny.Sequence), flag)) {
    return Companion_FieldMatch_.Create_MatchExact_((tokens).Select(1).(_dafny.Sequence))
  } else {
    var _in0 _dafny.Sequence = (tokens).Drop(1)
    _ = _in0
    var _in1 _dafny.Sequence = flag
    _ = _in1
    tokens = _in0
    flag = _in1
    goto TAIL_CALL_START
  }
}
func (_static *CompanionStruct_Default___) ValidTokens(tokens _dafny.Sequence) bool {
  goto TAIL_CALL_START
  TAIL_CALL_START:
  if ((_dafny.IntOfUint32((tokens).Cardinality())).Cmp(_dafny.One) < 0) {
    return true
  } else {
    var _0_t _dafny.Sequence = (tokens).Select(0).(_dafny.Sequence)
    _ = _0_t
    if ((((((((((((((((((((((_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("-s"))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("-d")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("-p")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--sport")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--dport")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("-j")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("-i")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("-o")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("-t")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("-m")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--to-destination")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--to-source")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--log-prefix")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--log-level")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--limit")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--limit-burst")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--ctstate")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--seconds")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--hitcount")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--name")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--dports")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--sports")))) {
      if ((_dafny.IntOfUint32((tokens).Cardinality())).Cmp(_dafny.IntOfInt64(2)) >= 0) {
        var _in0 _dafny.Sequence = (tokens).Drop(2)
        _ = _in0
        tokens = _in0
        goto TAIL_CALL_START
      } else {
        return false
      }
    } else if ((((((_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--set"))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--update")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--rcheck")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--remove")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--rsource")))) || (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("--rdest")))) {
      var _in1 _dafny.Sequence = (tokens).Drop(1)
      _ = _in1
      tokens = _in1
      goto TAIL_CALL_START
    } else if (_dafny.Companion_Sequence_.Equal(_0_t, _dafny.UnicodeSeqOfUtf8Bytes("-A"))) {
      if ((_dafny.IntOfUint32((tokens).Cardinality())).Cmp(_dafny.IntOfInt64(2)) >= 0) {
        var _in2 _dafny.Sequence = (tokens).Drop(2)
        _ = _in2
        tokens = _in2
        goto TAIL_CALL_START
      } else {
        return false
      }
    } else {
      return false
    }
  }
}
func (_static *CompanionStruct_Default___) ParseRuleTokens(tokens _dafny.Sequence, lineNumber _dafny.Int, rawLine _dafny.Sequence) (Rule, bool) {
  var rule Rule = Companion_Rule_.Default()
  _ = rule
  var ok bool = false
  _ = ok
  if ((_dafny.IntOfUint32((tokens).Cardinality())).Cmp(_dafny.IntOfInt64(3)) < 0) {
    rule = Companion_Rule_.Create_Rule_(_dafny.UnicodeSeqOfUtf8Bytes(""), Companion_FieldMatch_.Create_MatchAny_(), Companion_FieldMatch_.Create_MatchAny_(), Companion_FieldMatch_.Create_MatchAny_(), Companion_FieldMatch_.Create_MatchAny_(), Companion_FieldMatch_.Create_MatchAny_(), Companion_Target_.Create_TargetReturn_(), lineNumber, rawLine)
    ok = false
    return rule, ok
  }
  var _0_chain _dafny.Sequence
  _ = _0_chain
  _0_chain = (tokens).Select(1).(_dafny.Sequence)
  var _1_rest _dafny.Sequence
  _ = _1_rest
  _1_rest = (tokens).Drop(2)
  var _2_src FieldMatch
  _ = _2_src
  _2_src = Companion_Default___.FindMatch(_1_rest, _dafny.UnicodeSeqOfUtf8Bytes("-s"))
  var _3_dst FieldMatch
  _ = _3_dst
  _3_dst = Companion_Default___.FindMatch(_1_rest, _dafny.UnicodeSeqOfUtf8Bytes("-d"))
  var _4_proto FieldMatch
  _ = _4_proto
  _4_proto = Companion_Default___.FindMatch(_1_rest, _dafny.UnicodeSeqOfUtf8Bytes("-p"))
  var _5_srcPort FieldMatch
  _ = _5_srcPort
  _5_srcPort = Companion_Default___.FindMatch(_1_rest, _dafny.UnicodeSeqOfUtf8Bytes("--sport"))
  var _6_dstPort FieldMatch
  _ = _6_dstPort
  _6_dstPort = Companion_Default___.FindMatch(_1_rest, _dafny.UnicodeSeqOfUtf8Bytes("--dport"))
  var _7_targetMatch FieldMatch
  _ = _7_targetMatch
  _7_targetMatch = Companion_Default___.FindMatch(_1_rest, _dafny.UnicodeSeqOfUtf8Bytes("-j"))
  var _8_target Target = Companion_Target_.Default()
  _ = _8_target
  var _9_targetOk bool
  _ = _9_targetOk
  _9_targetOk = false
  if ((_7_targetMatch).Is_MatchExact()) {
    var _out0 Target
    _ = _out0
    _out0 = Companion_Default___.ParseTarget((_7_targetMatch).Dtor_value())
    _8_target = _out0
    _9_targetOk = true
  } else {
    _8_target = Companion_Target_.Create_TargetReturn_()
  }
  var _10_validFlags bool
  _ = _10_validFlags
  _10_validFlags = Companion_Default___.ValidTokens(_1_rest)
  if ((_10_validFlags) && (_9_targetOk)) {
    rule = Companion_Rule_.Create_Rule_(_0_chain, _2_src, _3_dst, _4_proto, _5_srcPort, _6_dstPort, _8_target, lineNumber, rawLine)
    ok = true
  } else {
    rule = Companion_Rule_.Create_Rule_(_0_chain, _2_src, _3_dst, _4_proto, _5_srcPort, _6_dstPort, _8_target, lineNumber, rawLine)
    ok = false
  }
  return rule, ok
}
func (_static *CompanionStruct_Default___) BuildSmtDocument(rules _dafny.Sequence) _dafny.Sequence {
  var doc _dafny.Sequence = _dafny.EmptySeq
  _ = doc
  if ((_dafny.IntOfUint32((rules).Cardinality())).Sign() == 0) {
    doc = _dafny.UnicodeSeqOfUtf8Bytes("")
    return doc
  }
  var _0_builder _dafny.Sequence
  _ = _0_builder
  _0_builder = _dafny.UnicodeSeqOfUtf8Bytes("(set-logic ALL)\n")
  _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_chain String)\n"))
  _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_src String)\n"))
  _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_dst String)\n"))
  _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_proto String)\n"))
  _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_sport String)\n"))
  _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_dport String)\n"))
  _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_action String)\n"))
  _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("(assert (distinct packet_chain \"\"))\n"))
  _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("\n"))
  var _1_i _dafny.Int
  _ = _1_i
  _1_i = _dafny.Zero
  var _2_prevNegations _dafny.Sequence
  _ = _2_prevNegations
  _2_prevNegations = _dafny.UnicodeSeqOfUtf8Bytes("")
  for (_1_i).Cmp(_dafny.IntOfUint32((rules).Cardinality())) < 0 {
    var _3_formatted _dafny.Sequence
    _ = _3_formatted
    _3_formatted = Companion_Default___.FormatRule((rules).Select((_1_i).Uint32()).(Rule), _1_i, _2_prevNegations)
    _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _3_formatted)
    _0_builder = _dafny.Companion_Sequence_.Concatenate(_0_builder, _dafny.UnicodeSeqOfUtf8Bytes("\n"))
    var _4_indexText _dafny.Sequence
    _ = _4_indexText
    _4_indexText = Companion_Default___.IntToString(_1_i)
    var _5_call _dafny.Sequence
    _ = _5_call
    _5_call = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(rule"), _4_indexText), _dafny.UnicodeSeqOfUtf8Bytes(" packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)"))
    _2_prevNegations = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2_prevNegations, _dafny.UnicodeSeqOfUtf8Bytes(" (not ")), _5_call), _dafny.UnicodeSeqOfUtf8Bytes(")"))
    _1_i = (_1_i).Plus(_dafny.One)
  }
  doc = _0_builder
  return doc
}
func (_static *CompanionStruct_Default___) IntToString(n _dafny.Int) _dafny.Sequence {
  if ((n).Sign() == 0) {
    return _dafny.UnicodeSeqOfUtf8Bytes("0")
  } else {
    var _0_value _dafny.Int = (func () _dafny.Int { if (n).Sign() == -1 { return (_dafny.Zero).Minus(n) }; return n })() 
    _ = _0_value
    var _1_digits _dafny.Sequence = Companion_Default___.DigitsToString(_0_value)
    _ = _1_digits
    if ((n).Sign() == -1) {
      return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("-"), _1_digits)
    } else {
      return _1_digits
    }
  }
}
func (_static *CompanionStruct_Default___) DigitsToString(n _dafny.Int) _dafny.Sequence {
  var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
  _ = _0___accumulator
  goto TAIL_CALL_START
  TAIL_CALL_START:
  var _1_digit _dafny.Int = (n).Modulo(_dafny.IntOfInt64(10))
  _ = _1_digit
  var _2_rest _dafny.Int = (n).DivBy(_dafny.IntOfInt64(10))
  _ = _2_rest
  var _3_dStr _dafny.Sequence = Companion_Default___.DigitToString(_1_digit)
  _ = _3_dStr
  if ((_2_rest).Sign() == 0) {
    return _dafny.Companion_Sequence_.Concatenate(_3_dStr, _0___accumulator)
  } else {
    _0___accumulator = _dafny.Companion_Sequence_.Concatenate(_3_dStr, _0___accumulator)
    var _in0 _dafny.Int = _2_rest
    _ = _in0
    n = _in0
    goto TAIL_CALL_START
  }
}
func (_static *CompanionStruct_Default___) DigitToString(d _dafny.Int) _dafny.Sequence {
  return (_dafny.UnicodeSeqOfUtf8Bytes("0123456789")).Subsequence((d).Uint32(), ((d).Plus(_dafny.One)).Uint32())
}
func (_static *CompanionStruct_Default___) BuildRuleConditions(rule Rule, prefix _dafny.Sequence) _dafny.Sequence {
  var _0_chainLiteral _dafny.Sequence = Companion_Default___.FormatStringLiteral((rule).Dtor_chain())
  _ = _0_chainLiteral
  var _1_c1 _dafny.Sequence = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(= "), prefix), _dafny.UnicodeSeqOfUtf8Bytes("chain ")), _0_chainLiteral), _dafny.UnicodeSeqOfUtf8Bytes(")")))
  _ = _1_c1
  var _2_c2 _dafny.Sequence = Companion_Default___.MaybeAppend(_1_c1, _dafny.Companion_Sequence_.Concatenate(prefix, _dafny.UnicodeSeqOfUtf8Bytes("src")), (rule).Dtor_source())
  _ = _2_c2
  var _3_c3 _dafny.Sequence = Companion_Default___.MaybeAppend(_2_c2, _dafny.Companion_Sequence_.Concatenate(prefix, _dafny.UnicodeSeqOfUtf8Bytes("dst")), (rule).Dtor_destination())
  _ = _3_c3
  var _4_c4 _dafny.Sequence = Companion_Default___.MaybeAppend(_3_c3, _dafny.Companion_Sequence_.Concatenate(prefix, _dafny.UnicodeSeqOfUtf8Bytes("proto")), (rule).Dtor_protocol())
  _ = _4_c4
  var _5_c5 _dafny.Sequence = Companion_Default___.MaybeAppend(_4_c4, _dafny.Companion_Sequence_.Concatenate(prefix, _dafny.UnicodeSeqOfUtf8Bytes("sport")), (rule).Dtor_srcPort())
  _ = _5_c5
  var _6_c6 _dafny.Sequence = Companion_Default___.MaybeAppend(_5_c5, _dafny.Companion_Sequence_.Concatenate(prefix, _dafny.UnicodeSeqOfUtf8Bytes("dport")), (rule).Dtor_dstPort())
  _ = _6_c6
  return _6_c6
}
func (_static *CompanionStruct_Default___) MaybeAppend(conds _dafny.Sequence, fieldName _dafny.Sequence, crit FieldMatch) _dafny.Sequence {
  var _source0 FieldMatch = crit
  _ = _source0
  {
    if (_source0.Is_MatchExact()) {
      var _0_value _dafny.Sequence = _source0.Get_().(FieldMatch_MatchExact).Value
      _ = _0_value
      var _1_literal _dafny.Sequence = Companion_Default___.FormatStringLiteral(_0_value)
      _ = _1_literal
      var _2_fragment _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(= "), fieldName), _dafny.UnicodeSeqOfUtf8Bytes(" ")), _1_literal), _dafny.UnicodeSeqOfUtf8Bytes(")"))
      _ = _2_fragment
      return _dafny.Companion_Sequence_.Concatenate(conds, _dafny.SeqOf(_2_fragment))
    }
  }
  {
    return conds
  }
}
func (_static *CompanionStruct_Default___) FormatTargetLiteral(target Target) _dafny.Sequence {
  return Companion_Default___.FormatStringLiteral(Companion_Default___.TargetToString(target))
}
func (_static *CompanionStruct_Default___) TargetToString(target Target) _dafny.Sequence {
  var _source0 Target = target
  _ = _source0
  {
    if (_source0.Is_TargetAccept()) {
      return _dafny.UnicodeSeqOfUtf8Bytes("ACCEPT")
    }
  }
  {
    if (_source0.Is_TargetDrop()) {
      return _dafny.UnicodeSeqOfUtf8Bytes("DROP")
    }
  }
  {
    if (_source0.Is_TargetReject()) {
      return _dafny.UnicodeSeqOfUtf8Bytes("REJECT")
    }
  }
  {
    if (_source0.Is_TargetReturn()) {
      return _dafny.UnicodeSeqOfUtf8Bytes("RETURN")
    }
  }
  {
    var _0_name _dafny.Sequence = _source0.Get_().(Target_TargetJump).Name
    _ = _0_name
    return _0_name
  }
}
func (_static *CompanionStruct_Default___) FormatStringLiteral(text _dafny.Sequence) _dafny.Sequence {
  return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("\""), Companion_Default___.EscapeForSmt(text)), _dafny.UnicodeSeqOfUtf8Bytes("\""))
}
func (_static *CompanionStruct_Default___) EscapeForSmt(text _dafny.Sequence) _dafny.Sequence {
  return Companion_Default___.EscapeForSmtRecursive(text)
}
func (_static *CompanionStruct_Default___) EscapeForSmtRecursive(text _dafny.Sequence) _dafny.Sequence {
  if ((_dafny.IntOfUint32((text).Cardinality())).Sign() == 0) {
    return _dafny.UnicodeSeqOfUtf8Bytes("")
  } else {
    var _0_ch _dafny.CodePoint = (text).Select(0).(_dafny.CodePoint)
    _ = _0_ch
    var _1_rest _dafny.Sequence = Companion_Default___.EscapeForSmtRecursive((text).Drop(1))
    _ = _1_rest
    if (((_0_ch) == (_dafny.CodePoint('"'))) || ((_0_ch) == (_dafny.CodePoint('\\')))) {
      return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("\\"), (text).Subsequence(0, 1)), _1_rest)
    } else {
      return _dafny.Companion_Sequence_.Concatenate((text).Subsequence(0, 1), _1_rest)
    }
  }
}
func (_static *CompanionStruct_Default___) FormatRule(rule Rule, index _dafny.Int, prevNegations _dafny.Sequence) _dafny.Sequence {
  var _0_def _dafny.Sequence = Companion_Default___.DefineRuleFunction(rule, index)
  _ = _0_def
  var _1_assertion _dafny.Sequence = Companion_Default___.AssertRuleLogic(rule, index, prevNegations, _dafny.UnicodeSeqOfUtf8Bytes("packet_action"))
  _ = _1_assertion
  return _dafny.Companion_Sequence_.Concatenate(_0_def, _1_assertion)
}
func (_static *CompanionStruct_Default___) DefineRuleFunction(rule Rule, index _dafny.Int) _dafny.Sequence {
  var _0_indexText _dafny.Sequence = Companion_Default___.IntToString(index)
  _ = _0_indexText
  var _1_lineText _dafny.Sequence = Companion_Default___.IntToString((rule).Dtor_lineNumber())
  _ = _1_lineText
  var _2_header _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("; Rule "), _0_indexText), _dafny.UnicodeSeqOfUtf8Bytes(" (line ")), _1_lineText), _dafny.UnicodeSeqOfUtf8Bytes("): ")), (rule).Dtor_original()), _dafny.UnicodeSeqOfUtf8Bytes("\n"))
  _ = _2_header
  var _3_def _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(define-fun rule"), _0_indexText), _dafny.UnicodeSeqOfUtf8Bytes(" ((pkt_chain String) (pkt_src String) (pkt_dst String) (pkt_proto String) (pkt_sport String) (pkt_dport String)) Bool\n"))
  _ = _3_def
  var _4_conditions _dafny.Sequence = Companion_Default___.BuildRuleConditions(rule, _dafny.UnicodeSeqOfUtf8Bytes("pkt_"))
  _ = _4_conditions
  var _5_condBuilder _dafny.Sequence = (func () _dafny.Sequence { if (_dafny.IntOfUint32((_4_conditions).Cardinality())).Sign() == 0 { return _dafny.UnicodeSeqOfUtf8Bytes("  true\n") }; return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("  (and\n"), Companion_Default___.JoinLinesIndent(_4_conditions)), _dafny.UnicodeSeqOfUtf8Bytes("  )\n")) })() 
  _ = _5_condBuilder
  var _6_closing _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes(")\n")
  _ = _6_closing
  return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2_header, _3_def), _5_condBuilder), _6_closing)
}
func (_static *CompanionStruct_Default___) AssertRuleLogic(rule Rule, index _dafny.Int, prevNegations _dafny.Sequence, actionVar _dafny.Sequence) _dafny.Sequence {
  var _0_indexText _dafny.Sequence = Companion_Default___.IntToString(index)
  _ = _0_indexText
  var _1_targetLiteral _dafny.Sequence = Companion_Default___.FormatTargetLiteral((rule).Dtor_target())
  _ = _1_targetLiteral
  var _2_ruleCall _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(rule"), _0_indexText), _dafny.UnicodeSeqOfUtf8Bytes(" packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)"))
  _ = _2_ruleCall
  var _3_premise _dafny.Sequence = (func () _dafny.Sequence { if _dafny.Companion_Sequence_.Equal(prevNegations, _dafny.UnicodeSeqOfUtf8Bytes("")) { return _2_ruleCall }; return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(and "), _2_ruleCall), prevNegations), _dafny.UnicodeSeqOfUtf8Bytes(")")) })() 
  _ = _3_premise
  return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(assert (=> "), _3_premise), _dafny.UnicodeSeqOfUtf8Bytes("\n")), _dafny.UnicodeSeqOfUtf8Bytes("            (= ")), actionVar), _dafny.UnicodeSeqOfUtf8Bytes(" ")), _1_targetLiteral), _dafny.UnicodeSeqOfUtf8Bytes(")))\n"))
}
func (_static *CompanionStruct_Default___) JoinLinesIndent(lines _dafny.Sequence) _dafny.Sequence {
  var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
  _ = _0___accumulator
  goto TAIL_CALL_START
  TAIL_CALL_START:
  if ((_dafny.IntOfUint32((lines).Cardinality())).Sign() == 0) {
    return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.UnicodeSeqOfUtf8Bytes(""))
  } else {
    _0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("    "), (lines).Select(0).(_dafny.Sequence)), _dafny.UnicodeSeqOfUtf8Bytes("\n")))
    var _in0 _dafny.Sequence = (lines).Drop(1)
    _ = _in0
    lines = _in0
    goto TAIL_CALL_START
  }
}
func (_static *CompanionStruct_Default___) ToUp(ch _dafny.CodePoint) _dafny.CodePoint {
  if (((_dafny.CodePoint('a')) <= (ch)) && ((ch) <= (_dafny.CodePoint('z')))) {
    return _dafny.CodePoint((((_dafny.IntOfInt32(rune(ch))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('a'))))).Plus(_dafny.IntOfInt32(rune(_dafny.CodePoint('A'))))).Int32())
  } else {
    return ch
  }
}
func (_static *CompanionStruct_Default___) EqualsIgnoreCase(s1 _dafny.Sequence, s2 _dafny.Sequence) bool {
  return ((_dafny.IntOfUint32((s1).Cardinality())).Cmp(_dafny.IntOfUint32((s2).Cardinality())) == 0) && (_dafny.Quantifier(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfUint32((s1).Cardinality())), true, func (_forall_var_0 _dafny.Int) bool {
    var _0_i _dafny.Int
    _0_i = interface{}(_forall_var_0).(_dafny.Int)
    return !(((_0_i).Sign() != -1) && ((_0_i).Cmp(_dafny.IntOfUint32((s1).Cardinality())) < 0)) || ((Companion_Default___.ToUp((s1).Select((_0_i).Uint32()).(_dafny.CodePoint))) == (Companion_Default___.ToUp((s2).Select((_0_i).Uint32()).(_dafny.CodePoint))))
  }))
}
func (_static *CompanionStruct_Default___) MatchesTarget(target Target, raw _dafny.Sequence) bool {
  var _source0 Target = target
  _ = _source0
  {
    if (_source0.Is_TargetAccept()) {
      return Companion_Default___.EqualsIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("ACCEPT"))
    }
  }
  {
    if (_source0.Is_TargetDrop()) {
      return Companion_Default___.EqualsIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("DROP"))
    }
  }
  {
    if (_source0.Is_TargetReject()) {
      return Companion_Default___.EqualsIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("REJECT"))
    }
  }
  {
    if (_source0.Is_TargetReturn()) {
      return Companion_Default___.EqualsIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("RETURN"))
    }
  }
  {
    var _0_name _dafny.Sequence = _source0.Get_().(Target_TargetJump).Name
    _ = _0_name
    return ((((_dafny.Companion_Sequence_.Equal(_0_name, raw)) && (!(Companion_Default___.EqualsIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("ACCEPT"))))) && (!(Companion_Default___.EqualsIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("DROP"))))) && (!(Companion_Default___.EqualsIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("REJECT"))))) && (!(Companion_Default___.EqualsIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("RETURN"))))
  }
}
func (_static *CompanionStruct_Default___) ParseTarget(raw _dafny.Sequence) Target {
  var target Target = Companion_Target_.Default()
  _ = target
  var _0_isAccept bool
  _ = _0_isAccept
  var _out0 bool
  _ = _out0
  _out0 = Companion_Default___.StringsEqualIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("ACCEPT"))
  _0_isAccept = _out0
  if (_0_isAccept) {
    target = Companion_Target_.Create_TargetAccept_()
    return target
  }
  var _1_isDrop bool
  _ = _1_isDrop
  var _out1 bool
  _ = _out1
  _out1 = Companion_Default___.StringsEqualIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("DROP"))
  _1_isDrop = _out1
  if (_1_isDrop) {
    target = Companion_Target_.Create_TargetDrop_()
    return target
  }
  var _2_isReject bool
  _ = _2_isReject
  var _out2 bool
  _ = _out2
  _out2 = Companion_Default___.StringsEqualIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("REJECT"))
  _2_isReject = _out2
  if (_2_isReject) {
    target = Companion_Target_.Create_TargetReject_()
    return target
  }
  var _3_isReturn bool
  _ = _3_isReturn
  var _out3 bool
  _ = _out3
  _out3 = Companion_Default___.StringsEqualIgnoreCase(raw, _dafny.UnicodeSeqOfUtf8Bytes("RETURN"))
  _3_isReturn = _out3
  if (_3_isReturn) {
    target = Companion_Target_.Create_TargetReturn_()
    return target
  }
  target = Companion_Target_.Create_TargetJump_(raw)
  return target
}
func (_static *CompanionStruct_Default___) StringsEqualIgnoreCase(left _dafny.Sequence, right _dafny.Sequence) bool {
  var eq bool = false
  _ = eq
  if ((_dafny.IntOfUint32((left).Cardinality())).Cmp(_dafny.IntOfUint32((right).Cardinality())) != 0) {
    eq = false
    return eq
  }
  var _0_i _dafny.Int
  _ = _0_i
  _0_i = _dafny.Zero
  eq = true
  for (_0_i).Cmp(_dafny.IntOfUint32((left).Cardinality())) < 0 {
    var _1_leftUp _dafny.CodePoint
    _ = _1_leftUp
    var _out0 _dafny.CodePoint
    _ = _out0
    _out0 = Companion_Default___.ToUpperChar((left).Select((_0_i).Uint32()).(_dafny.CodePoint))
    _1_leftUp = _out0
    var _2_rightUp _dafny.CodePoint
    _ = _2_rightUp
    var _out1 _dafny.CodePoint
    _ = _out1
    _out1 = Companion_Default___.ToUpperChar((right).Select((_0_i).Uint32()).(_dafny.CodePoint))
    _2_rightUp = _out1
    if ((_1_leftUp) != (_2_rightUp)/* dircomp */) {
      eq = false
      return eq
    }
    _0_i = (_0_i).Plus(_dafny.One)
  }
  return eq
}
func (_static *CompanionStruct_Default___) ToUpperChar(ch _dafny.CodePoint) _dafny.CodePoint {
  var upper _dafny.CodePoint = _dafny.CodePoint('D')
  _ = upper
  if (((_dafny.CodePoint('a')) <= (ch)) && ((ch) <= (_dafny.CodePoint('z')))) {
    upper = _dafny.CodePoint((((_dafny.IntOfInt32(rune(ch))).Minus(_dafny.IntOfInt32(rune(_dafny.CodePoint('a'))))).Plus(_dafny.IntOfInt32(rune(_dafny.CodePoint('A'))))).Int32())
  } else {
    upper = ch
  }
  return upper
}
func (_static *CompanionStruct_Default___) Join(parts _dafny.Sequence, delim _dafny.Sequence) _dafny.Sequence {
  var _0___accumulator _dafny.Sequence = _dafny.SeqOf()
  _ = _0___accumulator
  goto TAIL_CALL_START
  TAIL_CALL_START:
  if ((_dafny.IntOfUint32((parts).Cardinality())).Sign() == 0) {
    return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.UnicodeSeqOfUtf8Bytes(""))
  } else if ((_dafny.IntOfUint32((parts).Cardinality())).Cmp(_dafny.One) == 0) {
    return _dafny.Companion_Sequence_.Concatenate(_0___accumulator, (parts).Select(0).(_dafny.Sequence))
  } else {
    _0___accumulator = _dafny.Companion_Sequence_.Concatenate(_0___accumulator, _dafny.Companion_Sequence_.Concatenate((parts).Select(0).(_dafny.Sequence), delim))
    var _in0 _dafny.Sequence = (parts).Drop(1)
    _ = _in0
    var _in1 _dafny.Sequence = delim
    _ = _in1
    parts = _in0
    delim = _in1
    goto TAIL_CALL_START
  }
}
func (_static *CompanionStruct_Default___) SplitLines(text _dafny.Sequence) _dafny.Sequence {
  var lines _dafny.Sequence = _dafny.EmptySeq
  _ = lines
  if ((_dafny.IntOfUint32((text).Cardinality())).Sign() == 0) {
    lines = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes(""))
    return lines
  }
  var _0_segments _dafny.Sequence
  _ = _0_segments
  _0_segments = _dafny.SeqOf()
  var _1_start _dafny.Int
  _ = _1_start
  _1_start = _dafny.Zero
  var _2_i _dafny.Int
  _ = _2_i
  _2_i = _dafny.Zero
  for (_2_i).Cmp(_dafny.IntOfUint32((text).Cardinality())) < 0 {
    var _3_ch _dafny.CodePoint
    _ = _3_ch
    _3_ch = (text).Select((_2_i).Uint32()).(_dafny.CodePoint)
    if ((_3_ch) == (_dafny.CodePoint('\n'))) {
      var _4_part _dafny.Sequence
      _ = _4_part
      _4_part = (text).Subsequence((_1_start).Uint32(), (_2_i).Uint32())
      _0_segments = _dafny.Companion_Sequence_.Concatenate(_0_segments, _dafny.SeqOf(_4_part))
      _1_start = (_2_i).Plus(_dafny.One)
    }
    _2_i = (_2_i).Plus(_dafny.One)
  }
  var _5_lastPart _dafny.Sequence
  _ = _5_lastPart
  _5_lastPart = (text).Subsequence((_1_start).Uint32(), (_dafny.IntOfUint32((text).Cardinality())).Uint32())
  _0_segments = _dafny.Companion_Sequence_.Concatenate(_0_segments, _dafny.SeqOf(_5_lastPart))
  lines = _0_segments
  return lines
}
func (_static *CompanionStruct_Default___) SplitOnWhitespace(text _dafny.Sequence) _dafny.Sequence {
  var tokens _dafny.Sequence = _dafny.EmptySeq
  _ = tokens
  var _0_parts _dafny.Sequence
  _ = _0_parts
  _0_parts = _dafny.SeqOf()
  var _1_i _dafny.Int
  _ = _1_i
  _1_i = _dafny.Zero
  {
    for (_1_i).Cmp(_dafny.IntOfUint32((text).Cardinality())) < 0 {
      {
        for ((_1_i).Cmp(_dafny.IntOfUint32((text).Cardinality())) < 0) && (Companion_Default___.IsWhitespace((text).Select((_1_i).Uint32()).(_dafny.CodePoint))) {
          _1_i = (_1_i).Plus(_dafny.One)
        }
        if ((_1_i).Cmp(_dafny.IntOfUint32((text).Cardinality())) >= 0) {
          goto L1
        }
        if (((text).Select((_1_i).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('"'))) {
          var _2_start _dafny.Int
          _ = _2_start
          _2_start = (_1_i).Plus(_dafny.One)
          _1_i = (_1_i).Plus(_dafny.One)
          for ((_1_i).Cmp(_dafny.IntOfUint32((text).Cardinality())) < 0) && (((text).Select((_1_i).Uint32()).(_dafny.CodePoint)) != (_dafny.CodePoint('"'))/* dircomp */) {
            if ((((text).Select((_1_i).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('\\'))) && (((_1_i).Plus(_dafny.One)).Cmp(_dafny.IntOfUint32((text).Cardinality())) < 0)) {
              _1_i = (_1_i).Plus(_dafny.IntOfInt64(2))
            } else {
              _1_i = (_1_i).Plus(_dafny.One)
            }
          }
          var _3_value _dafny.Sequence
          _ = _3_value
          _3_value = (text).Subsequence((_2_start).Uint32(), (_1_i).Uint32())
          if ((_dafny.IntOfUint32((_3_value).Cardinality())).Sign() == 1) {
            _0_parts = _dafny.Companion_Sequence_.Concatenate(_0_parts, _dafny.SeqOf(_3_value))
          }
          if (((_1_i).Cmp(_dafny.IntOfUint32((text).Cardinality())) < 0) && (((text).Select((_1_i).Uint32()).(_dafny.CodePoint)) == (_dafny.CodePoint('"')))) {
            _1_i = (_1_i).Plus(_dafny.One)
          }
        } else {
          var _4_start2 _dafny.Int
          _ = _4_start2
          _4_start2 = _1_i
          for ((_1_i).Cmp(_dafny.IntOfUint32((text).Cardinality())) < 0) && (!(Companion_Default___.IsWhitespace((text).Select((_1_i).Uint32()).(_dafny.CodePoint)))) {
            _1_i = (_1_i).Plus(_dafny.One)
          }
          var _5_token _dafny.Sequence
          _ = _5_token
          _5_token = (text).Subsequence((_4_start2).Uint32(), (_1_i).Uint32())
          if ((_dafny.IntOfUint32((_5_token).Cardinality())).Sign() == 1) {
            _0_parts = _dafny.Companion_Sequence_.Concatenate(_0_parts, _dafny.SeqOf(_5_token))
          }
        }
        goto C1;
      }
    C1:
    }
    goto L1;
  }
L1:
  tokens = _0_parts
  return tokens
}
func (_static *CompanionStruct_Default___) Trim(text _dafny.Sequence) _dafny.Sequence {
  var trimmed _dafny.Sequence = _dafny.EmptySeq
  _ = trimmed
  if ((_dafny.IntOfUint32((text).Cardinality())).Sign() == 0) {
    trimmed = text
    return trimmed
  }
  var _0_start _dafny.Int
  _ = _0_start
  _0_start = _dafny.Zero
  var _1_stop _dafny.Int
  _ = _1_stop
  _1_stop = _dafny.IntOfUint32((text).Cardinality())
  for ((_0_start).Cmp(_1_stop) < 0) && (Companion_Default___.IsWhitespace((text).Select((_0_start).Uint32()).(_dafny.CodePoint))) {
    _0_start = (_0_start).Plus(_dafny.One)
  }
  for ((_1_stop).Cmp(_0_start) > 0) && (Companion_Default___.IsWhitespace((text).Select(((_1_stop).Minus(_dafny.One)).Uint32()).(_dafny.CodePoint))) {
    _1_stop = (_1_stop).Minus(_dafny.One)
  }
  trimmed = (text).Subsequence((_0_start).Uint32(), (_1_stop).Uint32())
  return trimmed
}
func (_static *CompanionStruct_Default___) IsWhitespace(ch _dafny.CodePoint) bool {
  return ((((ch) == (_dafny.CodePoint(' '))) || ((ch) == (_dafny.CodePoint('\t')))) || ((ch) == (_dafny.CodePoint('\r')))) || ((ch) == (_dafny.CodePoint('\n')))
}
func (_static *CompanionStruct_Default___) IsSubstring(haystack _dafny.Sequence, needle _dafny.Sequence) bool {
  goto TAIL_CALL_START
  TAIL_CALL_START:
  if ((_dafny.IntOfUint32((needle).Cardinality())).Cmp(_dafny.IntOfUint32((haystack).Cardinality())) > 0) {
    return false
  } else if ((_dafny.IntOfUint32((needle).Cardinality())).Sign() == 0) {
    return true
  } else if (_dafny.Companion_Sequence_.Equal((haystack).Subsequence(0, (_dafny.IntOfUint32((needle).Cardinality())).Uint32()), needle)) {
    return true
  } else {
    var _in0 _dafny.Sequence = (haystack).Drop(1)
    _ = _in0
    var _in1 _dafny.Sequence = needle
    _ = _in1
    haystack = _in0
    needle = _in1
    goto TAIL_CALL_START
  }
}
// End of class Default__

// Definition of datatype Target
type Target struct {
  Data_Target_
}

func (_this Target) Get_() Data_Target_ {
  return _this.Data_Target_
}

type Data_Target_ interface {
  isTarget()
}

type CompanionStruct_Target_ struct {
}
var Companion_Target_ = CompanionStruct_Target_ {
}

type Target_TargetAccept struct {
}

func (Target_TargetAccept) isTarget() {}

func (CompanionStruct_Target_) Create_TargetAccept_() Target {
  return Target{Target_TargetAccept{}}
}

func (_this Target) Is_TargetAccept() bool {
  _, ok := _this.Get_().(Target_TargetAccept)
  return ok
}

type Target_TargetDrop struct {
}

func (Target_TargetDrop) isTarget() {}

func (CompanionStruct_Target_) Create_TargetDrop_() Target {
  return Target{Target_TargetDrop{}}
}

func (_this Target) Is_TargetDrop() bool {
  _, ok := _this.Get_().(Target_TargetDrop)
  return ok
}

type Target_TargetReject struct {
}

func (Target_TargetReject) isTarget() {}

func (CompanionStruct_Target_) Create_TargetReject_() Target {
  return Target{Target_TargetReject{}}
}

func (_this Target) Is_TargetReject() bool {
  _, ok := _this.Get_().(Target_TargetReject)
  return ok
}

type Target_TargetReturn struct {
}

func (Target_TargetReturn) isTarget() {}

func (CompanionStruct_Target_) Create_TargetReturn_() Target {
  return Target{Target_TargetReturn{}}
}

func (_this Target) Is_TargetReturn() bool {
  _, ok := _this.Get_().(Target_TargetReturn)
  return ok
}

type Target_TargetJump struct {
  Name _dafny.Sequence
}

func (Target_TargetJump) isTarget() {}

func (CompanionStruct_Target_) Create_TargetJump_(Name _dafny.Sequence) Target {
  return Target{Target_TargetJump{Name}}
}

func (_this Target) Is_TargetJump() bool {
  _, ok := _this.Get_().(Target_TargetJump)
  return ok
}

func (CompanionStruct_Target_) Default() Target {
  return Companion_Target_.Create_TargetAccept_()
}

func (_this Target) Dtor_name() _dafny.Sequence {
  return _this.Get_().(Target_TargetJump).Name
}

func (_this Target) String() string {
  switch data := _this.Get_().(type) {
    case nil: return "null"
    case Target_TargetAccept: {
      return "IptablesToSmt.Target.TargetAccept"
    }
    case Target_TargetDrop: {
      return "IptablesToSmt.Target.TargetDrop"
    }
    case Target_TargetReject: {
      return "IptablesToSmt.Target.TargetReject"
    }
    case Target_TargetReturn: {
      return "IptablesToSmt.Target.TargetReturn"
    }
    case Target_TargetJump: {
      return "IptablesToSmt.Target.TargetJump" + "(" + data.Name.VerbatimString(true) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this Target) Equals(other Target) bool {
  switch data1 := _this.Get_().(type) {
    case Target_TargetAccept: {
      _, ok := other.Get_().(Target_TargetAccept)
      return ok
    }
    case Target_TargetDrop: {
      _, ok := other.Get_().(Target_TargetDrop)
      return ok
    }
    case Target_TargetReject: {
      _, ok := other.Get_().(Target_TargetReject)
      return ok
    }
    case Target_TargetReturn: {
      _, ok := other.Get_().(Target_TargetReturn)
      return ok
    }
    case Target_TargetJump: {
      data2, ok := other.Get_().(Target_TargetJump)
      return ok && data1.Name.Equals(data2.Name)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this Target) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(Target)
  return ok && _this.Equals(typed)
}

func Type_Target_() _dafny.TypeDescriptor {
  return type_Target_{}
}

type type_Target_ struct {
}

func (_this type_Target_) Default() interface{} {
  return Companion_Target_.Default();
}

func (_this type_Target_) String() string {
  return "IptablesToSmt.Target"
}
func (_this Target) ParentTraits_() []*_dafny.TraitID {
  return [](*_dafny.TraitID){};
}
var _ _dafny.TraitOffspring = Target{}

// End of datatype Target

// Definition of datatype FieldMatch
type FieldMatch struct {
  Data_FieldMatch_
}

func (_this FieldMatch) Get_() Data_FieldMatch_ {
  return _this.Data_FieldMatch_
}

type Data_FieldMatch_ interface {
  isFieldMatch()
}

type CompanionStruct_FieldMatch_ struct {
}
var Companion_FieldMatch_ = CompanionStruct_FieldMatch_ {
}

type FieldMatch_MatchAny struct {
}

func (FieldMatch_MatchAny) isFieldMatch() {}

func (CompanionStruct_FieldMatch_) Create_MatchAny_() FieldMatch {
  return FieldMatch{FieldMatch_MatchAny{}}
}

func (_this FieldMatch) Is_MatchAny() bool {
  _, ok := _this.Get_().(FieldMatch_MatchAny)
  return ok
}

type FieldMatch_MatchExact struct {
  Value _dafny.Sequence
}

func (FieldMatch_MatchExact) isFieldMatch() {}

func (CompanionStruct_FieldMatch_) Create_MatchExact_(Value _dafny.Sequence) FieldMatch {
  return FieldMatch{FieldMatch_MatchExact{Value}}
}

func (_this FieldMatch) Is_MatchExact() bool {
  _, ok := _this.Get_().(FieldMatch_MatchExact)
  return ok
}

func (CompanionStruct_FieldMatch_) Default() FieldMatch {
  return Companion_FieldMatch_.Create_MatchAny_()
}

func (_this FieldMatch) Dtor_value() _dafny.Sequence {
  return _this.Get_().(FieldMatch_MatchExact).Value
}

func (_this FieldMatch) String() string {
  switch data := _this.Get_().(type) {
    case nil: return "null"
    case FieldMatch_MatchAny: {
      return "IptablesToSmt.FieldMatch.MatchAny"
    }
    case FieldMatch_MatchExact: {
      return "IptablesToSmt.FieldMatch.MatchExact" + "(" + data.Value.VerbatimString(true) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this FieldMatch) Equals(other FieldMatch) bool {
  switch data1 := _this.Get_().(type) {
    case FieldMatch_MatchAny: {
      _, ok := other.Get_().(FieldMatch_MatchAny)
      return ok
    }
    case FieldMatch_MatchExact: {
      data2, ok := other.Get_().(FieldMatch_MatchExact)
      return ok && data1.Value.Equals(data2.Value)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this FieldMatch) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(FieldMatch)
  return ok && _this.Equals(typed)
}

func Type_FieldMatch_() _dafny.TypeDescriptor {
  return type_FieldMatch_{}
}

type type_FieldMatch_ struct {
}

func (_this type_FieldMatch_) Default() interface{} {
  return Companion_FieldMatch_.Default();
}

func (_this type_FieldMatch_) String() string {
  return "IptablesToSmt.FieldMatch"
}
func (_this FieldMatch) ParentTraits_() []*_dafny.TraitID {
  return [](*_dafny.TraitID){};
}
var _ _dafny.TraitOffspring = FieldMatch{}

// End of datatype FieldMatch

// Definition of datatype Rule
type Rule struct {
  Data_Rule_
}

func (_this Rule) Get_() Data_Rule_ {
  return _this.Data_Rule_
}

type Data_Rule_ interface {
  isRule()
}

type CompanionStruct_Rule_ struct {
}
var Companion_Rule_ = CompanionStruct_Rule_ {
}

type Rule_Rule struct {
  Chain _dafny.Sequence
  Source FieldMatch
  Destination FieldMatch
  Protocol FieldMatch
  SrcPort FieldMatch
  DstPort FieldMatch
  Target Target
  LineNumber _dafny.Int
  Original _dafny.Sequence
}

func (Rule_Rule) isRule() {}

func (CompanionStruct_Rule_) Create_Rule_(Chain _dafny.Sequence, Source FieldMatch, Destination FieldMatch, Protocol FieldMatch, SrcPort FieldMatch, DstPort FieldMatch, Target Target, LineNumber _dafny.Int, Original _dafny.Sequence) Rule {
  return Rule{Rule_Rule{Chain, Source, Destination, Protocol, SrcPort, DstPort, Target, LineNumber, Original}}
}

func (_this Rule) Is_Rule() bool {
  _, ok := _this.Get_().(Rule_Rule)
  return ok
}

func (CompanionStruct_Rule_) Default() Rule {
  return Companion_Rule_.Create_Rule_(_dafny.EmptySeq, Companion_FieldMatch_.Default(), Companion_FieldMatch_.Default(), Companion_FieldMatch_.Default(), Companion_FieldMatch_.Default(), Companion_FieldMatch_.Default(), Companion_Target_.Default(), _dafny.Zero, _dafny.EmptySeq)
}

func (_this Rule) Dtor_chain() _dafny.Sequence {
  return _this.Get_().(Rule_Rule).Chain
}

func (_this Rule) Dtor_source() FieldMatch {
  return _this.Get_().(Rule_Rule).Source
}

func (_this Rule) Dtor_destination() FieldMatch {
  return _this.Get_().(Rule_Rule).Destination
}

func (_this Rule) Dtor_protocol() FieldMatch {
  return _this.Get_().(Rule_Rule).Protocol
}

func (_this Rule) Dtor_srcPort() FieldMatch {
  return _this.Get_().(Rule_Rule).SrcPort
}

func (_this Rule) Dtor_dstPort() FieldMatch {
  return _this.Get_().(Rule_Rule).DstPort
}

func (_this Rule) Dtor_target() Target {
  return _this.Get_().(Rule_Rule).Target
}

func (_this Rule) Dtor_lineNumber() _dafny.Int {
  return _this.Get_().(Rule_Rule).LineNumber
}

func (_this Rule) Dtor_original() _dafny.Sequence {
  return _this.Get_().(Rule_Rule).Original
}

func (_this Rule) String() string {
  switch data := _this.Get_().(type) {
    case nil: return "null"
    case Rule_Rule: {
      return "IptablesToSmt.Rule.Rule" + "(" + data.Chain.VerbatimString(true) + ", " + _dafny.String(data.Source) + ", " + _dafny.String(data.Destination) + ", " + _dafny.String(data.Protocol) + ", " + _dafny.String(data.SrcPort) + ", " + _dafny.String(data.DstPort) + ", " + _dafny.String(data.Target) + ", " + _dafny.String(data.LineNumber) + ", " + data.Original.VerbatimString(true) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this Rule) Equals(other Rule) bool {
  switch data1 := _this.Get_().(type) {
    case Rule_Rule: {
      data2, ok := other.Get_().(Rule_Rule)
      return ok && data1.Chain.Equals(data2.Chain) && data1.Source.Equals(data2.Source) && data1.Destination.Equals(data2.Destination) && data1.Protocol.Equals(data2.Protocol) && data1.SrcPort.Equals(data2.SrcPort) && data1.DstPort.Equals(data2.DstPort) && data1.Target.Equals(data2.Target) && data1.LineNumber.Cmp(data2.LineNumber) == 0 && data1.Original.Equals(data2.Original)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this Rule) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(Rule)
  return ok && _this.Equals(typed)
}

func Type_Rule_() _dafny.TypeDescriptor {
  return type_Rule_{}
}

type type_Rule_ struct {
}

func (_this type_Rule_) Default() interface{} {
  return Companion_Rule_.Default();
}

func (_this type_Rule_) String() string {
  return "IptablesToSmt.Rule"
}
func (_this Rule) ParentTraits_() []*_dafny.TraitID {
  return [](*_dafny.TraitID){};
}
var _ _dafny.TraitOffspring = Rule{}

// End of datatype Rule

// Definition of datatype SmtResult
type SmtResult struct {
  Data_SmtResult_
}

func (_this SmtResult) Get_() Data_SmtResult_ {
  return _this.Data_SmtResult_
}

type Data_SmtResult_ interface {
  isSmtResult()
}

type CompanionStruct_SmtResult_ struct {
}
var Companion_SmtResult_ = CompanionStruct_SmtResult_ {
}

type SmtResult_Sat struct {
}

func (SmtResult_Sat) isSmtResult() {}

func (CompanionStruct_SmtResult_) Create_Sat_() SmtResult {
  return SmtResult{SmtResult_Sat{}}
}

func (_this SmtResult) Is_Sat() bool {
  _, ok := _this.Get_().(SmtResult_Sat)
  return ok
}

type SmtResult_Unsat struct {
}

func (SmtResult_Unsat) isSmtResult() {}

func (CompanionStruct_SmtResult_) Create_Unsat_() SmtResult {
  return SmtResult{SmtResult_Unsat{}}
}

func (_this SmtResult) Is_Unsat() bool {
  _, ok := _this.Get_().(SmtResult_Unsat)
  return ok
}

type SmtResult_Unknown struct {
}

func (SmtResult_Unknown) isSmtResult() {}

func (CompanionStruct_SmtResult_) Create_Unknown_() SmtResult {
  return SmtResult{SmtResult_Unknown{}}
}

func (_this SmtResult) Is_Unknown() bool {
  _, ok := _this.Get_().(SmtResult_Unknown)
  return ok
}

type SmtResult_Error struct {
  Msg _dafny.Sequence
}

func (SmtResult_Error) isSmtResult() {}

func (CompanionStruct_SmtResult_) Create_Error_(Msg _dafny.Sequence) SmtResult {
  return SmtResult{SmtResult_Error{Msg}}
}

func (_this SmtResult) Is_Error() bool {
  _, ok := _this.Get_().(SmtResult_Error)
  return ok
}

func (CompanionStruct_SmtResult_) Default() SmtResult {
  return Companion_SmtResult_.Create_Sat_()
}

func (_this SmtResult) Dtor_msg() _dafny.Sequence {
  return _this.Get_().(SmtResult_Error).Msg
}

func (_this SmtResult) String() string {
  switch data := _this.Get_().(type) {
    case nil: return "null"
    case SmtResult_Sat: {
      return "IptablesToSmt.SmtResult.Sat"
    }
    case SmtResult_Unsat: {
      return "IptablesToSmt.SmtResult.Unsat"
    }
    case SmtResult_Unknown: {
      return "IptablesToSmt.SmtResult.Unknown"
    }
    case SmtResult_Error: {
      return "IptablesToSmt.SmtResult.Error" + "(" + data.Msg.VerbatimString(true) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this SmtResult) Equals(other SmtResult) bool {
  switch data1 := _this.Get_().(type) {
    case SmtResult_Sat: {
      _, ok := other.Get_().(SmtResult_Sat)
      return ok
    }
    case SmtResult_Unsat: {
      _, ok := other.Get_().(SmtResult_Unsat)
      return ok
    }
    case SmtResult_Unknown: {
      _, ok := other.Get_().(SmtResult_Unknown)
      return ok
    }
    case SmtResult_Error: {
      data2, ok := other.Get_().(SmtResult_Error)
      return ok && data1.Msg.Equals(data2.Msg)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this SmtResult) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(SmtResult)
  return ok && _this.Equals(typed)
}

func Type_SmtResult_() _dafny.TypeDescriptor {
  return type_SmtResult_{}
}

type type_SmtResult_ struct {
}

func (_this type_SmtResult_) Default() interface{} {
  return Companion_SmtResult_.Default();
}

func (_this type_SmtResult_) String() string {
  return "IptablesToSmt.SmtResult"
}
func (_this SmtResult) ParentTraits_() []*_dafny.TraitID {
  return [](*_dafny.TraitID){};
}
var _ _dafny.TraitOffspring = SmtResult{}

// End of datatype SmtResult

// Definition of trait SmtSolver
type SmtSolver interface {
  String() string
  SendCommand(cmd _dafny.Sequence) bool
  CheckSat() SmtResult
  GetValue(variable _dafny.Sequence) _dafny.Sequence
}
type CompanionStruct_SmtSolver_ struct {
  TraitID_ *_dafny.TraitID
}
var Companion_SmtSolver_ = CompanionStruct_SmtSolver_ {
  TraitID_: &_dafny.TraitID{},
}
func (CompanionStruct_SmtSolver_) CastTo_(x interface{}) SmtSolver {
  var t SmtSolver
  t, _ = x.(SmtSolver)
  return t
}
// End of trait SmtSolver
