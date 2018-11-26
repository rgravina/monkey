package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	other := &String{Value: "another string"}
	if hello1.HashKey() != hello2.HashKey() {
		t.Fatalf("hash keys do not match")
	}
	if hello1.HashKey() == other.HashKey() {
		t.Fatalf("hash keys should not match")
	}
}
