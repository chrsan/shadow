package internal

import (
	"unsafe"
)

func f138(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var l7 float32
	_ = l7
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
	s1i32 = 112
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l4 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = f24(ctx, s1i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		s1i32 = l2
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l2
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l4 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l1
		s0i32 = f97(ctx, s0i32)
		f12(ctx, s0i32)
	lbl3:
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l2
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = -4
		s1i32 = s1i32 & s2i32
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+10)])
		s3i32 = 3
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		l1 = s1i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = l1
		s2i32 = 251
		s1i32 = s1i32 & s2i32
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+10)])
		s3i32 = 4
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+8)])
		ctx.Mem[int(s0i32+8)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+9)])
		ctx.Mem[int(s0i32+9)] = uint8(s1i32)
		goto lbl1
	}
	s0i32 = l4
	s1i32 = 192
	s0i32 = s0i32 & s1i32
	s1i32 = 128
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = f67(ctx, s1i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l2
	s1i32 = l0
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l2 = s0i32
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = f137(ctx)
		l4 = s0i32
		s0i32 = l3
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint16(s1i32)
		s0i32 = l3
		s1i32 = -1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		ctx.Mem[int(s0i32+98)] = uint8(s1i32)
		s0i32 = f137(ctx)
		l4 = s0i32
		s0i32 = l3
		s1i32 = 0
		ctx.Mem[int(s0i32+82)] = uint8(s1i32)
		s0i32 = l3
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint16(s1i32)
		s0i32 = l3
		s1i32 = -1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l3
		s2i32 = 72
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s3i32 = l0
		s4i32 = l1
		s5i32 = l3
		s6i32 = 72
		s5i32 = s5i32 + s6i32
		s3i32 = f2090(ctx, s3i32, s4i32, s5i32)
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l0 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l4
		if s1i32 != 0 {
			s1i32 = l4
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)]))
			s1i32 = s1i32 + s2i32
		} else {
			s1i32 = 0
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l0 = s0i32
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s1i32 = 0
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l3
		s1i32 = 0
		ctx.Mem[int(s0i32+67)] = uint8(s1i32)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l0
		s2i32 = -4
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		s3i32 = l0
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	lbl7:
		s0i32 = l3
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s0i32 = f131(ctx, s0i32, s1i32)
		l0 = s0i32
		s1i32 = 6
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l0
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
		case 5:
			goto lbl8
		default:
			goto lbl14
		}
	lbl14:
		s0i32 = l3
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i64
		s0i32 = l3
		s1i32 = 104
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = 88
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)]))
		s1i32 = 0
		s2f32 = 0
		s0i32 = f49(ctx, s0i32, s1i32, s2f32)
		s1i64 = l6
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint16(s1i32)
		goto lbl7
	lbl13:
		s0i32 = l3
		s1i32 = 88
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l3
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
		goto lbl7
	lbl12:
		s0i32 = l3
		s1f32 = 1
		s2i32 = l1
		s0f32 = f443(ctx, s0i32, s1f32, s2i32)
		l7 = s0f32
		s0i32 = l3
		s1i32 = 88
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l3
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s3i32 = l3
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
		s4i32 = l3
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
		s5f32 = l7
		s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
		goto lbl7
	lbl11:
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l1
		s0f32 = f443(ctx, s0i32, s1f32, s2i32)
		l7 = s0f32
		s0i32 = l3
		s1i32 = 88
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l3
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s3i32 = l3
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
		s4i32 = l3
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
		s5f32 = l7
		s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
		goto lbl7
	lbl10:
		s0i32 = l3
		s1i32 = 88
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = 2
		f440(ctx, s0i32, s1i32, s2i32)
		goto lbl7
	lbl9:
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
		l0 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l4 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l4
		s1i32 = l0
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
			goto lbl15
		}
		s0i32 = l3
		s1i32 = 104
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = 88
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)]))
		s1i32 = 5
		s2f32 = 0
		s0i32 = f49(ctx, s0i32, s1i32, s2f32)
	lbl15:
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+92)]))
		l0 = s1i32
		s2i32 = l0
		s3i32 = 31
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s1i32 = s1i32 ^ s2i32
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		goto lbl7
	lbl8:
		s0i32 = l3
		s1i32 = 88
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l0 = s0i32
			s0i32 = l2
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
			s0i32 = l2
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l0 = s0i32
			s0i32 = l2
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+92)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
			s0i32 = l2
			s1i32 = l3
			s1i32 = int32(ctx.Mem[int(s1i32+98)])
			l0 = s1i32
			s2i32 = 3
			s1i32 = s1i32 & s2i32
			s2i32 = l2
			s2i32 = int32(ctx.Mem[int(s2i32+10)])
			l4 = s2i32
			s3i32 = 248
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 | s2i32
			s2i32 = l0
			s3i32 = 4
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+10)] = uint8(s1i32)
			s0i32 = l3
			s1i32 = l4
			s2i32 = 4
			s1i32 = s1i32 & s2i32
			s2i32 = l0
			s3i32 = 248
			s2i32 = s2i32 & s3i32
			s3i32 = l4
			s4i32 = 3
			s3i32 = s3i32 & s4i32
			s2i32 = s2i32 | s3i32
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+98)] = uint8(s1i32)
			s0i32 = l2
			s0i32 = int32(ctx.Mem[int(s0i32+8)])
			l0 = s0i32
			s0i32 = l2
			s1i32 = l3
			s1i32 = int32(ctx.Mem[int(s1i32+96)])
			ctx.Mem[int(s0i32+8)] = uint8(s1i32)
			s0i32 = l3
			s1i32 = l0
			ctx.Mem[int(s0i32+96)] = uint8(s1i32)
			s0i32 = l2
			s0i32 = int32(ctx.Mem[int(s0i32+9)])
			l0 = s0i32
			s0i32 = l2
			s1i32 = l3
			s1i32 = int32(ctx.Mem[int(s1i32+97)])
			ctx.Mem[int(s0i32+9)] = uint8(s1i32)
			s0i32 = l3
			s1i32 = l0
			ctx.Mem[int(s0i32+97)] = uint8(s1i32)
		}
		s0i32 = l3
		s1i32 = 104
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		s1i32 = 0
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+86)])) = uint16(s1i32)
		s0i32 = l1
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l1 = s1i32
		s2i32 = l1
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l2
		s1i32 = 2
		ctx.Mem[int(s0i32+9)] = uint8(s1i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
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
			goto lbl17
		}
		s0i32 = l0
		s0i32 = f97(ctx, s0i32)
		f12(ctx, s0i32)
	lbl17:
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
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
			goto lbl1
		}
		s0i32 = l0
		s0i32 = f97(ctx, s0i32)
		f12(ctx, s0i32)
		goto lbl1
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+8)])
	l4 = s0i32
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l1
	f2078(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = l2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l2
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = -4
		s1i32 = s1i32 & s2i32
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+10)])
		s3i32 = 3
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		l5 = s1i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = l5
		s2i32 = 251
		s1i32 = s1i32 & s2i32
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+10)])
		s3i32 = 4
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l5 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = f24(ctx, s1i32)
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l5
	s1i32 = 12
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+86)])
		s1i32 = l5
		s1i32 = int32(ctx.Mem[int(s1i32+87)])
		s0i32 = s0i32 | s1i32
		if s0i32 != 0 {
			goto lbl20
		}
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 0
		s2i32 = l3
		s3i32 = l3
		s4i32 = 88
		s3i32 = s3i32 + s4i32
		s4i32 = 0
		s5i32 = 0
		s6i32 = l3
		s7i32 = 32
		s6i32 = s6i32 + s7i32
		s0i32 = f332(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		if s0i32 != 0 {
			goto lbl20
		}
	}
	s0i32 = 0
	l4 = s0i32
lbl20:
	s0i32 = l2
	s1i32 = l4
	ctx.Mem[int(s0i32+8)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+9)])
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 2
		ctx.Mem[int(s0i32+9)] = uint8(s1i32)
		goto lbl1
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l7 = s0f32
	s1f32 = 0
	if s0f32 < s1f32 {
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
		s0i32 = l2
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+9)])
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 4348
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		ctx.Mem[int(s0i32+9)] = uint8(s1i32)
		goto lbl1
	}
	s0f32 = l7
	s1f32 = 0
	if s0f32 > s1f32 {
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
		s0i32 = l2
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+9)])
		ctx.Mem[int(s0i32+9)] = uint8(s1i32)
		goto lbl1
	}
	s0i32 = l2
	s1i32 = 2
	ctx.Mem[int(s0i32+9)] = uint8(s1i32)
lbl1:
	s0i32 = l3
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
