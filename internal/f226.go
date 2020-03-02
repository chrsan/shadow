package internal

import (
	"unsafe"
)

func f226(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	l3 = s0i32
	s0i32 = l0
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+40)])
	l4 = s1i32
	ctx.Mem[int(s0i32+40)] = uint8(s1i32)
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		f223(ctx, s0i32, s1i32)
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l1
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	f512(ctx, s0i32, s1i32)
lbl0:
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+41)])
	ctx.Mem[int(s0i32+41)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+42)])
	ctx.Mem[int(s0i32+42)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
}
