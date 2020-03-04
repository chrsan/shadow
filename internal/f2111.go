package internal

import (
	"unsafe"
)

func f2111(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 float32
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = l3
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l5 = s0i32
	s0i32 = l0
	s1i32 = l3
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s0i32 = f446(ctx, s0i32, s1i32)
	l6 = s0i32
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l2 = s0i32
		goto lbl0
	}
	s0i32 = 0
	l2 = s0i32
lbl2:
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0f32
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l7
	s1f32 = 1
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l5
	s1i32 = l2
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1f32 = l7
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
lbl3:
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl0:
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl4
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l5
	s3i32 = l2
	f447(ctx, s0i32, s1i32, s2i32, s3i32)
lbl4:
	s0i32 = l3
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	return s0i32
}
