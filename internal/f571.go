package internal

import (
	"math"
	"unsafe"
)

func f571(ctx *Context, l0 float64, l1 float64, l2 float64, l3 float64, l4 int32) int32 {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
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
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l7
	s1f64 = l0
	s2f64 = l1
	s3f64 = l2
	s4f64 = l3
	s5i32 = l7
	s1i32 = f1421(ctx, s1f64, s2f64, s3f64, s4f64, s5i32)
	l9 = s1i32
	s2i32 = l4
	s0i32 = f1419(ctx, s0i32, s1i32, s2i32)
	l5 = s0i32
	s0i32 = l9
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl1:
		s0i32 = l7
		s1i32 = l8
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0f64
		s1f64 = 1.0000001192092896
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0f64 = 1
		s1f64 = l0
		s0f64 = s0f64 - s1f64
		s1f64 = 1.00005
		s2f64 = l0
		s1f64 = s1f64 - s2f64
		s0f64 = s0f64 * s1f64
		s1f64 = 0
		if s0f64 <= s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = 0
		l6 = s0i32
		s0i32 = l5
		s1i32 = 0
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	lbl5:
		s0i32 = l4
		s1i32 = l6
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1f64 = -1
		s0f64 = s0f64 + s1f64
		s0f64 = math.Abs(s0f64)
		s1f64 = 1.1920928955078125e-07
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l5
		s1i32 = l6
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		goto lbl3
	lbl4:
		s0f64 = l0
		s1f64 = -1.1920928955078125e-07
		if s0f64 > s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0f64 = -5e-05
		s1f64 = l0
		s0f64 = s0f64 - s1f64
		s1f64 = 0
		s2f64 = l0
		s1f64 = s1f64 - s2f64
		s0f64 = s0f64 * s1f64
		s1f64 = 0
		if s0f64 <= s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = 0
		l6 = s0i32
		s0i32 = l5
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl7:
			s0i32 = l4
			s1i32 = l6
			s2i32 = 3
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s0f64 = math.Abs(s0f64)
			s1f64 = 1.1920928955078125e-07
			if s0f64 < s1f64 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl2
			}
			s0i32 = l6
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l6 = s0i32
			s1i32 = l5
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl7
			}
		}
		s0i32 = l4
		s1i32 = l5
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		goto lbl2
	lbl3:
		s0i32 = l4
		s1i32 = l5
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i64 = 4607182418800017408
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
	lbl2:
		s0i32 = l8
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s1i32 = l9
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l7
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l5
	return s0i32
}
