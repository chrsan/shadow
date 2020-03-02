package internal

import (
	"unsafe"
)

func f1791(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
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
	s0i32 = ctx.g0
	s1i32 = 1120
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		s3i32 = l3
		f630(ctx, s0i32, s1i32, s2i32, s3i32)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l4
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s4i32 = l3
	s2i32 = f151(ctx, s2i32, s3i32, s4i32)
	l0 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+1100)]))
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+1104)]))
	f630(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l0
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = f74(ctx, s0i32)
	s0i32 = l0
	f40(ctx, s0i32)
lbl0:
	s0i32 = l4
	s1i32 = 1120
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
