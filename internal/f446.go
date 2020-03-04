package internal

import (
	"math"
	"unsafe"
)

func f446(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
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
	var l12 float32
	_ = l12
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
	var s6f32 float32
	_ = s6f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l8 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l7 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l6 = s2f32
	s3f32 = l6
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s1f32 = l5
	s2f32 = l5
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l6
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	l10 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s3f32 = l6
	s4f32 = l7
	s3f32 = s3f32 - s4f32
	s4f32 = 3
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l8
	s2f32 = s2f32 - s3f32
	l8 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l9 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	l11 = s2f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	l7 = s3f32
	s4f32 = l7
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 + s2f32
	l6 = s1f32
	s2f32 = l6
	s3f32 = l6
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l7
	s3f32 = l9
	s2f32 = s2f32 - s3f32
	l12 = s2f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+24)]))
	s4f32 = l7
	s5f32 = l11
	s4f32 = s4f32 - s5f32
	s5f32 = 3
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l9
	s3f32 = s3f32 - s4f32
	l7 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	l9 = s0f32
	s0f32 = l10
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	s1f32 = l12
	s2f32 = l6
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l10 = s0f32
	s0f32 = l5
	s1f32 = 3
	s0f32 = s0f32 * s1f32
	s1f32 = l8
	s0f32 = s0f32 * s1f32
	s1f32 = l6
	s2f32 = 3
	s1f32 = s1f32 * s2f32
	s2f32 = l7
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s0f32 = l8
	s1f32 = l8
	s0f32 = s0f32 * s1f32
	s1f32 = l7
	s2f32 = l7
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l6 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00024414062
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
		s0f32 = l5
		s1f32 = l9
		s2f32 = l10
		s3i32 = l1
		s0i32 = f122(ctx, s0f32, s1f32, s2f32, s3i32)
		l0 = s0i32
		goto lbl0
	}
	s0f32 = 1
	s1f32 = l6
	s0f32 = s0f32 / s1f32
	l6 = s0f32
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	l5 = s0f32
	s1f32 = 3
	s0f32 = s0f32 / s1f32
	l8 = s0f32
	s0f32 = l6
	s1f32 = l10
	s0f32 = s0f32 * s1f32
	s1f32 = 27
	s0f32 = s0f32 * s1f32
	s1f32 = l5
	s2f32 = l5
	s3f32 = l5
	s4f32 = l5
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l5
	s3f32 = 9
	s2f32 = s2f32 * s3f32
	s3f32 = l6
	s4f32 = l9
	s3f32 = s3f32 * s4f32
	l7 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 54
	s0f32 = s0f32 / s1f32
	l6 = s0f32
	s1f32 = l6
	s0f32 = s0f32 * s1f32
	s1f32 = l5
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	s2f32 = l7
	s3f32 = 3
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = 9
	s1f32 = s1f32 / s2f32
	l5 = s1f32
	s2f32 = l5
	s3f32 = l5
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	l7 = s1f32
	s0f32 = s0f32 - s1f32
	l9 = s0f32
	s1f32 = 0
	if s0f32 < s1f32 {
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
		s0i32 = l2
		s1f32 = l6
		s2f32 = l7
		s2f32 = float32(math.Sqrt(float64(s2f32)))
		s1f32 = s1f32 / s2f32
		l6 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l2
		s1i32 = -1082130432
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1065353216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1f32 = l5
		s1f32 = float32(math.Sqrt(float64(s1f32)))
		s2f32 = -2
		s1f32 = s1f32 * s2f32
		l5 = s1f32
		s2i32 = l2
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = l2
		s4i32 = 12
		s3i32 = s3i32 + s4i32
		s4i32 = l2
		s5i32 = 4
		s4i32 = s4i32 + s5i32
		s5f32 = l6
		s6f32 = 1
		if s5f32 < s6f32 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		s4f32 = l6
		s5f32 = -1
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
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2f32 = f1359(ctx, s2f32)
		l6 = s2f32
		s3f32 = 3
		s2f32 = s2f32 / s3f32
		s2f32 = f86(ctx, s2f32)
		s1f32 = s1f32 * s2f32
		s2f32 = l8
		s1f32 = s1f32 - s2f32
		l7 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l2
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1065353216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 12
		s2i32 = s2i32 + s3i32
		s3i32 = l2
		s4i32 = 4
		s3i32 = s3i32 + s4i32
		s4f32 = l7
		s5f32 = 1
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
		s3f32 = l7
		s4f32 = 0
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
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1f32 = l5
		s2f32 = l6
		s3f32 = 6.2831855
		s2f32 = s2f32 + s3f32
		s3f32 = 3
		s2f32 = s2f32 / s3f32
		s2f32 = f86(ctx, s2f32)
		s1f32 = s1f32 * s2f32
		s2f32 = l8
		s1f32 = s1f32 - s2f32
		l7 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l2
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1065353216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 12
		s2i32 = s2i32 + s3i32
		s3i32 = l2
		s4i32 = 4
		s3i32 = s3i32 + s4i32
		s4f32 = l7
		s5f32 = 1
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
		s3f32 = l7
		s4f32 = 0
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
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l0 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1f32 = l5
		s2f32 = l6
		s3f32 = -6.2831855
		s2f32 = s2f32 + s3f32
		s3f32 = 3
		s2f32 = s2f32 / s3f32
		s2f32 = f86(ctx, s2f32)
		s1f32 = s1f32 * s2f32
		s2f32 = l8
		s1f32 = s1f32 - s2f32
		l5 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l2
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1065353216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 12
		s2i32 = s2i32 + s3i32
		s3i32 = l2
		s4i32 = 4
		s3i32 = s3i32 + s4i32
		s4f32 = l5
		s5f32 = 1
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
		s3f32 = l5
		s4f32 = 0
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
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l4
		s0f32 = math.Float32frombits(uint32(s0i32))
		l8 = s0f32
		s0i32 = l3
		s0f32 = math.Float32frombits(uint32(s0i32))
		l5 = s0f32
		s1i32 = l0
		s1f32 = math.Float32frombits(uint32(s1i32))
		l6 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			s0i32 = l3
			l0 = s0i32
			s0f32 = l5
			l7 = s0f32
			s0f32 = l6
			l5 = s0f32
			goto lbl5
		}
		s0i32 = l1
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0f32 = l6
		l7 = s0f32
	lbl5:
		s0f32 = l5
		s1f32 = l8
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0f32 = l5
		l6 = s0f32
		s0f32 = l8
		l5 = s0f32
		goto lbl2
	}
	s0i32 = l2
	s1f32 = l6
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = l9
	s2f32 = float32(math.Sqrt(float64(s2f32)))
	s1f32 = s1f32 + s2f32
	s2f32 = 0.3333333
	s1f32 = f251(ctx, s1f32, s2f32)
	l7 = s1f32
	s1f32 = -s1f32
	s2f32 = l7
	s3f32 = l6
	s4f32 = 0
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
	s2f32 = 0
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l6
		s2f32 = l5
		s3f32 = l6
		s2f32 = s2f32 / s3f32
		s1f32 = s1f32 + s2f32
	} else {
		s1f32 = l6
	}
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	l5 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1065353216
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l2
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = 12
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s4i32 = 4
	s3i32 = s3i32 + s4i32
	s4f32 = l5
	s5f32 = 1
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
	s3f32 = l5
	s4f32 = 0
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
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 1
	l0 = s0i32
	goto lbl0
lbl3:
	s0i32 = l1
	s1f32 = l5
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l1
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0f32 = l8
	l6 = s0f32
lbl2:
	s0f32 = l6
	s1f32 = l5
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
		s0f32 = l6
		l8 = s0f32
		s0f32 = l5
		l6 = s0f32
		goto lbl8
	}
	s0i32 = l1
	s1f32 = l6
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l1
	s1f32 = l5
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0f32 = l5
	l8 = s0f32
lbl8:
	s0f32 = l6
	s1f32 = l8
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = 3
		goto lbl10
	}
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l8
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s0f32 = math.Float32frombits(uint32(s0i32))
	l7 = s0f32
	s0i32 = 2
lbl10:
	l0 = s0i32
	s0f32 = l8
	s1f32 = l7
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l0 = s0i32
lbl0:
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
