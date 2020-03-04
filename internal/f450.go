package internal

import (
	"unsafe"
)

func f450(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 int64
	_ = l5
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s1i32 = l1
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	f1291(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = 19200
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 156
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s0i32 = f179(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint32(s1i32)
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l5 = s0i64
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1032
	s0i32 = s0i32 + s1i32
	s1i32 = 52
	s2i32 = l0
	s3i32 = 200
	s2i32 = s2i32 + s3i32
	s3i32 = 832
	s0i32 = f368(ctx, s0i32, s1i32, s2i32, s3i32)
	l2 = s0i32
	s0i32 = l0
	s1i32 = 1076
	s0i32 = s0i32 + s1i32
	s1i64 = l5
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 1072
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1064
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l2
	s1i32 = f94(ctx, s1i32)
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s0i32 = f225(ctx, s0i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s1i32 = l0
	s2i32 = 1068
	s1i32 = s1i32 + s2i32
	f435(ctx, s0i32, s1i32)
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1084)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1088
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 28
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = f259(ctx)
	f1077(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 36
	s0i32 = f17(ctx, s0i32)
	l2 = s0i32
	s1i32 = l4
	s0i32 = f179(ctx, s0i32, s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1084)]))
	l1 = s0i32
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1084)])) = uint32(s1i32)
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = f41(ctx, s0i32)
	f12(ctx, s0i32)
lbl0:
	s0i32 = l0
	return s0i32
}
