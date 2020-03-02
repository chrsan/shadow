package internal

import (
	"unsafe"
)

func f185(ctx *Context, l0 int32, l1 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 24
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
}
