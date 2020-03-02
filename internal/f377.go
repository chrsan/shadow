package internal

import (
	"unsafe"
)

func f377(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 int32
	_ = l5
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
	s0i32 = ctx.g0
	s1i32 = 208
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)])) = uint32(s1i32)
	s0i32 = 0
	l2 = s0i32
	s0i32 = l5
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	s2i32 = 40
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+204)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint32(s1i32)
	s0i32 = 0
	s1i32 = l1
	s2i32 = l5
	s3i32 = 200
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 80
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = 160
	s4i32 = s4i32 + s5i32
	s5i32 = l3
	s6i32 = l4
	s0i32 = f376(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = -1
		l1 = s0i32
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	s1i32 = 0
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l2 = s0i32
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = l0
	s0i32 = int32(int8(ctx.Mem[int(s0i32+74)]))
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l6
		s2i32 = -33
		s1i32 = s1i32 & s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l6
	s1i32 = 32
	s0i32 = s0i32 & s1i32
	l7 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l5
		s3i32 = 200
		s2i32 = s2i32 + s3i32
		s3i32 = l5
		s4i32 = 80
		s3i32 = s3i32 + s4i32
		s4i32 = l5
		s5i32 = 160
		s4i32 = s4i32 + s5i32
		s5i32 = l3
		s6i32 = l4
		s0i32 = f376(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		goto lbl4
	}
	s0i32 = l0
	s1i32 = 80
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l5
	s2i32 = 80
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l6 = s0i32
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = l5
	s3i32 = 200
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 80
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = 160
	s4i32 = s4i32 + s5i32
	s5i32 = l3
	s6i32 = l4
	s0i32 = f376(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	l1 = s0i32
	s1i32 = l6
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl4
	}
	s0i32 = l0
	s1i32 = 0
	s2i32 = 0
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+36)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32) int32)(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l3 = s0i32
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = -1
	s2i32 = l3
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
lbl4:
	l1 = s0i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = l7
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = -1
	s1i32 = l1
	s2i32 = l0
	s3i32 = 32
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl0:
	s0i32 = l5
	s1i32 = 208
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l1
	return s0i32
}
