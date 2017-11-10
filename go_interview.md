


 - ##### 下面的代码输出什么？
    ```go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	defer_call()
    }
    
    func defer_call() {
    	defer func() { fmt.Println("打印前1") }()
    	defer func() {
    		fmt.Println("打印中1.1")
    		recover()
    		fmt.Println("打印中1.2")
    	}()
    	defer func() { fmt.Println("打印中2") }()
    	defer func() { fmt.Println("打印后3") }()
    
    	panic("触发异常")
    }
    ```
    ```
    打印后3
    打印中2
    打印中1.1
    打印中1.2
    打印前1
    ```
> 注意recover后，程序还可以继续执行,注意defer函数的入栈和出栈顺序，先入后出。

 - ##### 输出什么，并说明原因

    ```
    package main
    
    import (
    	"fmt"
    	"sync"
    	"runtime"
    )
    
    func main() {
    	runtime.GOMAXPROCS(1)
    	wg := sync.WaitGroup{}
    	wg.Add(20)
    	for i := 0; i < 10; i++ {
    		go func() {
    			fmt.Println("iiii: ", i)
    			wg.Done()
    		}()
    	}
    	for i := 0; i < 10; i++ {
    		go func(i int) {
    			fmt.Println("j: ", i)
    			wg.Done()
    		}(i)
    	}
    	wg.Wait()
    }
    ```
    ```
    //GoMaxProcs(1)
    j:  9
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    j:  0
    j:  1
    j:  2
    j:  3
    j:  4
    j:  5
    j:  6
    j:  7
    j:  8
    ```
    ```
    //无GoMaxProc(1)
    iiii:  5
    iiii:  10
    j:  6
    j:  5
    j:  2
    j:  1
    j:  7
    iiii:  10
    j:  9
    j:  3
    j:  4
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    iiii:  10
    j:  0
    iiii:  10
    j:  8
    ```
> 解析：如果没有GoMaxProc是1的限制，i的输出并不一定会是10，考虑到for循环的执行速度可能会比开启goroutine快，而j的输出一定会是0-9，顺序可能会不同。

 - ##### 下面代码会输出什么？
    ```
    package main
    
    import (
    	"fmt"
    )
    
    type People struct{}
    
    func (p *People) ShowA() {
    	fmt.Println("showA")
    	p.ShowB()
    }
    func (p *People) ShowB() {
    	fmt.Println("showB")
    }
    
    type Teacher struct {
    	People
    }
    
    func (t *Teacher) ShowB() {
    	fmt.Println("teacher showB")
    }
    
    func main() {
    	t := Teacher{}
    	t.ShowA()
    }
    ```
    ```
    showA
    showB
    ```
 - #### 下面代码会触发异常吗？
    ```
    package main
    
    import (
    	"fmt"
    	"runtime"
    )
    
    func main() {
    	runtime.GOMAXPROCS(1)
    	int_chan := make(chan int, 1)
    	string_chan := make(chan string, 1)
    	int_chan <- 1
    	string_chan <- "hello"
    	select {
    	case value := <-int_chan:
    		fmt.Println(value)
    	case value := <-string_chan:
    		panic(value)
    	}
    }
    ```
> 可能会触发异常。
    ```
    select关键字用于多个channel的结合，这些channel会通过类似于are-you-ready polling的机制来工作。select中会有case代码块，用于发送或接收数据——不论通过<-操作符指定的发送还是接收操作准备好时，channel也就准备好了。在select中也可以有一个default代码块，其一直是准备好的。那么，在select中，哪一个代码块被执行的算法大致如下：
    检查每个case代码块:
    1. 如果任意一个case代码块准备好发送或接收，执行对应内容
    2. 如果多余一个case代码块准备好发送或接收，随机选取一个并执行对应内容
    3. 如果任何一个case代码块都没有准备好，等待
    4. 如果有default代码块，并且没有任何case代码块准备好，执行default代码块对应内容
    ```
    
 - #### 写出输出内容
    ```
    package main
    
    import (
    	"fmt"
    )
    
    func calc(index string, a, b int) int {
    	ret := a + b
    	fmt.Println(index, a, b, ret)
    	return ret
    }
    
    func main() {
    	a := 1
    	b := 2
    	defer calc("1", a, calc("10", a, b))
    	a = 0
    	defer calc("2", a, calc("20", a, b))
    	b = 1
    }
    ```
    ```
    10 1 2 3
    20 0 2 2
    2 0 2 2
    1 1 3 4
    ```
> 注意嵌套调用的问题，另外注意defer传参的问题

 - #### 输出内容是什么？
    ```
    package main
    
    import (
        "fmt"
    )
    
    func main() {
	    s := make([]int, 5)
	    s = append(s, 1, 2, 3)
	    fmt.Println(s)
    }
    ```
    ```
    [0 0 0 0 0 1 2 3]
    ```
    > slice 的初始化
    
 - #### 下面代码有什么问题？
    ```
    package main
    
    import  "sync"
    
    type UserAges struct {
    	ages map[string]int
    	sync.Mutex
    }
    
    func (ua *UserAges) Add(name string, age int) {
    	ua.Lock()
    	defer ua.Unlock()
    	ua.ages[name] = age
    }
    
    func (ua *UserAges) Get(name string) int {
    	if age, ok := ua.ages[name]; ok {
    		return age
    	}
    	return -1
    }
    ```
    > 改进后的版本
    
    ```
    package main
    
    import "sync"
    
    type UserAges struct {
    	ages map[string]int
    	sync.RWMutex
    }
    
    func NewUserAges() *UserAges {
    	return &UserAges{ages: make(map[string]int)}
    }
    
    func (ua *UserAges) Add(name string, age int) {
    	ua.Lock()
    	defer ua.Unlock()
    	ua.ages[name] = age
    }
    
    func (ua *UserAges) Get(name string) int {
    	ua.RLock()
    	defer ua.RUnlock()
    	if age, ok := ua.ages[name]; ok {
    		return age
    	}
    	return -1
    }
    
    func main() {
    	ua := NewUserAges()
    	ua.Add("li", 25)
    	ua.Get("li")
    }
    ```
 - 程序输出结果
    ```
    package main
    
    import (
    	"fmt"
    )
    
    type People interface {
    	Show()
    }
    
    type Student struct{}
    
    func (stu *Student) Show() {
    }
    
    func live() People {
    	var stu *Student
    	return stu
    }
    
    func main() {
    	if live() == nil {
    		fmt.Println("AAAAAAA")
    	} else {
    		fmt.Println("BBBBBBB")
    	}
    }
    ```
    ```
    BBBBBBB
    ```
    > 经典的nil值问题，如果为nil,则val nil，type nil