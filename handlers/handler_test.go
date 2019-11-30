package handlers

import "testing"

func TestGenerateSupplyCommands(t *testing.T) {
	expectedValue := "select * from low_supply where service_area='a' and service_type='b'"
	value := generateSupplyCommands("select * from low_supply ", "a", "b")
	if expectedValue != value {
		t.Errorf("Expected %s, returned %s", expectedValue, value)
	}
	expectedValue = "select * from low_supply where service_area='a'"
	value = generateSupplyCommands("select * from low_supply ", "a", "")
	if expectedValue != value {
		t.Errorf("Expected %s, returned %s", expectedValue, value)
	}
	expectedValue = "select * from low_supply where service_type='b'"
	value = generateSupplyCommands("select * from low_supply ", "", "b")
	if expectedValue != value {
		t.Errorf("Expected %s, returned %s", expectedValue, value)
	}
}

func TestGetServiceAreas(t *testing.T) {

}
