package internal

import (
	"unsafe"
)

func f1827(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1096)]))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1092)]))
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1096)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
	s3i32 = l2
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+24)]))
	s3i32 = s3i32 - s4i32
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s1i32 = s1i32 - s2i32
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1092)])) = uint32(s1i32)
lbl0:
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	l0 = s1i32
	s2i32 = l0
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
}
