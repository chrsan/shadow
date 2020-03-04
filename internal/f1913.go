package internal

import (
	"unsafe"
)

func f1913(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+196)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1064
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	l2 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		f112(ctx, s0i32)
		s0i32 = l0
		s1i32 = 1032
		s0i32 = s0i32 + s1i32
		f367(ctx, s0i32)
		s0i32 = l0
		s1i32 = l0
		s2i32 = 1056
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		if s1i32 != 0 {
			s1i32 = l0
			s2i32 = 1036
			s1i32 = s1i32 + s2i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		} else {
			s1i32 = 0
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1064)])) = uint32(s1i32)
	}
}
