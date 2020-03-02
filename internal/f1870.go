package internal

import (
	"unsafe"
)

func f1870(ctx *Context, l0 float32) float32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 float32
	_ = l3
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
	var s3f32 float32
	_ = s3f32
lbl3:
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 5192
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = l0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	goto lbl1
lbl2:
	s0i32 = l1
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = 5204
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	return s0f32
lbl0:
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 5200
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		return s0f32
	}
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l1 = s0i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 5200
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0f32
	s1f32 = l0
	s2i32 = l2
	s3i32 = 5192
	s2i32 = s2i32 + s3i32
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l0 = s2f32
	s1f32 = s1f32 - s2f32
	s2i32 = l1
	s3i32 = 5192
	s2i32 = s2i32 + s3i32
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3f32 = l0
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	s2i32 = l1
	s3i32 = 5200
	s2i32 = s2i32 + s3i32
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3f32 = l3
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	return s0f32
}
