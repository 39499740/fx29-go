package fx29

import (
	"testing"
)

func TestEncode(t *testing.T) {
	encoded := Encode([]byte("你好😡"), nil)
	expected := "5L2g5aW98J%2BYoQ"
	if encoded != expected {
		t.Errorf("Expected %v, got %v", expected, encoded)
	}

	encoded = Encode([]byte("你好😡"), []byte{17, 2, 255})
	expected = "9b9f9KdC4Z1nsA"
	if encoded != expected {
		t.Errorf("Expected %v, got %v", expected, encoded)
	}
}

func TestDecode(t *testing.T) {
	decoded, err := Decode("5L2g5aW98J%2BYoQ", nil)
	pie(err)
	expected := "你好😡"
	if string(decoded) != expected {
		t.Errorf("Expected %v, got %v", expected, decoded)
	}

	decoded, err = Decode("9b9f9KdC4Z1nsA", []byte{17, 2, 255})
	pie(err)
	expected = "你好😡"
	if string(decoded) != expected {
		t.Errorf("Expected %v, got %v", expected, decoded)
	}
}

// panic if error occurred
func pie(err error) {
	if err != nil {
		panic(err)
	}
}
