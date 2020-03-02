package internal

import (
	"unsafe"
)

func f1784(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s1i32 = l2
	s2i32 = 32768
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s3i32 = l5
	s4i32 = 8
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 255
	s3i32 = s3i32 & s4i32
	l1 = s3i32
	s4i32 = 255
	s3i32 = s3i32 ^ s4i32
	s4i32 = l4
	s3i32 = s3i32 * s4i32
	s4i32 = 6
	s3i32 = s3i32 >> (uint32(s4i32) & 31)
	s4i32 = l1
	s5i32 = l4
	s4i32 = s4i32 * s5i32
	s5i32 = 6
	s4i32 = s4i32 >> (uint32(s5i32) & 31)
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+36)]))
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
	s0i32 = l2
	s1i32 = l3
	s0i32 = s0i32 + s1i32
	return s0i32
}
