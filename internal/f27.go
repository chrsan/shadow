package internal

import (
	"unsafe"
)

func f27(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
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
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = 3
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = -2
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	s0i32 = l3
	s1i32 = -3
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+2)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = 7
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+3)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = 9
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 0
	s2i32 = l0
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l4 = s1i32
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l1
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = 16843009
	s1i32 = s1i32 * s2i32
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l2
	s2i32 = l4
	s1i32 = s1i32 - s2i32
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	l4 = s1i32
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 9
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = -12
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 25
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = -16
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = -20
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = -24
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = -28
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l3
	s2i32 = 4
	s1i32 = s1i32 & s2i32
	s2i32 = 24
	s1i32 = s1i32 | s2i32
	l4 = s1i32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s1i32 = 32
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i64 = int64(uint32(s0i32))
	l5 = s0i64
	s1i64 = 32
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	s1i64 = l5
	s0i64 = s0i64 | s1i64
	l5 = s0i64
	s0i32 = l3
	s1i32 = l4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
lbl1:
	s0i32 = l1
	s1i64 = l5
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l1
	s1i64 = l5
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l1
	s1i64 = l5
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s1i64 = l5
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l2
	s1i32 = -32
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 31
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl0:
	s0i32 = l0
	return s0i32
}
