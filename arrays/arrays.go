package arrays

import "fmt"

func Array(){
	var sayilar [5]int
	var multi [3][2]int
	sayilar1 := [5]int{5,4,3,2,1}
	for i := 0; i < 5; i++ {
		sayilar[i] = i
	}
	fmt.Println(sayilar)
	fmt.Println(sayilar1)
	fmt.Println(len(sayilar1))
	fmt.Println(multi)
}