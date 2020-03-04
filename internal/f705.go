package internal

import (
	"unsafe"
)

func f705(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l9 int64
	_ = l9
	var l10 int64
	_ = l10
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
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i64
	s0i32 = int32(s0i64)
	l6 = s0i32
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0i64
	s1i64 = 32
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	s0i32 = int32(s0i64)
	l3 = s0i32
	s1i64 = l9
	s2i64 = 32
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	l4 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i64 = l10
		s1i32 = int32(s1i64)
		l3 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0i32 = s0i32 - s1i32
		s1i32 = 20
		s0i32 = i32DivS(s0i32, s1i32)
		s1i32 = l3
		s2i32 = l6
		s1i32 = s1i32 - s2i32
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = 204
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1i32 = s1i32 - s2i32
		s2i32 = -20
		s1i32 = i32DivS(s1i32, s2i32)
		s0i32 = s0i32 + s1i32
		l5 = s0i32
	}
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l8 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l3 = s2i32
	s1i32 = s1i32 - s2i32
	l4 = s1i32
	s2i32 = 2
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = 204
	s1i32 = s1i32 * s2i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s2i32 = 0
	s3i32 = l4
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	s2i32 = s2i32 + s3i32
	l7 = s2i32
	s1i32 = s1i32 - s2i32
	l4 = s1i32
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l5
		s2i32 = l4
		s1i32 = s1i32 - s2i32
		f701(ctx, s0i32, s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l8 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l3 = s0i32
	}
	s0i32 = l3
	s1i32 = l7
	s2i32 = 204
	s1i32 = i32DivU(s1i32, s2i32)
	l5 = s1i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l3
	s1i32 = l8
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
	} else {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l7
		s2i32 = l5
		s3i32 = 204
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 20
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
	}
	l5 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl5:
		s0i32 = l5
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0i32 = s0i32 - s1i32
		s1i32 = 4080
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l5 = s0i32
			s0i32 = l4
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l4 = s0i32
		}
		s0i32 = l1
		s1i32 = l3
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0i32 = s0i32 - s1i32
		s1i32 = 4080
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l6
			s2i32 = 4
			s1i32 = s1i32 + s2i32
			l2 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = l6
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l3 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l2
			l6 = s0i32
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l7
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
	}
}
