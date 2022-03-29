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

func TestProductGetTooManyRequested(t *testing.T) {
	p := productList {}
	p.add(7)
	p.add(0)
	p.add(2)
	p.add(5)
	p.add(4)
	result, err := p.get(10)
	expected := 0

	if result != expected {
		t.Errorf("When getting too many products, incorrect result given, got: %v, expected: %v", result, expected)
	}

	if err == nil {
		t.Errorf("When getting too many products, got an error thrown: %v", err)
	}
}

func TestProductGetNoProducts(t *testing.T) {
	p := productList {}
	p.add(7)
	p.add(0)
	p.add(2)
	p.add(5)
	p.add(4)
	result, err := p.get(0)
	expected := 0

	if result != expected {
		t.Errorf("When getting no products, incorrect result given, got: %v, expected: %v", result, expected)
	}

	if err != nil {
		t.Errorf("When getting no products, got an error thrown: %v", err)
	}
}

func BenchmarkProductListAddGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := productList {}
		p.add(7)
		p.add(0)
		p.add(2)
		p.add(5)
		p.add(4)
		p.get(3)
	}
}
