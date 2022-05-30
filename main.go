package main

import (
	_ "goAssignments/arraysAndSlicesExample"
	_ "goAssignments/hangmanGame"
	interfacesexample "goAssignments/interfacesExample"
	_ "goAssignments/mapExample"
	_ "goAssignments/pointersLearning"
	_ "goAssignments/receiverFunctionExample"
)

func main() {
	//pointerslearning.PointersBasics()
	//hangmangame.Hangman()
	//arraysandslicesexample.ArrayExample()
	//arraysandslicesexample.SliceExample()
	//mapexample.MapUsingMakeAndNewKeyword()
	s := interfacesexample.Square{20}
	c := interfacesexample.Circle{20}
	interfacesexample.PrintAreas(s)
	interfacesexample.PrintAreas(c)
	//receiverfunctionexample.Receiverfunctionexample()

}
