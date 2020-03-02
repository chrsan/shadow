package internal

import (
	"unsafe"
)

func f679(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
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
	var l23 float32
	_ = l23
	var l24 float32
	_ = l24
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
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = 1
	l14 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l12 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l10
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l11 = s0i32
lbl1:
	s0i32 = l11
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s1i32 = 4
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l14 = s0i32
		goto lbl0
	}
	s0i32 = l11
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l8
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = l12
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl0:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l8 = s0i32
	s1i32 = 2
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 2
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l1
			s2i32 = l2
			s3i32 = l3
			s4i32 = 1
			s3i32 = s3i32 + s4i32
			s4i32 = 1
			s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
			f680(ctx, s0i32, s1i32, s2i32, s3i32)
			goto lbl3
		}
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		s3i32 = l3
		s4i32 = 1
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		f330(ctx, s0i32, s1i32, s2i32, s3i32)
		goto lbl3
	}
	s0i32 = l0
	s1i32 = l12
	s2i32 = 1
	if s1i32 >= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l10
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		l11 = s1i32
		s1i32 = 0
		l8 = s1i32
	lbl8:
		s1i32 = 2
		s2i32 = l11
		s2i32 = int32(ctx.Mem[int(s2i32+0)])
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s3i32 = 4
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl6
		}
		s1i32 = l11
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l11 = s1i32
		s1i32 = l8
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s2i32 = l12
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl8
		}
	}
	s1i32 = l2
lbl6:
	ctx.Mem[int(s0i32+9)] = uint8(s1i32)
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l16 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l17 = s1f32
	s2f32 = l16
	s3f32 = l17
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l8 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l20 = s0f32
	s0f32 = l17
	s1f32 = l16
	s2i32 = l8
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l21 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l17 = s1f32
	s2f32 = l16
	s3f32 = l17
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l8 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l22 = s0f32
	s0f32 = l17
	s1f32 = l16
	s2i32 = l8
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l23 = s0f32
	s0i32 = l10
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l10
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l8 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l12 = s0i32
		s0i32 = 1
		l11 = s0i32
		goto lbl9
	}
	s0i32 = l10
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l8 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l12 = s0i32
	s0i32 = l10
	s0i32 = int32(ctx.Mem[int(s0i32+85)])
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l11 = s0i32
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l8
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l10
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l16 = s0f32
	s1f32 = l20
	s2f32 = l20
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
	l20 = s0f32
	s0i32 = l10
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l16 = s0f32
	s1f32 = l22
	s2f32 = l22
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
	l22 = s0f32
	s0i32 = l10
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l16 = s0f32
	s1f32 = l21
	s2f32 = l16
	s3f32 = l21
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
	l21 = s0f32
	s0i32 = l10
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l16 = s0f32
	s1f32 = l23
	s2f32 = l16
	s3f32 = l23
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
	l23 = s0f32
	s0i32 = 0
	l11 = s0i32
lbl9:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+9)])
	l10 = s0i32
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 9
	s3i32 = 10
	s4i32 = l2
	if s4i32 == 0 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	s5i32 = l3
	s6i32 = 1
	s5i32 = s5i32 & s6i32
	if s5i32 == 0 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	s4i32 = s4i32 ^ s5i32
	l5 = s4i32
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l13 = s2i32
	s3i32 = l13
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l4
	s1i32 = 7
	s2i32 = 1
	s3i32 = l2
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l3
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l13 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l16 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l17 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l18 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l19 = s0f32
	s0i32 = l4
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l24 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = s1f32
	s0i32 = l4
	s1f32 = l19
	s2f32 = l18
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l18 = s0f32
	s0i32 = l4
	s1f32 = l24
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = s1f32
	s0i32 = l4
	s1f32 = l17
	s2f32 = l18
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = s1f32
	s0i32 = l4
	s1f32 = l24
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = s1f32
	s0i32 = l4
	s1f32 = l17
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = s1f32
	s0i32 = l4
	s1f32 = l16
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = s1f32
	s0i32 = l4
	s1f32 = l17
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = s1f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l18 = s0f32
	s0i32 = l4
	s1f32 = l16
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = s1f32
	s0i32 = l4
	s1f32 = l17
	s2f32 = l18
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = s1f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l18 = s0f32
	s0i32 = l4
	s1f32 = l16
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = s1f32
	s0i32 = l4
	s1f32 = l19
	s2f32 = l18
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = s1f32
	s0i32 = l4
	s1f32 = l16
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = s1f32
	s0i32 = l4
	s1f32 = l19
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = s1f32
	s0i32 = l4
	s1f32 = l24
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = s1f32
	s0i32 = l4
	s1f32 = l19
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = s1f32
	s0i32 = l4
	s1f32 = l16
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
	s0i32 = l4
	s1f32 = l16
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l4
	s1f32 = l24
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l4
	s1i32 = 3
	s2i32 = 1
	s3i32 = l2
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l3
	s2i32 = 1
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = l2
	s3i32 = 0
	if s2i32 != s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s1i32 = s1i32 + s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l4
	s1f32 = l19
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l4
	s1f32 = l17
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l4
	s1f32 = l17
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l4
	s1f32 = l24
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l4
	s1f32 = l19
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l15 = s0i64
	s0i32 = l4
	s1i32 = 120
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 0
	s3i32 = 0
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	l1 = s0i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 0
	s2f32 = 0
	s0i32 = f49(ctx, s0i32, s1i32, s2f32)
	s1i64 = l15
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 512
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l1 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l3 = s0i32
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		l5 = s1i32
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s2i32 = 7
		s1i32 = s1i32 & s2i32
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l4
		s2i32 = 48
		s1i32 = s1i32 + s2i32
		s2i32 = l7
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l7
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
		l7 = s0i32
		s0i32 = l4
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l3
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l16 = s0f32
		s0i32 = l9
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l17 = s0f32
		s0i32 = l4
		s1i32 = l5
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		s2i32 = 7
		s1i32 = s1i32 & s2i32
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = l7
		s1f32 = l16
		s2f32 = l17
		s3i32 = l4
		s4i32 = 48
		s3i32 = s3i32 + s4i32
		s4i32 = l5
		s5i32 = 3
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s3i32 = s3i32 + s4i32
		l5 = s3i32
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = l5
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		s5f32 = 0.70710677
		s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
		l5 = s0i32
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		l6 = s1i32
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		s2i32 = 7
		s1i32 = s1i32 & s2i32
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l4
		s2i32 = 48
		s1i32 = s1i32 + s2i32
		s2i32 = l9
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
		l5 = s0i32
		s0i32 = l4
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l3
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l16 = s0f32
		s0i32 = l9
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l17 = s0f32
		s0i32 = l4
		s1i32 = l6
		s2i32 = l7
		s1i32 = s1i32 + s2i32
		s2i32 = 7
		s1i32 = s1i32 & s2i32
		l6 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = l5
		s1f32 = l16
		s2f32 = l17
		s3i32 = l4
		s4i32 = 48
		s3i32 = s3i32 + s4i32
		s4i32 = l6
		s5i32 = 3
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s3i32 = s3i32 + s4i32
		l5 = s3i32
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = l5
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		s5f32 = 0.70710677
		s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
		l5 = s0i32
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		l6 = s1i32
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		s2i32 = 7
		s1i32 = s1i32 & s2i32
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l4
		s2i32 = 48
		s1i32 = s1i32 + s2i32
		s2i32 = l9
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
		l5 = s0i32
		s0i32 = l4
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l3
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l16 = s0f32
		s0i32 = l9
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l17 = s0f32
		s0i32 = l4
		s1i32 = l6
		s2i32 = l7
		s1i32 = s1i32 + s2i32
		s2i32 = 7
		s1i32 = s1i32 & s2i32
		l6 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = l5
		s1f32 = l16
		s2f32 = l17
		s3i32 = l4
		s4i32 = 48
		s3i32 = s3i32 + s4i32
		s4i32 = l6
		s5i32 = 3
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s3i32 = s3i32 + s4i32
		l5 = s3i32
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = l5
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		s5f32 = 0.70710677
		s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
		l5 = s0i32
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		l6 = s1i32
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		s2i32 = 7
		s1i32 = s1i32 & s2i32
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l4
		s2i32 = 48
		s1i32 = s1i32 + s2i32
		s2i32 = l9
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
		s0i32 = l6
		s1i32 = l7
		s0i32 = s0i32 + s1i32
		goto lbl11
	}
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l17 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	l5 = s1i32
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l0
	s1f32 = l16
	s2f32 = l17
	s3i32 = l4
	s4i32 = 48
	s3i32 = s3i32 + s4i32
	s4i32 = l7
	s5i32 = 3
	s4i32 = s4i32 << (uint32(s5i32) & 31)
	s3i32 = s3i32 + s4i32
	l7 = s3i32
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l7
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s5f32 = 0.70710677
	s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
	l7 = s0i32
	s0i32 = l4
	s1i32 = l5
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l4
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l5
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	l5 = s0i32
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l17 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	l6 = s1i32
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l9 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l5
	s1f32 = l16
	s2f32 = l17
	s3i32 = l4
	s4i32 = 48
	s3i32 = s3i32 + s4i32
	s4i32 = l9
	s5i32 = 3
	s4i32 = s4i32 << (uint32(s5i32) & 31)
	s3i32 = s3i32 + s4i32
	l5 = s3i32
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l5
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s5f32 = 0.70710677
	s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
	l5 = s0i32
	s0i32 = l4
	s1i32 = l6
	s2i32 = l7
	s1i32 = s1i32 + s2i32
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l5
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	l5 = s0i32
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l17 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	l6 = s1i32
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l9 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l5
	s1f32 = l16
	s2f32 = l17
	s3i32 = l4
	s4i32 = 48
	s3i32 = s3i32 + s4i32
	s4i32 = l9
	s5i32 = 3
	s4i32 = s4i32 << (uint32(s5i32) & 31)
	s3i32 = s3i32 + s4i32
	l5 = s3i32
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l5
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s5f32 = 0.70710677
	s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
	l5 = s0i32
	s0i32 = l4
	s1i32 = l6
	s2i32 = l7
	s1i32 = s1i32 + s2i32
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l5
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)]))
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
	s0i32 = s0i32 + s1i32
lbl11:
	l5 = s0i32
	s0i32 = l8
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l8 = s0i32
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l1 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l17 = s0f32
	s0i32 = l4
	s1i32 = l5
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l3 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l0
	s1f32 = l16
	s2f32 = l17
	s3i32 = l4
	s4i32 = 48
	s3i32 = s3i32 + s4i32
	s4i32 = l3
	s5i32 = 3
	s4i32 = s4i32 << (uint32(s5i32) & 31)
	s3i32 = s3i32 + s4i32
	l3 = s3i32
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l3
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s5f32 = 0.70710677
	s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
	s0i32 = l4
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l3
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	s0i32 = s0i32 + s1i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l4
	s1i32 = 120
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 0
	s3i32 = 0
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 5
	s2f32 = 0
	s0i32 = f49(ctx, s0i32, s1i32, s2f32)
lbl13:
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s2i32 = l1
	s3i32 = 31
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s1i32 = s1i32 ^ s2i32
	s2i32 = -1
	s1i32 = s1i32 ^ s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 120
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 0
	s3i32 = 0
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s1i32 = l13
	ctx.Mem[int(s0i32+89)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = l2
	s2i32 = 1
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	ctx.Mem[int(s0i32+88)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = l14
	ctx.Mem[int(s0i32+87)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l8
	ctx.Mem[int(s0i32+8)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l10
	ctx.Mem[int(s0i32+9)] = uint8(s1i32)
	s0i32 = 0
	s1i32 = l12
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l11
	s3i32 = 1
	s2i32 = s2i32 ^ s3i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l23
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1f32 = l21
	s0f32 = s0f32 * s1f32
	s1f32 = l22
	s0f32 = s0f32 * s1f32
	s1f32 = l20
	s0f32 = s0f32 * s1f32
	l16 = s0f32
	s1f32 = l16
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 0
	s3i32 = 0
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s1i32 = 256
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint16(s1i32)
	s0i32 = l0
	s1f32 = l23
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l0
	s1f32 = l20
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l0
	s1f32 = l22
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l0
	s1f32 = l21
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
lbl3:
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
