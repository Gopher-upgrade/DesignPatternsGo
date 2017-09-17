package FactoryMethod

import "fmt"

const (
	CHEAP = "Cheap"
	FAST  = "Fast"
)


type ItalianFactory struct {
}

func (this *ItalianFactory) Create(action string) VehicleInterface{
	obj := this.CreateVehicle(action)
	return obj
}

func (this *ItalianFactory) CreateVehicle(action string) VehicleInterface {
	switch action {
	case CHEAP:
		return new(Bicycle)
	case FAST:
		return new(CarFerrari)
	default:
		return nil
	}
}

type VehicleInterface interface {
	SetColor(action string)
}
type Bicycle struct{}
func (b *Bicycle) SetColor(action string) {
	fmt.Println(action)
}

type CarFerrari struct{}
func (b *CarFerrari) SetColor(action string) {
}
func (b *CarFerrari) AddAMGTuning(){

}
