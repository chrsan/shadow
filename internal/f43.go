package internal

import (
	"unsafe"
)

func f43(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
	var l2 int64
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 != 0 {
	lbl1:
		s0i32 = l0
		s1i32 = l0
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l2 = s1i64
		s2i64 = 6
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		s1i32 = int32(s1i64)
		if int(s1i32) < 0 || int(s1i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s1i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s1i32].numParams != 1 {
			panic("argument count mismatch")
		}
		s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
		l1 = s0i32
		s1i64 = l2
		s1i32 = int32(s1i64)
		s2i32 = 63
		s1i32 = s1i32 & s2i32
		s0i32 = s0i32 - s1i32
		l0 = s0i32
		s0i32 = l1
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
