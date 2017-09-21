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

## 设计原则

<div>
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
</div>

## 行为型模式实例 Behavioral
<div>
  <details>
      <summary> Chain Of Responsibility 职责链模式</summary>
      <code>
          使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它为止
      </code>
      <p>
        <a href="/Gopher-upgrade/DesignPatternsGo/blob/master/Gopher-upgrade/DesignPatternsGo">StaticFactory</a>
      </p>
  </details>
  <details>
      <summary> Command 命令模式</summary>
      <code>
          将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化；对请求排队或者记录请求日志，以及支持可撤销的操作 Interpreter 解释器模式：给定一个语言，定义它的文法的一种表示，并定义一个解释器，这个解释器使用该表示来解释语言中的句子
      </code>
  </details>
  <details>
      <summary> Iterator 迭代器模式</summary>
      <code>
          提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露该对象的内部表示
      </code>
  </details>
  <details>
      <summary> Mediator 中介者模式</summary>
      <code>
          用一个中介对象来封装一系列的对象交互。中介这使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互。
      </code>
  </details>
  <details>
      <summary> Memento 备忘录模式</summary>
      <code>
          在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可以将该对象恢复到原先保存的状态
      </code>
  </details>
  <details>
      <summary> Observer 观察者模式</summary>
      <code>
          定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象。这个主题对象在状态发生改变时，会通知所有观察者对象，使它们能够自动更新自己。
      </code>
  </details>
  <details>
      <summary> State 状态模式</summary>
      <code>
          当一个对象的内在状态改变时，允许改变其行为，这个对象看起来像是改变了其类
      </code>
  </details>
  <details>
      <summary> Strategy 策略模式</summary>
      <code>
          它定义了算法家族，分别封装起来，让它们可以相互替换，此模式让算法的变化，不会影响到使用算法的客户。
      </code>
  </details>
  <details>
      <summary> Template Methed模板方法</summary>
      <code>
          定义一个操作中的算法的骨架，而将一些具体步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
      </code>
  </details>
  <details>
      <summary> Visitor 访问者模式</summary>
      <code>
          表示一个作用于某对象结构中的各元素的操作，它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作
      </code>
  <details>
</div>

## 创建型模式实例 Creational

<div>
    <details>
        <summary> Abstract Factory抽象工厂模式</summary>
        <code> 提供一个创建一系列相关或者相互依赖对象的接口，而无需指定他们具体的类</code>
        <p> <a href="https://github.com/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Creational/AbstractFactory/abstractFactory.go">Abstract Factory</a></p>
    </details>
    <details>
        <summary> Builder生成器模式</summary>
        <code> 将一个复杂对象的构建与它表示分离，使得同样的构建过程可以创建不同的表示</code>
        <p>  <a href="https://github.com/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Creational/Builder/Builder_test.go">Builder</a></p>
    </details>
    <details>
        <summary> Factory Method工厂方法模式</summary>
        <code> 定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类</code>
        <p>  <a href="https://github.com/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Creational/FactoryMethod/Factory_test.go">Factory Method</a></p>
    </details>
    <details>
        <summary> Prototype 原型模式</summary>
        <code> 用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象（clone浅复制、深复制）</code>
        <p>  <a href="https://github.com/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Creational/Prototype/prototype_test.go">Factory Method</a></p>
    </details>
    <details>
        <summary> Pool 对象池模式</summary>
        <code>
对象池设计模式 是创建型设计模式，它会对新创建的对象应用一系列的初始化操作，
让对象保持立即可使用的状态 - 一个存放对象的 “池子” - 而不是对对象进行一次性的的使用(创建并使用，完成之后立即销毁)。
 对象池的使用者会对对象池发起请求，以期望获取一个对象，并使用获取到的对象进行一系列操作，
 当使用者对对象的使用完成之后，使用者会将由对象池的对象创建工厂创建的对象返回给对象池，而不是用完之后销毁获取到的对象。
 对象池在某些情况下会带来重要的性能提升，比如耗费资源的对象初始化操作，实例化类的代价很高，但每次实例化的数量较少的情况下。
 对象池中将被创建的对象会在真正被使用时被提前创建，
  避免在使用时让使用者浪费对象创建所需的大量时间(比如在对象某些操作需要访问网络资源的情况下)从池子中取得对象的时间是可预测的，但新建一个实例所需的时间是不确定。
 总之，对象池会为你节省宝贵的程序执行时间，
 比如像数据库连接，socket连接，大量耗费资源的代表数字资源的对象，像字体或者位图。
 不过，在特定情况下，简单的对象创建池(没有请求外部的资源，仅仅将自身保存在内存中)或许并不会提升效率和性能，这时候，就需要使用者酌情考虑了。
 </code>
        <p>  <a href="https://github.com/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Creational/Prototype/prototype_test.go">Factory Method</a></p>
    </details>
    <details>
        <summary> Singleton单例模式</summary>
        <code> 保证一个类仅有一个实例，并提供一个访问它的全局访问点。</code>
        <p>  <a href="https://github.com/Gopher-upgrade/DesignPatternsGo/blob/master/Package/Creational/doubleLock/main.go">Factory Method</a></p>
    </details>
</div>

##  结构型模式实例 tructural

<div>
    <details>
        <summary>Adapter适配器模式</summary>
        <code>将一个类的接口转换成客户端希望的另一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作</code>
        <p><a href="/">正在梳理..</a></p>
    </details>
    <details>
        <summary>Bridge桥接模式</summary>
        <code>将抽象化(Abstraction)与实现化(Implementation)脱耦，使得二者可以独立地变化；</code>
        <p><a href="/">正在梳理..</a></p>
    </details>
    <details>
        <summary>Composite合成/组合模式</summary>
        <code>将对象组合成树形结构，以表示“部分-整体”的层次结构。组合模式使得用户对单个对象和组合对象的使用具有一致性</code>
        <p><a href="/">正在梳理..</a></p>
    </details>
    <details>
        <summary>Decorator装饰模式</summary>
        <code>动态地给一个对象添加一些额外的职责，就增加功能来说，装饰模式比生成子类更为灵活。</code>
        <p><a href="/">正在梳理..</a></p>
    </details>
    <details>
        <summary>Facade 外观模式</summary>
        <code>为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这一子系统更加容易使用</code>
        <p><a href="/">正在梳理..</a></p>
    </details>
    <details>
        <summary>Flyweight 享元模式</summary>
        <code>运用共享技术有效地支持大量细粒度的对象</code>
        <p><a href="/">正在梳理..</a></p>
    </details>
    <details>
        <summary>Proxy 代理模式</summary>
        <code>为其他对象提供一种代理，以控制对这个对象的访问。</code>
        <p><a href="/">正在梳理..</a></p>
    </details>
</div>



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
