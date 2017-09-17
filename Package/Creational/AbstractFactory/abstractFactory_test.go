package AbstractFactory

import (
	"testing"
)

/**
 *  Abstract Factory 抽象工厂模式：
 *
 * @author   Pu ShaoWei <marco0727@gamil.com>
 * @date     2017/9/14
 * @license  MIT
 * -------------------------------------------------------------
 * 创建一系列互相关联或依赖的对象时不需要指定将要创建的对象对应的类，
 * 因为这些将被创建的对象对应的类都实现了同一个接口。
 * 抽象工厂的使用者不需要关心对象的创建过程，它只需要知道这些对象是如何协调工作的
 * -------------------------------------------------------------
 */

func TestAbstractFactory(t *testing.T) {
	var f CreatorFactory
		f = new(JsonFactory)
		j := f.CreateText()
		jsonRender := "我是一个 Json \n"
		j.Production(&jsonRender)

		f = new(HtmlFactory)
		h := f.CreateText()
		htmlRender := "我是一个 Html \n"
	    h.Production(&htmlRender)
}
