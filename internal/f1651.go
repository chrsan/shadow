package internal

import (
	"math"
	"unsafe"
)

func f1651(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 int64
	_ = l12
	var l13 int64
	_ = l13
	var l14 float32
	_ = l14
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
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0i32 = ctx.g0
	s1i32 = 528
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+520)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+512)])) = uint64(s1i64)
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l5
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		f101(ctx, s0i32, s1i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l3 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l4 = s0i32
		s0i32 = l5
		s1i32 = 472
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		f101(ctx, s0i32, s1i32)
		s0i32 = l2
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])))
		l10 = s0i64
		s0i32 = l2
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])))
		l11 = s0i64
		s0i32 = l5
		s1i32 = l1
		s2i32 = l4
		s1i32 = s1i32 - s2i32
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = uint32(s1i32)
		s0i32 = l5
		s1i64 = l11
		s2i32 = l1
		s2i64 = int64(s2i32)
		s1i64 = s1i64 + s2i64
		l11 = s1i64
		s2i64 = 2147483647
		s3i64 = l11
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
		l11 = s1i64
		s2i64 = -2147483647
		s3i64 = l11
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
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+288)])) = uint32(s1i64)
		s0i32 = l5
		s1i32 = l3
		s2i32 = l5
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+476)]))
		s1i32 = s1i32 - s2i32
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+284)])) = uint32(s1i32)
		s0i32 = l5
		s1i64 = l10
		s2i32 = l1
		s2i64 = int64(s2i32)
		s1i64 = s1i64 + s2i64
		l10 = s1i64
		s2i64 = 2147483647
		s3i64 = l10
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
		l10 = s1i64
		s2i64 = -2147483647
		s3i64 = l10
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
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)])) = uint32(s1i64)
		s0i32 = l5
		s1i32 = 280
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s2i32 = 280
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s3i32 = 512
		s2i32 = s2i32 + s3i32
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l5
		s1i32 = 472
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = l5
		s3i32 = 280
		s2i32 = s2i32 + s3i32
		s3i32 = 0
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+188)]))
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
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+472)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l5
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 13195221663744
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+472)]))
		s2i32 = 0
		s3i32 = 0
		s4i32 = l0
		s5i32 = 0
		s6i32 = 21916
		s7i32 = l2
		s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+0)]))
		s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+176)]))
		if int(s7i32) < 0 || int(s7i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s7i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s7i32].numParams != 7 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32, int32, int32))(table[s7i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		s0i32 = l0
		s0i32 = f23(ctx, s0i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+472)]))
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
		goto lbl0
	}
	s0i32 = l5
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+504)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+496)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+488)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+480)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+472)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+464)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+456)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+448)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+440)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+432)])) = uint64(s1i64)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l7 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l4
		s1i32 = f24(ctx, s1i32)
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l7
	s1i32 = 12
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = f351(ctx, s0i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	}
	s0i32 = l5
	s1i32 = 21948
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+504)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21940
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+496)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21932
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+488)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21924
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+480)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21916
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+472)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+440)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+448)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+456)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+464)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+432)])) = uint64(s1i64)
	goto lbl3
lbl4:
	s0i32 = l4
	s1i32 = l5
	s2i32 = 424
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 472
	s2i32 = s2i32 + s3i32
	s0i32 = f347(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = l5
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+424)]))
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+428)]))
		f240(ctx, s0i32, s1f32, s2f32)
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+152)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+464)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+144)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+456)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+136)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+448)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+440)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+432)])) = uint64(s1i64)
		goto lbl3
	}
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+504)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+496)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+488)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+480)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+472)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21948
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+464)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21940
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+456)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21932
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+448)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21924
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+440)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21916
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+432)])) = uint64(s1i64)
lbl3:
	s0i32 = l5
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+416)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+408)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+400)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+392)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+384)])) = uint64(s1i64)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+508)]))
	l4 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l5
		s2i32 = 472
		s1i32 = s1i32 + s2i32
		s1i32 = f24(ctx, s1i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+508)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+416)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+408)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+400)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+392)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+384)])) = uint64(s1i64)
		goto lbl8
	}
	s0i32 = l5
	s1i32 = 472
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 384
	s1i32 = s1i32 + s2i32
	s0i32 = f84(ctx, s0i32, s1i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl8:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l4 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l7 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = l5
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1f32 = float32(s1i32)
	l14 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+284)])) = s1f32
	s0i32 = l5
	s1f32 = l14
	s2i32 = l7
	s2f32 = float32(s2i32)
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)])) = s1f32
	s0i32 = l5
	s1i32 = l6
	s1f32 = float32(s1i32)
	l14 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = s1f32
	s0i32 = l5
	s1f32 = l14
	s2i32 = l4
	s2f32 = float32(s2i32)
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+288)])) = s1f32
	s0i32 = l5
	s1i32 = 384
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 120
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 280
	s2i32 = s2i32 + s3i32
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l14 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l14
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl10
	}
	s1i32 = -2147483648
lbl10:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+380)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l14 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l14
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl12
	}
	s1i32 = -2147483648
lbl12:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+376)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
	s1f32 = float32(math.Floor(float64(s1f32)))
	l14 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l14
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl14
	}
	s1i32 = -2147483648
lbl14:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+372)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
	s1f32 = float32(math.Floor(float64(s1f32)))
	l14 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l14
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl16
	}
	s1i32 = -2147483648
lbl16:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+368)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 352
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l5
	s3i32 = 368
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 432
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = 368
	s4i32 = s4i32 + s5i32
	f496(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+352)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+356)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+284)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+360)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+288)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+364)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)])) = s1f32
	s0i32 = l5
	s1i32 = 472
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 120
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 280
	s2i32 = s2i32 + s3i32
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l14 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l14
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl18
	}
	s1i32 = -2147483648
lbl18:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+348)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l14 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l14
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl20
	}
	s1i32 = -2147483648
lbl20:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+344)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
	s1f32 = float32(math.Floor(float64(s1f32)))
	l14 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l14
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl22
	}
	s1i32 = -2147483648
lbl22:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+340)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
	s1f32 = float32(math.Floor(float64(s1f32)))
	l14 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l14
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl24
	}
	s1i32 = -2147483648
lbl24:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+336)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 120
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	f101(ctx, s0i32, s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
	l4 = s0i32
	s0i32 = l5
	s1i32 = 280
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	f101(ctx, s0i32, s1i32)
	s0i32 = l5
	s1i32 = 0
	s2i32 = l4
	s1i32 = s1i32 - s2i32
	s1i64 = int64(s1i32)
	l10 = s1i64
	s2i32 = l5
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+336)])))
	s1i64 = s1i64 + s2i64
	l11 = s1i64
	s2i64 = 2147483647
	s3i64 = l11
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
	l11 = s1i64
	s2i64 = -2147483647
	s3i64 = l11
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
	s0i32 = l5
	s1i32 = l5
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+344)])))
	s2i64 = l10
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+344)])) = uint32(s1i64)
	s0i32 = l5
	s1i32 = 0
	s2i32 = l5
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+284)]))
	s1i32 = s1i32 - s2i32
	s1i64 = int64(s1i32)
	l10 = s1i64
	s2i32 = l5
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+340)])))
	s1i64 = s1i64 + s2i64
	l11 = s1i64
	s2i64 = 2147483647
	s3i64 = l11
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
	l11 = s1i64
	s2i64 = -2147483647
	s3i64 = l11
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
	s0i32 = l5
	s1i32 = l5
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+348)])))
	s2i64 = l10
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+348)])) = uint32(s1i64)
	s0i32 = l5
	s1i32 = 336
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 336
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 512
	s2i32 = s2i32 + s3i32
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 328
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l5
	s3i32 = 336
	s2i32 = s2i32 + s3i32
	s3i32 = 0
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+188)]))
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
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i32
	s0i32 = l5
	s1i32 = 280
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 13195221663744
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l4
	l7 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+508)]))
	l4 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l5
		s2i32 = 472
		s1i32 = s1i32 + s2i32
		s1i32 = f24(ctx, s1i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+508)])) = uint32(s1i32)
	}
	s0i32 = l6
	s1i32 = 4
	s2i32 = l6
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l9 = s0i32
	s0i32 = l4
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = -769
		s1i32 = s1i32 & s2i32
		s2i32 = 768
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
		l4 = s0i32
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+364)]))
		s2i32 = l5
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+356)]))
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+360)]))
		s2i32 = l5
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+352)]))
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l9
		s3i32 = l8
		s4i32 = l5
		s5i32 = 120
		s4i32 = s4i32 + s5i32
		s5i32 = 2
		s6i32 = 0
		s7i32 = l4
		s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+0)]))
		s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+44)]))
		if int(s7i32) < 0 || int(s7i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s7i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s7i32].numParams != 7 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32, int32, int32))(table[s7i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		l4 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl27
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l6 = s0i32
		s0i32 = l5
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = 13195221663744
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = 0
		f196(ctx, s0i32, s1i32)
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = -261121
		s1i32 = s1i32 & s2i32
		s2i32 = 1024
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l4
		f606(ctx, s0i32, s1i32)
		s0i32 = l4
		s0i32 = f23(ctx, s0i32)
		s0i32 = l6
		s1i32 = 0
		s2i32 = l5
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+352)]))
		s1i32 = s1i32 - s2i32
		s1f32 = float32(s1i32)
		s2i32 = 0
		s3i32 = l5
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+356)]))
		s2i32 = s2i32 - s3i32
		s2f32 = float32(s2i32)
		f169(ctx, s0i32, s1f32, s2f32)
		s0i32 = l6
		s1i32 = l5
		s2i32 = 384
		s1i32 = s1i32 + s2i32
		f190(ctx, s0i32, s1i32)
		s0i32 = l5
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		f101(ctx, s0i32, s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l4 = s0i32
		s0i32 = l5
		s1i32 = 240
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		f101(ctx, s0i32, s1i32)
		s0i32 = l6
		s1i32 = l4
		s1f32 = float32(s1i32)
		s2i32 = l5
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+244)]))
		s2f32 = float32(s2i32)
		f169(ctx, s0i32, s1f32, s2f32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
		l0 = s0i32
		s0i32 = ctx.g0
		s1i32 = 16
		s0i32 = s0i32 - s1i32
		l4 = s0i32
		ctx.g0 = s0i32
		s0i32 = l5
		s1i32 = 40
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 0
		s3i32 = l0
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
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
		l0 = s0i32
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+336)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+240)])) = s1f32
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+340)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+244)])) = s1f32
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+344)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = s1f32
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+348)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+252)])) = s1f32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l4 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl31
		}
		s0i32 = l5
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = s1f32
		s0i32 = l5
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = s1f32
		s0i32 = l5
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = s1f32
		s0i32 = l5
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = s1f32
		s0i32 = l6
		s1i32 = l4
		s2i32 = l5
		s3i32 = 120
		s2i32 = s2i32 + s3i32
		s3i32 = l5
		s4i32 = 240
		s3i32 = s3i32 + s4i32
		s4i32 = l7
		s5i32 = 0
		f599(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl31
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l4 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl31
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
	lbl31:
		s0i32 = l5
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		l0 = s1i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
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
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l4 = s0i32
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l4
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l4
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
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l4 = s0i32
		s0i32 = l5
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
		l0 = s0i32
		s0i32 = l5
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)])) = uint32(s1i32)
		s0i32 = l0
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl33
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l4 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
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
		}
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl33
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l4 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl33
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
	lbl33:
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl28
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l4 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl28
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
		goto lbl28
	}
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+344)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+360)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+336)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+352)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 120
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	f101(ctx, s0i32, s1i32)
	s0i32 = l5
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])))
	l10 = s0i64
	s0i32 = l5
	s1i32 = 240
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	f101(ctx, s0i32, s1i32)
	s0i32 = l5
	s1i64 = l10
	s2i32 = l5
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+360)])))
	s1i64 = s1i64 + s2i64
	l11 = s1i64
	s2i64 = 2147483647
	s3i64 = l11
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
	l11 = s1i64
	s2i64 = -2147483647
	s3i64 = l11
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+360)])) = uint32(s1i64)
	s0i32 = l5
	s1i64 = l10
	s2i32 = l5
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+352)])))
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+352)])) = uint32(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+244)])))
	l10 = s1i64
	s2i32 = l5
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+356)])))
	s1i64 = s1i64 + s2i64
	l11 = s1i64
	s2i64 = 2147483647
	s3i64 = l11
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
	l11 = s1i64
	s2i64 = -2147483647
	s3i64 = l11
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+356)])) = uint32(s1i64)
	s0i32 = l5
	s1i64 = l10
	s2i32 = l5
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+364)])))
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+364)])) = uint32(s1i64)
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -769
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
lbl28:
	s0i32 = 0
	l4 = s0i32
	s0i32 = l5
	s1i32 = 0
	s2i32 = l5
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+352)]))
	s1i32 = s1i32 - s2i32
	l0 = s1i32
	s1i64 = int64(s1i32)
	l10 = s1i64
	s2i32 = l5
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+368)])))
	s1i64 = s1i64 + s2i64
	l11 = s1i64
	s2i64 = 2147483647
	s3i64 = l11
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
	l11 = s1i64
	s2i64 = -2147483647
	s3i64 = l11
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+368)])) = uint32(s1i64)
	s0i32 = l5
	s1i32 = 0
	s2i32 = l5
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+356)]))
	s1i32 = s1i32 - s2i32
	l6 = s1i32
	s1i64 = int64(s1i32)
	l11 = s1i64
	s2i32 = l5
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+372)])))
	s1i64 = s1i64 + s2i64
	l12 = s1i64
	s2i64 = 2147483647
	s3i64 = l12
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
	l12 = s1i64
	s2i64 = -2147483647
	s3i64 = l12
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+372)])) = uint32(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+376)])))
	s2i64 = l10
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+376)])) = uint32(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+380)])))
	s2i64 = l11
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+380)])) = uint32(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+464)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+272)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+456)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+264)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+448)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+256)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+440)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+432)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+240)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 240
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1f32 = float32(s1i32)
	s2i32 = l6
	s2f32 = float32(s2i32)
	f142(ctx, s0i32, s1f32, s2f32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
	l0 = s0i32
	s0i32 = l5
	s1i32 = 21948
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21940
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21932
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 21924
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
	s0i32 = 21916
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0i64
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+248)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+256)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+264)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+272)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = l10
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+240)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+376)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+208)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+368)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+220)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+216)])) = uint32(s1i32)
	s0i32 = l0
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	}
	s0i64 = 0
	l10 = s0i64
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+232)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+228)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l5
	s3i32 = 120
	s2i32 = s2i32 + s3i32
	f352(ctx, s0i32, s1i32, s2i32)
	s0i64 = 0
	l11 = s0i64
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l10 = s0i64
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i64 = l10
		s1i64 = 32
		s0i64 = s0i64 >> (uint64(s1i64) & 63)
		l11 = s0i64
		s0i64 = l10
		s1i64 = 32
		s0i64 = s0i64 << (uint64(s1i64) & 63)
		s1i64 = 32
		s0i64 = s0i64 >> (uint64(s1i64) & 63)
		l10 = s0i64
		s0i32 = l0
		l4 = s0i32
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
	l0 = s0i32
	s0i32 = l5
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)])) = uint32(s1i32)
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl37
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
		goto lbl37
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
lbl37:
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl38
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
		goto lbl38
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
lbl38:
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
	if s0i32 != 0 {
		s0i32 = l5
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+352)])))
		l12 = s0i64
		s0i32 = l5
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+356)])))
		l13 = s0i64
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+504)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+496)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+488)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+480)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+472)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1i32 = s1i32 - s2i32
		s1f32 = float32(s1i32)
		s2i32 = 0
		s3i32 = l3
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
		s2i32 = s2i32 - s3i32
		s2f32 = float32(s2i32)
		f142(ctx, s0i32, s1f32, s2f32)
		s0i32 = l5
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+328)]))
		l0 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l5
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l5
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l5
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
		s0i32 = ctx.g0
		s1i32 = 16
		s0i32 = s0i32 - s1i32
		l3 = s0i32
		ctx.g0 = s0i32
		s0i32 = l5
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 0
		s3i32 = l0
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
		s0i32 = l3
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		ctx.g0 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l3 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
		l0 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l4 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l6 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l8 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l0 = s0i32
		s0i32 = l5
		s1i64 = l11
		s2i64 = l13
		s1i64 = s1i64 + s2i64
		l11 = s1i64
		s2i64 = 2147483647
		s3i64 = l11
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
		l11 = s1i64
		s2i64 = -2147483647
		s3i64 = l11
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
		s1i32 = int32(s1i64)
		s1f32 = float32(s1i32)
		l14 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l5
		s1f32 = l14
		s2i32 = l8
		s3i32 = l6
		s2i32 = s2i32 - s3i32
		s2f32 = float32(s2i32)
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l5
		s1i64 = l10
		s2i64 = l12
		s1i64 = s1i64 + s2i64
		l10 = s1i64
		s2i64 = 2147483647
		s3i64 = l10
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
		l10 = s1i64
		s2i64 = -2147483647
		s3i64 = l10
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
		s1i32 = int32(s1i64)
		s1f32 = float32(s1i32)
		l14 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l5
		s1f32 = l14
		s2i32 = l4
		s3i32 = l0
		s2i32 = s2i32 - s3i32
		s2f32 = float32(s2i32)
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l2
		s1i32 = l3
		s2i32 = l5
		s3i32 = 24
		s2i32 = s2i32 + s3i32
		s3i32 = l5
		s4i32 = l7
		s5i32 = 0
		s6i32 = l2
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
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl40
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l3 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl40
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
	lbl40:
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+228)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl27
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
		goto lbl27
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
lbl27:
	s0i32 = l7
	s0i32 = f23(ctx, s0i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+328)]))
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
	s0i32 = l5
	s1i32 = 528
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
