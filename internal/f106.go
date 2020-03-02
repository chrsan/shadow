package internal

import (
	"unsafe"
)

func f106(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	s0i32 = l0
	s1i32 = l1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l1
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	return s0i32
lbl1:
	s0i32 = l0
	s1i32 = l1
	s0i32 = s0i32 ^ s1i32
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l3 = s0i32
	s0i32 = l0
	s1i32 = l1
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		if s0i32 != 0 {
			s0i32 = l0
			l3 = s0i32
			goto lbl3
		}
		s0i32 = l0
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			l3 = s0i32
			goto lbl4
		}
		s0i32 = l0
		l3 = s0i32
	lbl8:
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l3
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l2
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			goto lbl8
		}
		goto lbl4
	}
	s0i32 = l3
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l4
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
	lbl11:
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l0
		s1i32 = l2
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = l1
		s2i32 = l2
		s1i32 = s1i32 + s2i32
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l3
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			goto lbl11
		}
	}
	s0i32 = l2
	s1i32 = 3
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
lbl12:
	s0i32 = l0
	s1i32 = l2
	s2i32 = -4
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
lbl9:
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl13:
	s0i32 = l0
	s1i32 = l2
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l2
	if s0i32 != 0 {
		goto lbl13
	}
	goto lbl0
lbl4:
	s0i32 = l2
	s1i32 = 3
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l2
	l4 = s0i32
lbl14:
	s0i32 = l3
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l3
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l4
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l2
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l2 = s0i32
lbl3:
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl15:
	s0i32 = l3
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l2
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	if s0i32 != 0 {
		goto lbl15
	}
lbl0:
	s0i32 = l0
	return s0i32
}
