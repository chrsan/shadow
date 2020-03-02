package internal

import (
	"math"
	"unsafe"
)

func f2085(ctx *Context, l0 int32, l1 int32) {
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
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
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
	var s7i32 int32
	_ = s7i32
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
	s1i32 = 144
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl0
	}
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+90)])
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
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
			l1 = s0i32
			s0i32 = l4
			s1i32 = 0
			ctx.Mem[int(s0i32+84)] = uint8(s1i32)
			s0i32 = l4
			s1i32 = l1
			ctx.Mem[int(s0i32+85)] = uint8(s1i32)
		}
		s0i32 = l0
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl0
	}
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l2
	l9 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	f275(ctx, s0i32, s1i32)
	s0i32 = 0
	l4 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l1 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l5 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
	}
	s0i32 = l5
	l7 = s0i32
	s0i32 = l5
	s0f32 = math.Float32frombits(uint32(s0i32))
	l11 = s0f32
	l12 = s0f32
	s0i32 = l4
	l8 = s0i32
	s0i32 = l4
	s0f32 = math.Float32frombits(uint32(s0i32))
	l13 = s0f32
	l14 = s0f32
lbl6:
	s0i32 = l9
	s1i32 = l2
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s0i32 = f276(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i32 = 6
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
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
		goto lbl6
	case 5:
		goto lbl7
	default:
		goto lbl14
	}
lbl14:
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = 1
	goto lbl8
lbl13:
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = 1
	goto lbl8
lbl12:
	s0i32 = 0
	l3 = s0i32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	s3i32 = l2
	s4i32 = 104
	s3i32 = s3i32 + s4i32
	s0i32 = f693(ctx, s0f32, s1f32, s2f32, s3i32)
	l1 = s0i32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s3i32 = l2
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+36)]))
	s4i32 = l2
	s5i32 = 104
	s4i32 = s4i32 + s5i32
	s5i32 = l1
	s6i32 = 2
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	s1i32 = f693(ctx, s1f32, s2f32, s3f32, s4i32)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl16:
		s0i32 = l2
		s1i32 = 112
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 104
		s2i32 = s2i32 + s3i32
		s3i32 = l3
		s4i32 = 2
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		f339(ctx, s0i32, s1i32, s2f32)
		s0i32 = l2
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = l1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
	}
	s0i32 = l2
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	goto lbl9
lbl11:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
	s0i32 = 0
	l3 = s0i32
	s0i32 = l2
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 104
	s1i32 = s1i32 + s2i32
	s0i32 = f2107(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i32 = l2
	s2i32 = 112
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = 104
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s4i32 = 2
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s1i32 = f2105(ctx, s1i32, s2i32)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	if s0i32 != 0 {
	lbl18:
		s0i32 = l2
		s1i32 = 96
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 112
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 104
		s2i32 = s2i32 + s3i32
		s3i32 = l3
		s4i32 = 2
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		f2108(ctx, s0i32, s1i32, s2f32)
		s0i32 = l2
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = l1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
	}
	s0i32 = l2
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	goto lbl8
lbl10:
	s0i32 = 0
	l3 = s0i32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	s3i32 = l2
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	s4i32 = l2
	s5i32 = 112
	s4i32 = s4i32 + s5i32
	s0i32 = f691(ctx, s0f32, s1f32, s2f32, s3f32, s4i32)
	l1 = s0i32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s3i32 = l2
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+36)]))
	s4i32 = l2
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+44)]))
	s5i32 = l2
	s6i32 = 112
	s5i32 = s5i32 + s6i32
	s6i32 = l1
	s7i32 = 2
	s6i32 = s6i32 << (uint32(s7i32) & 31)
	s5i32 = s5i32 + s6i32
	s1i32 = f691(ctx, s1f32, s2f32, s3f32, s4f32, s5i32)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl20:
		s0i32 = l2
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 112
		s1i32 = s1i32 + s2i32
		s2i32 = l3
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l2
		s3i32 = 48
		s2i32 = s2i32 + s3i32
		s3i32 = l3
		s4i32 = 3
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		s3i32 = 0
		f337(ctx, s0i32, s1f32, s2i32, s3i32)
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = l1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
	}
	s0i32 = l2
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
lbl9:
	s0i32 = l1
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
lbl8:
	l10 = s0i32
	s0i32 = 0
	l3 = s0i32
lbl21:
	s0i32 = l5
	s1i32 = l2
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s2f32 = l11
	s3i32 = l1
	s3f32 = math.Float32frombits(uint32(s3i32))
	l15 = s3f32
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l5 = s0i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l11 = s0f32
	s0i32 = l4
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	s2f32 = l13
	s3i32 = l6
	s3f32 = math.Float32frombits(uint32(s3i32))
	l16 = s3f32
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l4 = s0i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l13 = s0f32
	s0i32 = l7
	s1i32 = l1
	s2f32 = l12
	s3f32 = l15
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l7 = s0i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l12 = s0f32
	s0i32 = l8
	s1i32 = l6
	s2f32 = l14
	s3f32 = l16
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l8 = s0i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l14 = s0f32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l10
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	goto lbl6
lbl7:
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl0:
	s0i32 = l2
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
