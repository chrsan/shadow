package internal

import (
	"unsafe"
)

func f1100(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+24)])
	l3 = s0i32
	s1i32 = 5
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l2 = s0i32
		s0i32 = 1
		l0 = s0i32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl1
		case 1:
			goto lbl1
		case 2:
			goto lbl2
		case 3:
			goto lbl4
		case 4:
			goto lbl1
		default:
			goto lbl5
		}
	lbl5:
		s0i32 = l2
		s1i32 = 7
		s0i32 = s0i32 + s1i32
		s1i32 = 3
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		goto lbl0
	lbl4:
		s0i32 = 2
		l0 = s0i32
		goto lbl1
	}
	s0i32 = l1
	s1i32 = 60
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 3323
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 3282
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 3256
	s1i32 = l1
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl2:
	s0i32 = 4
	l0 = s0i32
lbl1:
	s0i32 = l0
	s1i32 = l2
	s0i32 = s0i32 * s1i32
lbl0:
	l0 = s0i32
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
