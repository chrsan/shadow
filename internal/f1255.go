package internal

import (
	"unsafe"
)

func f1255(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l6 = s1i32
	s2i32 = l2
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 65535
	s2i32 = l3
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	l2 = s2i32
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+44)]))
	l3 = s3i32
	s4i32 = 24
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 - s2i32
	l5 = s1i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = l5
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l5 = s1i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l7 = s2i32
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 * s2i32
	s2i32 = l3
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l2
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l5
	s3i32 = l7
	s4i32 = 8
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 16711935
	s3i32 = s3i32 & s4i32
	s2i32 = s2i32 * s3i32
	s3i32 = l3
	s4i32 = 8
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 16711935
	s3i32 = s3i32 & s4i32
	s4i32 = l2
	s3i32 = s3i32 * s4i32
	s2i32 = s2i32 + s3i32
	s3i32 = -16711936
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 65535
	s2i32 = l4
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	l1 = s2i32
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+44)]))
	l0 = s3i32
	s4i32 = 24
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 - s2i32
	l3 = s1i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l3 = s1i32
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l2 = s2i32
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 * s2i32
	s2i32 = l0
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l1
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l3
	s3i32 = l2
	s4i32 = 8
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 16711935
	s3i32 = s3i32 & s4i32
	s2i32 = s2i32 * s3i32
	s3i32 = l0
	s4i32 = 8
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 16711935
	s3i32 = s3i32 & s4i32
	s4i32 = l1
	s3i32 = s3i32 * s4i32
	s2i32 = s2i32 + s3i32
	s3i32 = -16711936
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
