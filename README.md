# dafny_iptables_to_smtlib2

A formally verified Dafny program that parses a subset of the `iptables-save` format and emits an SMT-LIB 2.0 encoding. This ensures that the translation from firewall rules to logical assertions is mathematically correct.

## Features

- **Formally Verified Parser**: The input parser is verified against strong functional specifications. We prove that the `Rule` object returned exactly matches the input tokens.
- **Verified Output Generation**: The SMT-LIB printer is verified to produce safe, non-empty, and structurally correct output.
- **Unit Proofs**: Instead of runtime unit tests, this project uses compile-time "Unit Proofs" (`src/Tests.dfy`) which guarantee correctness for specific inputs mathematically.
- **Precondition Propagation**: The codebase explicitly handles preconditions (like "no carriage returns") from the core logic all the way to the `Main` entry point.

## Usage

The tool processes iptables rules and supports several modes of operation, including translation to SMT-LIB, equivalence checking, packet querying, and optimization.

### 1. Translate to SMT-LIB (Default)
Translates a rules file into an SMT-LIB 2.0 document.

```bash
# Using the Go binary
./iptables-verifier rules.txt > output.smt2

# Using Dafny directly
dafny run src/Program.dfy -- rules.txt > output.smt2
```

### 2. Check Equivalence
Determines if two rule sets are logically equivalent (i.e., they treat every possible packet identically).

```bash
./iptables-verifier check-eq rules1.txt rules2.txt
```
**Output:**
- `RESULT: EQUIVALENT`: The rules are identical in behavior.
- `RESULT: NOT EQUIVALENT`: A counterexample packet is printed showing where they differ.

### 3. Query Packet
Queries how a specific packet would be treated by the rules. The packet constraint is specified using iptables syntax (e.g., `-p tcp --dport 80`).

```bash
./iptables-verifier query rules.txt "-p tcp --dport 80 -s 10.0.0.1"
```
**Output:** `Packet Action: ACCEPT` (or DROP, REJECT, etc.)

### 4. Optimize Rules
Identifies redundant rules that can be removed without changing the firewall's behavior.

```bash
./iptables-verifier optimize rules.txt
```
**Output:** A list of redundant rules with their line numbers.

## Building with Go

To build a standalone executable using the Go backend (requires Go 1.16+):

```bash
./scripts/build_go.sh
```

This script:
1. Translates the Dafny code to Go.
2. Initializes a local Go module.
3. Compiles the binary to `./iptables-verifier`.

## Verification

To verified the correctness of the parser and printer:

```bash
dafny run src/Tests.dfy
```

This will run the suite of Unit Proofs. A successful run (exit code 0) means the code is mathematically proven to satisfy its specifications.

## Project Structure

- `src/IptablesToSmt.dfy`: Core logic for parsing and SMT generation.
- `src/Analysis.dfy`: Logic for equivalence checking, querying, and optimization.
- `src/Program.dfy`: Main entry point and CLI argument parsing.
- `src/Tests.dfy`: The suite of Unit Proofs.
- `samples/`: Sample iptables dump files.

## Modeling Details

- **Supported Flags (Parsed & Modeled)**: `-s`, `-d`, `-p`, `--sport`, `--dport`, `-j`.
- **Supported Flags (Parsed as Wildcard)**: `-i`, `-o`, `-t`, `-m` (and modules like `conntrack`, `limit`, `recent`), `--to-destination`, `--log-prefix`, etc. These are parsed to allow valid `iptables-save` files to process, but are currently treated as "always true" in the SMT encoding.
- **Targets**: `ACCEPT`, `DROP`, `REJECT`, `RETURN` are modeled explicitly. Others are generic jumps.
- **Logic**: Each `-A` rule is translated into a function `(define-fun ruleN ...)` and an assertion `(assert (=> (ruleN ...) action))`.
- **Empty Input**: If the input contains no valid `-A` rules, the tool outputs nothing (empty string) instead of an SMT preamble.

## License

MIT
