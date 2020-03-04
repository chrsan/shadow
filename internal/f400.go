package internal

import (
	"unsafe"
)

func f400(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
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
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s9i32 int32
	_ = s9i32
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 108
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l0
		s2i32 = 96
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s2i32 = l0
		s3i32 = 44
		s2i32 = s2i32 + s3i32
		s3i32 = l0
		s4i32 = 60
		s3i32 = s3i32 + s4i32
		s4i32 = l0
		s5i32 = 36
		s4i32 = s4i32 + s5i32
		s5i32 = l0
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s6i32 = l0
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
		s7i32 = l0
		s7i32 = int32(ctx.Mem[int(s7i32+84)])
		s8i32 = l2
		s9i32 = l0
		s9i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s9i32+92)]))
		if int(s9i32) < 0 || int(s9i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s9i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s9i32].numParams != 9 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32, float32, float32, int32, int32))(table[s9i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5f32, s6f32, s7i32, s8i32)
		s0i32 = l4
		f230(ctx, s0i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+85)])
		if s0i32 != 0 {
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l1 = s0i32
			s0i32 = int32(ctx.Mem[int(s0i32+84)])
			if s0i32 != 0 {
				s0i32 = l1
				s1i32 = 4
				s0i32 = s0i32 + s1i32
				s1i32 = l1
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
				s2i32 = l1
				s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
				s0i32 = f28(ctx, s0i32, s1i32, s2i32)
				l2 = s0i32
				s0i32 = l1
				s1i32 = 0
				ctx.Mem[int(s0i32+84)] = uint8(s1i32)
				s0i32 = l1
				s1i32 = l2
				ctx.Mem[int(s0i32+85)] = uint8(s1i32)
			}
			s0i32 = l4
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l2 = s0i32
			s0i32 = int32(ctx.Mem[int(s0i32+84)])
			if s0i32 != 0 {
				s0i32 = l2
				s1i32 = 4
				s0i32 = s0i32 + s1i32
				s1i32 = l2
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
				s2i32 = l2
				s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
				s0i32 = f28(ctx, s0i32, s1i32, s2i32)
				l6 = s0i32
				s0i32 = l2
				s1i32 = 0
				ctx.Mem[int(s0i32+84)] = uint8(s1i32)
				s0i32 = l2
				s1i32 = l6
				ctx.Mem[int(s0i32+85)] = uint8(s1i32)
			}
			s0i32 = l2
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l7 = s0f32
			s1i32 = l2
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
			l9 = s1f32
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl1
			}
			s0i32 = l2
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l10 = s0f32
			s1i32 = l2
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			l11 = s1f32
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl1
			}
			s0i32 = l1
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l8 = s0f32
			s1i32 = l1
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
			l12 = s1f32
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl1
			}
			s0f32 = l8
			s1f32 = l7
			if s0f32 <= s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl1
			}
			s0i32 = l1
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l7 = s0f32
			s1i32 = l1
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			l8 = s1f32
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl1
			}
			s0f32 = l8
			s1f32 = l11
			if s0f32 >= s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl1
			}
			s0f32 = l12
			s1f32 = l9
			if s0f32 >= s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl1
			}
			s0f32 = l7
			s1f32 = l10
			if s0f32 <= s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl1
			}
			s0i32 = l5
			s1i32 = l4
			f194(ctx, s0i32, s1i32)
			goto lbl1
		}
		s0i32 = l5
		s1i32 = l3
		f683(ctx, s0i32, s1i32)
		s0i32 = l4
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l3
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		f176(ctx, s0i32, s1f32, s2f32)
		s0i32 = l4
		s1i32 = l5
		f676(ctx, s0i32, s1i32)
		s0i32 = l4
		f230(ctx, s0i32)
		goto lbl1
	}
	s0i32 = l0
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l3
	f683(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = 108
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l0
	s2i32 = 60
	s1i32 = s1i32 + s2i32
	s2i32 = l0
	s3i32 = 28
	s2i32 = s2i32 + s3i32
	s3i32 = l3
	s4i32 = l1
	s5i32 = 0
	s6i32 = l2
	if s6i32 != 0 {
		// s4i32 = s4i32
	} else {
		s4i32 = s5i32
	}
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+88)]))
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
	s0i32 = l4
	s1i32 = l1
	f676(ctx, s0i32, s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
	l2 = s0i32
	s0i32 = l3
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s1f32 = -s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l3
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s1f32 = -s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l4
	s1i32 = l0
	s2i32 = 52
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = l0
	s4i32 = 68
	s3i32 = s3i32 + s4i32
	s4i32 = l1
	s5i32 = 0
	s6i32 = l0
	s6i32 = int32(ctx.Mem[int(s6i32+84)])
	if s6i32 != 0 {
		// s4i32 = s4i32
	} else {
		s4i32 = s5i32
	}
	s5i32 = l2
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
	s0i32 = l4
	f230(ctx, s0i32)
lbl1:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 108
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 120
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = l3
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	f684(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	f278(ctx, s0i32)
lbl0:
	s0i32 = l0
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	f278(ctx, s0i32)
	s0i32 = l0
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+108)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
