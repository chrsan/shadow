package internal

import (
	"math"
	"unsafe"
)

func f1406(ctx *Context, l0 int32, l1 int32, l2 float32, l3 int32, l4 int32, l5 int32) int32 {
	var l6 int32
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
	var s5f32 float32
	_ = s5f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l8 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l9 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l10 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l11 = s0f32
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
	s0f32 = l7
	s1f32 = l10
	s0f32 = s0f32 * s1f32
	s1f32 = l8
	s2f32 = l9
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l7 = s0f32
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
	s0f32 = l7
	s1f32 = l11
	s0f32 = f298(ctx, s0f32, s1f32)
	l7 = s0f32
	s1f32 = l2
	s0f32 = s0f32 * s1f32
	s1f32 = 0.25
	s0f32 = s0f32 * s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l2 = s0f32
	s1f32 = 65535
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 1
	l6 = s0i32
	s0i32 = l3
	s1f32 = 0
	s2f32 = l2
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	l2 = s2f32
	s3f32 = 2.1474835e+09
	s4f32 = l2
	s5f32 = 2.1474835e+09
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l2 = s2f32
	s3f32 = -2.1474835e+09
	s4f32 = l2
	s5f32 = -2.1474835e+09
	if s4f32 > s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l2 = s2f32
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 2.1474836e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		s2f32 = l2
		s2i32 = int32(math.Trunc(float64(s2f32)))
		goto lbl2
	}
	s2i32 = -2147483648
lbl2:
	l0 = s2i32
	s3i32 = 1
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl1
	}
	s1f32 = l7
	s2i32 = l0
	s2f32 = float32(s2i32)
	s1f32 = s1f32 / s2f32
lbl1:
	l2 = s1f32
	s1f32 = f107(ctx, s1f32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l4
	s1f32 = l2
	s1f32 = f86(ctx, s1f32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l5
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl0:
	s0i32 = l6
	return s0i32
}
