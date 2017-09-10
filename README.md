<h1 align="center">:rocket: golang 实战23种设计模式 </h1>

<p align="center">
<a href="https://github.com/PuShaoWei/designPatterns-go#简易结构">
  <img src="https://img.shields.io/badge/php-done-brightgreen.svg" alt="php">
</a>
<a href="https://github.com/PuShaoWei/designPatterns-go">
    <img src="https://img.shields.io/github/issues-pr-raw/designPatterns-go/cdnjs.svg">
</a>
<a href="https://github.com/PuShaoWei/designPatterns-go">
    <img src="https://img.shields.io/codacy/grade/e27821fb6289410b8f58338c7e0bc686.svg">
</a>
<a href="https://github.com/PuShaoWei/designPatterns-go">
    <img src="https://img.shields.io/travis/rust-lang/rust.svg">
</a>
<a href="https://github.com/PuShaoWei/designPatterns-go">
    <img src="https://img.shields.io/github/license/mashape/apistatus.svg">
</a>
</p>
<p align="center"> <a href="./README-EN.md">English</a>　<p>

## 关于

>  知行合一，学以致用

这里记录在Golang 学习中遇到的一些有趣的设计模式，代码段来自一些开源库源码，以及个人与参与者实现的一些较好的案例，我认为设计模式的最大问题是大家知道这是什么而往往总是不知道何时能用上

## 简易结构
```
├── Package
│   ├── Behavioral
│   │   ├──  [ ] ChainOfResponsibilities
│   │   ├──  [ ] Command
│   │   ├──  [ ] Iterator
│   │   ├──  [ ] Mediator
│   │   ├──  [ ] Memento
│   │   ├──  [ ] NullObject
│   │   ├──  [ ] Observer
│   │   ├──  [ ] Specification
│   │   ├──  [ ] State
│   │   ├──  [ ] Strategy
│   │   ├──  [ ] TemplateMethod
│   │   └──  [ ] Visitor
│   ├── Creational
│   │   ├──  [ ] AbstractFactory
│   │   ├──  [ ] Builder
│   │   ├──  [ ] FactoryMethod
│   │   ├──  [ ] Multiton
│   │   ├──  [ ] Pool
│   │   ├──  [ ] Prototype
│   │   ├──  [ ] SimpleFactory
│   │   ├──  [√] Singleton
│   │   │   ├── doubleLock
│   │   │   │   └── main.go
│   │   │   ├── hungry
│   │   │   │   └── main.go
│   │   │   └── lazy
│   │   │       └── main.go
│   │   └── StaticFactory
│   └── Structural
│       ├──  [ ] Adapter
│       ├──  [ ] Bridge
│       ├──  [ ] Composite
│       ├──  [ ] DataMapper
│       ├──  [ ] Decorator
│       ├──  [ ] DependencyInjection
│       ├──  [ ] Facade
│       ├──  [ ] FluentInterface
│       ├──  [ ] Flyweight
│       ├──  [ ] Proxy
│       └──  [ ] Registry
│     
├── LICENSE
├── README-CN.md
└── README.md
```

## 感谢

- [DesignPatternsPHP](https://github.com/domnikl/DesignPatternsPHP)

## 纠错

如果你发现有问题，你可以发起 [issue](https://github.com/PuShaoWei/designPatterns-go/issues) 或者 [pull request](https://github.com/PuShaoWei/designPatterns-go/pulls),我会及时纠正

> 补充:发起pull request的commit message请参考文章[Commit message 和 Change log 编写指南](http://www.ruanyifeng.com/blog/2016/01/commit_message_change_log.html)

## 贡献者
<table>
    <tbody>
        <tr>
            <td ><a href="https://github.com/luin"><img src="https://avatars2.githubusercontent.com/u/18391791?v=1" /></a>
            <p align="center">Marco</p>
            </td>
        </tr>
    </tbody>
</table>

## License

MIT
