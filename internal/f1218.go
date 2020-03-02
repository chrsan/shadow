package internal

import (
	"math"
	"unsafe"
)

func f1218(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
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
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
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
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s9i32 int32
	_ = s9i32
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
	var s7f32 float32
	_ = s7f32
	var s8f32 float32
	_ = s8f32
	var s9f32 float32
	_ = s9f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s0f32 = float32(math.Floor(float64(s0f32)))
	l13 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l13
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
	l13 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l13
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
		goto lbl0
	}
	s0i32 = -2147483648
lbl0:
	l2 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s0f32 = float32(math.Ceil(float64(s0f32)))
	l13 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l13
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
	l13 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l13
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
	l10 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s0f32 = float32(math.Ceil(float64(s0f32)))
	l13 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l13
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
	l13 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l13
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
		goto lbl4
	}
	s0i32 = -2147483648
lbl4:
	l8 = s0i32
	s0i32 = l10
	s1i32 = l2
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0f32 = float32(math.Floor(float64(s0f32)))
	l13 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l13
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
	l13 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l13
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
		goto lbl6
	}
	s0i32 = -2147483648
lbl6:
	l7 = s0i32
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l0
	s1i32 = l8
	s2i32 = l7
	s1i32 = s1i32 - s2i32
	l9 = s1i32
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	s2i32 = 3
	s1i32 = s1i32 * s2i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+52)]))
	if int(s2i32) < 0 || int(s2i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s2i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s2i32].numParams != 2 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
	l5 = s0i32
	s1i32 = l9
	s2i32 = -2
	s1i32 = s1i32 + s2i32
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
	s0i32 = l5
	s1i32 = 1
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l5
	s1i32 = l9
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l11 = s1i32
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = 1
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l5
	s1i32 = l9
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l16 = s0f32
	s1i32 = l2
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l12 = s1i32
	s1f32 = float32(s1i32)
	s2i32 = l6
	s3i32 = 1
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	l13 = s0f32
	s1f32 = 255
	s0f32 = s0f32 * s1f32
	l14 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l14
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l14
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl9
	}
	s0i32 = 0
lbl9:
	l3 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l17 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l14 = s0f32
	s0i32 = l5
	s1i32 = l4
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = -1
	s2i32 = 0
	s3i32 = l3
	s4i32 = l3
	s5i32 = 8
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l3
	s4i32 = 247
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = -1
	s2i32 = 0
	s3i32 = l7
	s4i32 = 1
	s3i32 = s3i32 + s4i32
	s3f32 = float32(s3i32)
	s4f32 = l14
	s3f32 = s3f32 - s4f32
	l14 = s3f32
	s4f32 = l13
	s3f32 = s3f32 * s4f32
	s4f32 = 255
	s3f32 = s3f32 * s4f32
	l15 = s3f32
	s4f32 = 4.2949673e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4f32 = l15
	s5f32 = 0
	if s4f32 >= s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	s3i32 = s3i32 & s4i32
	if s3i32 != 0 {
		s3f32 = l15
		s3i32 = int32(uint32(math.Trunc(float64(s3f32))))
		goto lbl11
	}
	s3i32 = 0
lbl11:
	l3 = s3i32
	s4i32 = l3
	s5i32 = 8
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l3
	s4i32 = 247
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0f32 = l17
	s1i32 = l8
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = float32(s1i32)
	s0f32 = s0f32 - s1f32
	l15 = s0f32
	s1f32 = l13
	s0f32 = s0f32 * s1f32
	s1f32 = 255
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l13
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl13
	}
	s0i32 = 0
lbl13:
	l3 = s0i32
	s0i32 = l4
	s1i32 = l11
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = -1
	s2i32 = 0
	s3i32 = l3
	s4i32 = l3
	s5i32 = 8
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l3
	s4i32 = 247
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l7
	s2i32 = l2
	s3i32 = l4
	s4i32 = l5
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l6
	s1i32 = 3
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = -1
		s1i32 = 0
		s2f32 = l15
		s3f32 = 255
		s2f32 = s2f32 * s3f32
		l13 = s2f32
		s3f32 = 4.2949673e+09
		if s2f32 < s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s3f32 = l13
		s4f32 = 0
		if s3f32 >= s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		s2i32 = s2i32 & s3i32
		if s2i32 != 0 {
			s2f32 = l13
			s2i32 = int32(uint32(math.Trunc(float64(s2f32))))
			goto lbl17
		}
		s2i32 = 0
	lbl17:
		l2 = s2i32
		s3i32 = l2
		s4i32 = 8
		if uint32(s3i32) < uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l2
		s3i32 = 247
		if uint32(s2i32) > uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		l1 = s0i32
		s0i32 = l0
		s1i32 = l7
		s2i32 = l12
		s3i32 = l9
		s4i32 = -2
		s3i32 = s3i32 + s4i32
		s4i32 = l6
		s5i32 = -2
		s4i32 = s4i32 + s5i32
		s5i32 = -1
		s6i32 = 0
		s7f32 = l14
		s8f32 = 255
		s7f32 = s7f32 * s8f32
		l13 = s7f32
		s8f32 = 4.2949673e+09
		if s7f32 < s8f32 {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		s8f32 = l13
		s9f32 = 0
		if s8f32 >= s9f32 {
			s8i32 = 1
		} else {
			s8i32 = 0
		}
		s7i32 = s7i32 & s8i32
		if s7i32 != 0 {
			s7f32 = l13
			s7i32 = int32(uint32(math.Trunc(float64(s7f32))))
			goto lbl19
		}
		s7i32 = 0
	lbl19:
		l2 = s7i32
		s8i32 = l2
		s9i32 = 8
		if uint32(s8i32) < uint32(s9i32) {
			s8i32 = 1
		} else {
			s8i32 = 0
		}
		if s8i32 != 0 {
			// s6i32 = s6i32
		} else {
			s6i32 = s7i32
		}
		s7i32 = l2
		s8i32 = 247
		if uint32(s7i32) > uint32(s8i32) {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i32 = s5i32
		} else {
			s5i32 = s6i32
		}
		s6i32 = 255
		s5i32 = s5i32 & s6i32
		s6i32 = l1
		s7i32 = l0
		s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+0)]))
		s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+24)]))
		if int(s7i32) < 0 || int(s7i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s7i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s7i32].numParams != 7 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32, int32, int32))(table[s7i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		goto lbl15
	}
	s0i32 = l6
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
lbl15:
	s0i32 = l4
	s1i32 = -1
	s2i32 = 0
	s3f32 = l16
	s4i32 = l10
	s5i32 = -1
	s4i32 = s4i32 + s5i32
	l2 = s4i32
	s4f32 = float32(s4i32)
	s3f32 = s3f32 - s4f32
	l13 = s3f32
	s4f32 = 255
	s3f32 = s3f32 * s4f32
	l16 = s3f32
	s4f32 = 4.2949673e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4f32 = l16
	s5f32 = 0
	if s4f32 >= s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	s3i32 = s3i32 & s4i32
	if s3i32 != 0 {
		s3f32 = l16
		s3i32 = int32(uint32(math.Trunc(float64(s3f32))))
		goto lbl21
	}
	s3i32 = 0
lbl21:
	l1 = s3i32
	s4i32 = l1
	s5i32 = 8
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l1
	s4i32 = 247
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = -1
	s2i32 = 0
	s3f32 = l14
	s4f32 = l13
	s3f32 = s3f32 * s4f32
	s4f32 = 255
	s3f32 = s3f32 * s4f32
	l14 = s3f32
	s4f32 = 4.2949673e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4f32 = l14
	s5f32 = 0
	if s4f32 >= s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	s3i32 = s3i32 & s4i32
	if s3i32 != 0 {
		s3f32 = l14
		s3i32 = int32(uint32(math.Trunc(float64(s3f32))))
		goto lbl23
	}
	s3i32 = 0
lbl23:
	l1 = s3i32
	s4i32 = l1
	s5i32 = 8
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l1
	s4i32 = 247
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l8
	s1i32 = -1
	s2i32 = 0
	s3f32 = l15
	s4f32 = l13
	s3f32 = s3f32 * s4f32
	s4f32 = 255
	s3f32 = s3f32 * s4f32
	l13 = s3f32
	s4f32 = 4.2949673e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4f32 = l13
	s5f32 = 0
	if s4f32 >= s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	s3i32 = s3i32 & s4i32
	if s3i32 != 0 {
		s3f32 = l13
		s3i32 = int32(uint32(math.Trunc(float64(s3f32))))
		goto lbl25
	}
	s3i32 = 0
lbl25:
	l1 = s3i32
	s4i32 = l1
	s5i32 = 8
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l1
	s4i32 = 247
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l7
	s2i32 = l2
	s3i32 = l4
	s4i32 = l5
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
lbl8:
}
