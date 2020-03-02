package internal

import (
	"unsafe"
)

func f137(ctx *Context) int32 {
	var l0 int32
	_ = l0
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	s0i32 = 28785
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l0 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl1
	case 1:
		goto lbl0
	default:
		goto lbl2
	}
lbl2:
	s0i32 = 28785
	s1i32 = 28785
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l0 = s1i32
	s2i32 = 1
	s3i32 = l0
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 92
	s0i32 = f17(ctx, s0i32)
	l0 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+64)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+90)] = uint8(s1i32)
	s0i32 = l0
	s1i64 = 4294967297
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = -1409286144
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+86)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint64(s1i64)
	s0i32 = 28788
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	s2i32 = 0
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	s0i32 = 28785
	s1i32 = 2
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	goto lbl0
lbl1:
lbl3:
	s0i32 = 28785
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl0:
	s0i32 = 28788
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	return s0i32
}
