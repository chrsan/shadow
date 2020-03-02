package internal

import (
	"unsafe"
)

func f594(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
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
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l1
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 0
	s3i32 = l3
	s4i32 = 4
	s3i32 = s3i32 & s4i32
	s4i32 = 2
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	l8 = s3i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 0
	s3i32 = l3
	s4i32 = 2
	s3i32 = s3i32 & s4i32
	s4i32 = 1
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	l9 = s3i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l10 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 0
	s2i32 = l3
	s3i32 = 1
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 - s2i32
	s2i32 = l7
	s1i32 = s1i32 & s2i32
	l11 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 268435455
	s0i32 = s0i32 & s1i32
	s1i32 = l1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l8
	s0i32 = s0i32 & s1i32
	s1i32 = l1
	s2i32 = 1073741823
	s1i32 = s1i32 & s2i32
	s2i32 = l1
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l9
	s1i32 = s1i32 & s2i32
	s2i32 = l1
	s3i32 = 536870911
	s2i32 = s2i32 & s3i32
	s3i32 = l1
	if s2i32 != s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s1i32 = s1i32 | s2i32
	s0i32 = s0i32 | s1i32
	s1i32 = l6
	s2i32 = l2
	if uint32(s1i32) < uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 | s1i32
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	l8 = s0i32
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	l9 = s0i32
	s0i32 = l4
	s1i32 = l5
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s2i32 = l5
	s1i32 = s1i32 + s2i32
	l12 = s1i32
	s2i32 = l10
	s1i32 = s1i32 + s2i32
	l13 = s1i32
	s2i32 = l11
	s1i32 = s1i32 + s2i32
	l14 = s1i32
	s2i32 = l7
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s2i32 = -64
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = -65
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l8
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l6
	s1i32 = l5
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l12
	s1i32 = l5
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l13
	s1i32 = l10
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l14
	s1i32 = l11
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = l7
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	goto lbl0
lbl1:
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
lbl0:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = l9
	if s3i32 == 0 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4i32 = l4
	s5i32 = 8
	s4i32 = s4i32 + s5i32
	f593(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
