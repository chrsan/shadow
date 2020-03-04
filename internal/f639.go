package internal

import (
	"unsafe"
)

func f639(ctx *Context, l0 int32, l1 int32) int32 {
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
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+52)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+55)])
	l6 = s0i32
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+55)])
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l2 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l0 = s0i32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
			goto lbl1
		}
		s0i32 = l2
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		s0i32 = s0i32 - s1i32
		l2 = s0i32
		s1i32 = l2
		s2i32 = 31
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l2 = s1i32
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s0i32 = s0i32 ^ s1i32
		s1i32 = 255
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l1
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		goto lbl1
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l2 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	l4 = s1i32
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = l3
	s2i32 = 31
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l3 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s0i32 = s0i32 ^ s1i32
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l0 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l7 = s1i32
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = l5
	s2i32 = 31
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l5 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s0i32 = s0i32 ^ s1i32
	s1i32 = 255
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 2
		l8 = s0i32
		s0i32 = l3
		s1i32 = 256
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s1i32 = l4
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
			goto lbl1
		}
		s0i32 = l1
		s1i32 = l6
		ctx.Mem[int(s0i32+55)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		goto lbl1
	}
	s0i32 = l3
	s1i32 = 255
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l7
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		goto lbl1
	}
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l6
	ctx.Mem[int(s0i32+55)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
lbl1:
	s0i32 = 1
	l8 = s0i32
lbl0:
	s0i32 = l8
	return s0i32
}
