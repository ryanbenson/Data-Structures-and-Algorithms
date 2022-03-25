package productlist

import "testing"

func TestProductGet(t *testing.T) {
	p := productList {}
	p.add(7)
	p.add(0)
	p.add(2)
	p.add(5)
	p.add(4)
	result, err := p.get(3)
	expected := 40

	if result != expected {
		t.Errorf("When finding the product total value, incorrect result given, got: %v, expected: %v", result, expected)
	}

	if err != nil {
		t.Errorf("When finding the product total value, got an error thrown: %v", err)
	}
}

// func BenchmarkSmallestTimeIntervalSimple(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		times := []string{"01:00", "08:15", "11:30", "13:45", "14:10", "20:05"}
// 		smallestTimeIntervalSimple(times)
// 	}
// }
