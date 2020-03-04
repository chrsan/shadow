package internal

import (
	"unsafe"
)

func f1711(ctx *Context, l0 int32, l1 int32) {
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
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
	var s3i64 int64
	_ = s3i64
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 4575657221408423936
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = 16
	s2i32 = s2i32 + s3i32
	l0 = s2i32
	s3i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])) = uint64(s3i64)
	s2i32 = l0
	s3i64 = 13195221663744
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)])) = uint64(s3i64)
	s2i32 = l0
	s3i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)])) = uint64(s3i64)
	s2i32 = l0
	s3i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)])) = uint64(s3i64)
	s2i32 = l0
	s3i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)])) = uint64(s3i64)
	s2i32 = l0
	s3i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)])) = uint64(s3i64)
	s2i32 = l0
	s3i32 = l2
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s3i32 = f212(ctx, s3i32)
	s4i32 = 3
	s5i32 = 21916
	f401(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	s0i32 = l0
	s0i32 = f23(ctx, s0i32)
	s0i32 = l2
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
}
