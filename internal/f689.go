package internal

import (
	"unsafe"
)

func f689(ctx *Context, l0 int32, l1 float32, l2 int32, l3 int32) int32 {
	var l4 int32
	_ = l4
	var l5 float64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1f32 float32
	_ = s1f32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	s0i32 = ctx.g0
	s1i32 = 208
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1f64 = float64(s1f32)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = s1f64
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1f64 = float64(s1f32)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = s1f64
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s1f64 = float64(s1f32)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = s1f64
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s1f64 = float64(s1f32)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = s1f64
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s1f64 = float64(s1f32)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = s1f64
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s1f64 = float64(s1f32)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = s1f64
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s1f64 = float64(s1f32)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = s1f64
	s0i32 = l4
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s1f64 = float64(s1f32)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = s1f64
	s0i32 = 1
	l0 = s0i32
	s0i32 = l4
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	s1f32 = l1
	s1f64 = float64(s1f32)
	s2i32 = l4
	s3i32 = 112
	s2i32 = s2i32 + s3i32
	s3i32 = l3
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, float64, int32) int32)(table[s3i32].f()))(ctx, s0i32, s1f64, s2i32)
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l0 = s0i32
		goto lbl0
	}
	s0i32 = l4
	s1i32 = l4
	s2i32 = 144
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
	f1422(ctx, s0i32, s1i32, s2f64)
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0f64
	s0i32 = l2
	s1i32 = l4
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l2
	s1f64 = l5
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l5 = s0f64
	s0i32 = l2
	s1i32 = l4
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l2
	s1f64 = l5
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l5 = s0f64
	s0i32 = l2
	s1i32 = l4
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l2
	s1f64 = l5
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l5 = s0f64
	s0i32 = l2
	s1i32 = l4
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l2
	s1f64 = l5
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l5 = s0f64
	s0i32 = l2
	s1i32 = l4
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
	s0i32 = l2
	s1f64 = l5
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l5 = s0f64
	s0i32 = l2
	s1i32 = l4
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
	s0i32 = l2
	s1f64 = l5
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
	l5 = s0f64
	s0i32 = l2
	s1i32 = l4
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = s1f32
	s0i32 = l2
	s1f64 = l5
	s1f32 = float32(s1f64)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
lbl0:
	s0i32 = l4
	s1i32 = 208
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
