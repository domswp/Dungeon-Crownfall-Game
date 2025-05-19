package main

import (
	"fmt"
	"roguelike_rpg/game"

)

func main() {
	fmt.Println("==========================")
	fmt.Println(" SELAMAT DATANG DI DUNGEON CROWNFALL ")
	fmt.Println("==========================")
	fmt.Println("Masukkan namamu, petualang: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("\nPilih kelas karaktermu:")
	fmt.Println("1. Ksatria")
	fmt.Println("2. Pencuri")
	fmt.Println("3. Penyihir")
	fmt.Println("Pilihlah (1-3): ")
	var choice int
	fmt.Scanln(&choice)

	class := ""
	switch choice {
	case 1:
		class = "Ksatria"
	case 2:
		class = "Pencuri"
	case 3:
		class = "Penyihir"
	default:
		class = "Petani"
	}

	player := game.NewPlayer(name, class)
	game.GenerateDungeon()

	fmt.Printf("\nSalam hormat, %s sang %s!\n", player.Name, player.Class)


