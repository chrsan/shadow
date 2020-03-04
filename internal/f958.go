package internal

import (
	"unsafe"
)

func f958(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 float32
	_ = l2
	var l3 float32
	_ = l3
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
	s0i32 = l1
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l2 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0f32
	s0i32 = l1
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l2
	s2f32 = l3
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l1
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l2 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0f32
	s0i32 = l1
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l1
	s1f32 = l2
	s2f32 = l3
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s0i32 = l1
	s1i64 = 550821167104
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = 1
	return s0i32
}
