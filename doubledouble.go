package doubledouble

// double double arithmetic
// based on https://gist.github.com/grd/4050062

type DoubleFloat64 [2]float64

func Add(a, b DoubleFloat64) (sum DoubleFloat64) {
	s1, s2 := twoSum(a[0], b[0])
	t1, t2 := twoSum(a[1], b[1])
	s2 += t1
	s1, s2 = quickTwoSum(s1, s2)
	s2 += t2
	sum[0], sum[1] = quickTwoSum(s1, s2)
	return
}

func (a DoubleFloat64) Cmp(b DoubleFloat64) int {
	switch {
	case a[0] < b[0]:
		return -1
	case a[0] > b[0]:
		return 1
	}
	switch {
	case a[1] < b[1]:
		return -1
	case a[1] > b[1]:
		return 1
	}
	return 0
}

func Div(a, b DoubleFloat64) (quo DoubleFloat64) {
	q1 := a[0] / b[0]
	r := Sub(a, Mul(DoubleFloat64{q1}, b))
	q2 := r[0] / b[0]
	t2 := Mul(DoubleFloat64{q2}, b)
	r = Sub(r, t2)
	quo[0], quo[1] = quickTwoSum(q1, q2)
	quo = Add(quo, DoubleFloat64{r[0] / b[0]})
	return
}

func (a DoubleFloat64) IsLT(b DoubleFloat64) bool {
	return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
}

func (a DoubleFloat64) IsLE(b DoubleFloat64) bool {
	return a[0] < b[0] || (a[0] == b[0] && a[1] <= b[1])
}

func (a DoubleFloat64) IsGT(b DoubleFloat64) bool {
	return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
}

func (a DoubleFloat64) IsGE(b DoubleFloat64) bool {
	return a[0] > b[0] || (a[0] == b[0] && a[1] >= b[1])
}

func (a DoubleFloat64) IsNE(b DoubleFloat64) bool {
	return a[0] != b[0] || a[1] != b[1]
}

func Mul(a, b DoubleFloat64) (prod DoubleFloat64) {
	p1, p2 := twoProd(a[0], b[0])
	p2 += a[0]*b[1] + a[1]*b[0]
	prod[0], prod[1] = quickTwoSum(p1, p2)
	return
}

func Sqr(f DoubleFloat64) (s DoubleFloat64) {
	p1, p2 := twoSqr(f[0])
	p2 += 2.0 * f[0] * f[1]
	p2 += f[1] * f[1]
	s[0], s[1] = quickTwoSum(p1, p2)
	return
}

func Sub(a, b DoubleFloat64) (diff DoubleFloat64) {
	s1, s2 := twoDiff(a[0], b[0])
	t1, t2 := twoDiff(a[1], b[1])
	s2 += t1
	s1, s2 = quickTwoSum(s1, s2)
	s2 += t2
	diff[0], diff[1] = quickTwoSum(s1, s2)
	return
}

func twoSum(a, b float64) (s, e float64) {
	s = a + b
	bb := s - a
	e = (a - (s - bb)) + (b - bb)
	return
}

func twoDiff(a, b float64) (d, e float64) {
	d = a - b
	bb := d - a
	e = (a - (d - bb)) - (b + bb)
	return
}

func quickTwoSum(a, b float64) (s, e float64) {
	s = a + b
	e = b - (s - a)
	return
}

const (
	splitter       = 134217729.0           // 2^27 + 1
	splitThreshold = 6.69692879491417e+299 // 2^996
)

func split(f float64) (hi, lo float64) {
	if f > splitThreshold || f < -splitThreshold {
		f *= 3.7252902984619140625e-09 // 2^-28
		temp := splitter * f
		hi = temp - (temp - f)
		lo = f - hi
		hi *= 268435456.0 // 2^28
		lo *= 268435456.0 // 2^28
	} else {
		temp := splitter * f
		hi = temp - (temp - f)
		lo = f - hi
	}
	return
}

func twoProd(a, b float64) (p, e float64) {
	p = a * b
	aHi, aLo := split(a)
	bHi, bLo := split(b)
	e = ((aHi*bHi - p) + aHi*bLo + aLo*bHi) + aLo*bLo
	return
}

func twoSqr(a float64) (q, e float64) {
	q = a * a
	hi, lo := split(a)
	e = ((hi*hi - q) + 2.0*hi*lo) + lo*lo
	return
}
