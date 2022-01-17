package adapter

import "testing"

func TestAdapter(t *testing.T) {
	worker := &worker{}
	cement := &cement{}
	longWood := &longWood{}
	longWoodAdapter := &longWoodAdapter{longWood: longWood}
	worker.inputRoomForMaterial(cement)
	worker.inputRoomForMaterial(longWoodAdapter)
}
