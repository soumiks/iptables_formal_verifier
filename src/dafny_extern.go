package Program

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "sync"
    "strings"

    m_IptablesToSmt "IptablesToSmt"
    _dafny "dafny"
)

// Global map to hold state for each solver instance
// Since Dafny generates the struct, we can't add fields to it.
var solverMap sync.Map // map[*RealSmtSolver]*Z3Wrapper

type Z3Wrapper struct {
    cmd    *exec.Cmd
    stdin  *bufio.Writer
    stdout *bufio.Scanner
}

func (_this *RealSmtSolver) Initialize() {
    cmd := exec.Command("z3", "-in")
    
    stdinPipe, err := cmd.StdinPipe()
    if err != nil {
        fmt.Printf("Error creating stdin pipe: %v\n", err)
        return
    }
    
    stdoutPipe, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Printf("Error creating stdout pipe: %v\n", err)
        return
    }
    
    if err := cmd.Start(); err != nil {
        fmt.Printf("Error starting z3: %v\n", err)
        return
    }
    
    wrapper := &Z3Wrapper{
        cmd:    cmd,
        stdin:  bufio.NewWriter(stdinPipe),
        stdout: bufio.NewScanner(stdoutPipe),
    }
    
    solverMap.Store(_this, wrapper)
}

func (_this *RealSmtSolver) SendCommand(cmdSeq _dafny.Sequence) bool {
    wrapperInterface, ok := solverMap.Load(_this)
    if !ok {
        return false
    }
    wrapper := wrapperInterface.(*Z3Wrapper)
    
    cmdStr := cmdSeq.VerbatimString(false)
    _, err := wrapper.stdin.WriteString(cmdStr + "\n")
    if err != nil {
        return false
    }
    err = wrapper.stdin.Flush()
    return err == nil
}

func (_this *RealSmtSolver) CheckSat() m_IptablesToSmt.SmtResult {
    wrapperInterface, ok := solverMap.Load(_this)
    if !ok {
        return m_IptablesToSmt.Companion_SmtResult_.Create_Error_(_dafny.UnicodeSeqOfUtf8Bytes("Solver not initialized"))
    }
    wrapper := wrapperInterface.(*Z3Wrapper)
    
    // Explicitly send check-sat because the Dfany orchestrator might have just sent logic without check-sat
    // Wait, the Orchestrator does NOT send (check-sat) via SendCommand?
    // Analysis.dfy calls `var result := solver.CheckSat();`
    // It is `CheckSat`'s job to send the command and parse the result.
    
    _, err := wrapper.stdin.WriteString("(check-sat)\n")
    if err != nil {
        return m_IptablesToSmt.Companion_SmtResult_.Create_Error_(_dafny.UnicodeSeqOfUtf8Bytes(err.Error()))
    }
    wrapper.stdin.Flush()
    
    if wrapper.stdout.Scan() {
        line := wrapper.stdout.Text()
        line = strings.TrimSpace(line)
        if line == "sat" {
            return m_IptablesToSmt.Companion_SmtResult_.Create_Sat_()
        } else if line == "unsat" {
            return m_IptablesToSmt.Companion_SmtResult_.Create_Unsat_()
        } else {
            fmt.Println("Z3 Output:", line)
            return m_IptablesToSmt.Companion_SmtResult_.Create_Unknown_()
        }
    }
    
    return m_IptablesToSmt.Companion_SmtResult_.Create_Error_(_dafny.UnicodeSeqOfUtf8Bytes("No output from Z3"))
}

func (_this *RealSmtSolver) GetValue(variable _dafny.Sequence) _dafny.Sequence {
    wrapperInterface, ok := solverMap.Load(_this)
    if !ok {
        return _dafny.UnicodeSeqOfUtf8Bytes("Error: Solver not initialized")
    }
    wrapper := wrapperInterface.(*Z3Wrapper)
    
    varName := variable.VerbatimString(false)
    cmd := fmt.Sprintf("(get-value (%s))", varName)
    
    _, err := wrapper.stdin.WriteString(cmd + "\n")
    if err != nil {
        return _dafny.UnicodeSeqOfUtf8Bytes("Error: Failed to write to Z3")
    }
    wrapper.stdin.Flush()
    
    // Parse output which looks like ((packet_action "ACCEPT"))
    // Or for complex types: ((n 0))
    // We expect 1 line for simple scalar values in standard Z3 -in mode usually?
    // Actually, (get-value) output can match parenthesis balancing.
    // For this tool, we only look for (var value).
    
    if wrapper.stdout.Scan() {
        line := wrapper.stdout.Text()
        // Format: ((packet_action "ACCEPT"))
        // We want to extract "ACCEPT" (including quotes) or just ACCEPT.
        // Let's rely on simple string parsing for now.
        
        // Find the second token
        parts := strings.Split(line, " ")
        if len(parts) >= 2 {
            // ((var "VAL"))
            // last part is "VAL"))
            last := parts[len(parts)-1]
            // Trim ))
            last = strings.TrimSuffix(last, "))")
            last = strings.TrimSuffix(last, ")")
            // The value might be quoted.
            return _dafny.UnicodeSeqOfUtf8Bytes(last)
        }
    }
    
    return _dafny.UnicodeSeqOfUtf8Bytes("Error: Parse failed")
}

func (_this *RealSmtSolver) ReadFile(path _dafny.Sequence) (_dafny.Sequence, bool) {
    p := path.VerbatimString(false)
    data, err := os.ReadFile(p)
    if err != nil {
        return _dafny.EmptySeq, false
    }
    s := string(data)
    s = strings.ReplaceAll(s, "\r", "")
    return _dafny.UnicodeSeqOfUtf8Bytes(s), true
}

func (_this *RealSmtSolver) GetUnsatCore() _dafny.Sequence {
    wrapperInterface, ok := solverMap.Load(_this)
    if !ok {
        return _dafny.EmptySeq
    }
    wrapper := wrapperInterface.(*Z3Wrapper)

    cmd := "(get-unsat-core)"
    
    _, err := wrapper.stdin.WriteString(cmd + "\n")
    if err != nil {
        return _dafny.EmptySeq
    }
    wrapper.stdin.Flush()
    
    // Output: (r_0 r_5 ...)
    if wrapper.stdout.Scan() {
        line := wrapper.stdout.Text()
        line = strings.TrimSpace(line)
        line = strings.TrimPrefix(line, "(")
        line = strings.TrimSuffix(line, ")")
        
        parts := strings.Split(line, " ")
        
        seq := make([]interface{}, 0, len(parts))
        for _, p := range parts {
            if len(p) > 0 {
                seq = append(seq, _dafny.UnicodeSeqOfUtf8Bytes(p))
            }
        }
        return _dafny.SeqOf(seq...)
    }
    
    return _dafny.EmptySeq
}
