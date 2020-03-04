package internal

import (
	"math"
	"unsafe"
)

func f758(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
	var l6 int32
	_ = l6
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
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l6 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3i32 = -1
	s2i32 = s2i32 + s3i32
	s2f32 = float32(uint32(s2i32))
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	l3 = s2f32
	s3f32 = 4.2949673e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3f32 = l3
	s4f32 = 0
	if s3f32 >= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		s2f32 = l3
		s2i32 = int32(uint32(math.Trunc(float64(s2f32))))
		goto lbl0
	}
	s2i32 = 0
lbl0:
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	l0 = s2i32
	s3i32 = l6
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s2i32 = s2i32 + s3i32
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s3i32 = l6
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
	s4i32 = l0
	s3i32 = s3i32 + s4i32
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s2f32 = s2f32 + s3f32
	s3i32 = l6
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s4i32 = l0
	s3i32 = s3i32 + s4i32
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s4i32 = l6
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+24)]))
	s5i32 = l0
	s4i32 = s4i32 + s5i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s3f32 = s3f32 + s4f32
	s4i32 = l6
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s5i32 = l0
	s4i32 = s4i32 + s5i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s5f32 = l2
	s4f32 = s4f32 * s5f32
	s5i32 = l6
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+28)]))
	s6i32 = l0
	s5i32 = s5i32 + s6i32
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s4f32 = s4f32 + s5f32
	s5i32 = l6
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
	s6i32 = l0
	s5i32 = s5i32 + s6i32
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s6f32 = l2
	s5f32 = s5f32 * s6f32
	s6i32 = l6
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+32)]))
	s7i32 = l0
	s6i32 = s6i32 + s7i32
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
	s5f32 = s5f32 + s6f32
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
