package FactoryMethod

import (
	"testing"
)
/**
 *  Factory 工厂模式：
 *
 * @author   Pu ShaoWei <marco0727@gamil.com>
 * @date     2017/9/14
 * @license  MIT
 * -------------------------------------------------------------
 * 比简单工厂模式好的一点是工厂方法可以通过继承实现不同的创建对象的逻辑。
 * 举个简单的例子，这些抽象类都仅仅是一个接口
 * 这个模式是一个 “真正” 的设计模式，
 * 因为它遵循了依赖反转原则（Dependency Inversion Principle） 众所周知这个 “D” 代表了真正的面向对象程序设计。
 * 它意味着工厂方法类依赖于类的抽象，而不是具体将被创建的类，这是工厂方法模式与简单工厂模式和静态工厂模式最重要的区别
 * -------------------------------------------------------------
 */
func TestFactory(t *testing.T) {
	i := new(ItalianFactory)
	i.Create(CHEAP).SetColor("test")
}
