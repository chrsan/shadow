package internal

import (
	"unsafe"
)

func f592(ctx *Context, l0 int32, l1 int32) int32 {
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
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l1
	s0i32 = f1626(ctx, s0i32, s1i32)
	l7 = s0i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	l7 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	l4 = s0i32
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l5 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l9 = s1i32
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	l10 = s1i32
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	l3 = s1i32
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = -4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l9
	s1i32 = -17
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 536870911
	s0i32 = s0i32 & s1i32
	s1i32 = l5
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 1073741823
	s0i32 = s0i32 & s1i32
	s1i32 = l4
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s1i32 = l10
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l8
	s1i32 = l6
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l8
	s1i32 = 3
	s0i32 = s0i32 + s1i32
	s1i32 = -4
	s0i32 = s0i32 & s1i32
	l7 = s0i32
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	l6 = s0i32
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l6
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	s2i32 = 8
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 5
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l2
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	f82(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l2
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	f82(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l2
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	f82(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l2
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	f82(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = 3
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l1
		f82(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l4
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = l1
		f82(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		s2i32 = l3
		f82(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l2
	f673(ctx, s0i32)
lbl0:
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l7
	return s0i32
}
