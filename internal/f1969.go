package internal

import (
	"unsafe"
)

func f1969(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l4
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l2
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])))
		s2i32 = l1
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)])))
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l2
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+2)])))
		s2i32 = l1
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+2)])))
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l2
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		s2i32 = l1
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l2
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+6)])))
		s2i32 = l1
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+6)])))
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+6)])) = uint16(s1i32)
		s0i32 = l2
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
