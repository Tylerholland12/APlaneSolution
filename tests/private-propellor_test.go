package tests

import "testing"

func TestScrapePropellorPlane(t *testing.T) {
	tables := []struct {
		url      string 
		expected bool
	}{
		{"https://www.aircraftcompare.com/aircraft-categories/private-propellor-planes/", true}
		{"https://www.aircraftcompare.com/aircraft-categories/private-jets/", false}
		{"https://www.aircraftcompare.com/aircraft-categories/commercial-airplanes/", false}
	}
}
