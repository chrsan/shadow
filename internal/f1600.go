package internal

import (
	"unsafe"
)

func f1600(ctx *Context, l0 int32, l1 int32) {
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
	var l11 int32
	_ = l11
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
	s0i32 = l1
	s1i32 = 14204
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l5 = s0i32
	s0i32 = -1
	s1i32 = 0
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s0i32 = f539(ctx, s0i32)
		l1 = s0i32
	}
	s0i32 = l1
	s1i32 = 5
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = l1
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = -4
	s0i32 = s0i32 & s1i32
	l9 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l10 = s1i32
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l2 = s1i32
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l11 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l11
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l6 = s0i32
	}
	s0i32 = l0
	s1i32 = l2
	s2i32 = 1
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s2i32 = l4
	s3i32 = l2
	s4i32 = l4
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = 4096
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l3 = s0i32
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l2
		s0i32 = f29(ctx, s0i32, s1i32)
		l3 = s0i32
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l2 = s0i32
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l2
		f13(ctx, s0i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l3 = s0i32
	}
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
lbl1:
	s0i32 = l0
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l7
	s1i32 = l9
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = -4
		s0i32 = s0i32 & s1i32
		s1i32 = l0
		s0i32 = s0i32 + s1i32
		s1i32 = -4
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = l1
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
}
