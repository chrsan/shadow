package internal

import (
	"math"
	"unsafe"
)

func f750(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
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
	var s6i32 int32
	_ = s6i32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2f32 = l2
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s3f32 = l3
	s4f32 = l3
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s2f32 = float32(math.Sqrt(float64(s2f32)))
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = -64
	s3i32 = s3i32 - s4i32
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 - s3f32
	s3f32 = l3
	s4f32 = l4
	s5f32 = l5
	s6i32 = l1
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
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
