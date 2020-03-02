package internal

import (
	"unsafe"
)

func f1435(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
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
	s0i32 = ctx.g0
	s1i32 = 112
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 2
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	s4i32 = 2
	s0i32 = f95(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	l1 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	l3 = s0i32
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+8)])
	l4 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+3)])
	l5 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+4)])
	l0 = s0i32
	s0i32 = l2
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = l3
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 | s1i32
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 | s1i32
	s1i32 = l5
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 | s1i32
	s1i32 = l0
	s2i32 = 4
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 | s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	return s0i32
}
