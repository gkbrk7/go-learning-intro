package slices

import "fmt"

func Slice(){
	names := make([]string, 3)
	cities := []string{"Ankara","İstanbul","İzmir"}

	names[0] = "Gokberk"
	names[1] = "Guven"
	names[2] = "Gulcin"
	names = append(names, "Gizem")
	citiesCopy := make([]string, len(cities))
	copy(citiesCopy, cities)
	fmt.Println(names)
	fmt.Println(citiesCopy)
	fmt.Println(cities[0:2])
	fmt.Println(cities[0:])
	fmt.Println(cities[:2])
}