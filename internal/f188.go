package internal

import (
	"math"
	"unsafe"
)

func f188(ctx *Context, l0 int32, l1 float32) float32 {
	var l2 int32
	_ = l2
	var l3 float32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
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
	var s6i32 int32
	_ = s6i32
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
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	s0f32 = -1
	s1f32 = 1
	s2f32 = l1
	s3f32 = 0
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
	l6 = s0f32
	s1f32 = l1
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl5
	}
	s0f32 = l1
	s1f32 = l1
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l1
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl6
	}
	s1i32 = -2147483648
lbl6:
	l2 = s1i32
	s1f32 = float32(s1i32)
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l2
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl4
	case 1:
		goto lbl1
	default:
		goto lbl3
	}
lbl5:
	s0f32 = l1
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l5 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	l8 = s2f32
	s1f32 = s1f32 + s2f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l9 = s2f32
	s1f32 = s1f32 + s2f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	l7 = s2f32
	s1f32 = s1f32 + s2f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	l10 = s2f32
	s1f32 = s1f32 + s2f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	l11 = s2f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l5
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l9
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l1
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0f32 = 0
		return s0f32
	}
	s0f32 = l7
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l8
	s1f32 = l5
	s2f32 = l7
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l3
	s1f32 = l7
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0f32 = l6
	s1f32 = l3
	s2f32 = l9
	s1f32 = s1f32 * s2f32
	s2f32 = l11
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	return s0f32
lbl4:
	s0f32 = l3
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 * s1f32
	l1 = s0f32
	s1f32 = 1
	if s0f32 <= s1f32 {
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
		s0f32 = l1
		s1f32 = 0
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0f32 = l1
		s1f32 = 1
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1f32 = l1
		s1i32 = int32(math.Float32bits(s1f32))
		l0 = s1i32
		s1f32 = float32(s1i32)
		s2f32 = 1.1920929e-07
		s1f32 = s1f32 * s2f32
		s2f32 = -124.22552
		s1f32 = s1f32 + s2f32
		s2i32 = l0
		s3i32 = 8388607
		s2i32 = s2i32 & s3i32
		s3i32 = 1056964608
		s2i32 = s2i32 | s3i32
		s2f32 = math.Float32frombits(uint32(s2i32))
		l1 = s2f32
		s3f32 = -1.4980303
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = -1.72588
		s3f32 = l1
		s4f32 = 0.35208872
		s3f32 = s3f32 + s4f32
		s2f32 = s2f32 / s3f32
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 * s1f32
		l3 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l3
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl11
		}
		s0i32 = -2147483648
	lbl11:
		l0 = s0i32
		s0f32 = float32(math.Inf(0))
		l1 = s0f32
		s0f32 = l3
		s1f32 = 121.274055
		s0f32 = s0f32 + s1f32
		s1f32 = l3
		s2i32 = l0
		s2f32 = float32(s2i32)
		l4 = s2f32
		s3f32 = -1
		s2f32 = s2f32 + s3f32
		s3f32 = l4
		s4f32 = l3
		s5f32 = l4
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
		s1f32 = s1f32 - s2f32
		l3 = s1f32
		s2f32 = -1.4901291
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 27.728024
		s2f32 = 4.8425255
		s3f32 = l3
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 / s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 8.388608e+06
		s0f32 = s0f32 * s1f32
		l3 = s0f32
		s1f32 = 2.1474836e+09
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0f32 = float32(math.Inf(0))
		l1 = s0f32
		s0f32 = l3
		s1f32 = -2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0f32 = l3
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l6
			s1f32 = l3
			s1i32 = int32(math.Trunc(float64(s1f32)))
			s1f32 = math.Float32frombits(uint32(s1i32))
			s0f32 = s0f32 * s1f32
			return s0f32
		}
		s0f32 = l6
		s1f32 = -0
		s0f32 = s0f32 * s1f32
		return s0f32
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l5 = s0f32
	s0f32 = l3
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s0f32 = s0f32 - s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0f32 = s0f32 * s1f32
	s1f32 = 1.442695
	s0f32 = s0f32 * s1f32
	l1 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l1
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl14
	}
	s0i32 = -2147483648
lbl14:
	l0 = s0i32
	s0f32 = l5
	s1f32 = float32(math.Inf(0))
	s2f32 = l1
	s3f32 = 121.274055
	s2f32 = s2f32 + s3f32
	s3f32 = l1
	s4i32 = l0
	s4f32 = float32(s4i32)
	l4 = s4f32
	s5f32 = -1
	s4f32 = s4f32 + s5f32
	s5f32 = l4
	s6f32 = l1
	s7f32 = l4
	if s6f32 < s7f32 {
		s6i32 = 1
	} else {
		s6i32 = 0
	}
	if s6i32 != 0 {
		// s4f32 = s4f32
	} else {
		s4f32 = s5f32
	}
	s3f32 = s3f32 - s4f32
	l1 = s3f32
	s4f32 = -1.4901291
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = 27.728024
	s4f32 = 4.8425255
	s5f32 = l1
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 / s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = 8.388608e+06
	s2f32 = s2f32 * s3f32
	l1 = s2f32
	s3f32 = 2.1474836e+09
	if s2f32 >= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl16
	}
	s1f32 = float32(math.Inf(0))
	s2f32 = l1
	s3f32 = -2.1474836e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl16
	}
	s1f32 = l1
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l1
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl17
	}
	s1i32 = -2147483648
lbl17:
	s1f32 = math.Float32frombits(uint32(s1i32))
lbl16:
	s0f32 = s0f32 + s1f32
	l1 = s0f32
lbl9:
	s0f32 = l6
	s1f32 = l1
	s0f32 = s0f32 * s1f32
	return s0f32
lbl3:
	s0f32 = l3
	s1f32 = 1
	if s0f32 <= s1f32 {
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
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l5 = s0f32
		s0f32 = l3
		s1f32 = 0
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0f32 = l3
		s1f32 = 1
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0f32 = l3
		s0i32 = int32(math.Float32bits(s0f32))
		l2 = s0i32
		s0f32 = float32(s0i32)
		s1f32 = 1.1920929e-07
		s0f32 = s0f32 * s1f32
		s1f32 = -124.22552
		s0f32 = s0f32 + s1f32
		s1i32 = l2
		s2i32 = 8388607
		s1i32 = s1i32 & s2i32
		s2i32 = 1056964608
		s1i32 = s1i32 | s2i32
		s1f32 = math.Float32frombits(uint32(s1i32))
		l1 = s1f32
		s2f32 = -1.4980303
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = -1.72588
		s2f32 = l1
		s3f32 = 0.35208872
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 / s2f32
		s0f32 = s0f32 + s1f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0f32 = s0f32 * s1f32
		l1 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l1
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl21
		}
		s0i32 = -2147483648
	lbl21:
		l0 = s0i32
		s0f32 = float32(math.Inf(0))
		l3 = s0f32
		s0f32 = l1
		s1f32 = 121.274055
		s0f32 = s0f32 + s1f32
		s1f32 = l1
		s2i32 = l0
		s2f32 = float32(s2i32)
		l4 = s2f32
		s3f32 = -1
		s2f32 = s2f32 + s3f32
		s3f32 = l4
		s4f32 = l1
		s5f32 = l4
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
		s1f32 = s1f32 - s2f32
		l1 = s1f32
		s2f32 = -1.4901291
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 27.728024
		s2f32 = 4.8425255
		s3f32 = l1
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 / s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 8.388608e+06
		s0f32 = s0f32 * s1f32
		l1 = s0f32
		s1f32 = 2.1474836e+09
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0f32 = float32(math.Inf(0))
		l3 = s0f32
		s0f32 = l1
		s1f32 = -2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0f32 = l1
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l1
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl23
		}
		s0i32 = -2147483648
	lbl23:
		s0f32 = math.Float32frombits(uint32(s0i32))
		l3 = s0f32
	lbl20:
		s0f32 = l6
		s1f32 = l5
		s2f32 = l3
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 * s1f32
		return s0f32
	}
	s0f32 = l6
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3f32 = l3
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
	s3f32 = s3f32 - s4f32
	s3i32 = int32(math.Float32bits(s3f32))
	l0 = s3i32
	s3f32 = float32(s3i32)
	s4f32 = 1.1920929e-07
	s3f32 = s3f32 * s4f32
	s4f32 = -124.22552
	s3f32 = s3f32 + s4f32
	s4i32 = l0
	s5i32 = 8388607
	s4i32 = s4i32 & s5i32
	s5i32 = 1056964608
	s4i32 = s4i32 | s5i32
	s4f32 = math.Float32frombits(uint32(s4i32))
	l1 = s4f32
	s5f32 = -1.4980303
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = -1.72588
	s5f32 = l1
	s6f32 = 0.35208872
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 / s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = 0.6931472
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	return s0f32
lbl2:
	s0f32 = l3
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	s1f32 = l8
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl25
	}
	s0f32 = l4
	s1f32 = 1
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl25
	}
	s0f32 = l1
	s1f32 = l4
	s1i32 = int32(math.Float32bits(s1f32))
	l0 = s1i32
	s1f32 = float32(s1i32)
	s2f32 = 1.1920929e-07
	s1f32 = s1f32 * s2f32
	s2f32 = -124.22552
	s1f32 = s1f32 + s2f32
	s2i32 = l0
	s3i32 = 8388607
	s2i32 = s2i32 & s3i32
	s3i32 = 1056964608
	s2i32 = s2i32 | s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l1 = s2f32
	s3f32 = -1.4980303
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = -1.72588
	s3f32 = l1
	s4f32 = 0.35208872
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 / s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	l1 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l1
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl26
	}
	s0i32 = -2147483648
lbl26:
	l0 = s0i32
	s0f32 = float32(math.Inf(0))
	l4 = s0f32
	s0f32 = l1
	s1f32 = 121.274055
	s0f32 = s0f32 + s1f32
	s1f32 = l1
	s2i32 = l0
	s2f32 = float32(s2i32)
	l3 = s2f32
	s3f32 = -1
	s2f32 = s2f32 + s3f32
	s3f32 = l3
	s4f32 = l1
	s5f32 = l3
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
	s1f32 = s1f32 - s2f32
	l1 = s1f32
	s2f32 = -1.4901291
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 27.728024
	s2f32 = 4.8425255
	s3f32 = l1
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 8.388608e+06
	s0f32 = s0f32 * s1f32
	l1 = s0f32
	s1f32 = 2.1474836e+09
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl25
	}
	s0f32 = float32(math.Inf(0))
	l4 = s0f32
	s0f32 = l1
	s1f32 = -2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl25
	}
	s0f32 = l1
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l1
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl28
	}
	s0i32 = -2147483648
lbl28:
	s0f32 = math.Float32frombits(uint32(s0i32))
	l4 = s0f32
lbl25:
	s0f32 = l6
	s1f32 = l4
	s2f32 = l10
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	return s0f32
lbl1:
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l8 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l9 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l10 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l5 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0f32
	s0f32 = l3
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl34
	}
	s0f32 = l3
	s1f32 = 1
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl34
	}
	s0f32 = l3
	s0i32 = int32(math.Float32bits(s0f32))
	l2 = s0i32
	s0f32 = float32(s0i32)
	s1f32 = 1.1920929e-07
	s0f32 = s0f32 * s1f32
	s1f32 = -124.22552
	s0f32 = s0f32 + s1f32
	s1i32 = l2
	s2i32 = 8388607
	s1i32 = s1i32 & s2i32
	s2i32 = 1056964608
	s1i32 = s1i32 | s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l1 = s1f32
	s2f32 = -1.4980303
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = -1.72588
	s2f32 = l1
	s3f32 = 0.35208872
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0f32 = s0f32 * s1f32
	l1 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l1
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl35
	}
	s0i32 = -2147483648
lbl35:
	l0 = s0i32
	s0f32 = l1
	s1f32 = 121.274055
	s0f32 = s0f32 + s1f32
	s1f32 = l1
	s2i32 = l0
	s2f32 = float32(s2i32)
	l3 = s2f32
	s3f32 = -1
	s2f32 = s2f32 + s3f32
	s3f32 = l3
	s4f32 = l1
	s5f32 = l3
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
	s1f32 = s1f32 - s2f32
	l1 = s1f32
	s2f32 = -1.4901291
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 27.728024
	s2f32 = 4.8425255
	s3f32 = l1
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 8.388608e+06
	s0f32 = s0f32 * s1f32
	l4 = s0f32
	s1f32 = 2.1474836e+09
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl33
	}
	s0f32 = l4
	s1f32 = -2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0f32 = float32(math.Inf(0))
	l3 = s0f32
	s0f32 = l7
	s1f32 = l5
	s2f32 = float32(math.Inf(0))
	s3f32 = l4
	s3f32 = float32(math.Abs(float64(s3f32)))
	l11 = s3f32
	s4f32 = 2.1474836e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		s3f32 = l4
		s3i32 = int32(math.Trunc(float64(s3f32)))
		goto lbl37
	}
	s3i32 = -2147483648
lbl37:
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4i32 = l0
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l1 = s0f32
	s1f32 = 0
	s2f32 = l1
	s3f32 = 0
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
	l1 = s0f32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl30
	}
	s0f32 = l11
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl32
	}
	s0f32 = l4
	s0i32 = int32(math.Trunc(float64(s0f32)))
	goto lbl31
lbl34:
	s0f32 = l7
	s1f32 = l3
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l1 = s0f32
	s1f32 = 0
	s2f32 = l1
	s3f32 = 0
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
	l1 = s0f32
	goto lbl30
lbl33:
	s0f32 = float32(math.Inf(0))
	l3 = s0f32
	s0f32 = l7
	s1f32 = l5
	s2f32 = float32(math.Inf(0))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l1 = s0f32
	s1f32 = 0
	s2f32 = l1
	s3f32 = 0
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
	l1 = s0f32
	goto lbl30
lbl32:
	s0i32 = -2147483648
lbl31:
	s0f32 = math.Float32frombits(uint32(s0i32))
	l3 = s0f32
lbl30:
	s0f32 = l1
	s1f32 = l10
	s2f32 = l9
	s3f32 = l3
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 / s1f32
	l1 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0f32 = l1
	s1f32 = 1
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0f32 = l8
	s1f32 = l1
	s1i32 = int32(math.Float32bits(s1f32))
	l0 = s1i32
	s1f32 = float32(s1i32)
	s2f32 = 1.1920929e-07
	s1f32 = s1f32 * s2f32
	s2f32 = -124.22552
	s1f32 = s1f32 + s2f32
	s2i32 = l0
	s3i32 = 8388607
	s2i32 = s2i32 & s3i32
	s3i32 = 1056964608
	s2i32 = s2i32 | s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l1 = s2f32
	s3f32 = -1.4980303
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = -1.72588
	s3f32 = l1
	s4f32 = 0.35208872
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 / s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l3
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl40
	}
	s0i32 = -2147483648
lbl40:
	l0 = s0i32
	s0f32 = float32(math.Inf(0))
	l1 = s0f32
	s0f32 = l3
	s1f32 = 121.274055
	s0f32 = s0f32 + s1f32
	s1f32 = l3
	s2i32 = l0
	s2f32 = float32(s2i32)
	l4 = s2f32
	s3f32 = -1
	s2f32 = s2f32 + s3f32
	s3f32 = l4
	s4f32 = l3
	s5f32 = l4
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
	s1f32 = s1f32 - s2f32
	l3 = s1f32
	s2f32 = -1.4901291
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 27.728024
	s2f32 = 4.8425255
	s3f32 = l3
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 8.388608e+06
	s0f32 = s0f32 * s1f32
	l3 = s0f32
	s1f32 = 2.1474836e+09
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0f32 = float32(math.Inf(0))
	l1 = s0f32
	s0f32 = l3
	s1f32 = -2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0f32 = l3
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l3
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl42
	}
	s0i32 = -2147483648
lbl42:
	s0f32 = math.Float32frombits(uint32(s0i32))
	l1 = s0f32
lbl39:
	s0f32 = l6
	s1f32 = l1
	s0f32 = s0f32 * s1f32
	l4 = s0f32
lbl0:
	s0f32 = l4
	return s0f32
}
