package internal

import (
	"unsafe"
)

func f299(ctx *Context, l0 int32, l1 int32) int32 {
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
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float64
	_ = l12
	var l13 float64
	_ = l13
	var l14 float64
	_ = l14
	var l15 float64
	_ = l15
	var l16 float64
	_ = l16
	var l17 float64
	_ = l17
	var l18 float64
	_ = l18
	var l19 float64
	_ = l19
	var l20 float64
	_ = l20
	var l21 float64
	_ = l21
	var l22 float64
	_ = l22
	var l23 float64
	_ = l23
	var l24 float64
	_ = l24
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
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s0f64 = float64(s0f32)
	l15 = s0f64
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s1f64 = float64(s1f32)
	l16 = s1f64
	s0f64 = s0f64 * s1f64
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s1f64 = float64(s1f32)
	l17 = s1f64
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	s2f64 = float64(s2f32)
	l18 = s2f64
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 - s1f64
	l13 = s0f64
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s1f64 = float64(s1f32)
	l21 = s1f64
	s0f64 = s0f64 * s1f64
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1f64 = float64(s1f32)
	l19 = s1f64
	s2f64 = l18
	s1f64 = s1f64 * s2f64
	s2f64 = l15
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s3f64 = float64(s3f32)
	l20 = s3f64
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	l14 = s1f64
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	s2f64 = float64(s2f32)
	l22 = s2f64
	s1f64 = s1f64 * s2f64
	s2f64 = l19
	s3f64 = l16
	s2f64 = s2f64 * s3f64
	s3f64 = l17
	s4f64 = l20
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 - s3f64
	l23 = s2f64
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
	s3f64 = float64(s3f32)
	l24 = s3f64
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s0f64 = s0f64 + s1f64
	l12 = s0f64
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f64 = 1
	s1f64 = l12
	s0f64 = s0f64 / s1f64
	l12 = s0f64
	s1f64 = 3.4028234663852886e+38
	if s0f64 > s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f64 = l12
	s1f64 = -3.4028234663852886e+38
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f64 = l12
	s0f32 = float32(s0f64)
	s1f32 = 0
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
	s0i32 = l1
	s1f64 = l13
	s2f64 = l12
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l3 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l1
	s1f64 = l14
	s2f64 = l12
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l4 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l1
	s1f64 = l12
	s2f64 = l23
	s2f64 = -s2f64
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l5 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l1
	s1f64 = l12
	s2f64 = l24
	s1f64 = s1f64 * s2f64
	l13 = s1f64
	s2f64 = l17
	s1f64 = s1f64 * s2f64
	s2f64 = l12
	s3f64 = l22
	s2f64 = s2f64 * s3f64
	l14 = s2f64
	s3f64 = l15
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s1f32 = float32(s1f64)
	l6 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l1
	s1f64 = l14
	s2f64 = l18
	s1f64 = s1f64 * s2f64
	s2f64 = l13
	s3f64 = l16
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s1f32 = float32(s1f64)
	l7 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l1
	s1f64 = l14
	s2f64 = l19
	s1f64 = s1f64 * s2f64
	s2f64 = l12
	s3f64 = l21
	s2f64 = s2f64 * s3f64
	l12 = s2f64
	s3f64 = l17
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s1f32 = float32(s1f64)
	l8 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l1
	s1f64 = l12
	s2f64 = l16
	s1f64 = s1f64 * s2f64
	s2f64 = l14
	s3f64 = l20
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s1f32 = float32(s1f64)
	l9 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l1
	s1f64 = l12
	s2f64 = l15
	s1f64 = s1f64 * s2f64
	s2f64 = l13
	s3f64 = l19
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s1f32 = float32(s1f64)
	l10 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l1
	s1f64 = l13
	s2f64 = l20
	s1f64 = s1f64 * s2f64
	s2f64 = l12
	s3f64 = l18
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s1f32 = float32(s1f64)
	l11 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0f32 = l7
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l9
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1f32 = l11
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1f32 = l6
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1f32 = l8
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1f32 = l10
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1f32 = l3
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1f32 = l5
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1f32 = l4
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	l2 = s0i32
lbl0:
	s0i32 = l2
	return s0i32
}
