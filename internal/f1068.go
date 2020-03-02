package internal

import (
	"unsafe"
)

func f1068(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var s6i32 int32
	_ = s6i32
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 192
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = 21948
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 21940
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 21932
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 21924
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 21916
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l2
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l2
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 80
			s0i32 = s0i32 + s1i32
			s1i32 = l7
			s2i32 = l6
			s3i32 = 112
			s2i32 = s2i32 + s3i32
			s3i32 = l6
			s4i32 = 96
			s3i32 = s3i32 + s4i32
			s4i32 = l7
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+64)]))
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
			s0i32 = l6
			s1i32 = l7
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = -64
			s0i32 = s0i32 - s1i32
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l7
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l6
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l6
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 56
			s0i32 = s0i32 + s1i32
			s1i32 = l6
			s2i32 = 40
			s1i32 = s1i32 + s2i32
			s2i32 = l3
			s3i32 = l7
			s4i32 = l7
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+52)]))
			if int(s4i32) < 0 || int(s4i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s4i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s4i32].numParams != 1 {
				panic("argument count mismatch")
			}
			s3i32 = (*(*func(*Context, int32) int32)(table[s4i32].f()))(ctx, s3i32)
			s4i32 = l6
			s5i32 = 8
			s4i32 = s4i32 + s5i32
			f495(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
			goto lbl4
		}
		s0i32 = l6
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	lbl4:
		s0i32 = l6
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = 1
		l4 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1i32 = 1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	lbl6:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l4
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = 21948
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 21940
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 21932
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 21924
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 21916
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l2
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l2
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 80
			s0i32 = s0i32 + s1i32
			s1i32 = l5
			s2i32 = l6
			s3i32 = 112
			s2i32 = s2i32 + s3i32
			s3i32 = l6
			s4i32 = 96
			s3i32 = s3i32 + s4i32
			s4i32 = l5
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+64)]))
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
			s0i32 = l6
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = -64
			s0i32 = s0i32 - s1i32
			s1i32 = l5
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l5
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l6
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l6
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = 56
			s0i32 = s0i32 + s1i32
			s1i32 = l6
			s2i32 = 40
			s1i32 = s1i32 + s2i32
			s2i32 = l3
			s3i32 = l5
			s4i32 = l5
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+52)]))
			if int(s4i32) < 0 || int(s4i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s4i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s4i32].numParams != 1 {
				panic("argument count mismatch")
			}
			s3i32 = (*(*func(*Context, int32) int32)(table[s4i32].f()))(ctx, s3i32)
			s4i32 = l6
			s5i32 = 8
			s4i32 = s4i32 + s5i32
			f495(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
			goto lbl7
		}
		s0i32 = l6
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	lbl7:
		s0i32 = l6
		s1i32 = 24
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		f430(ctx, s0i32, s1i32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		goto lbl2
	}
	s0i32 = l5
	s1i32 = l2
	s2i32 = l5
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l4 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 21948
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21940
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21932
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21924
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21916
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = l6
		s3i32 = 112
		s2i32 = s2i32 + s3i32
		s3i32 = l6
		s4i32 = 56
		s3i32 = s3i32 + s4i32
		s4i32 = l6
		s5i32 = 96
		s4i32 = s4i32 + s5i32
		s5i32 = 1
		s6i32 = l7
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+60)]))
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
		s0i32 = l6
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl9
	}
	s0i32 = l6
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
lbl9:
	s0i32 = l6
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = 1
	l5 = s0i32
lbl11:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l5
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 21948
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21940
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21932
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21924
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 21916
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = l6
		s3i32 = 112
		s2i32 = s2i32 + s3i32
		s3i32 = l6
		s4i32 = 56
		s3i32 = s3i32 + s4i32
		s4i32 = l6
		s5i32 = 96
		s4i32 = s4i32 + s5i32
		s5i32 = 1
		s6i32 = l7
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+60)]))
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
		s0i32 = l6
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l6
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl12
	}
	s0i32 = l6
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
lbl12:
	s0i32 = l6
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	f430(ctx, s0i32, s1i32)
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
lbl2:
	s0i32 = l0
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
lbl0:
	s0i32 = l6
	s1i32 = 192
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
