package internal

import (
	"math"
	"unsafe"
)

func f327(ctx *Context, l0 int32, l1 float32, l2 float32) int32 {
	var l3 int32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float64
	_ = l6
	var l7 float64
	_ = l7
	var l8 float64
	_ = l8
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
	var s4f32 float32
	_ = s4f32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	s0f64 = 1
	s1f32 = l1
	s1f64 = float64(s1f32)
	l6 = s1f64
	s2f64 = l6
	s1f64 = s1f64 * s2f64
	s2f32 = l2
	s2f64 = float64(s2f32)
	l7 = s2f64
	s3f64 = l7
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s1f64 = math.Sqrt(s1f64)
	s0f64 = s0f64 / s1f64
	l8 = s0f64
	s1f64 = l6
	s0f64 = s0f64 * s1f64
	s0f32 = float32(s0f64)
	l2 = s0f32
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
		s0f32 = 0
		l1 = s0f32
		goto lbl0
	}
	s0f32 = 0
	l1 = s0f32
	s0f64 = l8
	s1f64 = l7
	s0f64 = s0f64 * s1f64
	s0f32 = float32(s0f64)
	l5 = s0f32
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
		goto lbl0
	}
	s0f32 = 0
	s1f32 = l5
	s2f32 = l2
	s3f32 = 0
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3f32 = l5
	s4f32 = 0
	if s3f32 == s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	l3 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l1 = s0f32
	s0f32 = 0
	s1f32 = l2
	s2i32 = l3
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l4 = s0f32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	l3 = s0i32
lbl0:
	s0i32 = l0
	s1f32 = l1
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l0
	s1f32 = l4
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l3
	return s0i32
}
