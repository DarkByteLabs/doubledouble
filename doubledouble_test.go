package doubledouble_test

import (
	"testing"

	"../doubledouble"
)

func BenchmarkAdd(b *testing.B) {
	x, y := doubledouble.DoubleFloat64{1.1}, doubledouble.DoubleFloat64{2.001}
	for i := 0; i < b.N; i++ {
		x = doubledouble.Add(x, y)
	}
}

func TestAdd(t *testing.T) {
	var a, b doubledouble.DoubleFloat64
	a = doubledouble.Add(a, b)
	if a != (doubledouble.DoubleFloat64{}) {
		t.Errorf("got %v; want %v", a, doubledouble.DoubleFloat64{})
	}
	a = doubledouble.DoubleFloat64{5000000000000.0400}
	b = doubledouble.DoubleFloat64{10000000000000000.0}
	sum := doubledouble.Add(a, b)
	if sum[0] < b[0] {
		t.Error("got lt b; expected gt b")
	}
}

func BenchmarkDiv(b *testing.B) {
	x, y := doubledouble.DoubleFloat64{1000000000.9}, doubledouble.DoubleFloat64{2.0}
	for i := 0; i < b.N; i++ {
		x = doubledouble.Div(x, y)
	}
}

func TestDiv(t *testing.T) {
	var a, b doubledouble.DoubleFloat64
	b[0] = 2.0
	a = doubledouble.Div(a, b)
	if a != (doubledouble.DoubleFloat64{}) {
		t.Errorf("got %v; want %v", a, doubledouble.DoubleFloat64{})
	}
	a = doubledouble.DoubleFloat64{50000000000000000.0}
	b = doubledouble.DoubleFloat64{2.0}
	quo := doubledouble.Div(a, b)
	if quo[0] >= a[0] {
		t.Error("got ge a; expected lt a")
	}
}

func BenchmarkMul(b *testing.B) {
	x, y := doubledouble.DoubleFloat64{1.9}, doubledouble.DoubleFloat64{2.0}
	for i := 0; i < b.N; i++ {
		x = doubledouble.Mul(x, y)
	}
}

func TestMul(t *testing.T) {
	var a, b doubledouble.DoubleFloat64
	a = doubledouble.Mul(a, b)
	if a != (doubledouble.DoubleFloat64{}) {
		t.Errorf("got %v; want %v", a, doubledouble.DoubleFloat64{})
	}
	a = doubledouble.DoubleFloat64{50000000000000000.0}
	b = doubledouble.DoubleFloat64{2.0}
	prod := doubledouble.Mul(a, b)
	if prod[0] <= a[0] {
		t.Error("got le a; expected gt a")
	}
}

func BenchmarkSub(b *testing.B) {
	x, y := doubledouble.DoubleFloat64{1.1}, doubledouble.DoubleFloat64{2.001}
	for i := 0; i < b.N; i++ {
		x = doubledouble.Sub(x, y)
	}
}

func TestSub(t *testing.T) {
	var a, b doubledouble.DoubleFloat64
	a = doubledouble.Sub(a, b)
	if a != (doubledouble.DoubleFloat64{}) {
		t.Errorf("got %v; want %v", a, doubledouble.DoubleFloat64{})
	}
	a = doubledouble.DoubleFloat64{5000000000000.0400}
	b = doubledouble.DoubleFloat64{10000000000000000.0}
	diff := doubledouble.Sub(a, b)
	if diff[0] >= 0.0 {
		t.Error("got ge z; expected lt z")
	}
}

func TestHighPrecision(t *testing.T) {
	small := doubledouble.DoubleFloat64{0.001}
	big := doubledouble.DoubleFloat64{80000000000000000000000000.0}
	sum := doubledouble.Add(big, small)
	if sum.IsLE(big) {
		t.Error("not high precision sum")
	}
	diff := doubledouble.Sub(big, small)
	if diff.IsGE(big) {
		t.Error("not high precision diff")
	}
	diff2 := doubledouble.Sub(big, doubledouble.Mul(small, doubledouble.DoubleFloat64{2.0}))
	if diff2.IsGE(diff) {
		t.Error("not high precision diff2")
	}
}
