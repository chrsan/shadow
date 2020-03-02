package internal

import (
	"unsafe"
)

func f181(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int64
	_ = l3
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
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = -1
	s1i32 = -1
	s2i32 = -1
	s3i32 = -2147483648
	s4i32 = l2
	s5i32 = -1
	s4i32 = s4i32 + s5i32
	s5i32 = l2
	s6i32 = -2147483648
	if s5i32 == s6i32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	l2 = s5i32
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	s3i64 = int64(uint32(s3i32))
	s4i32 = l1
	s4i64 = int64(uint32(s4i32))
	s3i64 = s3i64 * s4i64
	l3 = s3i64
	s3i32 = int32(s3i64)
	l1 = s3i32
	s4i32 = l0
	s4i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)])))
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+4)]))
	s6i32 = 2
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s6i32 = 3728
	s5i32 = s5i32 + s6i32
	s5i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)])))
	s4i64 = s4i64 * s5i64
	l4 = s4i64
	s4i32 = int32(s4i64)
	s3i32 = s3i32 + s4i32
	l0 = s3i32
	s4i64 = l4
	s5i64 = 32
	s4i64 = int64(uint64(s4i64) >> (uint64(s5i64) & 63))
	s4i32 = int32(s4i64)
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i64 = l3
	s4i64 = 32
	s3i64 = int64(uint64(s3i64) >> (uint64(s4i64) & 63))
	s3i32 = int32(s3i64)
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = -1
	s2i32 = l0
	s3i32 = l1
	if uint32(s2i32) >= uint32(s3i32) {
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
