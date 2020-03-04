package internal

import (
	"unsafe"
)

func f258(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = 19480
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4380)]))
	l1 = s0i32
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = 0
		l1 = s0i32
	lbl1:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l2 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
		l4 = s0i32
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4380)]))
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4380)])) = uint32(s1i32)
			s0i32 = l2
			s1i32 = l4
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
			goto lbl2
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		s1i32 = 2
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		if int(s1i32) < 0 || int(s1i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s1i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s1i32].numParams != 1 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4380)]))
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4380)])) = uint32(s1i32)
		s0i32 = l0
		f54(ctx, s0i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)]))
		if int(s1i32) < 0 || int(s1i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s1i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s1i32].numParams != 1 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
	lbl2:
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l0
	f54(ctx, s0i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4428)]))
	l2 = s0i32
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4428)])) = uint32(s1i32)
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l1 = s0i32
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = l1
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
			s0i32 = l1
			f12(ctx, s0i32)
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l4 = s0i32
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l4
			s2i32 = l2
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
			l1 = s2i32
			if s1i32 == s2i32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				goto lbl7
			}
		lbl8:
			s0i32 = l1
			l3 = s0i32
			s1i32 = -48
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s0i32 = l3
			s1i32 = -20
			s0i32 = s0i32 + s1i32
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l3 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl9
			}
			s0i32 = l3
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l5 = s1i32
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l5
			s1i32 = 1
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl9
			}
			s0i32 = l3
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			if int(s1i32) < 0 || int(s1i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s1i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s1i32].numParams != 1 {
				panic("argument count mismatch")
			}
			(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
		lbl9:
			s0i32 = l1
			s1i32 = l4
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl8
			}
			s0i32 = l2
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		lbl7:
			l1 = s0i32
			s0i32 = l2
			s1i32 = l4
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
			s0i32 = l1
			f12(ctx, s0i32)
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l1 = s0i32
		s0i32 = l2
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		if s0i32 != 0 {
			s0i32 = l1
			f13(ctx, s0i32)
		}
		s0i32 = l2
		f12(ctx, s0i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4384)]))
	l1 = s0i32
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4384)])) = uint32(s1i32)
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		if int(s1i32) < 0 || int(s1i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s1i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s1i32].numParams != 1 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l1
		f12(ctx, s0i32)
	}
	s0i32 = l0
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	f522(ctx, s0i32)
	s0i32 = l0
	return s0i32
}
