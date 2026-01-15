include "IptablesToSmt.dfy"
include "Analysis.dfy"

module Program {
  import opened IptablesToSmt
  import opened Analysis

  // Concrete implementation of SMT Solver that calls into Go
  class {:extern} RealSmtSolver extends SmtSolver {
      constructor() {}

      method {:extern} Initialize()

      method {:extern} SendCommand(cmd: string) returns (success: bool) 
      
      method {:extern} CheckSat() returns (result: SmtResult) 

      method {:extern} GetValue(variable: string) returns (val: string) 

      method {:extern} ReadFile(path: string) returns (content: string, ok: bool)
          ensures ok ==> forall k :: 0 <= k < |content| ==> content[k] != '\r'
  }

  method Main(args: seq<string>) {
      if |args| <= 1 {
          print "Usage:\n";
          print "  dafny-iptables <rules_file>               (Translate to SMT)\n";
          print "  dafny-iptables check-eq <f1> <f2>         (Check Equivalence)\n";
          print "  dafny-iptables query <f> \"constraint\"     (Query Packet)\n";
          print "  dafny-iptables optimize <f>               (Find Redundant)\n";
          return;
      }
      
      var command := args[1];
      
      if command == "check-eq" && |args| >= 4 {
          var solver := new RealSmtSolver();
          solver.Initialize();
          
          var path1 := args[2];
          var path2 := args[3];
          
          var content1, ok1 := solver.ReadFile(path1);
          if !ok1 {
             print "Error: Could not read file: ", path1, "\n";
             return;
          }
          
          var content2, ok2 := solver.ReadFile(path2);
          if !ok2 {
             print "Error: Could not read file: ", path2, "\n";
             return;
          }
          
          var lines1 := SplitLines(content1);
          var rules1 := ParseRules(lines1);
          
          var lines2 := SplitLines(content2);
          var rules2 := ParseRules(lines2);
          
          if |rules1| == 0 || |rules2| == 0 {
             print "Error: One of the rule files is empty or invalid.\n";
             return;
          }
          
          
          // Append default DROP rule to ensure defined behavior for unmatched packets
          // For MVP we assume INPUT chain.
          var defaultRaw := "-A INPUT -j DROP";
          var defaultToks := SplitOnWhitespace(defaultRaw);
          var defaultRule, defOk := ParseRuleTokens(defaultToks, 99999, defaultRaw);
          
          var rules1Final := if defOk then rules1 + [defaultRule] else rules1;
          var rules2Final := if defOk then rules2 + [defaultRule] else rules2;
          
          print "Checking equivalence between ", |rules1Final|, " rules and ", |rules2Final|, " rules (including default DROP)...\n";
          
          var equivalent := CheckEquivalence(rules1Final, rules2Final, solver);
          if equivalent {
              print "RESULT: EQUIVALENT\n";
          } else {
              print "RESULT: NOT EQUIVALENT\n";
              print "Counterexample:\n";
              var c := solver.GetValue("packet_chain"); print "  Chain: ", c, "\n";
              var s := solver.GetValue("packet_src"); print "  Src: ", s, "\n";
              var d := solver.GetValue("packet_dst"); print "  Dst: ", d, "\n";
              var p := solver.GetValue("packet_proto"); print "  Proto: ", p, "\n";
              var sp := solver.GetValue("packet_sport"); print "  SPort: ", sp, "\n";
              var dp := solver.GetValue("packet_dport"); print "  DPort: ", dp, "\n";
              var a1 := solver.GetValue("packet_action_1"); print "  Action 1: ", a1, "\n";
              var a2 := solver.GetValue("packet_action_2"); print "  Action 2: ", a2, "\n";
          }

      } else if command == "query" && |args| >= 4 {
          var solver := new RealSmtSolver();
          solver.Initialize();
          
          var path := args[2];
          var spec := args[3];
          
          var content, ok := solver.ReadFile(path);
          if !ok {
             print "Error: Could not read file: ", path, "\n";
             return;
          }
          
          var lines := SplitLines(content);
          var rules := ParseRules(lines);
          if |rules| == 0 {
             print "Error: Rules file is empty or invalid.\n";
             return;
          }
          
          // Parse spec
          // We construct a dummy rule line to parse the match conditions
          // Force INPUT chain and ACCEPT target to be valid
          var rawSpec := "-A INPUT " + spec + " -j ACCEPT";
          var specTokens := SplitOnWhitespace(rawSpec);
          
          // Pass lineNumber 1 to satisfy precondition
          var packetRule, parseOk := ParseRuleTokens(specTokens, 1, rawSpec);
          
          if !parseOk {
              print "Error: Invalid packet spec format: ", spec, "\n";
              return;
          }

          var action := QueryPacket(rules, packetRule, solver);
          print "Packet Action: ", action, "\n";

      } else if command == "optimize" && |args| >= 3 {
          var solver := new RealSmtSolver();
          solver.Initialize();
          
          var path := args[2];
          var content, ok := solver.ReadFile(path);
          if !ok {
             print "Error: Could not read file: ", path, "\n";
             return;
          }
          
          var lines := SplitLines(content);
          var rules := ParseRules(lines);
          if |rules| == 0 {
             print "Error: Rules file is empty or invalid.\n";
             return;
          }
          
          var redundant := FindRedundantRules(rules, solver);
          if |redundant| == 0 {
              print "No redundant rules found.\n";
          } else {
              print "Redundant rules found:\n";
              var i := 0;
              while i < |redundant| {
                  var idx := redundant[i];
                  if 0 <= idx < |rules| {
                      print "Rule ", idx, " (line ", rules[idx].lineNumber, "): ", rules[idx].original, "\n";
                  }
                  i := i + 1;
              }
          }

      } else {
          // Default: Translate
          // We used to accept payload as string. Now we expect a file path.
          // args[1] is the file path.
          var path := args[1];
          
          // We use RealSmtSolver only for ReadFile
          var solver := new RealSmtSolver();
          // No Initialize needed for ReadFile
          
          var content, ok := solver.ReadFile(path);
          if !ok {
              print "Error: Could not read file: ", path, "\n";
              return;
          }
          
          var smt := ConvertIptablesToSmt(content);
          print smt;
      }
  }
}
