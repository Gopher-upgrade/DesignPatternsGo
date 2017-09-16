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

## 行为型模式实例 Behavioral

- [ ] [ChainOfResponsibilities](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)
- [ ] [Command](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Command/Command.go)
- [ ] [Iterator](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Iterator/Iterator.go)
- [ ] [Mediator](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)
- [ ] [Memento](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)
- [ ] [NullObject](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)
- [ ] [Observer](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)
- [ ] [Specification](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)
- [ ] [State](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)
- [ ] [Strategy](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)
- [ ] [TemplateMethod](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)
- [ ] [Visitor](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Behavioral/Behavioral_test.go)

## 创建型模式实例 Creational

- [ ] [AbstractFactory](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Builder](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [FactoryMethod](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Multiton](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Pool](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Prototype](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [SimpleFactory](/Gopher-upgrade/DesignPatternsGo/)
- [x] [Singleton](/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Creational/Singleton/doubleLock/main.go)
- [ ] [StaticFactory](/Gopher-upgrade/DesignPatternsGo/)


##  结构型模式实例 tructural

- [ ] [Adapter](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Bridge](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Composite](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [DataMapper](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Decorator](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [DependencyInjection](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Facade](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [FluentInterface](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Flyweight](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Proxy](/Gopher-upgrade/DesignPatternsGo/)
- [ ] [Registry](/Gopher-upgrade/DesignPatternsGo/)



### 开闭原则（Open Close Principle）
```
开闭原则就是说对扩展开放，对修改关闭。在程序需要进行拓展的时候，不能去修改原有的代码，实现一个热插拔的效果。 所以一句话概括就是：为了使程序的扩展性好，易于维护和升级
```

### 里氏代换原则（Liskov Substitution Principle）
```
里氏代换原则(Liskov Substitution Principle LSP)面向对象设计的基本原则之一。 里氏代换原则中说，任何 基类可以出现的地方，子类一定可以出现。 LSP是继承复用的基石，只有当衍生类可以替换掉基类，软件单位的功能不受 到影响时，基类才能真正被复用，而衍生类也能够在基类的基础上增加新的行为。里氏代换原则是对“开-闭”原则的补充。 实现“开-闭”原则的关键步骤就是抽象化。而基类与子类的继承关系就是抽象化的具体实现，所以里氏代换原则是对实现抽 象化的具体步骤的规范
```
### 依赖倒转原则（Dependence Inversion Principle）
```
这个是开闭原则的基础，具体内容：真对接口编程，依赖于抽象而不依赖于具体。
```
### 接口隔离原则（Interface Segregation Principle）
```
这个原则的意思是：使用多个隔离的接口，比使用单个接口要好。还是一个降低类之间的耦合度的意思，从这儿我们看出， 其实设计模式就是一个软件的设计思想，从大型软件架构出发，为了升级和维护方便。所以上文中多次出现：降低依赖，降低耦合。
```
### 迪米特法则（最少知道原则）（Demeter Principle）
```
为什么叫最少知道原则，就是说：一个实体应当尽量少的与其他实体之间发生相互作用，使得系统功能模块相对独立。
```
### 合成复用原则（Composite Reuse Principle）
```
原则是尽量使用合成/聚合的方式，而不是使用继承。
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
            <td ><a href="https://github.com/PuShaoWei"><img src="https://avatars2.githubusercontent.com/u/18391791?v=1" /></a>
            <p align="center">Marco</p>
            </td>
        </tr>
    </tbody>
</table>

## License

MIT
