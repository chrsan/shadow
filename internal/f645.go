package internal

import (
	"math"
	"unsafe"
)

func f645(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 float32
	_ = l9
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
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l7 = s1i32
	if s1i32 != 0 {
		s1i32 = l7
	} else {
		s1i32 = f612(ctx)
	}
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l7 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l4
		s1i32 = f24(ctx, s1i32)
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l5
	s1i32 = l7
	s2i32 = 2
	s1i32 = s1i32 & s2i32
	if s1i32 != 0 {
		s1i32 = l5
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s3f32 = 1024
		s2f32 = s2f32 * s3f32
		s3f32 = 0.5
		s2f32 = s2f32 + s3f32
		s2f32 = float32(math.Floor(float64(s2f32)))
		s3f32 = 0.0009765625
		s2f32 = s2f32 * s3f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = s2f32
		s1i32 = 1
		l8 = s1i32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2f32 = 1024
		s1f32 = s1f32 * s2f32
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		s2f32 = 0.0009765625
		s1f32 = s1f32 * s2f32
		goto lbl2
	}
	s1i32 = l5
	s2i32 = 1065353216
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
	s1f32 = 1
lbl2:
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l5
	s1i32 = l7
	s2i32 = 4
	s1i32 = s1i32 & s2i32
	if s1i32 != 0 {
		s1i32 = l5
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s3f32 = 1024
		s2f32 = s2f32 * s3f32
		s3f32 = 0.5
		s2f32 = s2f32 + s3f32
		s2f32 = float32(math.Floor(float64(s2f32)))
		s3f32 = 0.0009765625
		s2f32 = s2f32 * s3f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = s2f32
		s1i32 = 1
		l8 = s1i32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2f32 = 1024
		s1f32 = s1f32 * s2f32
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s1f32 = float32(math.Floor(float64(s1f32)))
		s2f32 = 0.0009765625
		s1f32 = s1f32 * s2f32
		goto lbl4
	}
	s1i32 = l5
	s2i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = uint32(s2i32)
	s1f32 = 0
lbl4:
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	s1i32 = 6
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l9 = s0f32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+16)])
	s1i32 = 16
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l9
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s1f32 = f1870(ctx, s1f32)
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 * s2f32
		l9 = s1f32
		s0f32 = s0f32 + s1f32
		s1f32 = l9
		s2i32 = l4
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l9 = s0f32
		s0i32 = l4
		s1i32 = 2
		s2i32 = l4
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
	}
	s0f32 = l9
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l5
	s1f32 = l9
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l5
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+44)])
	s2i32 = 4
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l7 = s1i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	s2i32 = l5
	s2i32 = int32(ctx.Mem[int(s2i32+49)])
	s3i32 = 240
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+49)] = uint8(s1i32)
	s0i32 = l5
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+44)])
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l7
	s1i32 = s1i32 | s2i32
	s2i32 = 51
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+49)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	goto lbl7
lbl8:
	s0i32 = l5
	s1i32 = 0
	ctx.Mem[int(s0i32+49)] = uint8(s1i32)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = 0
lbl7:
	l7 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+17)])
	l4 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl12
	case 1:
		goto lbl10
	default:
		goto lbl11
	}
lbl12:
	s0i32 = 1
	l4 = s0i32
lbl11:
	s0i32 = l5
	s1i32 = l4
	ctx.Mem[int(s0i32+48)] = uint8(s1i32)
	goto lbl9
lbl10:
	s0i32 = l5
	s1i32 = 4
	ctx.Mem[int(s0i32+48)] = uint8(s1i32)
	s0i32 = l8
	if s0i32 != 0 {
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s0f32 = s0f32 * s1f32
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 - s1f32
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l9 = s1f32
		s2f32 = l9
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 * s1f32
		s1f32 = 2304
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		goto lbl13
	}
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = 48
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl13
	}
lbl14:
	s0i32 = l5
	s1i32 = 1
	ctx.Mem[int(s0i32+48)] = uint8(s1i32)
	s0i32 = l7
	s1i32 = 2048
	s0i32 = s0i32 | s1i32
	l7 = s0i32
	goto lbl9
lbl13:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	s1i32 = 4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl9
	case 1:
		goto lbl18
	case 2:
		goto lbl17
	case 3:
		goto lbl16
	default:
		goto lbl19
	}
lbl19:
	s0i32 = l5
	s1i32 = 1
	ctx.Mem[int(s0i32+48)] = uint8(s1i32)
	s0i32 = l7
	s1i32 = 2048
	s0i32 = s0i32 | s1i32
	l7 = s0i32
	goto lbl9
lbl18:
	s0i32 = l7
	s1i32 = 1024
	s0i32 = s0i32 | s1i32
	l7 = s0i32
	goto lbl9
lbl17:
	s0i32 = l7
	s1i32 = 512
	s0i32 = s0i32 | s1i32
	l7 = s0i32
	goto lbl9
lbl16:
	s0i32 = l7
	s1i32 = 1536
	s0i32 = s0i32 | s1i32
	l7 = s0i32
lbl9:
	s0i32 = l5
	s1i32 = l7
	s2i32 = l0
	s2i32 = int32(ctx.Mem[int(s2i32+16)])
	l2 = s2i32
	s3i32 = 1
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 4
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s2i32 = l2
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 16
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s2i32 = l2
	s3i32 = 5
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 32
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s2i32 = l2
	s3i32 = 8
	s2i32 = s2i32 & s3i32
	s3i32 = 9
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	s2i32 = l2
	s3i32 = 32
	s2i32 = s2i32 & s3i32
	s3i32 = 8
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	s2i32 = 15999
	s1i32 = s1i32 & s2i32
	s2i32 = l0
	s2i32 = int32(ctx.Mem[int(s2i32+18)])
	s3i32 = 7
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+50)])) = uint16(s1i32)
	s0i32 = l1
	s0i32 = f2119(ctx, s0i32)
	l0 = s0i32
	s0i32 = l5
	s1i32 = 128
	ctx.Mem[int(s0i32+46)] = uint8(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint16(s1i32)
	s0i32 = l5
	s1i32 = l0
	s2i32 = 5
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l2 = s1i32
	s2i32 = 5
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l2
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	s2i32 = l2
	s3i32 = 1
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 | s2i32
	s2i32 = l0
	s3i32 = 13
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 7
	s2i32 = s2i32 & s3i32
	l2 = s2i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = l2
	s4i32 = 5
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 | s3i32
	s3i32 = l2
	s4i32 = 1
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 | s3i32
	s3i32 = 8
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	s2i32 = l0
	s3i32 = 21
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 7
	s2i32 = s2i32 & s3i32
	l0 = s2i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = l0
	s4i32 = 5
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 | s3i32
	s3i32 = l0
	s4i32 = 1
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 | s3i32
	s3i32 = 16
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	s2i32 = -16777216
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = -16777216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 16448
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint16(s1i32)
	}
	s0i32 = l3
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 0
		ctx.Mem[int(s0i32+46)] = uint8(s1i32)
	}
	s0i32 = l6
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
}
