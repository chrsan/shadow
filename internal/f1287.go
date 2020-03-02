package internal

import (
	"unsafe"
)

func f1287(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 int32, l5 int32) {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
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
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l6
	s0i32 = f38(ctx, s0i32)
	l7 = s0i32
	s1i32 = l1
	s2f32 = l2
	s3f32 = l3
	s4i32 = l4
	s5i32 = l5
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s6i32 = l5
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+44)]))
	s7i32 = 192
	s6i32 = s6i32 & s7i32
	s5i32 = s5i32 | s6i32
	if s5i32 == 0 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	f2086(ctx, s0i32, s1i32, s2f32, s3f32, s4i32, s5i32)
	s0i32 = l0
	s1i32 = l7
	s2i32 = l5
	s3i32 = 0
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+104)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 4 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l7
	f34(ctx, s0i32)
	s0i32 = l6
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
