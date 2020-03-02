package internal

import (
	"math"
	"unsafe"
)

func f743(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
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
	s0f32 = l2
	s1f32 = 0.33333334
	s0f32 = s0f32 + s1f32
	l6 = s0f32
	s1f32 = l6
	s1f32 = float32(math.Floor(float64(s1f32)))
	s0f32 = s0f32 - s1f32
	s1f32 = 6
	s0f32 = s0f32 * s1f32
	s1f32 = -3
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = -1
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1f32 = 1
	s0f32 = f14(ctx, s0f32, s1f32)
	l6 = s0f32
	s0f32 = l2
	s1f32 = 0.6666667
	s0f32 = s0f32 + s1f32
	l7 = s0f32
	s1f32 = l7
	s1f32 = float32(math.Floor(float64(s1f32)))
	s0f32 = s0f32 - s1f32
	s1f32 = 6
	s0f32 = s0f32 * s1f32
	s1f32 = -3
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = -1
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1f32 = 1
	s0f32 = f14(ctx, s0f32, s1f32)
	l7 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2f32 = 1
	s3f32 = l4
	s4f32 = l4
	s3f32 = s3f32 + s4f32
	s4f32 = -1
	s3f32 = s3f32 + s4f32
	s3f32 = float32(math.Abs(float64(s3f32)))
	s2f32 = s2f32 - s3f32
	s3f32 = l3
	s2f32 = s2f32 * s3f32
	l3 = s2f32
	s3f32 = l2
	s4f32 = 0
	s3f32 = s3f32 + s4f32
	l2 = s3f32
	s4f32 = l2
	s4f32 = float32(math.Floor(float64(s4f32)))
	s3f32 = s3f32 - s4f32
	s4f32 = 6
	s3f32 = s3f32 * s4f32
	s4f32 = -3
	s3f32 = s3f32 + s4f32
	s3f32 = float32(math.Abs(float64(s3f32)))
	s4f32 = -1
	s3f32 = s3f32 + s4f32
	s4f32 = 0
	s3f32 = f15(ctx, s3f32, s4f32)
	s4f32 = 1
	s3f32 = f14(ctx, s3f32, s4f32)
	s4f32 = -0.5
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = l4
	s2f32 = s2f32 + s3f32
	s3f32 = l3
	s4f32 = l7
	s5f32 = -0.5
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 * s4f32
	s4f32 = l4
	s3f32 = s3f32 + s4f32
	s4f32 = l3
	s5f32 = l6
	s6f32 = -0.5
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l4
	s4f32 = s4f32 + s5f32
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
