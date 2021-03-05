package maps_forrange

import "fmt"

func Map()  {
	// Like Dictionary, Hashset 
	// Key-Value Pair
	dictionary := make(map[string]string)
	dict := map[string]int{"test":1, "doe":2, "john":3}
	dictionary["table"] = "masa"
	dictionary["apple"] = "elma"
	dictionary["car"] = "araba"

	fmt.Println(dictionary["car"])
	delete(dictionary, "car")
	fmt.Println(dictionary)
	value, control := dictionary["apple"] 
	fmt.Println(dict)
	fmt.Println(value)
	fmt.Println("Exists : ", control)
}

func Range(){
	cities := map[string]int{"Ankara":1, "İstanbul":2, "İzmir":3}
	for _, city := range cities {
		fmt.Println(city)
	}
}