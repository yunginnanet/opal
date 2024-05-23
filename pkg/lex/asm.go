//go:build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"

	"git.tcp.direct/kayos/opal/pkg/lex"
)

var bs = []String{
	String(' '), lex.EQ, lex.LB, lex.RB, lex.LP, lex.RP, lex.SEMIC, lex.COMMA,
	lex.PLUS, lex.MINUS, lex.PIPE,
}

func main() {
	bytes := GLOBL("bytes", RODATA|NOPTR)
	offset := 0
	for _, b := range bs {
		if len([]byte(b)) != 1 {
			panic("bad byte")
		}
		DATA(int(string(b)[0]), b)
		offset += len([]byte(b))
	}
	TEXT("lookup", NOSPLIT, "func(i int) byte")
	Doc("lookup returns byte i in the 'bytes' global data section.")
	i := Load(Param("i"), GP64())
	ptr := Mem{Base: GP64()}
	LEAQ(bytes, ptr.Base)
	b := GP8()
	MOVB(ptr.Idx(i, 1), b)
	Store(b, ReturnIndex(0))
	RET()

	Generate()
}
