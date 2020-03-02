package internal

import (
	"math"
	"unsafe"
)

func f1504(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int64
	_ = l8
	var l9 int64
	_ = l9
	var l10 int64
	_ = l10
	var l11 int64
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
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s1f32 = float32(s1i32)
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s2i32 = l4
	s2f32 = float32(s2i32)
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	s3i32 = l5
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+260)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 4 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, float32, float32, int32))(table[s4i32].f()))(ctx, s0i32, s1f32, s2f32, s3i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		s1f32 = 0
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l6 = s0i32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		s1f32 = 0
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)]))
	s1i32 = 1
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+284)]))
	s1i32 = 1
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
lbl0:
	l4 = s0i32
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = 4.2949673e+09
	s0f32 = s0f32 * s1f32
	l12 = s0f32
	s1f32 = 9.2233715e+18
	s2f32 = l12
	s3f32 = 9.2233715e+18
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l12 = s0f32
	s1f32 = -9.2233715e+18
	s2f32 = l12
	s3f32 = -9.2233715e+18
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l12 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 9.223372e+18
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l12
		s0i64 = int64(math.Trunc(float64(s0f32)))
		goto lbl2
	}
	s0i64 = -9223372036854775808
lbl2:
	l8 = s0i64
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = l8
	s2i32 = l4
	s2i64 = int64(uint32(s2i32))
	s3i64 = 16
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 - s2i64
	s2i64 = 16
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	l4 = s1i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l5
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 12
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = l7
	s5i32 = l3
	if s4i32 < s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l4
	s4i32 = 0
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0f32 = l12
	s1f32 = 4.2949673e+09
	s0f32 = s0f32 * s1f32
	l12 = s0f32
	s1f32 = 9.2233715e+18
	s2f32 = l12
	s3f32 = 9.2233715e+18
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l12 = s0f32
	s1f32 = -9.2233715e+18
	s2f32 = l12
	s3f32 = -9.2233715e+18
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l12 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 9.223372e+18
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l12
		s0i64 = int64(math.Trunc(float64(s0f32)))
		goto lbl4
	}
	s0i64 = -9223372036854775808
lbl4:
	l8 = s0i64
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 0
		s2i32 = l2
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s0i32 = f27(ctx, s0i32, s1i32, s2i32)
		goto lbl6
	}
	s0i32 = l0
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+264)]))
	l10 = s0i64
	s0i64 = l8
	s1i32 = l6
	s1i64 = int64(s1i32)
	s2i64 = 16
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	s0i64 = s0i64 - s1i64
	l8 = s0i64
	s1i64 = 16
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	l9 = s0i64
	s0i32 = int32(s0i64)
	l0 = s0i32
	s1i32 = 16
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l3
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i64 = l10
	s1i64 = 16
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	l11 = s0i64
	s0i32 = int32(s0i64)
	l4 = s0i32
	s1i32 = 257
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l2
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s0i64 = int64(s0i32)
	s1i64 = l11
	s2i64 = 4294967295
	s1i64 = s1i64 & s2i64
	s0i64 = s0i64 * s1i64
	s1i64 = l9
	s2i64 = 32
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	s2i64 = 32
	s1i64 = s1i64 >> (uint64(s2i64) & 63)
	s0i64 = s0i64 + s1i64
	l9 = s0i64
	s1i64 = 2147483647
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i64 = l9
	s0i32 = int32(s0i64)
	s1i32 = 16
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l3
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
lbl11:
	s0i32 = l2
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l2
	l0 = s0i32
	goto lbl8
lbl10:
	s0i32 = l2
	s1i32 = 3
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		l3 = s0i32
		goto lbl12
	}
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l6 = s0i32
lbl14:
	s0i32 = l1
	s1i32 = l0
	s2i32 = l4
	s1i32 = s1i32 + s2i32
	s2i32 = -65536
	s1i32 = s1i32 & s2i32
	s2i32 = l0
	s3i32 = 16
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l2
	s1i32 = 4
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l2
	s1i32 = -2
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	l2 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl14
	}
lbl12:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 1
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s2i32 = l4
		s1i32 = s1i32 + s2i32
		s2i32 = 16
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
	}
	s0i32 = l1
	s1i32 = l0
	s2i32 = 16
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	goto lbl6
lbl9:
	s0i64 = l10
	s1i64 = 1
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	l9 = s0i64
lbl16:
	s0i32 = l5
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = l8
	s2i64 = 16
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	l0 = s1i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = l4
	s4i32 = l3
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l0
	s3i32 = 0
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l5
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = l8
	s2i64 = l10
	s1i64 = s1i64 + s2i64
	s2i64 = 16
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	l4 = s1i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l0
	s2i32 = l5
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 12
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = l6
	s6i32 = l3
	if s5i32 < s6i32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	s4i32 = l4
	s5i32 = 0
	if s4i32 < s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3i32 = 16
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i64 = l8
	s1i64 = l9
	s0i64 = s0i64 + s1i64
	l8 = s0i64
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l2
	s1i32 = 3
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l4 = s0i32
	s0i32 = l2
	s1i32 = -2
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	l2 = s0i32
	s0i32 = l4
	if s0i32 != 0 {
		goto lbl16
	}
lbl8:
	s0i32 = l0
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
lbl17:
	s0i32 = l5
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = l8
	s2i64 = 16
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	l2 = s1i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l5
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 12
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = l4
	s5i32 = l3
	if s4i32 < s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l2
	s4i32 = 0
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i64 = l8
	s1i64 = l10
	s0i64 = s0i64 + s1i64
	l8 = s0i64
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
	s0i32 = l0
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l2
	if s0i32 != 0 {
		goto lbl17
	}
lbl6:
	s0i32 = l5
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
