package internal

import (
	"unsafe"
)

func f1095(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s0i32 = f37(ctx, s0i32)
		l4 = s0i32
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s2i32 = l4
		s0i32 = f421(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l1
			s2i32 = l4
			f1097(ctx, s0i32, s1i32, s2i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l3 = s0i32
			goto lbl1
		}
		s0i32 = l1
		s1i32 = 28
		s2i32 = 4
		s0i32 = f56(ctx, s0i32, s1i32, s2i32)
		l3 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l2 = s0i32
		s0i32 = l1
		s1i32 = l3
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 359
		s2i32 = l3
		s3i32 = l2
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
		s0i32 = l3
		s1i32 = 0
		ctx.Mem[int(s0i32+16)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	lbl1:
		s0i32 = l3
		s0i32 = int32(ctx.Mem[int(s0i32+16)])
		l0 = s0i32
		s0i32 = l4
		f34(ctx, s0i32)
		s0i32 = l0
		s1i32 = 0
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l3 = s0i32
	}
	s0i32 = l5
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l3
	return s0i32
}
