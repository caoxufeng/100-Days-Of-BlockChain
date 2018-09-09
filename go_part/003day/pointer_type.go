package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in menory: %p\n", i1, &i1)
	//var p *int = nil
	//*p = 0 空指针的方向引用是不合法

	var l *int = &i1
	*l = 10
	fmt.Println(*l)
}
