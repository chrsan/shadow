package internal

import (
	"unsafe"
)

func f1363(ctx *Context, l0 int32, l1 int64, l2 int64, l3 int32) {
	var l4 int64
	_ = l4
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
	s0i32 = l3
	s1i32 = 64
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i64 = l1
		s1i32 = l3
		s2i32 = -64
		s1i32 = s1i32 + s2i32
		s1i64 = int64(uint32(s1i32))
		s0i64 = s0i64 << (uint64(s1i64) & 63)
		l2 = s0i64
		s0i64 = 0
		l1 = s0i64
		goto lbl0
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l2
	s1i32 = l3
	s1i64 = int64(uint32(s1i32))
	l4 = s1i64
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	s1i64 = l1
	s2i32 = 64
	s3i32 = l3
	s2i32 = s2i32 - s3i32
	s2i64 = int64(uint32(s2i32))
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s0i64 = s0i64 | s1i64
	l2 = s0i64
	s0i64 = l1
	s1i64 = l4
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	l1 = s0i64
lbl0:
	s0i32 = l0
	s1i64 = l1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = l2
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
}
