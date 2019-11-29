package faker

import (
	"testing"

	handler "github.com/lakshay2395/kepler-demo/handlers"
)

func TestIsLatLongPresentInServiceArea(t *testing.T) {
	serviceArea := handler.ServiceArea{}
	expectedValue := false
	value := IsLatLongPresentInServiceArea(serviceArea)
	if value != expectedValue {
		t.Errorf("Expected Value %v , returned %v", expectedValue, value)
	}
}
