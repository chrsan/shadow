package internal

import (
	"math"
	"unsafe"
)

func f906(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
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
	s0i32 = l0
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2f32 = l2
	s2i32 = int32(math.Float32bits(s2f32))
	l0 = s2i32
	s3i32 = -2147483648
	s2i32 = s2i32 & s3i32
	s3i32 = l0
	s4i32 = 2147483647
	s3i32 = s3i32 & s4i32
	s3f32 = math.Float32frombits(uint32(s3i32))
	l2 = s3f32
	s4f32 = 0.07739938
	s3f32 = s3f32 * s4f32
	s4f32 = l2
	s5f32 = l2
	s4f32 = s4f32 * s5f32
	s5f32 = l2
	s6f32 = 0.3
	s5f32 = s5f32 * s6f32
	s6f32 = 0.6975
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = 0.0025
	s4f32 = s4f32 + s5f32
	s5f32 = l2
	s6f32 = 0.055
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
	s3i32 = int32(math.Float32bits(s3f32))
	s2i32 = s2i32 | s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	s3f32 = l3
	s3i32 = int32(math.Float32bits(s3f32))
	l0 = s3i32
	s4i32 = -2147483648
	s3i32 = s3i32 & s4i32
	s4i32 = l0
	s5i32 = 2147483647
	s4i32 = s4i32 & s5i32
	s4f32 = math.Float32frombits(uint32(s4i32))
	l2 = s4f32
	s5f32 = 0.07739938
	s4f32 = s4f32 * s5f32
	s5f32 = l2
	s6f32 = l2
	s5f32 = s5f32 * s6f32
	s6f32 = l2
	s7f32 = 0.3
	s6f32 = s6f32 * s7f32
	s7f32 = 0.6975
	s6f32 = s6f32 + s7f32
	s5f32 = s5f32 * s6f32
	s6f32 = 0.0025
	s5f32 = s5f32 + s6f32
	s6f32 = l2
	s7f32 = 0.055
	if s6f32 < s7f32 {
		s6i32 = 1
	} else {
		s6i32 = 0
	}
	if s6i32 != 0 {
		// s4f32 = s4f32
	} else {
		s4f32 = s5f32
	}
	s4i32 = int32(math.Float32bits(s4f32))
	s3i32 = s3i32 | s4i32
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4f32 = l4
	s4i32 = int32(math.Float32bits(s4f32))
	l0 = s4i32
	s5i32 = -2147483648
	s4i32 = s4i32 & s5i32
	s5i32 = l0
	s6i32 = 2147483647
	s5i32 = s5i32 & s6i32
	s5f32 = math.Float32frombits(uint32(s5i32))
	l2 = s5f32
	s6f32 = 0.07739938
	s5f32 = s5f32 * s6f32
	s6f32 = l2
	s7f32 = l2
	s6f32 = s6f32 * s7f32
	s7f32 = l2
	s8f32 = 0.3
	s7f32 = s7f32 * s8f32
	s8f32 = 0.6975
	s7f32 = s7f32 + s8f32
	s6f32 = s6f32 * s7f32
	s7f32 = 0.0025
	s6f32 = s6f32 + s7f32
	s7f32 = l2
	s8f32 = 0.055
	if s7f32 < s8f32 {
		s7i32 = 1
	} else {
		s7i32 = 0
	}
	if s7i32 != 0 {
		// s5f32 = s5f32
	} else {
		s5f32 = s6f32
	}
	s5i32 = int32(math.Float32bits(s5f32))
	s4i32 = s4i32 | s5i32
	s4f32 = math.Float32frombits(uint32(s4i32))
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
