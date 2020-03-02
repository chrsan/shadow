package internal

import (
	"unsafe"
)

func f1368(ctx *Context, l0 int32, l1 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s1f64 float64
	_ = s1f64
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 15
	s1i32 = s1i32 + s2i32
	s2i32 = -16
	s1i32 = s1i32 & s2i32
	l1 = s1i32
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l1
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1f64 = f1362(ctx, s1i64, s2i64)
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
}
