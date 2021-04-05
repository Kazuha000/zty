package main

func Receiver(v interface{})  {
	switch v.(type) {
	case string:
		println("this is string")
	case int:
		println("this is int")
	case bool:
		println("this is bool")
	case float64:
		println("this is float64")

	}
}
func main()  {
	Receiver("你好")
	Receiver(32)
	Receiver(0.5)
	Receiver(true)

}
