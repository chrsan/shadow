package internal

import (
	"unsafe"
)

func f1967(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
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
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 208
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint32(s1i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 152
	s0i32 = s0i32 + s1i32
	s1i32 = 4
	s0i32 = s0i32 | s1i32
	l6 = s0i32
	s1i32 = l4
	s0i32 = f46(ctx, s0i32, s1i32)
	l7 = s0i32
	s0i32 = l5
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+164)]))
	s2i32 = l0
	s3i32 = 116
	s2i32 = s2i32 + s3i32
	f477(ctx, s0i32, s1i32, s2i32)
	s0i32 = l7
	s1i32 = l5
	s2i32 = 144
	s1i32 = s1i32 + s2i32
	f235(ctx, s0i32, s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)]))
	l6 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
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
lbl0:
	s0i32 = l1
	s1i32 = 1084
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 96
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21300
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+160)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l0
		s2i32 = 156
		s1i32 = s1i32 + s2i32
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)]))
		l8 = s0i32
		if s0i32 != 0 {
			s0i32 = l8
			s1i32 = l8
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		}
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l7 = s0i32
		s0i32 = l6
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l7
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l7
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l8 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l7
		f12(ctx, s0i32)
	lbl3:
		s0i32 = l6
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+172)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+180)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21916
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l0
		s2i32 = 1064
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s0i32 = f46(ctx, s0i32, s1i32)
		l4 = s0i32
		s0i32 = l5
		s1i32 = 40
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = 156
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		s3i32 = 0
		f581(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l4
		s1i32 = l5
		s2i32 = 40
		s1i32 = s1i32 + s2i32
		f83(ctx, s0i32, s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l7 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l0
		s1i32 = l0
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
	lbl5:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1084)]))
		l0 = s0i32
		s0i32 = l5
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l2
		s1f32 = float32(s1i32)
		s2i32 = l3
		s2f32 = float32(s2i32)
		f117(ctx, s0i32, s1f32, s2f32)
		s0i32 = l6
		s1i32 = l0
		s2i32 = l5
		s3i32 = 0
		s4i32 = l4
		f502(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l4
		s0i32 = f23(ctx, s0i32)
		s0i32 = l6
		s1i32 = 21300
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		f12(ctx, s0i32)
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = 156
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = l3
	s4i32 = l5
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+152)]))
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+108)]))
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
lbl1:
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = f23(ctx, s0i32)
		s0i32 = l5
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)])) = uint32(s1i32)
	}
	s0i32 = l5
	s1i32 = 208
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
