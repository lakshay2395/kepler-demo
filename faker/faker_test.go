package faker

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	client "github.com/influxdata/influxdb1-client"
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

func TestGenerateClientPoint(t *testing.T) {
	now := time.Now()
	expectedValue := client.Point{
		Measurement: "balle",
		Tags: map[string]string{
			"service_area": "A",
			"service_type": "T",
		},
		Fields: map[string]interface{}{
			"lat":   100,
			"long":  -100,
			"value": 0.4,
		},
		Time: now,
	}
	value := GenerateClientPoint("balle", handler.ServiceArea{
		ID:   1,
		Name: "A",
		Lat:  -100,
		Long: 100,
	}, handler.ServiceType{
		ID:   1,
		Name: "T",
	}, 0.4, now)
	if cmp.Equal(value, expectedValue) {
		t.Errorf("Expected Value %v , returned %v", expectedValue, value)
	}
}
