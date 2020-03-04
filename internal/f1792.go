package internal

import (
	"unsafe"
)

func f1792(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	s0i32 = ctx.g0
	s1i32 = 1136
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		f634(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1120)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1124)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = 65535
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1128)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = 65535
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1132)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l3
	s2i32 = 1120
	s1i32 = s1i32 + s2i32
	s0i32 = f153(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		s2i32 = 8
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = 128
		s2i32 = s2i32 + s3i32
		s3i32 = 8
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
		s4i32 = 128
		s3i32 = s3i32 + s4i32
		s4i32 = 8
		s3i32 = s3i32 >> (uint32(s4i32) & 31)
		s4i32 = l2
		s5i32 = 1
		f148(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l3
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s3i32 = l2
	s1i32 = f151(ctx, s1i32, s2i32, s3i32)
	l0 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1100)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+1104)]))
	f634(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = f74(ctx, s0i32)
	s0i32 = l0
	f40(ctx, s0i32)
lbl0:
	s0i32 = l3
	s1i32 = 1136
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
