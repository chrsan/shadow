package internal

import (
	"unsafe"
)

func f1118(ctx *Context, l0 float32, l1 float32, l2 float32, l3 float32, l4 float32, l5 float32) int32 {
	var l6 int32
	_ = l6
	var s0i32 int32
	_ = s0i32
	var s1i64 int64
	_ = s1i64
	var s1f32 float32
	_ = s1f32
	s0i32 = 40
	s0i32 = f17(ctx, s0i32)
	l6 = s0i32
	s1i64 = 550821167104
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l6
	s1f32 = l5
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l6
	s1f32 = l4
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l6
	s1f32 = l3
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l6
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l6
	s1f32 = l1
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l6
	s1f32 = l0
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l6
	return s0i32
}
