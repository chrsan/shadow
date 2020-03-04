package internal

import (
	"unsafe"
)

func f1146(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 3392
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+3384)])) = uint64(s1i64)
		s0i32 = l2
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+3376)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 3344
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = 4
		s1i32 = s1i32 | s2i32
		s2i32 = 3332
		s3i32 = 3332
		s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
		l3 = s0i32
		s0i32 = l2
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 0
		s3i32 = l1
		s0i32 = f182(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l2
		s1i32 = 3376
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		f420(ctx, s0i32, s1i32, s2i32)
		s0i32 = l3
		f43(ctx, s0i32)
	}
	s0i32 = l2
	s1i32 = 3392
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
