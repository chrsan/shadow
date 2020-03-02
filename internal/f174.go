package internal

import (
	"math"
	"unsafe"
)

func f174(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var s2i64 int64
	_ = s2i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0i32 = l0
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l8 = s1f32
	s2f32 = 0
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2f32 = 0
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2f32 = 0
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2f32 = 1
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l2 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l3 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l1 = s1i32
	s2i32 = 12
	s1i32 = s1i32 | s2i32
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l3
	s2i32 = s2i32 - s3i32
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l2
		s2i32 = 12
		s3i32 = 4
		f18(ctx, s1i32, s2i32, s3i32)
		s1i32 = 0
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l3 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l1 = s1i32
	}
	s1i32 = l2
	s2i32 = l1
	s3i32 = l3
	s2i32 = s2i32 + s3i32
	l1 = s2i32
	s3i32 = 12
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l1
	s2i64 = 19
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint64(s2i64)
	s1i32 = l1
	s2i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = 1
	goto lbl0
lbl1:
	s1f32 = l8
	s2f32 = 1
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2f32 = 1
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2f32 = 1
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2f32 = 1
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l2 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l3 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l1 = s1i32
	s2i32 = 12
	s1i32 = s1i32 | s2i32
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l3
	s2i32 = s2i32 - s3i32
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l2
		s2i32 = 12
		s3i32 = 4
		f18(ctx, s1i32, s2i32, s3i32)
		s1i32 = 0
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l3 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l1 = s1i32
	}
	s1i32 = l2
	s2i32 = l1
	s3i32 = l3
	s2i32 = s2i32 + s3i32
	l1 = s2i32
	s3i32 = 12
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l1
	s2i64 = 20
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint64(s2i64)
	s1i32 = l1
	s2i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = 1
	goto lbl0
lbl3:
	s1i32 = 0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l4 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l3 = s1i32
	s2i32 = 24
	s1i32 = s1i32 | s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l4
	s2i32 = s2i32 - s3i32
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l1
		s2i32 = 24
		s3i32 = 4
		f18(ctx, s1i32, s2i32, s3i32)
		s1i32 = 0
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l4 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l3 = s1i32
	}
	s1i32 = l1
	s2i32 = l3
	s3i32 = l4
	s2i32 = s2i32 + s3i32
	l1 = s2i32
	s3i32 = 24
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l1
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
	s1i32 = l1
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
	s1i32 = l1
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l3 = s1i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l5 = s1i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l6 = s1i32
	s1i32 = l1
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l7 = s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
	s1i32 = l1
	s2i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = l1
	s2i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l1
	s2i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l8 = s1f32
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl6
	}
	s1f32 = l8
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l8 = s2f32
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl6
	}
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l9 = s1f32
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl6
	}
	s1f32 = l9
	s2f32 = l8
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl6
	}
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l9 = s1f32
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl6
	}
	s1f32 = l9
	s2f32 = l8
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl6
	}
	s1i32 = l1
	s2i32 = l7
	s2f32 = math.Float32frombits(uint32(s2i32))
	s3f32 = 255
	s2f32 = s2f32 * s3f32
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	l8 = s2f32
	s3f32 = 4.2949673e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3f32 = l8
	s4f32 = 0
	if s3f32 >= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		s2f32 = l8
		s2i32 = int32(uint32(math.Trunc(float64(s2f32))))
		goto lbl7
	}
	s2i32 = 0
lbl7:
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+22)])) = uint16(s2i32)
	s1i32 = l1
	s2i32 = l6
	s2f32 = math.Float32frombits(uint32(s2i32))
	s3f32 = 255
	s2f32 = s2f32 * s3f32
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	l8 = s2f32
	s3f32 = 4.2949673e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3f32 = l8
	s4f32 = 0
	if s3f32 >= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		s2f32 = l8
		s2i32 = int32(uint32(math.Trunc(float64(s2f32))))
		goto lbl9
	}
	s2i32 = 0
lbl9:
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = uint16(s2i32)
	s1i32 = l1
	s2i32 = l5
	s2f32 = math.Float32frombits(uint32(s2i32))
	s3f32 = 255
	s2f32 = s2f32 * s3f32
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	l8 = s2f32
	s3f32 = 4.2949673e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3f32 = l8
	s4f32 = 0
	if s3f32 >= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		s2f32 = l8
		s2i32 = int32(uint32(math.Trunc(float64(s2f32))))
		goto lbl11
	}
	s2i32 = 0
lbl11:
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+18)])) = uint16(s2i32)
	s1i32 = l1
	s2i32 = l3
	s2f32 = math.Float32frombits(uint32(s2i32))
	s3f32 = 255
	s2f32 = s2f32 * s3f32
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	l8 = s2f32
	s3f32 = 4.2949673e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3f32 = l8
	s4f32 = 0
	if s3f32 >= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		s2f32 = l8
		s2i32 = int32(uint32(math.Trunc(float64(s2f32))))
		goto lbl13
	}
	s2i32 = 0
lbl13:
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint16(s2i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l5 = s1i32
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l3 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l2 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l4 = s1i32
	s2i32 = 12
	s1i32 = s1i32 | s2i32
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l2
	s2i32 = s2i32 - s3i32
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = 12
		s3i32 = 4
		f18(ctx, s1i32, s2i32, s3i32)
		s1i32 = 0
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l2 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l4 = s1i32
	}
	s1i32 = l3
	s2i32 = l2
	s3i32 = l4
	s2i32 = s2i32 + s3i32
	l2 = s2i32
	s3i32 = 12
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = 21
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = 2
	goto lbl0
lbl6:
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l6 = s1i32
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l3 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l2 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l5 = s1i32
	s2i32 = 12
	s1i32 = s1i32 | s2i32
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l2
	s2i32 = s2i32 - s3i32
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = 12
		s3i32 = 4
		f18(ctx, s1i32, s2i32, s3i32)
		s1i32 = 0
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l2 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l5 = s1i32
	}
	s1i32 = l3
	s2i32 = l2
	s3i32 = l5
	s2i32 = s2i32 + s3i32
	l2 = s2i32
	s3i32 = 12
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = 22
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = 2
	s2i32 = 1
	s3i32 = l4
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
lbl0:
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
}
