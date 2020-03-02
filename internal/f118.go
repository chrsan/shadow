package internal

import (
	"unsafe"
)

func f118(ctx *Context, l0 int32) int32 {
	var l1 int64
	_ = l1
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
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	s0i32 = l0
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1i32 = s1i32 - s2i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 * s1i64
	l1 = s0i64
	s0i32 = int32(s0i64)
	s1i32 = 0
	s2i64 = l1
	s3i64 = 2147483648
	if uint64(s2i64) < uint64(s3i64) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = 0
	s2i64 = l1
	s3i64 = 0
	if s2i64 > s3i64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	return s0i32
}
