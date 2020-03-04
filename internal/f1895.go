package internal

import (
	"unsafe"
)

func f1895(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	f101(ctx, s0i32, s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	l6 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l4 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	l5 = s1i32
	s0i32 = s0i32 | s1i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 0
		s2i32 = l5
		s1i32 = s1i32 - s2i32
		s2i32 = 0
		s3i32 = l4
		s2i32 = s2i32 - s3i32
		s3i32 = l6
		f657(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l6
		l1 = s0i32
	}
	s0i32 = l0
	s1i32 = 1064
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l5 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l5
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s2i32 = 1032
		s1i32 = s1i32 + s2i32
		s1i32 = f94(ctx, s1i32)
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1064)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l4
		f226(ctx, s0i32, s1i32)
		s0i32 = l5
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1064)]))
		l4 = s0i32
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i32 = l1
	s2i32 = l2
	f663(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1064)]))
		s1i32 = l0
		s2i32 = 1068
		s1i32 = s1i32 + s2i32
		f273(ctx, s0i32, s1i32)
	}
	s0i32 = l6
	f40(ctx, s0i32)
	s0i32 = l3
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
