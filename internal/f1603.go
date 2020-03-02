package internal

import (
	"unsafe"
)

func f1603(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s2i32 = 3
	s1i32 = s1i32 * s2i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 4
		s3i32 = l4
		s4i32 = 0
		if s3i32 > s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		f1602(ctx, s0i32, s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l4 = s0i32
	}
	s0i32 = l4
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l1
	s0i32 = s0i32 ^ s1i32
	s1i32 = -2048144789
	s0i32 = s0i32 * s1i32
	l3 = s0i32
	s1i32 = 13
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l3
	s0i32 = s0i32 ^ s1i32
	s1i32 = -1028477387
	s0i32 = s0i32 * s1i32
	l3 = s0i32
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l3
	s0i32 = s0i32 ^ s1i32
	l3 = s0i32
	s1i32 = 1
	s2i32 = l3
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l6 = s0i32
	s1i32 = l4
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 & s1i32
	l5 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l7 = s0i32
lbl2:
	s0i32 = l7
	s1i32 = l5
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		s1i64 = int64(uint32(s1i32))
		s2i32 = l2
		s2i64 = int64(uint32(s2i32))
		s3i64 = 32
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s1i64 = s1i64 | s2i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl0
	}
	s0i32 = l6
	s1i32 = l8
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l3
	s1i32 = l1
	s1i64 = int64(uint32(s1i32))
	s2i32 = l2
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	goto lbl0
lbl4:
	s0i32 = l5
	s1i32 = 0
	s2i32 = l4
	s3i32 = l5
	s4i32 = 0
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = s0i32 + s1i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l9
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s1i32 = l4
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl0:
}
