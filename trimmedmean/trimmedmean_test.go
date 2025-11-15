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

func TestAutoTrimmedMean(t *testing.T) {
	// Test symmetric case (approximately normal distribution)
	symmetricData := []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	mean, rec, err := AutoTrimmedMean(symmetricData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// For symmetric data, trimming should be symmetric
	if !approxEqual(rec.LowTrim, rec.HighTrim, 1e-9) {
		t.Fatalf("expected symmetric trimming for symmetric data, got low=%v high=%v", rec.LowTrim, rec.HighTrim)
	}

	// The mean should be reasonable
	if mean < 12 || mean > 17 {
		t.Fatalf("expected mean between 12 and 17, got %v", mean)
	}

	// Test asymmetric case (right-skewed data)
	rightSkewedData := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}
	mean2, rec2, err := AutoTrimmedMean(rightSkewedData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// For right-skewed data, high trim should be greater than low trim
	if rec2.HighTrim <= rec2.LowTrim {
		t.Fatalf("expected high trim > low trim for right-skewed data, got low=%v high=%v", rec2.LowTrim, rec2.HighTrim)
	}

	// The trimmed mean should be less affected by the outlier (100)
	if mean2 >= 20 { // Arithmetic mean would be ~14.5, but trimmed should be lower due to high outlier
		t.Fatalf("expected trimmed mean to handle outlier, got %v", mean2)
	}
}

func TestAutoTrimmedMeanInts(t *testing.T) {
	// Test with integer data
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100} // right-skewed
	mean, rec, err := AutoTrimmedMeanInts(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should recommend asymmetric trimming for skewed data
	if rec.HighTrim <= rec.LowTrim {
		t.Fatalf("expected high trim > low trim for right-skewed data, got low=%v high=%v", rec.LowTrim, rec.HighTrim)
	}

	// Mean should be reasonable (not overly influenced by outlier)
	if mean >= 20 {
		t.Fatalf("expected trimmed mean to handle outlier, got %v", mean)
	}
}

func TestAsymmetricTrimmingCalculation(t *testing.T) {
	// Specific test for asymmetric trimming calculation
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 100, 200, 300} // Highly right-skewed

	// Manually apply asymmetric trimming based on expected recommendation
	// For highly right-skewed data, we expect something like (0.05, 0.25)
	mean, err := TrimmedMean(data, 0.05, 0.25)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// With 12 elements:
	// lowDrop = floor(12 * 0.05) = 0
	// highDrop = floor(12 * 0.25) = 3
	// Remaining elements: data[0:9] = [1, 2, 3, 4, 5, 6, 7, 8, 9]
	// Mean = (1+2+3+4+5+6+7+8+9)/9 = 5.0
	expected := 5.0

	if !approxEqual(mean, expected, 1e-9) {
		t.Fatalf("expected %v got %v", expected, mean)
	}
}
