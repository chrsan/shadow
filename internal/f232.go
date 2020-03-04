package internal

import (
	"math"
	"unsafe"
)

func f232(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
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
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 5
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l3
	f688(ctx, s0i32, s1i32)
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s0f32 = s0f32 - s1f32
	l8 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s0f32 = s0f32 - s1f32
	l9 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0f32 = l8
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l9
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	s0f32 = s0f32 - s1f32
	l8 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	s0f32 = s0f32 - s1f32
	l9 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l8
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l9
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl3:
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l7 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l1
	s1i64 = l7
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l1
	s1i64 = l7
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = 1
	l2 = s0i32
	goto lbl0
lbl1:
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s0i32 = f444(ctx, s0i32, s1i32, s2i32)
lbl0:
	s0i32 = 1
	s1i32 = l2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l5 = s0i32
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 2
	s0i32 = s0i32 | s1i32
	l6 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f32 = 0
	l8 = s0f32
lbl5:
	s0f32 = l8
	s1i32 = l1
	s2i32 = l4
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 * s1f32
	l8 = s0f32
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
		goto lbl5
	}
	s0f32 = l8
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	s1i32 = 31
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l2 = s0i32
	s0i32 = 1
	l4 = s0i32
lbl6:
	s0i32 = l1
	s1i32 = l4
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
lbl4:
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
	s0i32 = l5
	return s0i32
}
