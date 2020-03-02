package internal

import (
	"unsafe"
)

func f1098(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 65535
	s0i32 = s0i32 & s1i32
	s1i32 = 8190
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l1
	f1099(ctx, s0i32, s1i32)
	s0i32 = l2
	s1i32 = l0
	f1860(ctx, s0i32, s1i32)
	s0i32 = 1
	l3 = s0i32
lbl0:
	s0i32 = l3
	return s0i32
}
