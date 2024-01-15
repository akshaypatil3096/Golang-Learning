package main

import "fmt"

type Shami int

type player interface {
	Stats()
}

func (s Shami) Stats() {
	fmt.Printf("Shami took %d wickets\n", s)
}

type Player struct {
	stamina int
	power   int
}

type ViratPlayer struct {
	Player
	centuries int
}

func (v *ViratPlayer) Stats() {
	fmt.Printf("Kohli's stats are stamina:%8v power:%8v centuries:%8v\n", v.stamina, v.power, v.centuries)
}

func (p *Player) Stats() {
	fmt.Printf("stamina:%8v power:%8v\n", p.stamina, p.power)
}

func main() {

	players := make([]Player, 11)
	for i := 0; i < len(players); i++ {
		players[i] = Player{
			stamina: 100 + i,
			power:   100,
		}

		printDetails(&players[i])
	}

	virat := &ViratPlayer{
		Player: Player{
			stamina: 100,
			power:   120,
		},
		centuries: 100,
	}

	printDetails(virat)
	var sStats Shami = 24
	printDetails(sStats)

}

func printDetails(player player) {
	player.Stats()
}
