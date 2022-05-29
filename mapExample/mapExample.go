package mapexample

import "fmt"

func MapUsingMakeAndNewKeyword() {
	//map using make keyword
	mapUsingMake := make(map[int]int, 3)
	fmt.Println("mapUsingMake -->", len(mapUsingMake))
}
