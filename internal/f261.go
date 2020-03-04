package internal

import (
	"math"
	"unsafe"
)

func f261(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float64
	_ = l16
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	s0i32 = ctx.g0
	s1i32 = 144
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = -1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 104
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = -1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 70364449226751
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = -70360154259455
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l4 = s0i32
		s0i64 = int64(s0i32)
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l5 = s1i32
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l10 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l7 = s0i32
		s0i64 = int64(s0i32)
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l8 = s1i32
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l11 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i64 = l10
		s1i64 = l11
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967295
		if uint64(s0i64) > uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l7
		s1i32 = 16383
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l4
		s1i32 = 16383
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l5
		s1i32 = -16383
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l8
		s1i32 = -16383
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	lbl3:
		s0i32 = l6
		s1i32 = l1
		s2i32 = l3
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = 1
		f268(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l6
		l1 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = -1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	lbl2:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+84)])
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			s1i32 = l4
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
			s2i32 = l4
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
			s0i32 = f28(ctx, s0i32, s1i32, s2i32)
			l5 = s0i32
			s0i32 = l4
			s1i32 = 0
			ctx.Mem[int(s0i32+84)] = uint8(s1i32)
			s0i32 = l4
			s1i32 = l5
			ctx.Mem[int(s0i32+85)] = uint8(s1i32)
		}
		s0i32 = l3
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
		l12 = s0f32
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
		l13 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl6
		}
		s0f32 = l12
		s1f32 = -5.368709e+08
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl6
		}
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l14 = s0f32
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+100)]))
		l15 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl6
		}
		s0f32 = l15
		s1f32 = 5.368709e+08
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl6
		}
		s0f32 = l13
		s1f32 = 5.368709e+08
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl6
		}
		s0i32 = 0
		s1f32 = l14
		s2f32 = -5.368709e+08
		if s1f32 >= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl5
		}
	lbl6:
		s0i32 = l3
		s1i64 = 5620492336267001856
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = -3602879698440290304
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = 88
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s0i32 = f135(ctx, s0i32, s1i32)
		if s0i32 != 0 {
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)]))
			l15 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
			l13 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
			l14 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
			l12 = s0f32
			s0i32 = 1
			goto lbl5
		}
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
		s0f32 = 0
		l15 = s0f32
		s0f32 = 0
		l13 = s0f32
		s0f32 = 0
		l14 = s0f32
		s0f32 = 0
		l12 = s0f32
		s0i32 = 1
	lbl5:
		l4 = s0i32
		s0i32 = l3
		s1f32 = l12
		s1f64 = float64(s1f32)
		s2f64 = -0.5234375
		s1f64 = s1f64 + s2f64
		s1f64 = math.Ceil(s1f64)
		l16 = s1f64
		s2f64 = 2.147483647e+09
		s3f64 = l16
		s4f64 = 2.147483647e+09
		if s3f64 < s4f64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f64 = s1f64
		} else {
			s1f64 = s2f64
		}
		l16 = s1f64
		s2f64 = -2.147483647e+09
		s3f64 = l16
		s4f64 = -2.147483647e+09
		if s3f64 > s4f64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f64 = s1f64
		} else {
			s1f64 = s2f64
		}
		l16 = s1f64
		s1f64 = math.Abs(s1f64)
		s2f64 = 2.147483648e+09
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f64 = l16
			s1i32 = int32(math.Trunc(s1f64))
			goto lbl8
		}
		s1i32 = -2147483648
	lbl8:
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
		s0i32 = l5
		s0i64 = int64(s0i32)
		l10 = s0i64
		s0i32 = l3
		s1f32 = l13
		s1f64 = float64(s1f32)
		s2f64 = 0.5234375
		s1f64 = s1f64 + s2f64
		s1f64 = math.Floor(s1f64)
		l16 = s1f64
		s2f64 = 2.147483647e+09
		s3f64 = l16
		s4f64 = 2.147483647e+09
		if s3f64 < s4f64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f64 = s1f64
		} else {
			s1f64 = s2f64
		}
		l16 = s1f64
		s2f64 = -2.147483647e+09
		s3f64 = l16
		s4f64 = -2.147483647e+09
		if s3f64 > s4f64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f64 = s1f64
		} else {
			s1f64 = s2f64
		}
		l16 = s1f64
		s1f64 = math.Abs(s1f64)
		s2f64 = 2.147483648e+09
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f64 = l16
			s1i32 = int32(math.Trunc(s1f64))
			goto lbl10
		}
		s1i32 = -2147483648
	lbl10:
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
		s0i32 = l5
		s0i64 = int64(s0i32)
		s1i64 = l10
		s0i64 = s0i64 - s1i64
		l10 = s0i64
		s0i32 = l3
		s1f32 = l14
		s1f64 = float64(s1f32)
		s2f64 = -0.5234375
		s1f64 = s1f64 + s2f64
		s1f64 = math.Ceil(s1f64)
		l16 = s1f64
		s2f64 = 2.147483647e+09
		s3f64 = l16
		s4f64 = 2.147483647e+09
		if s3f64 < s4f64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f64 = s1f64
		} else {
			s1f64 = s2f64
		}
		l16 = s1f64
		s2f64 = -2.147483647e+09
		s3f64 = l16
		s4f64 = -2.147483647e+09
		if s3f64 > s4f64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f64 = s1f64
		} else {
			s1f64 = s2f64
		}
		l16 = s1f64
		s1f64 = math.Abs(s1f64)
		s2f64 = 2.147483648e+09
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f64 = l16
			s1i32 = int32(math.Trunc(s1f64))
			goto lbl12
		}
		s1i32 = -2147483648
	lbl12:
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l3
		s1f32 = l15
		s1f64 = float64(s1f32)
		s2f64 = 0.5234375
		s1f64 = s1f64 + s2f64
		s1f64 = math.Floor(s1f64)
		l16 = s1f64
		s2f64 = 2.147483647e+09
		s3f64 = l16
		s4f64 = 2.147483647e+09
		if s3f64 < s4f64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f64 = s1f64
		} else {
			s1f64 = s2f64
		}
		l16 = s1f64
		s2f64 = -2.147483647e+09
		s3f64 = l16
		s4f64 = -2.147483647e+09
		if s3f64 > s4f64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f64 = s1f64
		} else {
			s1f64 = s2f64
		}
		l16 = s1f64
		s1f64 = math.Abs(s1f64)
		s2f64 = 2.147483648e+09
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f64 = l16
			s1i32 = int32(math.Trunc(s1f64))
			goto lbl14
		}
		s1i32 = -2147483648
	lbl14:
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
		s0i64 = l10
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		s0i32 = l7
		s0i64 = int64(s0i32)
		s1i32 = l5
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l11 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		s0i64 = l10
		s1i64 = l11
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967296
		if uint64(s0i64) < uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
	lbl17:
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+10)])
		s1i32 = 2
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l2
		s1i32 = l1
		f364(ctx, s0i32, s1i32)
		goto lbl1
	lbl16:
		s0i32 = l3
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = l1
		s3i32 = l3
		s4i32 = 72
		s3i32 = s3i32 + s4i32
		s4i32 = l0
		s4i32 = int32(ctx.Mem[int(s4i32+10)])
		s5i32 = 2
		s4i32 = s4i32 & s5i32
		s5i32 = 1
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s5i32 = l4
		s0i32 = f407(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		l2 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		l4 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+10)])
		s1i32 = 2
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l8 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
		s0i32 = l7
		s0i64 = int64(s0i32)
		s1i32 = l5
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l10 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
		s0i32 = l9
		s0i64 = int64(s0i32)
		s1i32 = l8
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l11 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
		s0i64 = l10
		s1i64 = l11
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967295
		if uint64(s0i64) > uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
		s0i32 = l4
		s1i32 = l3
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		f293(ctx, s0i32, s1i32, s2i32)
	lbl19:
		s0i32 = l0
		s1i32 = l1
		s2i32 = l4
		s3i32 = l3
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+76)]))
		s4i32 = l3
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+84)]))
		s5i32 = 0
		s6i32 = l2
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+56)]))
		if s6i32 == 0 {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		f409(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+10)])
		s1i32 = 2
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l0 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l8 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint32(s1i32)
		s0i32 = l5
		s0i64 = int64(s0i32)
		s1i32 = l0
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l10 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		s0i32 = l8
		s0i64 = int64(s0i32)
		s1i32 = l7
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l11 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		s0i64 = l10
		s1i64 = l11
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967295
		if uint64(s0i64) > uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		s0i32 = l4
		s1i32 = l3
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		f293(ctx, s0i32, s1i32, s2i32)
	lbl18:
		s0i32 = l2
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		s0i32 = f35(ctx, s0i32)
		s0i32 = l2
		s0i32 = f35(ctx, s0i32)
	lbl1:
		s0i32 = l6
		f40(ctx, s0i32)
	}
	s0i32 = l3
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
