package passwordphilosophy

import (
	"testing"
)

func TestPasswordsValidNumExample(t *testing.T) {
	list := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	result := passwordsValidNum(list)
	expected := 0

	if result != expected {
		t.Errorf("When getting the number of valid passwords, got: %v, want: %v.", result, expected)
	}
}
