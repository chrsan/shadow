package internal

import (
	"unsafe"
)

func f929(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 7
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl2:
		s0i32 = l2
		l6 = s0i32
		s0i32 = l0
		l5 = s0i32
	lbl3:
		s0i32 = l5
		s1i32 = l1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+14)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+10)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+6)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l6
		l7 = s0i32
		s1i32 = -8
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l5
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l7
		s1i32 = 15
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l7
		s1i32 = 8
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl5:
			s0i32 = l5
			s1i32 = l1
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
			s0i32 = l5
			s1i32 = 2
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l6
			s1i32 = 1
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l7 = s0i32
			s0i32 = l6
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l6 = s0i32
			s0i32 = l7
			if s0i32 != 0 {
				goto lbl5
			}
		}
		s0i32 = l4
		l6 = s0i32
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l0
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l6
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		goto lbl0
		panic("unreachable executed")
		panic("unreachable executed")
	}
	s0i32 = l2
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl6:
	s0i32 = l2
	l5 = s0i32
	s0i32 = l0
	l6 = s0i32
lbl7:
	s0i32 = l6
	s1i32 = l1
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l6
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l5
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l5
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s1i32 = l3
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l6 = s0i32
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l6
	if s0i32 != 0 {
		goto lbl6
	}
lbl0:
}
