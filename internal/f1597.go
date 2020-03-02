package internal

import (
	"unsafe"
)

func f1597(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 336
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l5
	s2i32 = 304
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l5
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	s2i32 = 256
	s3i32 = 256
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
	l7 = s0i32
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 4
	s2i32 = l5
	s3i32 = 16
	s2i32 = s2i32 + s3i32
	f434(ctx, s0i32, s1i32, s2i32)
	s0i32 = l6
	s1i32 = 4
	s2i32 = l5
	s3i32 = 24
	s2i32 = s2i32 + s3i32
	f149(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = 0
	s0i32 = f205(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		s0i32 = l4
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = 90
			s2i32 = l5
			s3i32 = 8
			s2i32 = s2i32 + s3i32
			f16(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1i32 = l6
		f99(ctx, s0i32, s1i32)
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l6
	f99(ctx, s0i32, s1i32)
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s1i32 = 93
	s2i32 = l5
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	f16(ctx, s0i32, s1i32, s2i32)
lbl0:
	s0i32 = l6
	s1i32 = 4
	s2i32 = l5
	s3i32 = 24
	s2i32 = s2i32 + s3i32
	f173(ctx, s0i32, s1i32, s2i32)
	s0i32 = l6
	s1i32 = l3
	s2i32 = 1
	f272(ctx, s0i32, s1i32, s2i32)
	s0i32 = l7
	f43(ctx, s0i32)
	s0i32 = l5
	s1i32 = 336
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
