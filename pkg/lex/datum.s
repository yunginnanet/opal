// Code generated by command: go run asm.go -out datum.s -stubs datum.go. DO NOT EDIT.

#include "textflag.h"

DATA bytes<>+32(SB)/1, $" "
DATA bytes<>+61(SB)/1, $"="
DATA bytes<>+123(SB)/1, $"{"
DATA bytes<>+125(SB)/1, $"}"
DATA bytes<>+40(SB)/1, $"("
DATA bytes<>+41(SB)/1, $")"
DATA bytes<>+59(SB)/1, $";"
DATA bytes<>+44(SB)/1, $","
DATA bytes<>+43(SB)/1, $"+"
DATA bytes<>+45(SB)/1, $"-"
DATA bytes<>+124(SB)/1, $"|"
GLOBL bytes<>(SB), RODATA|NOPTR, $126

// func lookup(i int) byte
TEXT ·lookup(SB), NOSPLIT, $0-9
	MOVQ i+0(FP), AX
	LEAQ bytes<>+0(SB), CX
	MOVB (CX)(AX*1), AL
	MOVB AL, ret+8(FP)
	RET
