package internal

import (
	"unsafe"
)

func f1574(ctx *Context, l0 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s1i32 = 1
	ctx.Mem[int(s0i32+52)] = uint8(s1i32)
}
