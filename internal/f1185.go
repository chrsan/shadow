package internal

import (
	"unsafe"
)

func f1185(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l0
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	lbl2:
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 16
	s0i32 = f102(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
	s3i32 = l2
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	s4i32 = l2
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+44)]))
	s0i32 = f244(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s0i32 = f360(ctx, s0i32, s1i32)
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s0i32 = f360(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = l0
	s2i32 = l2
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = 1
	f358(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	f13(ctx, s0i32)
lbl0:
	s0i32 = l2
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
