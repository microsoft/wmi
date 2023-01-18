package errors

import (
	"math"
	"math/rand"
	"testing"

	perrors "github.com/pkg/errors"
)

func init() {
}

func randUint16() uint16 {
	return uint16(rand.Intn(math.MaxUint16 + 1))
}

func TestWmiError(t *testing.T) {
	err := NewWMIError(0) // Boundary check - Use the min uint value. - should return true
	validateWmiError(t, err)

	err = NewWMIError(math.MaxUint16) // Boundary check - Use the max uint value. - should return true
	validateWmiError(t, err)

	err = NewWMIError(randUint16()) // Use a random number as an error. - should return true
	validateWmiError(t, err)

	result := IsWMIError(nil) // Test with a nil error - should return false
	if result != false {
		t.Fatal("Failed IsWMIError(nil) returned true")
	}

	result = IsWMIError(Failed) // Test with another error type - should return false
	if result != false {
		t.Fatal("Failed IsWMIError(Failed) returned true")
	}
}

func validateWmiError(t *testing.T, err error) {
	result := IsWMIError(err)
	if result != true {
		t.Fatal("Failed IsWMIError(" + err.Error() + ") returned false")
	}

	err = perrors.Wrapf(err, "An error happened. ")
	result = IsWMIError(err)
	if result != true {
		t.Fatal("Failed Wrapf(IsWMIError(" + err.Error() + ")) returned false")
	}

	err = perrors.Wrapf(err, "An error happened. ")
	result = IsWMIError(err)
	if result != true {
		t.Fatal("Failed Wrapf(Wrapf(IsWMIError(" + err.Error() + "))) returned false")
	}
}
