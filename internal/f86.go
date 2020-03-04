package internal

import (
	"math"
	"unsafe"
)

func f86(ctx *Context, l0 float32) float32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 float64
	_ = l4
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
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0f32 = l0
	s0i32 = int32(math.Float32bits(s0f32))
	l3 = s0i32
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	s1i32 = 1061752794
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 1
		s1i32 = l1
		s2i32 = 964689920
		if uint32(s1i32) < uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl0
		}
		s0f32 = l0
		s0f64 = float64(s0f32)
		s0f32 = f126(ctx, s0f64)
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 1081824209
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l0
		s0f64 = float64(s0f32)
		l4 = s0f64
		s0i32 = l1
		s1i32 = 1075235812
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f64 = 3.141592653589793
			s1f64 = -3.141592653589793
			s2i32 = l3
			s3i32 = 0
			if s2i32 < s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0f64 = s0f64
			} else {
				s0f64 = s1f64
			}
			s1f64 = l4
			s0f64 = s0f64 + s1f64
			s0f32 = f126(ctx, s0f64)
			s0f32 = -s0f32
			goto lbl0
		}
		s0i32 = l3
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f64 = l4
			s1f64 = 1.5707963267948966
			s0f64 = s0f64 + s1f64
			s0f32 = f125(ctx, s0f64)
			goto lbl0
		}
		s0f64 = 1.5707963267948966
		s1f64 = l4
		s0f64 = s0f64 - s1f64
		s0f32 = f125(ctx, s0f64)
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 1088565717
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 1085271520
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f64 = 6.283185307179586
			s1f64 = -6.283185307179586
			s2i32 = l3
			s3i32 = 0
			if s2i32 < s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0f64 = s0f64
			} else {
				s0f64 = s1f64
			}
			s1f32 = l0
			s1f64 = float64(s1f32)
			s0f64 = s0f64 + s1f64
			s0f32 = f126(ctx, s0f64)
			goto lbl0
		}
		s0i32 = l3
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f64 = -4.71238898038469
			s1f32 = l0
			s1f64 = float64(s1f32)
			s0f64 = s0f64 - s1f64
			s0f32 = f125(ctx, s0f64)
			goto lbl0
		}
		s0f32 = l0
		s0f64 = float64(s0f32)
		s1f64 = -4.71238898038469
		s0f64 = s0f64 + s1f64
		s0f32 = f125(ctx, s0f64)
		goto lbl0
	}
	s0f32 = l0
	s1f32 = l0
	s0f32 = s0f32 - s1f32
	s1i32 = l1
	s2i32 = 2139095040
	if uint32(s1i32) >= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = l0
	s1i32 = l2
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s0i32 = f371(ctx, s0f32, s1i32)
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	s1i32 = 2
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl10
		case 1:
			goto lbl9
		default:
			goto lbl11
		}
	lbl11:
		s0i32 = l2
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s0f32 = f126(ctx, s0f64)
		goto lbl0
	lbl10:
		s0i32 = l2
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s0f64 = -s0f64
		s0f32 = f125(ctx, s0f64)
		goto lbl0
	lbl9:
		s0i32 = l2
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s0f32 = f126(ctx, s0f64)
		s0f32 = -s0f32
		goto lbl0
	}
	s0i32 = l2
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s0f32 = f125(ctx, s0f64)
lbl0:
	l0 = s0f32
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0f32 = l0
	return s0f32
}
