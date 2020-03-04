package internal

import (
	"unsafe"
)

func f1415(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	s0i32 = 27292
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	return s0i32
}
