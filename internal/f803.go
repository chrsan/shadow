package internal

import (
	"unsafe"
)

func f803(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
	var l6 int32
	_ = l6
	var l7 int32
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
	var s9f32 float32
	_ = s9f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l19 = s0f32
	s1f32 = l5
	s0f32 = s0f32 * s1f32
	l8 = s0f32
	s1f32 = l2
	s2f32 = 0.3
	s1f32 = s1f32 * s2f32
	s2f32 = l3
	s3f32 = 0.59
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s3f32 = 0.11
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	l14 = s2f32
	s1f32 = s1f32 * s2f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l13 = s2f32
	s3f32 = l5
	s2f32 = s2f32 * s3f32
	l11 = s2f32
	s3f32 = 0.3
	s2f32 = s2f32 * s3f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	l20 = s3f32
	s4f32 = l5
	s3f32 = s3f32 * s4f32
	l9 = s3f32
	s4f32 = 0.59
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l8
	s4f32 = 0.11
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	s0f32 = s0f32 + s1f32
	l10 = s0f32
	s1f32 = 0.11
	s0f32 = s0f32 * s1f32
	s1f32 = l11
	s2f32 = l8
	s1f32 = s1f32 + s2f32
	l11 = s1f32
	s2f32 = 0.3
	s1f32 = s1f32 * s2f32
	s2f32 = l9
	s3f32 = l8
	s2f32 = s2f32 + s3f32
	l9 = s2f32
	s3f32 = 0.59
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	l8 = s0f32
	s1f32 = l14
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	l16 = s1f32
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	l17 = s1f32
	s2f32 = l10
	s3f32 = l8
	s4f32 = l8
	s5f32 = l10
	s6f32 = l8
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l8
	s6f32 = l11
	s7f32 = l9
	s8f32 = l10
	s7f32 = f14(ctx, s7f32, s8f32)
	s6f32 = f14(ctx, s6f32, s7f32)
	l12 = s6f32
	s5f32 = s5f32 - s6f32
	l18 = s5f32
	s4f32 = s4f32 / s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l12
	s5f32 = 0
	if s4f32 >= s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l6 = s4i32
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l12 = s2f32
	s3f32 = l8
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l11
	s3f32 = l9
	s4f32 = l10
	s3f32 = f15(ctx, s3f32, s4f32)
	s2f32 = f15(ctx, s2f32, s3f32)
	l15 = s2f32
	s3f32 = l8
	s2f32 = s2f32 - s3f32
	l10 = s2f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l12
	s2f32 = l15
	s3f32 = l16
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l7 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	l12 = s0f32
	s0f32 = l8
	s1f32 = l17
	s2f32 = l9
	s3f32 = l8
	s4f32 = l8
	s5f32 = l9
	s6f32 = l8
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l18
	s4f32 = s4f32 / s5f32
	s3f32 = s3f32 + s4f32
	s4i32 = l6
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l9 = s2f32
	s3f32 = l8
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l10
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l9
	s2i32 = l7
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	l15 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2f32 = l13
	s3f32 = 1
	s4f32 = l5
	s3f32 = s3f32 - s4f32
	l9 = s3f32
	s2f32 = s2f32 * s3f32
	s3f32 = 1
	s4f32 = l14
	s3f32 = s3f32 - s4f32
	l13 = s3f32
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l8
	s4f32 = l17
	s5f32 = l11
	s6f32 = l8
	s7f32 = l8
	s8f32 = l11
	s9f32 = l8
	s8f32 = s8f32 - s9f32
	s7f32 = s7f32 * s8f32
	s8f32 = l18
	s7f32 = s7f32 / s8f32
	s6f32 = s6f32 + s7f32
	s7i32 = l6
	if s7i32 != 0 {
		// s5f32 = s5f32
	} else {
		s5f32 = s6f32
	}
	l2 = s5f32
	s6f32 = l8
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l10
	s4f32 = s4f32 / s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l2
	s5i32 = l7
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	s4f32 = 0
	s3f32 = f15(ctx, s3f32, s4f32)
	s2f32 = s2f32 + s3f32
	s3f32 = l15
	s4f32 = l9
	s5f32 = l20
	s4f32 = s4f32 * s5f32
	s5f32 = l13
	s6f32 = l3
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l12
	s5f32 = l9
	s6f32 = l19
	s5f32 = s5f32 * s6f32
	s6f32 = l13
	s7f32 = l4
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 + s5f32
	s5f32 = l14
	s6f32 = l5
	s5f32 = s5f32 + s6f32
	s6f32 = l16
	s5f32 = s5f32 - s6f32
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
