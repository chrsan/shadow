package internal

import (
	"math"
	"unsafe"
)

func f1356(ctx *Context, l0 float64) float64 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int64
	_ = l3
	var l4 float64
	_ = l4
	var l5 float64
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	var s5f64 float64
	_ = s5f64
	var s6f64 float64
	_ = s6f64
	var s7f64 float64
	_ = s7f64
	var s8f64 float64
	_ = s8f64
	s0f64 = l0
	s0i64 = int64(math.Float64bits(s0f64))
	l3 = s0i64
	s1i64 = 63
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	s0i32 = int32(s0i64)
	l1 = s0i32
	s0f64 = l0
	s1i64 = l3
	s2i64 = 32
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	s2i32 = 2147483647
	s1i32 = s1i32 & s2i32
	l2 = s1i32
	s2i32 = 1082532651
	if uint32(s1i32) >= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i64 = l3
		s2i64 = 9223372036854775807
		s1i64 = s1i64 & s2i64
		s2i64 = 9218868437227405312
		if uint64(s1i64) > uint64(s2i64) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f64 = l0
			return s1f64
		}
		s1f64 = l0
		s2f64 = 709.782712893384
		if s1f64 > s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f64 = l0
			s2f64 = 8.98846567431158e+307
			s1f64 = s1f64 * s2f64
			return s1f64
		}
		s1f64 = l0
		s2f64 = -708.3964185322641
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 != 0 {
			goto lbl6
		}
		s1f64 = l0
		s2f64 = -745.1332191019411
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl6
		}
		goto lbl1
	}
	s1i32 = l2
	s2i32 = 1071001155
	if uint32(s1i32) < uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s1i32 = l2
	s2i32 = 1072734898
	if uint32(s1i32) < uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl5
	}
lbl6:
	s1f64 = l0
	s2f64 = 1.4426950408889634
	s1f64 = s1f64 * s2f64
	s2i32 = l1
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 19088
	s2i32 = s2i32 + s3i32
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f64 = s1f64 + s2f64
	l4 = s1f64
	s1f64 = math.Abs(s1f64)
	s2f64 = 2.147483648e+09
	if s1f64 < s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f64 = l4
		s1i32 = int32(math.Trunc(s1f64))
		goto lbl4
	}
	s1i32 = -2147483648
	goto lbl4
lbl5:
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	s2i32 = l1
	s1i32 = s1i32 - s2i32
lbl4:
	l1 = s1i32
	s1f64 = float64(s1i32)
	l4 = s1f64
	s2f64 = -0.6931471803691238
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	l0 = s0f64
	s1f64 = l4
	s2f64 = 1.9082149292705877e-10
	s1f64 = s1f64 * s2f64
	l5 = s1f64
	s0f64 = s0f64 - s1f64
	goto lbl2
lbl3:
	s0i32 = l2
	s1i32 = 1043333120
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	l1 = s0i32
	s0f64 = l0
lbl2:
	l4 = s0f64
	s0f64 = l0
	s1f64 = l4
	s2f64 = l4
	s3f64 = l4
	s4f64 = l4
	s3f64 = s3f64 * s4f64
	l0 = s3f64
	s4f64 = l0
	s5f64 = l0
	s6f64 = l0
	s7f64 = l0
	s8f64 = 4.1381367970572385e-08
	s7f64 = s7f64 * s8f64
	s8f64 = -1.6533902205465252e-06
	s7f64 = s7f64 + s8f64
	s6f64 = s6f64 * s7f64
	s7f64 = 6.613756321437934e-05
	s6f64 = s6f64 + s7f64
	s5f64 = s5f64 * s6f64
	s6f64 = -0.0027777777777015593
	s5f64 = s5f64 + s6f64
	s4f64 = s4f64 * s5f64
	s5f64 = 0.16666666666666602
	s4f64 = s4f64 + s5f64
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 - s3f64
	l0 = s2f64
	s1f64 = s1f64 * s2f64
	s2f64 = 2
	s3f64 = l0
	s2f64 = s2f64 - s3f64
	s1f64 = s1f64 / s2f64
	s2f64 = l5
	s1f64 = s1f64 - s2f64
	s0f64 = s0f64 + s1f64
	s1f64 = 1
	s0f64 = s0f64 + s1f64
	l4 = s0f64
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f64 = l4
	s1i32 = l1
	s0f64 = f250(ctx, s0f64, s1i32)
	l4 = s0f64
lbl1:
	s0f64 = l4
	return s0f64
lbl0:
	s0f64 = l0
	s1f64 = 1
	s0f64 = s0f64 + s1f64
	return s0f64
}
