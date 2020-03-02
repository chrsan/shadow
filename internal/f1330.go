package internal

import (
	"unsafe"
)

func f1330(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 int32) {
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
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 112
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s0i32 = f65(ctx, s0i32)
	if s0i32 != 0 {
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l10 = s0f32
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l11 = s0f32
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l12 = s0f32
		s0i32 = l5
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2f32 = l2
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = s1f32
		s0i32 = l5
		s1f32 = l12
		s2f32 = l3
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = s1f32
		s0i32 = l5
		s1f32 = l11
		s2f32 = l2
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = s1f32
		s0i32 = l5
		s1f32 = l10
		s2f32 = l3
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = s1f32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		s1i32 = 6
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l6 = s0i32
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s1i32 = s1i32 | s2i32
		s0i32 = s0i32 | s1i32
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l5
		s1i32 = 96
		s0i32 = s0i32 + s1i32
		goto lbl2
	lbl3:
		s0i32 = l4
		s1i32 = l5
		s2i32 = 96
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s3i32 = 24
		s2i32 = s2i32 + s3i32
		s3i32 = l6
		s0i32 = f64(ctx, s0i32, s1i32, s2i32, s3i32)
	lbl2:
		l7 = s0i32
		s0i32 = l5
		s1i32 = 96
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l0
		s1i32 = l7
		s0i32 = f70(ctx, s0i32, s1i32)
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4388)]))
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 1
		f57(ctx, s0i32, s1i32)
	}
	s0i32 = l5
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l4
	s3i32 = 0
	s4i32 = l6
	s0i32 = f61(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl6:
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l8 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4428)]))
	l9 = s0i32
	s0i32 = l5
	s1f32 = l3
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l5
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l9
	s1i32 = l8
	s2i32 = l1
	s3i32 = l5
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s4i32 = l7
	f1086(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l4
	if s0i32 != 0 {
		goto lbl6
	}
lbl5:
	s0i32 = l6
	s0i32 = int32(ctx.Mem[int(s0i32+64)])
	if s0i32 != 0 {
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		f54(ctx, s0i32)
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
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
	s0i32 = f23(ctx, s0i32)
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
lbl0:
	s0i32 = l5
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
