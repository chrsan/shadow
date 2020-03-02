package internal

import (
	"unsafe"
)

func f183(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
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
	s0i32 = ctx.g0
	s1i32 = 144
	s0i32 = s0i32 - s1i32
	l8 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l3 = s0i32
	s1i32 = 4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl1
	case 1:
		goto lbl1
	case 2:
		goto lbl1
	case 3:
		goto lbl0
	default:
		goto lbl2
	}
lbl2:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l3 = s1i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l9 = s2i32
	s1i32 = s1i32 - s2i32
	l10 = s1i32
	s2i32 = 3
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l11 = s1i32
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l4 = s2i32
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s2i32 = s2i32 - s3i32
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l4
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l7 = s0i32
	s0i32 = l3
	s1i32 = l9
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l7
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l6
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	l2 = s0i32
lbl4:
	s0i32 = l0
	s1i32 = l3
	s2i32 = l4
	s3i32 = l5
	s4i32 = 255
	s5i32 = l7
	s6i32 = l2
	s7i32 = -1
	s6i32 = s6i32 ^ s7i32
	s5i32 = s5i32 + s6i32
	l2 = s5i32
	s6i32 = 3
	s5i32 = s5i32 >> (uint32(s6i32) & 31)
	s6i32 = 1
	s5i32 = s5i32 + s6i32
	s6i32 = 32640
	s7i32 = l2
	s8i32 = 7
	s7i32 = s7i32 & s8i32
	s6i32 = int32(uint32(s6i32) >> (uint32(s7i32) & 31))
	s7i32 = 255
	s6i32 = s6i32 & s7i32
	f517(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	s0i32 = l6
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l5
	s1i32 = l11
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	goto lbl4
	panic("unreachable executed")
	panic("unreachable executed")
lbl3:
	s0i32 = l6
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	s1i32 = l3
	s2i32 = l10
	s3i32 = 7
	s2i32 = s2i32 & s3i32
	l2 = s2i32
	s1i32 = s1i32 - s2i32
	l9 = s1i32
	s2i32 = -1
	s1i32 = s1i32 ^ s2i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 3
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = 255
	s1i32 = l2
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l3 = s0i32
	s0i32 = 32640
	s1i32 = l1
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l2 = s0i32
lbl5:
	s0i32 = l0
	s1i32 = l9
	s2i32 = l4
	s3i32 = l5
	s4i32 = l3
	s5i32 = l10
	s6i32 = l2
	f517(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l5
	s1i32 = l11
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l6
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l1 = s0i32
	s0i32 = l6
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl5
	}
	goto lbl0
lbl1:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1i32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = 65
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = l3
		s2i32 = 2
		s1i32 = f50(ctx, s1i32, s2i32)
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		goto lbl6
	}
	s0i32 = l8
	s1i32 = l3
	if s1i32 != 0 {
		s1i32 = l8
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = 4
		s1i32 = s1i32 | s2i32
	} else {
		s1i32 = 0
	}
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
lbl6:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l12 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l11 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l9 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0i32
	s0i32 = l7
	s1i32 = 1
	s2i32 = l4
	s3i32 = 22120
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
	s0i32 = l7
	s1i32 = l4
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l10
		s0i32 = s0i32 + s1i32
		s1i32 = l9
		s0i32 = s0i32 - s1i32
		s1i32 = l3
		s2i32 = l12
		s1i32 = s1i32 - s2i32
		s2i32 = l11
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l5 = s0i32
	lbl10:
		s0i32 = l0
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l4
		s3i32 = l5
		s4i32 = l7
		s5i32 = l0
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l5
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l6
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l3 = s0i32
		s0i32 = l6
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l3
		if s0i32 != 0 {
			goto lbl10
		}
	}
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l0 = s0i32
	s1i32 = l8
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	f13(ctx, s0i32)
lbl0:
	s0i32 = l8
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
