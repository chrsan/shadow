package internal

import (
	"math"
	"unsafe"
)

func f557(ctx *Context, l0 int32, l1 float32, l2 float32, l3 int32) int32 {
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
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int64
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
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l0
		f1400(ctx, s0i32)
	}
	s0i32 = l0
	s1i32 = 116
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	l5 = s1i32
	s2i32 = l5
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s0f32 = f666(ctx, s0i32, s1i32, s2i32)
	l19 = s0f32
	s0i32 = l4
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+72)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
	l8 = s0i32
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l6 = s0i32
	lbl2:
		s0i32 = l4
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		l7 = s1i32
		s2i32 = l6
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l7
		s2i32 = 0
		s3i32 = l6
		s4i32 = 1
		s3i32 = s3i32 + s4i32
		l5 = s3i32
		s4i32 = l6
		s5i32 = l8
		s6i32 = -1
		s5i32 = s5i32 + s6i32
		if s4i32 == s5i32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l10
		s1i32 = l4
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = 40
		s2i32 = s2i32 + s3i32
		s0f32 = f666(ctx, s0i32, s1i32, s2i32)
		l18 = s0f32
		s1f32 = l19
		s2f32 = l18
		s3f32 = l19
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
		l19 = s0f32
		s0i32 = l5
		l6 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
		l8 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l0
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	l10 = s0i32
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = -16777216
	l9 = s0i32
	s0f32 = l1
	s1f32 = 0.00024414062
	if s0f32 > s1f32 {
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
		s0f32 = l1
		l18 = s0f32
		goto lbl3
	}
	s0f32 = l19
	s1f32 = l1
	s2f32 = 0.01
	s1f32 = s1f32 + s2f32
	l18 = s1f32
	s2f32 = l18
	s1f32 = s1f32 * s2f32
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0f32 = l1
		l18 = s0f32
		goto lbl5
	}
	s0f32 = l19
	s0f32 = float32(math.Sqrt(float64(s0f32)))
	s1f32 = -0.01
	s0f32 = s0f32 + s1f32
	l18 = s0f32
	s1f32 = l1
	s0f32 = s0f32 / s1f32
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	s1f32 = 128
	s0f32 = s0f32 * s1f32
	l1 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l1
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l1
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl7
	}
	s0i32 = 0
lbl7:
	s1i32 = 16711680
	s0i32 = s0i32 * s1i32
	s1i32 = -16777216
	s0i32 = s0i32 & s1i32
	l9 = s0i32
lbl5:
	s0i32 = l10
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l8
	s2f32 = l18
	s3i32 = l4
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s0i32 = f1407(ctx, s0i32, s1i32, s2f32, s3i32)
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+159)] = uint8(s1i32)
lbl3:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+157)])
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 28
		s0i32 = s0i32 + s1i32
		s0i32 = f45(ctx, s0i32)
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 40
		s0i32 = s0i32 + s1i32
		s0i32 = f88(ctx, s0i32)
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l7 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l20 = s0f32
	s0i32 = l7
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	l11 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l15 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l6 = s1i32
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l19 = s0f32
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+160)]))
	l1 = s1f32
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3i32 = l5
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
	s0i32 = l4
	s1f32 = l1
	s2f32 = l19
	s3f32 = l20
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
	s0i32 = 0
	s1i32 = l4
	s2i32 = 40
	s1i32 = s1i32 + s2i32
	s1i32 = f227(ctx, s1i32)
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl10
	}
	s0i32 = l0
	s1i32 = 140
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l1 = s0f32
	s0i32 = l0
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2f32 = l2
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = s1f32
	s0i32 = l0
	s1f32 = l1
	s2f32 = l2
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = s1f32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l17 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = l17
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+188)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+164)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s2i32 = 148
	s1i32 = s1i32 + s2i32
	s2i32 = l9
	s3i32 = l4
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s4i32 = l10
	s5f32 = l18
	s6f32 = 0.00024414062
	if s5f32 > s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	l13 = s3i32
	s4i32 = l0
	s5i32 = 164
	s4i32 = s4i32 + s5i32
	s0i32 = f556(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l0
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	l14 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+157)])
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+136)]))
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l25 = s1f32
	s0f32 = s0f32 - s1f32
	l21 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)]))
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l26 = s1f32
	s0f32 = s0f32 - s1f32
	l22 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
	l16 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	l8 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+172)]))
	l10 = s0i32
	l6 = s0i32
	s0i32 = l0
lbl14:
	s1f32 = l21
	s2f32 = l26
	s3i32 = l8
	s4i32 = l6
	s5i32 = 3
	s4i32 = s4i32 << (uint32(s5i32) & 31)
	l7 = s4i32
	s3i32 = s3i32 + s4i32
	l5 = s3i32
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s2f32 = s2f32 - s3f32
	l20 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l22
	s3f32 = l25
	s4i32 = l5
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s3f32 = s3f32 - s4f32
	l19 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l23 = s1f32
	s1f32 = l21
	s2i32 = l7
	s3i32 = l16
	s2i32 = s2i32 + s3i32
	l5 = s2i32
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l18 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l22
	s3i32 = l5
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	l1 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l24 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 0.00024414062
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l23
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 0.00024414062
		if s1f32 <= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 != 0 {
			goto lbl15
		}
		s1i32 = 0
		goto lbl12
	}
	s1f32 = l23
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl15
	}
	s1f32 = l23
	s2f32 = l24
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl15
	}
	s1f32 = l20
	s2f32 = l1
	s1f32 = s1f32 * s2f32
	s2f32 = l19
	s3f32 = l18
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l1 = s1f32
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl15
	}
	s1f32 = l1
	s2f32 = l24
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl13
	}
lbl15:
	s1i32 = l0
	s2i32 = l6
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+84)]))
	s2i32 = i32RemS(s2i32, s3i32)
	l6 = s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+172)])) = uint32(s2i32)
	s1i32 = l6
	s2i32 = l10
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl14
	}
	s1i32 = 0
	goto lbl12
lbl13:
	s1i32 = l14
	s1i32 = f45(ctx, s1i32)
	l5 = s1i32
	s2f32 = l25
	s3f32 = l21
	s4f32 = l1
	s5f32 = l24
	s4f32 = s4f32 / s5f32
	l1 = s4f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = s2f32
	s1i32 = l5
	s2f32 = l26
	s3f32 = l22
	s4f32 = l1
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = s2f32
	s1i32 = l0
	s2i32 = 40
	s1i32 = s1i32 + s2i32
	s1i32 = f88(ctx, s1i32)
	s2i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = 1
lbl12:
	l6 = s1i32
	ctx.Mem[int(s0i32+177)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l6
	ctx.Mem[int(s0i32+176)] = uint8(s1i32)
lbl11:
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)]))
	l20 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)]))
	l19 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)]))
	l18 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)]))
	l1 = s0f32
	s0i32 = l14
	s0i32 = f45(ctx, s0i32)
	l5 = s0i32
	s1f32 = l20
	s2f32 = l1
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l5
	s1f32 = l18
	s2f32 = l19
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s0i32 = f88(ctx, s0i32)
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	s2i32 = l12
	s3i32 = l9
	s4i32 = l13
	s5i32 = 0
	s6i32 = l3
	f555(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	s0i32 = l11
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l6 = s0i32
	lbl19:
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)]))
		l19 = s0f32
		s0i32 = l6
		s1i32 = 3
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l7 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l18 = s0f32
		s0i32 = l4
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+160)]))
		l1 = s1f32
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+188)]))
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
		s0i32 = l4
		s1f32 = l1
		s2f32 = l19
		s3f32 = l18
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
		s0i32 = l4
		s1i32 = 40
		s0i32 = s0i32 + s1i32
		s0i32 = f227(ctx, s0i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		s0i32 = l4
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2f32 = l2
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l4
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		s2f32 = l2
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s1i32 = l4
		s2f32 = l2
		s3i32 = 1
		s0i32 = f554(ctx, s0i32, s1i32, s2f32, s3i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		s2i32 = l7
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = l9
		s4i32 = l13
		s5i32 = l6
		s6i32 = l15
		if s5i32 == s6i32 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		s6i32 = l3
		f555(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l11
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
	}
	s0i32 = l0
	s1i32 = l12
	s2f32 = l2
	s3i32 = 0
	s0i32 = f554(ctx, s0i32, s1i32, s2f32, s3i32)
	l3 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+177)])
	l6 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])))
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
		l5 = s0i32
		s0i32 = l6
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = 52
			s0i32 = s0i32 + s1i32
			s1i32 = 3
			s0i32 = f89(ctx, s0i32, s1i32)
			l0 = s0i32
			s1i32 = l5
			s2i32 = 2
			s1i32 = s1i32 + s2i32
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint16(s1i32)
			s0i32 = l0
			s1i32 = l3
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
			s0i32 = l0
			s1i32 = l5
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
			goto lbl20
		}
		s0i32 = l0
		s1i32 = 52
		s0i32 = s0i32 + s1i32
		s1i32 = 3
		s0i32 = f89(ctx, s0i32, s1i32)
		l0 = s0i32
		s1i32 = l5
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint16(s1i32)
		s0i32 = l0
		s1i32 = l3
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
		s0i32 = l0
		s1i32 = l5
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = 1
		goto lbl10
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l5 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
	l0 = s0i32
	s0i32 = l6
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 3
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l5
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		goto lbl20
	}
	s0i32 = l0
	s1i32 = 3
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
lbl20:
	s0i32 = 1
	goto lbl10
lbl17:
	s0i32 = 0
lbl10:
	l8 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	f13(ctx, s0i32)
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l8
	return s0i32
}
