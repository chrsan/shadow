package internal

import (
	"unsafe"
)

func f1524(ctx *Context, l0 int32) int32 {
	var l1 int64
	_ = l1
	var l2 int64
	_ = l2
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l0 = s0i32
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])))
	s1i32 = l0
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])))
	s0i64 = s0i64 - s1i64
	l1 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l0
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])))
	s2i32 = l0
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)])))
	s1i64 = s1i64 - s2i64
	l2 = s1i64
	s2i64 = 1
	if s1i64 < s2i64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 | s1i32
	s1i64 = l1
	s2i64 = l2
	s1i64 = s1i64 | s2i64
	s2i64 = 2147483648
	s1i64 = s1i64 + s2i64
	s2i64 = 4294967295
	if uint64(s1i64) > uint64(s2i64) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 | s1i32
	return s0i32
}
