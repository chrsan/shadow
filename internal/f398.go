package internal

import (
	"unsafe"
)

func f398(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l1 = s1i32
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 3
	s1i32 = s1i32 + s2i32
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	s3i32 = 3
	s2i32 = s2i32 & s3i32
	s3i32 = 14024
	s2i32 = s2i32 + s3i32
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	s1i32 = s1i32 * s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = 36
	s0i32 = s0i32 + s1i32
	return s0i32
}
