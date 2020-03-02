package internal

import (
	"unsafe"
)

func f1693(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int64
	_ = l5
	var l6 int64
	_ = l6
	var l7 int64
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
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
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 88
	s1i32 = s1i32 + s2i32
	s0i32 = f1697(ctx, s0i32, s1i32)
	l4 = s0i32
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl3
	case 1:
		goto lbl1
	case 2:
		goto lbl2
	default:
		goto lbl4
	}
lbl4:
	s0i32 = l0
	s1i32 = l2
	s2i32 = 0
	f78(ctx, s0i32, s1i32, s2i32)
	goto lbl0
lbl3:
	s0i32 = l0
	s1i32 = l2
	s2i32 = 0
	f78(ctx, s0i32, s1i32, s2i32)
	goto lbl0
lbl2:
	s0i32 = l0
	s1i32 = l3
	s2i32 = 88
	s1i32 = s1i32 + s2i32
	s2i32 = 0
	f78(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
	l1 = s0i32
	s0i32 = l0
	s1i32 = 25480
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	s2i32 = 0
	f78(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	goto lbl0
lbl1:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l3
	s3i32 = 80
	s2i32 = s2i32 + s3i32
	s3i32 = l3
	s4i32 = 72
	s3i32 = s3i32 + s4i32
	s4i32 = 0
	s0i32 = f313(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s2i32 = 0
		f78(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+140)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 1065353216
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 4539628424389459968
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint16(s1i32)
	s0i32 = l0
	s1i32 = l3
	s2i32 = 96
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s0i32 = f311(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+140)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 1065353216
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 4539628424389459968
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint16(s1i32)
	s0i32 = l0
	s1i32 = l3
	s2i32 = 96
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s0i32 = f311(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0f32
	s0i32 = l3
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l8 = s1f32
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+104)]))
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l8
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+116)]))
	s4i32 = l3
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+108)]))
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s0i32 = f327(ctx, s0i32, s1f32, s2f32)
	if s0i32 != 0 {
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l8 = s0f32
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
		s0i32 = l3
		s1f32 = l8
		s1f32 = -s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = s1f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l8 = s0f32
		s0i32 = l3
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
		s2f32 = l9
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l8
		s2f32 = l9
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l3
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
		l5 = s0i64
		s0i32 = l3
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l6 = s0i64
		goto lbl6
	}
	s0i32 = l3
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	l5 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	l6 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
lbl6:
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+141)] = uint8(s1i32)
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i64
	s0i32 = l0
	s1i64 = l6
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = l7
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = l5
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
lbl0:
	s0i32 = l3
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
