package internal

import (
	"unsafe"
)

func f1885(ctx *Context, l0 int64) {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	ctx.g0 = s0i32
	s0i64 = l0
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i64 = l0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		f1884(ctx, s0i32)
	}
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
