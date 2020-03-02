package internal

import (
	"unsafe"
)

func f1485(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l2
		s1i32 = 4
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l2
		l5 = s0i32
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l0
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+288)])))
	l0 = s1i32
	s2i32 = l7
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l1 = s2i32
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l1
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l0
	s2i32 = s2i32 * s3i32
	s3i32 = -16711936
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s2i32 = l2
	s3i32 = 22124
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	return
lbl1:
	s0i32 = l0
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+288)])))
	l6 = s0i32
lbl3:
	s0i32 = l7
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l4 = s1i32
	s2i32 = 65535
	s1i32 = s1i32 & s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s0i32 = l7
	s1i32 = l4
	s2i32 = 14
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 262140
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = l7
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l9 = s1i32
	s2i32 = 65535
	s1i32 = s1i32 & s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s0i32 = l3
	s1i32 = l7
	s2i32 = l9
	s3i32 = 14
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 262140
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l9 = s1i32
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l6
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l9
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l6
	s2i32 = s2i32 * s3i32
	s3i32 = -16711936
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l8
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l6
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l8
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l6
	s2i32 = s2i32 * s3i32
	s3i32 = -16711936
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l4
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l6
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l4
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l6
	s2i32 = s2i32 * s3i32
	s3i32 = -16711936
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l5
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l6
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 16711935
	s1i32 = s1i32 & s2i32
	s2i32 = l5
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 16711935
	s2i32 = s2i32 & s3i32
	s3i32 = l6
	s2i32 = s2i32 * s3i32
	s3i32 = -16711936
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l1
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l2
	s1i32 = 7
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l8 = s0i32
	s0i32 = l2
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	l2 = s0i32
	s0i32 = l8
	if s0i32 != 0 {
		goto lbl3
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
		s0i32 = l0
		s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+288)])))
		l0 = s0i32
	lbl5:
		s0i32 = l3
		s1i32 = l7
		s2i32 = l4
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = 16711935
		s1i32 = s1i32 & s2i32
		s2i32 = l0
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 16711935
		s1i32 = s1i32 & s2i32
		s2i32 = l1
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 16711935
		s2i32 = s2i32 & s3i32
		s3i32 = l0
		s2i32 = s2i32 * s3i32
		s3i32 = -16711936
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l4
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l5
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l1 = s0i32
		s0i32 = l5
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l1
		if s0i32 != 0 {
			goto lbl5
		}
	}
}
