package internal

import (
	"math"
	"unsafe"
)

func f559(ctx *Context, l0 int32, l1 int32, l2 int32, l3 float32, l4 int32, l5 int32) int32 {
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
	var l25 int64
	_ = l25
	var l26 int64
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
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
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
	s0i32 = ctx.g0
	s1i32 = 560
	s0i32 = s0i32 - s1i32
	l9 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s1i32 = -3
	s0i32 = s0i32 + s1i32
	s1i32 = 65531
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l3
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
		goto lbl1
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	l27 = s0f32
	s1f32 = -0.5
	s0f32 = s0f32 * s1f32
	s1f32 = l27
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	l27 = s1f32
	s2f32 = l27
	s3f32 = 0
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l27 = s0f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f32 = s1f32 - s2f32
	l28 = s1f32
	s2f32 = -0.5
	s1f32 = s1f32 * s2f32
	s2f32 = l28
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	l28 = s2f32
	s3f32 = l28
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l28 = s1f32
	s2f32 = l27
	s3f32 = l28
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s1f32 = l3
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l3
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00024414062
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
		s0i32 = 0
		l2 = s0i32
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl4:
			s0i32 = 1
			l11 = s0i32
			s0i32 = l4
			s0i32 = f45(ctx, s0i32)
			s1i32 = l0
			s2i32 = l2
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s1i32 = l1
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl4
			}
			goto lbl1
			panic("unreachable executed")
			panic("unreachable executed")
		}
	lbl5:
		s0i32 = 1
		l11 = s0i32
		s0i32 = l4
		s0i32 = f45(ctx, s0i32)
		s1i32 = l0
		s2i32 = l2
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s0i32 = f382(ctx, s0i32)
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		goto lbl1
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l31 = s1f32
	s0f32 = s0f32 - s1f32
	l27 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l32 = s1f32
	s0f32 = s0f32 - s1f32
	l30 = s0f32
	s0i32 = 2
	l2 = s0i32
lbl6:
	s0f32 = l30
	s1i32 = l0
	s2i32 = l2
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = l31
	s1f32 = s1f32 - s2f32
	l28 = s1f32
	s0f32 = s0f32 * s1f32
	l33 = s0f32
	s0f32 = l29
	s1f32 = l27
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3f32 = l32
	s2f32 = s2f32 - s3f32
	l30 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l33
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 + s1f32
	l29 = s0f32
	s0f32 = l28
	l27 = s0f32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0f32 = l29
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 5.9604645e-08
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 0
	s1i32 = 1
	s2i32 = -1
	s3f32 = l29
	s4f32 = 0
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l19 = s0i32
	s0i32 = l1
	s1i32 = 64
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		if s0i32 != 0 {
			s0i32 = l9
			s1i32 = l9
			s2i32 = 32
			s1i32 = s1i32 + s2i32
			s2i32 = 4
			s1i32 = s1i32 | s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l7 = s0i32
			goto lbl9
		}
		s0i32 = l9
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = -1
		l7 = s0i32
		goto lbl8
	}
	s0i32 = l9
	s1i32 = l1
	s2i32 = 8
	s1i32 = f50(ctx, s1i32, s2i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = l1
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	goto lbl8
lbl9:
	s0i32 = l19
	s0f32 = float32(s0i32)
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	l27 = s0f32
	s0i32 = l7
	l2 = s0i32
lbl12:
	s0i32 = l2
	l10 = s0i32
	s0i32 = l0
	s1i32 = l6
	l2 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l8 = s1i32
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l28 = s0f32
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l29 = s1f32
	s0f32 = s0f32 * s1f32
	l30 = s0f32
	s1f32 = l30
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l9
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l13 = s0i32
	s0i32 = l0
	s1i32 = 0
	s2i32 = l2
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	l6 = s2i32
	s3i32 = l1
	s4i32 = l6
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l15 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l16 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l30 = s0f32
	s0i32 = l9
	s1i32 = l16
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = l28
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+556)])) = s1f32
	s0i32 = l9
	s1f32 = l29
	s2f32 = l30
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+552)])) = s1f32
	s0i32 = l9
	s1i32 = 552
	s0i32 = s0i32 + s1i32
	s1f32 = l27
	s0i32 = f113(ctx, s0i32, s1f32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l8
	s1i32 = l13
	s0i32 = s0i32 + s1i32
	s1i32 = l9
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+552)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = 0
	s1i32 = 1
	s2i32 = -1
	s3i32 = l0
	s4i32 = l2
	s5i32 = 65535
	s4i32 = s4i32 & s5i32
	s5i32 = 3
	s4i32 = s4i32 << (uint32(s5i32) & 31)
	s3i32 = s3i32 + s4i32
	l13 = s3i32
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l0
	s5i32 = l10
	s6i32 = 65535
	s5i32 = s5i32 & s6i32
	s6i32 = 3
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	l16 = s4i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	l28 = s4f32
	s3f32 = s3f32 - s4f32
	s4i32 = l0
	s5i32 = l15
	s6i32 = 65535
	s5i32 = s5i32 & s6i32
	s6i32 = 3
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	l15 = s4i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s5i32 = l16
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+4)]))
	l29 = s5f32
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s4i32 = l13
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s5f32 = l29
	s4f32 = s4f32 - s5f32
	s5i32 = l15
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s6f32 = l28
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l28 = s3f32
	s4f32 = 0
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2f32 = l28
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 5.9604645e-08
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l19
	s0i32 = s0i32 * s1i32
	s0f32 = float32(s0i32)
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l9
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l13 = s0i32
	s1i32 = l10
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l28 = s0f32
	s1i32 = l8
	s2i32 = l13
	s1i32 = s1i32 + s2i32
	l8 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l29 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l10
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l30 = s1f32
	s2i32 = l8
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l31 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l32 = s0f32
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
		goto lbl7
	}
	s0f32 = l28
	s1f32 = l31
	s0f32 = s0f32 * s1f32
	s1f32 = l29
	s2f32 = l30
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l28 = s0f32
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
		goto lbl7
	}
	s0f32 = l28
	s1f32 = l32
	s0f32 = f298(ctx, s0f32, s1f32)
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	s1f32 = 0.25
	s0f32 = s0f32 * s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	l28 = s0f32
	s1f32 = 65535
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0f32 = l28
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	l28 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l28
	s3f32 = 2.1474835e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l28 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l28
	s3f32 = -2.1474835e+09
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l28 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l28
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl14
	}
	s0i32 = -2147483648
lbl14:
	l10 = s0i32
	s1i32 = 1
	s2i32 = l10
	s3i32 = 1
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l12
	s0i32 = s0i32 + s1i32
	l12 = s0i32
lbl13:
	s0i32 = 1
	l8 = s0i32
	s0i32 = l12
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s0i32 = l1
	s1i32 = l6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
lbl8:
	s0i32 = 0
	s1i32 = 1
	s2i32 = -1
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l0
	s5i32 = l7
	s6i32 = 65535
	s5i32 = s5i32 & s6i32
	s6i32 = 3
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	l2 = s4i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	l27 = s4f32
	s3f32 = s3f32 - s4f32
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s5i32 = l2
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+4)]))
	l28 = s5f32
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s5f32 = l28
	s4f32 = s4f32 - s5f32
	s5i32 = l0
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+8)]))
	s6f32 = l27
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l27 = s3f32
	s4f32 = 0
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2f32 = l27
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 5.9604645e-08
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l19
	s0i32 = s0i32 * s1i32
	s0f32 = float32(s0i32)
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 < s1f32 {
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
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l2 = s0i32
		s1i32 = l7
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l27 = s0f32
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l28 = s1f32
		s0f32 = s0f32 * s1f32
		s1i32 = l10
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l29 = s1f32
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l30 = s2f32
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l31 = s0f32
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
			goto lbl7
		}
		s0f32 = l27
		s1f32 = l30
		s0f32 = s0f32 * s1f32
		s1f32 = l28
		s2f32 = l29
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 - s1f32
		l27 = s0f32
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
			goto lbl7
		}
		s0f32 = l27
		s1f32 = l31
		s0f32 = f298(ctx, s0f32, s1f32)
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = 0.25
		s0f32 = s0f32 * s1f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		l27 = s0f32
		s1f32 = 65535
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0f32 = l27
		s1f32 = 0.5
		s0f32 = s0f32 + s1f32
		s0f32 = float32(math.Floor(float64(s0f32)))
		l27 = s0f32
		s1f32 = 2.1474835e+09
		s2f32 = l27
		s3f32 = 2.1474835e+09
		if s2f32 < s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l27 = s0f32
		s1f32 = -2.1474835e+09
		s2f32 = l27
		s3f32 = -2.1474835e+09
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l27 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l27
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl17
		}
		s0i32 = -2147483648
	lbl17:
		l2 = s0i32
		s1i32 = 1
		s2i32 = l2
		s3i32 = 1
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l12
		s0i32 = s0i32 + s1i32
		l12 = s0i32
	}
	s0i32 = l12
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = 0
	l6 = s0i32
	s0i32 = l12
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l14 = s0i32
		s0i32 = 0
		goto lbl19
	}
	s0i32 = 3
	l14 = s0i32
	s0i32 = l12
	s1i32 = 8
	s2i32 = l12
	s3i32 = 8
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l17 = s0i32
	s1i32 = 40
	s0i32 = f50(ctx, s0i32, s1i32)
lbl19:
	l10 = s0i32
	s0i32 = l8
	if s0i32 != 0 {
		s0i32 = 0
		l16 = s0i32
	lbl23:
		s0i32 = 0
		s1i32 = 1
		s2i32 = -1
		s3i32 = l0
		s4i32 = l20
		l13 = s4i32
		s5i32 = 65535
		s4i32 = s4i32 & s5i32
		s5i32 = 3
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s3i32 = s3i32 + s4i32
		l2 = s3i32
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = l0
		s5i32 = l7
		s6i32 = 65535
		s5i32 = s5i32 & s6i32
		s6i32 = 3
		s5i32 = s5i32 << (uint32(s6i32) & 31)
		s4i32 = s4i32 + s5i32
		l8 = s4i32
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		l27 = s4f32
		s3f32 = s3f32 - s4f32
		s4i32 = l0
		s5i32 = 0
		s6i32 = l13
		s7i32 = 1
		s6i32 = s6i32 + s7i32
		l20 = s6i32
		s7i32 = l1
		s8i32 = l20
		if s7i32 == s8i32 {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i32 = s5i32
		} else {
			s5i32 = s6i32
		}
		l22 = s5i32
		s6i32 = 65535
		s5i32 = s5i32 & s6i32
		s6i32 = 3
		s5i32 = s5i32 << (uint32(s6i32) & 31)
		s4i32 = s4i32 + s5i32
		l15 = s4i32
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		s5i32 = l8
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+4)]))
		l28 = s5f32
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 * s4f32
		s4i32 = l2
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		s5f32 = l28
		s4f32 = s4f32 - s5f32
		s5i32 = l15
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s6f32 = l27
		s5f32 = s5f32 - s6f32
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 - s4f32
		l27 = s3f32
		s4f32 = 0
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2f32 = l27
		s2f32 = float32(math.Abs(float64(s2f32)))
		s3f32 = 5.9604645e-08
		if s2f32 <= s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l19
		s0i32 = s0i32 * s1i32
		s0f32 = float32(s0i32)
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = 0
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			s0i32 = l10
			l7 = s0i32
			s0i32 = l16
			l15 = s0i32
			s0i32 = l6
			l8 = s0i32
			goto lbl24
		}
		s0i32 = 0
		l11 = s0i32
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l2 = s0i32
		s1i32 = l13
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l18 = s1i32
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l29 = s0f32
		s1i32 = l2
		s2i32 = l7
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l23 = s1i32
		s1f32 = math.Float32frombits(uint32(s1i32))
		l27 = s1f32
		s0f32 = s0f32 * s1f32
		s1i32 = l8
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l30 = s1f32
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l24 = s2i32
		s2f32 = math.Float32frombits(uint32(s2i32))
		l28 = s2f32
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l31 = s0f32
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
			goto lbl21
		}
		s0f32 = l30
		s1f32 = l27
		s0f32 = s0f32 * s1f32
		s1f32 = l29
		s2f32 = l28
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 - s1f32
		l29 = s0f32
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
			goto lbl21
		}
		s0f32 = l29
		s1f32 = l31
		s0f32 = f298(ctx, s0f32, s1f32)
		l29 = s0f32
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = 0.25
		s0f32 = s0f32 * s1f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		l30 = s0f32
		s1f32 = 65535
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl21
		}
		s0f32 = 0
		s1f32 = l30
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l30 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l30
		s4f32 = 2.1474835e+09
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l30 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l30
		s4f32 = -2.1474835e+09
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l30 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l30
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl27
		}
		s1i32 = -2147483648
	lbl27:
		l11 = s1i32
		s2i32 = 1
		if s1i32 < s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl26
		}
		s0f32 = l29
		s1i32 = l11
		s1f32 = float32(s1i32)
		s0f32 = s0f32 / s1f32
	lbl26:
		l29 = s0f32
		s0i32 = l11
		s1i32 = 1
		s2i32 = l11
		s3i32 = 1
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l16
		s0i32 = s0i32 + s1i32
		l15 = s0i32
		s0i64 = int64(s0i32)
		l25 = s0i64
		s0i32 = l17
		s0i64 = int64(s0i32)
		l26 = s0i64
		s0f32 = l29
		s0f32 = f86(ctx, s0f32)
		l30 = s0f32
		s0f32 = l29
		s0f32 = f107(ctx, s0f32)
		l29 = s0f32
		s0i32 = 0
		l2 = s0i32
		s0i32 = l14
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		l21 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl29
		}
		s0i64 = l25
		s1i64 = 3
		s0i64 = s0i64 * s1i64
		s1i64 = l26
		if s0i64 >= s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl29
		}
		s0i32 = l14
		s1i32 = 2
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l2 = s0i32
	lbl29:
		s0i32 = l15
		s1i32 = l17
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl31
		}
		s0i32 = l2
		if s0i32 != 0 {
			goto lbl31
		}
		s0i32 = l10
		l7 = s0i32
		goto lbl30
	lbl31:
		s0i64 = l26
		s1i64 = l25
		s2i64 = l25
		s3i64 = 1
		s2i64 = s2i64 + s3i64
		s3i64 = 1
		s2i64 = s2i64 >> (uint64(s3i64) & 63)
		s1i64 = s1i64 + s2i64
		s2i64 = 7
		s1i64 = s1i64 + s2i64
		s2i64 = -8
		s1i64 = s1i64 & s2i64
		l25 = s1i64
		if s0i64 == s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l10
			l7 = s0i32
			goto lbl30
		}
		s0i64 = l25
		s1i64 = 2147483647
		s2i64 = l25
		s3i64 = 2147483647
		if s2i64 < s3i64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		l25 = s0i64
		s1i64 = -2147483647
		s2i64 = l25
		s3i64 = -2147483647
		if s2i64 > s3i64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		s0i32 = int32(s0i64)
		l17 = s0i32
		s1i32 = 40
		s0i32 = f50(ctx, s0i32, s1i32)
		l7 = s0i32
		s0i32 = 0
		l8 = s0i32
		s0i32 = l16
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl34:
			s0i32 = l7
			s1i32 = l8
			s2i32 = 40
			s1i32 = s1i32 * s2i32
			l14 = s1i32
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s1i32 = l10
			s2i32 = l14
			s1i32 = s1i32 + s2i32
			l14 = s1i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l14
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l14
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l14
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l14
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l8
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l8 = s0i32
			s1i32 = l16
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl34
			}
		}
		s0i32 = l21
		if s0i32 != 0 {
			s0i32 = l10
			f13(ctx, s0i32)
		}
		s0i32 = 1
		l14 = s0i32
	lbl30:
		s0i32 = l7
		s1i32 = l16
		s2i32 = 40
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l11
		s1i32 = 2
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			l8 = s0i32
			goto lbl36
		}
		s0i32 = l11
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s0i32 = l0
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l16 = s0i32
		l21 = s0i32
		s0i32 = 0
		l11 = s0i32
	lbl38:
		s0i32 = l16
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l27 = s0f32
		s0i32 = l21
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l28 = s0f32
		s0i32 = l2
		s1i32 = l13
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+38)])) = uint16(s1i32)
		s0i32 = l2
		s1i32 = l13
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint16(s1i32)
		s0i32 = l2
		s1i32 = -8388609
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1f32 = l28
		s2i32 = l24
		s2f32 = math.Float32frombits(uint32(s2i32))
		l31 = s2f32
		s1f32 = s1f32 + s2f32
		l32 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l2
		s1f32 = l27
		s2i32 = l23
		s2f32 = math.Float32frombits(uint32(s2i32))
		l33 = s2f32
		s1f32 = s1f32 + s2f32
		l34 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l2
		s1f32 = l32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l2
		s1f32 = l34
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l2
		s1f32 = l28
		s2f32 = l29
		s3f32 = l33
		s2f32 = s2f32 * s3f32
		s3f32 = l30
		s4f32 = l31
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		l28 = s2f32
		s1f32 = s1f32 + s2f32
		s2f32 = l32
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l2
		s1f32 = l27
		s2f32 = l30
		s3f32 = l33
		s2f32 = s2f32 * s3f32
		s3f32 = l29
		s4f32 = l31
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 - s3f32
		l27 = s2f32
		s1f32 = s1f32 + s2f32
		s2f32 = l34
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l6
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		}
		s0f32 = l28
		s0i32 = int32(math.Float32bits(s0f32))
		l24 = s0i32
		s0f32 = l27
		s0i32 = int32(math.Float32bits(s0f32))
		l23 = s0i32
		s0i32 = l2
		l6 = s0i32
		s1i32 = 40
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		l2 = s0i32
		s0i32 = l11
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s1i32 = l10
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl38
		}
	lbl36:
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l31 = s0f32
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l32 = s0f32
		s0i32 = l0
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l29 = s0f32
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l30 = s0f32
		s0i32 = l8
		s1i32 = l13
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+38)])) = uint16(s1i32)
		s0i32 = l8
		s1i32 = l13
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint16(s1i32)
		s0i32 = l8
		s1i32 = -8388609
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l8
		s1f32 = l30
		s2f32 = l28
		s1f32 = s1f32 + s2f32
		l28 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l8
		s1f32 = l29
		s2f32 = l27
		s1f32 = s1f32 + s2f32
		l27 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l8
		s1f32 = l28
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l8
		s1f32 = l27
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l8
		s1f32 = l30
		s2f32 = l32
		s1f32 = s1f32 + s2f32
		s2f32 = l28
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l8
		s1f32 = l29
		s2f32 = l31
		s1f32 = s1f32 + s2f32
		s2f32 = l27
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l6
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl24
		}
		s0i32 = l6
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	lbl24:
		s0i32 = l14
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		l18 = s0i32
		s0i32 = 0
		l2 = s0i32
		s0i32 = l15
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l16 = s0i32
		s0i64 = int64(s0i32)
		l25 = s0i64
		s1i64 = 3
		s0i64 = s0i64 * s1i64
		s1i32 = l17
		s1i64 = int64(s1i32)
		l26 = s1i64
		if s0i64 >= s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl40
		}
		s0i32 = l18
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl40
		}
		s0i32 = l14
		s1i32 = 2
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l2 = s0i32
	lbl40:
		s0i32 = l15
		s1i32 = l17
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl42
		}
		s0i32 = l2
		if s0i32 != 0 {
			goto lbl42
		}
		s0i32 = l7
		l10 = s0i32
		goto lbl41
	lbl42:
		s0i64 = l26
		s1i64 = l25
		s2i64 = l25
		s3i64 = 1
		s2i64 = s2i64 + s3i64
		s3i64 = 1
		s2i64 = s2i64 >> (uint64(s3i64) & 63)
		s1i64 = s1i64 + s2i64
		s2i64 = 7
		s1i64 = s1i64 + s2i64
		s2i64 = -8
		s1i64 = s1i64 & s2i64
		l25 = s1i64
		if s0i64 == s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l7
			l10 = s0i32
			goto lbl41
		}
		s0i64 = l25
		s1i64 = 2147483647
		s2i64 = l25
		s3i64 = 2147483647
		if s2i64 < s3i64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		l25 = s0i64
		s1i64 = -2147483647
		s2i64 = l25
		s3i64 = -2147483647
		if s2i64 > s3i64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		s0i32 = int32(s0i64)
		l17 = s0i32
		s1i32 = 40
		s0i32 = f50(ctx, s0i32, s1i32)
		l10 = s0i32
		s0i32 = 0
		l6 = s0i32
		s0i32 = l15
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl45:
			s0i32 = l10
			s1i32 = l6
			s2i32 = 40
			s1i32 = s1i32 * s2i32
			l11 = s1i32
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s1i32 = l7
			s2i32 = l11
			s1i32 = s1i32 + s2i32
			l11 = s1i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l11
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l11
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l11
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l11
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l6 = s0i32
			s1i32 = l15
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl45
			}
		}
		s0i32 = l18
		if s0i32 != 0 {
			s0i32 = l7
			f13(ctx, s0i32)
		}
		s0i32 = 1
		l14 = s0i32
	lbl41:
		s0i32 = l0
		s1i32 = l22
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l31 = s0f32
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l32 = s0f32
		s0i32 = l13
		s1i32 = 3
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l2 = s0i32
		s1i32 = l9
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l27 = s0f32
		s0i32 = l0
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l30 = s0f32
		s0i32 = l7
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l28 = s0f32
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l29 = s0f32
		s0i32 = l10
		s1i32 = l15
		s2i32 = 40
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l22
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+38)])) = uint16(s1i32)
		s0i32 = l6
		s1i32 = l13
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint16(s1i32)
		s0i32 = l6
		s1i32 = -8388609
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l6
		s1f32 = l29
		s2f32 = l28
		s1f32 = s1f32 + s2f32
		l29 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l6
		s1f32 = l30
		s2f32 = l27
		s1f32 = s1f32 + s2f32
		l30 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l6
		s1f32 = l29
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l6
		s1f32 = l30
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l6
		s1f32 = l28
		s2f32 = l32
		s1f32 = s1f32 + s2f32
		s2f32 = l29
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l6
		s1f32 = l27
		s2f32 = l31
		s1f32 = s1f32 + s2f32
		s2f32 = l30
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l8
		if s0i32 != 0 {
			s0i32 = l8
			s1i32 = l6
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		}
		s0i32 = l13
		l7 = s0i32
		s0i32 = l1
		s1i32 = l20
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl23
		}
	}
	s0i32 = l6
	s1i32 = l10
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l10
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l10
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l10
		l1 = s0i32
		goto lbl48
	}
	s0i32 = l12
	s0i64 = int64(uint32(s0i32))
	l25 = s0i64
	s1i64 = l25
	s0i64 = s0i64 * s1i64
	l26 = s0i64
	s0i64 = 0
	l25 = s0i64
	s0i32 = l10
	l0 = s0i32
	l1 = s0i32
lbl50:
	s0i32 = l12
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl48
	}
	s0i32 = l0
	l2 = s0i32
	s1i32 = l6
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl48
	}
	s0i64 = l25
	s1i64 = l26
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l11 = s0i32
		goto lbl21
	}
	s0i32 = l6
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+38)])))
	s1i32 = l2
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)])))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl55
	}
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s0f32 = s0f32 + s1f32
	l3 = s0f32
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s0f32 = s0f32 + s1f32
	l27 = s0f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0f32 = s0f32 - s1f32
	l28 = s0f32
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
		goto lbl56
	}
	s0f32 = l3
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0f32 = s0f32 - s1f32
	l29 = s0f32
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
		goto lbl56
	}
	s0f32 = l28
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl55
	}
	s0f32 = l29
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl55
	}
lbl56:
	s0i32 = l9
	s1f32 = l3
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+556)])) = s1f32
	s0i32 = l9
	s1f32 = l27
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+552)])) = s1f32
	s0i32 = l9
	s1i32 = 1065353216
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0f32 = 1
	goto lbl54
lbl55:
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l9
	s3i32 = 552
	s2i32 = s2i32 + s3i32
	s3i32 = l9
	s4i32 = 28
	s3i32 = s3i32 + s4i32
	s4i32 = l9
	s5i32 = 24
	s4i32 = s4i32 + s5i32
	s0i32 = f564(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl53
	}
	s0i32 = l9
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
lbl54:
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	if s0f32 < s1f32 {
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
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l0 = s0i32
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = 0
		s1i32 = l0
		s2i32 = l0
		s3i32 = l6
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l1
		s2i32 = l1
		s3i32 = l6
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l1 = s0i32
		s0i32 = l12
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		s0i32 = l2
		l0 = s0i32
		goto lbl52
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	s1f32 = -3.4028235e+38
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl58
	}
	s0i32 = l9
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+552)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 1e-06
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl58
	}
	s0i32 = l9
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+556)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 1e-06
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
		goto lbl48
	}
lbl58:
	s0i32 = l2
	s1i32 = l9
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+552)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l9
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l6
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+38)])))
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint16(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s0i32 = l2
	l6 = s0i32
	goto lbl52
lbl53:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s0f32 = 3.4028235e+38
	l27 = s0f32
	s0f32 = 3.4028235e+38
	l29 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l3 = s0f32
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l7 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l30 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l28 = s1f32
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	l31 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l32 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 5.9604645e-08
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l3
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = l28
		s2f32 = l28
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l30
		s2i32 = l7
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l2
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 * s2f32
		s2f32 = l31
		s3i32 = l7
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
		s4i32 = l2
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
		s3f32 = s3f32 - s4f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 - s2f32
		s2f32 = l32
		s1f32 = s1f32 / s2f32
		l29 = s1f32
		s1f32 = -s1f32
		s2f32 = l29
		s3f32 = -1
		s2f32 = s2f32 + s3f32
		s3f32 = l29
		s4f32 = 0
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l29 = s1f32
		s2f32 = l29
		s2f32 = float32(math.Abs(float64(s2f32)))
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 * s1f32
		l29 = s0f32
	}
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l32 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l34 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l33 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	l35 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l36 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 5.9604645e-08
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl61
	}
	s0f32 = l32
	s1f32 = l32
	s0f32 = s0f32 * s1f32
	s1f32 = l33
	s2f32 = l33
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l34
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	l32 = s2f32
	s3i32 = l6
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	l33 = s3f32
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l35
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	l34 = s3f32
	s4i32 = l6
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	l35 = s4f32
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = l36
	s1f32 = s1f32 / s2f32
	l27 = s1f32
	s1f32 = -s1f32
	s2f32 = l27
	s3f32 = -1
	s2f32 = s2f32 + s3f32
	s3f32 = l27
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l27 = s1f32
	s2f32 = l27
	s2f32 = float32(math.Abs(float64(s2f32)))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	l27 = s0f32
	s0f32 = l29
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl61
	}
	s0f32 = l27
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl61
	}
	s0i32 = 1
	l8 = s0i32
	s0i32 = 1
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2f32 = l31
	s1f32 = s1f32 + s2f32
	s2f32 = l33
	s1f32 = s1f32 - s2f32
	l31 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl62
	}
	s0i32 = 1
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2f32 = l30
	s1f32 = s1f32 + s2f32
	s2f32 = l35
	s1f32 = s1f32 - s2f32
	l30 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl62
	}
	s0f32 = l31
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l30
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
lbl62:
	l13 = s0i32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f32 = l3
	s0f32 = s0f32 + s1f32
	s1f32 = l32
	s0f32 = s0f32 - s1f32
	l3 = s0f32
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
		goto lbl63
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1f32 = l28
	s0f32 = s0f32 + s1f32
	s1f32 = l34
	s0f32 = s0f32 - s1f32
	l28 = s0f32
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
		goto lbl63
	}
	s0f32 = l3
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l28
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	l8 = s0i32
lbl63:
	s0i32 = l13
	s1i32 = l8
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	s0i32 = s0i32 | s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l7 = s0i32
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = 0
		s1i32 = l7
		s2i32 = l2
		s3i32 = l7
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l1
		s2i32 = l1
		s3i32 = l2
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l1 = s0i32
		goto lbl60
	}
	s0i32 = l8
	s1i32 = l13
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	s0i32 = s0i32 | s1i32
	if s0i32 != 0 {
		goto lbl61
	}
	s0i32 = l7
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 0
	s1i32 = l0
	s2i32 = l0
	s3i32 = l6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l1
	s2i32 = l1
	s3i32 = l6
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s0i32 = l12
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s0i32 = l2
	l0 = s0i32
	s0i32 = l7
	l6 = s0i32
	goto lbl52
lbl61:
	s0f32 = l29
	s1f32 = l27
	if s0f32 < s1f32 {
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
		s0i32 = l7
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l0 = s0i32
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl65
		}
		s0i32 = 0
		s1i32 = l0
		s2i32 = l0
		s3i32 = l1
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l1 = s0i32
		goto lbl65
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl60
	}
	s0i32 = 0
	s1i32 = l7
	s2i32 = l1
	s3i32 = l7
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	goto lbl60
lbl65:
	s0i32 = l2
	l0 = s0i32
	s0i32 = l7
	l6 = s0i32
lbl60:
	s0i32 = l12
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l12 = s0i32
lbl52:
	s0i64 = l25
	s1i64 = 1
	s0i64 = s0i64 + s1i64
	l25 = s0i64
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl50
	}
lbl48:
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	if s0i32 != 0 {
		s0i32 = l2
		f13(ctx, s0i32)
		s0i32 = l4
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l4
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = 0
		l2 = s0i32
	}
	s0i32 = 0
	l11 = s0i32
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l12
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 65533
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l12
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l12
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = 2
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l0
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l4
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l2
		s2i32 = l0
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = f29(ctx, s1i32, s2i32)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l4
	s0i32 = f45(ctx, s0i32)
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])))
		l0 = s0i32
		s0i32 = l5
		s0i32 = f382(ctx, s0i32)
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l2 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l0 = s0i32
		s0i32 = l5
		if s0i32 != 0 {
		lbl75:
			s0i32 = l2
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
			s1i32 = l4
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l0
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			l7 = s1i32
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s0f32 = s0f32 - s1f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 0.01
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
				s0i32 = l2
				s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
				s1i32 = l7
				s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
				s0f32 = s0f32 - s1f32
				s0f32 = float32(math.Abs(float64(s0f32)))
				s1f32 = 0.01
				if s0f32 <= s1f32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl76
				}
			}
			s0i32 = l4
			s0i32 = f45(ctx, s0i32)
			s1i32 = l2
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l2
			s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])))
			l7 = s0i32
			s0i32 = l5
			s0i32 = f382(ctx, s0i32)
			s1i32 = l7
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l0 = s0i32
		lbl76:
			s0i32 = l2
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l2 = s0i32
			s1i32 = l1
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl75
			}
			goto lbl73
			panic("unreachable executed")
			panic("unreachable executed")
		}
	lbl78:
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l0
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0f32 = s0f32 - s1f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 0.01
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
			s0i32 = l2
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
			s1i32 = l7
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			s0f32 = s0f32 - s1f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 0.01
			if s0f32 <= s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl79
			}
		}
		s0i32 = l4
		s0i32 = f45(ctx, s0i32)
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
	lbl79:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l2 = s0i32
		s1i32 = l1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl78
		}
	lbl73:
		s0i32 = l0
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl71
		}
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	goto lbl70
lbl71:
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l6
	s2i32 = l0
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l0 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.01
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl70
	}
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.01
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl70
	}
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl70
	}
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
lbl70:
	s0f32 = 0
	l29 = s0f32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l0 = s0i32
	s1i32 = 3
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l28 = s1f32
		s0f32 = s0f32 - s1f32
		l27 = s0f32
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l31 = s1f32
		s0f32 = s0f32 - s1f32
		l30 = s0f32
		s0i32 = 2
		l2 = s0i32
	lbl82:
		s0f32 = l30
		s1i32 = l6
		s2i32 = l2
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2f32 = l28
		s1f32 = s1f32 - s2f32
		l3 = s1f32
		s0f32 = s0f32 * s1f32
		l32 = s0f32
		s0f32 = l29
		s1f32 = l27
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s3f32 = l31
		s2f32 = s2f32 - s3f32
		l30 = s2f32
		s1f32 = s1f32 * s2f32
		s2f32 = l32
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 + s1f32
		l29 = s0f32
		s0f32 = l3
		l27 = s0f32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l0
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl82
		}
		s0f32 = 0
		s1f32 = 1
		s2f32 = -1
		s3f32 = l29
		s4f32 = 0
		if s3f32 > s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		s2f32 = l29
		s2f32 = float32(math.Abs(float64(s2f32)))
		s3f32 = 5.9604645e-08
		if s2f32 <= s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
	} else {
		s0f32 = 0
	}
	s1i32 = l19
	s1f32 = float32(s1i32)
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l6
	s1i32 = l0
	s0i32 = f563(ctx, s0i32, s1i32)
	l11 = s0i32
lbl21:
	s0i32 = l14
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l10
	f13(ctx, s0i32)
lbl7:
	s0i32 = l9
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l0 = s0i32
	s1i32 = l9
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	f13(ctx, s0i32)
lbl1:
	s0i32 = l9
	s1i32 = 560
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l11
	return s0i32
lbl0:
	s0i32 = l9
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 14655
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 14607
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 14581
	s1i32 = l9
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
