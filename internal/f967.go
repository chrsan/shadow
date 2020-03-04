package internal

import (
	"unsafe"
)

func f967(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = f2070(ctx, s0i32)
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+100)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l2 = s0i32
		s0i32 = l0
		s1i32 = 0
		ctx.Mem[int(s0i32+100)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l2
		ctx.Mem[int(s0i32+101)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		ctx.g0 = s0i32
		return
	}
	s0i32 = l1
	s1i32 = 1347
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 334
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1264
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = 1024
	s1i32 = l1
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl0:
	s0i32 = l1
	s1i32 = 1371
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 374
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1264
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 1024
	s1i32 = l1
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
