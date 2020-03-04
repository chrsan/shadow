package internal

import (
	"unsafe"
)

func f463(ctx *Context, l0 int32, l1 int32) {
	var l2 int64
	_ = l2
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
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	s0i32 = l0
	s1i32 = l1
	s2i32 = l1
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)])))
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s4i32 = 2
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s4i32 = 3728
	s3i32 = s3i32 + s4i32
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i64 = int64(s3i32)
	s2i64 = s2i64 * s3i64
	l2 = s2i64
	s2i32 = int32(s2i64)
	s3i32 = 0
	s4i64 = l2
	s5i64 = 2147483648
	if uint64(s4i64) < uint64(s5i64) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s0i32 = f198(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		f860(ctx)
		panic("unreachable executed")
	}
}
