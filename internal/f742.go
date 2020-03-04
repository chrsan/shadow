package internal

import (
	"unsafe"
)

func f742(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
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
	s0f32 = l2
	s1f32 = l3
	s2f32 = l4
	s1f32 = f14(ctx, s1f32, s2f32)
	s0f32 = f14(ctx, s0f32, s1f32)
	l7 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2f32 = l2
	s3f32 = l3
	s2f32 = s2f32 - s3f32
	s3f32 = 1
	s4f32 = l2
	s5f32 = l3
	s6f32 = l4
	s5f32 = f15(ctx, s5f32, s6f32)
	s4f32 = f15(ctx, s4f32, s5f32)
	l6 = s4f32
	s5f32 = l7
	s4f32 = s4f32 - s5f32
	l9 = s4f32
	s3f32 = s3f32 / s4f32
	l8 = s3f32
	s2f32 = s2f32 * s3f32
	s3f32 = 4
	s2f32 = s2f32 + s3f32
	s3f32 = l4
	s4f32 = l2
	s3f32 = s3f32 - s4f32
	s4f32 = l8
	s3f32 = s3f32 * s4f32
	s4f32 = 2
	s3f32 = s3f32 + s4f32
	s4f32 = l6
	s5f32 = l3
	if s4f32 != s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	s3f32 = 6
	s4f32 = 0
	s5f32 = l3
	s6f32 = l4
	if s5f32 < s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	s4f32 = l3
	s5f32 = l4
	s4f32 = s4f32 - s5f32
	s5f32 = l8
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l6
	s5f32 = l2
	if s4f32 != s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	s3f32 = 0.16666667
	s2f32 = s2f32 * s3f32
	s3f32 = 0
	s4f32 = l6
	s5f32 = l7
	if s4f32 != s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l0 = s4i32
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	s3f32 = l9
	s4f32 = 2
	s5f32 = l6
	s4f32 = s4f32 - s5f32
	s5f32 = l7
	s4f32 = s4f32 - s5f32
	s5f32 = l6
	s6f32 = l7
	s5f32 = s5f32 + s6f32
	l2 = s5f32
	s6f32 = l2
	s7f32 = 0.5
	s6f32 = s6f32 * s7f32
	l2 = s6f32
	s7f32 = 0.5
	if s6f32 > s7f32 {
		s6i32 = 1
	} else {
		s6i32 = 0
	}
	if s6i32 != 0 {
		// s4f32 = s4f32
	} else {
		s4f32 = s5f32
	}
	s3f32 = s3f32 / s4f32
	s4f32 = 0
	s5i32 = l0
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	s4f32 = l2
	s5f32 = l5
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
