package internal

import (
	"math"
	"unsafe"
)

func f413(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l16 int64
	_ = l16
	var l17 int64
	_ = l17
	var l18 float32
	_ = l18
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
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 176
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 20568
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 20632
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 20696
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 13860
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 13852
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l4
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
		s0i32 = l4
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
		s0i32 = l4
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
	}
	s0i32 = l1
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l15 = s0i32
	lbl2:
		s0i32 = l0
		s1i32 = l13
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = 48
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = 16
		s2i32 = s2i32 + s3i32
		s0i32 = f335(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l2
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 16
			s0i32 = s0i32 + s1i32
			s1i32 = l4
			s2i32 = 32
			s1i32 = s1i32 + s2i32
			s2i32 = l4
			s3i32 = 16
			s2i32 = s2i32 + s3i32
			s0i32 = f335(ctx, s0i32, s1i32, s2i32)
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl3
			}
		}
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		s1f32 = 64
		s0f32 = s0f32 * s1f32
		l18 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l18
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl5
		}
		s0i32 = -2147483648
	lbl5:
		l5 = s0i32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		s1f32 = 64
		s0f32 = s0f32 * s1f32
		l18 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l18
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl7
		}
		s0i32 = -2147483648
	lbl7:
		l8 = s0i32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		s1f32 = 64
		s0f32 = s0f32 * s1f32
		l18 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l18
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl9
		}
		s0i32 = -2147483648
	lbl9:
		l7 = s0i32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1f32 = 64
		s0f32 = s0f32 * s1f32
		l18 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l18
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl11
		}
		s0i32 = -2147483648
	lbl11:
		l9 = s0i32
		s0i32 = l3
		s1i32 = l2
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl13
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l6 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l10 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l11 = s0i32
		s0i32 = l4
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l11
		s2i32 = 6
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l10
		s2i32 = 6
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l6
		s2i32 = 6
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l1
		s2i32 = 6
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
		s0i32 = l8
		l1 = s0i32
		s1i32 = l9
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l9
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = l8
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
			s0i32 = l9
			l1 = s0i32
		}
		s0i32 = l5
		l6 = s0i32
		s0i32 = l4
		s1i32 = l5
		s2i32 = l7
		if s1i32 < s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l4
			s2i32 = l7
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)])) = uint32(s2i32)
			s1i32 = l4
			s2i32 = l5
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)])) = uint32(s2i32)
			s1i32 = l7
		} else {
			s1i32 = l6
		}
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l1
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 160
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		s2i32 = l4
		s3i32 = 80
		s2i32 = s2i32 + s3i32
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
		l1 = s0i32
		s0i64 = int64(s0i32)
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		l6 = s1i32
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l16 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
		l10 = s0i32
		s0i64 = int64(s0i32)
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)]))
		l11 = s1i32
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l17 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i64 = l16
		s1i64 = l17
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967295
		if uint64(s0i64) > uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l12 = s0i32
		s1i32 = l10
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
		l10 = s0i32
		s1i32 = l1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l1 = s0i32
		s1i32 = l11
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		l11 = s0i32
		s1i32 = l6
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l10
		s0i64 = int64(s0i32)
		s1i32 = l11
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l16 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l12
		s0i64 = int64(s0i32)
		s1i32 = l1
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l17 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l3
		s1i64 = l16
		s2i64 = l17
		s1i64 = s1i64 | s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 4294967296
		if uint64(s1i64) < uint64(s2i64) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl13
		}
	lbl16:
		s0i32 = l4
		s1i32 = 96
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l2
		s3i32 = 0
		s0i32 = f364(ctx, s0i32, s1i32, s2i32, s3i32)
	lbl13:
		l1 = s0i32
		s0i32 = l8
		s1i32 = l9
		s0i32 = s0i32 - s1i32
		l10 = s0i32
		s1i32 = l10
		s2i32 = 31
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l6 = s1i32
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s0i32 = s0i32 ^ s1i32
		s1i32 = l5
		s2i32 = l7
		s1i32 = s1i32 - s2i32
		l11 = s1i32
		s2i32 = l11
		s3i32 = 31
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l6 = s2i32
		s1i32 = s1i32 + s2i32
		s2i32 = l6
		s1i32 = s1i32 ^ s2i32
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l8
			s1i32 = l9
			s2i32 = l8
			s3i32 = l9
			if s2i32 < s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l12 = s2i32
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l14 = s0i32
			s1i32 = 32
			s0i32 = s0i32 + s1i32
			s1i32 = 6
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l6 = s0i32
			s1i32 = l9
			s2i32 = l8
			s3i32 = l12
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			s2i32 = 32
			s1i32 = s1i32 + s2i32
			s2i32 = 6
			s1i32 = s1i32 >> (uint32(s2i32) & 31)
			l8 = s1i32
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl3
			}
			s0i32 = l4
			s1i32 = l11
			s1i64 = int64(s1i32)
			s2i64 = 16
			s1i64 = s1i64 << (uint64(s2i64) & 63)
			s2i32 = l10
			s2i64 = int64(s2i32)
			s1i64 = i64DivS(s1i64, s2i64)
			l16 = s1i64
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
			s0i32 = l4
			s1i64 = -2147483647
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l4
			s1i64 = 2147483647
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l5
			s1i32 = l7
			s2i32 = l12
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			s1i32 = 10
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s1i32 = l4
			s2i32 = 8
			s1i32 = s1i32 + s2i32
			s2i32 = l4
			s3i32 = 160
			s2i32 = s2i32 + s3i32
			s3i32 = l4
			s4i64 = l16
			s5i64 = 2147483647
			if s4i64 < s5i64 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			if s4i32 != 0 {
				// s2i32 = s2i32
			} else {
				s2i32 = s3i32
			}
			s3i64 = l16
			s4i64 = -2147483647
			if s3i64 < s4i64 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			l7 = s1i32
			s2i32 = 32
			s3i32 = l14
			s2i32 = s2i32 - s3i32
			s3i32 = 63
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 * s2i32
			s2i32 = 6
			s1i32 = s1i32 >> (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			l5 = s0i32
		lbl18:
			s0i32 = l1
			s1i32 = l6
			s2i32 = l5
			s3i32 = 16
			s2i32 = s2i32 >> (uint32(s3i32) & 31)
			s3i32 = 1
			s4i32 = l1
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
			if int(s4i32) < 0 || int(s4i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s4i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s4i32].numParams != 4 {
				panic("argument count mismatch")
			}
			(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
			s0i32 = l5
			s1i32 = l7
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l6
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l6 = s0i32
			s1i32 = l8
			if s0i32 < s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl18
			}
			goto lbl3
		}
		s0i32 = l5
		s1i32 = l7
		s2i32 = l5
		s3i32 = l7
		if s2i32 < s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l12 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l14 = s0i32
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		s1i32 = 6
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l6 = s0i32
		s1i32 = l7
		s2i32 = l5
		s3i32 = l12
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = 6
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l7 = s1i32
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l4
		s1i32 = l10
		s1i64 = int64(s1i32)
		s2i64 = 16
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i32 = l11
		s2i64 = int64(s2i32)
		s1i64 = i64DivS(s1i64, s2i64)
		l16 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = -2147483647
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = 2147483647
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l8
		s1i32 = l9
		s2i32 = l12
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = 10
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l4
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = 160
		s2i32 = s2i32 + s3i32
		s3i32 = l4
		s4i64 = l16
		s5i64 = 2147483647
		if s4i64 < s5i64 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s3i64 = l16
		s4i64 = -2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l8 = s1i32
		s2i32 = 32
		s3i32 = l14
		s2i32 = s2i32 - s3i32
		s3i32 = 63
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 * s2i32
		s2i32 = 6
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l5 = s0i32
	lbl19:
		s0i32 = l1
		s1i32 = l5
		s2i32 = 16
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = l6
		s3i32 = 1
		s4i32 = l1
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
		if int(s4i32) < 0 || int(s4i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s4i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s4i32].numParams != 4 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l5
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l7
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
	lbl3:
		s0i32 = l13
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l13 = s0i32
		s1i32 = l15
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l4
	s1i32 = 140
	s0i32 = s0i32 + s1i32
	s0i32 = f35(ctx, s0i32)
	s0i32 = l4
	s1i32 = 108
	s0i32 = s0i32 + s1i32
	s0i32 = f35(ctx, s0i32)
	s0i32 = l4
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	s0i32 = f35(ctx, s0i32)
	s0i32 = l4
	s1i32 = 176
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
