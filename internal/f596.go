package internal

import (
	"unsafe"
)

func f596(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l2 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l1
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s2i32 = 0
		s3i32 = l1
		s4i32 = 2
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		l0 = s3i32
		s1i32 = f27(ctx, s1i32, s2i32, s3i32)
		s2i32 = l0
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		return
	}
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l4 = s1i32
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l7 = s0i32
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = 1073741824
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l2 = s0i32
		s0i32 = l5
		s1i32 = l3
		s2i32 = l4
		s1i32 = s1i32 - s2i32
		l3 = s1i32
		s2i32 = 1
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l8 = s1i32
		s2i32 = l8
		s3i32 = l5
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = 1073741823
		s2i32 = l3
		s3i32 = 2
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s3i32 = 536870911
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l3 = s0i32
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = 1073741824
			if uint32(s0i32) >= uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl1
			}
			s0i32 = l3
			s1i32 = 2
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s0i32 = f17(ctx, s0i32)
			l2 = s0i32
		}
		s0i32 = l7
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		s2i32 = l1
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s0i32 = f27(ctx, s0i32, s1i32, s2i32)
		s0i32 = l6
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = l4
			s2i32 = l6
			s0i32 = f22(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l0
		s1i32 = l2
		s2i32 = l3
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		s2i32 = l5
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		if s0i32 != 0 {
			s0i32 = l4
			f12(ctx, s0i32)
		}
		return
	}
	f104(ctx)
	panic("unreachable executed")
lbl1:
	f63(ctx)
	panic("unreachable executed")
}
