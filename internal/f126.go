package internal

func f126(ctx *Context, l0 float64) float32 {
	var l1 float64
	_ = l1
	var s0f32 float32
	_ = s0f32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	s0f64 = l0
	s1f64 = l0
	s0f64 = s0f64 * s1f64
	l0 = s0f64
	s1f64 = -0.499999997251031
	s0f64 = s0f64 * s1f64
	s1f64 = 1
	s0f64 = s0f64 + s1f64
	s1f64 = l0
	s2f64 = l0
	s1f64 = s1f64 * s2f64
	l1 = s1f64
	s2f64 = 0.04166662332373906
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	s1f64 = l0
	s2f64 = l1
	s1f64 = s1f64 * s2f64
	s2f64 = l0
	s3f64 = 2.439044879627741e-05
	s2f64 = s2f64 * s3f64
	s3f64 = -0.001388676377460993
	s2f64 = s2f64 + s3f64
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	s0f32 = float32(s0f64)
	return s0f32
}
