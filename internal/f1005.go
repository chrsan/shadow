package internal

import (
	"math"
	"unsafe"
)

func f1005(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int64
	_ = l14
	var l15 int64
	_ = l15
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
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
	s1i32 = 4752
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l5 = s0i32
	s1i32 = 3
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 2
		l5 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = 3
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l16 = s0f32
		s1f32 = -32767
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l19 = s0f32
		s1f32 = -32767
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l17 = s0f32
		s1f32 = 32767
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0f32 = l17
		s1f32 = l16
		s0f32 = s0f32 - s1f32
		s1f32 = 32767
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l18 = s0f32
		s1f32 = 32767
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0f32 = l18
		s1f32 = l19
		s0f32 = s0f32 - s1f32
		s1f32 = 32767
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l3
		s1f32 = l18
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l18 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l18
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl2
		}
		s1i32 = -2147483648
	lbl2:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint32(s1i32)
		s0i32 = l3
		s1f32 = l17
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l17 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l17
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
		l17 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l17
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
		l17 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l17
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl4
		}
		s1i32 = -2147483648
	lbl4:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint32(s1i32)
		s0i32 = l3
		s1f32 = l19
		s1f32 = float32(math.Floor(float64(s1f32)))
		l19 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l19
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
		l19 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l19
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
		l19 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l19
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl6
		}
		s1i32 = -2147483648
	lbl6:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = 4294967296
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+156)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint32(s1i32)
		s0i32 = l3
		s1f32 = l16
		s1f32 = float32(math.Floor(float64(s1f32)))
		l16 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l16
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
		l16 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l16
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
		l16 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l16
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl8
		}
		s1i32 = -2147483648
	lbl8:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l16 = s0f32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+16)])
		if s0i32 != 0 {
			s0i32 = l2
			s1f32 = l16
			s0f32 = f179(ctx, s0i32, s1f32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		}
		s0i32 = 0
		l5 = s0i32
		s0i32 = l0
		s1i32 = l3
		s2i32 = 104
		s1i32 = s1i32 + s2i32
		s2i32 = l3
		s3i32 = 136
		s2i32 = s2i32 + s3i32
		s3i32 = l2
		s4i32 = l3
		s5i32 = 168
		s4i32 = s4i32 + s5i32
		s5i32 = l0
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+32)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = 2
		l5 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l6 = s0i32
		s0f32 = math.Float32frombits(uint32(s0i32))
		l16 = s0f32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l7 = s1i32
		s1f32 = math.Float32frombits(uint32(s1i32))
		l19 = s1f32
		s2f32 = l19
		s3f32 = l16
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
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+168)]))
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s1f32 = float32(s1i32)
		l16 = s1f32
		s0f32 = s0f32 + s1f32
		l19 = s0f32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		l8 = s1i32
		s1f32 = math.Float32frombits(uint32(s1i32))
		l17 = s1f32
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
		l9 = s2i32
		s2f32 = math.Float32frombits(uint32(s2i32))
		l18 = s2f32
		s3f32 = l18
		s4f32 = l17
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
		s2f32 = l16
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 3
		s0f32 = s0f32 + s1f32
		l16 = s0f32
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1f32 = s1f32 - s2f32
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l10 = s0i32
		s0f32 = math.Float32frombits(uint32(s0i32))
		l17 = s0f32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l11 = s1i32
		s1f32 = math.Float32frombits(uint32(s1i32))
		l18 = s1f32
		s2f32 = l18
		s3f32 = l17
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
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+172)]))
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s1f32 = float32(s1i32)
		l17 = s1f32
		s0f32 = s0f32 + s1f32
		l18 = s0f32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		l12 = s1i32
		s1f32 = math.Float32frombits(uint32(s1i32))
		l20 = s1f32
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
		l13 = s2i32
		s2f32 = math.Float32frombits(uint32(s2i32))
		l21 = s2f32
		s3f32 = l21
		s4f32 = l20
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
		s2f32 = l17
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 3
		s0f32 = s0f32 + s1f32
		l17 = s0f32
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 - s2f32
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l3
		s1f32 = l17
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = s1f32
		s0i32 = l3
		s1f32 = l16
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = s1f32
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l13
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l12
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = 88
		s1i32 = s1i32 + s2i32
		s2i32 = l3
		f650(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l16 = s0f32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+16)])
		if s0i32 != 0 {
			s0i32 = l2
			s1f32 = l16
			s0f32 = f179(ctx, s0i32, s1f32)
		} else {
			s0f32 = l16
		}
		s1f32 = 128
		s0f32 = float32(math.Min(float64(s0f32), float64(s1f32)))
		l17 = s0f32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l3
		s3i32 = 32
		s2i32 = s2i32 + s3i32
		s3i32 = l4
		s0i32 = f1025(ctx, s0f32, s1i32, s2i32, s3i32)
		l1 = s0i32
		if s0i32 != 0 {
			goto lbl11
		}
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l16 = s0f32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+16)])
		if s0i32 != 0 {
			s0i32 = l2
			s1f32 = l16
			s0f32 = f179(ctx, s0i32, s1f32)
		} else {
			s0f32 = l16
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s0i32 = 1
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+216)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l3
			s2i32 = -64
			s1i32 = s1i32 - s2i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+208)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)]))
			l16 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)]))
			l20 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+188)]))
			l21 = s0f32
			s0i32 = l3
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+180)]))
			s1f32 = float32(math.Floor(float64(s1f32)))
			l22 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l22
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
			l22 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l22
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
			l22 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l22
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl16
			}
			s1i32 = -2147483648
		lbl16:
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint32(s1i32)
			s0i32 = l3
			s1f32 = l21
			s1f32 = float32(math.Ceil(float64(s1f32)))
			l21 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l21
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
			l21 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l21
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
			l21 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l21
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl18
			}
			s1i32 = -2147483648
		lbl18:
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint32(s1i32)
			s0i32 = l3
			s1f32 = l20
			s1f32 = float32(math.Floor(float64(s1f32)))
			l20 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l20
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
			l20 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l20
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
			l20 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l20
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl20
			}
			s1i32 = -2147483648
		lbl20:
			l1 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
			s0i32 = l3
			s1f32 = l16
			s1f32 = float32(math.Ceil(float64(s1f32)))
			l16 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l16
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
			l16 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l16
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
			l16 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l16
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl22
			}
			s1i32 = -2147483648
		lbl22:
			l5 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l5
			s2i32 = l1
			s1i32 = s1i32 - s2i32
			s2i32 = 3
			s1i32 = s1i32 + s2i32
			s2i32 = -4
			s1i32 = s1i32 & s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+156)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = 1
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l3
			s2i32 = 136
			s1i32 = s1i32 + s2i32
			s1i32 = f118(ctx, s1i32)
			s2i32 = 1
			s1i32 = f203(ctx, s1i32, s2i32)
			l1 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint32(s1i32)
			s0i32 = l1
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = 0
				l5 = s0i32
				goto lbl1
			}
			s0i32 = l3
			s1i32 = 4712
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l1
			s1i32 = 0
			ctx.Mem[int(s0i32+32)] = uint8(s1i32)
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
			l5 = s0i32
			s1i32 = l3
			s2i32 = 136
			s1i32 = s1i32 + s2i32
			f808(ctx, s0i32, s1i32)
			s0i32 = l3
			s1i32 = 280
			s0i32 = s0i32 + s1i32
			s1i32 = l5
			s0i32 = f408(ctx, s0i32, s1i32)
			l6 = s0i32
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
			s1f32 = float32(s1i32)
			s1f32 = -s1f32
			s2i32 = l3
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+144)]))
			s2f32 = float32(s2i32)
			s2f32 = -s2f32
			f170(ctx, s0i32, s1f32, s2f32)
			s0i32 = l3
			s1i32 = 232
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
			s0i32 = l1
			s1i32 = l1
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
			s2i32 = 1
			s1i32 = s1i32 | s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = l3
			s2i32 = 176
			s1i32 = s1i32 + s2i32
			s2i32 = l1
			f1481(ctx, s0i32, s1i32, s2i32)
			s0i32 = l1
			s0i32 = f23(ctx, s0i32)
			s0i32 = l6
			s0i32 = f259(ctx, s0i32)
			s0i32 = l5
			s0i32 = f41(ctx, s0i32)
			s0i32 = l3
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
			l1 = s0i32
			s0i32 = l0
			s1i32 = l4
			s2i32 = l3
			s3i32 = 136
			s2i32 = s2i32 + s3i32
			s3i32 = l2
			s4i32 = l3
			s5i32 = 168
			s4i32 = s4i32 + s5i32
			s5i32 = l0
			s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
			s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+32)]))
			if int(s5i32) < 0 || int(s5i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s5i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s5i32].numParams != 5 {
				panic("argument count mismatch")
			}
			s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
			l2 = s0i32
			s0i32 = l1
			if s0i32 != 0 {
				s0i32 = l1
				f13(ctx, s0i32)
			}
			s0i32 = 0
			l5 = s0i32
			s0i32 = l2
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl1
			}
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l0 = s0i32
		s0i32 = l4
		s0i32 = f487(ctx, s0i32)
		l2 = s0i32
		s0i32 = f651(ctx, s0i32)
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			l1 = s0i32
			goto lbl11
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l2
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		f13(ctx, s0i32)
		s0i32 = l4
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0f32 = l17
		s1i32 = l0
		s2i32 = l3
		s3i32 = 32
		s2i32 = s2i32 + s3i32
		s3i32 = l4
		s4i32 = l1
		f1023(ctx, s0f32, s1i32, s2i32, s3i32, s4i32)
	lbl11:
		s0i32 = l4
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])))
		l14 = s0i64
		s0i32 = l4
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])))
		l15 = s0i64
		s0i32 = l4
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l4
		s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])))
		s2i64 = l15
		s1i64 = s1i64 - s2i64
		l15 = s1i64
		s2i64 = 2147483647
		s3i64 = l15
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l15 = s1i64
		s2i64 = -2147483647
		s3i64 = l15
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i64)
		s0i32 = l4
		s1i32 = l4
		s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])))
		s2i64 = l14
		s1i64 = s1i64 - s2i64
		l14 = s1i64
		s2i64 = 2147483647
		s3i64 = l14
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l14 = s1i64
		s2i64 = -2147483647
		s3i64 = l14
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i64)
		s0i32 = l4
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+108)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = 1
		l5 = s0i32
		s0i32 = l4
		s1f32 = l18
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l16 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l16
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
		l16 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l16
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
		l16 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l16
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl28
		}
		s1i32 = -2147483648
	lbl28:
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l4
		s1f32 = l19
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l16 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l16
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
		l16 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l16
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
		l16 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l16
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl30
		}
		s1i32 = -2147483648
	lbl30:
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	lbl1:
		s0i32 = l3
		s1i32 = 4752
		s0i32 = s0i32 + s1i32
		ctx.g0 = s0i32
		s0i32 = l5
		return s0i32
	}
	s0i32 = l3
	s1i32 = 4752
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 4076
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	return s0i32
}
