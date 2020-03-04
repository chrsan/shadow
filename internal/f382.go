package internal

import (
	"math"
	"unsafe"
)

func f382(ctx *Context, l0 float64, l1 float64, l2 float64, l3 int32) int32 {
	var l4 int32
	_ = l4
	var l5 float64
	_ = l5
	var l6 float64
	_ = l6
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
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
	s0f64 = l0
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l1
		s0f64 = math.Abs(s0f64)
		s1f64 = 1.1920928955078125e-07
		if s0f64 < s1f64 {
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
		goto lbl0
	}
	s0f64 = l2
	s1f64 = l0
	s0f64 = s0f64 / s1f64
	l5 = s0f64
	s0f64 = l1
	s1f64 = l0
	s2f64 = l0
	s1f64 = s1f64 + s2f64
	s0f64 = s0f64 / s1f64
	l6 = s0f64
	s0f64 = l0
	s0f64 = math.Abs(s0f64)
	s1f64 = 1.1920928955078125e-07
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl4
	}
	s0f64 = l6
	s0f64 = math.Abs(s0f64)
	s1f64 = 8.388608e+06
	if s0f64 > s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f64 = l5
	s2f64 = math.Abs(s2f64)
	s3f64 = 8.388608e+06
	if s2f64 > s3f64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3i32 = 1
	s2i32 = s2i32 ^ s3i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f64 = l1
	s0f64 = math.Abs(s0f64)
	s1f64 = 1.1920928955078125e-07
	if s0f64 < s1f64 {
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
	goto lbl0
lbl4:
	s0f64 = l6
	s1f64 = l6
	s0f64 = s0f64 * s1f64
	l0 = s0f64
	s1f64 = l5
	s0i32 = f134(ctx, s0f64, s1f64)
	l4 = s0i32
	s0f64 = l0
	s1f64 = l5
	if s0f64 < s1f64 {
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
		s0i32 = 0
		s1i32 = l4
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl6
		}
	}
	s0i32 = l3
	s1f64 = l0
	s2f64 = l5
	s1f64 = s1f64 - s2f64
	s1f64 = math.Sqrt(s1f64)
	s2f64 = 0
	s3f64 = l0
	s4f64 = l5
	if s3f64 > s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	l0 = s1f64
	s2f64 = l6
	s1f64 = s1f64 - s2f64
	l1 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l3
	s1f64 = l0
	s1f64 = -s1f64
	s2f64 = l6
	s1f64 = s1f64 - s2f64
	l0 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
	s0i32 = 1
	s1i32 = 2
	s2f64 = l1
	s3f64 = l0
	s2i32 = f134(ctx, s2f64, s3f64)
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
lbl6:
	return s0i32
lbl1:
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0f64 = l2
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	return s0i32
lbl0:
	s0i32 = l3
	s1f64 = l2
	s1f64 = -s1f64
	s2f64 = l1
	s1f64 = s1f64 / s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = 1
	return s0i32
}
