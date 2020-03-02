package internal

import (
	"unsafe"
)

func f17(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = 1
	s2i32 = l0
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l0 = s0i32
lbl0:
	s0i32 = l0
	s0i32 = f164(ctx, s0i32)
	l1 = s0i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 30220
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l2
	if int(s0i32) < 0 || int(s0i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s0i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s0i32].numParams != 0 {
		panic("argument count mismatch")
	}
	(*(*func(*Context))(table[s0i32].f()))(ctx)
	goto lbl0
lbl1:
	s0i32 = l1
	return s0i32
}
