package internal

import (
	"math"
	"unsafe"
)

func f442(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int64
	_ = l8
	var l9 int64
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float64
	_ = l16
	var l17 float64
	_ = l17
	var l18 float64
	_ = l18
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
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
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
	var s3f64 float64
	_ = s3f64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l10 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l13 = s2f32
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	l6 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l11 = s1f32
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l6
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	l7 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l12 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l14 = s1f32
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0f32 = l12
	s1f32 = l11
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
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l12 = s0f32
		s0f32 = l13
		s1f32 = l10
		s0f32 = s0f32 - s1f32
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
			s0f32 = l12
			s1i32 = l0
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			s0f32 = s0f32 + s1f32
			s1f32 = 0.5
			s0f32 = s0f32 * s1f32
			goto lbl2
		}
		s0f32 = l11
		s0f64 = float64(s0f32)
		s1f32 = l10
		s1f64 = float64(s1f32)
		l17 = s1f64
		s0f64 = s0f64 - s1f64
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l15 = s1f32
		s1f64 = float64(s1f32)
		l18 = s1f64
		s2f32 = l12
		s2f64 = float64(s2f32)
		l16 = s2f64
		s1f64 = s1f64 - s2f64
		s0f64 = s0f64 * s1f64
		s1f32 = l13
		s1f64 = float64(s1f32)
		s2f64 = l17
		s1f64 = s1f64 - s2f64
		s0f64 = s0f64 / s1f64
		s1f64 = l16
		s0f64 = s0f64 + s1f64
		l17 = s0f64
		s1f64 = l18
		s1i64 = int64(math.Float64bits(s1f64))
		l8 = s1i64
		s2f64 = l16
		s2i64 = int64(math.Float64bits(s2f64))
		l9 = s2i64
		s3f32 = l15
		s4f32 = l12
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l5 = s3i32
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		s1f64 = math.Float64frombits(uint64(s1i64))
		l16 = s1f64
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0f64 = l17
		l16 = s0f64
		s1i64 = l9
		s2i64 = l8
		s3i32 = l5
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		s1f64 = math.Float64frombits(uint64(s1i64))
		l17 = s1f64
		if s0f64 > s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0f64 = l17
		l16 = s0f64
	lbl4:
		s0f64 = l16
		s0f32 = float32(s0f64)
	lbl2:
		l12 = s0f32
		s0i32 = l4
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 | s1i32
		l5 = s0i32
		s1f32 = l11
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l5
		s1f32 = l12
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	}
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 | s1i32
	l5 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = l14
	if s0f32 > s1f32 {
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
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l10 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0f32
		goto lbl5
	}
	s0i32 = l5
	s1i32 = 4
	s0i32 = s0i32 | s1i32
	l5 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l11 = s0f32
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 | s1i32
	s1f32 = l13
	s2f32 = l10
	s1f32 = s1f32 - s2f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 0.00024414062
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l11
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		l10 = s2f32
		s1f32 = s1f32 + s2f32
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		goto lbl7
	}
	s1f32 = l14
	s1f64 = float64(s1f32)
	s2f32 = l10
	s2f64 = float64(s2f32)
	l17 = s2f64
	s1f64 = s1f64 - s2f64
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	l10 = s2f32
	s2f64 = float64(s2f32)
	l18 = s2f64
	s3f32 = l11
	s3f64 = float64(s3f32)
	l16 = s3f64
	s2f64 = s2f64 - s3f64
	s1f64 = s1f64 * s2f64
	s2f32 = l13
	s2f64 = float64(s2f32)
	s3f64 = l17
	s2f64 = s2f64 - s3f64
	s1f64 = s1f64 / s2f64
	s2f64 = l16
	s1f64 = s1f64 + s2f64
	l17 = s1f64
	s2f64 = l18
	s2i64 = int64(math.Float64bits(s2f64))
	l8 = s2i64
	s3f64 = l16
	s3i64 = int64(math.Float64bits(s3f64))
	l9 = s3i64
	s4f32 = l10
	s5f32 = l11
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l0 = s4i32
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	s2f64 = math.Float64frombits(uint64(s2i64))
	l16 = s2f64
	if s1f64 < s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl9
	}
	s1f64 = l17
	l16 = s1f64
	s2i64 = l9
	s3i64 = l8
	s4i32 = l0
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	s2f64 = math.Float64frombits(uint64(s2i64))
	l17 = s2f64
	if s1f64 > s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl9
	}
	s1f64 = l17
	l16 = s1f64
lbl9:
	s1f64 = l16
	s1f32 = float32(s1f64)
lbl7:
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l5
	s1f32 = l14
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
lbl5:
	s0i32 = 1
	l5 = s0i32
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1f32 = l11
	s2f32 = l10
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	l0 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 | s1i32
	l6 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l15 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l13 = s1f32
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
		s0i32 = l4
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l4
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
		s0i32 = l4
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		goto lbl10
	}
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	l5 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 | s1i32
	l7 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l14 = s1f32
	if s0f32 >= s1f32 {
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
		s0i32 = l3
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l4
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l4
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
		s0i32 = 1
		l5 = s0i32
		s0i32 = l4
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		goto lbl10
	}
	s0f32 = l10
	s1f32 = l13
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
		s0i32 = l4
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l4
		s1i32 = l4
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 | s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l11 = s0f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		l10 = s1f32
		s0f32 = s0f32 - s1f32
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
			s0i32 = l4
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
			l10 = s0f32
			s1i32 = l4
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
			l12 = s1f32
			s0f32 = s0f32 + s1f32
			s1f32 = 0.5
			s0f32 = s0f32 * s1f32
			goto lbl16
		}
		s0f32 = l13
		s0f64 = float64(s0f32)
		s1f32 = l10
		s1f64 = float64(s1f32)
		l16 = s1f64
		s0f64 = s0f64 - s1f64
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		l12 = s1f32
		s1f64 = float64(s1f32)
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
		l10 = s2f32
		s2f64 = float64(s2f32)
		l17 = s2f64
		s1f64 = s1f64 - s2f64
		s0f64 = s0f64 * s1f64
		s1f32 = l11
		s1f64 = float64(s1f32)
		s2f64 = l16
		s1f64 = s1f64 - s2f64
		s0f64 = s0f64 / s1f64
		s1f64 = l17
		s0f64 = s0f64 + s1f64
		s0f32 = float32(s0f64)
	lbl16:
		l11 = s0f32
		s1f32 = l12
		s1i32 = int32(math.Float32bits(s1f32))
		l1 = s1i32
		s2f32 = l10
		s2i32 = int32(math.Float32bits(s2f32))
		l3 = s2i32
		s3f32 = l12
		s4f32 = l10
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l5 = s3i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1f32 = math.Float32frombits(uint32(s1i32))
		l10 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0f32 = l11
		l10 = s0f32
		s1i32 = l3
		s2i32 = l1
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1f32 = math.Float32frombits(uint32(s1i32))
		l11 = s1f32
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl15
		}
		s0f32 = l11
		l10 = s0f32
	lbl15:
		s0i32 = l4
		s1f32 = l10
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l4
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l4
		s1i32 = 8
		s0i32 = s0i32 | s1i32
		goto lbl13
	}
	s0i32 = l4
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
lbl13:
	l5 = s0i32
	s0f32 = l15
	s1f32 = l14
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
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l10 = s0f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		l11 = s1f32
		s0f32 = s0f32 - s1f32
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
			s0i32 = l4
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
			l11 = s0f32
			s1i32 = l4
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
			l12 = s1f32
			s0f32 = s0f32 + s1f32
			s1f32 = 0.5
			s0f32 = s0f32 * s1f32
			goto lbl21
		}
		s0f32 = l14
		s0f64 = float64(s0f32)
		s1f32 = l11
		s1f64 = float64(s1f32)
		l16 = s1f64
		s0f64 = s0f64 - s1f64
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		l12 = s1f32
		s1f64 = float64(s1f32)
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
		l11 = s2f32
		s2f64 = float64(s2f32)
		l17 = s2f64
		s1f64 = s1f64 - s2f64
		s0f64 = s0f64 * s1f64
		s1f32 = l10
		s1f64 = float64(s1f32)
		s2f64 = l16
		s1f64 = s1f64 - s2f64
		s0f64 = s0f64 / s1f64
		s1f64 = l17
		s0f64 = s0f64 + s1f64
		s0f32 = float32(s0f64)
	lbl21:
		l10 = s0f32
		s1f32 = l12
		s1i32 = int32(math.Float32bits(s1f32))
		l1 = s1i32
		s2f32 = l11
		s2i32 = int32(math.Float32bits(s2f32))
		l3 = s2i32
		s3f32 = l12
		s4f32 = l11
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l6 = s3i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1f32 = math.Float32frombits(uint32(s1i32))
		l11 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0f32 = l10
		l11 = s0f32
		s1i32 = l3
		s2i32 = l1
		s3i32 = l6
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1f32 = math.Float32frombits(uint32(s1i32))
		l10 = s1f32
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl20
		}
		s0f32 = l10
		l11 = s0f32
	lbl20:
		s0i32 = l5
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l5
		s1f32 = l11
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l5
		s1i32 = l4
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 | s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l5
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l5
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		goto lbl18
	}
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
lbl18:
	s0i32 = l5
	s1i32 = l4
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 3
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l5 = s0i32
	s0i32 = l4
	l1 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = 0
	l0 = s0i32
	s0i32 = l3
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl23:
	s0i32 = l2
	s1i32 = l5
	s2i32 = l0
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = l0
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l1 = s0i32
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl23
	}
	goto lbl0
lbl10:
	s0i32 = l2
	s1i32 = l1
	s2i32 = l5
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
lbl0:
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l5
	return s0i32
}
