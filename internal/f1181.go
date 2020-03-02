package internal

import (
	"unsafe"
)

func f1181(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l5 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l7
	s1i32 = s1i32 - s2i32
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0i32
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		l7 = s0i32
	lbl1:
		s0i32 = l7
		l4 = s0i32
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l6
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		s1i32 = l9
		s2i32 = l8
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s0i32 = s0i32 + s1i32
		l4 = s0i32
	}
	s0i32 = l1
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = l4
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l6 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		l7 = s0i32
		goto lbl2
	}
lbl4:
	s0i32 = l5
	s1i32 = l6
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	l6 = s0i32
	s0i32 = l4
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	l4 = s0i32
	s0i32 = l5
	s1i32 = l6
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
lbl2:
	s0i32 = l6
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = l3
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l7
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l4
	s1i32 = 255
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l0 = s0i32
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 4 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
	return
lbl6:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1076)]))
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l8 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l4 = s0i32
		goto lbl7
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s1i32 = s1i32 - s2i32
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 2
	s1i32 = f44(ctx, s1i32, s2i32)
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1076)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l4
	s2i32 = l6
	s3i32 = 1
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l8 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
lbl7:
	s0i32 = l4
	s1i32 = l3
	s2i32 = l5
	s3i32 = l5
	s4i32 = l3
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l5 = s1i32
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l8
	s1i32 = l7
	s1i32 = int32(ctx.Mem[int(s1i32+1)])
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = l5
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l3
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	if s0i32 != 0 {
	lbl10:
		s0i32 = l6
		s1i32 = l4
		s2i32 = l7
		s2i32 = int32(ctx.Mem[int(s2i32+2)])
		l3 = s2i32
		s3i32 = l4
		s4i32 = l3
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
		l3 = s1i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s1i32 = l7
		s1i32 = int32(ctx.Mem[int(s1i32+3)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l7
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l6
		s1i32 = l3
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l4
		s1i32 = l3
		l5 = s1i32
		s0i32 = s0i32 - s1i32
		l4 = s0i32
		if s0i32 != 0 {
			goto lbl10
		}
	}
	s0i32 = l6
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
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
lbl5:
}
