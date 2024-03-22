package utils

import (
	"testing"
)

func TestGenerateExpression(t *testing.T) {
	println(GenerateExpression(20, 0.2, 0.1))
}

func TestParseExpression(t *testing.T) {
	if value, e := ParseExpression("13 x 2"); e != nil {
		t.Fatal(e.Error())
	} else {
		if value != 26 {
			t.Fatal("failed to parse expression")
		}
	}
}

func TestRandRange(t *testing.T) {
	println(randRange(1, 4))
}
