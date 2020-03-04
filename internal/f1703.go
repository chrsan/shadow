package internal

import (
	"unsafe"
)

func f1703(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
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
	s0i32 = ctx.g0
	s1i32 = 1056
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 28864
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1052)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1048)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = 1024
	s2i32 = l1
	s0i32 = f540(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l1 = s0i32
		s0i32 = 13954
		s1i32 = 0
		f19(ctx, s0i32, s1i32)
		goto lbl2
	}
	s0i32 = l1
	s1i32 = 1024
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s1i32 = 0
	s2i32 = l1
	f620(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 28864
		l5 = s0i32
		goto lbl5
	}
	s0i32 = l4
	s1i32 = -10
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l4
	s2i32 = 9
	s1i32 = s1i32 + s2i32
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = -4
	s0i32 = s0i32 & s1i32
	s0i32 = f17(ctx, s0i32)
	l5 = s0i32
	s1i32 = 0
	ctx.Mem[int(s0i32+8)] = uint8(s1i32)
	s0i32 = l5
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	s1i32 = l4
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
lbl5:
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	s1i32 = 28864
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	f12(ctx, s0i32)
lbl4:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+1048)]))
	s0i32 = f540(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = l3
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = l0
	s1i32 = l3
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	f620(ctx, s0i32, s1i32, s2i32)
lbl1:
	s0i32 = l3
	s1i32 = 1056
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	return
lbl0:
	f621(ctx)
	panic("unreachable executed")
}
