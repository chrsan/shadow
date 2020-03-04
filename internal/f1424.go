package internal

import (
	"math"
	"unsafe"
)

func f1424(ctx *Context, l0 int32, l1 float64, l2 float64, l3 float64, l4 int32) float64 {
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
	var l10 float64
	_ = l10
	var l11 float64
	_ = l11
	var l12 float64
	_ = l12
	var l13 float64
	_ = l13
	var l14 float64
	_ = l14
	var l15 float64
	_ = l15
	var l16 float64
	_ = l16
	var l17 float64
	_ = l17
	var l18 float64
	_ = l18
	var l19 float64
	_ = l19
	var l20 float64
	_ = l20
	var l21 float64
	_ = l21
	var l22 float64
	_ = l22
	var l23 float64
	_ = l23
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
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s0f64 = l1
	s1f64 = l2
	s0f64 = s0f64 + s1f64
	s1f64 = 0.5
	s0f64 = s0f64 * s1f64
	l12 = s0f64
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		goto lbl0
	}
	s0f64 = l12
	s1f64 = 1
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		goto lbl0
	}
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l13 = s0f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l17 = s0f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l11 = s0f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f64
	s0i32 = l5
	s1f64 = 1
	s2f64 = l12
	s1f64 = s1f64 - s2f64
	l10 = s1f64
	s2f64 = l10
	s3f64 = l10
	s2f64 = s2f64 * s3f64
	l14 = s2f64
	s1f64 = s1f64 * s2f64
	l18 = s1f64
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1f64 = s1f64 * s2f64
	s2f64 = l12
	s3f64 = l14
	s4f64 = 3
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 * s3f64
	l14 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+24)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l12
	s3f64 = l12
	s2f64 = s2f64 * s3f64
	l15 = s2f64
	s3f64 = l10
	s4f64 = 3
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 * s3f64
	l10 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l12
	s3f64 = l15
	s2f64 = s2f64 * s3f64
	l15 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+56)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f64
	s0i32 = l5
	s1f64 = l18
	s2f64 = l16
	s1f64 = s1f64 * s2f64
	s2f64 = l14
	s3f64 = l11
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l10
	s3f64 = l17
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l15
	s3f64 = l13
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f64
lbl0:
	s0f64 = l12
	s1f64 = l1
	s0f64 = s0f64 - s1f64
	l17 = s0f64
	s0i32 = l0
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l4
	s1i32 = 3
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l4 = s0i32
	s1i32 = l5
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f64
	s1f64 = l3
	s0f64 = s0f64 - s1f64
	l13 = s0f64
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s0i32 = s0i32 + s1i32
	l4 = s0i32
lbl3:
	s0f64 = l12
	s1f64 = l17
	s2f64 = 0.5
	s1f64 = s1f64 * s2f64
	l17 = s1f64
	s0f64 = s0f64 - s1f64
	l10 = s0f64
	s1f64 = l1
	s2f64 = l10
	s3f64 = l1
	if s2f64 > s3f64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f64 = s0f64
	} else {
		s0f64 = s1f64
	}
	l10 = s0f64
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		goto lbl4
	}
	s0f64 = l10
	s1f64 = 1
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		goto lbl4
	}
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l14 = s0f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l18 = s0f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l15 = s0f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l22 = s0f64
	s0i32 = l5
	s1f64 = 1
	s2f64 = l10
	s1f64 = s1f64 - s2f64
	l11 = s1f64
	s2f64 = l11
	s3f64 = l11
	s2f64 = s2f64 * s3f64
	l19 = s2f64
	s1f64 = s1f64 * s2f64
	l20 = s1f64
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1f64 = s1f64 * s2f64
	s2f64 = l10
	s3f64 = l19
	s4f64 = 3
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 * s3f64
	l19 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+24)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l10
	s3f64 = l10
	s2f64 = s2f64 * s3f64
	l21 = s2f64
	s3f64 = l11
	s4f64 = 3
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 * s3f64
	l11 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l10
	s3f64 = l21
	s2f64 = s2f64 * s3f64
	l21 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+56)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f64
	s0i32 = l5
	s1f64 = l20
	s2f64 = l22
	s1f64 = s1f64 * s2f64
	s2f64 = l19
	s3f64 = l15
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l11
	s3f64 = l18
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l21
	s3f64 = l14
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f64
lbl4:
	s0i32 = l5
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l5
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	l14 = s1f64
	s0f64 = s0f64 - s1f64
	s0f64 = math.Abs(s0f64)
	s1f64 = 5.960464477539063e-08
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		s1i32 = l5
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		s0f64 = s0f64 - s1f64
		s0f64 = math.Abs(s0f64)
		s1f64 = 5.960464477539063e-08
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
	}
	s0i32 = l4
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f64 = l3
	s0f64 = s0f64 - s1f64
	l11 = s0f64
	s0f64 = l13
	s1f64 = 0
	if s0f64 > s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	l9 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l13
		s1f64 = l11
		if s0f64 > s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl13
		}
		goto lbl12
	}
	s0f64 = l13
	s1f64 = l11
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
lbl13:
	s0f64 = l12
	s1f64 = l17
	s0f64 = s0f64 + s1f64
	l10 = s0f64
	s1f64 = l2
	if s0f64 > s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0f64 = l10
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		goto lbl15
	}
	s0f64 = l10
	s1f64 = 1
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		goto lbl15
	}
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l18 = s0f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l15 = s0f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l22 = s0f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l19 = s0f64
	s0i32 = l5
	s1f64 = 1
	s2f64 = l10
	s1f64 = s1f64 - s2f64
	l11 = s1f64
	s2f64 = l11
	s3f64 = l11
	s2f64 = s2f64 * s3f64
	l20 = s2f64
	s1f64 = s1f64 * s2f64
	l21 = s1f64
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1f64 = s1f64 * s2f64
	s2f64 = l10
	s3f64 = l20
	s4f64 = 3
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 * s3f64
	l20 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+24)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l10
	s3f64 = l10
	s2f64 = s2f64 * s3f64
	l23 = s2f64
	s3f64 = l11
	s4f64 = 3
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 * s3f64
	l11 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l10
	s3f64 = l23
	s2f64 = s2f64 * s3f64
	l23 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+56)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
	s0i32 = l5
	s1f64 = l21
	s2f64 = l19
	s1f64 = s1f64 * s2f64
	s2f64 = l20
	s3f64 = l22
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l11
	s3f64 = l15
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l23
	s3f64 = l18
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
lbl15:
	s0i32 = l5
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f64 = l14
	s0f64 = s0f64 - s1f64
	s0f64 = math.Abs(s0f64)
	s1f64 = 5.960464477539063e-08
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1i32 = l5
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		s0f64 = s0f64 - s1f64
		s0f64 = math.Abs(s0f64)
		s1f64 = 5.960464477539063e-08
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
	}
	s0i32 = l8
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f64 = l3
	s0f64 = s0f64 - s1f64
	l11 = s0f64
	s0i32 = l9
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l13
		s1f64 = l11
		if s0f64 <= s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl12
		}
		goto lbl9
	}
	s0f64 = l13
	s1f64 = l11
	if s0f64 >= s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
lbl12:
	s0f64 = l10
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l13 = s0f64
		s0i32 = l0
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		goto lbl20
	}
	s0f64 = l10
	s1f64 = 1
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l13 = s0f64
		s0i32 = l0
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		goto lbl20
	}
	s0f64 = 1
	s1f64 = l10
	s0f64 = s0f64 - s1f64
	l12 = s0f64
	s1f64 = l12
	s2f64 = l12
	s1f64 = s1f64 * s2f64
	l13 = s1f64
	s0f64 = s0f64 * s1f64
	l11 = s0f64
	s1i32 = l0
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0f64 = s0f64 * s1f64
	s1f64 = l10
	s2f64 = l13
	s3f64 = 3
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 * s2f64
	l16 = s1f64
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	s1f64 = l10
	s2f64 = l10
	s1f64 = s1f64 * s2f64
	l13 = s1f64
	s2f64 = l12
	s3f64 = 3
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 * s2f64
	l12 = s1f64
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)]))
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	s1f64 = l10
	s2f64 = l13
	s1f64 = s1f64 * s2f64
	l14 = s1f64
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	l13 = s0f64
	s0f64 = l11
	s1i32 = l0
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f64 = s0f64 * s1f64
	s1f64 = l16
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	s1f64 = l12
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	s1f64 = l14
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
lbl20:
	l12 = s0f64
	s0i32 = l5
	s1f64 = l13
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f64
	s0i32 = l5
	s1f64 = l12
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f64
	s0i32 = l7
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f64
	s1f64 = l3
	s0f64 = s0f64 - s1f64
	l11 = s0f64
	l13 = s0f64
	s0f64 = l10
	l12 = s0f64
	goto lbl8
lbl10:
	s0f64 = -1
	l10 = s0f64
	goto lbl7
lbl9:
	s0f64 = l16
	s1f64 = l3
	s0f64 = s0f64 - s1f64
	l11 = s0f64
	s0f64 = l12
	l10 = s0f64
lbl8:
	s0f64 = l11
	s0f64 = math.Abs(s0f64)
	s1f64 = 1.1920928955078125e-07
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
lbl7:
	s0f64 = l10
	return s0f64
}
