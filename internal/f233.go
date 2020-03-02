package internal

import (
	"math"
	"unsafe"
)

func f233(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 float32
	_ = l3
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
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
	var s5f32 float32
	_ = s5f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0f32
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l3 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l5 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l8 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l9 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1f32 = -1
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s1f32 = l4
	s2f32 = 2
	s1f32 = s1f32 + s2f32
	s2f32 = 4
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 / s1f32
	l4 = s0f32
	s1f32 = l8
	s2f32 = l6
	s3f32 = l3
	s4f32 = l3
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	s1f32 = l4
	s2f32 = l9
	s3f32 = l7
	s4f32 = l5
	s5f32 = l5
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	l3 = s1f32
	s2f32 = l3
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Sqrt(float64(s0f32)))
	l3 = s0f32
	s1f32 = 0.25
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
lbl2:
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0f32 = l3
	s1f32 = 0.25
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s1f32 = 0.25
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 4
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l2
	l1 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl2
	}
lbl0:
	s0i32 = l2
	return s0i32
}
