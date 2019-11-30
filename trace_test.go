package locate

import (
	"testing"
)

func TestWhereAmI(t *testing.T) {

	expected := "github.com/invisiblethreat/location.TestWhereAmI"
	location := WhereAmI()
	if location != expected {
		t.Errorf("Expected %s, but got %s", expected, location)
	}
}
