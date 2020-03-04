package internal

import (
	"math"
	"unsafe"
)

func f1935(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 float32
	_ = l6
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
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = -2
		s0i32 = s0i32 + s1i32
		s1i32 = 4
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l5
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = 5
		l1 = s0i32
		goto lbl1
	}
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = 5
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l4 = s0i32
		s0i32 = l1
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = l4
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l5
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l5
	s2i32 = l5
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l6 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l6
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l6
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l6
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl4
	}
	s1i32 = -2147483648
lbl4:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l1 = s0i32
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l6 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l6
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l6
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l6
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl6
	}
	s1i32 = -2147483648
lbl6:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 4676
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1f32 = float32(math.Floor(float64(s1f32)))
	l6 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l6
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l6
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l6
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl8
	}
	s1i32 = -2147483648
lbl8:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1f32 = float32(math.Floor(float64(s1f32)))
	l6 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l6
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l6
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l6
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl10
	}
	s1i32 = -2147483648
lbl10:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
lbl1:
	s0i32 = l0
	s1i32 = l5
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	f436(ctx, s0i32, s1i32, s2i32)
lbl0:
	s0i32 = l5
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
