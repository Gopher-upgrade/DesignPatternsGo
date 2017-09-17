package Builder

import (
	"testing"
	"fmt"
)
/**
 *  Builder 生成器模式：
 *
 * @author   Pu ShaoWei <marco0727@gamil.com>
 * @date     2017/9/14
 * @license  MIT
 * -------------------------------------------------------------
 * 生成器的目的是将复杂对象的创建过程（流程）进行抽象，生成器表现为接口的形式。
 * 在特定的情况下，比如如果生成器对将要创建的对象有足够多的了解，
 * 那么代表生成器的接口(interface)可以是一个抽象类(也就是说可以有一定的具体实现，就像众所周知的适配器模式)。
 * 如果对象有复杂的继承树，理论上创建对象的生成器也同样具有复杂的继承树。
 * -------------------------------------------------------------
 */

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
