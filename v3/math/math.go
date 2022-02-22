package math

func Pow(v, n int) float64 {
	p := n
	if n < 1 {
		p = -n
	}
	r := 1.0
	for i := 0; i < p; i++ {
		r *= float64(v)
	}
	if n < 1 {
		r = 1 / r
	}
	return r
}
