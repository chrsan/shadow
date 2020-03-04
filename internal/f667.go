package internal

import (
	"math"
	"unsafe"
)

func f667(ctx *Context, l0 int32, l1 int32, l2 int32) float32 {
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
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l3 = s1f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l9 = s1f32
	s2f32 = l3
	s1f32 = s1f32 - s2f32
	l3 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l10 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l5 = s2f32
	s1f32 = s1f32 - s2f32
	l6 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l11 = s2f32
	s3f32 = l5
	s2f32 = s2f32 - s3f32
	l5 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l7 = s0f32
	s1f32 = 0
	if s0f32 <= s1f32 {
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
		s0f32 = l3
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = l5
		s2f32 = l5
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		return s0f32
	}
	s0f32 = l7
	s1f32 = l4
	s2f32 = l4
	s1f32 = s1f32 * s2f32
	s2f32 = l6
	s3f32 = l6
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l7 = s1f32
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
		s0f32 = l8
		s1f32 = l9
		s0f32 = s0f32 - s1f32
		l3 = s0f32
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = l10
		s2f32 = l11
		s1f32 = s1f32 - s2f32
		l3 = s1f32
		s2f32 = l3
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		return s0f32
	}
	s0f32 = l4
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	s1f32 = l6
	s2f32 = l3
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l4
	s2f32 = l7
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 * s1f32
	l4 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l3
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = l5
		s2f32 = l5
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
	} else {
		s0f32 = l4
	}
	return s0f32
}
