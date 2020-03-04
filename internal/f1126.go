package internal

import (
	"unsafe"
)

func f1126(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
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
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 208
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l7
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l8 = s0i32
		s0i32 = l7
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l7
		s1i32 = l8
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l7
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	s5i32 = 4
	s4i32 = s4i32 + s5i32
	s0i32 = f1127(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s0i32 = f118(ctx, s0i32)
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = l1
	s2i32 = 1
	s1i32 = f203(ctx, s1i32, s2i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+172)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+188)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+164)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = 21300
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	s1i32 = 4
	s0i32 = s0i32 | s1i32
	s1i32 = l4
	s0i32 = f1939(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 112
		s0i32 = s0i32 + s1i32
		s0i32 = f225(ctx, s0i32)
		l2 = s0i32
		s0i32 = l6
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		s0i32 = l6
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
		s0i32 = l6
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 24
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 13195221663744
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l3 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l7 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l9 = s0i32
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l8
		s2i32 = l7
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l3
		s2i32 = l9
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l6
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		f435(ctx, s0i32, s1i32)
		s0i32 = l6
		s1i32 = 72
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s1f32 = float32(s1i32)
		s1f32 = -s1f32
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s2f32 = float32(s2i32)
		s2f32 = -s2f32
		f117(ctx, s0i32, s1f32, s2f32)
		s0i32 = l6
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s2i32 = 72
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = 1
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l1
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
			s2i32 = -193
			s1i32 = s1i32 & s2i32
			s2i32 = 64
			s1i32 = s1i32 | s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		}
		s0i32 = l6
		s1i32 = 160
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = l1
		s3i32 = 0
		s4i32 = 0
		f132(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l1
		s0i32 = f23(ctx, s0i32)
		s0i32 = l2
		f112(ctx, s0i32)
	}
	s0i32 = l6
	s1i32 = 21300
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint32(s1i32)
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+172)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l0
	s1i32 = l0
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
		goto lbl4
	}
	s0i32 = l0
	f12(ctx, s0i32)
lbl4:
	s0i32 = 1
	l9 = s0i32
lbl0:
	s0i32 = l6
	s1i32 = 208
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l9
	return s0i32
}
