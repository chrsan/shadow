package internal

import (
	"math"
	"unsafe"
)

func f84(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float64
	_ = l8
	var l9 float64
	_ = l9
	var l10 float64
	_ = l10
	var l11 float64
	_ = l11
	var l12 float64
	_ = l12
	var l13 float64
	_ = l13
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
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l3 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = f24(ctx, s1i32)
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l3
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s1i32 = 3
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = 2
			s0i32 = s0i32 & s1i32
			if s0i32 != 0 {
				s0i32 = 0
				l3 = s0i32
				s0i32 = l0
				s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
				l6 = s0f32
				s1f32 = 0
				if s0f32 == s1f32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl1
				}
				s0i32 = l0
				s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
				l7 = s0f32
				s1f32 = 0
				if s0f32 == s1f32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl1
				}
				s0i32 = l1
				s1i32 = 0
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
				s0i32 = l1
				s1i64 = 4575657221408423936
				*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
				s0i32 = l1
				s1i32 = 0
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
				s0i32 = l1
				s1i32 = 0
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
				s0i32 = l1
				s1f32 = 1
				s2f32 = l7
				s1f32 = s1f32 / s2f32
				l7 = s1f32
				*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
				s0i32 = l1
				s1f32 = 1
				s2f32 = l6
				s1f32 = s1f32 / s2f32
				l6 = s1f32
				*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
				s0i32 = l1
				s1f32 = l6
				s2i32 = l0
				s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
				s2f32 = -s2f32
				s1f32 = s1f32 * s2f32
				*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
				s0i32 = l0
				s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
				l6 = s0f32
				s0i32 = l1
				s1i32 = l4
				s2i32 = 16
				s1i32 = s1i32 | s2i32
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
				s0i32 = l1
				s1f32 = l7
				s2f32 = l6
				s2f32 = -s2f32
				s1f32 = s1f32 * s2f32
				*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
				s0i32 = 1
				l3 = s0i32
				goto lbl1
			}
			s0i32 = l0
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l6 = s0f32
			s0i32 = l0
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
			l7 = s0f32
			s0i32 = l1
			s1i32 = 1065353216
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
			s0i32 = l1
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
			s0i32 = l1
			s1i64 = 4575657221408423936
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
			s0i32 = l1
			s1i64 = 1065353216
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l1
			s1f32 = l7
			s1f32 = -s1f32
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
			s0i32 = l1
			s1f32 = l6
			s1f32 = -s1f32
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
			s0i32 = l1
			s1i32 = 17
			s2i32 = 17
			s3i32 = 16
			s4f32 = l7
			s5f32 = 0
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
			s3f32 = l6
			s4f32 = 0
			if s3f32 != s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
			s0i32 = 1
			l3 = s0i32
			goto lbl1
		}
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1f32 = 0
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 1
			l3 = s0i32
			s0i32 = l0
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			s1f32 = 0
			if s0f32 != s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl1
			}
		}
		s0i32 = 0
		l3 = s0i32
		goto lbl1
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s0f64 = float64(s0f32)
	l8 = s0f64
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0f64 = float64(s0f32)
	l9 = s0f64
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		s0f64 = float64(s0f32)
		l10 = s0f64
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1f64 = float64(s1f32)
		l11 = s1f64
		s0f64 = s0f64 * s1f64
		s1f64 = l8
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
		s2f64 = float64(s2f32)
		l12 = s2f64
		s1f64 = s1f64 * s2f64
		s0f64 = s0f64 - s1f64
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s1f64 = float64(s1f32)
		s0f64 = s0f64 * s1f64
		s1f64 = l8
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
		s2f64 = float64(s2f32)
		l8 = s2f64
		s1f64 = s1f64 * s2f64
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s2f64 = float64(s2f32)
		l13 = s2f64
		s3f64 = l10
		s2f64 = s2f64 * s3f64
		s1f64 = s1f64 - s2f64
		s2f64 = l9
		s1f64 = s1f64 * s2f64
		s2f64 = l13
		s3f64 = l12
		s2f64 = s2f64 * s3f64
		s3f64 = l8
		s4f64 = l11
		s3f64 = s3f64 * s4f64
		s2f64 = s2f64 - s3f64
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
		s3f64 = float64(s3f32)
		s2f64 = s2f64 * s3f64
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 + s1f64
		goto lbl6
	}
	s0f64 = l9
	s1f64 = l8
	s0f64 = s0f64 * s1f64
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1f64 = float64(s1f32)
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 - s1f64
lbl6:
	l8 = s0f64
	s0i32 = 0
	l3 = s0i32
	s0f64 = 0
	s1f64 = 1
	s2f64 = l8
	s1f64 = s1f64 / s2f64
	s2f64 = l8
	s2f32 = float32(s2f64)
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 1.4551915e-11
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	l8 = s0f64
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l2
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
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
	l4 = s0i32
	s1i32 = l0
	s2f64 = l8
	s3i32 = l5
	s4i32 = 0
	if s3i32 != s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	f976(ctx, s0i32, s1i32, s2f64, s3i32)
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+36)])
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = 1
	l3 = s0i32
	s0i32 = l0
	s1i32 = l1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
lbl1:
	s0i32 = l2
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l3
	return s0i32
}
