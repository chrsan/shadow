package internal

import (
	"unsafe"
)

func f365(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int64
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 20760
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l5 = s0i64
	s0i32 = l0
	s1i32 = 20120
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = l5
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l2
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	s1i32 = f209(ctx, s1i32)
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = 24
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s2i32 = l1
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l3 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = l2
	s1i32 = s1i32 * s2i32
	l4 = s1i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = l2
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = 130816
	s1i32 = s1i32 & s2i32
	s2i32 = l2
	s3i32 = l1
	s4i32 = -16777216
	s3i32 = s3i32 & s4i32
	s2i32 = s2i32 | s3i32
	s1i32 = s1i32 | s2i32
	s2i32 = l3
	s3i32 = 16
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
}
