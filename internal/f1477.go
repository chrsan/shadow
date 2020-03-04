package internal

import (
	"unsafe"
)

func f1477(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l6
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l6
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s0i32 = f189(ctx, s0i32, s1i32, s2i32, s3i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 312
	s2i32 = 8
	s0i32 = f56(ctx, s0i32, s1i32, s2i32)
	l8 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	s0i32 = l5
	s1i32 = l8
	s2i32 = 304
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 1232
	s2i32 = l8
	s3i32 = l7
	s2i32 = s2i32 - s3i32
	f52(ctx, s0i32, s1i32, s2i32)
	s0i32 = l8
	s1i32 = l3
	s2i32 = l1
	s3i32 = l2
	s0i32 = f1491(ctx, s0i32, s1i32, s2i32, s3i32)
	l1 = s0i32
	s0i32 = l6
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = 0
	l7 = s0i32
	s0i32 = l1
	s1i32 = l6
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s0i32 = f1489(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l8
	s0i32 = f1488(ctx, s0i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l0
	s2i32 = l4
	s3i32 = l6
	s4i32 = 4
	s3i32 = s3i32 + s4i32
	s0i32 = f1475(ctx, s0i32, s1i32, s2i32, s3i32)
	l7 = s0i32
lbl0:
	s0i32 = l6
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l7
	return s0i32
}
