package internal

import (
	"unsafe"
)

func f540(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
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
	s0i32 = ctx.g0
	s1i32 = 160
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = 16024
	s2i32 = 144
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 2147483647
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = 1
		l1 = s0i32
		s0i32 = l3
		s1i32 = 159
		s0i32 = s0i32 + s1i32
		l0 = s0i32
	}
	s0i32 = l3
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = -2
	s2i32 = l0
	s1i32 = s1i32 - s2i32
	l4 = s1i32
	s2i32 = l1
	s3i32 = l1
	s4i32 = l4
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s2i32 = l1
	s1i32 = s1i32 + s2i32
	l0 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = 14096
	s2i32 = l2
	s0i32 = f374(ctx, s0i32, s1i32, s2i32)
	l0 = s0i32
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l1 = s0i32
	s1i32 = l1
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 - s1i32
	s1i32 = 0
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	goto lbl0
lbl1:
	s0i32 = 29100
	s1i32 = 61
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = -1
	l0 = s0i32
lbl0:
	s0i32 = l3
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
