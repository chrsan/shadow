package internal

import (
	"unsafe"
)

func f971(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
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
	s0i32 = l3
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l6 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l8 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l11 = s0f32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l1
		s1f32 = l10
		s2f32 = l6
		s3i32 = l2
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		l4 = s3f32
		s2f32 = s2f32 * s3f32
		s3f32 = l8
		s4i32 = l2
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		l5 = s4f32
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l11
		s2f32 = l9
		s3f32 = l4
		s2f32 = s2f32 * s3f32
		s3f32 = l7
		s4f32 = l5
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l2
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l1 = s0i32
	}
	s0i32 = l3
	s1i32 = 1
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	l0 = s0i32
lbl2:
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l4 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l5 = s0f32
	s0i32 = l1
	s1f32 = l10
	s2f32 = l6
	s3i32 = l2
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l12 = s3f32
	s2f32 = s2f32 * s3f32
	s3f32 = l8
	s4i32 = l2
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	l13 = s4f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l1
	s1f32 = l11
	s2f32 = l9
	s3f32 = l12
	s2f32 = s2f32 * s3f32
	s3f32 = l7
	s4f32 = l13
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l1
	s1f32 = l10
	s2f32 = l6
	s3f32 = l5
	s2f32 = s2f32 * s3f32
	s3f32 = l8
	s4f32 = l4
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l1
	s1f32 = l11
	s2f32 = l9
	s3f32 = l5
	s2f32 = s2f32 * s3f32
	s3f32 = l7
	s4f32 = l4
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i32 = l3
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl0:
}
