package main

import "fmt"

type Player struct {
	health int
}


func (p *Player) damages(dmg int){
	fmt.Println("player damage")
	p.health -= dmg
}


func damage(p *Player){
	fmt.Println("player damage")
	dmg := 10
	p.health -= dmg
}

func main() {
	player := &Player{
		health: 100,
	}
	fmt.Printf("before damage %+v\n", player)
	player.damages(20)
	fmt.Printf("after damage %+v\n", player)
}
