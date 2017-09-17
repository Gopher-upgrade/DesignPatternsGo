package Builder

type BuilderInterface interface {
	CreateVehicle(str string)
	AddDoors(str string)
	AddEngine(str string)
	AddWheel(str string)
	GetVehicle(str string)
}

type Director struct {
	create BuilderInterface
}

func (d *Director) build(str string) {
	d.create.CreateVehicle(str)
	d.create.AddDoors(str)
	d.create.AddEngine(str)
	d.create.AddWheel(str)
	d.create.GetVehicle(str)
}
