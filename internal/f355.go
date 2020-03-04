package internal

import (
	"unsafe"
)

func f355(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l4 = s0i32
	s0i32 = l1
	s0i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])))
	l6 = s0i64
	s0i32 = l1
	s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+6)])))
	l3 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])))
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = l6
	s2i32 = l5
	s2i64 = int64(uint32(s2i32))
	s3i64 = 48
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s3i64 = 48
	s2i64 = s2i64 >> (uint64(s3i64) & 63)
	s1i64 = s1i64 + s2i64
	l6 = s1i64
	s2i64 = 2147483647
	s3i64 = l6
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l6 = s1i64
	s2i64 = -2147483647
	s3i64 = l6
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i64)
	s0i32 = l0
	s1i32 = l4
	s1i64 = int64(uint32(s1i32))
	s2i32 = l3
	s2i64 = int64(uint32(s2i32))
	s3i64 = 48
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s3i64 = 48
	s2i64 = s2i64 >> (uint64(s3i64) & 63)
	s1i64 = s1i64 + s2i64
	l6 = s1i64
	s2i64 = 2147483647
	s3i64 = l6
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l6 = s1i64
	s2i64 = -2147483647
	s3i64 = l6
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i64)
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+24)])
	l3 = s0i32
	s1i32 = 5
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l1 = s0i32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl1
		case 1:
			goto lbl1
		case 2:
			goto lbl2
		case 3:
			goto lbl4
		case 4:
			goto lbl1
		default:
			goto lbl5
		}
	lbl5:
		s0i32 = l4
		s1i32 = 7
		s0i32 = s0i32 + s1i32
		s1i32 = 3
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		goto lbl0
	lbl4:
		s0i32 = 2
		l1 = s0i32
		goto lbl1
	}
	s0i32 = l2
	s1i32 = 60
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 3323
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 3282
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 3256
	s1i32 = l2
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl2:
	s0i32 = 4
	l1 = s0i32
lbl1:
	s0i32 = l1
	s1i32 = l4
	s0i32 = s0i32 * s1i32
lbl0:
	l1 = s0i32
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
