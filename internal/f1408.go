package internal

import (
	"math"
	"unsafe"
)

func f1408(ctx *Context, l0 int32, l1 int32, l2 float32, l3 int32) int32 {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
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
	var l14 float32
	_ = l14
	var l15 float32
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
	var s1i64 int64
	_ = s1i64
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
	s0i32 = ctx.g0
	s1i32 = 2608
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s1i32 = 3
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s1i32 = 65535
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l2
	s1f32 = -0.00024414062
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l2
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
	s0f32 = l2
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
	lbl3:
		s0i32 = 1
		l9 = s0i32
		s0i32 = l3
		s0i32 = f45(ctx, s0i32)
		s1i32 = l0
		s2i32 = l4
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		goto lbl1
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l15 = s1f32
	s0f32 = s0f32 - s1f32
	l18 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l16 = s1f32
	s0f32 = s0f32 - s1f32
	l14 = s0f32
	s0i32 = 2
	l4 = s0i32
lbl4:
	s0f32 = l14
	s1i32 = l0
	s2i32 = l4
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = l15
	s1f32 = s1f32 - s2f32
	l17 = s1f32
	s0f32 = s0f32 * s1f32
	l19 = s0f32
	s0f32 = l20
	s1f32 = l18
	s2i32 = l6
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3f32 = l16
	s2f32 = s2f32 - s3f32
	l14 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l19
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 + s1f32
	l20 = s0f32
	s0f32 = l17
	l18 = s0f32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f32 = l20
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 5.9604645e-08
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l4 = s0i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 0
	s1i32 = 1
	s2i32 = -1
	s3f32 = l20
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
	s2i32 = l4
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l12 = s0i32
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
			s0i32 = l5
			s1i32 = l5
			s2i32 = 40
			s1i32 = s1i32 + s2i32
			s2i32 = 4
			s1i32 = s1i32 | s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
			goto lbl9
		}
		s0i32 = 0
		l6 = s0i32
		s0i32 = l5
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = 0
		l1 = s0i32
		goto lbl7
	}
	s0i32 = l5
	s1i32 = l1
	s2i32 = 40
	s1i32 = f50(ctx, s1i32, s2i32)
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
lbl9:
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = l12
	s0f32 = float32(s0i32)
	s1f32 = l2
	s0f32 = s0f32 * s1f32
	l18 = s0f32
	s0i32 = 0
	l4 = s0i32
lbl12:
	s0i32 = l0
	s1i32 = l4
	l6 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0f32
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l17 = s1f32
	s0f32 = s0f32 * s1f32
	l14 = s0f32
	s1f32 = l14
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = 0
	l9 = s0i32
	s0i32 = 0
	s1i32 = 1
	s2i32 = -1
	s3f32 = l2
	s4i32 = l0
	s5i32 = l10
	s6i32 = 3
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	l11 = s4i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	l14 = s4f32
	s3f32 = s3f32 - s4f32
	s4i32 = l0
	s5i32 = 0
	s6i32 = l6
	s7i32 = 1
	s6i32 = s6i32 + s7i32
	l4 = s6i32
	s7i32 = l1
	s8i32 = l4
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
	l13 = s5i32
	s6i32 = 3
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	l8 = s4i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	l15 = s4f32
	s5i32 = l11
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+4)]))
	l16 = s5f32
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s4f32 = l17
	s5f32 = l16
	s4f32 = s4f32 - s5f32
	s5i32 = l8
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	l16 = s5f32
	s6f32 = l14
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l14 = s3f32
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
	s2f32 = l14
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
	s1i32 = l12
	s0i32 = s0i32 * s1i32
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l5
	s1f32 = l16
	s2f32 = l2
	s1f32 = s1f32 - s2f32
	l2 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
	s0i32 = l5
	s1f32 = l15
	s2f32 = l17
	s1f32 = s1f32 - s2f32
	l17 = s1f32
	s1f32 = -s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l5
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1f32 = l18
	s0i32 = f113(ctx, s0i32, s1f32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l11 = s0i32
	s1i32 = l6
	s2i32 = 40
	s1i32 = s1i32 * s2i32
	l8 = s1i32
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s1i32 = l11
	s2i32 = l13
	s3i32 = 40
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = l11
	s2i32 = l10
	s3i32 = 40
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l14 = s0f32
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l15 = s0f32
	s0i32 = l9
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l5
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l9
	s1f32 = l14
	s2f32 = l15
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	s1i32 = l8
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1f32 = l17
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l7
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	s1i32 = l8
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i64 = 4286578687
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l6
	l10 = s0i32
	s0i32 = l1
	s1i32 = l4
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l4 = s0i32
lbl8:
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		l6 = s0i32
		goto lbl7
	}
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		l6 = s0i32
		goto lbl7
	}
	s0i32 = l1
	s1i32 = l1
	s0i32 = s0i32 * s1i32
	l11 = s0i32
	s0i32 = 0
	l9 = s0i32
	s0i32 = 0
	l10 = s0i32
	s0i32 = l4
	l6 = s0i32
lbl15:
	s0i32 = l10
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s1i32 = l11
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 32
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 28
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = 24
	s4i32 = s4i32 + s5i32
	s0i32 = f565(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 != 0 {
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		s1i32 = l0
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
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l8 = s0i32
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l7 = s0i32
			s1i32 = l8
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = 0
			s1i32 = l7
			s2i32 = l0
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
			s1i32 = l6
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
			l6 = s0i32
			s0i32 = l1
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s0i32 = l4
			l7 = s0i32
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l4 = s0i32
			goto lbl16
		}
		s0i32 = l4
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
			goto lbl19
		}
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		s1i32 = l4
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
			goto lbl19
		}
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		s1i32 = l4
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
			goto lbl7
		}
	lbl19:
		s0i32 = l4
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l7 = s0i32
		goto lbl16
	}
	s0i32 = 0
	s1i32 = 1
	s2i32 = -1
	s3i32 = l4
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	l2 = s3f32
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	l17 = s4f32
	s5i32 = l4
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
	l18 = s5f32
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s4i32 = l4
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
	l14 = s4f32
	s5i32 = l0
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+8)]))
	l15 = s5f32
	s6i32 = l4
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+8)]))
	l16 = s6f32
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l19 = s3f32
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
	s2f32 = l19
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
	s1i32 = l12
	s0i32 = s0i32 * s1i32
	l7 = s0i32
	s1i32 = -1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l7
	s1i32 = 0
	s2i32 = 1
	s3i32 = -1
	s4f32 = l2
	s5f32 = l17
	s6i32 = l0
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+20)]))
	s5f32 = s5f32 + s6f32
	s6f32 = l18
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l14
	s6f32 = l15
	s7i32 = l0
	s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+16)]))
	s6f32 = s6f32 + s7f32
	s7f32 = l16
	s6f32 = s6f32 - s7f32
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 - s5f32
	l2 = s4f32
	s5f32 = 0
	if s4f32 > s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3f32 = l2
	s3f32 = float32(math.Abs(float64(s3f32)))
	s4f32 = 5.9604645e-08
	if s3f32 <= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l12
	s1i32 = s1i32 * s2i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 0
	s1i32 = l7
	s2i32 = l0
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
	s1i32 = l6
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
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l4
	l7 = s0i32
	goto lbl20
lbl21:
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 0
	s1i32 = l7
	s2i32 = l4
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
	s1i32 = l6
	s2i32 = l4
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
	l6 = s0i32
lbl20:
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	l4 = s0i32
lbl16:
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l4
	s1i32 = l7
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l8 = s0i32
	s0i32 = l4
	l0 = s0i32
	s0i32 = l7
	l4 = s0i32
	s0i32 = l8
	if s0i32 != 0 {
		goto lbl15
	}
lbl7:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		f13(ctx, s0i32)
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = 0
		l0 = s0i32
	}
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = 0
	l4 = s0i32
	s0i32 = l1
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 2
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s2i32 = l1
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = f29(ctx, s1i32, s2i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl23:
	s0i32 = l3
	s0i32 = f45(ctx, s0i32)
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl27:
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l4
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
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
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
			if s0i32 != 0 {
				goto lbl28
			}
		}
		s0i32 = l3
		s0i32 = f45(ctx, s0i32)
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
	lbl28:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l1 = s0i32
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl27
		}
		s0i32 = l4
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl25
		}
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	goto lbl24
lbl25:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l0
	s2i32 = l4
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l1 = s1i32
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
		goto lbl24
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l1
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
		goto lbl24
	}
	s0i32 = l3
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
lbl24:
	s0i32 = 0
	l9 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
	s1i32 = 3
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l21 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l22 = s1f32
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	l2 = s0f32
	s1f32 = l2
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l14 = s0f32
	s1f32 = l21
	s0f32 = s0f32 - s1f32
	l16 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l15 = s0f32
	s1f32 = l22
	s0f32 = s0f32 - s1f32
	l19 = s0f32
	s0f32 = l21
	s1i32 = l0
	s2i32 = l3
	s3i32 = -1
	s2i32 = s2i32 + s3i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	l23 = s0f32
	s0f32 = l22
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	l2 = s0f32
	s0f32 = l21
	s1f32 = l21
	s0f32 = s0f32 - s1f32
	l24 = s0f32
	s0f32 = l22
	s1f32 = l22
	s0f32 = s0f32 - s1f32
	l25 = s0f32
	s0i32 = 1
	l1 = s0i32
	s0f32 = 0
	l20 = s0f32
	s0i32 = 0
	l4 = s0i32
	s0f32 = 0
	l18 = s0f32
lbl30:
	s0f32 = l20
	s1f32 = l2
	s2f32 = l16
	l17 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l23
	s3f32 = l19
	l2 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l16 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0f32 = l18
	s1f32 = l14
	s2f32 = l21
	s1f32 = s1f32 - s2f32
	l23 = s1f32
	s2f32 = l25
	s1f32 = s1f32 * s2f32
	s2f32 = l15
	s3f32 = l22
	s2f32 = s2f32 - s3f32
	l25 = s2f32
	s3f32 = l24
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l19 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = 1
	l9 = s0i32
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s0i32 = i32RemS(s0i32, s1i32)
	l1 = s0i32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0f32 = l16
	s1f32 = l20
	s2f32 = l16
	s3f32 = 0
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l20 = s0f32
	s0f32 = l19
	s1f32 = l18
	s2f32 = l19
	s3f32 = 0
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l18 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l24 = s0f32
	s1f32 = l15
	s0f32 = s0f32 - s1f32
	l19 = s0f32
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l15 = s0f32
	s1f32 = l14
	s0f32 = s0f32 - s1f32
	l16 = s0f32
	s0f32 = l15
	l14 = s0f32
	s0f32 = l24
	l15 = s0f32
	s0f32 = l23
	l24 = s0f32
	s0f32 = l17
	l23 = s0f32
	s0i32 = l0
	s1i32 = l4
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	l17 = s0f32
	s1f32 = l17
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
lbl6:
	s0i32 = 0
	l9 = s0i32
lbl5:
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l0 = s0i32
	s1i32 = l5
	s2i32 = 40
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
	s0i32 = l5
	s1i32 = 2608
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l9
	return s0i32
lbl0:
	s0i32 = l5
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 14655
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 14607
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 14581
	s1i32 = l5
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
