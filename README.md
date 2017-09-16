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
<div>
<details>
    <summary> Chain Of Responsibility 职责链模式</summary>
    <li class="task-list-item">
        使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它为止
    </li>
- 使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它为止
- [ ] [AbstractFactory](/Gopher-upgrade/DesignPatternsGo/)

</details>
<details>
    <summary> Command 命令模式</summary>
    <li class="task-list-item">
        将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化；对请求排队或者记录请求日志，以及支持可撤销的操作 Interpreter 解释器模式：给定一个语言，定义它的文法的一种表示，并定义一个解释器，这个解释器使用该表示来解释语言中的句子
    </li>
</details>
<details>
    <summary> Iterator 迭代器模式</summary>
    <li class="task-list-item">
        提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露该对象的内部表示
    </li>
</details>
<details>
    <summary> Mediator 中介者模式</summary>
    <li class="task-list-item">
        用一个中介对象来封装一系列的对象交互。中介这使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互。
    </li>
</details>
<details>
    <summary> Memento 备忘录模式</summary>
    <li class="task-list-item">
        在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可以将该对象恢复到原先保存的状态
    </li>
</details>
<details>
    <summary> Observer 观察者模式</summary>
    <li class="task-list-item">
        定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象。这个主题对象在状态发生改变时，会通知所有观察者对象，使它们能够自动更新自己。
    </li>
</details>
<details>
    <summary> State 状态模式</summary>
    <li class="task-list-item">
        当一个对象的内在状态改变时，允许改变其行为，这个对象看起来像是改变了其类
    </li>
</details>
<details>
    <summary> Strategy 策略模式</summary>
    <li class="task-list-item">
        它定义了算法家族，分别封装起来，让它们可以相互替换，此模式让算法的变化，不会影响到使用算法的客户。
    </li>
</details>
<details>
    <summary> Template Methed模板方法</summary>
    <li class="task-list-item">
        定义一个操作中的算法的骨架，而将一些具体步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
    </li>
</details>
<details>
    <summary> Visitor 访问者模式</summary>
    <li class="task-list-item">
        表示一个作用于某对象结构中的各元素的操作，它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作
    </li>
<details>
</div>

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






## 设计原则

<details>
 <summary>开闭原则（Open Close Principle）</summary>
     <pre><code>
            开闭原则就是说对扩展开放，对修改关闭。在程序需要进行拓展的时候，不能去修改原有的代码，实现一个热插拔的效果。 所以一句话概括就是：为了使程序的扩展性好，易于维护和升级
     </code>
     </pre>
</details>
<details>
 <summary>里氏代换原则（Liskov Substitution Principle）</summary>
     <pre><code>
            里氏代换原则(Liskov Substitution Principle LSP)面向对象设计的基本原则之一。 里氏代换原则中说，任何 基类可以出现的地方，子类一定可以出现。 LSP是继承复用的基石，只有当衍生类可以替换掉基类，软件单位的功能不受 到影响时，基类才能真正被复用，而衍生类也能够在基类的基础上增加新的行为。里氏代换原则是对“开-闭”原则的补充。 实现“开-闭”原则的关键步骤就是抽象化。而基类与子类的继承关系就是抽象化的具体实现，所以里氏代换原则是对实现抽 象化的具体步骤的规范
     </code>
     </pre>
</details>
<details>
 <summary>依赖倒转原则（Dependence Inversion Principle）</summary>
     <pre><code>
          这个是开闭原则的基础，具体内容：真对接口编程，依赖于抽象而不依赖于具体。
     </code>
     </pre>
</details>
<details>
 <summary>接口隔离原则（Interface Segregation Principle）</summary>
     <pre><code>
        这个原则的意思是：使用多个隔离的接口，比使用单个接口要好。还是一个降低类之间的耦合度的意思，从这儿我们看出， 其实设计模式就是一个软件的设计思想，从大型软件架构出发，为了升级和维护方便。所以上文中多次出现：降低依赖，降低耦合。
     </code>
     </pre>
</details>
<details>
 <summary>迪米特法则（最少知道原则）（Demeter Principle）</summary>
     <pre>
     <code>
          为什么叫最少知道原则，就是说：一个实体应当尽量少的与其他实体之间发生相互作用，使得系统功能模块相对独立。
     </code>
     </pre>
</details>
<details>
 <summary>合成复用原则（Composite Reuse Principle）</summary>
     <pre>
     <code>
          原则是尽量使用合成/聚合的方式，而不是使用继承。
     </code>
     </pre>
</details>

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
