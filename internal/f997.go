package internal

import (
	"unsafe"
)

func f997(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	s1i32 = 1280
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 2
	l7 = s0i32
	s0i32 = l1
	s1i32 = l6
	s2i32 = 1248
	s1i32 = s1i32 + s2i32
	s0i32 = f2083(ctx, s0i32, s1i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l7 = s0i32
		s0i32 = l1
		s1i32 = l6
		s2i32 = 1248
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		s3i32 = 0
		s0i32 = f333(ctx, s0i32, s1i32, s2i32, s3i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l6
	s2i32 = 1248
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = l2
	s4i32 = l3
	s5i32 = l3
	s6i32 = 20
	s5i32 = s5i32 + s6i32
	s6i32 = l3
	s6i32 = int32(ctx.Mem[int(s6i32+40)])
	if s6i32 != 0 {
		// s4i32 = s4i32
	} else {
		s4i32 = s5i32
	}
	s5i32 = l6
	s6i32 = 72
	s5i32 = s5i32 + s6i32
	s6i32 = l0
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+44)]))
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32, int32) int32)(table[s6i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	l9 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l9
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 1
		l8 = s0i32
		s0i32 = l6
		s1i32 = 72
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = 100
		s1i32 = s1i32 + s2i32
		s2i32 = l6
		s3i32 = 116
		s2i32 = s2i32 + s3i32
		s3i32 = l7
		s4i32 = 1
		if s3i32 == s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		s4i32 = l3
		s5i32 = l4
		f478(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		goto lbl3
	}
	s0i32 = 1
	l10 = s0i32
lbl3:
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)]))
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 0
		f91(ctx, s0i32, s1i32)
		goto lbl5
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
	f13(ctx, s0i32)
lbl5:
	s0i32 = l10
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = 0
	l8 = s0i32
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1216)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1184)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l3
	s2i32 = l3
	s3i32 = 20
	s2i32 = s2i32 + s3i32
	s3i32 = l3
	s3i32 = int32(ctx.Mem[int(s3i32+40)])
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l0
	s3i32 = l2
	s4i32 = l6
	s5i32 = 1216
	s4i32 = s4i32 + s5i32
	s5i32 = l5
	s0i32 = f1126(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1216)]))
	l5 = s0i32
	s0i32 = l0
	s1i32 = l6
	s2i32 = 1184
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	s3i32 = 1216
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s4i32 = 0
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+32)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1184)]))
	l2 = s0i32
	s0i32 = l6
	s1i32 = 72
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = l4
	s0i32 = f151(ctx, s0i32, s1i32, s2i32)
	l0 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1104)]))
	l3 = s0i32
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1100)]))
	s2i32 = l6
	s3i32 = 1184
	s2i32 = s2i32 + s3i32
	s3i32 = 4
	s2i32 = s2i32 | s3i32
	s0i32 = f111(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 44
		s0i32 = s0i32 + s1i32
		l4 = s0i32
	lbl9:
		s0i32 = l3
		s1i32 = l6
		s2i32 = 1184
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = l3
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
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
		s0i32 = l1
		f110(ctx, s0i32)
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+60)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
	}
	s0i32 = l0
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = f74(ctx, s0i32)
	s0i32 = l0
	f40(ctx, s0i32)
	s0i32 = 1
	l8 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l2
	f13(ctx, s0i32)
lbl7:
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	f13(ctx, s0i32)
lbl0:
	s0i32 = l6
	s1i32 = 1280
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l8
	return s0i32
}
