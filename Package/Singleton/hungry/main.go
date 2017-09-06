package main

import "sync"

/**
 * @example  带锁线程安全的单例模式 （饿汉）
 * @author   ShaoWei Pu <pushaowei0727@gmail.com>
 * @date     2017/9/6
 * @license  MIT
 * -------------------------------------------------------------
 * sync/mutext是Go语言底层基础对象之一，用于构建多个goroutine间的同步逻辑，因此被大量高层对象所使用。
 * 其工作模型类似于Linux内核的futex对象，具体实现极为简洁，性能也有保证。
 * mutex对象仅有两个数值字段，分为为state（存储状态）和sema（用于计算休眠goroutine数量的信号量）。
 * 初始化时填入的0值将mutex设定在未锁定状态，同时保证时间开销最小。这一特性允许将mutex作为其它对象的子对象使用。
 *
 * 在之前的单例下跟新了一个带线程锁的机制，在getInstance 中，没次调用我们都会上锁，保证只有一个goroutine执行它，
 * 这样就解决并发问题了，不过现在什么情况都会上锁，而且加锁代价昂贵，也不是很优雅的做法
 *
 * @return instance
 */

type singleton struct{}

var instance *singleton
var lock *sync.Mutex = &sync.Mutex{}

func GetInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func main() {
	GetInstance()
}
