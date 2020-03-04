package internal

import (
	"unsafe"
)

func f1899(ctx *Context, l0 int32, l1 int32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l3 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l1
		s1i32 = l0
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		if int(s2i32) < 0 || int(s2i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s2i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s2i32].numParams != 2 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32))(table[s2i32].f()))(ctx, s0i32, s1i32)
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l0 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l3 = s0i32
lbl2:
	s0i32 = l0
	l5 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l6 = s0i32
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l0 = s0i32
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l0 = s0i32
		s0i32 = l2
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		if int(s2i32) < 0 || int(s2i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s2i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s2i32].numParams != 2 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32))(table[s2i32].f()))(ctx, s0i32, s1i32)
		s0i32 = l4
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		goto lbl3
	}
	s0i32 = l4
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l6
	s2i32 = 2
	if s1i32 < s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s0i32 = l3
	s1i32 = l5
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl6:
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = 0
		l9 = s0i32
		s0i32 = l4
		l0 = s0i32
	lbl7:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l7 = s0i32
		s0i32 = l2
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		if int(s2i32) < 0 || int(s2i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s2i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s2i32].numParams != 2 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32))(table[s2i32].f()))(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l9
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l8
		l3 = s0i32
		s1i32 = l5
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
	}
	s0i32 = l4
	s1i32 = l6
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
lbl3:
	l4 = s0i32
	s0i32 = l5
	l3 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s1i32 = 2147483647
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl0:
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
