package internal

import (
	"unsafe"
)

func f383(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32 {
	var l14 int32
	_ = l14
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
	var s11i32 int32
	_ = s11i32
	var s12i32 int32
	_ = s12i32
	var s13i32 int32
	_ = s13i32
	var s14i32 int32
	_ = s14i32
	var s2f32 float32
	_ = s2f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	s5i32 = l5
	s6i32 = l6
	s7i32 = l7
	s8i32 = l8
	s9i32 = l9
	s10i32 = l10
	s11i32 = l11
	s12i32 = l12
	s13i32 = l13
	s14i32 = l0
	s14i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s14i32+0)]))
	s14i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s14i32+68)]))
	if int(s14i32) < 0 || int(s14i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s14i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s14i32].numParams != 14 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(table[s14i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32, s10i32, s11i32, s12i32, s13i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 1
	l14 = s0i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l13
	s1i32 = l1
	s2f32 = 1
	s1i32 = f20(ctx, s1i32, s2f32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl0:
	s0i32 = l14
	return s0i32
}
