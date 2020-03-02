package internal

import (
	"unsafe"
)

func f1075(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0f32 = 0.5
	l4 = s0f32
	s0f32 = 0.5
	s1i32 = l1
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = 0
	l4 = s0f32
	s0f32 = 0
	s1i32 = l2
	s2i32 = 2
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l3 = s0i32
	s1i32 = 3380
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0f32
	s0i32 = l3
	s1i32 = 3368
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
lbl0:
	l5 = s0f32
	s0i32 = l0
	s1f32 = l4
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l0
	s1f32 = l5
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 786432
	s2i32 = 0
	s3i32 = l2
	s4i32 = 1
	if s3i32 != s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4i32 = l1
	s3i32 = s3i32 & s4i32
	l3 = s3i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 3
	s2i32 = 0
	s3i32 = l2
	s4i32 = 2
	if s3i32 != s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4i32 = l1
	s3i32 = s3i32 & s4i32
	l1 = s3i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	s2i32 = l3
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	s2i32 = l1
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
}
