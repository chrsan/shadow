package internal

import (
	"unsafe"
)

func f588(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 26040
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = f438(ctx)
		l2 = s0i32
	}
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+32)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 26060
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
