package main

// 初始化顺序：变量初始化->init()->main()
func init() {
	println("execute main init()")
}

func main() {
	//println("Hello World")
	app := &App{
		name:    "z-main-routine",
		version: "v1.0.0",
	}
	println(app.name, app.version)
}
