package main

/**
 * @example  非线程安全的单例模式 （懒汉）
 * @author   ShaoWei Pu <pushaowei0727@gmail.com>
 * @date     2017/9/6
 * @license  MIT
 * -------------------------------------------------------------
 *
 * 这就是一个最简单的单例了，对于singleton结构体，我们需要通过GetInstance取获取它的实例
 * 这个函数中首先判断instance这个全局变量是否存在，若不存在则将singleton结构体的引用地址传递给它
 * 若存在则返回实例体，固在并发情况下是不可靠的
 *
 * @return instance
 */

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func main() {
	GetInstance()
}
