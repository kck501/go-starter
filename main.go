package main

import "fmt"

func main() {
	//x := -1            //記住 :=(重點)
	if x := 2; x < 0 { //如果這邊if宣告為主
		fmt.Println("x 小於 0")

	} else {
		fmt.Println("x 大於或等於 0")
	}
	userAge := 20
	// 宣告判斷式
	if age := userAge; age >= 18 {
		//fmt.Printf("%dx%d=%d ", x, y, x*y) // 輸出：1x2=2
		fmt.Printf("你的年齡是%d,已滿18歲歡迎觀賞", age) //printf的使用方式
	}
	//fmt.Println(x) //如果x小於0
}
