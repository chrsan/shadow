package internal

import (
	"unsafe"
)

func f1388(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 float32, l5 float32, l6 int32) {
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
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		s0i32 = l6
		if s0i32 != 0 {
			s0i32 = l7
			s1i32 = 21948
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
			s0i32 = l7
			s1i32 = 21940
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
			s0i32 = l7
			s1i32 = -64
			s0i32 = s0i32 - s1i32
			s1i32 = 21932
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l7
			s1i32 = 21924
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
			s0i32 = l7
			s1i32 = 21916
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
			goto lbl1
		}
		s0i32 = l7
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1f32 = l4
		s2f32 = l5
		f117(ctx, s0i32, s1f32, s2f32)
		s0i32 = l7
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l7
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 116
		s1i32 = s1i32 + s2i32
		s2i32 = l7
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		f69(ctx, s0i32, s1i32, s2i32)
	lbl1:
		s0i32 = l7
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		s0i32 = l7
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l1
		s2i32 = 0
		s3i32 = 0
		s4i32 = l2
		s5i32 = l3
		s6i32 = l0
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+136)]))
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
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	}
	s0i32 = l7
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
