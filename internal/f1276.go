package internal

import (
	"unsafe"
)

func f1276(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var s6i32 int32
	_ = s6i32
	var s1i64 int64
	_ = s1i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l7
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s0i32 = f46(ctx, s0i32, s1i32)
	l10 = s0i32
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	l13 = s0i32
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl1:
		s0i32 = l10
		s1i32 = l10
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = -2
		s1i32 = s1i32 & s2i32
		s2i32 = l1
		s3i32 = l11
		s4i32 = 52
		s3i32 = s3i32 * s4i32
		s2i32 = s2i32 + s3i32
		l9 = s2i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
		s3i32 = 15
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l10
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		s2i32 = l9
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)]))
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = 0
		l8 = s0i32
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		s1i32 = 0
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
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
			s0i32 = l9
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
			l8 = s0i32
			s0i32 = l7
			s1i64 = 69784829952
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
			s0i32 = l7
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
			s0i32 = l7
			s1i64 = 1065353216
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
			s0i32 = l7
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l7
			s1i64 = 1065353216
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l7
			s1i32 = l7
			s2i32 = 40
			s1i32 = s1i32 + s2i32
			s2i32 = l4
			s3i32 = l8
			s4i32 = 40
			s3i32 = s3i32 * s4i32
			s2i32 = s2i32 + s3i32
			f69(ctx, s0i32, s1i32, s2i32)
			s0i32 = l0
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
			s0i32 = 1
			l8 = s0i32
		}
		s0i32 = l9
		s0i32 = int32(ctx.Mem[int(s0i32+48)])
		if s0i32 != 0 {
			s0i32 = l8
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l0
				s1i32 = l0
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
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
			s0i32 = f38(ctx, s0i32)
			l8 = s0i32
			s1i32 = l3
			s2i32 = l12
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = 4
			s3i32 = 1
			f277(ctx, s0i32, s1i32, s2i32, s3i32)
			s0i32 = l0
			s1i32 = l8
			s2i32 = 1
			s3i32 = l10
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+44)]))
			s4i32 = 1
			s3i32 = s3i32 & s4i32
			s4i32 = l0
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+40)]))
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
			s0i32 = l8
			f34(ctx, s0i32)
			s0i32 = l12
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l12 = s0i32
			s0i32 = 1
			l8 = s0i32
		}
		s0i32 = l0
		s1i32 = l9
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l9
		s3i32 = 4
		s2i32 = s2i32 + s3i32
		s3i32 = l9
		s4i32 = 20
		s3i32 = s3i32 + s4i32
		s4i32 = l10
		s5i32 = l6
		s6i32 = l0
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+124)]))
		if int(s6i32) < 0 || int(s6i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s6i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s6i32].numParams != 6 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32, int32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		s0i32 = l8
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
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
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l13
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
		}
		s0i32 = l11
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s1i32 = l2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l10
	s0i32 = f23(ctx, s0i32)
	s0i32 = l7
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
