package Builder

import (
	"testing"
	"fmt"
)

func TestCanBuildTruck(t *testing.T) {
	newVehicle := Director{&TruckBuilder{}}
	newVehicle.build("truck")
}

func TestCanBuildCar(t *testing.T) {
	newVehicle := Director{&CarBuilder{}}
	newVehicle.build("Car")
}

type data map[string]string

type TVehicle struct {
	data map[string]string
}

func (t *TVehicle) Set(key, value string) {
	if t.data == nil{
		t.data = map[string]string{key:value}
	}
	t.data[key] = value
}

func TestMap(t *testing.T) {
	ts := TVehicle{}
	ts.Set("a", "b")
	ts.Set("c", "b")
	fmt.Println(ts.data)
}
