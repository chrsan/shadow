package internal

import (
	"unsafe"
)

func f1382(ctx *Context, l0 int32, l1 int64, l2 int32) int64 {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = 0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
	s2i64 = l1
	s3i32 = l2
	s4i32 = 255
	s3i32 = s3i32 & s4i32
	s4i32 = l3
	s5i32 = 8
	s4i32 = s4i32 + s5i32
	s1i32 = f4(ctx, s1i32, s2i64, s3i32, s4i32)
	l0 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s0i32 = 29100
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = -1
lbl1:
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		goto lbl0
	}
	s0i32 = l3
	s1i64 = -1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i64 = -1
lbl0:
	l1 = s0i64
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i64 = l1
	return s0i64
}
