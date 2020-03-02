package internal

import (
	"unsafe"
)

func f936(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	s0i32 = l2
	s1i32 = 4
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		l5 = s0i32
		goto lbl0
	}
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 ^ s1i32
	l5 = s0i32
	s1i32 = 31
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l5
	s2i32 = 24
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s0i32 = s0i32 + s1i32
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s0i32 = l3
	s1i32 = 65280
	s0i32 = s0i32 & s1i32
	s1i32 = 128
	s0i32 = s0i32 | s1i32
	l8 = s0i32
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 128
	s0i32 = s0i32 | s1i32
	l9 = s0i32
	s0i32 = l3
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l5 = s0i32
	s1i32 = 65280
	s0i32 = s0i32 & s1i32
	s1i32 = 128
	s0i32 = s0i32 | s1i32
	l10 = s0i32
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 128
	s0i32 = s0i32 | s1i32
	l6 = s0i32
lbl2:
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+15)])
	l5 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+14)])
	l11 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+13)])
	l12 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	l13 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+11)])
	l14 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	l15 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+9)])
	l16 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+7)])
	l17 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+6)])
	l18 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+5)])
	l19 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+3)])
	l20 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+8)])
	l21 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+4)])
	l22 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	l23 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l24 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+1)])
	s2i32 = l4
	s1i32 = s1i32 * s2i32
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l24
	s1i32 = s1i32 * s2i32
	s2i32 = l9
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l23
	s1i32 = s1i32 * s2i32
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+2)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l20
	s1i32 = s1i32 * s2i32
	s2i32 = l10
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+3)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l22
	s1i32 = s1i32 * s2i32
	s2i32 = l9
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+4)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l19
	s1i32 = s1i32 * s2i32
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+5)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l18
	s1i32 = s1i32 * s2i32
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+6)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l17
	s1i32 = s1i32 * s2i32
	s2i32 = l10
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+7)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l21
	s1i32 = s1i32 * s2i32
	s2i32 = l9
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+8)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l16
	s1i32 = s1i32 * s2i32
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+9)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l15
	s1i32 = s1i32 * s2i32
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+10)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l14
	s1i32 = s1i32 * s2i32
	s2i32 = l10
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+11)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l13
	s1i32 = s1i32 * s2i32
	s2i32 = l9
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+12)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l12
	s1i32 = s1i32 * s2i32
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+13)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l11
	s1i32 = s1i32 * s2i32
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+14)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l5
	s1i32 = s1i32 * s2i32
	s2i32 = l10
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+15)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l2
	s1i32 = 7
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l11 = s0i32
	s0i32 = l2
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	l2 = s0i32
	s0i32 = l11
	if s0i32 != 0 {
		goto lbl2
	}
lbl0:
	s0i32 = l5
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = -1
		s0i32 = s0i32 ^ s1i32
		l2 = s0i32
		s1i32 = 31
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l2
		s2i32 = 24
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s0i32 = s0i32 + s1i32
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		l2 = s0i32
		s0i32 = l3
		s1i32 = 65280
		s0i32 = s0i32 & s1i32
		s1i32 = 128
		s0i32 = s0i32 | s1i32
		l4 = s0i32
		s0i32 = l3
		s1i32 = 8
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 128
		s0i32 = s0i32 | s1i32
		l8 = s0i32
		s0i32 = l3
		s1i32 = 16
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		l3 = s0i32
		s1i32 = 65280
		s0i32 = s0i32 & s1i32
		s1i32 = 128
		s0i32 = s0i32 | s1i32
		l9 = s0i32
		s0i32 = l3
		s1i32 = 8
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 128
		s0i32 = s0i32 | s1i32
		l10 = s0i32
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 | s1i32
		l3 = s0i32
	lbl4:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l7
		s1i32 = int32(ctx.Mem[int(s1i32+1)])
		s2i32 = l2
		s1i32 = s1i32 * s2i32
		s2i32 = l4
		s1i32 = s1i32 + s2i32
		s2i32 = 65280
		s1i32 = s1i32 & s2i32
		s2i32 = l7
		s2i32 = int32(ctx.Mem[int(s2i32+0)])
		s3i32 = l2
		s2i32 = s2i32 * s3i32
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		s3i32 = 65280
		s2i32 = s2i32 & s3i32
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 | s2i32
		s2i32 = l7
		s2i32 = int32(ctx.Mem[int(s2i32+3)])
		s3i32 = l2
		s2i32 = s2i32 * s3i32
		s3i32 = l9
		s2i32 = s2i32 + s3i32
		s3i32 = 65280
		s2i32 = s2i32 & s3i32
		s3i32 = l7
		s3i32 = int32(ctx.Mem[int(s3i32+2)])
		s4i32 = l2
		s3i32 = s3i32 * s4i32
		s4i32 = l10
		s3i32 = s3i32 + s4i32
		s4i32 = 65280
		s3i32 = s3i32 & s4i32
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s2i32 = s2i32 | s3i32
		s3i32 = 16
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l5
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l6 = s0i32
		s0i32 = l5
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l6
		if s0i32 != 0 {
			goto lbl4
		}
	}
}
