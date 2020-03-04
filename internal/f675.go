package internal

import (
	"math"
	"unsafe"
)

func f675(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int64
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
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
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = 1
	l3 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l4 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l5 = s0f32
	s0i64 = l4
	s1i64 = 32
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	s0i32 = int32(s0i64)
	s0f32 = math.Float32frombits(uint32(s0i32))
	l6 = s0f32
	s0i64 = l4
	s0i32 = int32(s0i64)
	s0f32 = math.Float32frombits(uint32(s0i32))
	l7 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l8 = s1f32
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1f32 = l5
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s1f32 = l6
	s2f32 = l5
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
	s0i32 = l0
	s1f32 = l7
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	goto lbl2
lbl3:
	s0i32 = l2
	s1f32 = l6
	s2f32 = l5
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l2
	s1f32 = l7
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l0
	s1i32 = l2
	s2i32 = 24
	s1i32 = s1i32 + s2i32
	s0i32 = f2087(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i32 = 5
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl8
	case 1:
		goto lbl8
	case 2:
		goto lbl2
	case 3:
		goto lbl5
	case 4:
		goto lbl6
	default:
		goto lbl7
	}
lbl8:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l3 = s0i32
	s1i32 = 5
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		s2i32 = 2
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		goto lbl4
	}
	s0i32 = l1
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l0
	s1i32 = 2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = 0
	l3 = s0i32
	goto lbl0
lbl7:
	s0i32 = 0
	l3 = s0i32
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+52)] = uint8(s1i32)
	goto lbl0
lbl6:
	s0i32 = l2
	s1i32 = 2139
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 4360
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 4298
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 4272
	s1i32 = l2
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl5:
	s0i32 = l0
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	l1 = s1i32
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = 0
	l3 = s0i32
	s0i32 = l1
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	goto lbl2
lbl4:
	s0i32 = l0
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
lbl2:
	s0i32 = l0
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = 1
	l3 = s0i32
lbl0:
	s0i32 = l2
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l3
	return s0i32
}
