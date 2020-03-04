package internal

import (
	"unsafe"
)

func f2092(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
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
	var l10 int32
	_ = l10
	var l11 int64
	_ = l11
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
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s0i64 int64
	_ = s0i64
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
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l0
	s2i32 = l0
	s3i32 = l1
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		s2i32 = l4
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = l0
		s2i32 = f677(ctx, s2i32, s3i32)
		l1 = s2i32
	}
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l2 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)]))
	s3i32 = l2
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		s0i32 = s0i32 + s1i32
		l5 = s0i32
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l2 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l3 = s0i32
	s0i32 = 0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	l6 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s0i32 = l6
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
lbl3:
	l8 = s0i32
	s0i32 = l5
	s1i32 = l7
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l2
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = 1
		l6 = s0i32
	lbl5:
		s0i32 = l5
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l9 = s0i32
		s1i32 = 4338
		s0i32 = s0i32 + s1i32
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l3 = s0i32
		s0i32 = l6
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l2 = s0i32
			goto lbl6
		}
		s0i32 = l1
		s1i32 = -8
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0i64
		s0i32 = l4
		s1i32 = 24
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		s1i32 = 0
		s2f32 = 0
		s0i32 = f49(ctx, s0i32, s1i32, s2f32)
		s1i64 = l11
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	lbl6:
		s0i32 = 0
		l6 = s0i32
		s0i32 = l2
		s1i32 = l3
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s0i32 = l9
		s1i32 = 5
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = 0
		s1i32 = l3
		s0i32 = s0i32 - s1i32
		l3 = s0i32
		s0i32 = l9
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl13
		case 1:
			goto lbl12
		case 2:
			goto lbl11
		case 3:
			goto lbl10
		case 4:
			goto lbl9
		default:
			goto lbl14
		}
	lbl14:
		s0i32 = l10
		if s0i32 != 0 {
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
				goto lbl16
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
				goto lbl16
			}
			s0i32 = l4
			s1i32 = 24
			s0i32 = s0i32 + s1i32
			s1i32 = l0
			s2i32 = 0
			s3i32 = 0
			s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
			s0i32 = l4
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
			s1i32 = 5
			s2f32 = 0
			s0i32 = f49(ctx, s0i32, s1i32, s2f32)
		lbl16:
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
		}
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = 0
		l10 = s0i32
		s0i32 = 1
		l6 = s0i32
		goto lbl8
	lbl13:
		s0i32 = l0
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l2
		s3i32 = l3
		s4i32 = 3
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
		goto lbl8
	lbl12:
		s0i32 = l0
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s3i32 = l1
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = l2
		s5i32 = l3
		s6i32 = 3
		s5i32 = s5i32 << (uint32(s6i32) & 31)
		s4i32 = s4i32 + s5i32
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		f130(ctx, s0i32, s1f32, s2f32, s3f32, s4f32)
		goto lbl8
	lbl11:
		s0i32 = l0
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s3i32 = l1
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = l2
		s5i32 = l3
		s6i32 = 3
		s5i32 = s5i32 << (uint32(s6i32) & 31)
		s4i32 = s4i32 + s5i32
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		s5i32 = l8
		s6i32 = -4
		s5i32 = s5i32 + s6i32
		l8 = s5i32
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
		goto lbl8
	lbl10:
		s0i32 = l0
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s3i32 = l1
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
		s4i32 = l1
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
		s5i32 = l1
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s6i32 = l2
		s7i32 = l3
		s8i32 = 3
		s7i32 = s7i32 << (uint32(s8i32) & 31)
		s6i32 = s6i32 + s7i32
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
		f175(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32, s6f32)
		goto lbl8
	lbl9:
		s0i32 = 1
		l10 = s0i32
	lbl8:
		s0i32 = l5
		s1i32 = l7
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		s0i32 = l0
		s0i32 = f97(ctx, s0i32)
		f12(ctx, s0i32)
	lbl18:
		s0i32 = l4
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
