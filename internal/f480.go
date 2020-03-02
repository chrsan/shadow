package internal

import (
	"math"
	"unsafe"
)

func f480(ctx *Context, l0 float32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
	var l23 float32
	_ = l23
	var l24 float32
	_ = l24
	var l25 float32
	_ = l25
	var l26 float32
	_ = l26
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
	s0f32 = l0
	s1f32 = 6
	s0f32 = s0f32 * s1f32
	s0f32 = float32(math.Ceil(float64(s0f32)))
	l21 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l21
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
	l21 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l21
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
	l21 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l21
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl0
	}
	s0i32 = -2147483648
lbl0:
	l8 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = l8
	s1i32 = 1
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l6 = s0i32
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l21 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l24 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l25 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l26 = s0f32
	s0i32 = l1
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l26
	s2i32 = l6
	s2f32 = float32(s2i32)
	l22 = s2f32
	s1f32 = s1f32 + s2f32
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l23 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l23
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
	l23 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l23
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
	l23 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l23
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl4
	}
	s1i32 = -2147483648
lbl4:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l25
	s2f32 = l22
	s1f32 = s1f32 + s2f32
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l23 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l23
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
	l23 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l23
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
	l23 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l23
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl6
	}
	s1i32 = -2147483648
lbl6:
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l24
	s2f32 = l22
	s1f32 = s1f32 - s2f32
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l23 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l23
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
	l23 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l23
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
	l23 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l23
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl8
	}
	s1i32 = -2147483648
lbl8:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l21
	s2f32 = l22
	s1f32 = s1f32 - s2f32
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l22 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l22
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
	l22 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l22
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
	l22 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l22
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl10
	}
	s1i32 = -2147483648
lbl10:
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l4
	s2i32 = l7
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0f32 = l26
	s1f32 = l24
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	l22 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l22
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
	l22 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l22
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
	l22 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l22
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl12
	}
	s0i32 = -2147483648
lbl12:
	l14 = s0i32
	s0f32 = l25
	s1f32 = l21
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	l22 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l22
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
	l22 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l22
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
	l22 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l22
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl14
	}
	s0i32 = -2147483648
lbl14:
	l9 = s0i32
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l15 = s0i32
		s0i32 = l3
		s1i32 = 3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l1
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l1
		s1f32 = l26
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l0 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l0
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
		l0 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l0
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
		l0 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l0
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl18
		}
		s1i32 = -2147483648
	lbl18:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l1
		s1f32 = l25
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l0 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l0
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
		l0 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l0
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
		l0 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l0
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl20
		}
		s1i32 = -2147483648
	lbl20:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l1
		s1f32 = l24
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l0 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l0
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
		l0 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l0
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
		l0 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l0
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl22
		}
		s1i32 = -2147483648
	lbl22:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0f32 = l21
		s1f32 = 0.5
		s0f32 = s0f32 + s1f32
		s0f32 = float32(math.Floor(float64(s0f32)))
		l0 = s0f32
		s1f32 = 2.1474835e+09
		s2f32 = l0
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
		l0 = s0f32
		s1f32 = -2.1474835e+09
		s2f32 = l0
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
		l0 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1f32 = l0
			s1i32 = int32(math.Trunc(float64(s1f32)))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = 1
			return s0i32
		}
		s0i32 = l1
		s1i32 = -2147483648
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = 1
		return s0i32
	}
	s0i32 = 1
	l4 = s0i32
	s0i32 = l8
	s1i32 = 1
	s0i32 = f50(ctx, s0i32, s1i32)
	l13 = s0i32
	s1i32 = 255
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l8
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 1
		s1f32 = l0
		s2f32 = l0
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 / s1f32
		l26 = s0f32
	lbl26:
		s0i32 = l4
		s1i32 = l13
		s0i32 = s0i32 + s1i32
		s1f32 = 0
		s2f32 = l26
		s3i32 = l6
		s4i32 = l4
		s3i32 = s3i32 - s4i32
		s3f32 = float32(s3i32)
		s4f32 = -0.5
		s3f32 = s3f32 + s4f32
		s2f32 = s2f32 * s3f32
		l21 = s2f32
		s3f32 = 1.5
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl28
		}
		s1f32 = 1
		s2f32 = l21
		s3f32 = -1.5
		if s2f32 < s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl28
		}
		s1f32 = 0.5625
		s2f32 = l21
		s3f32 = 1.125
		s2f32 = s2f32 * s3f32
		s3f32 = l21
		s4f32 = l21
		s5f32 = l21
		s4f32 = s4f32 * s5f32
		l24 = s4f32
		s3f32 = s3f32 * s4f32
		l25 = s3f32
		s4f32 = 6
		s3f32 = s3f32 / s4f32
		s4f32 = l24
		s5f32 = -3
		s4f32 = s4f32 * s5f32
		s5f32 = 0.25
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 + s4f32
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 - s2f32
		s2f32 = l21
		s3f32 = 0.5
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s3i32 = 1
		s2i32 = s2i32 ^ s3i32
		if s2i32 == 0 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl28
		}
		s1f32 = l25
		s2f32 = 3
		s1f32 = s1f32 / s2f32
		s2f32 = l21
		s3f32 = -0.75
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s2f32 = l21
		s3f32 = -0.5
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s3i32 = 1
		s2i32 = s2i32 ^ s3i32
		if s2i32 == 0 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl28
		}
		s1f32 = l25
		s2f32 = -6
		s1f32 = s1f32 / s2f32
		s2f32 = l24
		s3f32 = -3
		s2f32 = s2f32 * s3f32
		s3f32 = 0.25
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l21
		s3f32 = -1.125
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = 0.4375
		s1f32 = s1f32 + s2f32
	lbl28:
		s2f32 = 255
		s1f32 = s1f32 * s2f32
		l21 = s1f32
		s2f32 = 4.2949673e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2f32 = l21
		s3f32 = 0
		if s2f32 >= s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s1i32 = s1i32 & s2i32
		if s1i32 != 0 {
			s1f32 = l21
			s1i32 = int32(uint32(math.Trunc(float64(s1f32))))
			goto lbl27
		}
		s1i32 = 0
	lbl27:
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l8
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl26
		}
	}
	s0i32 = l1
	s0i32 = f118(ctx, s0i32)
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l1
	s1i32 = l4
	s2i32 = 0
	s1i32 = f203(ctx, s1i32, s2i32)
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l16 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l17 = s1i32
	s0i32 = s0i32 - s1i32
	l10 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 1
		s0i32 = f50(ctx, s0i32, s1i32)
		l11 = s0i32
	}
	s0i32 = l10
	if s0i32 != 0 {
		s0i32 = l10
		s1i32 = 1
		s0i32 = f50(ctx, s0i32, s1i32)
		l12 = s0i32
	}
	s0i32 = l11
	s1i32 = l13
	s2i32 = l7
	s3f32 = l0
	f481(ctx, s0i32, s1i32, s2i32, s3f32)
	s0i32 = l12
	s1i32 = l13
	s2i32 = l10
	s3f32 = l0
	f481(ctx, s0i32, s1i32, s2i32, s3f32)
	s0i32 = 1
	l15 = s0i32
	s0i32 = l10
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl33
	}
	s0i32 = l7
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl33
	}
	s0i32 = l5
	l8 = s0i32
lbl34:
	s0i32 = l12
	s1i32 = l18
	s0i32 = s0i32 + s1i32
	l19 = s0i32
	s0i32 = 0
	l4 = s0i32
lbl35:
	s0i32 = l8
	s1i32 = l19
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	s2i32 = l4
	s3i32 = l11
	s2i32 = s2i32 + s3i32
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	s1i32 = s1i32 * s2i32
	s2i32 = 128
	s1i32 = s1i32 + s2i32
	l20 = s1i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = l20
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l8
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l7
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl35
	}
	s0i32 = l18
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l18 = s0i32
	s1i32 = l10
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl34
	}
lbl33:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl36
	}
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl38
	case 1:
		goto lbl37
	default:
		goto lbl39
	}
lbl39:
	s0i32 = l6
	s1i32 = l10
	s2i32 = l6
	s1i32 = s1i32 - s2i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl36
	}
	s0i32 = l5
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l16
	s1i32 = l6
	s2i32 = l17
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
lbl40:
	s0i32 = l1
	s1i32 = l6
	s2i32 = l7
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = 255
	s2i32 = l9
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl40
	}
	goto lbl36
lbl38:
	s0i32 = l6
	s1i32 = l10
	s2i32 = l6
	s1i32 = s1i32 - s2i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl36
	}
	s0i32 = l5
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l16
	s1i32 = l6
	s2i32 = l17
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
lbl41:
	s0i32 = l1
	s1i32 = l6
	s2i32 = l7
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	s2i32 = l9
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl41
	}
	goto lbl36
lbl37:
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	l0 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l0
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l0
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl42
	}
	s0i32 = 0
lbl42:
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l15 = s0i32
		goto lbl36
	}
	s0i32 = l1
	s1i32 = l3
	s2i32 = 0
	s1i32 = f203(ctx, s1i32, s2i32)
	l3 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 1
	l4 = s0i32
	s0i32 = l14
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl45
	}
	s0i32 = l3
	s1i32 = l5
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s2i32 = l6
	s3i32 = l7
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = l9
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	s0i32 = l14
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl45
	}
lbl46:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l4
	s2i32 = l9
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = l4
	s3i32 = l6
	s2i32 = s2i32 + s3i32
	s3i32 = l7
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = l9
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l14
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl46
	}
lbl45:
	s0i32 = l5
	f13(ctx, s0i32)
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l21 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l24 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l25 = s0f32
	s0i32 = l1
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l25
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l25 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l25
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
	l25 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l25
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
	l25 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l25
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl47
	}
	s1i32 = -2147483648
lbl47:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l24
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l24 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l24
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
	l24 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l24
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
	l24 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l24
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl49
	}
	s1i32 = -2147483648
lbl49:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l21
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l21 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l21
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
	l21 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l21
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
	l21 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l21
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl51
	}
	s1i32 = -2147483648
lbl51:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l0
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l0 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l0
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
	l0 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l0
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
	l0 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l0
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl53
	}
	s1i32 = -2147483648
lbl53:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
lbl36:
	s0i32 = l12
	if s0i32 != 0 {
		s0i32 = l12
		f13(ctx, s0i32)
	}
	s0i32 = l11
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l11
	f13(ctx, s0i32)
lbl30:
	s0i32 = l13
	f13(ctx, s0i32)
lbl16:
	s0i32 = l15
	return s0i32
}
