package internal

import (
	"unsafe"
)

func f1286(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s0i32 = f38(ctx, s0i32)
	l4 = s0i32
	s1i32 = l1
	f331(ctx, s0i32, s1i32)
	s0i32 = l4
	s1i32 = l2
	f331(ctx, s0i32, s1i32)
	s0i32 = l4
	s1i32 = l4
	s1i32 = int32(ctx.Mem[int(s1i32+10)])
	s2i32 = 248
	s1i32 = s1i32 & s2i32
	s2i32 = 5
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+10)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l3
	s3i32 = 1
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
	s0i32 = l4
	f34(ctx, s0i32)
	s0i32 = l5
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
