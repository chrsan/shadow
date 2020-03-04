package internal

import (
	"math"
	"unsafe"
)

func f633(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
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
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = s1f32
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = s1f32
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = s1f32
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = s1f32
		s0i32 = l3
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s0i32 = f135(ctx, s0i32, s1i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l3
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+92)]))
		l7 = s1f32
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l8 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l8
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
		l8 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l8
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
		l8 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l8
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl2
		}
		s1i32 = -2147483648
	lbl2:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		l8 = s1f32
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l9 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l9
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
		l9 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l9
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
		l9 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l9
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl4
		}
		s1i32 = -2147483648
	lbl4:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
		l9 = s1f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l10 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l10
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
		l10 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l10
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
		l10 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l10
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl6
		}
		s1i32 = -2147483648
	lbl6:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		l10 = s1f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l11 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l11
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
		l11 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l11
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
		l11 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l11
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl8
		}
		s1i32 = -2147483648
	lbl8:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l7
			s1f32 = 65536
			s0f32 = s0f32 * s1f32
			l7 = s0f32
			s1f32 = 2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s1f32 = -2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 2.1474836e+09
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0f32 = l7
				s0i32 = int32(math.Trunc(float64(s0f32)))
				goto lbl11
			}
			s0i32 = -2147483648
		lbl11:
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = 8
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l4 = s0i32
			s0f32 = l8
			s1f32 = 65536
			s0f32 = s0f32 * s1f32
			l7 = s0f32
			s1f32 = 2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s1f32 = -2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 2.1474836e+09
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0f32 = l7
				s0i32 = int32(math.Trunc(float64(s0f32)))
				goto lbl13
			}
			s0i32 = -2147483648
		lbl13:
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = 8
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l0 = s0i32
			s0f32 = l9
			s1f32 = 65536
			s0f32 = s0f32 * s1f32
			l7 = s0f32
			s1f32 = 2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s1f32 = -2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 2.1474836e+09
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0f32 = l7
				s0i32 = int32(math.Trunc(float64(s0f32)))
				goto lbl15
			}
			s0i32 = -2147483648
		lbl15:
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = 8
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l5 = s0i32
			s0f32 = l10
			s1f32 = 65536
			s0f32 = s0f32 * s1f32
			l7 = s0f32
			s1f32 = 2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s1f32 = -2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 2.1474836e+09
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0f32 = l7
				s0i32 = int32(math.Trunc(float64(s0f32)))
				goto lbl17
			}
			s0i32 = -2147483648
		lbl17:
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = 8
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			s1i32 = l5
			s2i32 = l0
			s3i32 = l4
			s4i32 = l2
			s5i32 = 1
			f148(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
			goto lbl1
		}
		s0i32 = l3
		s1i32 = l1
		s2i32 = l3
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s0i32 = f111(ctx, s0i32, s1i32, s2i32)
		l4 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+60)])
		if s0i32 != 0 {
			goto lbl1
		}
	lbl19:
		s0i32 = l3
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = s1f32
		s0i32 = l3
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = s1f32
		s0i32 = l3
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = s1f32
		s0i32 = l3
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = s1f32
		s0i32 = l3
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s0i32 = f135(ctx, s0i32, s1i32)
		if s0i32 != 0 {
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
			s1f32 = 65536
			s0f32 = s0f32 * s1f32
			l7 = s0f32
			s1f32 = 2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s1f32 = -2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 2.1474836e+09
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0f32 = l7
				s0i32 = int32(math.Trunc(float64(s0f32)))
				goto lbl21
			}
			s0i32 = -2147483648
		lbl21:
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = 8
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l5 = s0i32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
			s1f32 = 65536
			s0f32 = s0f32 * s1f32
			l7 = s0f32
			s1f32 = 2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s1f32 = -2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 2.1474836e+09
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0f32 = l7
				s0i32 = int32(math.Trunc(float64(s0f32)))
				goto lbl23
			}
			s0i32 = -2147483648
		lbl23:
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = 8
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l1 = s0i32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
			s1f32 = 65536
			s0f32 = s0f32 * s1f32
			l7 = s0f32
			s1f32 = 2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s1f32 = -2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 2.1474836e+09
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0f32 = l7
				s0i32 = int32(math.Trunc(float64(s0f32)))
				goto lbl25
			}
			s0i32 = -2147483648
		lbl25:
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = 8
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l6 = s0i32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
			s1f32 = 65536
			s0f32 = s0f32 * s1f32
			l7 = s0f32
			s1f32 = 2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s1f32 = -2.1474835e+09
			s2f32 = l7
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
			l7 = s0f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 2.1474836e+09
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0f32 = l7
				s0i32 = int32(math.Trunc(float64(s0f32)))
				goto lbl27
			}
			s0i32 = -2147483648
		lbl27:
			s1i32 = 128
			s0i32 = s0i32 + s1i32
			s1i32 = 8
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			s1i32 = l6
			s2i32 = l1
			s3i32 = l5
			s4i32 = l2
			s5i32 = 1
			f148(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		}
		s0i32 = l4
		f110(ctx, s0i32)
		s0i32 = l4
		s0i32 = int32(ctx.Mem[int(s0i32+60)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
	lbl1:
		s0i32 = l3
		s1i32 = 96
		s0i32 = s0i32 + s1i32
		ctx.g0 = s0i32
		return
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1f32 = 65536
	s0f32 = s0f32 * s1f32
	l7 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l7
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
	l7 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l7
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
	l7 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l7
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl29
	}
	s0i32 = -2147483648
lbl29:
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l4 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f32 = 65536
	s0f32 = s0f32 * s1f32
	l7 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l7
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
	l7 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l7
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
	l7 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l7
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl31
	}
	s0i32 = -2147483648
lbl31:
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l5 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = 65536
	s0f32 = s0f32 * s1f32
	l7 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l7
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
	l7 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l7
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
	l7 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l7
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl33
	}
	s0i32 = -2147483648
lbl33:
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l1 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 65536
	s0f32 = s0f32 * s1f32
	l7 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l7
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
	l7 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l7
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
	l7 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l7
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl35
	}
	s0i32 = -2147483648
lbl35:
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l1
	s2i32 = l5
	s3i32 = l4
	s4i32 = l2
	s5i32 = 1
	f148(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	s0i32 = l3
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
