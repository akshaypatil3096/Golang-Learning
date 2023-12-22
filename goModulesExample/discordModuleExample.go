package main

import (
	"fmt"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/bwmarrin/discordgo"
)

func main() {
	DescordModule()
}

// DescordModule function will generate the discord session object
func DescordModule() {
	discord, err := discordgo.New("")
	if err != nil {
		fmt.Println("Could not start the session")
	}
	fmt.Println("Session created: ", discord)
}

// RandomData generates the random data
func RandomData() {
	fmt.Println("Printing ramdom data ", randomdata.SillyName())
}
