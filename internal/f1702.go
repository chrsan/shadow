package internal

import (
	"unsafe"
)

func f1702(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 int32, l5 int32, l6 float32, l7 int32) int32 {
	var l8 int32
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = l0
	s1i32 = l7
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	s0i32 = l0
	s1f32 = l6
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l0
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l7 = s0i32
	s0i32 = l0
	s1i32 = 108
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l8 = s0i32
	s0i32 = l0
	s1i32 = 120
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 2
	l5 = s0i32
	s0f32 = l3
	s1f32 = 1
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1f32 = 1
	s2f32 = l3
	s1f32 = s1f32 / s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = 0
	l5 = s0i32
lbl0:
	s0i32 = l0
	s1i32 = l4
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 25464
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l5
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 25476
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l0
	s1i64 = -4294967296
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint64(s1i64)
	s0i32 = l8
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s2i32 = 3
	s1i32 = s1i32 * s2i32
	f441(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+118)])
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+118)] = uint8(s1i32)
	s0i32 = l7
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	f441(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint32(s1i32)
	s0i32 = l0
	s1f32 = 1
	s2f32 = l6
	s3f32 = 4
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 / s2f32
	l2 = s1f32
	s2f32 = l2
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l0
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+106)])
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+106)] = uint8(s1i32)
	s0i32 = l0
	return s0i32
}
