package main

import "fmt"

// 定义一个方法
type Integer int // 给int类型指定了一个别名叫Integer,别名可以随便起，只要符合GO语言的命名规则就可以

/*
func (a Integer) Test(b Integer) Integer{
}
表示定义了一个方法，方法的定义与函数的区别
第一：在关键字后面加上( a Integer), 这个在方法中称之为接收者，所谓的接受者就是接收传递过来的第一个参数，然后复制a， a的类型是Integer ,由于Integer是int的别名，所以a的类型为int
第二：在表示参数的类型时，都使用了对应的别名。
通过方法的定义，可以看出方法其实就是给某个类型绑定的函数。在该案例中，是为整型绑定的函数，只不过在给整型绑定函数(方法)时，一定要通过type来指定一个别名，因为int类型是系统已经规定好了，无法直接绑定函数，所以只能通过别名的方式。


第三:调用方式不同
var result Interger=3
表示定义一个整型变量result，并赋值为3.
result.Test( 3)
通过result变量，完成方法的调用。因为，Test( )方法，是为int类型绑定的函数，而result变量为int类型。所以可以调用Test( )方法。result变量的值会传递给Test( )方法的接受者，也就是参数a,而实参Test( 3),会传递形参b.
当然，我们也可以将Test( )方法，理解成是为int类型扩展了，追加了的方法。因为系统在int类型时，是没有改方法的。

*/

func (a Integer) Test(b Integer) Integer {
	return a + b
}

func main() {
	var result Integer = 3
	r := result.Test(3)
	fmt.Println(r)
}
