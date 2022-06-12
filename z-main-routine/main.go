package main

// 初始化顺序：变量初始化->init()->main()
func init() {
	println("execute init()")
}

func main() {
	println("Hello World")
}
