package trimmedmean

import (
	"math"
	"testing"
)

func approxEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestTrimmedMeanSymmetric(t *testing.T) {
	data := []float64{1, 2, 3, 4, 100}
	// 20% trimmed symmetric => floor(5*0.2)=1 drop each side => remaining {2,3,4} mean=3
	got, err := TrimmedMean(data, 0.2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 3.0
	if !approxEqual(got, want, 1e-9) {
		t.Fatalf("expected %v got %v", want, got)
	}
}

func TestTrimmedMeanAsymmetric(t *testing.T) {
	data := []float64{1, 2, 3, 4, 100}
	// lowTrim=0.0, highTrim=0.2 => floor(5*0)=0 low drop, floor(5*0.2)=1 high drop => remaining {1,2,3,4} mean=2.5
	got, err := TrimmedMean(data, 0.0, 0.2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 2.5
	if !approxEqual(got, want, 1e-9) {
		t.Fatalf("expected %v got %v", want, got)
	}
}

func TestTrimmedMeanInts(t *testing.T) {
	data := []int{10, 20, 30, 40, 1000}
	got, err := TrimmedMeanInts(data, 0.2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// same logic as first test
	want := 30.0
	if !approxEqual(got, want, 1e-9) {
		t.Fatalf("expected %v got %v", want, got)
	}
}

func TestEvaluateDistribution(t *testing.T) {
	// right skewed sample
	data := []float64{1, 1, 2, 2, 3, 4, 10, 100, 200, 300}
	rec, err := EvaluateDistribution(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rec.Skewness <= 0 {
		t.Fatalf("expected positive skewness, got %v", rec.Skewness)
	}
	// Expect high trim recommended > low trim
	if rec.HighTrim <= rec.LowTrim {
		t.Fatalf("expected high trim > low trim for right-skewed data. got low=%v high=%v", rec.LowTrim, rec.HighTrim)
	}
}
