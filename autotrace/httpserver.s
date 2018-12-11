#include "go_asm.h"
#include "textflag.h"

// func ptrTestHookServerServe() unsafe.Pointer
TEXT ·ptrTestHookServerServe(SB),NOSPLIT,$0-8
        LEAQ net∕http·testHookServerServe(SB), AX
        MOVQ AX, r0+8(SP)
        RET
