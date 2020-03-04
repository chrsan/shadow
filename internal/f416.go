package internal

import (
	"unsafe"
)

func f416(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 1104
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = int32(ctx.Mem[int(s0i32+85)])
		l5 = s0i32
		goto lbl1
	}
	s0i32 = l4
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l5 = s0i32
	s0i32 = l4
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = l5
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
lbl1:
	s0i32 = l5
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		f261(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 1080
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 1024
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1076)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l3
	s2i32 = 52
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 20824
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	s2i32 = l1
	s3i32 = 20
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s3i32 = int32(ctx.Mem[int(s3i32+40)])
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = f81(ctx, s0i32, s1i32)
	s0i32 = l3
	s1i32 = l1
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l3
	f261(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s0i32 = f74(ctx, s0i32)
	s0i32 = l4
	f40(ctx, s0i32)
lbl0:
	s0i32 = l3
	s1i32 = 1104
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
