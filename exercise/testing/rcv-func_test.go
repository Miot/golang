//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
//
package main

import "testing"

func newPlayer() Player {
	return Player{
		name: "Player1", health: 100, maxHealth: 100, energy: 100, maxEnergy: 100,
	}
}
	
func TestHealth(t *testing.T) {
	player := newPlayer()
	player.modifyHealth(999)
	if player.health != 100 {
		t.Fatalf("Health went over limit: %v, want: %v", player.health, player.maxHealth)
	}
	player.modifyHealth(-999)
	if player.health != 0 {
		t.Fatalf("Health went under limit: %v, want: %v", player.health, 0)
	}
}

func TestEnergy(t *testing.T) {
	player := newPlayer()
	player.modifyEnergy(999)
	if player.energy != 100 {
		t.Fatalf("Energy went over limit: %v, want: %v", player.energy, player.maxEnergy)
	}
	player.modifyEnergy(-999)
	if player.energy != 0 {
		t.Fatalf("Energy went under limit: %v, want: %v", player.energy, 0)
	}
}	
