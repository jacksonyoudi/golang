package test

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct {
		a int
		b int
		c int
		d bool
	}{
		{3, 4, 5, true},
		{5, 12, 13, true},
		{8, 15, 17, false},
		{30000, 40000, 50000, true},
	}

	for _, tt := range tests {
		if actual := Triangle(tt.a, tt.b, tt.c); actual != tt.d {
			t.Errorf("err: %d, %d, %d", tt.a, tt.b, tt.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	m := struct {
		a int
		b int
		c int
		d bool
	}{30000, 40000, 50000, true}

	for i := 0; i < b.N; i++ {
		actul := Triangle(m.a, m.b, m.c)
		if actul != m.d {
			b.Errorf("error")
		}
	}

}

func ExampleTriangle() {

	//
}