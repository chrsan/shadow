package internal

import (
	"math"
	"unsafe"
)

func f780(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
	var l23 float32
	_ = l23
	var l24 float32
	_ = l24
	var l25 float32
	_ = l25
	var l26 float32
	_ = l26
	var l27 float32
	_ = l27
	var l28 float32
	_ = l28
	var l29 float32
	_ = l29
	var l30 float32
	_ = l30
	var l31 float32
	_ = l31
	var l32 float32
	_ = l32
	var l33 float32
	_ = l33
	var l34 float32
	_ = l34
	var l35 float32
	_ = l35
	var l36 float32
	_ = l36
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s6i32 int32
	_ = s6i32
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
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l19 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l5 = s0f32
	s0f32 = l2
	s1f32 = -0.5
	s0f32 = s0f32 + s1f32
	l18 = s0f32
	l20 = s0f32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l14 = s0i32
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl2
	case 1:
		goto lbl1
	default:
		goto lbl0
	}
lbl2:
	s0f32 = l18
	s1f32 = l5
	s2f32 = l18
	s3f32 = l19
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l20 = s0f32
	goto lbl0
lbl1:
	s0f32 = l18
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l5
	s2f32 = l5
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l19
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l20 = s0f32
lbl0:
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l21 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l16 = s0f32
	s0f32 = l3
	s1f32 = -0.5
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	l17 = s0f32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l15 = s0i32
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl5
	case 1:
		goto lbl4
	default:
		goto lbl3
	}
lbl5:
	s0f32 = l4
	s1f32 = l16
	s2f32 = l4
	s3f32 = l21
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l17 = s0f32
	goto lbl3
lbl4:
	s0f32 = l4
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	l17 = s0f32
	s1f32 = l16
	s2f32 = l16
	s1f32 = s1f32 + s2f32
	s2f32 = l17
	s3f32 = l21
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l17 = s0f32
lbl3:
	s0f32 = l5
	s0i32 = int32(math.Float32bits(s0f32))
	l10 = s0i32
	s0f32 = l16
	s0i32 = int32(math.Float32bits(s0f32))
	l11 = s0i32
	s0f32 = 0
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l9 = s1i32
	s2i32 = -4
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s2i32 = 2
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl6
	}
	s0f32 = 0
	s1i32 = l6
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl6
	}
	s0f32 = l17
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l11
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l17 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l17
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l17
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl7
	}
	s0i32 = 0
lbl7:
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l8 = s0i32
	s0f32 = l20
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l10
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l17 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l17
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l17
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl9
	}
	s0i32 = 0
lbl9:
	l6 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l6
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l8 = s0i32
	s1i32 = l6
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l12 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l13 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l22 = s0f32
	s0i32 = l6
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l23 = s0f32
	s0i32 = l6
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l25 = s0f32
	s0i32 = l12
	s1i32 = l8
	s2i32 = l13
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
lbl6:
	l24 = s0f32
	s0f32 = l3
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l20 = s0f32
	s0f32 = l2
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l27 = s0f32
	s0f32 = l18
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l17 = s0f32
	l3 = s0f32
	s0i32 = l14
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl12
	case 1:
		goto lbl13
	default:
		goto lbl11
	}
lbl13:
	s0f32 = l17
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	l2 = s0f32
	s1f32 = l5
	s2f32 = l5
	s1f32 = s1f32 + s2f32
	s2f32 = l2
	s3f32 = l19
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l3 = s0f32
	goto lbl11
lbl12:
	s0f32 = l17
	s1f32 = l5
	s2f32 = l17
	s3f32 = l19
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l3 = s0f32
lbl11:
	s0f32 = l20
	s0f32 = float32(math.Floor(float64(s0f32)))
	l28 = s0f32
	s0f32 = l27
	s0f32 = float32(math.Floor(float64(s0f32)))
	l26 = s0f32
	s0f32 = l4
	l2 = s0f32
	s0i32 = l15
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl15
	case 1:
		goto lbl16
	default:
		goto lbl14
	}
lbl16:
	s0f32 = l4
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	l2 = s0f32
	s1f32 = l16
	s2f32 = l16
	s1f32 = s1f32 + s2f32
	s2f32 = l2
	s3f32 = l21
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l2 = s0f32
	goto lbl14
lbl15:
	s0f32 = l4
	s1f32 = l16
	s2f32 = l4
	s3f32 = l21
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l2 = s0f32
lbl14:
	s0f32 = l20
	s1f32 = l28
	s0f32 = s0f32 - s1f32
	l20 = s0f32
	s0f32 = l27
	s1f32 = l26
	s0f32 = s0f32 - s1f32
	l27 = s0f32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		l26 = s0f32
		s0f32 = 0
		goto lbl17
	}
	s0f32 = 0
	l26 = s0f32
	s0f32 = 0
	s1i32 = l6
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl17
	}
	s0f32 = l2
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l11
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l2 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l2
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l2
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl19
	}
	s0i32 = 0
lbl19:
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l8 = s0i32
	s0f32 = l3
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l10
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l2 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l2
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l2
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl21
	}
	s0i32 = 0
lbl21:
	l6 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l6
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l8 = s0i32
	s1i32 = l6
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l12 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l13 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l26 = s0f32
	s0i32 = l12
	s1i32 = l8
	s2i32 = l13
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l30 = s0f32
	s0i32 = l6
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l29 = s0f32
	s0i32 = l6
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl17:
	l28 = s0f32
	s0f32 = 1
	s1f32 = l20
	s0f32 = s0f32 - s1f32
	l31 = s0f32
	s0f32 = 1
	s1f32 = l27
	s0f32 = s0f32 - s1f32
	l33 = s0f32
	s0i32 = l14
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl24
	case 1:
		goto lbl25
	default:
		goto lbl23
	}
lbl25:
	s0f32 = l18
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	l2 = s0f32
	s1f32 = l5
	s2f32 = l5
	s1f32 = s1f32 + s2f32
	s2f32 = l2
	s3f32 = l19
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l18 = s0f32
	goto lbl23
lbl24:
	s0f32 = l18
	s1f32 = l5
	s2f32 = l18
	s3f32 = l19
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l18 = s0f32
lbl23:
	s0f32 = l33
	s1f32 = l31
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0f32 = l4
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l3 = s0f32
	l4 = s0f32
	s0i32 = l15
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl27
	case 1:
		goto lbl28
	default:
		goto lbl26
	}
lbl28:
	s0f32 = l3
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l16
	s2f32 = l16
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l21
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l4 = s0f32
	goto lbl26
lbl27:
	s0f32 = l3
	s1f32 = l16
	s2f32 = l3
	s3f32 = l21
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
lbl26:
	s0f32 = l2
	s1f32 = l23
	s0f32 = s0f32 * s1f32
	l32 = s0f32
	s0f32 = l2
	s1f32 = l22
	s0f32 = s0f32 * s1f32
	l34 = s0f32
	s0f32 = l2
	s1f32 = l25
	s0f32 = s0f32 * s1f32
	l35 = s0f32
	s0f32 = l2
	s1f32 = l24
	s0f32 = s0f32 * s1f32
	l36 = s0f32
	s0f32 = l27
	s1f32 = l31
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0f32 = 0
	l24 = s0f32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		l22 = s0f32
		s0f32 = 0
		l23 = s0f32
		s0f32 = 0
		goto lbl29
	}
	s0f32 = 0
	l22 = s0f32
	s0f32 = 0
	l23 = s0f32
	s0f32 = 0
	s1i32 = l6
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl29
	}
	s0f32 = l4
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l11
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l4 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l4
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l4
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl31
	}
	s0i32 = 0
lbl31:
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l8 = s0i32
	s0f32 = l18
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l10
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l4 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l4
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l4
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl33
	}
	s0i32 = 0
lbl33:
	l6 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l6
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l8 = s0i32
	s1i32 = l6
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l12 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l13 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l22 = s0f32
	s0i32 = l12
	s1i32 = l8
	s2i32 = l13
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l24 = s0f32
	s0i32 = l6
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l23 = s0f32
	s0i32 = l6
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl29:
	l25 = s0f32
	s0f32 = l32
	s1f32 = 0
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s0f32 = l34
	s1f32 = 0
	s0f32 = s0f32 + s1f32
	l18 = s0f32
	s0f32 = l35
	s1f32 = 0
	s0f32 = s0f32 + s1f32
	l31 = s0f32
	s0f32 = l36
	s1f32 = 0
	s0f32 = s0f32 + s1f32
	l32 = s0f32
	s0f32 = l2
	s1f32 = l29
	s0f32 = s0f32 * s1f32
	l29 = s0f32
	s0f32 = l2
	s1f32 = l26
	s0f32 = s0f32 * s1f32
	l26 = s0f32
	s0f32 = l2
	s1f32 = l28
	s0f32 = s0f32 * s1f32
	l28 = s0f32
	s0f32 = l2
	s1f32 = l30
	s0f32 = s0f32 * s1f32
	l30 = s0f32
	s0f32 = l33
	s1f32 = l20
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0i32 = l14
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl36
	case 1:
		goto lbl37
	default:
		goto lbl35
	}
lbl37:
	s0f32 = l17
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	l17 = s0f32
	s1f32 = l5
	s2f32 = l5
	s1f32 = s1f32 + s2f32
	s2f32 = l17
	s3f32 = l19
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l17 = s0f32
	goto lbl35
lbl36:
	s0f32 = l17
	s1f32 = l5
	s2f32 = l17
	s3f32 = l19
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l17 = s0f32
lbl35:
	s0f32 = l4
	s1f32 = l29
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s0f32 = l18
	s1f32 = l26
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s0f32 = l31
	s1f32 = l28
	s0f32 = s0f32 + s1f32
	l18 = s0f32
	s0f32 = l32
	s1f32 = l30
	s0f32 = s0f32 + s1f32
	l19 = s0f32
	s0f32 = l2
	s1f32 = l23
	s0f32 = s0f32 * s1f32
	l23 = s0f32
	s0f32 = l2
	s1f32 = l22
	s0f32 = s0f32 * s1f32
	l22 = s0f32
	s0f32 = l2
	s1f32 = l25
	s0f32 = s0f32 * s1f32
	l25 = s0f32
	s0f32 = l2
	s1f32 = l24
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0i32 = l15
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl39
	case 1:
		goto lbl40
	default:
		goto lbl38
	}
lbl40:
	s0f32 = l3
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	l3 = s0f32
	s1f32 = l16
	s2f32 = l16
	s1f32 = s1f32 + s2f32
	s2f32 = l3
	s3f32 = l21
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l3 = s0f32
	goto lbl38
lbl39:
	s0f32 = l3
	s1f32 = l16
	s2f32 = l3
	s3f32 = l21
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l3 = s0f32
lbl38:
	s0f32 = l4
	s1f32 = l23
	s0f32 = s0f32 + s1f32
	l16 = s0f32
	s0f32 = l5
	s1f32 = l22
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s0f32 = l18
	s1f32 = l25
	s0f32 = s0f32 + s1f32
	l21 = s0f32
	s0f32 = l19
	s1f32 = l2
	s0f32 = s0f32 + s1f32
	l19 = s0f32
	s0f32 = 0
	l4 = s0f32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		l18 = s0f32
		s0f32 = 0
		l24 = s0f32
		s0f32 = 0
		goto lbl41
	}
	s0f32 = 0
	l18 = s0f32
	s0f32 = 0
	l24 = s0f32
	s0f32 = 0
	s1i32 = l6
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl41
	}
	s0f32 = l3
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l11
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l2 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l2
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l2
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl43
	}
	s0i32 = 0
lbl43:
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l11 = s0i32
	s0f32 = l17
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l10
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l2 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l2
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l2
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl45
	}
	s0i32 = 0
lbl45:
	l6 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l6
	s2i32 = l11
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l6 = s0i32
	s1i32 = l7
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l10 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l9 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l18 = s0f32
	s0i32 = l10
	s1i32 = l6
	s2i32 = l9
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l4 = s0f32
	s0i32 = l7
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l24 = s0f32
	s0i32 = l7
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl41:
	l2 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2f32 = l19
	s3f32 = l27
	s4f32 = l20
	s3f32 = s3f32 * s4f32
	l3 = s3f32
	s4f32 = l4
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l21
	s4f32 = l3
	s5f32 = l2
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l5
	s5f32 = l3
	s6f32 = l18
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s5f32 = l16
	s6f32 = l3
	s7f32 = l24
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 + s6f32
	s6i32 = l1
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, float32, float32, float32, float32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2f32, s3f32, s4f32, s5f32)
}
