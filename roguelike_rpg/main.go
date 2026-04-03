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
        fmt.Println("Kau terbangun dalam lorong tua yang remang-remangan... Penjara Crownfall menantimu.")

	for {
	     currentTile := game.WorldMap[player.Position]
	     if currentTile == nil {
		     fmt.Println("Kau melangkah ke puing-puing ruangan.. Tidak ada apa pun disini.")
		     } else {
			     if !currentTile.Visited {
				     fmt.Println(currentTile.Description)
				     currentTile.Visited = true

			     }
			     if currentTile.Enemy != nil {
				     fmt.Printf("Seekor %s muncul dihadapanmu!\n", currentTile.Enemy.Name)
				     game.StartBattle(player, currentTile.Enemy)
				     if player.HP <= 0 {
					     fmt.Println("Rohmu perlahan menghilang ke dalam kegelapan.. Tamat")
					     return
				     }
				     fmt.Printf("\nKamu mendapatkan %d Exp!\n", currentTile.Enemy.ExpReward)
				     player.Exp += currentTile.Enemy.ExpReward
				     player.CheckLevelUp()
				     currentTile.Enemy = nil
			     }
		    if len(currentTile.Loot) > 0 {
			    fmt.Println("Kamu menemukan:")
			    for _, item := range currentTile.Loot {
				    fmt.Printf("- %s: %s\n", item.Name, item.Description)
				    player.Inventory = append(player.Inventory, item)
			    }
			    currentTile.Loot = nil
		    }

			if currentTile.IsPortal {
				fmt.Printf("\nKamu berdiri di atas Portal. Lanjutkan ke Lantai %d? (y/n): ", game.CurrentLevel+1)
				var ans string
				fmt.Scanln(&ans)
				if ans == "y" {
					game.CurrentLevel++
					player.Position = game.Position{X: 0, Y: 0}
					game.GenerateDungeon()
					fmt.Printf("\n--- SELAMAT DATANG DI LANTAI %d ---\n", game.CurrentLevel)
					continue
				}
			}

		   }

		   game.DrawMap(player)
		   fmt.Println("\nArah gerakanmu? (w = atas, s = bawah, a = kiri, d = kanan, q = keluar): ")
		   var move string
		   fmt.Scanln(&move)
		   switch move {
		   case "w":
			   if !player.Move(0, -1) {
				   fmt.Println("Sebuah dinding bata tebal menghalangimu!")
			   }
		   case "s":
			   if !player.Move(0, 1) {
				   fmt.Println("Sebuah dinding bata tebal menghalangimu!")
			   }
		   case "a":
			   if !player.Move(-1, 0) {
				   fmt.Println("Sebuah dinding bata tebal menghalangimu!")
			   }
		   case "d":
			   if !player.Move(1, 0) {
				   fmt.Println("Sebuah dinding bata tebal menghalangimu!")
			   }
		   case "q":
			   fmt.Println("Kau meninggalkan dungeon Crownfall. Sampai Jumpa, pahlawan.")
			   return
	    	  default:
			  fmt.Println("Perintah tidak dikenali.")
			  		  }
				  }
			  }






