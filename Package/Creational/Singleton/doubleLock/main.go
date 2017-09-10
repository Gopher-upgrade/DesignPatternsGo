package main

import "sync"

/**
 * @example  do 单例
 * @author   ShaoWei Pu <pushaowei0727@gmail.com>
 * @date     2017/9/6
 * @license  MIT
 * -------------------------------------------------------------
 * 	其实sync.Once，有一个Do方法，在它中的函数go会只保证仅仅调用一次...
 *  呃，有话好好说。。别打脸
 *
 * 不能任性的用，这样的单例会一直存在内存中，一些我们用的不是那么频繁的东西使用了单例就造成了内存的浪费，
 * 大家在用单例的时候还是要多思考思考，这个模块适不适合用单例！
 *
 * @return instance
 */

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func(){
		instance = &singleton{}
	})
	return instance
}

func main() {
	GetInstance()
}
