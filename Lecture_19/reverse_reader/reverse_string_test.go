package reverseReader

import (
	"testing"
)

func TestNewReverseStringReader(t *testing.T) {
	//Arrange/Act
	expectedResult := "!dlrow olleH"
	result := NewReverseStringReader("Hello world!")

	//Assert
	if result.Text != expectedResult {
		t.Errorf("Expected: %s, got: %s", expectedResult, result.Text)
	}
}
