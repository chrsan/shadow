package internal

import (
	"unsafe"
)

func f319(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
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
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int64
	_ = l15
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
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+25)])
	l10 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l7 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l4 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l5 = s0i32
	s0i32 = l0
	s0i32 = int32(int8(ctx.Mem[int(s0i32+24)]))
	l1 = s0i32
lbl1:
	s0i32 = l1
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = l10
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = l4
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l7
		s1i32 = l10
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = l5
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		s1i32 = l7
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		goto lbl2
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l8 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	l9 = s0i32
lbl2:
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l4
	s1i32 = 10
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l12 = s0i32
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	l13 = s0i32
	s1i32 = 6
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l14 = s0i32
	s1i32 = l8
	s2i32 = 10
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l3 = s1i32
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s2i32 = 6
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l4 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l12
		s0i32 = s0i32 - s1i32
		l3 = s0i32
		s0i32 = l9
		s1i32 = 10
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = l5
		s2i32 = 10
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l5 = s1i32
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s1i32 = 32768
		s0i32 = s0i32 + s1i32
		s1i32 = 65535
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = 16
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s1i32 = l3
			s0i32 = i32DivS(s0i32, s1i32)
			goto lbl5
		}
		s0i32 = l2
		s1i32 = l1
		s1i64 = int64(s1i32)
		s2i64 = 16
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i32 = l3
		s2i64 = int64(s2i32)
		s1i64 = i64DivS(s1i64, s2i64)
		l15 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l2
		s1i64 = -2147483647
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l2
		s1i64 = 2147483647
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 24
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i64 = l15
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
		s2i64 = l15
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
		l1 = s0i32
		s0i32 = l0
		s1i32 = l14
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l4
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s2i32 = l1
		s2i64 = int64(s2i32)
		s3i32 = l13
		s4i32 = -64
		s3i32 = s3i32 & s4i32
		s4i32 = 32
		s3i32 = s3i32 | s4i32
		s4i32 = l12
		s3i32 = s3i32 - s4i32
		s3i64 = int64(s3i32)
		s2i64 = s2i64 * s3i64
		s3i64 = 16
		s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
		s2i32 = int32(s2i64)
		s1i32 = s1i32 + s2i32
		s2i32 = 10
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = 1
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l3 = s0i32
	s0i32 = l9
	l5 = s0i32
	s0i32 = l8
	l4 = s0i32
	s0i32 = l11
	l1 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 0
lbl0:
	l1 = s0i32
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l11
	ctx.Mem[int(s0i32+24)] = uint8(s1i32)
	s0i32 = l1
	return s0i32
}
