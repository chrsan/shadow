package internal

import (
	"unsafe"
)

func f1077(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = 2
		s1i32 = l3
		s1i32 = f297(ctx, s1i32)
		if s1i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = 3
lbl0:
	l3 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 17179869184
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s2i32 = 104
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 17179869184
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s2i32 = 128
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+100)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+100)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+124)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+124)] = uint8(s1i32)
}
