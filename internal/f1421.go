package internal

import (
	"math"
	"unsafe"
)

func f1421(ctx *Context, l0 float64, l1 float64, l2 float64, l3 float64, l4 int32) int32 {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 float64
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
	var s5f64 float64
	_ = s5f64
	var s6f64 float64
	_ = s6f64
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0f64 = l0
	s0f64 = math.Abs(s0f64)
	l8 = s0f64
	s1f64 = 1.1920928955078125e-07
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
	s0f64 = l0
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l8
		s1f64 = l1
		s2f64 = 1.1920928955078125e-07
		s1f64 = s1f64 * s2f64
		s1f64 = math.Abs(s1f64)
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
	}
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l8
		s1f64 = l2
		s2f64 = 1.1920928955078125e-07
		s1f64 = s1f64 * s2f64
		s1f64 = math.Abs(s1f64)
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
	}
	s0f64 = l0
	s1f64 = 0
	if s0f64 != s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l8
		s1f64 = l3
		s2f64 = 1.1920928955078125e-07
		s1f64 = s1f64 * s2f64
		s1f64 = math.Abs(s1f64)
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
	}
	s0f64 = l1
	s1f64 = l2
	s2f64 = l3
	s3i32 = l4
	s0i32 = f382(ctx, s0f64, s1f64, s2f64, s3i32)
	l6 = s0i32
	goto lbl0
lbl1:
	s0f64 = l3
	s0f64 = math.Abs(s0f64)
	l8 = s0f64
	s0f64 = l3
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l8
		s1f64 = l0
		s2f64 = 1.1920928955078125e-07
		s1f64 = s1f64 * s2f64
		s1f64 = math.Abs(s1f64)
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
	}
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l8
		s1f64 = l1
		s2f64 = 1.1920928955078125e-07
		s1f64 = s1f64 * s2f64
		s1f64 = math.Abs(s1f64)
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
	}
	s0f64 = l3
	s1f64 = 0
	if s0f64 != s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l8
		s1f64 = l2
		s2f64 = 1.1920928955078125e-07
		s1f64 = s1f64 * s2f64
		s1f64 = math.Abs(s1f64)
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
	}
	s0i32 = 0
	l5 = s0i32
	s0f64 = l0
	s1f64 = l1
	s2f64 = l2
	s3i32 = l4
	s0i32 = f382(ctx, s0f64, s1f64, s2f64, s3i32)
	l6 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl10:
		s0i32 = l4
		s1i32 = l5
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
			goto lbl0
		}
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
	}
	s0i32 = l4
	s1i32 = l6
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	goto lbl0
lbl5:
	s0f64 = l0
	s1f64 = l1
	s0f64 = s0f64 + s1f64
	l8 = s0f64
	s1f64 = l2
	s0f64 = s0f64 + s1f64
	s1f64 = l3
	s0f64 = s0f64 + s1f64
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
		s0i32 = 0
		l5 = s0i32
		s0f64 = l0
		s1f64 = l8
		s2f64 = l3
		s2f64 = -s2f64
		s3i32 = l4
		s0i32 = f382(ctx, s0f64, s1f64, s2f64, s3i32)
		l6 = s0i32
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl13:
			s0i32 = l4
			s1i32 = l5
			s2i32 = 3
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s1f64 = 1
			s0i32 = f134(ctx, s0f64, s1f64)
			if s0i32 != 0 {
				goto lbl0
			}
			s0i32 = l5
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s1i32 = l6
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl13
			}
		}
		s0i32 = l4
		s1i32 = l6
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i64 = 4607182418800017408
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		goto lbl0
	}
	s0f64 = 1
	s1f64 = l0
	s0f64 = s0f64 / s1f64
	l8 = s0f64
	s1f64 = l1
	s0f64 = s0f64 * s1f64
	l1 = s0f64
	s1f64 = 3
	s0f64 = s0f64 / s1f64
	l0 = s0f64
	s0f64 = l8
	s1f64 = l3
	s0f64 = s0f64 * s1f64
	s1f64 = 27
	s0f64 = s0f64 * s1f64
	s1f64 = l1
	s2f64 = l1
	s3f64 = l1
	s2f64 = s2f64 * s3f64
	l3 = s2f64
	s3f64 = l3
	s2f64 = s2f64 + s3f64
	s1f64 = s1f64 * s2f64
	s2f64 = l1
	s3f64 = 9
	s2f64 = s2f64 * s3f64
	s3f64 = l8
	s4f64 = l2
	s3f64 = s3f64 * s4f64
	l2 = s3f64
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s0f64 = s0f64 + s1f64
	s1f64 = 54
	s0f64 = s0f64 / s1f64
	l1 = s0f64
	s1f64 = l1
	s0f64 = s0f64 * s1f64
	l9 = s0f64
	s1f64 = l3
	s2f64 = l2
	s3f64 = 3
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s2f64 = 9
	s1f64 = s1f64 / s2f64
	l2 = s1f64
	s2f64 = l2
	s3f64 = l2
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 * s2f64
	l8 = s1f64
	s0f64 = s0f64 - s1f64
	l3 = s0f64
	s1f64 = 0
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
		s0i32 = l7
		s1f64 = l1
		s2f64 = l8
		s2f64 = math.Sqrt(s2f64)
		s1f64 = s1f64 / s2f64
		l1 = s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f64
		s0i32 = l7
		s1i64 = -4616189618054758400
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 4607182418800017408
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l4
		s1f64 = l2
		s1f64 = math.Sqrt(s1f64)
		s2f64 = -2
		s1f64 = s1f64 * s2f64
		l2 = s1f64
		s2i32 = l7
		s3i32 = 16
		s2i32 = s2i32 + s3i32
		s3i32 = l7
		s4i32 = 24
		s3i32 = s3i32 + s4i32
		s4i32 = l7
		s5i32 = 8
		s4i32 = s4i32 + s5i32
		s5f64 = l1
		s6f64 = 1
		if s5f64 < s6f64 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		s4f64 = l1
		s5f64 = -1
		if s4f64 < s5f64 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2f64 = f1361(ctx, s2f64)
		l1 = s2f64
		s3f64 = 3
		s2f64 = s2f64 / s3f64
		s2f64 = f372(ctx, s2f64)
		s1f64 = s1f64 * s2f64
		s2f64 = l0
		s1f64 = s1f64 - s2f64
		l3 = s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l4
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0f64 = l3
		s1f64 = l2
		s2f64 = l1
		s3f64 = 6.283185307179586
		s2f64 = s2f64 + s3f64
		s3f64 = 3
		s2f64 = s2f64 / s3f64
		s2f64 = f372(ctx, s2f64)
		s1f64 = s1f64 * s2f64
		s2f64 = l0
		s1f64 = s1f64 - s2f64
		l3 = s1f64
		s0i32 = f134(ctx, s0f64, s1f64)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1f64 = l3
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
			s0i32 = l4
			s1i32 = 16
			s0i32 = s0i32 + s1i32
			l5 = s0i32
		}
		s0f64 = l1
		s1f64 = -6.283185307179586
		s0f64 = s0f64 + s1f64
		s1f64 = 3
		s0f64 = s0f64 / s1f64
		s0f64 = f372(ctx, s0f64)
		l1 = s0f64
		s0i32 = l4
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1f64 = l2
		s2f64 = l1
		s1f64 = s1f64 * s2f64
		s2f64 = l0
		s1f64 = s1f64 - s2f64
		l0 = s1f64
		s0i32 = f134(ctx, s0f64, s1f64)
		if s0i32 != 0 {
			goto lbl14
		}
		s0i32 = l5
		s1i32 = l4
		s0i32 = s0i32 - s1i32
		s1i32 = 8
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			s1f64 = l0
			s0i32 = f134(ctx, s0f64, s1f64)
			if s0i32 != 0 {
				goto lbl14
			}
		}
		s0i32 = l5
		s1f64 = l0
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l5
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		goto lbl14
	}
	s0f64 = l1
	s0f64 = math.Abs(s0f64)
	s1f64 = l3
	s1f64 = math.Sqrt(s1f64)
	s0f64 = s0f64 + s1f64
	s0f64 = f1417(ctx, s0f64)
	l3 = s0f64
	s0f64 = -s0f64
	s1f64 = l3
	s2f64 = l1
	s3f64 = 0
	if s2f64 > s3f64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	l3 = s0f64
	s1f64 = 0
	if s0f64 != s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l3
		s1f64 = l2
		s2f64 = l3
		s1f64 = s1f64 / s2f64
		s0f64 = s0f64 + s1f64
		l3 = s0f64
	}
	s0i32 = l4
	s1f64 = l3
	s2f64 = l0
	s1f64 = s1f64 - s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0f64 = l9
	s1f64 = l8
	s0i32 = f134(ctx, s0f64, s1f64)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f64 = l3
	s2f64 = -0.5
	s1f64 = s1f64 * s2f64
	s2f64 = l0
	s1f64 = s1f64 - s2f64
	l0 = s1f64
	s0i32 = f134(ctx, s0f64, s1f64)
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l4
	s1f64 = l0
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l5 = s0i32
lbl14:
	s0i32 = l5
	s1i32 = l4
	s0i32 = s0i32 - s1i32
	s1i32 = 3
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l6 = s0i32
lbl0:
	s0i32 = l7
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l6
	return s0i32
}
