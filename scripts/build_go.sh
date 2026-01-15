#!/bin/bash
set -e

# Turn on Go Modules explicitly
export GO111MODULE=on

# Default output name
OUTPUT_NAME="dafny-iptables"

# Check for required tools
if ! command -v dafny &> /dev/null; then
    echo "Error: dafny is not installed or not in PATH."
    exit 1
fi

if ! command -v go &> /dev/null; then
    echo "Error: go is not installed or not in PATH."
    exit 1
fi

# Ensure goimports is installed (Dafny's Go backend often needs it)
if ! command -v goimports &> /dev/null; then
    echo "Warning: goimports not found. Installing..."
    go install golang.org/x/tools/cmd/goimports@latest
fi

# Add GOPATH/bin to PATH so Dafny can find goimports
export PATH=$PATH:$(go env GOPATH)/bin

echo "Building Dafny project to Go..."

# Clean up previous build artifacts
rm -rf dafny-go-output-go

# 1. Transpile Dafny to Go
# --output:dafny-go-output will create 'dafny-go-output-go' directory
dafny translate go src/Program.dfy --output:dafny-go-output --allow-external-contracts

# Copy extern implementation
echo "Copying externs..."
cp src/dafny_extern.go dafny-go-output-go/src/Program/dafny_extern.go

# Copy missing runtime System_ if needed
if [ ! -d "dafny-go-output-go/src/System_" ]; then
    echo "Restoring missing System_ runtime..."
    mkdir -p dafny-go-output-go/src/System_
    cp scripts/std_system.go dafny-go-output-go/src/System_/System_.go
fi

# 2. Build the final Go binary using Go Modules
echo "Compiling Go binary..."

cd dafny-go-output-go/src

# Initialize a temporary module
if [ ! -f go.mod ]; then
    go mod init dafny-output
fi

# Add the official Dafny Runtime for Go
echo "Adding Dafny Runtime dependency..."
go get github.com/dafny-lang/DafnyRuntimeGo/v4

MODULE_NAME="dafny-output"

# Recursively replace imports in all .go files
echo "Fixing imports..."

if [[ "$OSTYPE" == "darwin"* ]]; then
  # Mac version of sed requires '' for -i
  find . -name "*.go" | xargs sed -i '' \
      -e 's|"System_"|"'"$MODULE_NAME"'/System_"|g' \
      -e 's|"dafny"|"github.com/dafny-lang/DafnyRuntimeGo/v4/dafny"|g' \
      -e 's|"Program"|"'"$MODULE_NAME"'/Program"|g' \
      -e 's|"IptablesToSmt"|"'"$MODULE_NAME"'/IptablesToSmt"|g' \
      -e 's|"Analysis"|"'"$MODULE_NAME"'/Analysis"|g'
else
  # Linux version of sed
  find . -name "*.go" | xargs sed -i \
      -e 's|"System_"|"'"$MODULE_NAME"'/System_"|g' \
      -e 's|"dafny"|"github.com/dafny-lang/DafnyRuntimeGo/v4/dafny"|g' \
      -e 's|"Program"|"'"$MODULE_NAME"'/Program"|g' \
      -e 's|"IptablesToSmt"|"'"$MODULE_NAME"'/IptablesToSmt"|g' \
      -e 's|"Analysis"|"'"$MODULE_NAME"'/Analysis"|g'
fi

# Tidy up dependencies
go mod tidy

# Build the binary
go build -o ../../$OUTPUT_NAME dafny-go-output.go

cd ../..

echo "Build successful! Executable created at: ./$OUTPUT_NAME"
