package internal

import (
	"unsafe"
)

func f1349(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
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
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = f164(ctx, s0i32)
		return s0i32
	}
	s0i32 = l1
	s1i32 = -64
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 29100
		s1i32 = 48
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = 0
		return s0i32
	}
	s0i32 = l0
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	s1i32 = 16
	s2i32 = l1
	s3i32 = 11
	s2i32 = s2i32 + s3i32
	s3i32 = -8
	s2i32 = s2i32 & s3i32
	s3i32 = l1
	s4i32 = 11
	if uint32(s3i32) < uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = f1348(ctx, s0i32, s1i32)
	l2 = s0i32
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		return s0i32
	}
	s0i32 = l1
	s0i32 = f164(ctx, s0i32)
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = l2
	s1i32 = l0
	s2i32 = l0
	s3i32 = -4
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l3 = s2i32
	s3i32 = -8
	s2i32 = s2i32 & s3i32
	s3i32 = 4
	s4i32 = 8
	s5i32 = l3
	s6i32 = 3
	s5i32 = s5i32 & s6i32
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	s2i32 = s2i32 - s3i32
	l3 = s2i32
	s3i32 = l1
	s4i32 = l3
	s5i32 = l1
	if uint32(s4i32) < uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	f12(ctx, s0i32)
	s0i32 = l2
	return s0i32
}
