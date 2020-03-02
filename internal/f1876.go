package internal

import (
	"unsafe"
)

func f1876(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s1i32 = l1
	f1091(ctx, s0i32, s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l0
	s2i32 = 124
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	if int(s2i32) < 0 || int(s2i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s2i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s2i32].numParams != 2 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32))(table[s2i32].f()))(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = 188
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
	l1 = s1i32
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+54)])))
	s2i32 = 16
	s1i32 = s1i32 & s2i32
	s2i32 = 4
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = l1
	s2i32 = f1858(ctx, s2i32)
	f1075(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+220)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+216)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+212)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+228)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 236
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	s2i32 = 0
	s3i32 = 1024
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
}
