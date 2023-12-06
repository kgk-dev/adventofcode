package ulits

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadLineStream(t *testing.T) {
	input := `9cbncbxclbvkmfzdnldc
jjn1drdffhs
3six7`
	inputReader := strings.NewReader(input)
	expected := []string{
		"9cbncbxclbvkmfzdnldc",
		"jjn1drdffhs",
		"3six7",
	}
	result := make([]string, 0)

	for line := range ReadLineStream(inputReader) {
		result = append(result, line)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Must equal want %v and got %v", expected, result)
	}

}
