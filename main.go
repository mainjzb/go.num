package main

import (
	"fmt"

	"github.com/mainjzb/go.num/v2/zh"
)

func main() {
	/*
	var num zh.Uint64
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(uint64(num))

	 */

	var num2 zh.Uint64
	_, err2 := fmt.Sscan("十一", &num2)
	if err2 != nil{
		return
	}
	fmt.Println(num2)
}
