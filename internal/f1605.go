package internal

import (
	"unsafe"
)

func f1605(ctx *Context, l0 int32, l1 int32) {
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
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
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
		return
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l2 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s2i32 = l3
		s1i32 = f666(ctx, s1i32, s2i32)
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
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
		goto lbl1
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l4 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l3
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l3
	s0i32 = s0i32 ^ s1i32
	s1i32 = -2048144789
	s0i32 = s0i32 * s1i32
	l2 = s0i32
	s1i32 = 13
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l2
	s0i32 = s0i32 ^ s1i32
	s1i32 = -1028477387
	s0i32 = s0i32 * s1i32
	l2 = s0i32
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l2
	s0i32 = s0i32 ^ s1i32
	l2 = s0i32
	s1i32 = 1
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l7 = s0i32
	s1i32 = l4
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 & s1i32
	l2 = s0i32
	s0i32 = l0
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
lbl5:
	s0i32 = l5
	s1i32 = l2
	s2i32 = 12
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l7
	s1i32 = l9
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l3
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l2
	s1i32 = 0
	s2i32 = l4
	s3i32 = l2
	s4i32 = 0
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
	s0i32 = s0i32 + s1i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l4
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl4:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	if int(s2i32) < 0 || int(s2i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s2i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s2i32].numParams != 1 {
		panic("argument count mismatch")
	}
	s1i32 = (*(*func(*Context, int32) int32)(table[s2i32].f()))(ctx, s1i32)
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)]))
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
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	f1604(ctx, s0i32, s1i32, s2i32)
	goto lbl1
lbl3:
	s0i32 = l0
	s1i32 = l5
	s2i32 = l2
	s3i32 = 12
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = 8
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
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
lbl1:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l2 = s0i32
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 36
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		f66(ctx, s0i32, s1i32)
	}
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
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
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = l3
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
