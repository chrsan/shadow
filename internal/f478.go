package internal

import (
	"unsafe"
)

func f478(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var l11 int32
	_ = l11
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 int64
	_ = l25
	var l26 int64
	_ = l26
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	s0i32 = ctx.g0
	s1i32 = 5376
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l6
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1i32 = l4
	s2i32 = l5
	s0i32 = f151(ctx, s0i32, s1i32, s2i32)
	l21 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1104)]))
	l10 = s0i32
	s0i32 = l6
	s1i32 = l21
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1100)]))
	s2i32 = l1
	s0i32 = f111(ctx, s0i32, s1i32, s2i32)
	l22 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l18 = s0i32
		s0i32 = l22
		s1i32 = 44
		s0i32 = s0i32 + s1i32
		l17 = s0i32
		s0i32 = l6
		s1i32 = 1176
		s0i32 = s0i32 + s1i32
		s1i32 = 4
		s0i32 = s0i32 | s1i32
		l20 = s0i32
		s0i32 = l6
		s1i32 = 1232
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l6
		s1i32 = 5344
		s0i32 = s0i32 + s1i32
		s1i32 = 4
		s0i32 = s0i32 | s1i32
		l15 = s0i32
	lbl1:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l16 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l14 = s0i32
		s0i32 = l6
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5344)])) = uint32(s1i32)
		s0i32 = l18
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l25 = s0i64
		s0i32 = l6
		s1i32 = l16
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5360)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l14
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5356)])) = uint32(s1i32)
		s0i32 = l6
		s1i64 = l25
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+5348)])) = uint64(s1i64)
		s0i32 = l14
		s1i64 = l25
		s1i32 = int32(s1i64)
		s0i32 = s0i32 - s1i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l16
		s1i64 = l25
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		s1i32 = int32(s1i64)
		l5 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l9 = s0i32
		s0i32 = l6
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l8 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5364)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l7
		s2i32 = l8
		s3i32 = l5
		s4i32 = l9
		s3i32 = s3i32 - s4i32
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5344)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5368)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		s0i32 = l6
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5352)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5348)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l7
		s1i64 = int64(s1i32)
		s2i32 = l16
		s2i64 = int64(s2i32)
		s3i64 = l25
		s4i64 = 32
		s3i64 = s3i64 >> (uint64(s4i64) & 63)
		s2i64 = s2i64 - s3i64
		s1i64 = s1i64 + s2i64
		l26 = s1i64
		s2i64 = 2147483647
		s3i64 = l26
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l26 = s1i64
		s2i64 = -2147483647
		s3i64 = l26
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5360)])) = uint32(s1i64)
		s0i32 = l6
		s1i32 = l5
		s1i64 = int64(s1i32)
		s2i32 = l14
		s2i64 = int64(s2i32)
		s3i64 = l25
		s4i64 = 32
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 32
		s3i64 = s3i64 >> (uint64(s4i64) & 63)
		s2i64 = s2i64 - s3i64
		s1i64 = s1i64 + s2i64
		l25 = s1i64
		s2i64 = 2147483647
		s3i64 = l25
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l25 = s1i64
		s2i64 = -2147483647
		s3i64 = l25
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5356)])) = uint32(s1i64)
		s0i32 = l6
		s1i32 = 1224
		s0i32 = s0i32 + s1i32
		s1i32 = l15
		s2i32 = l17
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l10
		s1i32 = l6
		s2i32 = 5344
		s1i32 = s1i32 + s2i32
		s2i32 = l6
		s3i32 = 1224
		s2i32 = s2i32 + s3i32
		s3i32 = l10
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
	lbl2:
		s0i32 = l15
		s1i32 = l18
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l15
		s1i32 = l18
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l16
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5360)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l14
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5348)])) = uint32(s1i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5356)]))
		l9 = s0i32
		s1i32 = l7
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l16
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5352)]))
		l5 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l12 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l13 = s0i32
		s0i32 = l6
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l19 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5364)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l11
		s2i32 = l19
		s3i32 = l5
		s4i32 = l12
		s3i32 = s3i32 - s4i32
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l7
		s3i32 = l13
		s2i32 = s2i32 - s3i32
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5344)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5368)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l11 = s0i32
		s0i32 = l6
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l12 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5352)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l12
		s1i64 = int64(s1i32)
		s2i32 = l16
		s2i64 = int64(s2i32)
		s3i32 = l5
		s3i64 = int64(s3i32)
		s2i64 = s2i64 - s3i64
		s1i64 = s1i64 + s2i64
		l25 = s1i64
		s2i64 = 2147483647
		s3i64 = l25
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l25 = s1i64
		s2i64 = -2147483647
		s3i64 = l25
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5360)])) = uint32(s1i64)
		s0i32 = l6
		s1i32 = l11
		s2i32 = l8
		s1i32 = s1i32 - s2i32
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5348)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l5
		s1i64 = int64(s1i32)
		s2i32 = l9
		s2i64 = int64(s2i32)
		s3i32 = l14
		s4i32 = -1
		s3i32 = s3i32 ^ s4i32
		s3i64 = int64(s3i32)
		s2i64 = s2i64 + s3i64
		s1i64 = s1i64 + s2i64
		l25 = s1i64
		s2i64 = 2147483647
		s3i64 = l25
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l25 = s1i64
		s2i64 = -2147483647
		s3i64 = l25
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5356)])) = uint32(s1i64)
		s0i32 = l6
		s1i32 = 1224
		s0i32 = s0i32 + s1i32
		s1i32 = l15
		s2i32 = l17
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l10
		s1i32 = l6
		s2i32 = 5344
		s1i32 = s1i32 + s2i32
		s2i32 = l6
		s3i32 = 1224
		s2i32 = s2i32 + s3i32
		s3i32 = l10
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
	lbl3:
		s0i32 = l15
		s1i32 = l18
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l15
		s1i32 = l18
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l16
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5352)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l14
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5356)])) = uint32(s1i32)
		s0i32 = l14
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5348)]))
		l9 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5360)]))
		l8 = s0i32
		s1i32 = l5
		s0i32 = s0i32 - s1i32
		l11 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l12 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l13 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l19 = s0i32
		s0i32 = l6
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l23 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5364)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l12
		s2i32 = l23
		s3i32 = l5
		s4i32 = l13
		s3i32 = s3i32 - s4i32
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l9
		s3i32 = l19
		s2i32 = s2i32 - s3i32
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5344)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5368)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l12 = s0i32
		s0i32 = l6
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l13 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5348)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l13
		s1i64 = int64(s1i32)
		s2i32 = l14
		s2i64 = int64(s2i32)
		s3i32 = l9
		s3i64 = int64(s3i32)
		s2i64 = s2i64 - s3i64
		s1i64 = s1i64 + s2i64
		l25 = s1i64
		s2i64 = 2147483647
		s3i64 = l25
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l25 = s1i64
		s2i64 = -2147483647
		s3i64 = l25
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5356)])) = uint32(s1i64)
		s0i32 = l6
		s1i32 = l12
		s2i32 = l11
		s1i32 = s1i32 - s2i32
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5352)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l9
		s1i64 = int64(s1i32)
		s2i32 = l8
		s2i64 = int64(s2i32)
		s3i32 = l5
		s3i64 = int64(s3i32)
		s2i64 = s2i64 - s3i64
		s1i64 = s1i64 + s2i64
		l25 = s1i64
		s2i64 = 2147483647
		s3i64 = l25
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l25 = s1i64
		s2i64 = -2147483647
		s3i64 = l25
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5360)])) = uint32(s1i64)
		s0i32 = l6
		s1i32 = 1224
		s0i32 = s0i32 + s1i32
		s1i32 = l15
		s2i32 = l17
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l10
		s1i32 = l6
		s2i32 = 5344
		s1i32 = s1i32 + s2i32
		s2i32 = l6
		s3i32 = 1224
		s2i32 = s2i32 + s3i32
		s3i32 = l10
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
	lbl4:
		s0i32 = l15
		s1i32 = l18
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l15
		s1i32 = l18
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5352)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5348)])) = uint32(s1i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5356)]))
		l9 = s0i32
		s1i32 = l7
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5360)]))
		l11 = s0i32
		s1i32 = l5
		s0i32 = s0i32 - s1i32
		l12 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l13 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l19 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l23 = s0i32
		s0i32 = l6
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l24 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5364)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l13
		s2i32 = l24
		s3i32 = l5
		s4i32 = l19
		s3i32 = s3i32 - s4i32
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l7
		s3i32 = l23
		s2i32 = s2i32 - s3i32
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5344)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5368)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l13 = s0i32
		s0i32 = l6
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l12
		s1i32 = s1i32 - s2i32
		l12 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5352)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l12
		s1i64 = int64(s1i32)
		s2i32 = l11
		s2i64 = int64(s2i32)
		s3i32 = l5
		s3i64 = int64(s3i32)
		s2i64 = s2i64 - s3i64
		s1i64 = s1i64 + s2i64
		l25 = s1i64
		s2i64 = 2147483647
		s3i64 = l25
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l25 = s1i64
		s2i64 = -2147483647
		s3i64 = l25
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5360)])) = uint32(s1i64)
		s0i32 = l6
		s1i32 = l13
		s2i32 = l8
		s1i32 = s1i32 - s2i32
		l8 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5348)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l8
		s1i64 = int64(s1i32)
		s2i32 = l9
		s2i64 = int64(s2i32)
		s3i32 = l14
		s4i32 = -1
		s3i32 = s3i32 ^ s4i32
		s3i64 = int64(s3i32)
		s2i64 = s2i64 + s3i64
		s1i64 = s1i64 + s2i64
		l25 = s1i64
		s2i64 = 2147483647
		s3i64 = l25
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l25 = s1i64
		s2i64 = -2147483647
		s3i64 = l25
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5356)])) = uint32(s1i64)
		s0i32 = l6
		s1i32 = 1224
		s0i32 = s0i32 + s1i32
		s1i32 = l15
		s2i32 = l17
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l10
		s1i32 = l6
		s2i32 = 5344
		s1i32 = s1i32 + s2i32
		s2i32 = l6
		s3i32 = 1224
		s2i32 = s2i32 + s3i32
		s3i32 = l10
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
	lbl5:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l8 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l12 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l13 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l9 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l19 = s0i32
		s0i32 = l6
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l5
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5340)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l7
		s2i32 = l19
		s1i32 = s1i32 + s2i32
		s2i32 = l9
		s1i32 = s1i32 - s2i32
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5336)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l13
		s2i32 = l16
		s1i32 = s1i32 + s2i32
		s2i32 = l12
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5332)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l11
		s2i32 = l14
		s1i32 = s1i32 + s2i32
		s2i32 = l8
		s1i32 = s1i32 - s2i32
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5328)])) = uint32(s1i32)
		s0i32 = l3
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = 1224
			s0i32 = s0i32 + s1i32
			s1i32 = l6
			s2i32 = 5328
			s1i32 = s1i32 + s2i32
			s2i32 = l17
			s0i32 = f25(ctx, s0i32, s1i32, s2i32)
			if s0i32 != 0 {
				s0i32 = l10
				s1i32 = l6
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1224)]))
				l5 = s1i32
				s2i32 = l6
				s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+1228)]))
				l7 = s2i32
				s3i32 = l6
				s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+1232)]))
				s4i32 = l5
				s3i32 = s3i32 - s4i32
				s4i32 = l6
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+1236)]))
				s5i32 = l7
				s4i32 = s4i32 - s5i32
				s5i32 = l10
				s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
				s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
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
			}
			s0i32 = l6
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5336)]))
			l9 = s0i32
			s0i32 = l6
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5328)]))
			l7 = s0i32
		}
		s0i32 = l6
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1224)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 4096
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1228)])) = uint32(s1i32)
		s0i32 = l4
		l5 = s0i32
		s0i32 = l9
		s1i32 = l7
		s0i32 = s0i32 - s1i32
		l11 = s0i32
		s1i32 = 3
		s0i32 = s0i32 * s1i32
		s1i32 = 3
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s1i32 = 4097
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l8
			s1i32 = 4096
			s2i32 = l8
			s3i32 = 4096
			if uint32(s2i32) > uint32(s3i32) {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l7 = s0i32
			s1i32 = 2
			s0i32 = f44(ctx, s0i32, s1i32)
			l5 = s0i32
			s0i32 = l6
			s1i32 = l7
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1228)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = l5
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1224)])) = uint32(s1i32)
			s0i32 = l6
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5336)]))
			l9 = s0i32
			s0i32 = l6
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5328)]))
			l7 = s0i32
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l8 = s0i32
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5332)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1220)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1216)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1212)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1208)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l11
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l6
		s1i32 = 1208
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = 1208
		s1i32 = s1i32 + s2i32
		s2i32 = l17
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1220)]))
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1212)]))
		l7 = s1i32
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0i32 = l7
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 - s1i32
		l7 = s0i32
		s1i32 = 0
		s2i32 = l7
		s3i32 = 0
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l5
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1216)]))
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+1208)]))
		s1i32 = s1i32 - s2i32
		l11 = s1i32
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l12 = s0i32
	lbl10:
		s0i32 = l5
		s1i32 = l11
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l12
		s1i32 = 0
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l9
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l14
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1i32 = s1i32 - s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s3i32 = l7
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l10
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1208)]))
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s3i32 = l7
		s2i32 = s2i32 + s3i32
		s3i32 = l9
		s4i32 = l5
		s5i32 = l10
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
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
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l8
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
	lbl9:
		s0i32 = l6
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1220)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5336)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1216)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5340)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1212)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5328)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1208)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 1208
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = 1208
		s1i32 = s1i32 + s2i32
		s2i32 = l17
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1220)]))
		l7 = s0i32
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1212)]))
		l8 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l11 = s0i32
		s1i32 = l8
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s0i32 = l11
		s1i32 = l7
		s0i32 = s0i32 - s1i32
		l7 = s0i32
		s0i32 = l5
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1216)]))
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+1208)]))
		s1i32 = s1i32 - s2i32
		l11 = s1i32
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l12 = s0i32
	lbl12:
		s0i32 = l5
		s1i32 = l11
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l12
		s1i32 = 0
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l9
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l14
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1i32 = s1i32 - s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s3i32 = l7
		s4i32 = -1
		s3i32 = s3i32 ^ s4i32
		l13 = s3i32
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
		s3i32 = s3i32 + s4i32
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
		s3i32 = s3i32 - s4i32
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l10
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1208)]))
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s3i32 = l13
		s2i32 = s2i32 + s3i32
		s3i32 = l9
		s4i32 = l5
		s5i32 = l10
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
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
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l8
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
	lbl11:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5340)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1220)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5328)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1216)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5332)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1212)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1208)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 1208
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = 1208
		s1i32 = s1i32 + s2i32
		s2i32 = l17
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l5 = s2i32
			s3i32 = l6
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+1208)]))
			s2i32 = s2i32 + s3i32
			s3i32 = l1
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
			s2i32 = s2i32 - s3i32
			s1i32 = s1i32 + s2i32
			s2i32 = l5
			s1i32 = s1i32 - s2i32
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
			s3i32 = l16
			s2i32 = s2i32 * s3i32
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1176)])) = uint32(s1i32)
			s0i32 = l20
			s1i32 = l6
			s2i32 = 1216
			s1i32 = s1i32 + s2i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l20
			s1i32 = l6
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+1208)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l6
			s1i64 = 4294967296
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1196)])) = uint64(s1i64)
			s0i32 = l10
			s1i32 = l6
			s2i32 = 1176
			s1i32 = s1i32 + s2i32
			s2i32 = l6
			s3i32 = 1208
			s2i32 = s2i32 + s3i32
			s3i32 = l10
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
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l5 = s0i32
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5340)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1220)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1216)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5332)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1212)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5336)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1208)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 1208
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = 1208
		s1i32 = s1i32 + s2i32
		s2i32 = l17
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l6
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+1208)]))
			s3i32 = l0
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
			s4i32 = l1
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
			s3i32 = s3i32 - s4i32
			s2i32 = s2i32 + s3i32
			s1i32 = s1i32 + s2i32
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			s1i32 = s1i32 - s2i32
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
			s3i32 = l16
			s2i32 = s2i32 * s3i32
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1176)])) = uint32(s1i32)
			s0i32 = l20
			s1i32 = l6
			s2i32 = 1216
			s1i32 = s1i32 + s2i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l20
			s1i32 = l6
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+1208)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l6
			s1i64 = 4294967296
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1196)])) = uint64(s1i64)
			s0i32 = l10
			s1i32 = l6
			s2i32 = 1176
			s1i32 = s1i32 + s2i32
			s2i32 = l6
			s3i32 = 1208
			s2i32 = s2i32 + s3i32
			s3i32 = l10
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
		}
		s0i32 = l4
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1224)]))
		l5 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			f13(ctx, s0i32)
		}
		s0i32 = l22
		f110(ctx, s0i32)
		s0i32 = l22
		s0i32 = int32(ctx.Mem[int(s0i32+60)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l21
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = f74(ctx, s0i32)
	s0i32 = l21
	f40(ctx, s0i32)
	s0i32 = l6
	s1i32 = 5376
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
