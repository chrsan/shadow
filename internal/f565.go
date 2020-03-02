package internal

import (
	"math"
	"unsafe"
)

func f565(ctx *Context, l0 int32, l1 float32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	var l7 int32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
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
	var l26 float32
	_ = l26
	var l27 float32
	_ = l27
	var l28 float32
	_ = l28
	var l29 float32
	_ = l29
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
	var s8f32 float32
	_ = s8f32
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l11 = s1f32
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	s3i32 = l4
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l9 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0f32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l3 = s0i32
	s1i32 = 192
	s0i32 = s0i32 & s1i32
	s1i32 = 128
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l2
		s1i32 = f67(ctx, s1i32)
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0f32 = l8
	s1f32 = l9
	s0f32 = s0f32 + s1f32
	l21 = s0f32
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l10 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l11 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l8 = s0f32
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1064514355
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l7
		s1f32 = l21
		s2f32 = l8
		s3f32 = l21
		s2f32 = s2f32 - s3f32
		l12 = s2f32
		s1f32 = s1f32 / s2f32
		l9 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = s1f32
		s0i32 = l6
		s1i32 = l7
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = l7
		s3i32 = 92
		s2i32 = s2i32 + s3i32
		s3i32 = l7
		s4i32 = 48
		s3i32 = s3i32 + s4i32
		s4f32 = l9
		s5f32 = 0.95
		if s4f32 < s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s3f32 = l9
		s4f32 = 0
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l9 = s1f32
		s2f32 = l1
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l7
		s1i32 = 1065353216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1073322394
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l7
		s1f32 = l8
		s2f32 = l12
		s1f32 = s1f32 / s2f32
		l1 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = s1f32
		s0i32 = l5
		s1f32 = l10
		s2f32 = l9
		s2f32 = -s2f32
		l8 = s2f32
		s1f32 = s1f32 * s2f32
		l9 = s1f32
		s2f32 = 0
		if s1f32 != s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2f32 = l11
		s3f32 = l8
		s2f32 = s2f32 * s3f32
		l8 = s2f32
		s3f32 = 0
		if s2f32 != s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s1i32 = s1i32 | s2i32
		s2i32 = 18
		s3i32 = 16
		s4i32 = l7
		s5i32 = 8
		s4i32 = s4i32 + s5i32
		s5i32 = l7
		s6i32 = 92
		s5i32 = s5i32 + s6i32
		s6i32 = l7
		s7i32 = 48
		s6i32 = s6i32 + s7i32
		s7f32 = l1
		s8f32 = 1.95
		if s7f32 < s8f32 {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i32 = s5i32
		} else {
			s5i32 = s6i32
		}
		s6f32 = l1
		s7f32 = 1
		if s6f32 < s7f32 {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4i32 = s4i32
		} else {
			s4i32 = s5i32
		}
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		l1 = s4f32
		s5f32 = 1
		if s4f32 != s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 1065353216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l5
		s1f32 = l8
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l5
		s1f32 = l1
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l5
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l5
		s1f32 = l9
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l5
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l5
		s1f32 = l1
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l5
		s1i32 = l2
		f199(ctx, s0i32, s1i32)
		goto lbl2
	}
	s0i32 = 0
	l3 = s0i32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s1i32 = l7
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	f269(ctx, s0i32, s1i32)
	s0i32 = l2
	s1i32 = l7
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = 48
	s2i32 = s2i32 + s3i32
	s3i32 = 4
	f55(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l7
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s0i32 = f1408(ctx, s0i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0f32
	s1f32 = l8
	s2f32 = l10
	s3i32 = l4
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s2f32 = s2f32 * s3f32
	l12 = s2f32
	s3f32 = l11
	s4i32 = l4
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s3f32 = s3f32 * s4f32
	l15 = s3f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	l16 = s1f32
	s0f32 = s0f32 - s1f32
	l17 = s0f32
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l9
	s1f32 = l8
	s2f32 = l15
	s3f32 = l10
	s4i32 = l4
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
	s3f32 = s3f32 * s4f32
	l10 = s3f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	l13 = s1f32
	s0f32 = s0f32 - s1f32
	l19 = s0f32
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
		goto lbl1
	}
	s0f32 = l9
	s1f32 = l8
	s2f32 = l10
	s3f32 = l11
	s4i32 = l4
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s3f32 = s3f32 * s4f32
	l10 = s3f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	l11 = s1f32
	s0f32 = s0f32 - s1f32
	l15 = s0f32
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l9
	s1f32 = l8
	s2f32 = l12
	s3f32 = l10
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	l10 = s1f32
	s0f32 = s0f32 - s1f32
	l12 = s0f32
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	l9 = s0f32
	s1f32 = l16
	s2f32 = l17
	s1f32 = s1f32 / s2f32
	l14 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l8 = s2f32
	s3f32 = l9
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l9 = s0f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)]))
	l16 = s1f32
	s2f32 = l11
	s3f32 = l15
	s2f32 = s2f32 / s3f32
	l20 = s2f32
	s3f32 = l8
	s4f32 = l16
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l15 = s1f32
	s0f32 = s0f32 - s1f32
	l18 = s0f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	l11 = s1f32
	s2f32 = l10
	s3f32 = l12
	s2f32 = s2f32 / s3f32
	l12 = s2f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l10 = s3f32
	s4f32 = l11
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l16 = s1f32
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
	l11 = s2f32
	s3f32 = l13
	s4f32 = l19
	s3f32 = s3f32 / s4f32
	l17 = s3f32
	s4f32 = l10
	s5f32 = l11
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 - s3f32
	l11 = s2f32
	s1f32 = s1f32 - s2f32
	l19 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
	l13 = s1f32
	s2f32 = l17
	s3f32 = l8
	s4f32 = l13
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l17 = s1f32
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+76)]))
	l13 = s2f32
	s3f32 = l12
	s4f32 = l8
	s5f32 = l13
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 - s3f32
	l12 = s2f32
	s1f32 = s1f32 - s2f32
	l22 = s1f32
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+64)]))
	l8 = s2f32
	s3f32 = l20
	s4f32 = l10
	s5f32 = l8
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 - s3f32
	l13 = s2f32
	s3i32 = l7
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+48)]))
	l8 = s3f32
	s4f32 = l14
	s5f32 = l10
	s6f32 = l8
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l10 = s3f32
	s2f32 = s2f32 - s3f32
	l14 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l8 = s0f32
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
		goto lbl1
	}
	s0i32 = l5
	s1i64 = 550821167104
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l5
	s1f32 = l22
	s2f32 = l10
	s3f32 = l15
	s2f32 = s2f32 * s3f32
	s3f32 = l9
	s4f32 = l13
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 - s3f32
	l20 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l18
	s3f32 = l11
	s4f32 = l12
	s3f32 = s3f32 * s4f32
	s4f32 = l17
	s5f32 = l16
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l18 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = l8
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l5
	s1f32 = l18
	s2f32 = l14
	s1f32 = s1f32 * s2f32
	s2f32 = l19
	s3f32 = l20
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = l8
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l5
	s1f32 = l9
	s2f32 = l12
	s1f32 = s1f32 - s2f32
	l22 = s1f32
	s2f32 = l13
	s3f32 = l11
	s2f32 = s2f32 - s3f32
	l24 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l17
	s3f32 = l15
	s2f32 = s2f32 - s3f32
	l25 = s2f32
	s3f32 = l16
	s4f32 = l10
	s3f32 = s3f32 - s4f32
	l19 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l14 = s1f32
	s1f32 = -s1f32
	s2f32 = l14
	s3f32 = l11
	s4f32 = l10
	s3f32 = s3f32 - s4f32
	s4f32 = l9
	s5f32 = l11
	s4f32 = s4f32 * s5f32
	s5f32 = l17
	s6f32 = l10
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 - s5f32
	l18 = s4f32
	s5f32 = l15
	s6f32 = l12
	s5f32 = s5f32 - s6f32
	l26 = s5f32
	s4f32 = s4f32 * s5f32
	s5f32 = l17
	s6f32 = l9
	s5f32 = s5f32 - s6f32
	l14 = s5f32
	s6f32 = l12
	s7f32 = l13
	s6f32 = s6f32 * s7f32
	s7f32 = l15
	s8f32 = l16
	s7f32 = s7f32 * s8f32
	s6f32 = s6f32 - s7f32
	l23 = s6f32
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 - s5f32
	l20 = s4f32
	s5f32 = l9
	s4f32 = s4f32 - s5f32
	l27 = s4f32
	s3f32 = s3f32 * s4f32
	s4f32 = l14
	s5f32 = l10
	s6f32 = l11
	s5f32 = s5f32 - s6f32
	l28 = s5f32
	s6f32 = l23
	s5f32 = s5f32 * s6f32
	s6f32 = l18
	s7f32 = l16
	s8f32 = l13
	s7f32 = s7f32 - s8f32
	l23 = s7f32
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 - s6f32
	l18 = s5f32
	s6f32 = l10
	s5f32 = s5f32 - s6f32
	l29 = s5f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	l2 = s3i32
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = l8
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l5
	s1f32 = l14
	s2f32 = l23
	s1f32 = s1f32 * s2f32
	s2f32 = l28
	s3f32 = l26
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l14 = s1f32
	s1f32 = -s1f32
	s2f32 = l14
	s3f32 = l19
	s4f32 = l27
	s3f32 = s3f32 * s4f32
	s4f32 = l12
	s5f32 = l9
	s4f32 = s4f32 - s5f32
	s5f32 = l29
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	s4f32 = 0
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	l3 = s3i32
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = l8
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l5
	s1f32 = l25
	s2f32 = l10
	s3f32 = l12
	s2f32 = s2f32 * s3f32
	s3f32 = l9
	s4f32 = l16
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 - s3f32
	l9 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l22
	s3f32 = l11
	s4f32 = l15
	s3f32 = s3f32 * s4f32
	s4f32 = l17
	s5f32 = l13
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l10 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	s1f32 = -s1f32
	s2f32 = l11
	s3i32 = l2
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = l8
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l5
	s1f32 = l20
	s1f32 = -s1f32
	s2f32 = l20
	s3i32 = l3
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = l8
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l5
	s1f32 = l19
	s2f32 = l10
	s1f32 = s1f32 * s2f32
	s2f32 = l9
	s3f32 = l24
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l9 = s1f32
	s1f32 = -s1f32
	s2f32 = l9
	s3i32 = l2
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = l8
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l5
	s1f32 = l18
	s1f32 = -s1f32
	s2f32 = l18
	s3i32 = l3
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = l8
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0f32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l10 = s0f32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l9 = s0f32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l11 = s0f32
	s0i32 = l7
	s1i64 = 550821167104
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l7
	s1f32 = 2
	s2f32 = l11
	s3f32 = l9
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	l11 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l7
	s1f32 = -1
	s2f32 = l9
	s3f32 = l11
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l7
	s1f32 = 2
	s2f32 = l10
	s3f32 = l8
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	l9 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l7
	s1f32 = -1
	s2f32 = l8
	s3f32 = l9
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l5
	s1i32 = l7
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	f199(ctx, s0i32, s1i32)
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0f32
	s0i32 = l7
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 1064514355
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l7
	s1f32 = l21
	s2f32 = l8
	s3f32 = l21
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	l8 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = s1f32
	s0i32 = l6
	s1i32 = l7
	s2i32 = 92
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = 84
	s2i32 = s2i32 + s3i32
	s3i32 = l7
	s4i32 = 88
	s3i32 = s3i32 + s4i32
	s4f32 = l8
	s5f32 = 0.95
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3f32 = l8
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = l1
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
lbl2:
	s0i32 = 1
	l3 = s0i32
lbl1:
	s0i32 = l7
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l3
	return s0i32
}
