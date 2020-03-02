package internal

import (
	"math"
	"unsafe"
)

func f171(ctx *Context, l0 int32, l1 int32, l2 float32, l3 int32, l4 float32, l5 int32, l6 float32) {
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
	var s8f32 float32
	_ = s8f32
	s0i32 = l3
	s1f32 = l4
	s2i32 = l1
	s2f32 = float32(uint32(s2i32))
	s3f32 = 255
	s2f32 = s2f32 / s3f32
	l4 = s2f32
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0f32 = (*(*func(*Context, int32, float32, float32) float32)(table[s3i32].f()))(ctx, s0i32, s1f32, s2f32)
	l9 = s0f32
	s0i32 = l5
	s1f32 = l6
	s2f32 = 1
	s3f32 = l4
	s2f32 = s2f32 - s3f32
	l7 = s2f32
	s3i32 = l5
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0f32 = (*(*func(*Context, int32, float32, float32) float32)(table[s3i32].f()))(ctx, s0i32, s1f32, s2f32)
	l10 = s0f32
	s1f32 = l2
	s0f32 = s0f32 * s1f32
	l8 = s0f32
	s0i32 = 0
	l3 = s0i32
	s0f32 = 0
	l2 = s0f32
	s0f32 = l4
	s1f32 = l7
	s0f32 = s0f32 - s1f32
	l11 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00390625
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
	lbl2:
		s0i32 = l0
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		s1f32 = l2
		s2f32 = 255
		s1f32 = s1f32 / s2f32
		l4 = s1f32
		s2f32 = l4
		s3f32 = l8
		s4f32 = 1
		s5f32 = l4
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = 255
		s1f32 = s1f32 * s2f32
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		l4 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l4
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
		l4 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l4
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
		l4 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l4
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl3
		}
		s1i32 = -2147483648
	lbl3:
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0f32 = l2
		s1f32 = 1
		s0f32 = s0f32 + s1f32
		l2 = s0f32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = 256
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		goto lbl0
		panic("unreachable executed")
		panic("unreachable executed")
	}
lbl5:
	s0i32 = l0
	s1i32 = l3
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2f32 = l6
	s3f32 = l9
	s4f32 = l2
	s5f32 = 255
	s4f32 = s4f32 / s5f32
	l4 = s4f32
	s5f32 = l4
	s6f32 = l8
	s7f32 = 1
	s8f32 = l4
	s7f32 = s7f32 - s8f32
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	l4 = s4f32
	s3f32 = s3f32 * s4f32
	s4f32 = l10
	s5f32 = 1
	s6f32 = l4
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4i32 = l5
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s1f32 = (*(*func(*Context, int32, float32, float32) float32)(table[s4i32].f()))(ctx, s1i32, s2f32, s3f32)
	s2f32 = l7
	s1f32 = s1f32 - s2f32
	s2f32 = l11
	s1f32 = s1f32 / s2f32
	s2f32 = 255
	s1f32 = s1f32 * s2f32
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l4 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l4
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
	l4 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l4
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
	l4 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l4
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl6
	}
	s1i32 = -2147483648
lbl6:
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0f32 = l2
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	l2 = s0f32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = 256
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl0:
}
