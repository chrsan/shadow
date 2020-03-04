package internal

import (
	"unsafe"
)

func f552(ctx *Context, l0 int32, l1 int32, l2 int32, l3 float32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int64
	_ = l11
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 176
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l5 = s0i32
	s1i32 = 192
	s0i32 = s0i32 & s1i32
	s1i32 = 128
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = f67(ctx, s1i32)
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s1f32 = l3
		s2i32 = l1
		s0f32 = f443(ctx, s0i32, s1f32, s2i32)
		l3 = s0f32
	}
	s0i32 = l1
	s1i32 = l2
	s2i32 = l2
	s3i32 = 3
	f55(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
	s0i32 = l4
	s1f32 = l3
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = s1f32
	s0i32 = l4
	s1i32 = 1
	s2i32 = l4
	s3i32 = 144
	s2i32 = s2i32 + s3i32
	s2i32 = f233(ctx, s2i32)
	l6 = s2i32
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l7 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		f13(ctx, s0i32)
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 1
	s0i32 = s0i32 | s1i32
	l1 = s0i32
	s1i32 = 18
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l1
		s2i32 = 8
		s1i32 = f50(ctx, s1i32, s2i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl3
	}
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	l1 = s0i32
lbl3:
	s0i32 = l4
	s1i32 = l4
	s2i32 = 144
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s3i32 = l6
	s1i32 = f232(ctx, s1i32, s2i32, s3i32)
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l6
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l4
		s1i32 = 160
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l4
		s1i32 = 144
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = s0i32 | s1i32
		l10 = s0i32
		s0i32 = l1
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0i64
		s0i32 = 0
		l1 = s0i32
	lbl6:
		s0i32 = l4
		s1i64 = l11
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l7
		s2i32 = l2
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = l1
		s4i32 = l8
		if s3i32 == s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l11 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l10
		f103(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l9
		f103(ctx, s0i32, s1i32)
		s0i32 = l2
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
	}
	s0i32 = l5
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		f13(ctx, s0i32)
	}
	s0i32 = l4
	s1i32 = 176
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
