package internal

import (
	"math"
	"unsafe"
)

func f1281(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
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
	s0i32 = ctx.g0
	s1i32 = 224
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = l3
	s0i32 = f469(ctx, s0i32, s1i32, s2i32)
	l2 = s0i32
	s0i32 = l5
	s1i32 = 0
	ctx.Mem[int(s0i32+75)] = uint8(s1i32)
	s0i32 = l5
	s1i64 = 4294967297
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 12884901894
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l5
	s2i32 = l5
	s3i32 = 80
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 75
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = 76
	s4i32 = s4i32 + s5i32
	s0i32 = f158(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s1f32 = float32(s1i32)
	l9 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s1f32 = float32(s1i32)
	l10 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1f32 = float32(s1i32)
	l7 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1f32 = float32(s1i32)
	l8 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = s1f32
	s0i32 = l5
	s0i32 = int32(ctx.Mem[int(s0i32+75)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l10
		s1f32 = l8
		s0f32 = s0f32 - s1f32
		s1f32 = 1
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl3
		}
		s0f32 = l9
		s1f32 = l7
		s0f32 = s0f32 - s1f32
		s1f32 = 1
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl3
		}
		s0f32 = l7
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l7
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl5
		}
		s0i32 = -2147483648
	lbl5:
		l3 = s0i32
		s0i32 = l1
		s1i32 = l5
		s2i32 = 48
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s3i32 = 76
		s2i32 = s2i32 + s3i32
		s3i32 = 4
		s4f32 = l8
		s4f32 = float32(math.Abs(float64(s4f32)))
		s5f32 = 2.1474836e+09
		if s4f32 < s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			s4f32 = l8
			s4i32 = int32(math.Trunc(float64(s4f32)))
			goto lbl7
		}
		s4i32 = -2147483648
	lbl7:
		s5i32 = l3
		s6i32 = 0
		s7i32 = l1
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
		s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32, int32, int32) int32)(table[s7i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		s1i32 = 261120
		s0i32 = s0i32 & s1i32
		s1i32 = 3072
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l5
	s1i32 = l4
	s0i32 = f46(ctx, s0i32, s1i32)
	l3 = s0i32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	s2f32 = 255
	s1f32 = s1f32 * s2f32
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l7 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l7
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
	l7 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l7
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
	l7 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l7
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl10
	}
	s1i32 = -2147483648
lbl10:
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+76)]))
	l6 = s2i32
	s3i32 = 24
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 * s2i32
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = -16777216
	s1i32 = s1i32 & s2i32
	s2i32 = l6
	s3i32 = 16777215
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	f196(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = l5
	s2i32 = 80
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+80)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s0i32 = f23(ctx, s0i32)
	goto lbl2
lbl3:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l5
	s3i32 = 96
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 80
	s3i32 = s3i32 + s4i32
	s4i32 = l4
	s5i32 = 0
	s6i32 = l0
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+124)]))
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32, int32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
lbl2:
	s0i32 = l2
	s1i32 = l5
	s2i32 = l5
	s3i32 = 80
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 75
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = 76
	s4i32 = s4i32 + s5i32
	s0i32 = f158(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	f12(ctx, s0i32)
lbl0:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+92)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+76)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+44)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+28)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		f13(ctx, s0i32)
	}
	s0i32 = l5
	s1i32 = 224
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
