package internal

import (
	"math"
	"unsafe"
)

func f371(ctx *Context, l0 float32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 float64
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
	var s2f32 float32
	_ = s2f32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0f32 = l0
	s0i32 = int32(math.Float32bits(s0f32))
	l4 = s0i32
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l2 = s0i32
	s1i32 = 1305022426
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1f32 = l0
		s1f64 = float64(s1f32)
		l5 = s1f64
		s2f64 = l5
		s3f64 = 0.6366197723675814
		s2f64 = s2f64 * s3f64
		s3f64 = 6.755399441055744e+15
		s2f64 = s2f64 + s3f64
		s3f64 = -6.755399441055744e+15
		s2f64 = s2f64 + s3f64
		l5 = s2f64
		s3f64 = -1.5707963109016418
		s2f64 = s2f64 * s3f64
		s1f64 = s1f64 + s2f64
		s2f64 = l5
		s3f64 = -1.5893254773528196e-08
		s2f64 = s2f64 * s3f64
		s1f64 = s1f64 + s2f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0f64 = l5
		s0f64 = math.Abs(s0f64)
		s1f64 = 2.147483648e+09
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f64 = l5
			s0i32 = int32(math.Trunc(s0f64))
			l2 = s0i32
			goto lbl0
		}
		s0i32 = -2147483648
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 2139095040
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1f32 = l0
		s2f32 = l0
		s1f32 = s1f32 - s2f32
		s1f64 = float64(s1f32)
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = 0
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l2
	s2i32 = l2
	s3i32 = 23
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = -150
	s2i32 = s2i32 + s3i32
	l2 = s2i32
	s3i32 = 23
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 - s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s1f64 = float64(s1f32)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = l2
	s3i32 = 1
	s4i32 = 0
	s0i32 = f538(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l2 = s0i32
	s0i32 = l3
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0f64
	s0i32 = l4
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1f64 = l5
		s1f64 = -s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = 0
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l1
	s1f64 = l5
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
lbl0:
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l2
	return s0i32
}
