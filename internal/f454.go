package internal

import (
	"unsafe"
)

func f454(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 112
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 3
	s3i32 = f105(ctx)
	s4i32 = 3
	s0i32 = f95(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l2 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l0
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	f162(ctx, s0i32, s1i32)
	s0i32 = l3
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
