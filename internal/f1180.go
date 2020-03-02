package internal

import (
	"unsafe"
)

func f1180(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l6 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l8 = s1i32
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l8
	s1i32 = s1i32 - s2i32
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l9 = s0i32
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		l5 = s0i32
	lbl1:
		s0i32 = l5
		l8 = s0i32
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l7
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		s1i32 = l10
		s2i32 = l9
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s0i32 = s0i32 + s1i32
		l5 = s0i32
	}
	s0i32 = l1
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s1i32 = l5
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l7 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		l8 = s0i32
		goto lbl2
	}
lbl4:
	s0i32 = l6
	s1i32 = l7
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s0i32 = l5
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	l7 = s0i32
	s0i32 = l5
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	l5 = s0i32
	s0i32 = l6
	s1i32 = l7
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
lbl2:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1076)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s1i32 = s1i32 - s2i32
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l9 = s1i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 2
		s1i32 = f44(ctx, s1i32, s2i32)
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1076)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s2i32 = l9
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	}
	s0i32 = l4
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = l6
		s0i32 = s0i32 - s1i32
		l7 = s0i32
		s0i32 = l5
		s1i32 = 16
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 16
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l6 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l11 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l9 = s0i32
	lbl7:
		s0i32 = l8
		s0i32 = int32(ctx.Mem[int(s0i32+1)])
		l13 = s0i32
	lbl9:
		s0i32 = l3
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l12 = s0i32
		s0i32 = l9
		s1i32 = l7
		s2i32 = l6
		s3i32 = l7
		s4i32 = l6
		if s3i32 < s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l10 = s1i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l11
		s1i32 = l12
		s2i32 = l13
		s1i32 = s1i32 * s2i32
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		l12 = s1i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = l12
		s1i32 = s1i32 + s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l9
		s1i32 = l10
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l6
		s1i32 = l10
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l5
			s2i32 = 16
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s2i32 = 16
			s1i32 = s1i32 >> (uint32(s2i32) & 31)
			l6 = s1i32
			s2i32 = 1
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			l4 = s0i32
			s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
			l5 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl8
			}
			s0i32 = l3
			s1i32 = l6
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l5
			s1i32 = 16
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s1i32 = 16
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l6 = s0i32
		}
		s0i32 = l10
		s1i32 = l11
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s0i32 = l7
		s1i32 = l10
		s0i32 = s0i32 - s1i32
		l7 = s0i32
		if s0i32 != 0 {
			goto lbl9
		}
		s0i32 = l8
		s0i32 = int32(ctx.Mem[int(s0i32+2)])
		l7 = s0i32
		s0i32 = l8
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		goto lbl7
	lbl8:
		s0i32 = l9
		s1i32 = 0
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l3 = s0i32
	s1i32 = l1
	s2i32 = l2
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+36)]))
	s5i32 = l3
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
}
