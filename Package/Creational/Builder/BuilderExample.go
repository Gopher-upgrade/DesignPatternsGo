package Builder

import "fmt"

type Vehicle struct {
	data map[string]string
}

func (v *Vehicle) SetPart(key, value string) {
	if v.data == nil {
		v.data = map[string]string{key: value}
	}
	v.data[key] = value
}

// ------------------------------------------------
type TruckBuilder struct {
	Vehicle // 思考一下这里与 *Vehicle 的区别 ,另外此处并未实现将对象放进属性的操作，后期一定要改进
}

func (this *TruckBuilder) AddDoors(doors string) {
	this.Vehicle.SetPart("addPors", doors)
}

func (this *TruckBuilder) AddEngine(engine string) {
	this.Vehicle.SetPart("AddEngine", engine)
}

func (this *TruckBuilder) AddWheel(wheel string) {
	this.Vehicle.SetPart("AddWheel", wheel)
}

func (this *TruckBuilder) CreateVehicle(wheel string) {
	this.Vehicle.SetPart("CreateVehicle", wheel)
}

func (this *TruckBuilder) GetVehicle(doors string) {
	fmt.Println(this.data)
}

// ------------------------------------------------
type CarBuilder struct {
	Vehicle
}

func (this CarBuilder) AddDoors(doors string) {
	this.SetPart("addPors", doors)
}

func (this *CarBuilder) AddEngine(engine string) {
	this.SetPart("AddEngine", engine)
}

func (this *CarBuilder) AddWheel(wheel string) {
	this.SetPart("AddWheel", wheel)
}

func (this *CarBuilder) CreateVehicle(wheel string) {
	this.SetPart("CreateVehicle", wheel)
}

func (this *CarBuilder) GetVehicle(doors string) {
	fmt.Println(this.data)
}
