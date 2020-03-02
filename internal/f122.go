package internal

import (
	"math"
	"unsafe"
)

func f122(ctx *Context, l0 float32, l1 float32, l2 float32, l3 int32) int32 {
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
	var l9 float64
	_ = l9
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
	var s5f32 float32
	_ = s5f32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	s0f32 = l0
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l2
		s1f32 = l2
		s1f32 = -s1f32
		s2f32 = l2
		s3f32 = 0
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l5 = s2i32
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l0 = s0f32
		s1f32 = l1
		s1f32 = -s1f32
		s2f32 = l1
		s3i32 = l5
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l1 = s1f32
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0f32 = l1
		s1f32 = 0
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0f32 = l0
		s1f32 = 0
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0f32 = l0
		s1f32 = l1
		s0f32 = s0f32 / s1f32
		l0 = s0f32
		s1f32 = 0
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1f32 = l0
		s2f32 = l0
		if s1f32 != s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 | s1i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l3
		s1f32 = l0
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = 1
		return s0i32
	}
	s0f32 = l1
	s0f64 = float64(s0f32)
	l9 = s0f64
	s1f64 = l9
	s0f64 = s0f64 * s1f64
	s1f32 = l0
	s1f64 = float64(s1f32)
	s2f64 = -4
	s1f64 = s1f64 * s2f64
	s2f32 = l2
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	l9 = s0f64
	s1f64 = 0
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f64 = l9
	s0f64 = math.Sqrt(s0f64)
	s0f32 = float32(s0f64)
	l7 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 0
	s2f32 = l7
	s2f32 = -s2f32
	s3f32 = l7
	s4f32 = l1
	s5f32 = 0
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	s3f32 = l1
	s2f32 = s2f32 + s3f32
	l1 = s2f32
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	l8 = s2f32
	s3f32 = l1
	s4f32 = -0.5
	s3f32 = s3f32 * s4f32
	l1 = s3f32
	s4f32 = l1
	s5f32 = 0
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l6 = s4i32
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l7 = s2f32
	s3f32 = l0
	s3f32 = -s3f32
	s4f32 = l0
	s5i32 = l6
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	l0 = s3f32
	if s2f32 >= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl2
	}
	s1i32 = 0
	s2f32 = l0
	s3f32 = 0
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl2
	}
	s1i32 = 0
	s2f32 = l7
	s3f32 = 0
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl2
	}
	s1i32 = 0
	s2f32 = l7
	s3f32 = l0
	s2f32 = s2f32 / s3f32
	l0 = s2f32
	s3f32 = 0
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3f32 = l0
	s4f32 = l0
	if s3f32 != s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 | s3i32
	if s2i32 != 0 {
		goto lbl2
	}
	s1i32 = l3
	s2f32 = l0
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = s2f32
	s1i32 = 1
lbl2:
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0f32 = l2
	s0f32 = -s0f32
	s1f32 = l2
	s2f32 = l2
	s3f32 = 0
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l6 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l0 = s0f32
	s1f32 = l8
	s2f32 = l1
	s3i32 = l6
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l1 = s1f32
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l0
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l1
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l0
	s1f32 = l1
	s0f32 = s0f32 / s1f32
	l0 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l0
	s2f32 = l0
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 | s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l4
	s1f32 = l0
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = 1
	l5 = s0i32
lbl3:
	s0i32 = l4
	s1i32 = l5
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l3
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s1i32 = 8
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0f32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1f32
	if s0f32 > s1f32 {
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
		s0i32 = l3
		s1f32 = l0
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l3
		s1f32 = l1
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = 8
		l4 = s0i32
		goto lbl4
	}
	s0f32 = l0
	s1f32 = l1
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 8
		l4 = s0i32
		goto lbl4
	}
	s0i32 = l5
	s1i32 = l3
	s0i32 = s0i32 - s1i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l4 = s0i32
lbl4:
	s0i32 = l4
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l4 = s0i32
lbl0:
	s0i32 = l4
	return s0i32
}
