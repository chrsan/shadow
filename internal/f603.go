package internal

import (
	"unsafe"
)

func f603(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l1
		s2i32 = 24
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 14
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = 14
			s2i32 = -1
			s3i32 = -1
			s4i32 = -1
			s5i32 = l3
			s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
			s6i32 = l2
			s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
			s6i32 = 0
			s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
			return s0i32
		}
		s0i32 = l0
		s1i32 = 31
		s2i32 = l1
		s3i32 = -1
		s4i32 = -1
		s5i32 = l2
		s6i32 = 0
		s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	} else {
		s0i32 = l1
	}
	return s0i32
}
