package internal

import (
	"unsafe"
)

func f1880(ctx *Context, l0 int32, l1 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s1i32 = 1064
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s2i32 = l0
		s3i32 = 20
		s2i32 = s2i32 + s3i32
		s3i32 = l0
		s3i32 = int32(ctx.Mem[int(s3i32+40)])
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s0i32 = f81(ctx, s0i32, s1i32)
		return
	}
	s0i32 = l1
	s1i32 = l0
	f223(ctx, s0i32, s1i32)
}
