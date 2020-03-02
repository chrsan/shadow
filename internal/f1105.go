package internal

import (
	"unsafe"
)

func f1105(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 24
	s2i32 = 0
	f16(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+137)])
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 132
		s2i32 = l0
		s3i32 = 96
		s2i32 = s2i32 + s3i32
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 131
	s2i32 = l0
	s3i32 = 48
	s2i32 = s2i32 + s3i32
	f16(ctx, s0i32, s1i32, s2i32)
	s0i32 = 1
	return s0i32
}
