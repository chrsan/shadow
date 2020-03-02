package internal

import (
	"unsafe"
)

func f998(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 int32
	_ = l5
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
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l3
	s5i32 = 20
	s4i32 = s4i32 + s5i32
	s5i32 = l3
	s5i32 = int32(ctx.Mem[int(s5i32+40)])
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	s4i32 = l5
	s5i32 = 8
	s4i32 = s4i32 + s5i32
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+48)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s2i32 = 36
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s3i32 = 52
		s2i32 = s2i32 + s3i32
		s3i32 = 1
		s4i32 = l3
		s5i32 = l4
		f478(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		s0i32 = 1
		l6 = s0i32
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 0
		f91(ctx, s0i32, s1i32)
		goto lbl1
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	f13(ctx, s0i32)
lbl1:
	s0i32 = l5
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
	s0i32 = l6
	return s0i32
}
