// Dafny program Program.dfy compiled into Go
package main

import (
	m_Analysis "Analysis"
	m_IptablesToSmt "IptablesToSmt"
	m_Program "Program"
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

func main() {
	defer _dafny.CatchHalt()
	m_Program.Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
