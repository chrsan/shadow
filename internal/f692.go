package internal

import (
	"unsafe"
)

func f692(ctx *Context, l0 int32) float32 {
	var l1 float32
	_ = l1
	var l2 float32
	_ = l2
	var l3 float32
	_ = l3
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
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0f32 = 0
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l2 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l1 = s2f32
	s1f32 = s1f32 - s2f32
	s2f32 = l1
	s3f32 = l2
	s2f32 = s2f32 - s3f32
	s3f32 = l2
	s2f32 = s2f32 - s3f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	s2f32 = s2f32 + s3f32
	l2 = s2f32
	s1f32 = s1f32 * s2f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l1 = s2f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	l3 = s3f32
	s2f32 = s2f32 - s3f32
	s3f32 = l3
	s4f32 = l1
	s3f32 = s3f32 - s4f32
	s4f32 = l1
	s3f32 = s3f32 - s4f32
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
	s3f32 = s3f32 + s4f32
	l1 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l3 = s1f32
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = 1
	s1f32 = l2
	s2f32 = l2
	s1f32 = s1f32 * s2f32
	s2f32 = l1
	s3f32 = l1
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l2 = s1f32
	s2f32 = l3
	s2f32 = -s2f32
	l1 = s2f32
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = l1
	s1f32 = l2
	s0f32 = s0f32 / s1f32
lbl0:
	return s0f32
}
