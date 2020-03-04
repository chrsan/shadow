package internal

import (
	"unsafe"
)

func f545(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = int32(int8(ctx.Mem[int(s0i32+75)]))
		s1i32 = l0
		s2i32 = 255
		s1i32 = s1i32 & s2i32
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l2 = s0i32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		s1i32 = l2
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		return
	lbl1:
		s0i32 = l1
		s1i32 = l0
		f377(ctx, s0i32, s1i32)
		return
	}
	s0i32 = l1
	s0i32 = int32(int8(ctx.Mem[int(s0i32+75)]))
	s1i32 = l0
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l2 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l1
	s1i32 = l2
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l0
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	goto lbl2
lbl3:
	s0i32 = l1
	s1i32 = l0
	f377(ctx, s0i32, s1i32)
lbl2:
}
