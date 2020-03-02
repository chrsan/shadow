package internal

import (
	"math"
	"unsafe"
)

func f664(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
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
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = 4
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l5
		s2i32 = l1
		s0i32 = f42(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0f32
		s0i32 = l3
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l1 = s0i32
		s0i32 = l5
		s1i32 = l4
		if s1i32 != 0 {
			s1i32 = l5
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			s1f32 = float32(math.Floor(float64(s1f32)))
			l6 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l6
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
			l6 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l6
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
			l6 = s1f32
			s1f32 = l7
			s1f32 = float32(math.Floor(float64(s1f32)))
			l7 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l7
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
			l7 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l7
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
			l7 = s1f32
			s1i32 = l5
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
			s1f32 = float32(math.Ceil(float64(s1f32)))
			l9 = s1f32
			s1i32 = l5
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
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
			goto lbl2
		}
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l6 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l6
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
		l6 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l6
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
		l6 = s1f32
		s1f32 = l7
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l7 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l7
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
		l7 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l7
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
		l7 = s1f32
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l9 = s1f32
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
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
	lbl2:
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
			goto lbl1
		}
		s1i32 = -2147483648
	lbl1:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 4676
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l5
		s1f32 = l6
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l6
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl5
		}
		s1i32 = -2147483648
	lbl5:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l5
		s1f32 = l7
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l7
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl7
		}
		s1i32 = -2147483648
	lbl7:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l5
		s1f32 = l9
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
		l6 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l6
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
		l6 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l6
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl9
		}
		s1i32 = -2147483648
	lbl9:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		f437(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l5
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
