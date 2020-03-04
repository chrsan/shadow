package internal

import (
	"math"
	"unsafe"
)

func f1131(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l8 float32
	_ = l8
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
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 3440
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l4 = s0i32
	s1i32 = 192
	s0i32 = s0i32 & s1i32
	s1i32 = 64
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l7 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l4 = s0i32
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
		s0i32 = l4
		s1i32 = l4
		s1i32 = f67(ctx, s1i32)
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s1f32 = l7
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l3
	s1f32 = l7
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l4
	s1i32 = l3
	s2i32 = 3376
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = 2
	f238(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3376)]))
	s0f32 = float32(math.Abs(float64(s0f32)))
	l7 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	l4 = s0i32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+3380)]))
	s1f32 = float32(math.Abs(float64(s1f32)))
	l8 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	l5 = s1i32
	s2f32 = l7
	s3f32 = l8
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l6 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	s1i32 = l5
	s2i32 = l4
	s3i32 = l6
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = s0f32 + s1f32
	s1f32 = 1
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3384)]))
	s0f32 = float32(math.Abs(float64(s0f32)))
	l7 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	l4 = s0i32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+3388)]))
	s1f32 = float32(math.Abs(float64(s1f32)))
	l8 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	l5 = s1i32
	s2f32 = l7
	s3f32 = l8
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l6 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0f32 = math.Float32frombits(uint32(s0i32))
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	s1i32 = l5
	s2i32 = l4
	s3i32 = l6
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = s0f32 + s1f32
	s1f32 = 1
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+44)])
	s1i32 = 192
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = 3424
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 3416
	s0i32 = s0i32 + s1i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 3408
	s0i32 = s0i32 + s1i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 3400
	s0i32 = s0i32 + s1i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 3392
	s0i32 = s0i32 + s1i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = 3384
	s0i32 = s0i32 + s1i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+3376)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	s2i32 = l3
	s3i32 = 3376
	s2i32 = s2i32 + s3i32
	s0i32 = f1879(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 3336
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	s2i32 = 3332
	s3i32 = 3332
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
	l4 = s0i32
	s0i32 = l3
	s1i32 = l0
	s2i32 = 0
	s3i32 = l2
	s0i32 = f182(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l3
	s2i32 = 3376
	s1i32 = s1i32 + s2i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+36)]))
	s4i32 = l3
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s0i32 = f998(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l5 = s0i32
	s0i32 = l4
	f43(ctx, s0i32)
	s0i32 = l5
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = l3
	s0i32 = f37(ctx, s0i32)
	l4 = s0i32
	s1i32 = l1
	f331(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l2
	s3i32 = 0
	s4i32 = 1
	f132(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l4
	f34(ctx, s0i32)
lbl0:
	s0i32 = l3
	s1i32 = 3440
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
