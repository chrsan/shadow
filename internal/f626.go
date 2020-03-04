package internal

import (
	"math"
	"unsafe"
)

func f626(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
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
	s0i32 = ctx.g0
	s1i32 = 4144
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l15 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l16 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l12 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l17 = s0f32
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l13 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l18 = s2f32
	s3f32 = 0.6666667
	s2f32 = s2f32 * s3f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
	l19 = s3f32
	s4f32 = 0.33333334
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4124)])) = s1f32
	s0i32 = l4
	s1f32 = l12
	s2f32 = l10
	s3f32 = 0.6666667
	s2f32 = s2f32 * s3f32
	s3f32 = l17
	s4f32 = 0.33333334
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 - s2f32
	l14 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4128)])) = s1f32
	s0i32 = l4
	s1f32 = 0
	s2f32 = l14
	s1f32 = s1f32 - s2f32
	l20 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l4
	s1f32 = 0
	s2f32 = l11
	s1f32 = s1f32 - s2f32
	l21 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4140)])) = s1f32
	s0i32 = l4
	s1i32 = 4128
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2f32 = l14
	s3f32 = l20
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s0i32 = l4
	s1i32 = l4
	s2i32 = 4124
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = 4140
	s2i32 = s2i32 + s3i32
	s3f32 = l11
	s4f32 = l21
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
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4132)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4136)])) = uint32(s1i32)
	s0i32 = l4
	s1f32 = l15
	s2f32 = l18
	s3f32 = 0.33333334
	s2f32 = s2f32 * s3f32
	s3f32 = l19
	s4f32 = 0.6666667
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4108)])) = s1f32
	s0i32 = l4
	s1f32 = l16
	s2f32 = l10
	s3f32 = 0.33333334
	s2f32 = s2f32 * s3f32
	s3f32 = l17
	s4f32 = 0.6666667
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 - s2f32
	l10 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4112)])) = s1f32
	s0i32 = l4
	s1f32 = 0
	s2f32 = l10
	s1f32 = s1f32 - s2f32
	l14 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l4
	s1f32 = 0
	s2f32 = l11
	s1f32 = s1f32 - s2f32
	l18 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4140)])) = s1f32
	s0i32 = l4
	s1i32 = 4112
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2f32 = l10
	s3f32 = l14
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = l4
	s1i32 = l4
	s2i32 = 4108
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = 4140
	s2i32 = s2i32 + s3i32
	s3f32 = l11
	s4f32 = l18
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
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l8 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4116)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4120)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 4132
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 4116
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s2f32 = math.Float32frombits(uint32(s2i32))
	s3i32 = l8
	s3f32 = math.Float32frombits(uint32(s3i32))
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0f32
	s1i32 = l4
	s2i32 = 4136
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = 4120
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4i32 = l6
	s4f32 = math.Float32frombits(uint32(s4i32))
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
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l11 = s1f32
	s2f32 = l11
	s3f32 = l10
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
	l10 = s0f32
	s1f32 = 0.125
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
		s0i32 = 1
		l5 = s0i32
		s0f32 = l10
		s1f32 = 0.5
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 2
		l5 = s0i32
		s0f32 = l10
		s1f32 = 2
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 3
		l5 = s0i32
		s0f32 = l10
		s1f32 = 8
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 4
		l5 = s0i32
		s0f32 = l10
		s1f32 = 32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 5
		l5 = s0i32
		s0f32 = l10
		s1f32 = 128
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 6
		l5 = s0i32
		s0f32 = l10
		s1f32 = 512
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 7
		l5 = s0i32
		s0f32 = l10
		s1f32 = 2048
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 8
		l5 = s0i32
		s0i32 = 512
		s1f32 = l10
		s2f32 = 8192
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 != 0 {
			goto lbl2
		}
	lbl3:
		s0i32 = 1
		s1i32 = l5
		s0i32 = s0i32 << (uint32(s1i32) & 31)
	lbl2:
		l6 = s0i32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l10 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l11 = s0f32
		s0i32 = l4
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0f32 = 1
		s1i32 = l6
		s1f32 = float32(s1i32)
		s0f32 = s0f32 / s1f32
		l14 = s0f32
		s0f32 = l13
		s1f32 = l11
		s0f32 = s0f32 - s1f32
		s1f32 = 3
		s0f32 = s0f32 * s1f32
		l18 = s0f32
		s0f32 = l12
		s1f32 = l10
		s0f32 = s0f32 - s1f32
		s1f32 = 3
		s0f32 = s0f32 * s1f32
		l20 = s0f32
		s0f32 = l11
		s1f32 = l15
		s2f32 = l13
		s3f32 = l13
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 3
		s0f32 = s0f32 * s1f32
		l21 = s0f32
		s0f32 = l10
		s1f32 = l16
		s2f32 = l12
		s3f32 = l12
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 3
		s0f32 = s0f32 * s1f32
		l22 = s0f32
		s0f32 = l19
		s1f32 = l13
		s2f32 = l15
		s1f32 = s1f32 - s2f32
		s2f32 = 3
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l11
		s0f32 = s0f32 - s1f32
		l15 = s0f32
		s0f32 = l17
		s1f32 = l12
		s2f32 = l16
		s1f32 = s1f32 - s2f32
		s2f32 = 3
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l10
		s0f32 = s0f32 - s1f32
		l16 = s0f32
		s0i32 = -1
		l7 = s0i32
		s0f32 = 0
		l12 = s0f32
		s0i32 = 1
		l5 = s0i32
		s0f32 = 0
		l13 = s0f32
		s0i32 = -1
		l8 = s0i32
	lbl4:
		s0i32 = l4
		s1i32 = l5
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s1f32 = l11
		s2f32 = l14
		s3f32 = l12
		s2f32 = s2f32 + s3f32
		l12 = s2f32
		s3f32 = l18
		s4f32 = l12
		s5f32 = l21
		s6f32 = l15
		s7f32 = l12
		s6f32 = s6f32 * s7f32
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 + s4f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		l17 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l9
		s1f32 = l10
		s2f32 = l14
		s3f32 = l13
		s2f32 = s2f32 + s3f32
		l13 = s2f32
		s3f32 = l20
		s4f32 = l13
		s5f32 = l22
		s6f32 = l16
		s7f32 = l13
		s6f32 = s6f32 * s7f32
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 + s4f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		l19 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = 0
		s1i32 = l7
		s2f32 = l17
		s2i32 = int32(math.Float32bits(s2f32))
		s3i32 = 2139095040
		s2i32 = s2i32 & s3i32
		s3i32 = 2139095040
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = 0
		s1i32 = l8
		s2f32 = l19
		s2i32 = int32(math.Float32bits(s2f32))
		s3i32 = 2139095040
		s2i32 = s2i32 & s3i32
		s3i32 = 2139095040
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l8 = s0i32
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
			goto lbl4
		}
		s0i32 = l8
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l7
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l4
		s1i32 = l6
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l6
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s3i32 = l2
		s4i32 = l3
		if int(s4i32) < 0 || int(s4i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s4i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s4i32].numParams != 4 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
		goto lbl0
	}
	s0i32 = l4
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 2
	s2i32 = l1
	s3i32 = l2
	s4i32 = l3
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 4 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
lbl0:
	s0i32 = l4
	s1i32 = 4144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
