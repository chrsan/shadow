package internal

import (
	"unsafe"
)

func f1097(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
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
	s0i32 = l1
	s1i32 = 28
	s2i32 = 4
	s0i32 = f56(ctx, s0i32, s1i32, s2i32)
	l3 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
	s0i32 = l1
	s1i32 = l3
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 359
	s2i32 = l3
	s3i32 = l4
	s2i32 = s2i32 - s3i32
	f52(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l1 = s0i32
	s0i32 = l3
	s1i32 = 0
	ctx.Mem[int(s0i32+16)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		f334(ctx, s0i32, s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l2 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l1 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+84)])
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			s1i32 = l1
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
			s2i32 = l1
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
			s0i32 = f28(ctx, s0i32, s1i32, s2i32)
			l2 = s0i32
			s0i32 = l1
			s1i32 = 0
			ctx.Mem[int(s0i32+84)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = l2
			ctx.Mem[int(s0i32+85)] = uint8(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		} else {
			s0i32 = l2
		}
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		f2072(ctx, s0i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = 1
		ctx.Mem[int(s0i32+16)] = uint8(s1i32)
	}
}
