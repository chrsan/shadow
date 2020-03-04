package internal

import (
	"unsafe"
)

func f681(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l12 int64
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
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
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	l7 = s1i32
	s2i32 = 1
	if s1i32 >= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		l6 = s1i32
	lbl2:
		s1i32 = 2
		s2i32 = l6
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
			goto lbl0
		}
		s1i32 = l6
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s1i32 = l8
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s2i32 = l7
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl2
		}
	}
	s1i32 = l2
lbl0:
	l10 = s1i32
	ctx.Mem[int(s0i32+9)] = uint8(s1i32)
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l14 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l15 = s1f32
	s2f32 = l14
	s3f32 = l15
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l6 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l17 = s0f32
	s0f32 = l15
	s1f32 = l14
	s2i32 = l6
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l14 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l16 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l13 = s1f32
	s2f32 = l13
	s3f32 = l16
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l6 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l15 = s0f32
	s0f32 = l13
	s1f32 = l16
	s2i32 = l6
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l16 = s0f32
	s0i32 = l5
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l6 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l8 = s0i32
		s0i32 = 1
		l7 = s0i32
		goto lbl3
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l6 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l8 = s0i32
	s0i32 = l5
	s0i32 = int32(ctx.Mem[int(s0i32+85)])
	l9 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l9
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l13 = s0f32
	s1f32 = l17
	s2f32 = l17
	s3f32 = l13
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
	l17 = s0f32
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l13 = s0f32
	s1f32 = l16
	s2f32 = l16
	s3f32 = l13
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
	l16 = s0f32
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l13 = s0f32
	s1f32 = l14
	s2f32 = l13
	s3f32 = l14
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
	l14 = s0f32
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l13 = s0f32
	s1f32 = l15
	s2f32 = l13
	s3f32 = l15
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
	l15 = s0f32
	s0i32 = 0
	l7 = s0i32
lbl3:
	s0i32 = l4
	s1i32 = l0
	s2i32 = 5
	s3i32 = 5
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l4
	s1i32 = 3
	s2i32 = 1
	s3i32 = l2
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s0i32 = l4
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l11 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0i32
	s0i32 = l4
	s1i32 = l11
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l3
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l1 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0i64
	s0i32 = l4
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 0
	s3i32 = 0
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	l3 = s0i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 0
	s2f32 = 0
	s0i32 = f49(ctx, s0i32, s1i32, s2f32)
	s1i64 = l12
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 512
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l1
	s3i32 = l2
	s2i32 = s2i32 + s3i32
	l1 = s2i32
	s3i32 = 3
	s2i32 = s2i32 & s3i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s1i32 = l4
	s2i32 = l1
	s3i32 = l2
	s2i32 = s2i32 + s3i32
	l1 = s2i32
	s3i32 = 3
	s2i32 = s2i32 & s3i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	l3 = s0i32
	s0i32 = l4
	s1i32 = l1
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l4
	s2i32 = l1
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l6
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l3
	s1i32 = l2
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
		goto lbl5
	}
	s0i32 = l4
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 0
	s3i32 = 0
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 5
	s2f32 = 0
	s0i32 = f49(ctx, s0i32, s1i32, s2f32)
lbl5:
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l2 = s1i32
	s2i32 = l2
	s3i32 = 31
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s1i32 = s1i32 ^ s2i32
	s2i32 = -1
	s1i32 = s1i32 ^ s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+8)] = uint8(s1i32)
	s0i32 = 0
	s1i32 = l8
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l7
	s3i32 = 1
	s2i32 = s2i32 ^ s3i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0f32 = l17
	s1f32 = l16
	s2f32 = l14
	s3f32 = l15
	s4f32 = 0
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s1f32 = l13
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l4
	s1i32 = l0
	s2i32 = 0
	s3i32 = 0
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s1i32 = 256
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint16(s1i32)
	s0i32 = l1
	s1f32 = l15
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l1
	s1f32 = l17
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l1
	s1f32 = l16
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l1
	s1f32 = l14
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
lbl6:
	s0i32 = l0
	s1i32 = l10
	ctx.Mem[int(s0i32+9)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
