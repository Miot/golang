//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name string
	health, maxHealth int
	energy, maxEnergy int
}

func (p *Player) modifyHealth(amount int) {
	p.health += amount
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	} else if p.health < 0 {
		p.health = 0
	}
	fmt.Println(p.name, "health:", p.health)
}

func (p *Player) modifyEnergy(amount int) {
	p.energy += amount
	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	} else if p.energy < 0 {
		p.energy = 0
	}
	fmt.Println(p.name, "energy:", p.energy)
}

func main() {
	player := Player{name: "Player1", health: 10, maxHealth: 100, energy: 66, maxEnergy: 100}
	player.modifyHealth(100)
	player.modifyHealth(-10)
	player.modifyEnergy(50)
}

