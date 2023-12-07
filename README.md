# Hexagonal architecture

Created time: December 6, 2023 3:12 PM

Last edited time: December 7, 2023 9:53 PM


> 使用传统的分层架构，我们的所有依赖项都指向一个方向，上面的每一层都依赖于下面的层。传输层将依赖于交互器，交互器将依赖于持久层。

> 在六边形架构中，所有依赖项都指向内部——我们的核心业务逻辑对传输层或数据源一无所知。尽管如此，传输层知道如何使用交互器，数据源知道如何符合存储库接口。

## 概述

最近在想着写一个个人项目，但是在项目的结构上却犯了难，此时翻到了一个视频，采用Hexagonal architecture（六边形架构），也被称为Ports and Adapters，大致就是下面图片的结构：

![https://asdxz.oss-cn-beijing.aliyuncs.com/img/202312072050938.png](https://asdxz.oss-cn-beijing.aliyuncs.com/img/202312072050938.png)

一共分为三层：

Domain： 里面放的是处理的基本逻辑，可以理解为大纲，它决定着Application和Framework的选择和实现

Application:：它协调使用我们的Domain代码, 通过位于两者之间的方式，调整从framework到domain的请求

Framework： 为外部组件提供交互方式，驱动通常放在左边，被驱动放在右边

我们需要注意的是：

1. 高层模块不应该依赖于低层模块。两者都应该依赖于抽象。
2. 抽象不应该依赖于细节。细节应该依赖于抽象。
3. 在驱动侧，适配器依赖于端口，由应用程序服务实现，因此适配器不知道谁会响应其调用，它只知道有哪些方法是保证可用的，因此它依赖于抽象。
4. 在被驱动侧，应用程序服务依赖于端口，而适配器则实现端口的接口，有效地反转了依赖关系，因为“低级”适配器（即数据库存储库）被迫实现应用程序核心中定义的抽象，这是“高级”的。

所以我们的项目目录会像这样：

![https://asdxz.oss-cn-beijing.aliyuncs.com/img/202312072113108.png](https://asdxz.oss-cn-beijing.aliyuncs.com/img/202312072113108.png)

## 例子

这里我们也能看出六边形架构的另外一个称呼：Ports and Adapters的原因，适配器实现端口（通常为接口），以达到代码解耦的作用，下面将以上面的目录进行具体的例子讲解：

完整代码：[link](https://github.com/Renewdxin/hex-arch.git)

本项目很简单，就是实现一个简单加减乘除的运算和数据库保存，那么我们秉持着核心domain层统领一切，适配器实现端口的原则，我们先定义 ./ports/core.go:

```go
package ports

type ArithmeticPort interface {
	Addition(a int32, b int32) (int32, error)
	Subtraction(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
	Division(a int32, b int32) (int32, error)
}
```

有了接口我们就得配以适配器 ./adapters/core/arithmetic/arithmetic.go：

```go
type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (Arith Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (Arith Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (Arith Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (Arith Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
```

这便是我们的核心逻辑，当项目慢慢变大时，核心层逻辑也会越来越多。

接下来就到了应用层，当我们实现了运算，那么便需要拿到结果，注意：此时还用不到sql，所以我们把目的写进 ./ports/app.go:

```go
type APIPort interface {
	GetAddition(a int32, b int32) (int32, error)
	GetSubtraction(a int32, b int32) (int32, error)
	GetMultiplication(a int32, b int32) (int32, error)
	GetDivision(a int32, b int32) (int32, error)
}
```

之后适配器实现：

```go
type Adapter struct {
	// depedencies injection
	arith ports.ArithmeticPort
	db    ports.DBPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{
		db:    db,
		arith: arith,
	}
}

func (api Adapter) GetAddition(a int32, b int32) (int32, error) {
	answer, err := api.arith.Addition(a, b)
	err = api.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (api Adapter) GetSubtraction(a int32, b int32) (int32, error) {
	answer, err := api.arith.Subtraction(a, b)
	err = api.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (api Adapter) GetMultiplication(a int32, b int32) (int32, error) {
	return api.arith.Multiplication(a, b)
}

func (api Adapter) GetDivision(a int32, b int32) (int32, error) {
	return api.arith.Division(a, b)
}
```

然后就到了用依赖的时候了，也就是framework，本文就讲讲mysql的CRUD:

```go
// internal/ports/framework_right.go
package ports

type DBPort interface {
	CloseDBConnection()
	AddToHistory(answer int32, operation string) error
}
```

然后实现：

```go
//internal/adapters/framework/right/db/db.go
package db

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect to db
	db, err := sql.Open(driverName, dataSourceName)
	...
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	stmt, err := da.db.Prepare("INSERT INTO history (data, answer, opration) VALUES (?,?,?)")
	...
}
```

之后我们编写测试文件，进行测试，通常情况一个适配器配一个测试文件

基本都创建好之后，如何连接呢？

我们在cmd文件中创建一个main.go：连接所有端口和适配器代码的地方，将依赖项注入到需要的层中，例如将数据库注入到framework层

这样实现了代码的解耦，例如我们想换一个数据库，只需要更换数据库名和数据源名，其余不需要修改，同时我们的业务逻辑也不需要了解特定的数据源限制

## 总结

所以总结一下优点：

1. 能够封装数据源实现细节
2. 长期稳定性和可扩展性，因为只需少许更改，所以在微服务部署失败时很容易回滚，也可以直接通过配置文件决定数据源

但是他也并不是silver bullet，我们应该多多检测层之间的漏洞，预防逻辑泄露等问题

## 参考文章&&学习视频

[Ready for changes with Hexagonal Architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)

[Hexagonal Architecture, there are always two sides to every story](https://medium.com/ssense-tech/hexagonal-architecture-there-are-always-two-sides-to-every-story-bc0780ed7d9c)

[How To Structure Your Go App - Full Course [ Hex Arch + Tests ]](https://youtu.be/MpFog2kZsHk?si=NYbFCbEPZ9stCi7x)