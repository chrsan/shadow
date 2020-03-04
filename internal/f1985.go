package internal

import (
	"math"
	"unsafe"
)

func f1985(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var l7 int64
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
	var s7i32 int32
	_ = s7i32
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	var s6i64 int64
	_ = s6i64
	var s7i64 int64
	_ = s7i64
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0i32 = l3
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = 65535
		s2i32 = 0
		s3i64 = 2147475456
		s4i64 = 0
		s5i32 = l1
		s5i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)])))
		l6 = s5i64
		s6i64 = 32767
		s5i64 = s5i64 & s6i64
		l7 = s5i64
		s5i32 = int32(s5i64)
		s6i32 = 1023
		if uint32(s5i32) > uint32(s6i32) {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			// s3i64 = s3i64
		} else {
			s3i64 = s4i64
		}
		s4i64 = l7
		s5i64 = 13
		s4i64 = s4i64 << (uint64(s5i64) & 63)
		s5i64 = 939524096
		s4i64 = s4i64 + s5i64
		s3i64 = s3i64 & s4i64
		s4i64 = l6
		s5i64 = 16
		s4i64 = s4i64 << (uint64(s5i64) & 63)
		s5i64 = 2147483648
		s4i64 = s4i64 & s5i64
		s3i64 = s3i64 | s4i64
		s3i32 = int32(s3i64)
		s3f32 = math.Float32frombits(uint32(s3i32))
		s4i64 = 2147475456
		s5i64 = 0
		s6i32 = l2
		s6i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)])))
		l6 = s6i64
		s7i64 = 32767
		s6i64 = s6i64 & s7i64
		l7 = s6i64
		s6i32 = int32(s6i64)
		s7i32 = 1023
		if uint32(s6i32) > uint32(s7i32) {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4i64 = s4i64
		} else {
			s4i64 = s5i64
		}
		s5i64 = l7
		s6i64 = 13
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 939524096
		s5i64 = s5i64 + s6i64
		s4i64 = s4i64 & s5i64
		s5i64 = l6
		s6i64 = 16
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 2147483648
		s5i64 = s5i64 & s6i64
		s4i64 = s4i64 | s5i64
		s4i32 = int32(s4i64)
		s4f32 = math.Float32frombits(uint32(s4i32))
		s3f32 = s3f32 + s4f32
		s4i64 = 2147475456
		s5i64 = 0
		s6i32 = l1
		s6i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s6i32+2)])))
		l6 = s6i64
		s7i64 = 32767
		s6i64 = s6i64 & s7i64
		l7 = s6i64
		s6i32 = int32(s6i64)
		s7i32 = 1023
		if uint32(s6i32) > uint32(s7i32) {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4i64 = s4i64
		} else {
			s4i64 = s5i64
		}
		s5i64 = l7
		s6i64 = 13
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 939524096
		s5i64 = s5i64 + s6i64
		s4i64 = s4i64 & s5i64
		s5i64 = l6
		s6i64 = 16
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 2147483648
		s5i64 = s5i64 & s6i64
		s4i64 = s4i64 | s5i64
		s4i32 = int32(s4i64)
		s4f32 = math.Float32frombits(uint32(s4i32))
		s3f32 = s3f32 + s4f32
		s4i64 = 2147475456
		s5i64 = 0
		s6i32 = l2
		s6i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s6i32+2)])))
		l6 = s6i64
		s7i64 = 32767
		s6i64 = s6i64 & s7i64
		l7 = s6i64
		s6i32 = int32(s6i64)
		s7i32 = 1023
		if uint32(s6i32) > uint32(s7i32) {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4i64 = s4i64
		} else {
			s4i64 = s5i64
		}
		s5i64 = l7
		s6i64 = 13
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 939524096
		s5i64 = s5i64 + s6i64
		s4i64 = s4i64 & s5i64
		s5i64 = l6
		s6i64 = 16
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 2147483648
		s5i64 = s5i64 & s6i64
		s4i64 = s4i64 | s5i64
		s4i32 = int32(s4i64)
		s4f32 = math.Float32frombits(uint32(s4i32))
		s3f32 = s3f32 + s4f32
		s4f32 = 0.25
		s3f32 = s3f32 * s4f32
		s3i32 = int32(math.Float32bits(s3f32))
		l5 = s3i32
		s4i32 = 2147475456
		s3i32 = s3i32 & s4i32
		s4i32 = 947904511
		if uint32(s3i32) > uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l5
		s3i32 = -134217728
		s2i32 = s2i32 - s3i32
		s3i32 = 13
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 & s2i32
		s2i32 = l5
		s3i32 = 16
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 32768
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l2
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
