package internal

import (
	"math"
	"unsafe"
)

func f1041(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l35 int32
	_ = l35
	var l36 int32
	_ = l36
	var l37 int32
	_ = l37
	var l38 int32
	_ = l38
	var l39 int32
	_ = l39
	var l40 int32
	_ = l40
	var l41 int64
	_ = l41
	var l42 int64
	_ = l42
	var l43 float64
	_ = l43
	var l44 float64
	_ = l44
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
	var s9i32 int32
	_ = s9i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	s0i32 = ctx.g0
	s1i32 = 1200
	s0i32 = s0i32 - s1i32
	l16 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l43 = s0f64
	s1f64 = 2
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l44 = s0f64
	s1f64 = 2
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l16
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1f64 = l43
	s0i32 = f483(ctx, s0i32, s1f64)
	l9 = s0i32
	s0i32 = l16
	s1i32 = 1144
	s0i32 = s0i32 + s1i32
	s1f64 = l44
	s0i32 = f483(ctx, s0i32, s1f64)
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l5 = s0i32
	s0i32 = l9
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l11 = s0i32
	if s0i32 != 0 {
		s0i32 = l9
		s1i32 = l11
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l14 = s0i32
	lbl3:
		s0i32 = l16
		s1i32 = 1132
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l9
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2f64 = 65536
		s1f64 = s1f64 * s2f64
		s1f64 = f0(ctx, s1f64)
		l43 = s1f64
		s2f64 = 4.294967296e+09
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2f64 = l43
		s3f64 = 0
		if s2f64 >= s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s1i32 = s1i32 & s2i32
		if s1i32 != 0 {
			s1f64 = l43
			s1i32 = int32(uint32(math.Trunc(s1f64)))
			goto lbl4
		}
		s1i32 = 0
	lbl4:
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l9
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s1i32 = l14
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l5
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l11
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s0i32 = l5
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l5
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = 0
		l9 = s0i32
	lbl7:
		s0i32 = l16
		s1i32 = 1120
		s0i32 = s0i32 + s1i32
		s1i32 = l9
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2f64 = 65536
		s1f64 = s1f64 * s2f64
		s1f64 = f0(ctx, s1f64)
		l43 = s1f64
		s2f64 = 4.294967296e+09
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2f64 = l43
		s3f64 = 0
		if s2f64 >= s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s1i32 = s1i32 & s2i32
		if s1i32 != 0 {
			s1f64 = l43
			s1i32 = int32(uint32(math.Trunc(s1f64)))
			goto lbl8
		}
		s1i32 = 0
	lbl8:
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l9
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l4
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
	}
	s0i32 = l16
	s1i32 = 1088
	s0i32 = s0i32 + s1i32
	s1i32 = l12
	s2i32 = l8
	s3i32 = l2
	f486(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l3
	s1i32 = l16
	s2i32 = 1112
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l16
	s2i32 = 1104
	s1i32 = s1i32 + s2i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l16
	s2i32 = 1096
	s1i32 = s1i32 + s2i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l16
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+1088)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = 0
		l12 = s0i32
		s0i32 = 0
		l8 = s0i32
		goto lbl10
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l20 = s0i32
	s1i32 = 4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l10 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0i32 = s0i32 - s1i32
	l14 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l15 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l17 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l1 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l2 = s0i32
	s0i32 = l20
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl17
	case 1:
		goto lbl14
	case 2:
		goto lbl16
	case 3:
		goto lbl15
	default:
		goto lbl13
	}
lbl17:
	s0i32 = 0
	s1i32 = 8
	s2i32 = l8
	s3i32 = l16
	s4i32 = 1120
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = l2
	s6i32 = l4
	s7i32 = l14
	s8i32 = l9
	s9i32 = l12
	s8i32 = s8i32 + s9i32
	s9i32 = l1
	f286(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32)
	goto lbl12
lbl16:
	s0i32 = 385
	s1i32 = 32
	s2i32 = l8
	s3i32 = l16
	s4i32 = 1120
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = l2
	s6i32 = l4
	s7i32 = l14
	s8i32 = l9
	s9i32 = l12
	s8i32 = s8i32 + s9i32
	s9i32 = l1
	f286(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32)
	goto lbl12
lbl15:
	s0i32 = 386
	s1i32 = 16
	s2i32 = l8
	s3i32 = l16
	s4i32 = 1120
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = l2
	s6i32 = l4
	s7i32 = l14
	s8i32 = l9
	s9i32 = l12
	s8i32 = s8i32 + s9i32
	s9i32 = l1
	f286(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32)
	goto lbl12
lbl14:
	s0i32 = l16
	s1i32 = 943
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l16
	s1i32 = 3924
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l16
	s1i32 = 3874
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 3848
	s1i32 = l16
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl13:
	s0i32 = 387
	s1i32 = 1
	s2i32 = l8
	s3i32 = l16
	s4i32 = 1120
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = l2
	s6i32 = l4
	s7i32 = l14
	s8i32 = l9
	s9i32 = l12
	s8i32 = s8i32 + s9i32
	s9i32 = l1
	f286(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32)
lbl12:
	s0i32 = l11
	s1i32 = -2
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 21556
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l16
	s2i32 = 1132
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l2 = s2i32
	s3i32 = l12
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s4i32 = l4
	s5i32 = l2
	s6i32 = l1
	s7i32 = l15
	s8i32 = l17
	s7i32 = s7i32 - s8i32
	s8i32 = l6
	s9i32 = l10
	s8i32 = s8i32 - s9i32
	f1036(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32)
lbl10:
	s0i32 = l0
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l12
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	goto lbl0
lbl1:
	s0i32 = l16
	s1i32 = 1056
	s0i32 = s0i32 + s1i32
	s1i32 = l16
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s2i32 = 1024
	s3i32 = 1024
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
	l33 = s0i32
	s0i32 = l1
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f64 = 3
	s0f64 = s0f64 * s1f64
	s1f64 = 2.5066282746310002
	s0f64 = s0f64 * s1f64
	s1f64 = 0.25
	s0f64 = s0f64 * s1f64
	s1f64 = 0.5
	s0f64 = s0f64 + s1f64
	s0f64 = math.Floor(s0f64)
	l43 = s0f64
	s0f64 = math.Abs(s0f64)
	s1f64 = 2.147483648e+09
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l43
		s0i32 = int32(math.Trunc(s0f64))
		goto lbl18
	}
	s0i32 = -2147483648
lbl18:
	l4 = s0i32
	s1i32 = 1
	s2i32 = l4
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
	l4 = s0i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	l14 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 2
		s0i32 = i32DivS(s0i32, s1i32)
		s1i32 = 3
		s0i32 = s0i32 * s1i32
		goto lbl20
	}
	s0i32 = l4
	s1i32 = 1
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 3
	s0i32 = s0i32 * s1i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
lbl20:
	l32 = s0i32
	s0i32 = l1
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f64 = 3
	s0f64 = s0f64 * s1f64
	s1f64 = 2.5066282746310002
	s0f64 = s0f64 * s1f64
	s1f64 = 0.25
	s0f64 = s0f64 * s1f64
	s1f64 = 0.5
	s0f64 = s0f64 + s1f64
	s0f64 = math.Floor(s0f64)
	l43 = s0f64
	s0f64 = math.Abs(s0f64)
	s1f64 = 2.147483648e+09
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l43
		s0i32 = int32(math.Trunc(s0f64))
		goto lbl22
	}
	s0i32 = -2147483648
lbl22:
	l1 = s0i32
	s1i32 = 1
	s2i32 = l1
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
	l1 = s0i32
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	l9 = s0i32
	s0f64 = 1
	s1i32 = 0
	s2i32 = l4
	s3i32 = l4
	s2i32 = s2i32 * s3i32
	l11 = s2i32
	s3i32 = l14
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l4
	s3i32 = l11
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s1f64 = float64(s1i32)
	s0f64 = s0f64 / s1f64
	s1f64 = 4.294967296e+09
	s0f64 = s0f64 * s1f64
	s0f64 = f0(ctx, s0f64)
	l43 = s0f64
	s1f64 = 1.8446744073709552e+19
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f64 = l43
	s2f64 = 0
	if s1f64 >= s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f64 = l43
		s0i64 = int64(uint64(math.Trunc(s0f64)))
		goto lbl24
	}
	s0i64 = 0
lbl24:
	l41 = s0i64
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l29 = s0i32
	s0i32 = l16
	s1i32 = 1144
	s0i32 = s0i32 + s1i32
	s1i32 = l32
	s2i32 = l9
	if s2i32 != 0 {
		s2i32 = l29
		s3i32 = 2
		s2i32 = i32DivS(s2i32, s3i32)
		s3i32 = 3
		s2i32 = s2i32 * s3i32
		goto lbl26
	}
	s2i32 = l1
	s3i32 = 1
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 3
	s2i32 = s2i32 * s3i32
	s3i32 = -1
	s2i32 = s2i32 + s3i32
lbl26:
	l34 = s2i32
	s3i32 = l2
	f486(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l3
	s1i32 = l16
	s2i32 = 1168
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l16
	s2i32 = 1160
	s1i32 = s1i32 + s2i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l16
	s2i32 = 1152
	s1i32 = s1i32 + s2i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l16
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+1144)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0f64 = 1
	s1i32 = 0
	s2i32 = l1
	s3i32 = l1
	s2i32 = s2i32 * s3i32
	l11 = s2i32
	s3i32 = l9
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l1
	s3i32 = l11
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s1f64 = float64(s1i32)
	s0f64 = s0f64 / s1f64
	s1f64 = 4.294967296e+09
	s0f64 = s0f64 * s1f64
	s0f64 = f0(ctx, s0f64)
	l43 = s0f64
	s1f64 = 1.8446744073709552e+19
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f64 = l43
	s2f64 = 0
	if s1f64 >= s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f64 = l43
		s0i64 = int64(uint64(math.Trunc(s0f64)))
		goto lbl28
	}
	s0i64 = 0
lbl28:
	l42 = s0i64
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = 0
		l32 = s0i32
		s0i32 = 0
		l34 = s0i32
		goto lbl30
	}
	s0i32 = l29
	s1i32 = l1
	s2i32 = l9
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l35 = s0i32
	s1i32 = l29
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l5
	s2i32 = l4
	s3i32 = l14
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l11 = s1i32
	s2i32 = l5
	s3i32 = 1
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	s2i32 = l4
	s3i32 = l1
	if uint32(s2i32) < uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l4 = s0i32
	s1i32 = 1073741824
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l36 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l40 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0i32 = s0i32 - s1i32
		l18 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 - s1i32
		l37 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l14 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l20 = s0i32
		s0i32 = 0
		s1i32 = l16
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1060)]))
		l9 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l1 = s0i32
		s1i32 = l4
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l4 = s1i32
		s0i32 = s0i32 | s1i32
		s1i32 = l16
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1064)]))
		l7 = s1i32
		s2i32 = l9
		s1i32 = s1i32 - s2i32
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l33
			s1i32 = l4
			s2i32 = 4
			f18(ctx, s0i32, s1i32, s2i32)
			s0i32 = l16
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1064)]))
			l7 = s0i32
			s0i32 = 0
			s1i32 = l16
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1060)]))
			l9 = s1i32
			s0i32 = s0i32 - s1i32
			s1i32 = 3
			s0i32 = s0i32 & s1i32
			l1 = s0i32
		}
		s0i32 = l16
		s1i32 = l1
		s2i32 = l9
		s1i32 = s1i32 + s2i32
		l27 = s1i32
		s2i32 = l4
		s1i32 = s1i32 + s2i32
		l31 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1060)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l31
		s0i32 = s0i32 - s1i32
		s1i32 = l18
		s2i32 = l37
		s1i32 = s1i32 * s2i32
		l28 = s1i32
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l33
			s1i32 = l28
			s2i32 = 1
			f18(ctx, s0i32, s1i32, s2i32)
			s0i32 = l16
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1060)]))
			l31 = s0i32
		}
		s0i32 = l16
		s1i32 = l28
		s2i32 = l31
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1060)])) = uint32(s1i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l4 = s0i32
		s1i32 = 4
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl39
		}
		s0i32 = l32
		s1i32 = 1
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 1
		s0i32 = s0i32 | s1i32
		l9 = s0i32
		s1i32 = l14
		s2i32 = l20
		s1i32 = s1i32 - s2i32
		l1 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 0
		s2i32 = l9
		s3i32 = l1
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
		l30 = s0i32
		s0i32 = l27
		s1i32 = l5
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l5 = s1i32
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s1i32 = l5
		s0i32 = s0i32 + s1i32
		l14 = s0i32
		s1i32 = l11
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l21 = s0i32
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl33
		case 1:
			goto lbl39
		case 2:
			goto lbl34
		case 3:
			goto lbl35
		default:
			goto lbl40
		}
	lbl40:
		s0i32 = l1
		s1i32 = 8
		s0i32 = i32DivS(s0i32, s1i32)
		l4 = s0i32
		s0i32 = l18
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl32
		}
		s0i32 = l4
		s1i32 = 3
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l1
		s0i32 = s0i32 - s1i32
		s1i32 = 7
		s0i32 = s0i32 + s1i32
		l20 = s0i32
		s0i32 = 0
		s1i32 = l18
		s0i32 = s0i32 - s1i32
		l22 = s0i32
		s0i32 = l21
		s1i32 = l27
		s0i32 = s0i32 - s1i32
		l23 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l15 = s0i32
		s1i32 = l4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
	lbl41:
		s0i32 = 7
		l13 = s0i32
		s0i32 = l15
		l12 = s0i32
		s0i32 = l14
		l4 = s0i32
		s0i32 = 0
		l6 = s0i32
		s0i32 = 0
		l10 = s0i32
		s0i32 = 0
		l11 = s0i32
		s0i32 = l9
		l5 = s0i32
		s0i32 = l27
		s1i32 = 0
		s2i32 = l23
		s0i32 = f27(ctx, s0i32, s1i32, s2i32)
		l24 = s0i32
		l7 = s0i32
		s0i32 = l26
		s1i32 = l31
		s0i32 = s0i32 + s1i32
		l17 = s0i32
		l8 = s0i32
	lbl43:
		s0i32 = l1
		s1i32 = l12
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl45
		}
		s0i32 = l13
		s1i32 = l20
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 0
		s2i32 = l1
		s3i32 = l12
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
		if s0i32 != 0 {
			goto lbl45
		}
		s0i32 = l17
		s1i32 = l28
		s0i32 = s0i32 + s1i32
		l13 = s0i32
		s0i32 = 0
		l17 = s0i32
		s0i32 = l30
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl44
		}
		goto lbl42
	lbl45:
		s0i32 = l8
		s1i32 = l11
		s2i32 = 0
		s3i32 = l12
		s3i32 = int32(ctx.Mem[int(s3i32+0)])
		s4i32 = l13
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 1
		s3i32 = s3i32 & s4i32
		s2i32 = s2i32 - s3i32
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		l19 = s2i32
		s3i32 = l6
		s2i32 = s2i32 + s3i32
		l6 = s2i32
		s3i32 = l10
		s2i32 = s2i32 + s3i32
		l10 = s2i32
		s1i32 = s1i32 + s2i32
		l11 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l25 = s0i32
		s0i32 = l4
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l38 = s0i32
		s0i32 = l5
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l39 = s0i32
		s0i32 = l7
		s1i32 = l19
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l13
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		s1i32 = 7
		s2i32 = l13
		s3i32 = 0
		if s2i32 > s3i32 {
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
		l13 = s0i32
		s0i32 = l12
		s1i32 = l12
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		s2i32 = l19
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l12 = s0i32
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l24
		s2i32 = l7
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l8
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l11
		s1i32 = l25
		s0i32 = s0i32 - s1i32
		l11 = s0i32
		s0i32 = l10
		s1i32 = l38
		s0i32 = s0i32 - s1i32
		l10 = s0i32
		s0i32 = l6
		s1i32 = l39
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		goto lbl43
	lbl44:
	lbl46:
		s0i32 = l8
		s1i32 = l6
		s2i32 = l10
		s1i32 = s1i32 + s2i32
		l12 = s1i32
		s2i32 = l11
		s1i32 = s1i32 + s2i32
		l11 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l10 = s0i32
		s0i32 = l4
		s1i32 = l12
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l5
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l25 = s0i32
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l24
		s2i32 = l7
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l8
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l11
		s1i32 = l10
		s0i32 = s0i32 - s1i32
		l11 = s0i32
		s0i32 = l12
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l10 = s0i32
		s0i32 = l6
		s1i32 = l25
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l17
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l17 = s0i32
		s1i32 = l30
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl46
		}
	lbl42:
		s0i32 = 0
		l17 = s0i32
		s0i32 = l24
		s1i32 = 0
		s2i32 = l23
		s0i32 = f27(ctx, s0i32, s1i32, s2i32)
		l24 = s0i32
		s0i32 = 0
		l11 = s0i32
		s0i32 = 0
		l10 = s0i32
		s0i32 = l1
		l12 = s0i32
		s0i32 = l20
		l6 = s0i32
		s0i32 = l8
		s1i32 = l13
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl48:
			s0i32 = l13
			s1i32 = l22
			s0i32 = s0i32 + s1i32
			l13 = s0i32
			s1i32 = l17
			s2i32 = l11
			s3i32 = l10
			s4i32 = 0
			s5i32 = l12
			s6i32 = l12
			s7i32 = -1
			s6i32 = s6i32 + s7i32
			s7i32 = l6
			s8i32 = 7
			if s7i32 < s8i32 {
				s7i32 = 1
			} else {
				s7i32 = 0
			}
			l10 = s7i32
			if s7i32 != 0 {
				// s5i32 = s5i32
			} else {
				s5i32 = s6i32
			}
			l12 = s5i32
			s5i32 = int32(ctx.Mem[int(s5i32+0)])
			s6i32 = l6
			s7i32 = 1
			s6i32 = s6i32 + s7i32
			s7i32 = 0
			s8i32 = l10
			if s8i32 != 0 {
				// s6i32 = s6i32
			} else {
				s6i32 = s7i32
			}
			l6 = s6i32
			s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
			s6i32 = 1
			s5i32 = s5i32 & s6i32
			s4i32 = s4i32 - s5i32
			s5i32 = 255
			s4i32 = s4i32 & s5i32
			l17 = s4i32
			s3i32 = s3i32 + s4i32
			l10 = s3i32
			s2i32 = s2i32 + s3i32
			l11 = s2i32
			s1i32 = s1i32 + s2i32
			l19 = s1i32
			s1i64 = int64(uint32(s1i32))
			s2i64 = l41
			s1i64 = s1i64 * s2i64
			s2i64 = 2147483648
			s1i64 = s1i64 + s2i64
			s2i64 = 32
			s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
			ctx.Mem[int(s0i32+0)] = uint8(s1i64)
			s0i32 = l4
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l25 = s0i32
			s0i32 = l4
			s1i32 = l11
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l38 = s0i32
			s0i32 = l5
			s1i32 = l10
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l39 = s0i32
			s0i32 = l7
			s1i32 = l17
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l19
			s1i32 = l25
			s0i32 = s0i32 - s1i32
			l17 = s0i32
			s0i32 = l11
			s1i32 = l38
			s0i32 = s0i32 - s1i32
			l11 = s0i32
			s0i32 = l10
			s1i32 = l39
			s0i32 = s0i32 - s1i32
			l10 = s0i32
			s0i32 = l7
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l7 = s0i32
			s1i32 = l24
			s2i32 = l7
			s3i32 = l9
			if uint32(s2i32) < uint32(s3i32) {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l7 = s0i32
			s0i32 = l5
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s1i32 = l9
			s2i32 = l5
			s3i32 = l14
			if uint32(s2i32) < uint32(s3i32) {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l5 = s0i32
			s0i32 = l4
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l4 = s0i32
			s1i32 = l14
			s2i32 = l4
			s3i32 = l21
			if uint32(s2i32) < uint32(s3i32) {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l4 = s0i32
			s0i32 = l13
			s1i32 = l8
			if uint32(s0i32) > uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl48
			}
		}
		s0i32 = l1
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l4 = s1i32
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l4
		s1i32 = l15
		s0i32 = s0i32 + s1i32
		l15 = s0i32
		s0i32 = l26
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l26 = s0i32
		s1i32 = l18
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl41
		}
		goto lbl32
	lbl39:
		s0i32 = l16
		s1i32 = 1036
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l16
		s1i32 = 3924
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l16
		s1i32 = 3874
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = 3848
		s1i32 = l16
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		f19(ctx, s0i32, s1i32)
		panic("unreachable executed")
	}
	f63(ctx)
	panic("unreachable executed")
lbl35:
	s0i32 = l18
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl32
	}
	s0i32 = 0
	s1i32 = l18
	s0i32 = s0i32 - s1i32
	l23 = s0i32
	s0i32 = l21
	s1i32 = l27
	s0i32 = s0i32 - s1i32
	l24 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l20 = s0i32
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
lbl49:
	s0i32 = l14
	l4 = s0i32
	s0i32 = 0
	l13 = s0i32
	s0i32 = 0
	l8 = s0i32
	s0i32 = 0
	l6 = s0i32
	s0i32 = l9
	l5 = s0i32
	s0i32 = l27
	s1i32 = 0
	s2i32 = l24
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l15 = s0i32
	l7 = s0i32
	s0i32 = l26
	s1i32 = l31
	s0i32 = s0i32 + s1i32
	l22 = s0i32
	l12 = s0i32
	s0i32 = l20
	l11 = s0i32
	s1i32 = l1
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl51:
		s0i32 = l12
		s1i32 = l6
		s2i32 = l11
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
		l10 = s2i32
		s3i32 = 31
		s2i32 = s2i32 & s3i32
		l6 = s2i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = l6
		s4i32 = 2
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s2i32 = s2i32 | s3i32
		s3i32 = l10
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 248
		s3i32 = s3i32 & s4i32
		s4i32 = l10
		s5i32 = 13
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s3i32 = s3i32 | s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = l10
		s4i32 = 5
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 63
		s3i32 = s3i32 & s4i32
		l10 = s3i32
		s4i32 = 2
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s4i32 = l10
		s5i32 = 4
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s3i32 = s3i32 | s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = 3
		s2i32 = i32DivU(s2i32, s3i32)
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		l6 = s2i32
		s3i32 = l13
		s2i32 = s2i32 + s3i32
		l10 = s2i32
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		l17 = s2i32
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l13 = s0i32
		s0i32 = l4
		s1i32 = l17
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l5
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l25 = s0i32
		s0i32 = l7
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l13
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l17
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l10
		s1i32 = l25
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s1i32 = l15
		s2i32 = l10
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l12
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = l11
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s1i32 = l1
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl51
		}
	}
	s0i32 = l22
	s1i32 = l28
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = 0
	l11 = s0i32
	s0i32 = l30
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl53:
		s0i32 = l12
		s1i32 = l8
		s2i32 = l13
		s1i32 = s1i32 + s2i32
		l17 = s1i32
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
		s0i32 = l4
		s1i32 = l17
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l22 = s0i32
		s0i32 = l5
		s1i32 = l13
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l15
		s2i32 = l7
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l12
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = l6
		s1i32 = l8
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l17
		s1i32 = l22
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l13
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l11
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s1i32 = l30
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl53
		}
	}
	s0i32 = 0
	l6 = s0i32
	s0i32 = l15
	s1i32 = 0
	s2i32 = l24
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l13 = s0i32
	s0i32 = 0
	l8 = s0i32
	s0i32 = 0
	l11 = s0i32
	s0i32 = l1
	l17 = s0i32
	s0i32 = l12
	s1i32 = l10
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl55:
		s0i32 = l10
		s1i32 = l23
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s1i32 = l6
		s2i32 = l11
		s3i32 = l17
		s4i32 = -2
		s3i32 = s3i32 + s4i32
		l17 = s3i32
		s3i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)])))
		l6 = s3i32
		s4i32 = 31
		s3i32 = s3i32 & s4i32
		l15 = s3i32
		s4i32 = 3
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s4i32 = l15
		s5i32 = 2
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s3i32 = s3i32 | s4i32
		s4i32 = l6
		s5i32 = 8
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s5i32 = 248
		s4i32 = s4i32 & s5i32
		s5i32 = l6
		s6i32 = 13
		s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
		s4i32 = s4i32 | s5i32
		s3i32 = s3i32 + s4i32
		s4i32 = l6
		s5i32 = 5
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s5i32 = 63
		s4i32 = s4i32 & s5i32
		l6 = s4i32
		s5i32 = 2
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s5i32 = l6
		s6i32 = 4
		s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
		s4i32 = s4i32 | s5i32
		s3i32 = s3i32 + s4i32
		s4i32 = 3
		s3i32 = i32DivU(s3i32, s4i32)
		s4i32 = 255
		s3i32 = s3i32 & s4i32
		l6 = s3i32
		s2i32 = s2i32 + s3i32
		l11 = s2i32
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		l15 = s2i32
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l22 = s0i32
		s0i32 = l4
		s1i32 = l15
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l5
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l25 = s0i32
		s0i32 = l7
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l22
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l15
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l11
		s1i32 = l25
		s0i32 = s0i32 - s1i32
		l11 = s0i32
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l15 = s0i32
		s1i32 = l13
		s2i32 = l15
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l10
		s1i32 = l12
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl55
		}
	}
	s0i32 = l1
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l4 = s1i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l4
	s1i32 = l20
	s0i32 = s0i32 + s1i32
	l20 = s0i32
	s0i32 = l26
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l26 = s0i32
	s1i32 = l18
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl49
	}
	goto lbl32
lbl34:
	s0i32 = l18
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl32
	}
	s0i32 = 0
	s1i32 = l18
	s0i32 = s0i32 - s1i32
	l24 = s0i32
	s0i32 = l21
	s1i32 = l27
	s0i32 = s0i32 - s1i32
	l17 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l20 = s0i32
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
lbl56:
	s0i32 = l14
	l4 = s0i32
	s0i32 = 0
	l13 = s0i32
	s0i32 = 0
	l8 = s0i32
	s0i32 = 0
	l6 = s0i32
	s0i32 = l9
	l5 = s0i32
	s0i32 = l27
	s1i32 = 0
	s2i32 = l17
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l15 = s0i32
	l7 = s0i32
	s0i32 = l26
	s1i32 = l31
	s0i32 = s0i32 + s1i32
	l23 = s0i32
	l12 = s0i32
	s0i32 = l20
	l10 = s0i32
	s1i32 = l1
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl58:
		s0i32 = l12
		s1i32 = l6
		s2i32 = l10
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s3i32 = 24
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		l6 = s2i32
		s3i32 = l13
		s2i32 = s2i32 + s3i32
		l11 = s2i32
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		l8 = s2i32
		s1i32 = s1i32 + s2i32
		l13 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l22 = s0i32
		s0i32 = l4
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l5
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l25 = s0i32
		s0i32 = l7
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l15
		s2i32 = l6
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l12
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = l13
		s1i32 = l22
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l8
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l11
		s1i32 = l25
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l10
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s1i32 = l1
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl58
		}
	}
	s0i32 = l23
	s1i32 = l28
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = 0
	l11 = s0i32
	s0i32 = l30
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl60:
		s0i32 = l12
		s1i32 = l8
		s2i32 = l13
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l23 = s0i32
		s0i32 = l4
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l22 = s0i32
		s0i32 = l5
		s1i32 = l13
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l15
		s2i32 = l7
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l12
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = l6
		s1i32 = l23
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l8
		s1i32 = l22
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l13
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l11
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s1i32 = l30
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl60
		}
	}
	s0i32 = 0
	l6 = s0i32
	s0i32 = l15
	s1i32 = 0
	s2i32 = l17
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l23 = s0i32
	s0i32 = 0
	l8 = s0i32
	s0i32 = 0
	l13 = s0i32
	s0i32 = l1
	l11 = s0i32
	s0i32 = l12
	s1i32 = l10
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl62:
		s0i32 = l10
		s1i32 = l24
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s1i32 = l6
		s2i32 = l11
		s3i32 = -4
		s2i32 = s2i32 + s3i32
		l11 = s2i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s3i32 = 24
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		l6 = s2i32
		s3i32 = l13
		s2i32 = s2i32 + s3i32
		l15 = s2i32
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		l8 = s2i32
		s1i32 = s1i32 + s2i32
		l13 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l22 = s0i32
		s0i32 = l4
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l5
		s1i32 = l15
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l25 = s0i32
		s0i32 = l7
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l23
		s2i32 = l6
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l13
		s1i32 = l22
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l8
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l15
		s1i32 = l25
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l10
		s1i32 = l12
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl62
		}
	}
	s0i32 = l1
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l4 = s1i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l4
	s1i32 = l20
	s0i32 = s0i32 + s1i32
	l20 = s0i32
	s0i32 = l26
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l26 = s0i32
	s1i32 = l18
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl56
	}
	goto lbl32
lbl33:
	s0i32 = l18
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl32
	}
	s0i32 = 0
	s1i32 = l18
	s0i32 = s0i32 - s1i32
	l24 = s0i32
	s0i32 = l21
	s1i32 = l27
	s0i32 = s0i32 - s1i32
	l17 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l20 = s0i32
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
lbl63:
	s0i32 = l14
	l4 = s0i32
	s0i32 = 0
	l13 = s0i32
	s0i32 = 0
	l8 = s0i32
	s0i32 = 0
	l6 = s0i32
	s0i32 = l9
	l5 = s0i32
	s0i32 = l27
	s1i32 = 0
	s2i32 = l17
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l15 = s0i32
	l7 = s0i32
	s0i32 = l26
	s1i32 = l31
	s0i32 = s0i32 + s1i32
	l23 = s0i32
	l12 = s0i32
	s0i32 = l20
	l10 = s0i32
	s1i32 = l1
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl65:
		s0i32 = l12
		s1i32 = l6
		s2i32 = l13
		s3i32 = l10
		s3i32 = int32(ctx.Mem[int(s3i32+0)])
		l6 = s3i32
		s2i32 = s2i32 + s3i32
		l11 = s2i32
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		l8 = s2i32
		s1i32 = s1i32 + s2i32
		l13 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l22 = s0i32
		s0i32 = l4
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l5
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l25 = s0i32
		s0i32 = l7
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l15
		s2i32 = l6
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l12
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = l13
		s1i32 = l22
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l8
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l11
		s1i32 = l25
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l10
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s1i32 = l1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl65
		}
	}
	s0i32 = l23
	s1i32 = l28
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = 0
	l11 = s0i32
	s0i32 = l30
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl67:
		s0i32 = l12
		s1i32 = l8
		s2i32 = l13
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l23 = s0i32
		s0i32 = l4
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l22 = s0i32
		s0i32 = l5
		s1i32 = l13
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l15
		s2i32 = l7
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l12
		s1i32 = l18
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = l6
		s1i32 = l23
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l8
		s1i32 = l22
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l13
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l11
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s1i32 = l30
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl67
		}
	}
	s0i32 = 0
	l6 = s0i32
	s0i32 = l15
	s1i32 = 0
	s2i32 = l17
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l23 = s0i32
	s0i32 = 0
	l8 = s0i32
	s0i32 = 0
	l13 = s0i32
	s0i32 = l1
	l11 = s0i32
	s0i32 = l12
	s1i32 = l10
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl69:
		s0i32 = l10
		s1i32 = l24
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s1i32 = l6
		s2i32 = l13
		s3i32 = l11
		s4i32 = -1
		s3i32 = s3i32 + s4i32
		l11 = s3i32
		s3i32 = int32(ctx.Mem[int(s3i32+0)])
		l6 = s3i32
		s2i32 = s2i32 + s3i32
		l15 = s2i32
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		l8 = s2i32
		s1i32 = s1i32 + s2i32
		l13 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l41
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l22 = s0i32
		s0i32 = l4
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l19 = s0i32
		s0i32 = l5
		s1i32 = l15
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l25 = s0i32
		s0i32 = l7
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l23
		s2i32 = l6
		s3i32 = l9
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l9
		s2i32 = l5
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l14
		s2i32 = l4
		s3i32 = l21
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l13
		s1i32 = l22
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s0i32 = l8
		s1i32 = l19
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l15
		s1i32 = l25
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l10
		s1i32 = l12
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl69
		}
	}
	s0i32 = l1
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l4 = s1i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l4
	s1i32 = l20
	s0i32 = s0i32 + s1i32
	l20 = s0i32
	s0i32 = l26
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l26 = s0i32
	s1i32 = l18
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl63
	}
lbl32:
	s0i32 = l37
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l40
	s1i32 = l36
	s0i32 = s0i32 - s1i32
	l30 = s0i32
	s0i32 = l34
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 1
	s0i32 = s0i32 | s1i32
	l1 = s0i32
	s1i32 = l18
	s0i32 = s0i32 - s1i32
	s1i32 = 0
	s2i32 = l1
	s3i32 = l18
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
	l13 = s0i32
	s0i32 = l27
	s1i32 = l29
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l1 = s1i32
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	l14 = s0i32
	s1i32 = l35
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s1i32 = l27
	s0i32 = s0i32 - s1i32
	l21 = s0i32
	s0i32 = 0
	l20 = s0i32
lbl70:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s0i32 = 0
	l26 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s1i32 = l30
	s0i32 = s0i32 * s1i32
	l28 = s0i32
	s0i32 = l31
	s1i32 = l18
	s2i32 = l20
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l18
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l14
	l9 = s0i32
	s0i32 = 0
	l7 = s0i32
	s0i32 = 0
	l8 = s0i32
	s0i32 = 0
	l15 = s0i32
	s0i32 = l2
	l1 = s0i32
	s0i32 = l27
	s1i32 = 0
	s2i32 = l21
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l17 = s0i32
	l4 = s0i32
	s0i32 = l5
	s1i32 = l20
	s0i32 = s0i32 + s1i32
	l29 = s0i32
	l5 = s0i32
	s0i32 = l18
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl72:
		s0i32 = l5
		s1i32 = l15
		s2i32 = l7
		s3i32 = l6
		s3i32 = int32(ctx.Mem[int(s3i32+0)])
		l15 = s3i32
		s2i32 = s2i32 + s3i32
		l7 = s2i32
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		l8 = s2i32
		s1i32 = s1i32 + s2i32
		l24 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l42
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l23 = s0i32
		s0i32 = l9
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l35 = s0i32
		s0i32 = l1
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l36 = s0i32
		s0i32 = l4
		s1i32 = l15
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l17
		s2i32 = l4
		s3i32 = l2
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l2
		s2i32 = l1
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
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
		s0i32 = l9
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s1i32 = l14
		s2i32 = l9
		s3i32 = l12
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l9 = s0i32
		s0i32 = l5
		s1i32 = l10
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l24
		s1i32 = l23
		s0i32 = s0i32 - s1i32
		l15 = s0i32
		s0i32 = l8
		s1i32 = l35
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l7
		s1i32 = l36
		s0i32 = s0i32 - s1i32
		l7 = s0i32
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
			goto lbl72
		}
	}
	s0i32 = l28
	s1i32 = l29
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l13
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl74:
		s0i32 = l5
		s1i32 = l7
		s2i32 = l8
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s2i32 = l15
		s1i32 = s1i32 + s2i32
		l15 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l42
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l28 = s0i32
		s0i32 = l9
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l29 = s0i32
		s0i32 = l1
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l24 = s0i32
		s0i32 = l4
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l17
		s2i32 = l4
		s3i32 = l2
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l2
		s2i32 = l1
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
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
		s0i32 = l9
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s1i32 = l14
		s2i32 = l9
		s3i32 = l12
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l9 = s0i32
		s0i32 = l5
		s1i32 = l10
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l15
		s1i32 = l28
		s0i32 = s0i32 - s1i32
		l15 = s0i32
		s0i32 = l8
		s1i32 = l29
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l7
		s1i32 = l24
		s0i32 = s0i32 - s1i32
		l7 = s0i32
		s0i32 = l26
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l26 = s0i32
		s1i32 = l13
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl74
		}
	}
	s0i32 = l17
	s1i32 = 0
	s2i32 = l21
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l26 = s0i32
	s0i32 = l5
	s1i32 = l6
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l15 = s0i32
		s0i32 = 0
		s1i32 = l10
		s0i32 = s0i32 - s1i32
		l28 = s0i32
		s0i32 = 0
		l8 = s0i32
		s0i32 = 0
		l7 = s0i32
	lbl76:
		s0i32 = l6
		s1i32 = l28
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l15
		s2i32 = l7
		s3i32 = l11
		s4i32 = -1
		s3i32 = s3i32 + s4i32
		l11 = s3i32
		s3i32 = int32(ctx.Mem[int(s3i32+0)])
		l15 = s3i32
		s2i32 = s2i32 + s3i32
		l10 = s2i32
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		l17 = s2i32
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		s1i64 = int64(uint32(s1i32))
		s2i64 = l42
		s1i64 = s1i64 * s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
		s0i32 = l9
		s1i32 = l17
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l29 = s0i32
		s0i32 = l1
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l24 = s0i32
		s0i32 = l4
		s1i32 = l15
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l26
		s2i32 = l4
		s3i32 = l2
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l2
		s2i32 = l1
		s3i32 = l14
		if uint32(s2i32) < uint32(s3i32) {
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
		s0i32 = l9
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s1i32 = l14
		s2i32 = l9
		s3i32 = l12
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l9 = s0i32
		s0i32 = l7
		s1i32 = l8
		s0i32 = s0i32 - s1i32
		l15 = s0i32
		s0i32 = l17
		s1i32 = l29
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l10
		s1i32 = l24
		s0i32 = s0i32 - s1i32
		l7 = s0i32
		s0i32 = l6
		s1i32 = l5
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl76
		}
	}
	s0i32 = l20
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l20 = s0i32
	s1i32 = l37
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl70
	}
lbl30:
	s0i32 = l0
	s1i32 = l34
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l33
	f43(ctx, s0i32)
lbl0:
	s0i32 = l16
	s1i32 = 1200
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
