package internal

import (
	"unsafe"
)

func f1572(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 float32
	_ = l4
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
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0f32 = s0f32 * s1f32
	l4 = s0f32
	s1f32 = l4
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
		s1i32 = 1
		if s0i32 < s1i32 {
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
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
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
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l2 = s0i32
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s0i32 = f94(ctx, s0i32)
		l2 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l3 = s0i32
		s0i32 = l2
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s1i32 = l3
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l3 = s0i32
		s0i32 = l2
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l2 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	lbl2:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l3 = s0i32
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l3
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
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		if s0i32 != 0 {
			goto lbl2
		}
	lbl1:
		s0i32 = l0
		s1i32 = l1
		s2i32 = 1
		s3i32 = 0
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+208)]))
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
	}
}
