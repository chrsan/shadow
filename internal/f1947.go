package internal

import (
	"math"
	"unsafe"
)

func f1947(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var l13 int64
	_ = l13
	var l14 int64
	_ = l14
	var l15 int64
	_ = l15
	var l16 int64
	_ = l16
	var l17 int64
	_ = l17
	var l18 int64
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
	var l23 float32
	_ = l23
	var l24 float32
	_ = l24
	var l25 float32
	_ = l25
	var l26 float32
	_ = l26
	var l27 float32
	_ = l27
	var l28 float32
	_ = l28
	var l29 float32
	_ = l29
	var l30 float32
	_ = l30
	var l31 float32
	_ = l31
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0i32 = ctx.g0
	s1i32 = 464
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l7
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+460)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+408)])) = uint32(s1i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l8 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+272)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+264)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+256)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 248
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		s2i32 = l2
		s1i32 = s1i32 - s2i32
		l9 = s1i32
		s1f32 = float32(s1i32)
		s2i32 = 0
		s3i32 = l3
		s2i32 = s2i32 - s3i32
		l11 = s2i32
		s2f32 = float32(s2i32)
		f117(ctx, s0i32, s1f32, s2f32)
		s0i32 = l7
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+400)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+392)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+384)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+376)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+368)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 368
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = 248
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s3i32 = 116
		s2i32 = s2i32 + s3i32
		f69(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s1i32 = 1064
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s1i32 = l4
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s2i32 = int32(ctx.Mem[int(s2i32+40)])
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l13 = s0i64
		s0i32 = l4
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])))
		l14 = s0i64
		s0i32 = l4
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])))
		l16 = s0i64
		s0i32 = l4
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])))
		l15 = s0i64
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+240)]))
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
		l4 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+172)]))
		l10 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)]))
		l12 = s0i32
		s0i32 = l7
		s1i32 = 21948
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 21940
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+272)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 21932
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+264)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 21924
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+256)])) = uint64(s1i64)
		s0i32 = 21916
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l17 = s0i64
		s0i32 = l7
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+376)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+296)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+384)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+304)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+392)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+312)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+400)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+320)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = l15
		s2i32 = l11
		s2i64 = int64(s2i32)
		l18 = s2i64
		s1i64 = s1i64 + s2i64
		l15 = s1i64
		s2i64 = 2147483647
		s3i64 = l15
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
		l15 = s1i64
		s2i64 = -2147483647
		s3i64 = l15
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
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+340)])) = uint32(s1i64)
		s0i32 = l7
		s1i64 = l16
		s2i32 = l9
		s2i64 = int64(s2i32)
		l15 = s2i64
		s1i64 = s1i64 + s2i64
		l16 = s1i64
		s2i64 = 2147483647
		s3i64 = l16
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
		l16 = s1i64
		s2i64 = -2147483647
		s3i64 = l16
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
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+336)])) = uint32(s1i64)
		s0i32 = l7
		s1i64 = l14
		s2i64 = l18
		s1i64 = s1i64 + s2i64
		l14 = s1i64
		s2i64 = 2147483647
		s3i64 = l14
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
		l14 = s1i64
		s2i64 = -2147483647
		s3i64 = l14
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
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+332)])) = uint32(s1i64)
		s0i32 = l7
		s1i64 = l17
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+368)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+288)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l12
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+352)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+348)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+344)])) = uint32(s1i32)
		s0i32 = l7
		s1i64 = l13
		s2i64 = l15
		s1i64 = s1i64 + s2i64
		l13 = s1i64
		s2i64 = 2147483647
		s3i64 = l13
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
		l13 = s1i64
		s2i64 = -2147483647
		s3i64 = l13
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
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)])) = uint32(s1i64)
		s0i32 = l1
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l1
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		}
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+360)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+356)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 136
		s0i32 = s0i32 + s1i32
		s1i32 = l8
		s2i32 = l7
		s3i32 = 248
		s2i32 = s2i32 + s3i32
		f352(ctx, s0i32, s1i32, s2i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
		l9 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 1
			l11 = s0i32
			s0i32 = 0
			l9 = s0i32
			goto lbl3
		}
		s0i32 = l7
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)]))
		l13 = s0i64
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l9
			s1i32 = l9
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
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+408)]))
		l1 = s0i32
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+460)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l7
			s1i32 = 408
			s0i32 = s0i32 + s1i32
			s1i32 = 4
			s0i32 = s0i32 | s1i32
			l8 = s0i32
			s1i32 = l1
			s0i32 = f46(ctx, s0i32, s1i32)
			l1 = s0i32
			s0i32 = l7
			s1i32 = l8
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+408)])) = uint32(s1i32)
			s0i32 = l7
			s1i32 = l8
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+460)])) = uint32(s1i32)
		}
		s0i64 = l13
		s1i64 = 32
		s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
		s0i32 = int32(s0i64)
		l8 = s0i32
		s0i64 = l13
		s0i32 = int32(s0i64)
		l10 = s0i32
		s0i32 = 0
		l11 = s0i32
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+240)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l7
		s2i32 = 240
		s1i32 = s1i32 + s2i32
		f195(ctx, s0i32, s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+240)]))
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l12 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l12
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l1
		s1i32 = l1
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
	lbl7:
		s0i32 = l3
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l2
		s1i32 = l10
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l9
		l1 = s0i32
	lbl3:
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+356)]))
		l8 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l8
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l10 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l10
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l8
		s1i32 = l8
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
	lbl8:
		s0i32 = l4
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l8 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0i32 = l4
		s1i32 = l4
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
		s0i32 = l11
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+408)]))
		l4 = s0i32
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l7
	s1i32 = 232
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+460)]))
	if s1i32 != 0 {
		s1i32 = l8
	} else {
		s1i32 = l7
		s2i32 = 408
		s1i32 = s1i32 + s2i32
		s2i32 = 4
		s1i32 = s1i32 | s2i32
		l8 = s1i32
		s2i32 = l4
		s1i32 = f46(ctx, s1i32, s2i32)
		l4 = s1i32
		s1i32 = l7
		s2i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+408)])) = uint32(s2i32)
		s1i32 = l7
		s2i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+460)])) = uint32(s2i32)
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+420)]))
	}
	s2i32 = l0
	s3i32 = 116
	s2i32 = s2i32 + s3i32
	f477(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s1i32 = l7
	s2i32 = 232
	s1i32 = s1i32 + s2i32
	f235(ctx, s0i32, s1i32)
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+232)]))
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l8 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l4
	s1i32 = l4
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
lbl10:
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l7
		s2i32 = 248
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
		s1i32 = l4
		s2i32 = 0
		ctx.Mem[int(s1i32+32)] = uint8(s2i32)
		s1i32 = l4
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint64(s2i64)
		s1i32 = l4
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
		s1i32 = l4
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
		s1i32 = l4
		s2i32 = l1
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
		s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l4
			s2i32 = l2
			s3i32 = l3
			s4i32 = l7
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+408)]))
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
		}
		s0i32 = l4
		s0i32 = f41(ctx, s0i32)
		goto lbl0
	}
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l7
	s1i32 = 224
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 0
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+48)]))
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
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+400)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+392)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+384)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+376)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+368)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 368
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 116
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	f69(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l1 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l4 = s0i32
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l4
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+260)])) = s1f32
	s0i32 = l7
	s1i32 = l1
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+256)])) = s1f32
	s0i32 = l7
	s1i32 = 368
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = 208
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = 248
	s2i32 = s2i32 + s3i32
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)]))
	l1 = s0i32
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])))
	l16 = s0i64
	s0i32 = l1
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])))
	l15 = s0i64
	s0i32 = l7
	s1i32 = l3
	s1i64 = int64(s1i32)
	l13 = s1i64
	s2i64 = 2147483647
	s3i64 = l13
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
	l14 = s1i64
	s2i64 = -2147483647
	s3i64 = l14
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+196)])) = uint32(s1i64)
	s0i32 = l7
	s1i32 = l2
	s1i64 = int64(s1i32)
	l14 = s1i64
	s2i64 = 2147483647
	s3i64 = l14
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
	l17 = s1i64
	s2i64 = -2147483647
	s3i64 = l17
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint32(s1i64)
	s0i32 = l7
	s1i64 = l13
	s2i64 = l15
	s1i64 = s1i64 + s2i64
	l13 = s1i64
	s2i64 = 2147483647
	s3i64 = l13
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
	l13 = s1i64
	s2i64 = -2147483647
	s3i64 = l13
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)])) = uint32(s1i64)
	s0i32 = l7
	s1i64 = l14
	s2i64 = l16
	s1i64 = s1i64 + s2i64
	l13 = s1i64
	s2i64 = 2147483647
	s3i64 = l13
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
	l13 = s1i64
	s2i64 = -2147483647
	s3i64 = l13
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint32(s1i64)
	s0i32 = l7
	s1i32 = l0
	s2i32 = 1064
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = l1
	s3i32 = 20
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s3i32 = int32(ctx.Mem[int(s3i32+40)])
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l1 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+220)]))
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l19 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l19
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l19 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l19
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l19 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l19
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl15
	}
	s1i32 = -2147483648
lbl15:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+260)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+216)]))
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l19 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l19
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l19 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l19
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l19 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l19
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl17
	}
	s1i32 = -2147483648
lbl17:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+256)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+212)]))
	s1f32 = float32(math.Floor(float64(s1f32)))
	l19 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l19
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l19 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l19
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l19 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l19
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl19
	}
	s1i32 = -2147483648
lbl19:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+252)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+208)]))
	s1f32 = float32(math.Floor(float64(s1f32)))
	l19 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l19
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l19 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l19
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l19 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l19
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl21
	}
	s1i32 = -2147483648
lbl21:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 176
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = 176
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = 248
	s2i32 = s2i32 + s3i32
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l7
	s1i32 = 176
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = 176
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = 192
	s2i32 = s2i32 + s3i32
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l7
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)])) = uint32(s1i32)
	s0i32 = l7
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
	s0i32 = 327682
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = 2097151
	s1i32 = s1i32 & s2i32
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl25
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+404)]))
	l1 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = l7
		s2i32 = 368
		s1i32 = s1i32 + s2i32
		s1i32 = f24(ctx, s1i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+404)])) = uint32(s1i32)
	}
	s0i32 = l1
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		goto lbl27
	}
	s0i32 = l7
	s1i32 = 368
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = 96
	s1i32 = s1i32 + s2i32
	s0i32 = f84(ctx, s0i32, s1i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl25
	}
lbl27:
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+404)]))
	l4 = s0i32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+400)]))
	l19 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+396)]))
	l20 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+392)]))
	l21 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+388)]))
	l22 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+384)]))
	l23 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+380)]))
	l24 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+376)]))
	l25 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+372)]))
	l26 = s0f32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+368)]))
	l27 = s0f32
	s0i32 = l7
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s1f32 = float32(s1i32)
	s2i32 = l3
	s2f32 = float32(s2i32)
	f117(ctx, s0i32, s1f32, s2f32)
	s0i32 = l7
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = 96
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = 16
	s2i32 = s2i32 + s3i32
	f69(ctx, s0i32, s1i32, s2i32)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+208)]))
	l28 = s0f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+216)]))
	l29 = s1f32
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+212)]))
	l30 = s0f32
	s1i32 = l7
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+220)]))
	l31 = s1f32
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl30
	}
	s0f32 = l28
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+192)]))
	l1 = s1i32
	s1f32 = float32(s1i32)
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)]))
	l2 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l13 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+204)]))
	l1 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+196)]))
	l3 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l14 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	s0i64 = l13
	s1i64 = l14
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	s0f32 = l31
	s1i32 = l1
	s1f32 = float32(s1i32)
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl30
	}
	s0f32 = l29
	s1i32 = l2
	s1f32 = float32(s1i32)
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl30
	}
	s0f32 = l30
	s1i32 = l3
	s1f32 = float32(s1i32)
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl29
	}
lbl30:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)]))
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		s1i32 = l2
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
		s0i32 = l2
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)])) = uint32(s1i32)
	}
	s0i32 = l7
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+260)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+268)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+276)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+284)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+252)])) = uint64(s1i64)
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
	s0i32 = l0
	s1i32 = 21916
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 21948
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 21940
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 21932
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 21924
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+192)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = s1f32
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+196)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = s1f32
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+200)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = s1f32
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+204)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = s1f32
	s0i32 = l0
	s1i32 = l7
	s2i32 = 56
	s1i32 = s1i32 + s2i32
	s2i32 = 1
	s3i32 = 0
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+32)]))
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
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+284)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+276)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+268)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+260)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+252)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s2i32 = 248
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)])) = uint32(s1i32)
lbl29:
	s0i32 = l7
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)])))
	l13 = s0i64
	s0i32 = l7
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])))
	l14 = s0i64
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+184)])))
	s2i64 = l14
	s1i64 = s1i64 - s2i64
	l14 = s1i64
	s2i64 = 2147483647
	s3i64 = l14
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
	l14 = s1i64
	s2i64 = -2147483647
	s3i64 = l14
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint32(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+188)])))
	s2i64 = l13
	s1i64 = s1i64 - s2i64
	l13 = s1i64
	s2i64 = 2147483647
	s3i64 = l13
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
	l13 = s1i64
	s2i64 = -2147483647
	s3i64 = l13
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+188)])) = uint32(s1i64)
	goto lbl24
lbl25:
	s0i32 = l7
	s1i64 = 8589934593
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+184)]))
	s2i32 = l7
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+176)]))
	s1i32 = s1i32 - s2i32
	s1i64 = int64(uint32(s1i32))
	s2i32 = l7
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+188)]))
	s3i32 = l7
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+180)]))
	s2i32 = s2i32 - s3i32
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = 56
	s1i32 = s1i32 + s2i32
	s2i32 = 0
	f252(ctx, s0i32, s1i32, s2i32)
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl32
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l4 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl32
	}
	s0i32 = l1
	f12(ctx, s0i32)
lbl32:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s0i32 = f585(ctx, s0i32)
	l1 = s0i32
	s1i32 = 0
	s2i32 = l7
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+176)]))
	s1i32 = s1i32 - s2i32
	s1f32 = float32(s1i32)
	s2i32 = 0
	s3i32 = l7
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+180)]))
	s2i32 = s2i32 - s3i32
	s2f32 = float32(s2i32)
	f169(ctx, s0i32, s1f32, s2f32)
	s0i32 = l1
	s1i32 = l7
	s2i32 = 368
	s1i32 = s1i32 + s2i32
	f190(ctx, s0i32, s1i32)
	s0i32 = l1
	s1i32 = l5
	s2f32 = 0
	s3f32 = 0
	s4i32 = 0
	f574(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0i32 = l7
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	f1569(ctx, s0i32, s1i32)
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l5 = s0i32
	s0i32 = 21952
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = 21948
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l19 = s0f32
	s0i32 = 21944
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l20 = s0f32
	s0i32 = 21940
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l21 = s0f32
	s0i32 = 21936
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l22 = s0f32
	s0i32 = 21932
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l23 = s0f32
	s0i32 = 21928
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l24 = s0f32
	s0i32 = 21924
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l25 = s0f32
	s0i32 = 21920
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l26 = s0f32
	s0i32 = 21916
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l27 = s0f32
	s0i32 = l7
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = l7
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+176)]))
	s1i32 = s1i32 - s2i32
	s1f32 = float32(s1i32)
	s2i32 = l3
	s3i32 = l7
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+180)]))
	s2i32 = s2i32 - s3i32
	s2f32 = float32(s2i32)
	f117(ctx, s0i32, s1f32, s2f32)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l2 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l1
	s1i32 = l1
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
lbl24:
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint32(s1i32)
	s0i32 = l0
	s1f32 = l19
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = s1f32
	s0i32 = l0
	s1f32 = l20
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = s1f32
	s0i32 = l0
	s1f32 = l21
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = s1f32
	s0i32 = l0
	s1f32 = l22
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = s1f32
	s0i32 = l0
	s1f32 = l23
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = s1f32
	s0i32 = l0
	s1f32 = l24
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = s1f32
	s0i32 = l0
	s1f32 = l25
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = s1f32
	s0i32 = l0
	s1f32 = l26
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = s1f32
	s0i32 = l0
	s1f32 = l27
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = s1f32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+408)]))
	l1 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+460)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 408
		s0i32 = s0i32 + s1i32
		s1i32 = 4
		s0i32 = s0i32 | s1i32
		l2 = s0i32
		s1i32 = l1
		s0i32 = f46(ctx, s0i32, s1i32)
		l1 = s0i32
		s0i32 = l7
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+408)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+460)])) = uint32(s1i32)
	}
	s0i32 = l7
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+224)]))
	s2i32 = l7
	s3i32 = 136
	s2i32 = s2i32 + s3i32
	f391(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	s1i32 = l7
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	f83(ctx, s0i32, s1i32)
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl34
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l2 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl34
	}
	s0i32 = l1
	s1i32 = l1
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
lbl34:
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l1 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l2 = s0i32
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+180)]))
	s1f32 = float32(s1i32)
	l19 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l7
	s1f32 = l19
	s2i32 = l2
	s2f32 = float32(s2i32)
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+176)]))
	s1f32 = float32(s1i32)
	l19 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l7
	s1f32 = l19
	s2i32 = l1
	s2f32 = float32(s2i32)
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l0
	s1i32 = l5
	s2i32 = 0
	s3i32 = l7
	s4i32 = 16
	s3i32 = s3i32 + s4i32
	s4i32 = l7
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+408)]))
	s5i32 = 1
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
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s2i32 = -64
	s1i32 = s1i32 - s2i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s1i32 = l1
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
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)])) = uint32(s1i32)
	}
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
		goto lbl23
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
lbl23:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)]))
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
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
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
lbl0:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+460)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = f23(ctx, s0i32)
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+460)])) = uint32(s1i32)
	}
	s0i32 = l9
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl37
	}
	s0i32 = l9
	s1i32 = l9
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
		goto lbl37
	}
	s0i32 = l9
	s1i32 = l9
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
lbl37:
	s0i32 = l7
	s1i32 = 464
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
