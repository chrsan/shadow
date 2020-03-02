package internal

import (
	"unsafe"
)

func f316(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var s10i32 int32
	_ = s10i32
	var s11i32 int32
	_ = s11i32
	s0i32 = l2
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s1i32 = 128
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl2:
		s0i32 = l0
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = 2
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l0 = s0i32
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l8 = s0i32
			s1i32 = 1
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l2 = s0i32
			if s0i32 != 0 {
			lbl5:
				s0i32 = l1
				s1i32 = l2
				l7 = s1i32
				s2i32 = -1
				s1i32 = s1i32 + s2i32
				l2 = s1i32
				s2i32 = 2
				s1i32 = s1i32 << (uint32(s2i32) & 31)
				s0i32 = s0i32 + s1i32
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
				l9 = s0i32
				s0i32 = l2
				l5 = s0i32
				s0i32 = l7
				s1i32 = 1
				s0i32 = s0i32 << (uint32(s1i32) & 31)
				l4 = s0i32
				s1i32 = l8
				if uint32(s0i32) > uint32(s1i32) {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl6
				}
				s0i32 = l9
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
				l11 = s0i32
			lbl7:
				s0i32 = l9
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
				s1i32 = l9
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
				s2i32 = l11
				s3i32 = l11
				s4i32 = l1
				s5i32 = l4
				s6i32 = l0
				if uint32(s5i32) > uint32(s6i32) {
					s5i32 = 1
				} else {
					s5i32 = 0
				}
				if s5i32 != 0 {
					s5i32 = l4
				} else {
					s5i32 = l4
					s6i32 = l1
					s7i32 = l4
					s8i32 = 2
					s7i32 = s7i32 << (uint32(s8i32) & 31)
					s6i32 = s6i32 + s7i32
					l5 = s6i32
					s7i32 = -4
					s6i32 = s6i32 + s7i32
					s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
					l3 = s6i32
					s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+16)]))
					s7i32 = l3
					s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+12)]))
					s8i32 = l3
					s8i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s8i32+28)]))
					l10 = s8i32
					s9i32 = l10
					s10i32 = l5
					s10i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s10i32+0)]))
					l3 = s10i32
					s10i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s10i32+28)]))
					l5 = s10i32
					if s9i32 == s10i32 {
						s9i32 = 1
					} else {
						s9i32 = 0
					}
					l10 = s9i32
					if s9i32 != 0 {
						// s7i32 = s7i32
					} else {
						s7i32 = s8i32
					}
					l12 = s7i32
					s8i32 = l12
					s9i32 = l3
					s9i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s9i32+12)]))
					s10i32 = l5
					s11i32 = l10
					if s11i32 != 0 {
						// s9i32 = s9i32
					} else {
						s9i32 = s10i32
					}
					l5 = s9i32
					if s8i32 == s9i32 {
						s8i32 = 1
					} else {
						s8i32 = 0
					}
					l10 = s8i32
					if s8i32 != 0 {
						// s6i32 = s6i32
					} else {
						s6i32 = s7i32
					}
					s7i32 = l3
					s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+16)]))
					s8i32 = l5
					s9i32 = l10
					if s9i32 != 0 {
						// s7i32 = s7i32
					} else {
						s7i32 = s8i32
					}
					if s6i32 < s7i32 {
						s6i32 = 1
					} else {
						s6i32 = 0
					}
					s5i32 = s5i32 | s6i32
				}
				l3 = s5i32
				s6i32 = -1
				s5i32 = s5i32 + s6i32
				l5 = s5i32
				s6i32 = 2
				s5i32 = s5i32 << (uint32(s6i32) & 31)
				s4i32 = s4i32 + s5i32
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
				l4 = s4i32
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+28)]))
				l10 = s4i32
				if s3i32 == s4i32 {
					s3i32 = 1
				} else {
					s3i32 = 0
				}
				l12 = s3i32
				if s3i32 != 0 {
					// s1i32 = s1i32
				} else {
					s1i32 = s2i32
				}
				l13 = s1i32
				s2i32 = l13
				s3i32 = l4
				s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
				s4i32 = l10
				s5i32 = l12
				if s5i32 != 0 {
					// s3i32 = s3i32
				} else {
					s3i32 = s4i32
				}
				l10 = s3i32
				if s2i32 == s3i32 {
					s2i32 = 1
				} else {
					s2i32 = 0
				}
				l12 = s2i32
				if s2i32 != 0 {
					// s0i32 = s0i32
				} else {
					s0i32 = s1i32
				}
				s1i32 = l4
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
				s2i32 = l10
				s3i32 = l12
				if s3i32 != 0 {
					// s1i32 = s1i32
				} else {
					s1i32 = s2i32
				}
				if s0i32 >= s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					s0i32 = l7
					s1i32 = -1
					s0i32 = s0i32 + s1i32
					l5 = s0i32
					goto lbl6
				}
				s0i32 = l7
				s1i32 = 2
				s0i32 = s0i32 << (uint32(s1i32) & 31)
				s1i32 = l1
				s0i32 = s0i32 + s1i32
				s1i32 = -4
				s0i32 = s0i32 + s1i32
				s1i32 = l4
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
				s0i32 = l3
				l7 = s0i32
				s1i32 = 1
				s0i32 = s0i32 << (uint32(s1i32) & 31)
				l4 = s0i32
				s1i32 = l8
				if uint32(s0i32) <= uint32(s1i32) {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl7
				}
			lbl6:
				s0i32 = l1
				s1i32 = l5
				s2i32 = 2
				s1i32 = s1i32 << (uint32(s2i32) & 31)
				s0i32 = s0i32 + s1i32
				s1i32 = l9
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
				s0i32 = l2
				if s0i32 != 0 {
					goto lbl5
				}
				s0i32 = l6
				if s0i32 == 0 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl0
				}
			}
		lbl10:
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l2 = s0i32
			s0i32 = l1
			s1i32 = l1
			s2i32 = l0
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			l4 = s1i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = l0
			f1804(ctx, s0i32, s1i32)
			s0i32 = l0
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l0 = s0i32
			if s0i32 != 0 {
				goto lbl10
			}
			goto lbl0
		}
		s0i32 = l1
		s1i32 = l6
		s2i32 = 1
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 2147483644
		s1i32 = s1i32 & s2i32
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
		s0i32 = l4
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l1
		s1i32 = l2
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l3 = s0i32
			goto lbl11
		}
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l6 = s0i32
		s0i32 = l1
		l3 = s0i32
		l4 = s0i32
	lbl13:
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l5
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		l9 = s2i32
		s3i32 = l6
		s4i32 = l9
		if s3i32 == s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l9 = s3i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l8 = s1i32
		s2i32 = l8
		s3i32 = l7
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
		s4i32 = l6
		s5i32 = l9
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		l9 = s3i32
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l8 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l9
		s3i32 = l8
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l5
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l3 = s0i32
		}
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl13
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
	lbl11:
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s0i32 = l3
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		s2i32 = l3
		s3i32 = -4
		s2i32 = s2i32 + s3i32
		f316(ctx, s0i32, s1i32, s2i32)
		s0i32 = l2
		s1i32 = l3
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s1i32 = 128
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	l3 = s0i32
lbl15:
	s0i32 = l4
	l0 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	l9 = s0i32
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l7
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	l5 = s2i32
	s3i32 = l5
	s4i32 = l3
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	l3 = s4i32
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+28)]))
	l6 = s4i32
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	l8 = s3i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l11 = s1i32
	s2i32 = l11
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = l6
	s5i32 = l8
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	l6 = s3i32
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l8 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l6
	s3i32 = l8
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl17:
		s0i32 = l4
		s1i32 = l4
		s2i32 = -4
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l1
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		s0i32 = l4
		s1i32 = -8
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l3
		l4 = s0i32
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l5
		s3i32 = l5
		s4i32 = l6
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		l6 = s4i32
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+28)]))
		l8 = s4i32
		if s3i32 == s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l11 = s3i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l10 = s1i32
		s2i32 = l10
		s3i32 = l6
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
		s4i32 = l8
		s5i32 = l11
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		l8 = s3i32
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l11 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l8
		s3i32 = l11
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
	lbl18:
		s0i32 = l3
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l0
	l3 = s0i32
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l2
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
lbl0:
}
