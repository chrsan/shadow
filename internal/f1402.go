package internal

import (
	"math"
	"unsafe"
)

func f1402(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l12 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l17 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	l6 = s1i32
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l19 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	l7 = s2i32
	s3f32 = l17
	s4f32 = l19
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
	l14 = s1f32
	s2i32 = l6
	s3i32 = l2
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	l20 = s3f32
	s3i32 = int32(math.Float32bits(s3f32))
	l8 = s3i32
	s4f32 = l17
	s5f32 = l20
	if s4f32 > s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s2f32 = math.Float32frombits(uint32(s2i32))
	l13 = s2f32
	s3f32 = l13
	s4f32 = l14
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
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l14 = s2f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl2
	}
	s0i32 = -2147483648
lbl2:
	l9 = s0i32
	s0f32 = l12
	s1i32 = l6
	s2i32 = l7
	s3f32 = l17
	s4f32 = l19
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
	l12 = s1f32
	s2i32 = l6
	s3i32 = l8
	s4f32 = l17
	s5f32 = l20
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s2f32 = math.Float32frombits(uint32(s2i32))
	l13 = s2f32
	s3f32 = l12
	s4f32 = l13
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
	s2f32 = l14
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	l12 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l12
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl4
	}
	s0i32 = -2147483648
lbl4:
	l6 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l15 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l12 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	l1 = s1i32
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l14 = s2f32
	s2i32 = int32(math.Float32bits(s2f32))
	l3 = s2i32
	s3f32 = l12
	s4f32 = l14
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
	l18 = s1f32
	s2i32 = l1
	s3i32 = l2
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l13 = s3f32
	s3i32 = int32(math.Float32bits(s3f32))
	l7 = s3i32
	s4f32 = l12
	s5f32 = l13
	if s4f32 > s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s2f32 = math.Float32frombits(uint32(s2i32))
	l16 = s2f32
	s3f32 = l16
	s4f32 = l18
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
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l18 = s2f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	l16 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l16
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl6
	}
	s0i32 = -2147483648
lbl6:
	l2 = s0i32
	s0i32 = l6
	s1i32 = l9
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l8 = s0i32
	s0f32 = l15
	s1i32 = l1
	s2i32 = l3
	s3f32 = l12
	s4f32 = l14
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
	l15 = s1f32
	s2i32 = l1
	s3i32 = l7
	s4f32 = l12
	s5f32 = l13
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s2f32 = math.Float32frombits(uint32(s2i32))
	l16 = s2f32
	s3f32 = l15
	s4f32 = l16
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
	s2f32 = l18
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	l15 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l15
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl8
	}
	s0i32 = -2147483648
lbl8:
	l3 = s0i32
	s0i32 = l8
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = l2
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l13
	s1f32 = l12
	s0f32 = s0f32 - s1f32
	l21 = s0f32
	s1f32 = l19
	s2f32 = l20
	s1f32 = s1f32 - s2f32
	l22 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l20
	s2f32 = l17
	s1f32 = s1f32 - s2f32
	l23 = s1f32
	s2f32 = l14
	s3f32 = l13
	s2f32 = s2f32 - s3f32
	l24 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l15 = s0f32
	s0f32 = l17
	s1f32 = l19
	s0f32 = s0f32 - s1f32
	l25 = s0f32
	s0f32 = l12
	s1f32 = l14
	s0f32 = s0f32 - s1f32
	l26 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l8 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l10 = s0i32
lbl10:
	s0i32 = l6
	s1i32 = l10
	s0i32 = s0i32 * s1i32
	l11 = s0i32
	s0i32 = l3
	l1 = s0i32
lbl11:
	s0i32 = l8
	s1i32 = l1
	s2i32 = l11
	s1i32 = s1i32 + s2i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 != 0 {
	lbl13:
		s0i32 = l0
		s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])))
		l7 = s0i32
		s1i32 = l4
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		s0i32 = l5
		s1i32 = l7
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		s0f32 = l15
		s1f32 = l21
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		l18 = s2f32
		s3f32 = l17
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 * s2f32
		s2f32 = l23
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
		l16 = s3f32
		s4f32 = l12
		s3f32 = s3f32 - s4f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 * s1f32
		s1f32 = 0.00024414062
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		s0f32 = l15
		s1f32 = l24
		s2f32 = l18
		s3f32 = l20
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 * s2f32
		s2f32 = l22
		s3f32 = l16
		s4f32 = l13
		s3f32 = s3f32 - s4f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 * s1f32
		s1f32 = 0.00024414062
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		s0i32 = 1
		s1f32 = l15
		s2f32 = l26
		s3f32 = l18
		s4f32 = l19
		s3f32 = s3f32 - s4f32
		s2f32 = s2f32 * s3f32
		s3f32 = l25
		s4f32 = l16
		s5f32 = l14
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 * s2f32
		s2f32 = 0.00024414062
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 != 0 {
			goto lbl0
		}
	lbl14:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l0 = s0i32
		if s0i32 != 0 {
			goto lbl13
		}
	}
	s0i32 = l1
	s1i32 = l2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l6
	s1i32 = l9
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl10
	}
lbl1:
	s0i32 = 0
lbl0:
	return s0i32
}
