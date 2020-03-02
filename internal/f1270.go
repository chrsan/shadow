package internal

import (
	"unsafe"
)

func f1270(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l4
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l2
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
	lbl1:
		s0i32 = l5
		s1i32 = 16
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 16
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l1 = s0i32
		s0i32 = l3
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l2 = s0i32
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l2
			s2i32 = l1
			s0i32 = f27(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l1
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l4
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l5 = s0i32
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
