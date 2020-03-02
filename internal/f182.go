package internal

import (
	"unsafe"
)

func f182(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
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
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+32)]))
	s4i32 = l2
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l5 = s2i32
	s3i32 = l3
	s4i32 = l0
	s5i32 = 3336
	s4i32 = s4i32 + s5i32
	l6 = s4i32
	s5i32 = 0
	s1i32 = f292(ctx, s1i32, s2i32, s3i32, s4i32, s5i32)
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l5
		s2i32 = l4
		s3i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])) = uint64(s3i64)
		s2i32 = l4
		s3i64 = 13195221663744
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)])) = uint64(s3i64)
		s2i32 = l4
		s3i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)])) = uint64(s3i64)
		s2i32 = l4
		s3i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)])) = uint64(s3i64)
		s2i32 = l4
		s3i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)])) = uint64(s3i64)
		s2i32 = l4
		s3i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)])) = uint64(s3i64)
		s2i32 = l4
		l1 = s2i32
		s3i32 = l6
		s4i32 = 1
		s0i32 = f292(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		l5 = s0i32
		s0i32 = l1
		s0i32 = f23(ctx, s0i32)
		s0i32 = l6
		s1i32 = 28
		s2i32 = 4
		s0i32 = f56(ctx, s0i32, s1i32, s2i32)
		l2 = s0i32
		s0i32 = l0
		s1i32 = 3340
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l3
		s1i32 = l2
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 304
		s2i32 = l2
		s3i32 = l1
		s2i32 = s2i32 - s3i32
		f52(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l2
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 21212
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l2
	return s0i32
}
