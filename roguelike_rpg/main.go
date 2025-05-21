package main

import (
	"fmt"
	"roguelike_rpg/game"
)

func main() {
	fmt.Println("====================================")
	fmt.Println(" SELAMAT DATANG DI PENJARA CROWNFALL ")
	fmt.Println("====================================")
	fmt.Print("Masukkan namamu, petualang: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("\nPilih kelas karaktermu:")
	fmt.Println("1. Ksatria")
	fmt.Println("2. Pencuri")
	fmt.Println("3. Penyihir")
	fmt.Print("Pilihan (1-3): ")
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
	fmt.Println("Kau terbangun dalam lorong tua yang remang-remang... Penjara Crownfall menantimu.")
	game.DisplayPlayerStatus(player)

	for {
		currentTile := game.WorldMap[player.Position]
		if currentTile == nil {
			fmt.Println("Kau melangkah ke puing-puing ruangan runtuh... Tidak ada apa pun di sini.")
		} else {
			if !currentTile.Visited {
				fmt.Println(currentTile.Description)
				currentTile.Visited = true
			}
			if currentTile.Enemy != nil {
				fmt.Printf("⚠️  Musuh muncul: %s!\n", currentTile.Enemy.Name)
				game.StartBattle(player, currentTile.Enemy)
				if player.HP <= 0 {
					fmt.Println("⚰️  Rohmu lenyap ditelan kegelapan... Tamat.")
					return
				}
				currentTile.Enemy = nil
				game.DisplayPlayerStatus(player)
			}
			if len(currentTile.Loot) > 0 {
				fmt.Println("🎁 Kamu menemukan:")
				for _, item := range currentTile.Loot {
					fmt.Printf("- %s: %s\n", item.Name, item.Description)
					player.Inventory = append(player.Inventory, item)
				}
				currentTile.Loot = nil
				game.DisplayPlayerStatus(player)
			}
		}

		fmt.Println("\nArah gerakmu? (w = atas, s = bawah, a = kiri, d = kanan, q = keluar): ")
		var move string
		fmt.Scanln(&move)

		switch move {
		case "w":
			player.Move(0, -1)
		case "s":
			player.Move(0, 1)
		case "a":
			player.Move(-1, 0)
		case "d":
			player.Move(1, 0)
		case "q":
			fmt.Println("Kau memilih meninggalkan penjara ini... untuk saat ini. Sampai jumpa.")
			return
		default:
			fmt.Println("Perintah tidak dikenali. Coba gunakan: w/s/a/d atau q.")
		}

		game.DisplayPlayerStatus(player)
	}
}

