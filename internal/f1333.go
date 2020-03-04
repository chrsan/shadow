package internal

import (
	"unsafe"
)

func f1333(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l6
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 13195221663744
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l4
	if s1i32 != 0 {
		s1i32 = l5
		s2i32 = l4
		f114(ctx, s1i32, s2i32)
		s1i32 = l5
		s2i32 = l5
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
		s3i32 = -193
		s2i32 = s2i32 & s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)])) = uint32(s2i32)
		s1i32 = l6
		s2i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = l5
		s2i32 = l6
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		f123(ctx, s1i32, s2i32)
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l4 = s1i32
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl4
		}
		s1i32 = l4
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l7 = s2i32
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l7
		s2i32 = 1
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl4
		}
		s1i32 = l4
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		if int(s2i32) < 0 || int(s2i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s2i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s2i32].numParams != 1 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32))(table[s2i32].f()))(ctx, s1i32)
	lbl4:
		s1i32 = l5
		s1i32 = f65(ctx, s1i32)
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl1
		}
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = 6
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l4 = s1i32
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l3
			s2i32 = l5
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
			s3i32 = l5
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
			s4i32 = l5
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
			s3i32 = s3i32 | s4i32
			s2i32 = s2i32 | s3i32
			if s2i32 == 0 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				goto lbl2
			}
		}
		s1i32 = l5
		s2i32 = l3
		s3i32 = l6
		s4i32 = 8
		s3i32 = s3i32 + s4i32
		s4i32 = l4
		s1i32 = f64(ctx, s1i32, s2i32, s3i32, s4i32)
		goto lbl2
	}
	s1i32 = l3
lbl2:
	s0i32 = f70(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4388)]))
	l4 = s0i32
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 1
		f57(ctx, s0i32, s1i32)
	}
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l5
	s3i32 = 0
	s4i32 = l3
	s0i32 = f61(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l4 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
lbl8:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l7
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+56)]))
	s5i32 = l7
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+116)]))
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
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl8
	}
lbl7:
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+64)])
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		f54(ctx, s0i32)
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = f23(ctx, s0i32)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
lbl0:
	s0i32 = l5
	s0i32 = f23(ctx, s0i32)
	s0i32 = l6
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
