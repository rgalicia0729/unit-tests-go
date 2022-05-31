package calcular

import "testing"

func TestSumar(t *testing.T) {
	testCases := []struct {
		a int
		b int
		n int
	}{
		{2, 3, 5},
		{5, 5, 10},
		{75, 25, 100},
	}

	for _, item := range testCases {
		total := Sumar(item.a, item.b)

		if total != item.n {
			t.Errorf("La suma de %d y %d debe de ser %d", item.a, item.b, item.n)
		}
	}
}

func TestGetOlder(t *testing.T) {
	testCases := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{100, 75, 100},
		{25, 50, 50},
	}

	for _, item := range testCases {
		older := GetOlder(item.a, item.b)

		if older != item.n {
			t.Errorf("El mayor entre %d y %d debe de ser: %d", item.a, item.b, item.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{45, 1134903170},
	}

	for _, item := range testCases {
		fib := Fibonacci(item.n)

		if fib != item.r {
			t.Errorf("El valor fibonacci de %d debe de ser %d", item.n, item.r)
		}
	}
}
