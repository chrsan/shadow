package internal

import (
	"unsafe"
)

func f1475(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s1i32 = 112
	s2i32 = 4
	s0i32 = f56(ctx, s0i32, s1i32, s2i32)
	l4 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	s0i32 = l0
	s1i32 = l4
	s2i32 = 104
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1233
	s2i32 = l4
	s3i32 = l5
	s2i32 = s2i32 - s3i32
	f52(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l4
	s1i32 = l1
	s2i32 = l2
	f1436(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 26880
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+88)])
	s1i32 = 255
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
lbl0:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
	l1 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l4
		s2i32 = 48
		s1i32 = s1i32 + s2i32
		s1i32 = f24(ctx, s1i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	}
	s0i32 = l1
	s1i32 = 12
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
	s2i32 = 2
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
lbl1:
	s0i32 = l4
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 26856
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	return s0i32
}
