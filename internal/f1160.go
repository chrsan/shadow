package internal

import (
	"unsafe"
)

func f1160(ctx *Context, l0 int32) float32 {
	var s0i32 int32
	_ = s0i32
	var s0f32 float32
	_ = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	return s0f32
}
