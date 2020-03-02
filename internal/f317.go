package internal

import (
	"unsafe"
)

func f317(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+8)])
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = f177(ctx, s0i32)
		l3 = s0i32
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = int32(ctx.Mem[int(s0i32+90)])
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		s3i32 = l3
		s4i32 = 1
		if s3i32 != s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		s0i32 = f1830(ctx, s0i32, s1i32, s2i32, s3i32)
		return s0i32
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = 1
	if s3i32 != s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s0i32 = f1829(ctx, s0i32, s1i32, s2i32, s3i32)
	return s0i32
}
