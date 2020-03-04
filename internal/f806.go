package internal

import (
	"math"
	"unsafe"
)

func f806(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
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
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s9i32 int32
	_ = s9i32
	var s10i32 int32
	_ = s10i32
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
	var s9f32 float32
	_ = s9f32
	var s10f32 float32
	_ = s10f32
	var s11f32 float32
	_ = s11f32
	var s12f32 float32
	_ = s12f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2f32 = 1
	s3f32 = l5
	s2f32 = s2f32 - s3f32
	l10 = s2f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	l7 = s3f32
	s2f32 = s2f32 * s3f32
	s3f32 = 1
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+24)]))
	l8 = s4f32
	s3f32 = s3f32 - s4f32
	l11 = s3f32
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l7
	s4f32 = l2
	s5f32 = l2
	s4f32 = s4f32 + s5f32
	l12 = s4f32
	s5f32 = l5
	s4f32 = s4f32 - s5f32
	l9 = s4f32
	s5f32 = 1
	s6f32 = l7
	s7f32 = l8
	s6f32 = s6f32 / s7f32
	s7f32 = 0
	s8f32 = l8
	s9f32 = 0
	if s8f32 > s9f32 {
		s8i32 = 1
	} else {
		s8i32 = 0
	}
	l6 = s8i32
	if s8i32 != 0 {
		// s6f32 = s6f32
	} else {
		s6f32 = s7f32
	}
	l2 = s6f32
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l5
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 * s4f32
	s4f32 = l7
	s5f32 = l5
	s4f32 = s4f32 * s5f32
	s5f32 = l9
	s6f32 = l8
	s5f32 = s5f32 * s6f32
	s6f32 = l2
	s7f32 = 7
	s6f32 = s6f32 * s7f32
	s7f32 = l2
	s8f32 = -1
	s7f32 = s7f32 + s8f32
	s8f32 = l2
	s9f32 = l2
	s8f32 = s8f32 + s9f32
	l9 = s8f32
	s9f32 = l9
	s8f32 = s8f32 + s9f32
	l9 = s8f32
	s9f32 = l9
	s10f32 = l9
	s9f32 = s9f32 * s10f32
	s8f32 = s8f32 + s9f32
	s7f32 = s7f32 * s8f32
	s6f32 = s6f32 + s7f32
	s7f32 = 1
	s8f32 = 1
	s9f32 = l2
	s9f32 = float32(math.Sqrt(float64(s9f32)))
	s8f32 = s8f32 / s9f32
	s7f32 = s7f32 / s8f32
	s8f32 = l2
	s7f32 = s7f32 - s8f32
	s8f32 = l7
	s9f32 = l7
	s8f32 = s8f32 + s9f32
	l2 = s8f32
	s9f32 = l2
	s8f32 = s8f32 + s9f32
	s9f32 = l8
	if s8f32 <= s9f32 {
		s8i32 = 1
	} else {
		s8i32 = 0
	}
	if s8i32 != 0 {
		// s6f32 = s6f32
	} else {
		s6f32 = s7f32
	}
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s5f32 = l12
	s6f32 = l5
	if s5f32 <= s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	s2f32 = s2f32 + s3f32
	s3f32 = l10
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
	l7 = s4f32
	s3f32 = s3f32 * s4f32
	s4f32 = l11
	s5f32 = l3
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l7
	s5f32 = l3
	s6f32 = l3
	s5f32 = s5f32 + s6f32
	l9 = s5f32
	s6f32 = l5
	s5f32 = s5f32 - s6f32
	l3 = s5f32
	s6f32 = 1
	s7f32 = l7
	s8f32 = l8
	s7f32 = s7f32 / s8f32
	s8f32 = 0
	s9i32 = l6
	if s9i32 != 0 {
		// s7f32 = s7f32
	} else {
		s7f32 = s8f32
	}
	l2 = s7f32
	s6f32 = s6f32 - s7f32
	s5f32 = s5f32 * s6f32
	s6f32 = l5
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l7
	s6f32 = l5
	s5f32 = s5f32 * s6f32
	s6f32 = l3
	s7f32 = l8
	s6f32 = s6f32 * s7f32
	s7f32 = l2
	s8f32 = 7
	s7f32 = s7f32 * s8f32
	s8f32 = l2
	s9f32 = -1
	s8f32 = s8f32 + s9f32
	s9f32 = l2
	s10f32 = l2
	s9f32 = s9f32 + s10f32
	l3 = s9f32
	s10f32 = l3
	s9f32 = s9f32 + s10f32
	l3 = s9f32
	s10f32 = l3
	s11f32 = l3
	s10f32 = s10f32 * s11f32
	s9f32 = s9f32 + s10f32
	s8f32 = s8f32 * s9f32
	s7f32 = s7f32 + s8f32
	s8f32 = 1
	s9f32 = 1
	s10f32 = l2
	s10f32 = float32(math.Sqrt(float64(s10f32)))
	s9f32 = s9f32 / s10f32
	s8f32 = s8f32 / s9f32
	s9f32 = l2
	s8f32 = s8f32 - s9f32
	s9f32 = l7
	s10f32 = l7
	s9f32 = s9f32 + s10f32
	l2 = s9f32
	s10f32 = l2
	s9f32 = s9f32 + s10f32
	s10f32 = l8
	if s9f32 <= s10f32 {
		s9i32 = 1
	} else {
		s9i32 = 0
	}
	if s9i32 != 0 {
		// s7f32 = s7f32
	} else {
		s7f32 = s8f32
	}
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 + s6f32
	s6f32 = l9
	s7f32 = l5
	if s6f32 <= s7f32 {
		s6i32 = 1
	} else {
		s6i32 = 0
	}
	if s6i32 != 0 {
		// s4f32 = s4f32
	} else {
		s4f32 = s5f32
	}
	s3f32 = s3f32 + s4f32
	s4f32 = l11
	s5f32 = l4
	s4f32 = s4f32 * s5f32
	s5f32 = l10
	s6i32 = l0
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+20)]))
	l3 = s6f32
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s5f32 = l3
	s6f32 = l4
	s7f32 = l4
	s6f32 = s6f32 + s7f32
	l7 = s6f32
	s7f32 = l5
	s6f32 = s6f32 - s7f32
	l4 = s6f32
	s7f32 = 1
	s8f32 = l3
	s9f32 = l8
	s8f32 = s8f32 / s9f32
	s9f32 = 0
	s10i32 = l6
	if s10i32 != 0 {
		// s8f32 = s8f32
	} else {
		s8f32 = s9f32
	}
	l2 = s8f32
	s7f32 = s7f32 - s8f32
	s6f32 = s6f32 * s7f32
	s7f32 = l5
	s6f32 = s6f32 + s7f32
	s5f32 = s5f32 * s6f32
	s6f32 = l3
	s7f32 = l5
	s6f32 = s6f32 * s7f32
	s7f32 = l4
	s8f32 = l8
	s7f32 = s7f32 * s8f32
	s8f32 = l2
	s9f32 = 7
	s8f32 = s8f32 * s9f32
	s9f32 = l2
	s10f32 = -1
	s9f32 = s9f32 + s10f32
	s10f32 = l2
	s11f32 = l2
	s10f32 = s10f32 + s11f32
	l4 = s10f32
	s11f32 = l4
	s10f32 = s10f32 + s11f32
	l4 = s10f32
	s11f32 = l4
	s12f32 = l4
	s11f32 = s11f32 * s12f32
	s10f32 = s10f32 + s11f32
	s9f32 = s9f32 * s10f32
	s8f32 = s8f32 + s9f32
	s9f32 = 1
	s10f32 = 1
	s11f32 = l2
	s11f32 = float32(math.Sqrt(float64(s11f32)))
	s10f32 = s10f32 / s11f32
	s9f32 = s9f32 / s10f32
	s10f32 = l2
	s9f32 = s9f32 - s10f32
	s10f32 = l3
	s11f32 = l3
	s10f32 = s10f32 + s11f32
	l2 = s10f32
	s11f32 = l2
	s10f32 = s10f32 + s11f32
	s11f32 = l8
	if s10f32 <= s11f32 {
		s10i32 = 1
	} else {
		s10i32 = 0
	}
	if s10i32 != 0 {
		// s8f32 = s8f32
	} else {
		s8f32 = s9f32
	}
	s7f32 = s7f32 * s8f32
	s6f32 = s6f32 + s7f32
	s7f32 = l7
	s8f32 = l5
	if s7f32 <= s8f32 {
		s7i32 = 1
	} else {
		s7i32 = 0
	}
	if s7i32 != 0 {
		// s5f32 = s5f32
	} else {
		s5f32 = s6f32
	}
	s4f32 = s4f32 + s5f32
	s5f32 = l10
	s6f32 = l8
	s5f32 = s5f32 * s6f32
	s6f32 = l5
	s5f32 = s5f32 + s6f32
	s6i32 = l1
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, float32, float32, float32, float32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2f32, s3f32, s4f32, s5f32)
}
