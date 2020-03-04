package internal

import (
	"unsafe"
)

func f1480(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s6i32 int32
	_ = s6i32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l7 = s0i32
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l8 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)]))
	l9 = s0i32
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+284)]))
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = 0
	l2 = s0i32
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l5
	s3i32 = 4
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = 12
	s4i32 = s4i32 + s5i32
	s5i32 = l6
	s6i32 = l8
	if s5i32 < s6i32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	s4i32 = l6
	s5i32 = 0
	if s4i32 < s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l1
	s1i32 = l9
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i32 = -1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		l2 = s0i32
		goto lbl1
	}
	s0i32 = l3
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l4
	s3i32 = 0
	s4i32 = l0
	s3i32 = s3i32 - s4i32
	l0 = s3i32
	s4i32 = l0
	s5i32 = l4
	if s4i32 > s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l0 = s2i32
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
	s0i32 = l4
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l0
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l3 = s0i32
lbl1:
	s0i32 = l7
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l7
	s1i32 = l2
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l6
		s2i32 = l2
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s3i32 = l2
		s2i32 = s2i32 - s3i32
		l1 = s2i32
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		s3i32 = l4
		s4i32 = l4
		s5i32 = l1
		if s4i32 > s5i32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		l1 = s2i32
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l2 = s2i32
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
		l3 = s0i32
		s0i32 = l4
		s1i32 = l1
		s0i32 = s0i32 - s1i32
		l4 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s1i32 = l3
		s0i32 = s0i32 + s1i32
	} else {
		s0i32 = l3
	}
	s1i32 = l6
	s2i32 = l0
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l4
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
lbl0:
	s0i32 = l5
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
