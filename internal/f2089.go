package internal

import (
	"math"
	"unsafe"
)

func f2089(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int64
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
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
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
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
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l3
	s0i32 = f279(ctx, s0i32, s1i32)
	l7 = s0i32
	s1i32 = 6
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		goto lbl1
	}
	s0i32 = l3
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i64
	s0i32 = l3
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = 0
	s3i32 = 0
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	s1i32 = 0
	s2f32 = 0
	s0i32 = f49(ctx, s0i32, s1i32, s2f32)
	s1i64 = l8
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = 512
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	s0i32 = l2
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l8 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i64 = l8
	s0i32 = int32(s0i64)
	s0f32 = math.Float32frombits(uint32(s0i32))
lbl1:
	l9 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l10 = s0f32
	s0f32 = l9
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l10
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = l9
	s2f32 = l10
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
lbl3:
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 | s1i32
	l4 = s0i32
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l3
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l7
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l2
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	switch s1i32 {
	case 0:
		goto lbl8
	case 1:
		goto lbl5
	case 2:
		goto lbl9
	default:
		goto lbl7
	}
lbl9:
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = l3
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
	s5i32 = l3
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
	s6i32 = l3
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+24)]))
	s7i32 = l3
	s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+28)]))
	f175(ctx, s1i32, s2f32, s3f32, s4f32, s5f32, s6f32, s7f32)
	s1i32 = l6
	goto lbl6
lbl8:
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = l3
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
	s5i32 = l3
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
	f130(ctx, s1i32, s2f32, s3f32, s4f32, s5f32)
	s1i32 = l5
	goto lbl6
lbl7:
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s1i32 = f30(ctx, s1i32, s2f32, s3f32)
	s1i32 = l4
lbl6:
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
lbl5:
	s0i32 = l0
	s1i32 = l3
	s0i32 = f279(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i32 = 6
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl10:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l2
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	switch s1i32 {
	case 0:
		goto lbl14
	case 1:
		goto lbl11
	case 2:
		goto lbl13
	default:
		goto lbl15
	}
lbl15:
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s1i32 = f30(ctx, s1i32, s2f32, s3f32)
	s1i32 = l4
	goto lbl12
lbl14:
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = l3
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
	s5i32 = l3
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
	f130(ctx, s1i32, s2f32, s3f32, s4f32, s5f32)
	s1i32 = l5
	goto lbl12
lbl13:
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = l3
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
	s5i32 = l3
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
	s6i32 = l3
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+24)]))
	s7i32 = l3
	s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+28)]))
	f175(ctx, s1i32, s2f32, s3f32, s4f32, s5f32, s6f32, s7f32)
	s1i32 = l6
lbl12:
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
lbl11:
	s0i32 = l0
	s1i32 = l3
	s0i32 = f279(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i32 = 6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
lbl0:
	s0i32 = l3
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
