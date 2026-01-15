// Package Analysis
// Dafny module Analysis compiled into Go

package Analysis

import (
	m_IptablesToSmt "IptablesToSmt"
	m__System "System_"
	_dafny "dafny"
	os "os"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_IptablesToSmt.Dummy__

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
	return "Analysis.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) CheckEquivalence(rules1 _dafny.Sequence, rules2 _dafny.Sequence, solver m_IptablesToSmt.SmtSolver) bool {
	var areEquivalent bool = false
	_ = areEquivalent
	var _0_cmdOk bool
	_ = _0_cmdOk
	var _out0 bool
	_ = _out0
	_out0 = (solver).SendCommand(_dafny.UnicodeSeqOfUtf8Bytes("(reset)"))
	_0_cmdOk = _out0
	if !(_0_cmdOk) {
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Failed to reset solver\n").VerbatimString(false))
		areEquivalent = false
		return areEquivalent
	}
	var _out1 bool
	_ = _out1
	_out1 = (solver).SendCommand(_dafny.UnicodeSeqOfUtf8Bytes("(set-logic ALL)"))
	_0_cmdOk = _out1
	var _1_decls _dafny.Sequence
	_ = _1_decls
	_1_decls = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_chain String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_src String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_dst String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_proto String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_sport String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_dport String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_action_1 String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_action_2 String)"), _dafny.UnicodeSeqOfUtf8Bytes("(assert (distinct packet_chain \"\"))"))
	var _2_k _dafny.Int
	_ = _2_k
	_2_k = _dafny.Zero
	for (_2_k).Cmp(_dafny.IntOfUint32((_1_decls).Cardinality())) < 0 {
		var _3___v0 bool
		_ = _3___v0
		var _out2 bool
		_ = _out2
		_out2 = (solver).SendCommand((_1_decls).Select((_2_k).Uint32()).(_dafny.Sequence))
		_3___v0 = _out2
		_2_k = (_2_k).Plus(_dafny.One)
	}
	var _4_i _dafny.Int
	_ = _4_i
	_4_i = _dafny.Zero
	var _5_prevNegations1 _dafny.Sequence
	_ = _5_prevNegations1
	_5_prevNegations1 = _dafny.UnicodeSeqOfUtf8Bytes("")
	for (_4_i).Cmp(_dafny.IntOfUint32((rules1).Cardinality())) < 0 {
		var _6_rule m_IptablesToSmt.Rule
		_ = _6_rule
		_6_rule = (rules1).Select((_4_i).Uint32()).(m_IptablesToSmt.Rule)
		var _7_def _dafny.Sequence
		_ = _7_def
		_7_def = m_IptablesToSmt.Companion_Default___.DefineRuleFunction(_6_rule, _4_i)
		var _8___v1 bool
		_ = _8___v1
		var _out3 bool
		_ = _out3
		_out3 = (solver).SendCommand(_7_def)
		_8___v1 = _out3
		var _9_assertion _dafny.Sequence
		_ = _9_assertion
		_9_assertion = m_IptablesToSmt.Companion_Default___.AssertRuleLogic(_6_rule, _4_i, _5_prevNegations1, _dafny.UnicodeSeqOfUtf8Bytes("packet_action_1"))
		var _10___v2 bool
		_ = _10___v2
		var _out4 bool
		_ = _out4
		_out4 = (solver).SendCommand(_9_assertion)
		_10___v2 = _out4
		var _11_indexText _dafny.Sequence
		_ = _11_indexText
		_11_indexText = m_IptablesToSmt.Companion_Default___.IntToString(_4_i)
		var _12_call _dafny.Sequence
		_ = _12_call
		_12_call = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(rule"), _11_indexText), _dafny.UnicodeSeqOfUtf8Bytes(" packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)"))
		_5_prevNegations1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_5_prevNegations1, _dafny.UnicodeSeqOfUtf8Bytes(" (not ")), _12_call), _dafny.UnicodeSeqOfUtf8Bytes(")"))
		_4_i = (_4_i).Plus(_dafny.One)
	}
	var _13_finalNegations1 _dafny.Sequence
	_ = _13_finalNegations1
	_13_finalNegations1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(and"), _5_prevNegations1), _dafny.UnicodeSeqOfUtf8Bytes(")"))
	var _14___v3 bool
	_ = _14___v3
	var _out5 bool
	_ = _out5
	_out5 = (solver).SendCommand(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(assert (=> "), _13_finalNegations1), _dafny.UnicodeSeqOfUtf8Bytes(" (= packet_action_1 \"DEFAULT\")))")))
	_14___v3 = _out5
	var _15_j _dafny.Int
	_ = _15_j
	_15_j = _dafny.Zero
	var _16_prevNegations2 _dafny.Sequence
	_ = _16_prevNegations2
	_16_prevNegations2 = _dafny.UnicodeSeqOfUtf8Bytes("")
	var _17_offset _dafny.Int
	_ = _17_offset
	_17_offset = _dafny.IntOfUint32((rules1).Cardinality())
	for (_15_j).Cmp(_dafny.IntOfUint32((rules2).Cardinality())) < 0 {
		var _18_rule m_IptablesToSmt.Rule
		_ = _18_rule
		_18_rule = (rules2).Select((_15_j).Uint32()).(m_IptablesToSmt.Rule)
		var _19_actualIndex _dafny.Int
		_ = _19_actualIndex
		_19_actualIndex = (_17_offset).Plus(_15_j)
		var _20_def _dafny.Sequence
		_ = _20_def
		_20_def = m_IptablesToSmt.Companion_Default___.DefineRuleFunction(_18_rule, _19_actualIndex)
		var _21___v4 bool
		_ = _21___v4
		var _out6 bool
		_ = _out6
		_out6 = (solver).SendCommand(_20_def)
		_21___v4 = _out6
		var _22_assertion _dafny.Sequence
		_ = _22_assertion
		_22_assertion = m_IptablesToSmt.Companion_Default___.AssertRuleLogic(_18_rule, _19_actualIndex, _16_prevNegations2, _dafny.UnicodeSeqOfUtf8Bytes("packet_action_2"))
		var _23___v5 bool
		_ = _23___v5
		var _out7 bool
		_ = _out7
		_out7 = (solver).SendCommand(_22_assertion)
		_23___v5 = _out7
		var _24_indexText _dafny.Sequence
		_ = _24_indexText
		_24_indexText = m_IptablesToSmt.Companion_Default___.IntToString(_19_actualIndex)
		var _25_call _dafny.Sequence
		_ = _25_call
		_25_call = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(rule"), _24_indexText), _dafny.UnicodeSeqOfUtf8Bytes(" packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)"))
		_16_prevNegations2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_16_prevNegations2, _dafny.UnicodeSeqOfUtf8Bytes(" (not ")), _25_call), _dafny.UnicodeSeqOfUtf8Bytes(")"))
		_15_j = (_15_j).Plus(_dafny.One)
	}
	var _26_finalNegations2 _dafny.Sequence
	_ = _26_finalNegations2
	_26_finalNegations2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(and"), _16_prevNegations2), _dafny.UnicodeSeqOfUtf8Bytes(")"))
	var _27___v6 bool
	_ = _27___v6
	var _out8 bool
	_ = _out8
	_out8 = (solver).SendCommand(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(assert (=> "), _26_finalNegations2), _dafny.UnicodeSeqOfUtf8Bytes(" (= packet_action_2 \"DEFAULT\")))")))
	_27___v6 = _out8
	var _28___v7 bool
	_ = _28___v7
	var _out9 bool
	_ = _out9
	_out9 = (solver).SendCommand(_dafny.UnicodeSeqOfUtf8Bytes("(assert (not (= packet_action_1 packet_action_2)))"))
	_28___v7 = _out9
	var _29_result m_IptablesToSmt.SmtResult
	_ = _29_result
	var _out10 m_IptablesToSmt.SmtResult
	_ = _out10
	_out10 = (solver).CheckSat()
	_29_result = _out10
	var _source0 m_IptablesToSmt.SmtResult = _29_result
	_ = _source0
	{
		{
			if _source0.Is_Unsat() {
				areEquivalent = true
				goto Lmatch0
			}
		}
		{
			if _source0.Is_Sat() {
				areEquivalent = false
				goto Lmatch0
			}
		}
		{
			if _source0.Is_Unknown() {
				_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Warning: Solver returned unknown\n").VerbatimString(false))
				areEquivalent = false
				goto Lmatch0
			}
		}
		{
			var _30_msg _dafny.Sequence = _source0.Get_().(m_IptablesToSmt.SmtResult_Error).Msg
			_ = _30_msg
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Solver failed: ").VerbatimString(false))
			_dafny.Print(_30_msg.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			areEquivalent = false
		}
		goto Lmatch0
	}
Lmatch0:
	return areEquivalent
}
func (_static *CompanionStruct_Default___) QueryPacket(rules _dafny.Sequence, packetSpec m_IptablesToSmt.Rule, solver m_IptablesToSmt.SmtSolver) _dafny.Sequence {
	var action _dafny.Sequence = _dafny.EmptySeq
	_ = action
	var _0___v8 bool
	_ = _0___v8
	var _out0 bool
	_ = _out0
	_out0 = (solver).SendCommand(_dafny.UnicodeSeqOfUtf8Bytes("(reset)"))
	_0___v8 = _out0
	var _1___v9 bool
	_ = _1___v9
	var _out1 bool
	_ = _out1
	_out1 = (solver).SendCommand(_dafny.UnicodeSeqOfUtf8Bytes("(set-logic ALL)"))
	_1___v9 = _out1
	var _2_decls _dafny.Sequence
	_ = _2_decls
	_2_decls = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_chain String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_src String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_dst String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_proto String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_sport String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_dport String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_action String)"), _dafny.UnicodeSeqOfUtf8Bytes("(assert (distinct packet_chain \"\"))"))
	var _3_k _dafny.Int
	_ = _3_k
	_3_k = _dafny.Zero
	for (_3_k).Cmp(_dafny.IntOfUint32((_2_decls).Cardinality())) < 0 {
		var _4___v10 bool
		_ = _4___v10
		var _out2 bool
		_ = _out2
		_out2 = (solver).SendCommand((_2_decls).Select((_3_k).Uint32()).(_dafny.Sequence))
		_4___v10 = _out2
		_3_k = (_3_k).Plus(_dafny.One)
	}
	var _5_i _dafny.Int
	_ = _5_i
	_5_i = _dafny.Zero
	var _6_prevNegations _dafny.Sequence
	_ = _6_prevNegations
	_6_prevNegations = _dafny.UnicodeSeqOfUtf8Bytes("")
	for (_5_i).Cmp(_dafny.IntOfUint32((rules).Cardinality())) < 0 {
		var _7_rule m_IptablesToSmt.Rule
		_ = _7_rule
		_7_rule = (rules).Select((_5_i).Uint32()).(m_IptablesToSmt.Rule)
		var _8_def _dafny.Sequence
		_ = _8_def
		_8_def = m_IptablesToSmt.Companion_Default___.DefineRuleFunction(_7_rule, _5_i)
		var _9___v11 bool
		_ = _9___v11
		var _out3 bool
		_ = _out3
		_out3 = (solver).SendCommand(_8_def)
		_9___v11 = _out3
		var _10_assertion _dafny.Sequence
		_ = _10_assertion
		_10_assertion = m_IptablesToSmt.Companion_Default___.AssertRuleLogic(_7_rule, _5_i, _6_prevNegations, _dafny.UnicodeSeqOfUtf8Bytes("packet_action"))
		var _11___v12 bool
		_ = _11___v12
		var _out4 bool
		_ = _out4
		_out4 = (solver).SendCommand(_10_assertion)
		_11___v12 = _out4
		var _12_indexText _dafny.Sequence
		_ = _12_indexText
		_12_indexText = m_IptablesToSmt.Companion_Default___.IntToString(_5_i)
		var _13_call _dafny.Sequence
		_ = _13_call
		_13_call = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(rule"), _12_indexText), _dafny.UnicodeSeqOfUtf8Bytes(" packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)"))
		_6_prevNegations = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_6_prevNegations, _dafny.UnicodeSeqOfUtf8Bytes(" (not ")), _13_call), _dafny.UnicodeSeqOfUtf8Bytes(")"))
		_5_i = (_5_i).Plus(_dafny.One)
	}
	var _14_finalNegations _dafny.Sequence
	_ = _14_finalNegations
	_14_finalNegations = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(and"), _6_prevNegations), _dafny.UnicodeSeqOfUtf8Bytes(")"))
	var _15___v13 bool
	_ = _15___v13
	var _out5 bool
	_ = _out5
	_out5 = (solver).SendCommand(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(assert (=> "), _14_finalNegations), _dafny.UnicodeSeqOfUtf8Bytes(" (= packet_action \"NOMATCH\")))")))
	_15___v13 = _out5
	var _16_conditions _dafny.Sequence
	_ = _16_conditions
	_16_conditions = m_IptablesToSmt.Companion_Default___.BuildRuleConditions(packetSpec, _dafny.UnicodeSeqOfUtf8Bytes("packet_"))
	var _17_j _dafny.Int
	_ = _17_j
	_17_j = _dafny.Zero
	for (_17_j).Cmp(_dafny.IntOfUint32((_16_conditions).Cardinality())) < 0 {
		var _18___v14 bool
		_ = _18___v14
		var _out6 bool
		_ = _out6
		_out6 = (solver).SendCommand(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(assert "), (_16_conditions).Select((_17_j).Uint32()).(_dafny.Sequence)), _dafny.UnicodeSeqOfUtf8Bytes(")")))
		_18___v14 = _out6
		_17_j = (_17_j).Plus(_dafny.One)
	}
	var _19_res m_IptablesToSmt.SmtResult
	_ = _19_res
	var _out7 m_IptablesToSmt.SmtResult
	_ = _out7
	_out7 = (solver).CheckSat()
	_19_res = _out7
	var _source0 m_IptablesToSmt.SmtResult = _19_res
	_ = _source0
	{
		{
			if _source0.Is_Sat() {
				var _out8 _dafny.Sequence
				_ = _out8
				_out8 = (solver).GetValue(_dafny.UnicodeSeqOfUtf8Bytes("packet_action"))
				action = _out8
				goto Lmatch0
			}
		}
		{
			if _source0.Is_Unsat() {
				action = _dafny.UnicodeSeqOfUtf8Bytes("Unsat")
				goto Lmatch0
			}
		}
		{
			if _source0.Is_Unknown() {
				action = _dafny.UnicodeSeqOfUtf8Bytes("Unknown")
				goto Lmatch0
			}
		}
		{
			var _20_msg _dafny.Sequence = _source0.Get_().(m_IptablesToSmt.SmtResult_Error).Msg
			_ = _20_msg
			action = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("Error: "), _20_msg)
		}
		goto Lmatch0
	}
Lmatch0:
	return action
}
func (_static *CompanionStruct_Default___) FindRedundantRules(rules _dafny.Sequence, solver m_IptablesToSmt.SmtSolver) _dafny.Sequence {
	var redundantIndices _dafny.Sequence = _dafny.EmptySeq
	_ = redundantIndices
	var _0___v15 bool
	_ = _0___v15
	var _out0 bool
	_ = _out0
	_out0 = (solver).SendCommand(_dafny.UnicodeSeqOfUtf8Bytes("(reset)"))
	_0___v15 = _out0
	var _1___v16 bool
	_ = _1___v16
	var _out1 bool
	_ = _out1
	_out1 = (solver).SendCommand(_dafny.UnicodeSeqOfUtf8Bytes("(set-logic ALL)"))
	_1___v16 = _out1
	var _2_decls _dafny.Sequence
	_ = _2_decls
	_2_decls = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_chain String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_src String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_dst String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_proto String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_sport String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_dport String)"), _dafny.UnicodeSeqOfUtf8Bytes("(declare-const packet_action String)"), _dafny.UnicodeSeqOfUtf8Bytes("(assert (distinct packet_chain \"\"))"))
	var _3_k _dafny.Int
	_ = _3_k
	_3_k = _dafny.Zero
	for (_3_k).Cmp(_dafny.IntOfUint32((_2_decls).Cardinality())) < 0 {
		var _4___v17 bool
		_ = _4___v17
		var _out2 bool
		_ = _out2
		_out2 = (solver).SendCommand((_2_decls).Select((_3_k).Uint32()).(_dafny.Sequence))
		_4___v17 = _out2
		_3_k = (_3_k).Plus(_dafny.One)
	}
	redundantIndices = _dafny.SeqOf()
	var _5_i _dafny.Int
	_ = _5_i
	_5_i = _dafny.Zero
	for (_5_i).Cmp(_dafny.IntOfUint32((rules).Cardinality())) < 0 {
		var _6_rule m_IptablesToSmt.Rule
		_ = _6_rule
		_6_rule = (rules).Select((_5_i).Uint32()).(m_IptablesToSmt.Rule)
		var _7_indexText _dafny.Sequence
		_ = _7_indexText
		_7_indexText = m_IptablesToSmt.Companion_Default___.IntToString(_5_i)
		var _8_def _dafny.Sequence
		_ = _8_def
		_8_def = m_IptablesToSmt.Companion_Default___.DefineRuleFunction(_6_rule, _5_i)
		var _9___v18 bool
		_ = _9___v18
		var _out3 bool
		_ = _out3
		_out3 = (solver).SendCommand(_8_def)
		_9___v18 = _out3
		var _10_ruleCall _dafny.Sequence
		_ = _10_ruleCall
		_10_ruleCall = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(rule"), _7_indexText), _dafny.UnicodeSeqOfUtf8Bytes(" packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)"))
		var _11___v19 bool
		_ = _11___v19
		var _out4 bool
		_ = _out4
		_out4 = (solver).SendCommand(_dafny.UnicodeSeqOfUtf8Bytes("(push)"))
		_11___v19 = _out4
		var _12___v20 bool
		_ = _12___v20
		var _out5 bool
		_ = _out5
		_out5 = (solver).SendCommand(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(assert "), _10_ruleCall), _dafny.UnicodeSeqOfUtf8Bytes(")")))
		_12___v20 = _out5
		var _13_res m_IptablesToSmt.SmtResult
		_ = _13_res
		var _out6 m_IptablesToSmt.SmtResult
		_ = _out6
		_out6 = (solver).CheckSat()
		_13_res = _out6
		if (_13_res).Equals(m_IptablesToSmt.Companion_SmtResult_.Create_Unsat_()) {
			redundantIndices = _dafny.Companion_Sequence_.Concatenate(redundantIndices, _dafny.SeqOf(_5_i))
		}
		var _14___v21 bool
		_ = _14___v21
		var _out7 bool
		_ = _out7
		_out7 = (solver).SendCommand(_dafny.UnicodeSeqOfUtf8Bytes("(pop)"))
		_14___v21 = _out7
		var _15___v22 bool
		_ = _15___v22
		var _out8 bool
		_ = _out8
		_out8 = (solver).SendCommand(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("(assert (not "), _10_ruleCall), _dafny.UnicodeSeqOfUtf8Bytes("))")))
		_15___v22 = _out8
		_5_i = (_5_i).Plus(_dafny.One)
	}
	return redundantIndices
}

// End of class Default__
