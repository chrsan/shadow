package internal

import (
	"unsafe"
)

func f1246(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l2
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	l2 = s2i32
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l0 = s0i32
	s1i32 = -16777217
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = l1
	s1i32 = l1
	s2i32 = l3
	s3i32 = l0
	f366(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l5
	if s0i32 != 0 {
		goto lbl2
	}
lbl1:
	return
lbl0:
	s0i32 = l1
	s1i32 = l0
	s2i32 = l3
	s3i32 = l2
	s4i32 = l4
	s5i32 = 22132
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
}
