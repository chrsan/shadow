package internal

import (
	"math/bits"
	"unsafe"
)

func f2069(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
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
	var l25 int32
	_ = l25
	var l26 int64
	_ = l26
	var l27 int64
	_ = l27
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
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])))
	l26 = s0i64
	s0i32 = 0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l18 = s1i32
	s2i32 = 20
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s0i32 = l18
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl10
	case 1:
		goto lbl12
	case 2:
		goto lbl11
	case 3:
		goto lbl13
	case 4:
		goto lbl0
	case 5:
		goto lbl13
	case 6:
		goto lbl5
	case 7:
		goto lbl5
	case 8:
		goto lbl0
	case 9:
		goto lbl0
	case 10:
		goto lbl10
	case 11:
		goto lbl9
	case 12:
		goto lbl9
	case 13:
		goto lbl0
	case 14:
		goto lbl8
	case 15:
		goto lbl4
	case 16:
		goto lbl3
	case 17:
		goto lbl6
	case 18:
		goto lbl7
	case 19:
		goto lbl2
	default:
		goto lbl0
	}
lbl13:
	s0i32 = 719
	l9 = s0i32
	s0i32 = 720
	l10 = s0i32
	s0i32 = 721
	l11 = s0i32
	s0i32 = 722
	l12 = s0i32
	s0i32 = 724
	l13 = s0i32
	s0i32 = 725
	l14 = s0i32
	s0i32 = 726
	l15 = s0i32
	s0i32 = 723
	goto lbl1
lbl12:
	s0i32 = 727
	l9 = s0i32
	s0i32 = 728
	l10 = s0i32
	s0i32 = 729
	l11 = s0i32
	s0i32 = 730
	l12 = s0i32
	s0i32 = 732
	l13 = s0i32
	s0i32 = 733
	l14 = s0i32
	s0i32 = 734
	l15 = s0i32
	s0i32 = 731
	goto lbl1
lbl11:
	s0i32 = 735
	l9 = s0i32
	s0i32 = 736
	l10 = s0i32
	s0i32 = 737
	l11 = s0i32
	s0i32 = 738
	l12 = s0i32
	s0i32 = 740
	l13 = s0i32
	s0i32 = 741
	l14 = s0i32
	s0i32 = 742
	l15 = s0i32
	s0i32 = 739
	goto lbl1
lbl10:
	s0i32 = 743
	l9 = s0i32
	s0i32 = 744
	l10 = s0i32
	s0i32 = 745
	l11 = s0i32
	s0i32 = 746
	l12 = s0i32
	s0i32 = 748
	l13 = s0i32
	s0i32 = 749
	l14 = s0i32
	s0i32 = 750
	l15 = s0i32
	s0i32 = 747
	goto lbl1
lbl9:
	s0i32 = 751
	l9 = s0i32
	s0i32 = 752
	l10 = s0i32
	s0i32 = 753
	l11 = s0i32
	s0i32 = 754
	l12 = s0i32
	s0i32 = 756
	l13 = s0i32
	s0i32 = 757
	l14 = s0i32
	s0i32 = 758
	l15 = s0i32
	s0i32 = 755
	goto lbl1
lbl8:
	s0i32 = 759
	l9 = s0i32
	s0i32 = 760
	l10 = s0i32
	s0i32 = 761
	l11 = s0i32
	s0i32 = 762
	l12 = s0i32
	s0i32 = 764
	l13 = s0i32
	s0i32 = 765
	l14 = s0i32
	s0i32 = 766
	l15 = s0i32
	s0i32 = 763
	goto lbl1
lbl7:
	s0i32 = 767
	l9 = s0i32
	s0i32 = 768
	l10 = s0i32
	s0i32 = 769
	l11 = s0i32
	s0i32 = 770
	l12 = s0i32
	s0i32 = 772
	l13 = s0i32
	s0i32 = 773
	l14 = s0i32
	s0i32 = 774
	l15 = s0i32
	s0i32 = 771
	goto lbl1
lbl6:
	s0i32 = 775
	l9 = s0i32
	s0i32 = 776
	l10 = s0i32
	s0i32 = 777
	l11 = s0i32
	s0i32 = 778
	l12 = s0i32
	s0i32 = 780
	l13 = s0i32
	s0i32 = 781
	l14 = s0i32
	s0i32 = 782
	l15 = s0i32
	s0i32 = 779
	goto lbl1
lbl5:
	s0i32 = 783
	l9 = s0i32
	s0i32 = 784
	l10 = s0i32
	s0i32 = 785
	l11 = s0i32
	s0i32 = 786
	l12 = s0i32
	s0i32 = 788
	l13 = s0i32
	s0i32 = 789
	l14 = s0i32
	s0i32 = 790
	l15 = s0i32
	s0i32 = 787
	goto lbl1
lbl4:
	s0i32 = 791
	l9 = s0i32
	s0i32 = 792
	l10 = s0i32
	s0i32 = 793
	l11 = s0i32
	s0i32 = 794
	l12 = s0i32
	s0i32 = 796
	l13 = s0i32
	s0i32 = 797
	l14 = s0i32
	s0i32 = 798
	l15 = s0i32
	s0i32 = 795
	goto lbl1
lbl3:
	s0i32 = 799
	l9 = s0i32
	s0i32 = 800
	l10 = s0i32
	s0i32 = 801
	l11 = s0i32
	s0i32 = 802
	l12 = s0i32
	s0i32 = 804
	l13 = s0i32
	s0i32 = 805
	l14 = s0i32
	s0i32 = 806
	l15 = s0i32
	s0i32 = 803
	goto lbl1
lbl2:
	s0i32 = 807
	l9 = s0i32
	s0i32 = 808
	l10 = s0i32
	s0i32 = 809
	l11 = s0i32
	s0i32 = 810
	l12 = s0i32
	s0i32 = 812
	l13 = s0i32
	s0i32 = 813
	l14 = s0i32
	s0i32 = 814
	l15 = s0i32
	s0i32 = 811
lbl1:
	l21 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l2 = s0i32
	s1i32 = 1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	l7 = s2i32
	s3i32 = 2
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	s1i32 = l2
	s2i32 = 1
	if s1i32 < s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl14
	}
	s0i32 = 0
	s1i32 = l7
	s2i32 = 1
	if s1i32 < s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl14
	}
	s0i32 = 0
	s1i32 = l7
	s2i32 = l2
	s3i32 = l2
	s4i32 = l7
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l3 = s1i32
	s2i32 = 2
	if s1i32 < s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl14
	}
	s0i32 = 31
	s1i32 = l3
	s1i32 = int32(bits.LeadingZeros32(uint32(s1i32)))
	s0i32 = s0i32 - s1i32
lbl14:
	l17 = s0i32
	l3 = s0i32
lbl15:
	s0i32 = 0
	l16 = s0i32
	s0i32 = l2
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l4 = s0i32
		goto lbl16
	}
	s0i32 = 0
	l4 = s0i32
	s0i32 = l7
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = 0
	s1i32 = 31
	s2i32 = l7
	s3i32 = l2
	s4i32 = l2
	s5i32 = l7
	if s4i32 < s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l8 = s2i32
	s2i32 = int32(bits.LeadingZeros32(uint32(s2i32)))
	s1i32 = s1i32 - s2i32
	s2i32 = l8
	s3i32 = 2
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l3
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = l7
	s1i32 = l3
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l4 = s0i32
	s1i32 = 1
	s2i32 = l4
	s3i32 = 1
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
	l4 = s0i32
	s0i32 = l2
	s1i32 = l7
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l2 = s0i32
	s1i32 = 1
	s2i32 = l2
	s3i32 = 1
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
	l16 = s0i32
lbl16:
	s0i32 = l18
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 3728
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l4
	s2i32 = l16
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 * s1i32
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l7 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l2 = s0i32
		goto lbl15
	}
	s0i32 = 0
	l8 = s0i32
	s0i32 = l17
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l17
	s1i32 = 40
	s0i32 = s0i32 * s1i32
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s0i64 = int64(s0i32)
	s1i32 = l5
	s1i64 = int64(uint32(s1i32))
	s0i64 = s0i64 + s1i64
	l27 = s0i64
	s1i64 = 2147483647
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l27
	s0i32 = int32(s0i64)
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l1
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
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = 52
		s0i32 = f17(ctx, s0i32)
		l8 = s0i32
		s1i32 = l2
		s2i32 = l1
		s0i32 = f484(ctx, s0i32, s1i32, s2i32)
		goto lbl19
	}
	s0i32 = 52
	s0i32 = f17(ctx, s0i32)
	l8 = s0i32
	s1i32 = l2
	s2i32 = 2
	s1i32 = f44(ctx, s1i32, s2i32)
	s2i32 = l2
	s0i32 = f485(ctx, s0i32, s1i32, s2i32)
lbl19:
	s0i32 = l8
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = 23740
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		goto lbl21
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l2 = s0i32
	s0i32 = l8
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l2
	s1i32 = l2
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
		goto lbl21
	}
	s0i32 = l2
	f12(ctx, s0i32)
lbl21:
	s0i32 = l8
	s1i32 = l17
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l8
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l19 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l16 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l4 = s0i32
	s0i32 = l6
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l6
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
	s0i32 = l17
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l19
		s1i32 = l17
		s2i32 = 40
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l18
		s0i64 = int64(uint32(s0i32))
		s1i64 = l26
		s2i64 = 32
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s0i64 = s0i64 | s1i64
		l26 = s0i64
	lbl25:
		s0i32 = l16
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l13
			s1i32 = l11
			s2i32 = l4
			s3i32 = 1
			s2i32 = s2i32 & s3i32
			l2 = s2i32
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			s1i32 = l16
			s2i32 = 1
			if s1i32 == s2i32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				goto lbl26
			}
			s0i32 = l21
			s1i32 = l2
			if s1i32 == 0 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				goto lbl26
			}
			s0i32 = l10
			s1i32 = l15
			s2i32 = l4
			s3i32 = 1
			if s2i32 == s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			goto lbl26
		}
		s0i32 = l12
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 & s2i32
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl26
		}
		s0i32 = l9
		s1i32 = l14
		s2i32 = l4
		s3i32 = 1
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
	lbl26:
		l22 = s0i32
		s0i32 = l18
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 3728
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		s0i32 = 0
		l7 = s0i32
		s0i32 = l19
		s1i32 = l20
		s2i32 = 40
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s1i64 = l26
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l2
		s2i32 = l4
		s3i32 = 1
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l3 = s2i32
		s3i32 = 1
		s4i32 = l3
		s5i32 = 1
		if s4i32 > s5i32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		l4 = s2i32
		s1i32 = s1i32 * s2i32
		l23 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l4
		s1i64 = int64(uint32(s1i32))
		s2i32 = l16
		s3i32 = 1
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l2 = s2i32
		s3i32 = 1
		s4i32 = l2
		s5i32 = 1
		if s4i32 > s5i32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		l16 = s2i32
		s2i64 = int64(uint32(s2i32))
		s3i64 = 32
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s1i64 = s1i64 | s2i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l2 = s0i32
		s0i32 = l5
		s1i32 = l16
		s1f32 = float32(s1i32)
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
		s2f32 = float32(s2i32)
		s1f32 = s1f32 / s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l5
		s1i32 = l4
		s1f32 = float32(s1i32)
		s2i32 = l2
		s2f32 = float32(s2i32)
		s1f32 = s1f32 / s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l24 = s0i32
		s1i32 = 1
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l25 = s0i32
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
		s0i32 = l1
		l2 = s0i32
	lbl28:
		s0i32 = l2
		s1i32 = l3
		s2i32 = l24
		s3i32 = l4
		s4i32 = l22
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
		s0i32 = l3
		s1i32 = l25
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l2
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l16
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl28
		}
		s0i32 = l6
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l3 = s0i32
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		}
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l2 = s0i32
		s0i32 = l6
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl30
		}
		s0i32 = l2
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l3 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl30
		}
		s0i32 = l2
		f12(ctx, s0i32)
	lbl30:
		s0i32 = l6
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l16
		s2i32 = l23
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l20
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l20 = s0i32
		s1i32 = l17
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl25
		}
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l3 = s0i32
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	f12(ctx, s0i32)
lbl0:
	s0i32 = l6
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l8
	return s0i32
}
