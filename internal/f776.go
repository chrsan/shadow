package internal

import (
	"math"
	"unsafe"
)

func f776(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
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
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 int32
	_ = l25
	var l26 int32
	_ = l26
	var l27 int32
	_ = l27
	var l28 int32
	_ = l28
	var l29 int32
	_ = l29
	var l30 int32
	_ = l30
	var l31 int32
	_ = l31
	var l32 int32
	_ = l32
	var l33 int32
	_ = l33
	var l34 int32
	_ = l34
	var l35 float32
	_ = l35
	var l36 float32
	_ = l36
	var l37 float32
	_ = l37
	var l38 float32
	_ = l38
	var l39 float32
	_ = l39
	var l40 float32
	_ = l40
	var l41 float32
	_ = l41
	var l42 float32
	_ = l42
	var l43 float32
	_ = l43
	var l44 float32
	_ = l44
	var l45 float32
	_ = l45
	var l46 float32
	_ = l46
	var l47 float32
	_ = l47
	var l48 float32
	_ = l48
	var l49 float32
	_ = l49
	var l50 float32
	_ = l50
	var l51 float32
	_ = l51
	var l52 float32
	_ = l52
	var l53 float32
	_ = l53
	var l54 float32
	_ = l54
	var l55 float32
	_ = l55
	var l56 float32
	_ = l56
	var l57 float32
	_ = l57
	var l58 float32
	_ = l58
	var l59 float32
	_ = l59
	var l60 float32
	_ = l60
	var l61 float32
	_ = l61
	var l62 float32
	_ = l62
	var l63 float32
	_ = l63
	var l64 float32
	_ = l64
	var l65 float32
	_ = l65
	var l66 float32
	_ = l66
	var l67 float32
	_ = l67
	var l68 float32
	_ = l68
	var l69 float32
	_ = l69
	var l70 float32
	_ = l70
	var l71 float32
	_ = l71
	var l72 float32
	_ = l72
	var l73 float32
	_ = l73
	var l74 float32
	_ = l74
	var l75 float32
	_ = l75
	var l76 float32
	_ = l76
	var l77 float32
	_ = l77
	var l78 float32
	_ = l78
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
	l6 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l42 = s0f32
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l35 = s0f32
	s0f32 = l2
	s1f32 = -1.5
	s0f32 = s0f32 + s1f32
	l38 = s0f32
	l44 = s0f32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l16 = s0i32
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
	s0f32 = l38
	s1f32 = l35
	s2f32 = l38
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l44 = s0f32
	goto lbl0
lbl1:
	s0f32 = l38
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l44 = s0f32
lbl0:
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l43 = s0f32
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l36 = s0f32
	s0f32 = l3
	s1f32 = -1.5
	s0f32 = s0f32 + s1f32
	l37 = s0f32
	l40 = s0f32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l11 = s0i32
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
	s0f32 = l37
	s1f32 = l36
	s2f32 = l37
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l40 = s0f32
	goto lbl3
lbl4:
	s0f32 = l37
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l40 = s0f32
lbl3:
	s0f32 = l35
	s0i32 = int32(math.Float32bits(s0f32))
	l20 = s0i32
	s0f32 = l36
	s0i32 = int32(math.Float32bits(s0f32))
	l21 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l9 = s0i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		goto lbl6
	}
	s0f32 = 0
	s1i32 = l7
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
	s0f32 = l40
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
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
		goto lbl8
	}
	s0i32 = 0
lbl8:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l12 = s0i32
	s0f32 = l44
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl10
	}
	s0i32 = 0
lbl10:
	l7 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l7
	s2i32 = l12
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
	l12 = s0i32
	s1i32 = l7
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l8 = s1i32
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
	l23 = s0i32
	s0i32 = l8
	s1i32 = l12
	s2i32 = l13
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l12 = s0i32
	s0i32 = l7
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l66 = s0f32
	s0i32 = l7
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl6:
	l59 = s0f32
	s0f32 = l38
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l44 = s0f32
	l40 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl16
	case 1:
		goto lbl17
	default:
		goto lbl15
	}
lbl17:
	s0f32 = l44
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l40 = s0f32
	s0f32 = l37
	l45 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl13
	case 1:
		goto lbl14
	default:
		goto lbl12
	}
lbl16:
	s0f32 = l44
	s1f32 = l35
	s2f32 = l44
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l40 = s0f32
lbl15:
	s0f32 = l37
	l45 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl13
	case 1:
		goto lbl14
	default:
		goto lbl12
	}
lbl14:
	s0f32 = l37
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l45 = s0f32
	goto lbl12
lbl13:
	s0f32 = l37
	s1f32 = l36
	s2f32 = l37
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l45 = s0f32
lbl12:
	s0i32 = 0
	l7 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		goto lbl18
	}
	s0f32 = 0
	s1i32 = l8
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl18
	}
	s0f32 = l45
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
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
		goto lbl20
	}
	s0i32 = 0
lbl20:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l7 = s0i32
	s0f32 = l40
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl22
	}
	s0i32 = 0
lbl22:
	l8 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l7
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l7 = s0i32
	s1i32 = l8
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l13 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l10 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l24 = s0i32
	s0i32 = l13
	s1i32 = l7
	s2i32 = l10
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l7 = s0i32
	s0i32 = l8
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l47 = s0f32
	s0i32 = l8
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl18:
	l41 = s0f32
	s0f32 = l44
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l40 = s0f32
	l45 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl28
	case 1:
		goto lbl29
	default:
		goto lbl27
	}
lbl29:
	s0f32 = l40
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l45 = s0f32
	s0f32 = l37
	l4 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl25
	case 1:
		goto lbl26
	default:
		goto lbl24
	}
lbl28:
	s0f32 = l40
	s1f32 = l35
	s2f32 = l40
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l45 = s0f32
lbl27:
	s0f32 = l37
	l4 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl25
	case 1:
		goto lbl26
	default:
		goto lbl24
	}
lbl26:
	s0f32 = l37
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l4 = s0f32
	goto lbl24
lbl25:
	s0f32 = l37
	s1f32 = l36
	s2f32 = l37
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
lbl24:
	s0i32 = 0
	l8 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l13 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		goto lbl30
	}
	s0f32 = 0
	s1i32 = l13
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl30
	}
	s0f32 = l4
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
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
		goto lbl32
	}
	s0i32 = 0
lbl32:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l8 = s0i32
	s0f32 = l45
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl34
	}
	s0i32 = 0
lbl34:
	l13 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l8
	s2i32 = l13
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l13 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l8 = s0i32
	s1i32 = l13
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
	l14 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l26 = s0i32
	s0i32 = l10
	s1i32 = l8
	s2i32 = l14
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l8 = s0i32
	s0i32 = l13
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l49 = s0f32
	s0i32 = l13
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl30:
	l48 = s0f32
	s0f32 = l40
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l45 = s0f32
	l4 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl40
	case 1:
		goto lbl41
	default:
		goto lbl39
	}
lbl41:
	s0f32 = l45
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l4 = s0f32
	s0f32 = l37
	l5 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl37
	case 1:
		goto lbl38
	default:
		goto lbl36
	}
lbl40:
	s0f32 = l45
	s1f32 = l35
	s2f32 = l45
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
lbl39:
	s0f32 = l37
	l5 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl37
	case 1:
		goto lbl38
	default:
		goto lbl36
	}
lbl38:
	s0f32 = l37
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l5 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l5
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l5 = s0f32
	goto lbl36
lbl37:
	s0f32 = l37
	s1f32 = l36
	s2f32 = l37
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l5 = s0f32
lbl36:
	s0i32 = 0
	l13 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		goto lbl42
	}
	s0f32 = 0
	s1i32 = l10
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl42
	}
	s0f32 = l5
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l5 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l5
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l5
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl44
	}
	s0i32 = 0
lbl44:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l13 = s0i32
	s0f32 = l4
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl46
	}
	s0i32 = 0
lbl46:
	l10 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l10
	s2i32 = l13
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l13 = s0i32
	s1i32 = l10
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l14 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l15 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l27 = s0i32
	s0i32 = l14
	s1i32 = l13
	s2i32 = l15
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l13 = s0i32
	s0i32 = l10
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l56 = s0f32
	s0i32 = l10
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl42:
	l50 = s0f32
	s0f32 = l38
	l4 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl49
	case 1:
		goto lbl50
	default:
		goto lbl48
	}
lbl50:
	s0f32 = l38
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l4 = s0f32
	goto lbl48
lbl49:
	s0f32 = l38
	s1f32 = l35
	s2f32 = l38
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
lbl48:
	s0f32 = l37
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l37 = s0f32
	l5 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl52
	case 1:
		goto lbl53
	default:
		goto lbl51
	}
lbl53:
	s0f32 = l37
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l5 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l5
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l5 = s0f32
	goto lbl51
lbl52:
	s0f32 = l37
	s1f32 = l36
	s2f32 = l37
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l5 = s0f32
lbl51:
	s0i32 = 0
	l10 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l14 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l44
		l4 = s0f32
		s0i32 = l16
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl58
		case 1:
			goto lbl59
		default:
			goto lbl57
		}
	}
	s0i32 = l14
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	if s0i32 != 0 {
		s0f32 = l5
		s1f32 = 0
		s0f32 = f15(ctx, s0f32, s1f32)
		s1i32 = l21
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		s1f32 = math.Float32frombits(uint32(s1i32))
		s0f32 = f14(ctx, s0f32, s1f32)
		l5 = s0f32
		s1f32 = 4.2949673e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1f32 = l5
		s2f32 = 0
		if s1f32 >= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0f32 = l5
			s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
			goto lbl62
		}
		s0i32 = 0
	lbl62:
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 * s1i32
		l10 = s0i32
		s0f32 = l4
		s1f32 = 0
		s0f32 = f15(ctx, s0f32, s1f32)
		s1i32 = l20
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
			goto lbl64
		}
		s0i32 = 0
	lbl64:
		l14 = s0i32
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l10
		s2i32 = l14
		s1i32 = s1i32 + s2i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l14 = s0i32
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		s0f32 = float32(uint32(s0i32))
		s1f32 = 0.003921569
		s0f32 = s0f32 * s1f32
		s0i32 = int32(math.Float32bits(s0f32))
		l10 = s0i32
		s1i32 = l14
		s2i32 = 16
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 255
		s1i32 = s1i32 & s2i32
		s1f32 = float32(uint32(s1i32))
		s2f32 = 0.003921569
		s1f32 = s1f32 * s2f32
		s1i32 = int32(math.Float32bits(s1f32))
		l15 = s1i32
		s2i32 = l9
		s3i32 = 6
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l17 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l28 = s0i32
		s0i32 = l14
		s1i32 = 24
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s0f32 = float32(uint32(s0i32))
		s1f32 = 0.003921569
		s0f32 = s0f32 * s1f32
		l51 = s0f32
		s0i32 = l14
		s1i32 = 8
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		s0f32 = float32(uint32(s0i32))
		s1f32 = 0.003921569
		s0f32 = s0f32 * s1f32
		l60 = s0f32
		s0i32 = l15
		s1i32 = l10
		s2i32 = l17
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l10 = s0i32
	}
	s0f32 = l44
	l4 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl58
	case 1:
		goto lbl59
	default:
		goto lbl57
	}
lbl59:
	s0f32 = l44
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l4 = s0f32
	s0f32 = l37
	l5 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl55
	case 1:
		goto lbl56
	default:
		goto lbl54
	}
lbl58:
	s0f32 = l44
	s1f32 = l35
	s2f32 = l44
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
lbl57:
	s0f32 = l37
	l5 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl55
	case 1:
		goto lbl56
	default:
		goto lbl54
	}
lbl56:
	s0f32 = l37
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l5 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l5
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l5 = s0f32
	goto lbl54
lbl55:
	s0f32 = l37
	s1f32 = l36
	s2f32 = l37
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l5 = s0f32
lbl54:
	s0i32 = 0
	l14 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l15 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l40
		l4 = s0f32
		s0i32 = l16
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl70
		case 1:
			goto lbl71
		default:
			goto lbl69
		}
	}
	s0i32 = l15
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	if s0i32 != 0 {
		s0f32 = l5
		s1f32 = 0
		s0f32 = f15(ctx, s0f32, s1f32)
		s1i32 = l21
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		s1f32 = math.Float32frombits(uint32(s1i32))
		s0f32 = f14(ctx, s0f32, s1f32)
		l5 = s0f32
		s1f32 = 4.2949673e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1f32 = l5
		s2f32 = 0
		if s1f32 >= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0f32 = l5
			s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
			goto lbl74
		}
		s0i32 = 0
	lbl74:
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 * s1i32
		l14 = s0i32
		s0f32 = l4
		s1f32 = 0
		s0f32 = f15(ctx, s0f32, s1f32)
		s1i32 = l20
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
			goto lbl76
		}
		s0i32 = 0
	lbl76:
		l15 = s0i32
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l14
		s2i32 = l15
		s1i32 = s1i32 + s2i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l15 = s0i32
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		s0f32 = float32(uint32(s0i32))
		s1f32 = 0.003921569
		s0f32 = s0f32 * s1f32
		s0i32 = int32(math.Float32bits(s0f32))
		l14 = s0i32
		s1i32 = l15
		s2i32 = 16
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 255
		s1i32 = s1i32 & s2i32
		s1f32 = float32(uint32(s1i32))
		s2f32 = 0.003921569
		s1f32 = s1f32 * s2f32
		s1i32 = int32(math.Float32bits(s1f32))
		l17 = s1i32
		s2i32 = l9
		s3i32 = 6
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l18 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l29 = s0i32
		s0i32 = l15
		s1i32 = 24
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s0f32 = float32(uint32(s0i32))
		s1f32 = 0.003921569
		s0f32 = s0f32 * s1f32
		l54 = s0f32
		s0i32 = l15
		s1i32 = 8
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		s0f32 = float32(uint32(s0i32))
		s1f32 = 0.003921569
		s0f32 = s0f32 * s1f32
		l62 = s0f32
		s0i32 = l17
		s1i32 = l14
		s2i32 = l18
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l14 = s0i32
	}
	s0f32 = l40
	l4 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl70
	case 1:
		goto lbl71
	default:
		goto lbl69
	}
lbl71:
	s0f32 = l40
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l4 = s0f32
	s0f32 = l37
	l5 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl67
	case 1:
		goto lbl68
	default:
		goto lbl66
	}
lbl70:
	s0f32 = l40
	s1f32 = l35
	s2f32 = l40
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
lbl69:
	s0f32 = l37
	l5 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl67
	case 1:
		goto lbl68
	default:
		goto lbl66
	}
lbl68:
	s0f32 = l37
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l5 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l5
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l5 = s0f32
	goto lbl66
lbl67:
	s0f32 = l37
	s1f32 = l36
	s2f32 = l37
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l5 = s0f32
lbl66:
	s0i32 = 0
	l15 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l17 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l45
		l4 = s0f32
		s0i32 = l16
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl79
		case 1:
			goto lbl80
		default:
			goto lbl78
		}
	}
	s0i32 = l17
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	if s0i32 != 0 {
		s0f32 = l5
		s1f32 = 0
		s0f32 = f15(ctx, s0f32, s1f32)
		s1i32 = l21
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		s1f32 = math.Float32frombits(uint32(s1i32))
		s0f32 = f14(ctx, s0f32, s1f32)
		l5 = s0f32
		s1f32 = 4.2949673e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1f32 = l5
		s2f32 = 0
		if s1f32 >= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0f32 = l5
			s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
			goto lbl83
		}
		s0i32 = 0
	lbl83:
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 * s1i32
		l15 = s0i32
		s0f32 = l4
		s1f32 = 0
		s0f32 = f15(ctx, s0f32, s1f32)
		s1i32 = l20
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
			goto lbl85
		}
		s0i32 = 0
	lbl85:
		l17 = s0i32
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l15
		s2i32 = l17
		s1i32 = s1i32 + s2i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l17 = s0i32
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		s0f32 = float32(uint32(s0i32))
		s1f32 = 0.003921569
		s0f32 = s0f32 * s1f32
		s0i32 = int32(math.Float32bits(s0f32))
		l15 = s0i32
		s1i32 = l17
		s2i32 = 16
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 255
		s1i32 = s1i32 & s2i32
		s1f32 = float32(uint32(s1i32))
		s2f32 = 0.003921569
		s1f32 = s1f32 * s2f32
		s1i32 = int32(math.Float32bits(s1f32))
		l18 = s1i32
		s2i32 = l9
		s3i32 = 6
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l19 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l30 = s0i32
		s0i32 = l17
		s1i32 = 24
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s0f32 = float32(uint32(s0i32))
		s1f32 = 0.003921569
		s0f32 = s0f32 * s1f32
		l65 = s0f32
		s0i32 = l17
		s1i32 = 8
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		s0f32 = float32(uint32(s0i32))
		s1f32 = 0.003921569
		s0f32 = s0f32 * s1f32
		l63 = s0f32
		s0i32 = l18
		s1i32 = l15
		s2i32 = l19
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l15 = s0i32
	}
	s0f32 = l45
	l4 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl79
	case 1:
		goto lbl80
	default:
		goto lbl78
	}
lbl80:
	s0f32 = l45
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l4 = s0f32
	goto lbl78
lbl79:
	s0f32 = l45
	s1f32 = l35
	s2f32 = l45
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
lbl78:
	s0f32 = l2
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l2 = s0f32
	s0f32 = l37
	l5 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl88
	case 1:
		goto lbl89
	default:
		goto lbl87
	}
lbl89:
	s0f32 = l37
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l5 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l5
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l5 = s0f32
	goto lbl87
lbl88:
	s0f32 = l37
	s1f32 = l36
	s2f32 = l37
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l5 = s0f32
lbl87:
	s0f32 = l2
	s0f32 = float32(math.Floor(float64(s0f32)))
	l46 = s0f32
	s0i32 = 0
	l17 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l18 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		goto lbl90
	}
	s0f32 = 0
	s1i32 = l18
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl90
	}
	s0f32 = l5
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l5 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l5
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l5
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl92
	}
	s0i32 = 0
lbl92:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l17 = s0i32
	s0f32 = l4
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl94
	}
	s0i32 = 0
lbl94:
	l18 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l17
	s2i32 = l18
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l18 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l17 = s0i32
	s1i32 = l18
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l19 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l22 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l31 = s0i32
	s0i32 = l19
	s1i32 = l17
	s2i32 = l22
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l17 = s0i32
	s0i32 = l18
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l64 = s0f32
	s0i32 = l18
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl90:
	l72 = s0f32
	s0f32 = l3
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s0f32 = l2
	s1f32 = l46
	s0f32 = s0f32 - s1f32
	l57 = s0f32
	s0f32 = l38
	l5 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl97
	case 1:
		goto lbl98
	default:
		goto lbl96
	}
lbl98:
	s0f32 = l38
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l2 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l2
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l5 = s0f32
	goto lbl96
lbl97:
	s0f32 = l38
	s1f32 = l35
	s2f32 = l38
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l5 = s0f32
lbl96:
	s0f32 = l4
	s0f32 = float32(math.Floor(float64(s0f32)))
	l46 = s0f32
	s0f32 = 1
	s1f32 = l57
	s0f32 = s0f32 - s1f32
	l3 = s0f32
	s0f32 = l37
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l2 = s0f32
	l37 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl100
	case 1:
		goto lbl101
	default:
		goto lbl99
	}
lbl101:
	s0f32 = l2
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l37 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l37
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l37 = s0f32
	goto lbl99
lbl100:
	s0f32 = l2
	s1f32 = l36
	s2f32 = l2
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l37 = s0f32
lbl99:
	s0f32 = l4
	s1f32 = l46
	s0f32 = s0f32 - s1f32
	l55 = s0f32
	s0f32 = l3
	s1f32 = 1.1666666
	s0f32 = s0f32 * s1f32
	l46 = s0f32
	s0i32 = 0
	l18 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l19 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		goto lbl102
	}
	s0f32 = 0
	s1i32 = l19
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl102
	}
	s0f32 = l37
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
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
		goto lbl104
	}
	s0i32 = 0
lbl104:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l18 = s0i32
	s0f32 = l5
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl106
	}
	s0i32 = 0
lbl106:
	l19 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l18
	s2i32 = l19
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l19 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l18 = s0i32
	s1i32 = l19
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l22 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l25 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l32 = s0i32
	s0i32 = l22
	s1i32 = l18
	s2i32 = l25
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l18 = s0i32
	s0i32 = l19
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l67 = s0f32
	s0i32 = l19
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl102:
	l76 = s0f32
	s0f32 = 1
	s1f32 = l55
	s0f32 = s0f32 - s1f32
	l5 = s0f32
	s0f32 = 1.5
	s1f32 = l46
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s0f32 = l57
	s1f32 = 1.1666666
	s0f32 = s0f32 * s1f32
	l46 = s0f32
	s0f32 = l44
	l37 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl109
	case 1:
		goto lbl110
	default:
		goto lbl108
	}
lbl110:
	s0f32 = l44
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l37 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l37
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l37 = s0f32
	goto lbl108
lbl109:
	s0f32 = l44
	s1f32 = l35
	s2f32 = l44
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l37 = s0f32
lbl108:
	s0f32 = l5
	s1f32 = 0.3888889
	s0f32 = s0f32 * s1f32
	l39 = s0f32
	s0f32 = l3
	s1f32 = 0.3888889
	s0f32 = s0f32 * s1f32
	l68 = s0f32
	s0f32 = l3
	s1f32 = l4
	s0f32 = s0f32 * s1f32
	l4 = s0f32
	s0f32 = 1.5
	s1f32 = l46
	s0f32 = s0f32 - s1f32
	l73 = s0f32
	s0f32 = l2
	l46 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl112
	case 1:
		goto lbl113
	default:
		goto lbl111
	}
lbl113:
	s0f32 = l2
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l46 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l46
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l46 = s0f32
	goto lbl111
lbl112:
	s0f32 = l2
	s1f32 = l36
	s2f32 = l2
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l46 = s0f32
lbl111:
	s0f32 = l39
	s1f32 = -0.33333334
	s0f32 = s0f32 + s1f32
	l39 = s0f32
	s0f32 = l5
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	l69 = s0f32
	s0f32 = l68
	s1f32 = -0.33333334
	s0f32 = s0f32 + s1f32
	l68 = s0f32
	s0f32 = l3
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	l52 = s0f32
	s0f32 = l4
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l53 = s0f32
	s0f32 = l57
	s1f32 = l73
	s0f32 = s0f32 * s1f32
	l58 = s0f32
	s0f32 = l5
	s1f32 = 1.1666666
	s0f32 = s0f32 * s1f32
	l61 = s0f32
	s0i32 = 0
	l19 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l22 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l25 = s0i32
		s0f32 = 0
		goto lbl114
	}
	s0i32 = 0
	l25 = s0i32
	s0f32 = 0
	s1i32 = l22
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl114
	}
	s0f32 = l46
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
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
		goto lbl116
	}
	s0i32 = 0
lbl116:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l19 = s0i32
	s0f32 = l37
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl118
	}
	s0i32 = 0
lbl118:
	l22 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l19
	s2i32 = l22
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l22 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l19 = s0i32
	s1i32 = l22
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l33 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l34 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l25 = s0i32
	s0i32 = l33
	s1i32 = l19
	s2i32 = l34
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l19 = s0i32
	s0i32 = l22
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l75 = s0f32
	s0i32 = l22
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl114:
	l73 = s0f32
	s0f32 = l69
	s1f32 = l39
	s0f32 = s0f32 * s1f32
	l37 = s0f32
	s0f32 = l52
	s1f32 = l68
	s0f32 = s0f32 * s1f32
	l46 = s0f32
	s0f32 = l3
	s1f32 = l53
	s0f32 = s0f32 * s1f32
	l4 = s0f32
	s0f32 = l58
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l69 = s0f32
	s0f32 = 1.5
	s1f32 = l61
	s0f32 = s0f32 - s1f32
	l52 = s0f32
	s0f32 = l40
	l3 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl121
	case 1:
		goto lbl122
	default:
		goto lbl120
	}
lbl122:
	s0f32 = l40
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l3 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l3
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l3 = s0f32
	goto lbl120
lbl121:
	s0f32 = l40
	s1f32 = l35
	s2f32 = l40
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l3 = s0f32
lbl120:
	s0f32 = l46
	s1f32 = l37
	s0f32 = s0f32 * s1f32
	l39 = s0f32
	s0i32 = l23
	s0f32 = math.Float32frombits(uint32(s0i32))
	l53 = s0f32
	s0i32 = l12
	s0f32 = math.Float32frombits(uint32(s0i32))
	l58 = s0f32
	s0f32 = l4
	s1f32 = 0.055555556
	s0f32 = s0f32 + s1f32
	l68 = s0f32
	s0f32 = l57
	s1f32 = l69
	s0f32 = s0f32 * s1f32
	l69 = s0f32
	s0f32 = l57
	s1f32 = 0.3888889
	s0f32 = s0f32 * s1f32
	l61 = s0f32
	s0f32 = l5
	s1f32 = l52
	s0f32 = s0f32 * s1f32
	l52 = s0f32
	s0f32 = l2
	l4 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl124
	case 1:
		goto lbl125
	default:
		goto lbl123
	}
lbl125:
	s0f32 = l2
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l4 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l4 = s0f32
	goto lbl123
lbl124:
	s0f32 = l2
	s1f32 = l36
	s2f32 = l2
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
lbl123:
	s0f32 = l39
	s1f32 = l59
	s0f32 = s0f32 * s1f32
	l70 = s0f32
	s0f32 = l39
	s1f32 = l53
	s0f32 = s0f32 * s1f32
	l53 = s0f32
	s0f32 = l39
	s1f32 = l66
	s0f32 = s0f32 * s1f32
	l71 = s0f32
	s0f32 = l39
	s1f32 = l58
	s0f32 = s0f32 * s1f32
	l58 = s0f32
	s0f32 = l37
	s1f32 = l68
	s0f32 = s0f32 * s1f32
	l39 = s0f32
	s0i32 = l24
	s0f32 = math.Float32frombits(uint32(s0i32))
	l74 = s0f32
	s0i32 = l7
	s0f32 = math.Float32frombits(uint32(s0i32))
	l77 = s0f32
	s0f32 = l69
	s1f32 = 0.055555556
	s0f32 = s0f32 + s1f32
	l66 = s0f32
	s0f32 = l61
	s1f32 = -0.33333334
	s0f32 = s0f32 + s1f32
	l61 = s0f32
	s0f32 = l57
	s1f32 = l57
	s0f32 = s0f32 * s1f32
	l57 = s0f32
	s0f32 = l52
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l52 = s0f32
	s0f32 = 0
	l59 = s0f32
	s0i32 = 0
	l12 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l22 = s0i32
		s0f32 = 0
		goto lbl126
	}
	s0i32 = 0
	l22 = s0i32
	s0f32 = 0
	s1i32 = l7
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl126
	}
	s0f32 = l4
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
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
		goto lbl128
	}
	s0i32 = 0
lbl128:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l12 = s0i32
	s0f32 = l3
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l3 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l3
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l3
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl130
	}
	s0i32 = 0
lbl130:
	l7 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l7
	s2i32 = l12
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
	l12 = s0i32
	s1i32 = l7
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l23 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l24 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l22 = s0i32
	s0i32 = l23
	s1i32 = l12
	s2i32 = l24
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l12 = s0i32
	s0i32 = l7
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l59 = s0f32
	s0i32 = l7
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl126:
	l69 = s0f32
	s0f32 = l70
	s1f32 = 0
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s0f32 = l53
	s1f32 = 0
	s0f32 = s0f32 + s1f32
	l53 = s0f32
	s0f32 = l71
	s1f32 = 0
	s0f32 = s0f32 + s1f32
	l70 = s0f32
	s0f32 = l58
	s1f32 = 0
	s0f32 = s0f32 + s1f32
	l58 = s0f32
	s0f32 = l39
	s1f32 = l41
	s0f32 = s0f32 * s1f32
	l41 = s0f32
	s0f32 = l39
	s1f32 = l74
	s0f32 = s0f32 * s1f32
	l71 = s0f32
	s0f32 = l39
	s1f32 = l47
	s0f32 = s0f32 * s1f32
	l74 = s0f32
	s0f32 = l39
	s1f32 = l77
	s0f32 = s0f32 * s1f32
	l39 = s0f32
	s0f32 = l37
	s1f32 = l66
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0i32 = l26
	s0f32 = math.Float32frombits(uint32(s0i32))
	l77 = s0f32
	s0i32 = l8
	s0f32 = math.Float32frombits(uint32(s0i32))
	l78 = s0f32
	s0f32 = l57
	s1f32 = l61
	s0f32 = s0f32 * s1f32
	l57 = s0f32
	s0f32 = l5
	s1f32 = l52
	s0f32 = s0f32 * s1f32
	l5 = s0f32
	s0f32 = l55
	s1f32 = 1.1666666
	s0f32 = s0f32 * s1f32
	l52 = s0f32
	s0f32 = l45
	l47 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl133
	case 1:
		goto lbl134
	default:
		goto lbl132
	}
lbl134:
	s0f32 = l45
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l47 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l47
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l47 = s0f32
	goto lbl132
lbl133:
	s0f32 = l45
	s1f32 = l35
	s2f32 = l45
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l47 = s0f32
lbl132:
	s0f32 = l4
	s1f32 = l41
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s0f32 = l53
	s1f32 = l71
	s0f32 = s0f32 + s1f32
	l53 = s0f32
	s0f32 = l70
	s1f32 = l74
	s0f32 = s0f32 + s1f32
	l61 = s0f32
	s0f32 = l58
	s1f32 = l39
	s0f32 = s0f32 + s1f32
	l39 = s0f32
	s0f32 = l3
	s1f32 = l48
	s0f32 = s0f32 * s1f32
	l48 = s0f32
	s0f32 = l3
	s1f32 = l77
	s0f32 = s0f32 * s1f32
	l58 = s0f32
	s0f32 = l3
	s1f32 = l49
	s0f32 = s0f32 * s1f32
	l49 = s0f32
	s0f32 = l3
	s1f32 = l78
	s0f32 = s0f32 * s1f32
	l70 = s0f32
	s0f32 = l57
	s1f32 = l37
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0i32 = l27
	s0f32 = math.Float32frombits(uint32(s0i32))
	l37 = s0f32
	s0i32 = l13
	s0f32 = math.Float32frombits(uint32(s0i32))
	l71 = s0f32
	s0f32 = l5
	s1f32 = 0.055555556
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s0f32 = 1.5
	s1f32 = l52
	s0f32 = s0f32 - s1f32
	l52 = s0f32
	s0f32 = l2
	l41 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl136
	case 1:
		goto lbl137
	default:
		goto lbl135
	}
lbl137:
	s0f32 = l2
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l41 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l41
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l41 = s0f32
	goto lbl135
lbl136:
	s0f32 = l2
	s1f32 = l36
	s2f32 = l2
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l41 = s0f32
lbl135:
	s0f32 = l4
	s1f32 = l48
	s0f32 = s0f32 + s1f32
	l48 = s0f32
	s0f32 = l53
	s1f32 = l58
	s0f32 = s0f32 + s1f32
	l53 = s0f32
	s0f32 = l61
	s1f32 = l49
	s0f32 = s0f32 + s1f32
	l58 = s0f32
	s0f32 = l39
	s1f32 = l70
	s0f32 = s0f32 + s1f32
	l39 = s0f32
	s0f32 = l3
	s1f32 = l50
	s0f32 = s0f32 * s1f32
	l50 = s0f32
	s0f32 = l3
	s1f32 = l37
	s0f32 = s0f32 * s1f32
	l61 = s0f32
	s0f32 = l3
	s1f32 = l56
	s0f32 = s0f32 * s1f32
	l56 = s0f32
	s0f32 = l3
	s1f32 = l71
	s0f32 = s0f32 * s1f32
	l70 = s0f32
	s0f32 = l46
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0i32 = l28
	s0f32 = math.Float32frombits(uint32(s0i32))
	l71 = s0f32
	s0i32 = l10
	s0f32 = math.Float32frombits(uint32(s0i32))
	l74 = s0f32
	s0f32 = l55
	s1f32 = l52
	s0f32 = s0f32 * s1f32
	l52 = s0f32
	s0f32 = 0
	l37 = s0f32
	s0i32 = 0
	l7 = s0i32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l13 = s0i32
		s0f32 = 0
		goto lbl138
	}
	s0i32 = 0
	l13 = s0i32
	s0f32 = 0
	s1i32 = l8
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl138
	}
	s0f32 = l41
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
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
		goto lbl140
	}
	s0i32 = 0
lbl140:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l7 = s0i32
	s0f32 = l47
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl142
	}
	s0i32 = 0
lbl142:
	l8 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l7
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l7 = s0i32
	s1i32 = l8
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
	l23 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l13 = s0i32
	s0i32 = l10
	s1i32 = l7
	s2i32 = l23
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l7 = s0i32
	s0i32 = l8
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l37 = s0f32
	s0i32 = l8
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl138:
	l49 = s0f32
	s0f32 = l48
	s1f32 = l50
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s0f32 = l53
	s1f32 = l61
	s0f32 = s0f32 + s1f32
	l41 = s0f32
	s0f32 = l58
	s1f32 = l56
	s0f32 = s0f32 + s1f32
	l47 = s0f32
	s0f32 = l39
	s1f32 = l70
	s0f32 = s0f32 + s1f32
	l56 = s0f32
	s0f32 = l3
	s1f32 = l51
	s0f32 = s0f32 * s1f32
	l39 = s0f32
	s0f32 = l3
	s1f32 = l71
	s0f32 = s0f32 * s1f32
	l48 = s0f32
	s0f32 = l3
	s1f32 = l60
	s0f32 = s0f32 * s1f32
	l60 = s0f32
	s0f32 = l3
	s1f32 = l74
	s0f32 = s0f32 * s1f32
	l50 = s0f32
	s0f32 = l68
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0i32 = l29
	s0f32 = math.Float32frombits(uint32(s0i32))
	l51 = s0f32
	s0i32 = l14
	s0f32 = math.Float32frombits(uint32(s0i32))
	l53 = s0f32
	s0f32 = l52
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l52 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl145
	case 1:
		goto lbl146
	default:
		goto lbl144
	}
lbl146:
	s0f32 = l38
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l38 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l38
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l38 = s0f32
	goto lbl144
lbl145:
	s0f32 = l38
	s1f32 = l35
	s2f32 = l38
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l38 = s0f32
lbl144:
	s0f32 = l4
	s1f32 = l39
	s0f32 = s0f32 + s1f32
	l39 = s0f32
	s0f32 = l41
	s1f32 = l48
	s0f32 = s0f32 + s1f32
	l48 = s0f32
	s0f32 = l47
	s1f32 = l60
	s0f32 = s0f32 + s1f32
	l47 = s0f32
	s0f32 = l56
	s1f32 = l50
	s0f32 = s0f32 + s1f32
	l56 = s0f32
	s0f32 = l3
	s1f32 = l54
	s0f32 = s0f32 * s1f32
	l60 = s0f32
	s0f32 = l3
	s1f32 = l51
	s0f32 = s0f32 * s1f32
	l50 = s0f32
	s0f32 = l3
	s1f32 = l62
	s0f32 = s0f32 * s1f32
	l62 = s0f32
	s0f32 = l3
	s1f32 = l53
	s0f32 = s0f32 * s1f32
	l51 = s0f32
	s0f32 = l66
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0i32 = l30
	s0f32 = math.Float32frombits(uint32(s0i32))
	l54 = s0f32
	s0i32 = l15
	s0f32 = math.Float32frombits(uint32(s0i32))
	l53 = s0f32
	s0f32 = l55
	s1f32 = l52
	s0f32 = s0f32 * s1f32
	l52 = s0f32
	s0f32 = l2
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	l41 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl148
	case 1:
		goto lbl149
	default:
		goto lbl147
	}
lbl149:
	s0f32 = l4
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l2 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l2
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l41 = s0f32
	goto lbl147
lbl148:
	s0f32 = l4
	s1f32 = l36
	s2f32 = l4
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l41 = s0f32
lbl147:
	s0f32 = l39
	s1f32 = l60
	s0f32 = s0f32 + s1f32
	l39 = s0f32
	s0f32 = l48
	s1f32 = l50
	s0f32 = s0f32 + s1f32
	l48 = s0f32
	s0f32 = l47
	s1f32 = l62
	s0f32 = s0f32 + s1f32
	l50 = s0f32
	s0f32 = l56
	s1f32 = l51
	s0f32 = s0f32 + s1f32
	l51 = s0f32
	s0f32 = l3
	s1f32 = l65
	s0f32 = s0f32 * s1f32
	l65 = s0f32
	s0f32 = l3
	s1f32 = l54
	s0f32 = s0f32 * s1f32
	l54 = s0f32
	s0f32 = l3
	s1f32 = l63
	s0f32 = s0f32 * s1f32
	l63 = s0f32
	s0f32 = l3
	s1f32 = l53
	s0f32 = s0f32 * s1f32
	l53 = s0f32
	s0f32 = l57
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0i32 = l31
	s0f32 = math.Float32frombits(uint32(s0i32))
	l58 = s0f32
	s0i32 = l17
	s0f32 = math.Float32frombits(uint32(s0i32))
	l61 = s0f32
	s0f32 = l52
	s1f32 = 0.055555556
	s0f32 = s0f32 + s1f32
	l3 = s0f32
	s0f32 = 0
	l47 = s0f32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		l60 = s0f32
		s0f32 = 0
		l62 = s0f32
		s0f32 = 0
		goto lbl150
	}
	s0f32 = 0
	l60 = s0f32
	s0f32 = 0
	l62 = s0f32
	s0f32 = 0
	s1i32 = l8
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl150
	}
	s0f32 = l41
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l5 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l5
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l5
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl152
	}
	s0i32 = 0
lbl152:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l8 = s0i32
	s0f32 = l38
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l5 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l5
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l5
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl154
	}
	s0i32 = 0
lbl154:
	l10 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l8
	s2i32 = l10
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l10 = s0i32
	s1i32 = l8
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l14 = s1i32
	s2i32 = l9
	s3i32 = 6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l15 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l60 = s0f32
	s0i32 = l14
	s1i32 = l10
	s2i32 = l15
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l47 = s0f32
	s0i32 = l8
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l62 = s0f32
	s0i32 = l8
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl150:
	l56 = s0f32
	s0f32 = l39
	s1f32 = l65
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s0f32 = l48
	s1f32 = l54
	s0f32 = s0f32 + s1f32
	l38 = s0f32
	s0f32 = l50
	s1f32 = l63
	s0f32 = s0f32 + s1f32
	l41 = s0f32
	s0f32 = l51
	s1f32 = l53
	s0f32 = s0f32 + s1f32
	l63 = s0f32
	s0f32 = l2
	s1f32 = l72
	s0f32 = s0f32 * s1f32
	l39 = s0f32
	s0f32 = l2
	s1f32 = l58
	s0f32 = s0f32 * s1f32
	l48 = s0f32
	s0f32 = l2
	s1f32 = l64
	s0f32 = s0f32 * s1f32
	l64 = s0f32
	s0f32 = l2
	s1f32 = l61
	s0f32 = s0f32 * s1f32
	l50 = s0f32
	s0f32 = l46
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0i32 = l32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l51 = s0f32
	s0i32 = l18
	s0f32 = math.Float32frombits(uint32(s0i32))
	l54 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl157
	case 1:
		goto lbl158
	default:
		goto lbl156
	}
lbl158:
	s0f32 = l44
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l44 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l44
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l44 = s0f32
	goto lbl156
lbl157:
	s0f32 = l44
	s1f32 = l35
	s2f32 = l44
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l44 = s0f32
lbl156:
	s0f32 = l5
	s1f32 = l39
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s0f32 = l38
	s1f32 = l48
	s0f32 = s0f32 + s1f32
	l38 = s0f32
	s0f32 = l41
	s1f32 = l64
	s0f32 = s0f32 + s1f32
	l64 = s0f32
	s0f32 = l63
	s1f32 = l50
	s0f32 = s0f32 + s1f32
	l63 = s0f32
	s0f32 = l2
	s1f32 = l76
	s0f32 = s0f32 * s1f32
	l39 = s0f32
	s0f32 = l2
	s1f32 = l51
	s0f32 = s0f32 * s1f32
	l48 = s0f32
	s0f32 = l2
	s1f32 = l67
	s0f32 = s0f32 * s1f32
	l67 = s0f32
	s0f32 = l2
	s1f32 = l54
	s0f32 = s0f32 * s1f32
	l50 = s0f32
	s0f32 = l68
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0i32 = l25
	s0f32 = math.Float32frombits(uint32(s0i32))
	l51 = s0f32
	s0i32 = l19
	s0f32 = math.Float32frombits(uint32(s0i32))
	l54 = s0f32
	s0f32 = l55
	s1f32 = 0.3888889
	s0f32 = s0f32 * s1f32
	l65 = s0f32
	s0f32 = l4
	l41 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl160
	case 1:
		goto lbl161
	default:
		goto lbl159
	}
lbl161:
	s0f32 = l4
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l41 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l41
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l41 = s0f32
	goto lbl159
lbl160:
	s0f32 = l4
	s1f32 = l36
	s2f32 = l4
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l41 = s0f32
lbl159:
	s0f32 = l5
	s1f32 = l39
	s0f32 = s0f32 + s1f32
	l39 = s0f32
	s0f32 = l38
	s1f32 = l48
	s0f32 = s0f32 + s1f32
	l38 = s0f32
	s0f32 = l64
	s1f32 = l67
	s0f32 = s0f32 + s1f32
	l48 = s0f32
	s0f32 = l63
	s1f32 = l50
	s0f32 = s0f32 + s1f32
	l50 = s0f32
	s0f32 = l2
	s1f32 = l73
	s0f32 = s0f32 * s1f32
	l72 = s0f32
	s0f32 = l2
	s1f32 = l51
	s0f32 = s0f32 * s1f32
	l51 = s0f32
	s0f32 = l2
	s1f32 = l75
	s0f32 = s0f32 * s1f32
	l76 = s0f32
	s0f32 = l2
	s1f32 = l54
	s0f32 = s0f32 * s1f32
	l54 = s0f32
	s0f32 = l66
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	l5 = s0f32
	s0i32 = l22
	s0f32 = math.Float32frombits(uint32(s0i32))
	l73 = s0f32
	s0i32 = l12
	s0f32 = math.Float32frombits(uint32(s0i32))
	l75 = s0f32
	s0f32 = l65
	s1f32 = -0.33333334
	s0f32 = s0f32 + s1f32
	l65 = s0f32
	s0f32 = l55
	s1f32 = l55
	s0f32 = s0f32 * s1f32
	l55 = s0f32
	s0f32 = 0
	l2 = s0f32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		l64 = s0f32
		s0f32 = 0
		l67 = s0f32
		s0f32 = 0
		goto lbl162
	}
	s0f32 = 0
	l64 = s0f32
	s0f32 = 0
	l67 = s0f32
	s0f32 = 0
	s1i32 = l12
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl162
	}
	s0f32 = l41
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
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
		goto lbl164
	}
	s0i32 = 0
lbl164:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l12 = s0i32
	s0f32 = l44
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl166
	}
	s0i32 = 0
lbl166:
	l8 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l8
	s2i32 = l12
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l8 = s0i32
	s1i32 = l12
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
	l14 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l64 = s0f32
	s0i32 = l10
	s1i32 = l8
	s2i32 = l14
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l2 = s0f32
	s0i32 = l12
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l67 = s0f32
	s0i32 = l12
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl162:
	l63 = s0f32
	s0f32 = l39
	s1f32 = l72
	s0f32 = s0f32 + s1f32
	l41 = s0f32
	s0f32 = l38
	s1f32 = l51
	s0f32 = s0f32 + s1f32
	l38 = s0f32
	s0f32 = l48
	s1f32 = l76
	s0f32 = s0f32 + s1f32
	l39 = s0f32
	s0f32 = l50
	s1f32 = l54
	s0f32 = s0f32 + s1f32
	l48 = s0f32
	s0f32 = l5
	s1f32 = l69
	s0f32 = s0f32 * s1f32
	l50 = s0f32
	s0f32 = l5
	s1f32 = l73
	s0f32 = s0f32 * s1f32
	l51 = s0f32
	s0f32 = l5
	s1f32 = l59
	s0f32 = s0f32 * s1f32
	l59 = s0f32
	s0f32 = l5
	s1f32 = l75
	s0f32 = s0f32 * s1f32
	l5 = s0f32
	s0f32 = l57
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0i32 = l13
	s0f32 = math.Float32frombits(uint32(s0i32))
	l54 = s0f32
	s0i32 = l7
	s0f32 = math.Float32frombits(uint32(s0i32))
	l72 = s0f32
	s0f32 = l55
	s1f32 = l65
	s0f32 = s0f32 * s1f32
	l44 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl169
	case 1:
		goto lbl170
	default:
		goto lbl168
	}
lbl170:
	s0f32 = l40
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l40 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l40
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l40 = s0f32
	goto lbl168
lbl169:
	s0f32 = l40
	s1f32 = l35
	s2f32 = l40
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l40 = s0f32
lbl168:
	s0f32 = l41
	s1f32 = l50
	s0f32 = s0f32 + s1f32
	l41 = s0f32
	s0f32 = l38
	s1f32 = l51
	s0f32 = s0f32 + s1f32
	l38 = s0f32
	s0f32 = l39
	s1f32 = l59
	s0f32 = s0f32 + s1f32
	l55 = s0f32
	s0f32 = l48
	s1f32 = l5
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s0f32 = l3
	s1f32 = l49
	s0f32 = s0f32 * s1f32
	l59 = s0f32
	s0f32 = l3
	s1f32 = l54
	s0f32 = s0f32 * s1f32
	l39 = s0f32
	s0f32 = l3
	s1f32 = l37
	s0f32 = s0f32 * s1f32
	l37 = s0f32
	s0f32 = l3
	s1f32 = l72
	s0f32 = s0f32 * s1f32
	l48 = s0f32
	s0f32 = l46
	s1f32 = l44
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0f32 = l4
	l49 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl172
	case 1:
		goto lbl173
	default:
		goto lbl171
	}
lbl173:
	s0f32 = l4
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l49 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l49
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l49 = s0f32
	goto lbl171
lbl172:
	s0f32 = l4
	s1f32 = l36
	s2f32 = l4
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l49 = s0f32
lbl171:
	s0f32 = l41
	s1f32 = l59
	s0f32 = s0f32 + s1f32
	l41 = s0f32
	s0f32 = l38
	s1f32 = l39
	s0f32 = s0f32 + s1f32
	l59 = s0f32
	s0f32 = l55
	s1f32 = l37
	s0f32 = s0f32 + s1f32
	l55 = s0f32
	s0f32 = l5
	s1f32 = l48
	s0f32 = s0f32 + s1f32
	l39 = s0f32
	s0f32 = l3
	s1f32 = l62
	s0f32 = s0f32 * s1f32
	l62 = s0f32
	s0f32 = l3
	s1f32 = l60
	s0f32 = s0f32 * s1f32
	l60 = s0f32
	s0f32 = l3
	s1f32 = l56
	s0f32 = s0f32 * s1f32
	l56 = s0f32
	s0f32 = l3
	s1f32 = l47
	s0f32 = s0f32 * s1f32
	l47 = s0f32
	s0f32 = l44
	s1f32 = l68
	s0f32 = s0f32 * s1f32
	l38 = s0f32
	s0f32 = 0
	l3 = s0f32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		l37 = s0f32
		s0f32 = 0
		l46 = s0f32
		s0f32 = 0
		goto lbl174
	}
	s0f32 = 0
	l37 = s0f32
	s0f32 = 0
	l46 = s0f32
	s0f32 = 0
	s1i32 = l7
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl174
	}
	s0f32 = l49
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l3 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l3
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l3
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl176
	}
	s0i32 = 0
lbl176:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l12 = s0i32
	s0f32 = l40
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = f14(ctx, s0f32, s1f32)
	l3 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l3
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l3
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl178
	}
	s0i32 = 0
lbl178:
	l7 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l7
	s2i32 = l12
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
	l12 = s0i32
	s1i32 = l7
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l8 = s1i32
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
	l37 = s0f32
	s0i32 = l8
	s1i32 = l12
	s2i32 = l13
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l3 = s0f32
	s0i32 = l7
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l46 = s0f32
	s0i32 = l7
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl174:
	l5 = s0f32
	s0f32 = l41
	s1f32 = l62
	s0f32 = s0f32 + s1f32
	l40 = s0f32
	s0f32 = l59
	s1f32 = l60
	s0f32 = s0f32 + s1f32
	l41 = s0f32
	s0f32 = l55
	s1f32 = l56
	s0f32 = s0f32 + s1f32
	l49 = s0f32
	s0f32 = l39
	s1f32 = l47
	s0f32 = s0f32 + s1f32
	l47 = s0f32
	s0f32 = l38
	s1f32 = l67
	s0f32 = s0f32 * s1f32
	l55 = s0f32
	s0f32 = l38
	s1f32 = l64
	s0f32 = s0f32 * s1f32
	l59 = s0f32
	s0f32 = l38
	s1f32 = l63
	s0f32 = s0f32 * s1f32
	l56 = s0f32
	s0f32 = l38
	s1f32 = l2
	s0f32 = s0f32 * s1f32
	l38 = s0f32
	s0f32 = l44
	s1f32 = l66
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0i32 = l16
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl181
	case 1:
		goto lbl182
	default:
		goto lbl180
	}
lbl182:
	s0f32 = l45
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	l45 = s0f32
	s1f32 = l35
	s2f32 = l35
	s1f32 = s1f32 + s2f32
	s2f32 = l45
	s3f32 = l42
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l35
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l45 = s0f32
	goto lbl180
lbl181:
	s0f32 = l45
	s1f32 = l35
	s2f32 = l45
	s3f32 = l42
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l45 = s0f32
lbl180:
	s0f32 = l40
	s1f32 = l55
	s0f32 = s0f32 + s1f32
	l40 = s0f32
	s0f32 = l41
	s1f32 = l59
	s0f32 = s0f32 + s1f32
	l41 = s0f32
	s0f32 = l49
	s1f32 = l56
	s0f32 = s0f32 + s1f32
	l49 = s0f32
	s0f32 = l47
	s1f32 = l38
	s0f32 = s0f32 + s1f32
	l38 = s0f32
	s0f32 = l2
	s1f32 = l46
	s0f32 = s0f32 * s1f32
	l47 = s0f32
	s0f32 = l2
	s1f32 = l37
	s0f32 = s0f32 * s1f32
	l37 = s0f32
	s0f32 = l2
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	l5 = s0f32
	s0f32 = l2
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl184
	case 1:
		goto lbl185
	default:
		goto lbl183
	}
lbl185:
	s0f32 = l4
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	l3 = s0f32
	s1f32 = l36
	s2f32 = l36
	s1f32 = s1f32 + s2f32
	s2f32 = l3
	s3f32 = l43
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l36
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l4 = s0f32
	goto lbl183
lbl184:
	s0f32 = l4
	s1f32 = l36
	s2f32 = l4
	s3f32 = l43
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l4 = s0f32
lbl183:
	s0f32 = l40
	s1f32 = l47
	s0f32 = s0f32 + s1f32
	l40 = s0f32
	s0f32 = l41
	s1f32 = l37
	s0f32 = s0f32 + s1f32
	l37 = s0f32
	s0f32 = l49
	s1f32 = l5
	s0f32 = s0f32 + s1f32
	l41 = s0f32
	s0f32 = l38
	s1f32 = l2
	s0f32 = s0f32 + s1f32
	l49 = s0f32
	s0f32 = 0
	l5 = s0f32
	s0i32 = l9
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		l3 = s0f32
		s0f32 = 0
		l38 = s0f32
		s0f32 = 0
		goto lbl186
	}
	s0f32 = 0
	l3 = s0f32
	s0f32 = 0
	l38 = s0f32
	s0f32 = 0
	s1i32 = l11
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl186
	}
	s0f32 = l4
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l21
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
		goto lbl188
	}
	s0i32 = 0
lbl188:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 * s1i32
	l7 = s0i32
	s0f32 = l45
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l20
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
		goto lbl190
	}
	s0i32 = 0
lbl190:
	l11 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l7
	s2i32 = l11
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
	l11 = s0i32
	s1i32 = l6
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	l7 = s1i32
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
	l3 = s0f32
	s0i32 = l7
	s1i32 = l11
	s2i32 = l9
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	l5 = s0f32
	s0i32 = l6
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
	l38 = s0f32
	s0i32 = l6
	s1i32 = 8
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s0f32 = float32(uint32(s0i32))
	s1f32 = 0.003921569
	s0f32 = s0f32 * s1f32
lbl186:
	l2 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2f32 = l49
	s3f32 = l57
	s4f32 = l44
	s3f32 = s3f32 * s4f32
	l4 = s3f32
	s4f32 = l5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l41
	s4f32 = l4
	s5f32 = l2
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l37
	s5f32 = l4
	s6f32 = l3
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s5f32 = l40
	s6f32 = l4
	s7f32 = l38
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
