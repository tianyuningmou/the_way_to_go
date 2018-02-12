//当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil

package main

const PI = 3.1415926
const name string = "小明"

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Color int

const (
	RED Color = iota
	ORANGE
	YELLOW
	GREEN
	BLUE
)

func main() {
	println(a, b, c)
	println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	println(RED, ORANGE, YELLOW, GREEN, BLUE)
}
