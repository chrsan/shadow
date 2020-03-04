package internal

import (
	"unsafe"
)

func f942(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 float32
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
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 2
	s2i32 = 0
	f16(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	f185(ctx, s0i32, s1i32)
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = f105(ctx)
	s2i32 = 3
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = 3
	s0i32 = f95(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s1i32 = l2
	s2i32 = 112
	s1i32 = s1i32 + s2i32
	f162(ctx, s0i32, s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = l2
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+124)]))
	l5 = s2f32
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l2
	s1f32 = l5
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+116)]))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l2
	s1f32 = l5
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l2
	s1f32 = l5
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+120)]))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l4
	s1i32 = l3
	s2i32 = l2
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	f174(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	f99(ctx, s0i32, s1i32)
	s0i32 = l2
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = 1
	return s0i32
}
