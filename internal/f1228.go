package internal

import (
	"unsafe"
)

func f1228(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	f515(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l0
	s1i32 = 20376
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 2
	s1i32 = f44(ctx, s1i32, s2i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 216
	s2i32 = 217
	s3i32 = l3
	s4i32 = l3
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 1 {
		panic("argument count mismatch")
	}
	s3i32 = (*(*func(*Context, int32) int32)(table[s4i32].f()))(ctx, s3i32)
	s4i32 = 1
	s3i32 = s3i32 & s4i32
	s4i32 = l2
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+44)]))
	s5i32 = 261120
	s4i32 = s4i32 & s5i32
	s5i32 = 1024
	if s4i32 == s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	s3i32 = s3i32 | s4i32
	l1 = s3i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 218
	s2i32 = 219
	s3i32 = l1
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
}
