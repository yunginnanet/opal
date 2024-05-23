package lex

type (
	Mode int
)

const (
	modeNone Mode = iota
	modeControl
	modeModifier
	modeAssign
	modeCommand
	modeType
)

var modeToName = map[Mode]string{
	modeNone: "none", modeControl: "control",
	modeCommand: "command", modeType: "type",
}
