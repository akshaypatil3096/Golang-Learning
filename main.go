package main

import (
	_ "goAssignments/arraysAndSlicesExample"
	gomodulesexample "goAssignments/goModulesExample"
	_ "goAssignments/goRoutinesExample"
	_ "goAssignments/hangmanGame"
	_ "goAssignments/interfacesExample"
	_ "goAssignments/mapExample"
	_ "goAssignments/receiverFunctionExample"
)

func main() {
	//pointerslearning.PointersBasics()
	//hangmangame.Hangman()
	//arraysandslicesexample.ArrayExample()
	//arraysandslicesexample.SliceExample()
	//mapexample.MapUsingMakeAndNewKeyword()
	/* s := interfacesexample.Square{20, "fas"}
	c := interfacesexample.Circle{20}
	interfacesexample.PrintAreas(s)
	interfacesexample.PrintAreas(c) */
	//receiverfunctionexample.Receiverfunctionexample()
	/* goroutinesexample.GoRoutines()
	goroutinesexample.Channels() */

	gomodulesexample.DescordModule()
	gomodulesexample.RandomData()
}
