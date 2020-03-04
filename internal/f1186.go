package internal

import (
	"unsafe"
)

func f1186(ctx *Context, l0 int32, l1 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l1
	s1i32 = 2
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = -13
		s1i32 = s1i32 & s2i32
		s2i32 = l1
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = 12
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	}
}
