package internal

import (
	"unsafe"
)

func f589(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l1
	s1i32 = 28
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 28928
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l2 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl2
	case 1:
		goto lbl1
	default:
		goto lbl3
	}
lbl3:
	s0i32 = l1
	s1i32 = 28928
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l2
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l2 = s1i32
	s2i32 = 1
	s3i32 = l2
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l2
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s1i32 = 22092
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
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
	l2 = s0i32
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = 12
	s0i32 = f17(ctx, s0i32)
	l2 = s0i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 26016
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl4:
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 28960
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 28928
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	goto lbl1
lbl2:
lbl6:
	s0i32 = l1
	s1i32 = 28928
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
lbl1:
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 28960
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
lbl0:
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
