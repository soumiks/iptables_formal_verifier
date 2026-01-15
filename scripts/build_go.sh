#!/bin/bash
set -e

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
rm -rf dafny-iptables-go

# 1. Transpile Dafny to Go
# --output:dafny-iptables will create 'dafny-iptables-go' directory
# 1. Transpile Dafny to Go
# --output:dafny-go-output will create 'dafny-go-output-go' directory
dafny translate go src/Program.dfy --output:dafny-go-output --allow-external-contracts

# Copy extern implementation
echo "Copying externs..."
cp src/dafny_extern.go dafny-go-output-go/src/Program/dafny_extern.go

# 2. Build the final Go binary
echo "Compiling Go binary..."
# Dafny generates code that expects a GOPATH structure (src/PackageName),
# so we set GOPATH to the generated directory.
export GOPATH=$(pwd)/dafny-go-output-go
export GO111MODULE=off

cd dafny-go-output-go/src

# Build the binary
go build -o ../../$OUTPUT_NAME dafny-go-output.go

cd ../..

echo "Build successful! Executable created at: ./$OUTPUT_NAME"
