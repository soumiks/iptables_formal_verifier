// Package Program
// Dafny module Program compiled into Go

package Program

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m__System "System_"
	_dafny "dafny"
	os "os"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_IptablesToSmt.Dummy__
var _ m_Analysis.Dummy__

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
	return "Program.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Main(args _dafny.Sequence) {
	if (_dafny.IntOfUint32((args).Cardinality())).Cmp(_dafny.One) <= 0 {
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Usage:\n").VerbatimString(false))
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  dafny-iptables <rules_file>               (Translate to SMT)\n").VerbatimString(false))
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  dafny-iptables check-eq <f1> <f2>         (Check Equivalence)\n").VerbatimString(false))
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  dafny-iptables query <f> \"constraint\"     (Query Packet)\n").VerbatimString(false))
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  dafny-iptables optimize <f>               (Find Redundant)\n").VerbatimString(false))
		return
	}
	var _0_command _dafny.Sequence
	_ = _0_command
	_0_command = (args).Select(1).(_dafny.Sequence)
	if (_dafny.Companion_Sequence_.Equal(_0_command, _dafny.UnicodeSeqOfUtf8Bytes("check-eq"))) && ((_dafny.IntOfUint32((args).Cardinality())).Cmp(_dafny.IntOfInt64(4)) >= 0) {
		var _1_solver *RealSmtSolver
		_ = _1_solver
		var _nw0 *RealSmtSolver = New_RealSmtSolver_()
		_ = _nw0
		_nw0.Ctor__()
		_1_solver = _nw0
		(_1_solver).Initialize()
		var _2_path1 _dafny.Sequence
		_ = _2_path1
		_2_path1 = (args).Select(2).(_dafny.Sequence)
		var _3_path2 _dafny.Sequence
		_ = _3_path2
		_3_path2 = (args).Select(3).(_dafny.Sequence)
		var _4_content1 _dafny.Sequence
		_ = _4_content1
		var _5_ok1 bool
		_ = _5_ok1
		var _out0 _dafny.Sequence
		_ = _out0
		var _out1 bool
		_ = _out1
		_out0, _out1 = (_1_solver).ReadFile(_2_path1)
		_4_content1 = _out0
		_5_ok1 = _out1
		if !(_5_ok1) {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Could not read file: ").VerbatimString(false))
			_dafny.Print(_2_path1.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			return
		}
		var _6_content2 _dafny.Sequence
		_ = _6_content2
		var _7_ok2 bool
		_ = _7_ok2
		var _out2 _dafny.Sequence
		_ = _out2
		var _out3 bool
		_ = _out3
		_out2, _out3 = (_1_solver).ReadFile(_3_path2)
		_6_content2 = _out2
		_7_ok2 = _out3
		if !(_7_ok2) {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Could not read file: ").VerbatimString(false))
			_dafny.Print(_3_path2.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			return
		}
		var _8_lines1 _dafny.Sequence
		_ = _8_lines1
		var _out4 _dafny.Sequence
		_ = _out4
		_out4 = m_IptablesToSmt.Companion_Default___.SplitLines(_4_content1)
		_8_lines1 = _out4
		var _9_rules1 _dafny.Sequence
		_ = _9_rules1
		var _out5 _dafny.Sequence
		_ = _out5
		_out5 = m_IptablesToSmt.Companion_Default___.ParseRules(_8_lines1)
		_9_rules1 = _out5
		var _10_lines2 _dafny.Sequence
		_ = _10_lines2
		var _out6 _dafny.Sequence
		_ = _out6
		_out6 = m_IptablesToSmt.Companion_Default___.SplitLines(_6_content2)
		_10_lines2 = _out6
		var _11_rules2 _dafny.Sequence
		_ = _11_rules2
		var _out7 _dafny.Sequence
		_ = _out7
		_out7 = m_IptablesToSmt.Companion_Default___.ParseRules(_10_lines2)
		_11_rules2 = _out7
		if ((_dafny.IntOfUint32((_9_rules1).Cardinality())).Sign() == 0) || ((_dafny.IntOfUint32((_11_rules2).Cardinality())).Sign() == 0) {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: One of the rule files is empty or invalid.\n").VerbatimString(false))
			return
		}
		var _12_defaultRaw _dafny.Sequence
		_ = _12_defaultRaw
		_12_defaultRaw = _dafny.UnicodeSeqOfUtf8Bytes("-A INPUT -j DROP")
		var _13_defaultToks _dafny.Sequence
		_ = _13_defaultToks
		var _out8 _dafny.Sequence
		_ = _out8
		_out8 = m_IptablesToSmt.Companion_Default___.SplitOnWhitespace(_12_defaultRaw)
		_13_defaultToks = _out8
		var _14_defaultRule m_IptablesToSmt.Rule
		_ = _14_defaultRule
		var _15_defOk bool
		_ = _15_defOk
		var _out9 m_IptablesToSmt.Rule
		_ = _out9
		var _out10 bool
		_ = _out10
		_out9, _out10 = m_IptablesToSmt.Companion_Default___.ParseRuleTokens(_13_defaultToks, _dafny.IntOfInt64(99999), _12_defaultRaw)
		_14_defaultRule = _out9
		_15_defOk = _out10
		var _16_rules1Final _dafny.Sequence
		_ = _16_rules1Final
		if _15_defOk {
			_16_rules1Final = _dafny.Companion_Sequence_.Concatenate(_9_rules1, _dafny.SeqOf(_14_defaultRule))
		} else {
			_16_rules1Final = _9_rules1
		}
		var _17_rules2Final _dafny.Sequence
		_ = _17_rules2Final
		if _15_defOk {
			_17_rules2Final = _dafny.Companion_Sequence_.Concatenate(_11_rules2, _dafny.SeqOf(_14_defaultRule))
		} else {
			_17_rules2Final = _11_rules2
		}
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Checking equivalence between ").VerbatimString(false))
		_dafny.Print(_dafny.IntOfUint32((_16_rules1Final).Cardinality()))
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" rules and ").VerbatimString(false))
		_dafny.Print(_dafny.IntOfUint32((_17_rules2Final).Cardinality()))
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" rules (including default DROP)...\n").VerbatimString(false))
		var _18_equivalent bool
		_ = _18_equivalent
		var _out11 bool
		_ = _out11
		_out11 = m_Analysis.Companion_Default___.CheckEquivalence(_16_rules1Final, _17_rules2Final, _1_solver)
		_18_equivalent = _out11
		if _18_equivalent {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("RESULT: EQUIVALENT\n").VerbatimString(false))
		} else {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("RESULT: NOT EQUIVALENT\n").VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Counterexample:\n").VerbatimString(false))
			var _19_c _dafny.Sequence
			_ = _19_c
			var _out12 _dafny.Sequence
			_ = _out12
			_out12 = (_1_solver).GetValue(_dafny.UnicodeSeqOfUtf8Bytes("packet_chain"))
			_19_c = _out12
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  Chain: ").VerbatimString(false))
			_dafny.Print(_19_c.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			var _20_s _dafny.Sequence
			_ = _20_s
			var _out13 _dafny.Sequence
			_ = _out13
			_out13 = (_1_solver).GetValue(_dafny.UnicodeSeqOfUtf8Bytes("packet_src"))
			_20_s = _out13
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  Src: ").VerbatimString(false))
			_dafny.Print(_20_s.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			var _21_d _dafny.Sequence
			_ = _21_d
			var _out14 _dafny.Sequence
			_ = _out14
			_out14 = (_1_solver).GetValue(_dafny.UnicodeSeqOfUtf8Bytes("packet_dst"))
			_21_d = _out14
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  Dst: ").VerbatimString(false))
			_dafny.Print(_21_d.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			var _22_p _dafny.Sequence
			_ = _22_p
			var _out15 _dafny.Sequence
			_ = _out15
			_out15 = (_1_solver).GetValue(_dafny.UnicodeSeqOfUtf8Bytes("packet_proto"))
			_22_p = _out15
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  Proto: ").VerbatimString(false))
			_dafny.Print(_22_p.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			var _23_sp _dafny.Sequence
			_ = _23_sp
			var _out16 _dafny.Sequence
			_ = _out16
			_out16 = (_1_solver).GetValue(_dafny.UnicodeSeqOfUtf8Bytes("packet_sport"))
			_23_sp = _out16
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  SPort: ").VerbatimString(false))
			_dafny.Print(_23_sp.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			var _24_dp _dafny.Sequence
			_ = _24_dp
			var _out17 _dafny.Sequence
			_ = _out17
			_out17 = (_1_solver).GetValue(_dafny.UnicodeSeqOfUtf8Bytes("packet_dport"))
			_24_dp = _out17
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  DPort: ").VerbatimString(false))
			_dafny.Print(_24_dp.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			var _25_a1 _dafny.Sequence
			_ = _25_a1
			var _out18 _dafny.Sequence
			_ = _out18
			_out18 = (_1_solver).GetValue(_dafny.UnicodeSeqOfUtf8Bytes("packet_action_1"))
			_25_a1 = _out18
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  Action 1: ").VerbatimString(false))
			_dafny.Print(_25_a1.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			var _26_a2 _dafny.Sequence
			_ = _26_a2
			var _out19 _dafny.Sequence
			_ = _out19
			_out19 = (_1_solver).GetValue(_dafny.UnicodeSeqOfUtf8Bytes("packet_action_2"))
			_26_a2 = _out19
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("  Action 2: ").VerbatimString(false))
			_dafny.Print(_26_a2.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
		}
	} else if (_dafny.Companion_Sequence_.Equal(_0_command, _dafny.UnicodeSeqOfUtf8Bytes("query"))) && ((_dafny.IntOfUint32((args).Cardinality())).Cmp(_dafny.IntOfInt64(4)) >= 0) {
		var _27_solver *RealSmtSolver
		_ = _27_solver
		var _nw1 *RealSmtSolver = New_RealSmtSolver_()
		_ = _nw1
		_nw1.Ctor__()
		_27_solver = _nw1
		(_27_solver).Initialize()
		var _28_path _dafny.Sequence
		_ = _28_path
		_28_path = (args).Select(2).(_dafny.Sequence)
		var _29_spec _dafny.Sequence
		_ = _29_spec
		_29_spec = (args).Select(3).(_dafny.Sequence)
		var _30_content _dafny.Sequence
		_ = _30_content
		var _31_ok bool
		_ = _31_ok
		var _out20 _dafny.Sequence
		_ = _out20
		var _out21 bool
		_ = _out21
		_out20, _out21 = (_27_solver).ReadFile(_28_path)
		_30_content = _out20
		_31_ok = _out21
		if !(_31_ok) {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Could not read file: ").VerbatimString(false))
			_dafny.Print(_28_path.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			return
		}
		var _32_lines _dafny.Sequence
		_ = _32_lines
		var _out22 _dafny.Sequence
		_ = _out22
		_out22 = m_IptablesToSmt.Companion_Default___.SplitLines(_30_content)
		_32_lines = _out22
		var _33_rules _dafny.Sequence
		_ = _33_rules
		var _out23 _dafny.Sequence
		_ = _out23
		_out23 = m_IptablesToSmt.Companion_Default___.ParseRules(_32_lines)
		_33_rules = _out23
		if (_dafny.IntOfUint32((_33_rules).Cardinality())).Sign() == 0 {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Rules file is empty or invalid.\n").VerbatimString(false))
			return
		}
		var _34_rawSpec _dafny.Sequence
		_ = _34_rawSpec
		_34_rawSpec = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("-A INPUT "), _29_spec), _dafny.UnicodeSeqOfUtf8Bytes(" -j ACCEPT"))
		var _35_specTokens _dafny.Sequence
		_ = _35_specTokens
		var _out24 _dafny.Sequence
		_ = _out24
		_out24 = m_IptablesToSmt.Companion_Default___.SplitOnWhitespace(_34_rawSpec)
		_35_specTokens = _out24
		var _36_packetRule m_IptablesToSmt.Rule
		_ = _36_packetRule
		var _37_parseOk bool
		_ = _37_parseOk
		var _out25 m_IptablesToSmt.Rule
		_ = _out25
		var _out26 bool
		_ = _out26
		_out25, _out26 = m_IptablesToSmt.Companion_Default___.ParseRuleTokens(_35_specTokens, _dafny.One, _34_rawSpec)
		_36_packetRule = _out25
		_37_parseOk = _out26
		if !(_37_parseOk) {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Invalid packet spec format: ").VerbatimString(false))
			_dafny.Print(_29_spec.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			return
		}
		var _38_action _dafny.Sequence
		_ = _38_action
		var _out27 _dafny.Sequence
		_ = _out27
		_out27 = m_Analysis.Companion_Default___.QueryPacket(_33_rules, _36_packetRule, _27_solver)
		_38_action = _out27
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Packet Action: ").VerbatimString(false))
		_dafny.Print(_38_action.VerbatimString(false))
		_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	} else if (_dafny.Companion_Sequence_.Equal(_0_command, _dafny.UnicodeSeqOfUtf8Bytes("optimize"))) && ((_dafny.IntOfUint32((args).Cardinality())).Cmp(_dafny.IntOfInt64(3)) >= 0) {
		var _39_solver *RealSmtSolver
		_ = _39_solver
		var _nw2 *RealSmtSolver = New_RealSmtSolver_()
		_ = _nw2
		_nw2.Ctor__()
		_39_solver = _nw2
		(_39_solver).Initialize()
		var _40_path _dafny.Sequence
		_ = _40_path
		_40_path = (args).Select(2).(_dafny.Sequence)
		var _41_content _dafny.Sequence
		_ = _41_content
		var _42_ok bool
		_ = _42_ok
		var _out28 _dafny.Sequence
		_ = _out28
		var _out29 bool
		_ = _out29
		_out28, _out29 = (_39_solver).ReadFile(_40_path)
		_41_content = _out28
		_42_ok = _out29
		if !(_42_ok) {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Could not read file: ").VerbatimString(false))
			_dafny.Print(_40_path.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			return
		}
		var _43_lines _dafny.Sequence
		_ = _43_lines
		var _out30 _dafny.Sequence
		_ = _out30
		_out30 = m_IptablesToSmt.Companion_Default___.SplitLines(_41_content)
		_43_lines = _out30
		var _44_rules _dafny.Sequence
		_ = _44_rules
		var _out31 _dafny.Sequence
		_ = _out31
		_out31 = m_IptablesToSmt.Companion_Default___.ParseRules(_43_lines)
		_44_rules = _out31
		if (_dafny.IntOfUint32((_44_rules).Cardinality())).Sign() == 0 {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Rules file is empty or invalid.\n").VerbatimString(false))
			return
		}
		var _45_redundant _dafny.Sequence
		_ = _45_redundant
		var _out32 _dafny.Sequence
		_ = _out32
		_out32 = m_Analysis.Companion_Default___.FindRedundantRules(_44_rules, _39_solver)
		_45_redundant = _out32
		if (_dafny.IntOfUint32((_45_redundant).Cardinality())).Sign() == 0 {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("No redundant rules found.\n").VerbatimString(false))
		} else {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Redundant rules found:\n").VerbatimString(false))
			var _46_i _dafny.Int
			_ = _46_i
			_46_i = _dafny.Zero
			for (_46_i).Cmp(_dafny.IntOfUint32((_45_redundant).Cardinality())) < 0 {
				var _47_idx _dafny.Int
				_ = _47_idx
				_47_idx = (_45_redundant).Select((_46_i).Uint32()).(_dafny.Int)
				if ((_47_idx).Sign() != -1) && ((_47_idx).Cmp(_dafny.IntOfUint32((_44_rules).Cardinality())) < 0) {
					_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Rule ").VerbatimString(false))
					_dafny.Print(_47_idx)
					_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes(" (line ").VerbatimString(false))
					_dafny.Print(((_44_rules).Select((_47_idx).Uint32()).(m_IptablesToSmt.Rule)).Dtor_lineNumber())
					_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("): ").VerbatimString(false))
					_dafny.Print(((_44_rules).Select((_47_idx).Uint32()).(m_IptablesToSmt.Rule)).Dtor_original().VerbatimString(false))
					_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
				}
				_46_i = (_46_i).Plus(_dafny.One)
			}
		}
	} else {
		var _48_path _dafny.Sequence
		_ = _48_path
		_48_path = (args).Select(1).(_dafny.Sequence)
		var _49_solver *RealSmtSolver
		_ = _49_solver
		var _nw3 *RealSmtSolver = New_RealSmtSolver_()
		_ = _nw3
		_nw3.Ctor__()
		_49_solver = _nw3
		var _50_content _dafny.Sequence
		_ = _50_content
		var _51_ok bool
		_ = _51_ok
		var _out33 _dafny.Sequence
		_ = _out33
		var _out34 bool
		_ = _out34
		_out33, _out34 = (_49_solver).ReadFile(_48_path)
		_50_content = _out33
		_51_ok = _out34
		if !(_51_ok) {
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("Error: Could not read file: ").VerbatimString(false))
			_dafny.Print(_48_path.VerbatimString(false))
			_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
			return
		}
		var _52_smt _dafny.Sequence
		_ = _52_smt
		var _out35 _dafny.Sequence
		_ = _out35
		_out35 = m_IptablesToSmt.Companion_Default___.ConvertIptablesToSmt(_50_content)
		_52_smt = _out35
		_dafny.Print(_52_smt.VerbatimString(false))
	}
}

// End of class Default__

// Definition of class RealSmtSolver
type RealSmtSolver struct {
	dummy byte
}

func New_RealSmtSolver_() *RealSmtSolver {
	_this := RealSmtSolver{}

	return &_this
}

type CompanionStruct_RealSmtSolver_ struct {
}

var Companion_RealSmtSolver_ = CompanionStruct_RealSmtSolver_{}

func (_this *RealSmtSolver) Equals(other *RealSmtSolver) bool {
	return _this == other
}

func (_this *RealSmtSolver) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*RealSmtSolver)
	return ok && _this.Equals(other)
}

func (*RealSmtSolver) String() string {
	return "Program.RealSmtSolver"
}

func Type_RealSmtSolver_() _dafny.TypeDescriptor {
	return type_RealSmtSolver_{}
}

type type_RealSmtSolver_ struct {
}

func (_this type_RealSmtSolver_) Default() interface{} {
	return (*RealSmtSolver)(nil)
}

func (_this type_RealSmtSolver_) String() string {
	return "Program.RealSmtSolver"
}
func (_this *RealSmtSolver) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){m_IptablesToSmt.Companion_SmtSolver_.TraitID_}
}

var _ m_IptablesToSmt.SmtSolver = &RealSmtSolver{}
var _ _dafny.TraitOffspring = &RealSmtSolver{}

func (_this *RealSmtSolver) Ctor__() {
	{
	}
}

// End of class RealSmtSolver
