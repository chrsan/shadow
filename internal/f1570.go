package internal

import (
	"unsafe"
)

func f1570(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
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
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 0
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+32)]))
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
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l5 = s0i32
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l6 = s0i32
		s0i32 = l0
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l6
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = l6
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l5 = s1i32
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l5
			s1i32 = 1
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l6
				s1i32 = l6
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
			}
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l5 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl4
			}
			s0i32 = l5
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l6 = s1i32
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = 1
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl4
			}
			s0i32 = l5
			s1i32 = l5
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
		lbl4:
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
			l5 = s0i32
		}
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l5
	s2f32 = l2
	s3f32 = l3
	s4i32 = l4
	f574(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l5
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
lbl0:
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
