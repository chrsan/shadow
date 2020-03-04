package internal

import (
	"unsafe"
)

func f613(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s1i32 = 25604
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l1 = s0i32
	if s0i32 != 0 {
	lbl1:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		s0i32 = l1
		f13(ctx, s0i32)
		s0i32 = l2
		l1 = s0i32
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l0
	return s0i32
}
