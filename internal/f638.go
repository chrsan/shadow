package internal

import (
	"unsafe"
)

func f638(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int64
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	s0i32 = l2
	s1i32 = l4
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		l6 = s0i32
		s0i32 = l2
		l7 = s0i32
		s0i32 = l3
		l1 = s0i32
		s0i32 = l4
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 0
	s2i32 = l0
	s2i32 = int32(ctx.Mem[int(s2i32+55)])
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+55)] = uint8(s1i32)
	s0i32 = l3
	l6 = s0i32
	s0i32 = l4
	l7 = s0i32
lbl0:
	s0i32 = l2
	s1i32 = l7
	s0i32 = s0i32 - s1i32
	s1i32 = 10
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = 2147483647
	l2 = s0i32
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l1
	s1i32 = l6
	s0i32 = s0i32 - s1i32
	s1i32 = 10
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l5
	s1i32 = 31
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l2 = s0i32
	s1i32 = l5
	s2i32 = 10
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s0i32 = s0i32 ^ s1i32
	l2 = s0i32
	s1i32 = 1023
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 9456
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		goto lbl3
	}
	s0i32 = l3
	s1i32 = l3
	s2i32 = 31
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l2 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s0i32 = s0i32 ^ s1i32
	s1i32 = 4095
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s1i32 = l1
	s2i32 = 31
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l2 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s0i32 = s0i32 ^ s1i32
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	s1i32 = 1015
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 9456
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l3
	s0i32 = s0i32 * s1i32
	s1i32 = 6
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	goto lbl5
lbl6:
	s0i32 = l3
	s1i32 = 32768
	s0i32 = s0i32 + s1i32
	s1i32 = 65535
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 16
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l1
		s0i32 = i32DivS(s0i32, s1i32)
		goto lbl5
	}
	s0i32 = l4
	s1i32 = l3
	s1i64 = int64(s1i32)
	s2i64 = 16
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	s2i32 = l1
	s2i64 = int64(s2i32)
	s1i64 = i64DivS(s1i64, s2i64)
	l8 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = -2147483647
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 2147483647
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 24
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i64 = l8
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i64 = l8
	s3i64 = -2147483647
	if s2i64 < s3i64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
lbl5:
	l2 = s0i32
	s1i32 = 31
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l1 = s0i32
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s0i32 = s0i32 ^ s1i32
	l2 = s0i32
lbl3:
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = 1
	return s0i32
}
