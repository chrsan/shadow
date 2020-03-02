package internal

import (
	"unsafe"
)

func f194(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s1i32 = l1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		s0i32 = l0
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l2 = s0i32
		s0i32 = l0
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+10)])
		l2 = s2i32
		s3i32 = 252
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		l3 = s1i32
		s2i32 = 252
		s1i32 = s1i32 & s2i32
		s2i32 = l2
		s3i32 = 3
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		l2 = s1i32
		s2i32 = 251
		s1i32 = s1i32 & s2i32
		s2i32 = l3
		s3i32 = 4
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = 251
		s1i32 = s1i32 & s2i32
		s2i32 = l2
		s3i32 = 4
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+8)])
		l2 = s0i32
		s0i32 = l0
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+8)])
		ctx.Mem[int(s0i32+8)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = l2
		ctx.Mem[int(s0i32+8)] = uint8(s1i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+9)])
		l2 = s0i32
		s0i32 = l0
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+9)])
		ctx.Mem[int(s0i32+9)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = l2
		ctx.Mem[int(s0i32+9)] = uint8(s1i32)
	}
}
