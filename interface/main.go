package main

import "fmt"

type Shami int

type player interface {
	Stats()
}

func (s Shami) Stats() {
	fmt.Println("shami stats", s)
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
	fmt.Printf("stamina:%8v power:%8v centuries:%8v\n", v.stamina, v.power, v.centuries)
}

func (p *Player) Stats() {
	fmt.Printf("stamina:%8v power:%8v\n", p.stamina, p.power)
}

func main() {
	var sStats Shami = 1
	sStats.Stats()

	players := make([]Player, 11)
	for i := 0; i < len(players); i++ {
		players[i] = Player{
			stamina: 100 + i,
			power:   100,
		}

		players[i].Stats()

	}

	virat := &ViratPlayer{
		Player: Player{
			stamina: 100,
			power:   120,
		},
		centuries: 100,
	}

	virat.Stats()

}
