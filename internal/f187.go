package internal

import (
	"unsafe"
)

func f187(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = 30736
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s1i32 = l0
	s2i32 = 3
	s1i32 = s1i32 + s2i32
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 29100
		s1i32 = 48
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = -1
		return s0i32
	}
	s0i32 = l0
	s1i32 = int32(len(ctx.Mem) / 65536)
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0i32 = f1346(ctx, s0i32)
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 29100
	s1i32 = 48
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = -1
	return s0i32
lbl1:
	s0i32 = 30736
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	return s0i32
}
