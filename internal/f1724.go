package internal

import (
	"unsafe"
)

func f1724(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	s0i32 = f422(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l2 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+4)])
		l3 = s0i32
		s0i32 = l2
		s1i32 = 1
		ctx.Mem[int(s0i32+4)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l2
			f146(ctx, s0i32)
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+296)]))
		s2i32 = l1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+296)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+300)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l0 = s0i32
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
			s2i32 = l1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		}
		s0i32 = l2
		s1i32 = 0
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	}
}
