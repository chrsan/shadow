package internal

import (
	"math"
	"unsafe"
)

func f1361(ctx *Context, l0 float64, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var l7 float64
	_ = l7
	var l8 float64
	_ = l8
	var l9 float64
	_ = l9
	var l10 float64
	_ = l10
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
	var s4f64 float64
	_ = s4f64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0f64 = l0
	s0i64 = int64(math.Float64bits(s0f64))
	l6 = s0i64
	s1i64 = 32
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	s0i32 = int32(s0i64)
	l2 = s0i32
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l3 = s0i32
	s1i32 = 1074752122
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 1048575
		s0i32 = s0i32 & s1i32
		s1i32 = 598523
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l3
		s1i32 = 1073928572
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i64 = l6
			s1i64 = 0
			if s0i64 >= s1i64 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l1
				s1f64 = l0
				s2f64 = -1.5707963267341256
				s1f64 = s1f64 + s2f64
				l0 = s1f64
				s2f64 = -6.077100506506192e-11
				s1f64 = s1f64 + s2f64
				l7 = s1f64
				*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
				s0i32 = l1
				s1f64 = l0
				s2f64 = l7
				s1f64 = s1f64 - s2f64
				s2f64 = -6.077100506506192e-11
				s1f64 = s1f64 + s2f64
				*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
				s0i32 = 1
				l2 = s0i32
				goto lbl0
			}
			s0i32 = l1
			s1f64 = l0
			s2f64 = 1.5707963267341256
			s1f64 = s1f64 + s2f64
			l0 = s1f64
			s2f64 = 6.077100506506192e-11
			s1f64 = s1f64 + s2f64
			l7 = s1f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
			s0i32 = l1
			s1f64 = l0
			s2f64 = l7
			s1f64 = s1f64 - s2f64
			s2f64 = 6.077100506506192e-11
			s1f64 = s1f64 + s2f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
			s0i32 = -1
			l2 = s0i32
			goto lbl0
		}
		s0i64 = l6
		s1i64 = 0
		if s0i64 >= s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1f64 = l0
			s2f64 = -3.1415926534682512
			s1f64 = s1f64 + s2f64
			l0 = s1f64
			s2f64 = -1.2154201013012384e-10
			s1f64 = s1f64 + s2f64
			l7 = s1f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
			s0i32 = l1
			s1f64 = l0
			s2f64 = l7
			s1f64 = s1f64 - s2f64
			s2f64 = -1.2154201013012384e-10
			s1f64 = s1f64 + s2f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
			s0i32 = 2
			l2 = s0i32
			goto lbl0
		}
		s0i32 = l1
		s1f64 = l0
		s2f64 = 3.1415926534682512
		s1f64 = s1f64 + s2f64
		l0 = s1f64
		s2f64 = 1.2154201013012384e-10
		s1f64 = s1f64 + s2f64
		l7 = s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l1
		s1f64 = l0
		s2f64 = l7
		s1f64 = s1f64 - s2f64
		s2f64 = 1.2154201013012384e-10
		s1f64 = s1f64 + s2f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
		s0i32 = -2
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 1075594811
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 1075183036
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = 1074977148
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl2
			}
			s0i64 = l6
			s1i64 = 0
			if s0i64 >= s1i64 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l1
				s1f64 = l0
				s2f64 = -4.712388980202377
				s1f64 = s1f64 + s2f64
				l0 = s1f64
				s2f64 = -1.8231301519518578e-10
				s1f64 = s1f64 + s2f64
				l7 = s1f64
				*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
				s0i32 = l1
				s1f64 = l0
				s2f64 = l7
				s1f64 = s1f64 - s2f64
				s2f64 = -1.8231301519518578e-10
				s1f64 = s1f64 + s2f64
				*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
				s0i32 = 3
				l2 = s0i32
				goto lbl0
			}
			s0i32 = l1
			s1f64 = l0
			s2f64 = 4.712388980202377
			s1f64 = s1f64 + s2f64
			l0 = s1f64
			s2f64 = 1.8231301519518578e-10
			s1f64 = s1f64 + s2f64
			l7 = s1f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
			s0i32 = l1
			s1f64 = l0
			s2f64 = l7
			s1f64 = s1f64 - s2f64
			s2f64 = 1.8231301519518578e-10
			s1f64 = s1f64 + s2f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
			s0i32 = -3
			l2 = s0i32
			goto lbl0
		}
		s0i32 = l3
		s1i32 = 1075388923
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i64 = l6
		s1i64 = 0
		if s0i64 >= s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1f64 = l0
			s2f64 = -6.2831853069365025
			s1f64 = s1f64 + s2f64
			l0 = s1f64
			s2f64 = -2.430840202602477e-10
			s1f64 = s1f64 + s2f64
			l7 = s1f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
			s0i32 = l1
			s1f64 = l0
			s2f64 = l7
			s1f64 = s1f64 - s2f64
			s2f64 = -2.430840202602477e-10
			s1f64 = s1f64 + s2f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
			s0i32 = 4
			l2 = s0i32
			goto lbl0
		}
		s0i32 = l1
		s1f64 = l0
		s2f64 = 6.2831853069365025
		s1f64 = s1f64 + s2f64
		l0 = s1f64
		s2f64 = 2.430840202602477e-10
		s1f64 = s1f64 + s2f64
		l7 = s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l1
		s1f64 = l0
		s2f64 = l7
		s1f64 = s1f64 - s2f64
		s2f64 = 2.430840202602477e-10
		s1f64 = s1f64 + s2f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
		s0i32 = -4
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 1094263290
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = l1
	s1f64 = l0
	s2f64 = l0
	s3f64 = 0.6366197723675814
	s2f64 = s2f64 * s3f64
	s3f64 = 6.755399441055744e+15
	s2f64 = s2f64 + s3f64
	s3f64 = -6.755399441055744e+15
	s2f64 = s2f64 + s3f64
	l8 = s2f64
	s3f64 = -1.5707963267341256
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	l7 = s1f64
	s2f64 = l8
	s3f64 = 6.077100506506192e-11
	s2f64 = s2f64 * s3f64
	l10 = s2f64
	s1f64 = s1f64 - s2f64
	l0 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l3
	s1i32 = 20
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l5 = s0i32
	s1f64 = l0
	s1i64 = int64(math.Float64bits(s1f64))
	s2i64 = 52
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	s2i32 = 2047
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 - s1i32
	s1i32 = 17
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l3 = s0i32
	s0f64 = l8
	s0f64 = math.Abs(s0f64)
	s1f64 = 2.147483648e+09
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l8
		s0i32 = int32(math.Trunc(s0f64))
		goto lbl11
	}
	s0i32 = -2147483648
lbl11:
	l2 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l1
	s1f64 = l7
	s2f64 = l8
	s3f64 = 6.077100506303966e-11
	s2f64 = s2f64 * s3f64
	l0 = s2f64
	s1f64 = s1f64 - s2f64
	l9 = s1f64
	s2f64 = l8
	s3f64 = 2.0222662487959506e-21
	s2f64 = s2f64 * s3f64
	s3f64 = l7
	s4f64 = l9
	s3f64 = s3f64 - s4f64
	s4f64 = l0
	s3f64 = s3f64 - s4f64
	s2f64 = s2f64 - s3f64
	l10 = s2f64
	s1f64 = s1f64 - s2f64
	l0 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l5
	s1f64 = l0
	s1i64 = int64(math.Float64bits(s1f64))
	s2i64 = 52
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	s2i32 = 2047
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 - s1i32
	s1i32 = 50
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l9
		l7 = s0f64
		goto lbl13
	}
	s0i32 = l1
	s1f64 = l9
	s2f64 = l8
	s3f64 = 2.0222662487111665e-21
	s2f64 = s2f64 * s3f64
	l0 = s2f64
	s1f64 = s1f64 - s2f64
	l7 = s1f64
	s2f64 = l8
	s3f64 = 8.4784276603689e-32
	s2f64 = s2f64 * s3f64
	s3f64 = l9
	s4f64 = l7
	s3f64 = s3f64 - s4f64
	s4f64 = l0
	s3f64 = s3f64 - s4f64
	s2f64 = s2f64 - s3f64
	l10 = s2f64
	s1f64 = s1f64 - s2f64
	l0 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
lbl13:
	s0i32 = l1
	s1f64 = l7
	s2f64 = l0
	s1f64 = s1f64 - s2f64
	s2f64 = l10
	s1f64 = s1f64 - s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
	goto lbl0
lbl1:
	s0i32 = l3
	s1i32 = 2146435072
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1f64 = l0
		s2f64 = l0
		s1f64 = s1f64 - s2f64
		l0 = s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l1
		s1f64 = l0
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
		s0i32 = 0
		l2 = s0i32
		goto lbl0
	}
	s0i64 = l6
	s1i64 = 4503599627370495
	s0i64 = s0i64 & s1i64
	s1i64 = 4710765210229538816
	s0i64 = s0i64 | s1i64
	s0f64 = math.Float64frombits(uint64(s0i64))
	l0 = s0f64
	s0i32 = 0
	l2 = s0i32
lbl16:
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	l5 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1f64 = l0
	s1f64 = math.Abs(s1f64)
	s2f64 = 2.147483648e+09
	if s1f64 < s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f64 = l0
		s1i32 = int32(math.Trunc(s1f64))
		goto lbl17
	}
	s1i32 = -2147483648
lbl17:
	s1f64 = float64(s1i32)
	l7 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0f64 = l0
	s1f64 = l7
	s0f64 = s0f64 - s1f64
	s1f64 = 1.6777216e+07
	s0f64 = s0f64 * s1f64
	l0 = s0f64
	s0i32 = 1
	l2 = s0i32
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = l4
	s1f64 = l0
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f64
	s0f64 = l0
	s1f64 = 0
	if s0f64 != s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 2
		l2 = s0i32
		goto lbl19
	}
	s0i32 = 1
	l5 = s0i32
lbl21:
	s0i32 = l5
	l2 = s0i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
lbl19:
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = l3
	s3i32 = 20
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = -1046
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s4i32 = 1
	s3i32 = s3i32 + s4i32
	s4i32 = 1
	s0i32 = f538(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l2 = s0i32
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0f64
	s0i64 = l6
	s1i64 = -1
	if s0i64 <= s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1f64 = l0
		s1f64 = -s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l1
		s1i32 = l4
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s1f64 = -s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
		s0i32 = 0
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l1
	s1f64 = l0
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l1
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
lbl0:
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l2
	return s0i32
}
