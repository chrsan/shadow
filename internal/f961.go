package internal

import (
	"unsafe"
)

func f961(ctx *Context, l0 int32, l1 float32, l2 float32, l3 int32) {
	var s0i32 int32
	_ = s0i32
	var s1f32 float32
	_ = s1f32
	s0i32 = l3
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l3
	s1f32 = l1
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
}
