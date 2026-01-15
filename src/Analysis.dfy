include "IptablesToSmt.dfy"

module Analysis {
  import opened IptablesToSmt

  // Orchestrator Logic for Equivalence Checking
  // Returns true if the two rulesets result in the same action for all packets.
  method CheckEquivalence(rules1: seq<Rule>, rules2: seq<Rule>, solver: SmtSolver) returns (areEquivalent: bool)
    requires |rules1| > 0 && |rules2| > 0
    modifies solver
  {
    // 1. Reset Solver
    var cmdOk := solver.SendCommand("(reset)");
    if !cmdOk {
        print "Error: Failed to reset solver\n";
        return false;
    }
    
    cmdOk := solver.SendCommand("(set-logic ALL)");

    // 2. Define Packet Fields (Common)
    var decls := [
      "(declare-const packet_chain String)",
      "(declare-const packet_src String)",
      "(declare-const packet_dst String)",
      "(declare-const packet_proto String)",
      "(declare-const packet_sport String)",
      "(declare-const packet_dport String)",
      "(declare-const packet_action_1 String)", // Action for ruleset 1
      "(declare-const packet_action_2 String)", // Action for ruleset 2
      "(assert (distinct packet_chain \"\"))"
    ];
    
    var k := 0;
    while k < |decls| {
        var _ := solver.SendCommand(decls[k]);
        k := k + 1;
    }

    // 3. Define Rules for Set 1
    var i := 0;
    var prevNegations1 := "";
    while i < |rules1|
      invariant 0 <= i <= |rules1|
    {
       var rule := rules1[i];
       var def := DefineRuleFunction(rule, i);
       var _ := solver.SendCommand(def);
       
       var assertion := AssertRuleLogic(rule, i, prevNegations1, "packet_action_1");
       var _ := solver.SendCommand(assertion);
       
       var indexText := IntToString(i);
       var call := "(rule" + indexText + " packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)";
       prevNegations1 := prevNegations1 + " (not " + call + ")";
       
       i := i + 1;
    }

    var finalNegations1 := "(and" + prevNegations1 + ")";
    var _ := solver.SendCommand("(assert (=> " + finalNegations1 + " (= packet_action_1 \"DEFAULT\")))");

    // 4. Define Rules for Set 2
    // We start index at |rules1| to avoid name collisions (rule0, rule1...)
    var j := 0;
    var prevNegations2 := "";
    var offset := |rules1|;
    
    while j < |rules2|
      invariant 0 <= j <= |rules2|
    {
       var rule := rules2[j];
       var actualIndex := offset + j;
       
       var def := DefineRuleFunction(rule, actualIndex);
       var _ := solver.SendCommand(def);
       
       var assertion := AssertRuleLogic(rule, actualIndex, prevNegations2, "packet_action_2");
       var _ := solver.SendCommand(assertion);
       
       var indexText := IntToString(actualIndex);
       var call := "(rule" + indexText + " packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)";
       prevNegations2 := prevNegations2 + " (not " + call + ")";
       
       j := j + 1;
    }
    
    var finalNegations2 := "(and" + prevNegations2 + ")";
    var _ := solver.SendCommand("(assert (=> " + finalNegations2 + " (= packet_action_2 \"DEFAULT\")))");
    
    // 5. Assert Inequality
    // We want to find a packet where action1 != action2
    var _ := solver.SendCommand("(assert (not (= packet_action_1 packet_action_2)))");
    
    // 6. Check Satisfiability
    var result := solver.CheckSat();
    match result {
        case Unsat => areEquivalent := true; // No counterexample found -> Equivalent
        case Sat => areEquivalent := false; // Counterexample found -> Not Equivalent
        case Unknown => 
             print "Warning: Solver returned unknown\n";
             areEquivalent := false;
        case Error(msg) => 
             print "Error: Solver failed: ", msg, "\n";
             areEquivalent := false;
    }
  }

  // Queries how a specific packet is handled by the ruleset.
  // Returns the action string (e.g. "ACCEPT", "DROP")
  method QueryPacket(rules: seq<Rule>, packetSpec: Rule, solver: SmtSolver) returns (action: string)
    requires |rules| > 0
    modifies solver
  {
    // 1. Reset & Setup
    var _ := solver.SendCommand("(reset)");
    var _ := solver.SendCommand("(set-logic ALL)");
    
    // Define fields
    var decls := [
      "(declare-const packet_chain String)",
      "(declare-const packet_src String)",
      "(declare-const packet_dst String)",
      "(declare-const packet_proto String)",
      "(declare-const packet_sport String)",
      "(declare-const packet_dport String)",
      "(declare-const packet_action String)",
      "(assert (distinct packet_chain \"\"))"
    ];
    var k := 0; 
    while k < |decls| { var _ := solver.SendCommand(decls[k]); k := k + 1; }

    // 2. Assert Rules (Standard Logic)
    var i := 0;
    var prevNegations := "";
    while i < |rules|
      invariant 0 <= i <= |rules|
    {
       var rule := rules[i];
       var def := DefineRuleFunction(rule, i);
       var _ := solver.SendCommand(def);
       var assertion := AssertRuleLogic(rule, i, prevNegations, "packet_action");
       var _ := solver.SendCommand(assertion);
       
       var indexText := IntToString(i);
       var call := "(rule" + indexText + " packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)";
       prevNegations := prevNegations + " (not " + call + ")";
       i := i + 1;
    }
    
    var finalNegations := "(and" + prevNegations + ")";
    var _ := solver.SendCommand("(assert (=> " + finalNegations + " (= packet_action \"NOMATCH\")))");

    // 3. Assert Packet Spec Constraints
    // We treat packetSpec as a rule to get its match conditions
    var conditions := BuildRuleConditions(packetSpec, "packet_");
    var j := 0;
    while j < |conditions| 
      invariant 0 <= j <= |conditions|
    {
       var _ := solver.SendCommand("(assert " + conditions[j] + ")");
       j := j + 1;
    }

    // 4. Query Value
    var res := solver.CheckSat();
    match res {
        case Sat => action := solver.GetValue("packet_action");
        case Unsat => action := "Unsat";
        case Unknown => action := "Unknown";
        case Error(msg) => action := "Error: " + msg;
    }
  }

  // Finds indices of rules that are redundant (shadowed by previous rules).
  method FindRedundantRules(rules: seq<Rule>, solver: SmtSolver) returns (redundantIndices: seq<int>)
    modifies solver
  {
    // 1. Reset & Setup
    var _ := solver.SendCommand("(reset)");
    var _ := solver.SendCommand("(set-logic ALL)");
    var decls := [
      "(declare-const packet_chain String)",
      "(declare-const packet_src String)",
      "(declare-const packet_dst String)",
      "(declare-const packet_proto String)",
      "(declare-const packet_sport String)",
      "(declare-const packet_dport String)",
      "(declare-const packet_action String)", // Not strictly needed for reachability but keeps consistency
      "(assert (distinct packet_chain \"\"))"
    ];
    var k := 0; 
    while k < |decls| { var _ := solver.SendCommand(decls[k]); k := k + 1; }
    
    redundantIndices := [];
    
    // 2. Loop through rules
    // Pattern: 
    //   Context = (not rule0) AND (not rule1) ...
    //   Test Rule K: Assert Context AND ruleK.
    //   If Unsat => Rule K matches nothing new => Redundant.
    //   Update Context: Add (not ruleK).
    
    var i := 0;
    while i < |rules|
      invariant 0 <= i <= |rules|
    {
       var rule := rules[i];
       var indexText := IntToString(i);
       
       // Define the rule predicate
       var def := DefineRuleFunction(rule, i);
       var _ := solver.SendCommand(def);
       
       var ruleCall := "(rule" + indexText + " packet_chain packet_src packet_dst packet_proto packet_sport packet_dport)";
       
       // Test if this rule is reachable given current context
       var _ := solver.SendCommand("(push)"); // Save context
       var _ := solver.SendCommand("(assert " + ruleCall + ")");
       
       var res := solver.CheckSat();
       if res == Unsat {
          // Reachable? No. -> Redundant
          redundantIndices := redundantIndices + [i];
       }
       
       var _ := solver.SendCommand("(pop)"); // Restore to just context
       
       // Add this rule's negation to the context for future rules
       // "If we reached here, we didn't match previous rules AND we didn't match this one"
       // Actually, for reachability of NEXT rule, we assume we didn't match THIS rule.
       var _ := solver.SendCommand("(assert (not " + ruleCall + "))");
       
       i := i + 1;
    }
  }
}
