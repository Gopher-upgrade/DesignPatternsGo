package Prototype

import (
	"testing"
	"strconv"
	"fmt"
)

/**
 *  Factory 工厂模式：
 *
 * @author   Pu ShaoWei <marco0727@gamil.com>
 * @date     2017/9/14
 * @license  MIT
 * -------------------------------------------------------------
 * 通过创建一个原型对象，然后复制原型对象来避免通过标准的方式创建大量的对象产生的开销
 * -------------------------------------------------------------
 */
func Test_Prototype(t *testing.T) {
	fooPrototype := new(FooBookPrototype)
	barPrototype := new(BarBookPrototype)
	for i := 0; i < 5; i++ {
		book := fooPrototype.Clone()
		book.SetTitle("Hello " + strconv.Itoa(i))
		fmt.Printf("%p |%v \n",book,book.title)
	}
	fmt.Printf("-----------------------\n")
	for i := 0; i < 5; i++ {
		book := barPrototype.Clone()
		book.SetTitle("Hello " + strconv.Itoa(i))
		fmt.Printf("%p | %v \n",book,book.title)
	}
}
