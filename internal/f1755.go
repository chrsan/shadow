package internal

import (
	"unsafe"
)

func f1755(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s2i64 int64
	_ = s2i64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 36
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
		s1i32 = l1
		s2i32 = 0
		ctx.Mem[int(s1i32+32)] = uint8(s2i32)
		s1i32 = l1
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint64(s2i64)
		s1i32 = l1
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
		s1i32 = l1
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
		s1i32 = l1
		s2i32 = l2
		s0i32 = f342(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l1
			s0i32 = f41(ctx, s0i32)
			goto lbl0
		}
		s0i32 = l0
		s1i32 = l1
		f253(ctx, s0i32, s1i32)
		s0i32 = l1
		s0i32 = f41(ctx, s0i32)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = 36
	s1i32 = s1i32 + s2i32
	f253(ctx, s0i32, s1i32)
lbl0:
	s0i32 = l3
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
