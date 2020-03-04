package internal

import (
	"unsafe"
)

func f678(ctx *Context, l0 int32, l1 float32, l2 float32, l3 float32) {
	var l4 int32
	_ = l4
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0f32 = l3
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1f32 = l2
		s2f32 = l3
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l4
		s1f32 = l1
		s2f32 = l3
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l4
		s1f32 = l2
		s2f32 = l3
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l4
		s1f32 = l1
		s2f32 = l3
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s1i32 = l4
		s2i32 = 0
		s3i32 = 1
		f330(ctx, s0i32, s1i32, s2i32, s3i32)
	}
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
